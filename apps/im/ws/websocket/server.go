package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	sync.RWMutex

	*threading.TaskRunner
	opt            *serverOption
	authentication Authentication

	routes map[string]HandlerFunc
	addr   string
	patten string

	connToUser map[*Conn]string
	userToConn map[string]*Conn

	upGrader websocket.Upgrader
	logx.Logger
}

func NewServer(addr string, opts ...ServerOption) *Server {
	opt := newOption(opts...)
	return &Server{
		addr: addr,

		opt:            &opt,
		patten:         opt.patten,
		authentication: opt.Authentication,

		routes:     make(map[string]HandlerFunc),
		connToUser: make(map[*Conn]string),
		userToConn: make(map[string]*Conn),

		upGrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		Logger:     logx.WithContext(context.Background()),
		TaskRunner: threading.NewTaskRunner(opt.concurrency),
	}
}

func (s *Server) ServerWs(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			s.Errorf("server handler ws recover err : %v", err)
		}
	}()

	conn := NewConn(s, w, r)
	if conn == nil {
		return
	}
	if !s.authentication.Auth(w, r) {
		err := s.Send(&Message{
			Data: fmt.Sprint("No access rights"),
		}, conn)
		if err != nil {
			s.Error("Send message err", err)
		}
		s.Info("authentication failed")
		return
	}

	// todo: 读取消息，完成请求，建立连接
	s.addConn(conn, r)
	go s.handleConn(conn)
}

func (s *Server) handleConn(conn *Conn) {
	ids := s.GetUsers(conn)
	conn.Uid = ids[0]

	// 处理任务
	go s.handlerWrite(conn)

	if s.isAck(nil) {
		go s.readAck(conn)
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			s.Errorf("websocket conn read message : err %v, userId : %v", err, "")
			s.Close(conn)
			return
		}

		var message Message
		err = json.Unmarshal(msg, &message)
		if err != nil {
			s.Errorf("websocket conn read message : err %v, userId : %v", err, "")
			s.Close(conn)
			return
		}
		// todo: 给客户端回复一个ack

		// 依据消息进行处理
		if s.isAck(&message) {
			s.Infof("conn message read ack msg %v", message)
			conn.appendMsgMq(&message)
		} else {
			conn.message <- &message
		}
	}
}

func (s *Server) isAck(message *Message) bool {
	if message == nil {
		return s.opt.ack != NoAck
	}
	return s.opt.ack != NoAck && message.FrameType != FrameNoAck
}

// 读取消息的确认
func (s *Server) readAck(conn *Conn) {
	for {
		select {
		case <-conn.done:
			s.Infof("close message ack uid %v", conn.Uid)
			return
		default:

		}
		// 从队列中读取新的消息
		conn.messageMu.Lock()
		if len(conn.readMessage) == 0 {
			conn.messageMu.Unlock()
			// 增加睡眠
			time.Sleep(100 * time.Microsecond)
			continue
		}
		//读取第一条消息
		message := conn.readMessage[0]

		// 判断ack的方式
		switch s.opt.ack {
		case OnlyAck:
			// 直接给客户端回复
			s.Send(&Message{
				Id:        message.Id,
				FrameType: FrameAck,
				AckSeq:    message.AckSeq + 1,
			}, conn)
			// 进行业务处理
			// 从队列中移除
			conn.readMessage = conn.readMessage[1:]
			conn.messageMu.Unlock()
			conn.message <- message
		case RigorAck:
			// 先回
			if message.AckSeq == 0 {
				// 还未确认
				conn.readMessage[0].AckSeq++
				conn.readMessage[0].ackTime = time.Now()
				s.Send(&Message{
					Id:        message.Id,
					FrameType: FrameAck,
					AckSeq:    message.AckSeq,
				}, conn)
				s.Infof("message ack RigorAck send mid %v,seq %v,time %v", message.Id,
					message.AckSeq, message.ackTime)
				conn.messageMu.Unlock()
				continue
			}
			// 再验证
			// 1.客户端返回结果，再一次确认
			// 得到客户端的序号
			msgSeq := conn.readMessageSeq[message.Id]
			if msgSeq.AckSeq > message.AckSeq {
				// 确认
				conn.readMessage = conn.readMessage[1:]
				conn.messageMu.Unlock()
				conn.message <- message
				s.Infof("message ack RigorAck success mid %v", message.Id)
				continue
			}
			// 2.客户端没有确认，考虑是否超过了ack的确认时间
			val := s.opt.ackTimeout - time.Since(message.ackTime)
			if !message.ackTime.IsZero() && val <= 0 {
				// 		2.2 超过，结束确认
				delete(conn.readMessageSeq, message.Id)
				conn.readMessage = conn.readMessage[1:]
				conn.messageMu.Unlock()
				continue
			}
			// 		2.1 未超过，重新发送
			s.Send(&Message{
				Id:        message.Id,
				FrameType: FrameAck,
				AckSeq:    message.AckSeq,
			}, conn)
			time.Sleep(300 * time.Microsecond)
			conn.messageMu.Unlock()
		}
	}
}

// 任务的处理
func (s *Server) handlerWrite(conn *Conn) {
	for {
		select {
		case <-conn.done:
			fmt.Println("close conn handler write")
			return
		case msg := <-conn.message:
			switch msg.FrameType {
			case FramePing:
				err := s.Send(&Message{FrameType: FramePing}, conn)
				if err != nil {
					s.Error("send message err", err)
				}
			case FrameData:
				if handler, ok := s.routes[msg.Method]; ok {
					handler(s, conn, msg)
				} else {
					err := s.Send(&Message{FrameType: FrameData, Data: "no found handler"}, conn)
					if err != nil {
						s.Error("send message err", err)
					}
				}
			}
			if s.isAck(msg) {
				conn.messageMu.Lock()
				delete(conn.readMessageSeq, msg.Id)
				conn.messageMu.Unlock()

			}
		}
	}
}

func (s *Server) addConn(conn *Conn, req *http.Request) {
	uid := s.authentication.UserId(req)

	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	// 验证用户是否之前登入过
	if c := s.userToConn[uid]; c != nil {
		// 关闭之前的连接
		_ = c.Close()
	}

	s.connToUser[conn] = uid
	s.userToConn[uid] = conn
}

func (s *Server) Close(conn *Conn) {

	s.RWMutex.Lock()
	defer s.Unlock()

	uid := s.connToUser[conn]
	if uid == "" {
		// 已经被关闭了
		return
	}
	delete(s.connToUser, conn)
	delete(s.userToConn, uid)

	_ = conn.Close()
}

func (s *Server) SendByUserId(msg any, sendIds ...string) error {
	if len(sendIds) == 0 {
		return nil
	}
	return s.Send(msg, s.GetConn(sendIds...)...)
}

func (s *Server) Send(msg any, conns ...*Conn) error {
	if len(conns) == 0 {
		return nil
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	for _, conn := range conns {
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) GetConn(uid ...string) []*Conn {
	s.RWMutex.RLock()
	defer s.RUnlock()
	conn := make([]*Conn, 0, len(uid))
	for _, id := range uid {
		conn = append(conn, s.userToConn[id])
	}
	return conn
}

func (s *Server) GetUsers(conns ...*Conn) []string {
	s.RLock()
	defer s.RUnlock()

	var res []string
	if len(conns) == 0 {
		res = make([]string, 0, len(s.connToUser))
		for _, conn := range s.connToUser {
			res = append(res, conn)
		}
	} else {
		res = make([]string, 0, len(conns))
		for _, conn := range conns {
			res = append(res, s.connToUser[conn])
		}
	}
	return res
}

func (s *Server) AddRoutes(rs []Route) {
	for _, r := range rs {
		s.routes[r.Method] = r.Handler
	}
}

func (s *Server) Start() {
	http.HandleFunc(s.opt.patten, s.ServerWs)
	_ = http.ListenAndServe(s.addr, nil)
}

func (s *Server) Stop() {
	fmt.Println("stop")
}
