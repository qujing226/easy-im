package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"sync"
)

type Server struct {
	sync.RWMutex

	opt            *option
	authentication Authentication

	routes map[string]HandlerFunc
	addr   string
	patten string

	connToUser map[*Conn]string
	userToConn map[string]*Conn

	upGrader websocket.Upgrader
	logx.Logger
}

func NewServer(addr string, opts ...Options) *Server {
	opt := newOption(opts...)
	return &Server{
		addr: addr,

		opt:            &opt,
		patten:         opt.patten,
		authentication: opt.Authentication,

		routes:     make(map[string]HandlerFunc),
		connToUser: make(map[*Conn]string),
		userToConn: make(map[string]*Conn),

		upGrader: websocket.Upgrader{},
		Logger:   logx.WithContext(context.Background()),
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

		// 处理解析完成后的消息
		switch message.FrameType {
		case FramePing:
			s.Send(&Message{FrameType: FramePing}, conn)
		case FrameData:
			if handler, ok := s.routes[message.Method]; ok {
				handler(s, conn, &message)
			} else {
				s.Send(&Message{FrameType: FrameData, Data: "no found handler"}, conn)
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
