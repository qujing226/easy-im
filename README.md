# easy · im —— 下一代轻量化安全通讯APP
* 构建模块： [[构建]]、[[设计模式]]
* 安全模块： [[加密]]
* 部署模块： [[部署]] 、[[etcd]]
## Content:
- [[#一. 前言|一. 前言]]
    - [[#一. 前言#0 - 概述|0 - 概述]]
    - [[#一. 前言#1 - 技术栈分析|1 - 技术栈分析]]
- [[#二. 设计思路|二. 设计思路]]
    - [[#二. 设计思路#0 - 业务分析|0 - 业务分析]]
    - [[#二. 设计思路#1 - 表结构及关键字段|1 - 表结构及关键字段]]
    - [[#二. 设计思路#2 - 存储设计|2 - 存储设计]]
    - [[#二. 设计思路#3 - 通讯机制|3 - 通讯机制]]
    - [[#二. 设计思路#4 - 鉴权设计|4 - 鉴权设计]]
    - [[#二. 设计思路#5 - 心跳检测|5 - 心跳检测]]
    - [[#二. 设计思路#6 - ACK 序列机制|6 - ACK 序列机制]]
    - [[#二. 设计思路#7 - 安全设计|7 - 安全设计]]
    - [[#二. 设计思路#8 - 扩展与高可用设计|8 - 扩展与高可用设计]]
- [[#三. 关键元素|三. 关键元素]]
    - [[#三. 关键元素#1 - 高可用、热重载|1 - 高可用、热重载]]
    - [[#三. 关键元素#2 - 消息序列|2 - 消息序列]]
    - [[#三. 关键元素#3 - 长连接支持|3 - 长连接支持]]
    - [[#三. 关键元素#4 - 安全通讯|4 - 安全通讯]]
## 一. 前言
### 0 - 概述
> 本项目是微服务架构中的信息收发的关键组成，采用websocket、grpc、go-zero、etcd、mysql、redis、kafka、mongoDB、apisix、docker等实现，设置了特殊的Ack确认机制，并使用bitmap确保消息已读未读，具有快速、轻量、高可用、安全等特点。
> 1. 采用更加高效的消息已读未读算法降低在多用户群聊场景下系统的负载。
> 2. 在理论场景下，项目支持百万级别长连接以及高可用。
> 3. 实现可靠的用户心跳检测。
> 4. 参考TCP协议的三次握手实现通讯消息序列以及可靠传输。
> 5. 独特的“安全私聊”模块为用户提供了点对点的通讯服务，并设置了特殊的安全机制确保消息的真实性、同时采用抗量子算法加密伪实现信息的存储时效性。
> 6. 在安全通讯中设置了加密的TAG确保保障消息可被信任。

### 1 - 竞品分析
#### Telegram -- 安全的通讯应用
>为什么说电报通讯安全？
1. Tg 实现了用户间点对点进行通讯，信息不在服务端进行留存。
2. 以一种更加安全的加密方式使得信息即使被截取也很难破解。

> 本应用对比 Tg 有什么优势？
> 此应用
随着互联网的不断壮大，用户越来越注重自己的隐私、内容所有权是必要的发展趋势。

### 2 - 技术栈分析
#### websocket协议
WebSocket是一种基于tcp实现在Web应用中实现实时双向通信的协议。它通过在客户端和服务器之间建立持久连接，实现了服务器主动向客户端推送数据的能力
##### 特点：
1. websocket可以在浏览器里使用
2. 支持双向通信
3. 使用很简单
##### 为什么使用websocket
当一个请求需要长时间处理时，网关可能会返回超时信息，对于处理过程我们可以交给异步任务，可以在处理完成时服务端主动通知客户端，这就需要websocket。
注意：虽然HTTP/2也具备服务器推送功能，但HTTP/2 只能推送静态资源，无法推送指定的信息。

#####  适合场景
1. 存在及时通信功能需求
2. 对数据变化要求实时性
3. 用于用户消息推送通知比如异步任务结果的通知
## 二. 设计思路
### 0 - 业务分析
#### 主要业务

在im服务中我们需要提供的内容有
1. 对外需考虑提供对消息记录的查询接口
2. 对内需要考虑其他模块的rpc调度推送消息
3. 还需要考虑用户的websocket连接服务

为实现通讯服务的基本功能，可分为如下业务
1. 私聊
2. 群聊
3. 消息已读未读
4. 用户在线离线
5. 历史与离线消息
### 1 - 表结构及关键字段
#### 业务关键表
1. 用户表
2. 社交表
3. 群聊表
#### 聊天关键表
1. 聊天记录表
2. 用户会话列表

```Bash
goctl model mongo --type chatLog --dir ./apps/im/models/
```
#### 关键字段


##### 会话id
聊天记录的查询分为发起人和接收人即send_id与recv_id，针对好友之间的消息查询在实现的方案上采通常采用会话id来查询。

> 基于sendId与recvId进行计算后生成一个特定的值作为id。

### 2 - 存储设计
基于客户端缓存+服务端信息存储
- 方式：本地先缓存当前的消息，基于消息时序号记录消息的起始节点，在需要拉取历史消息时根据时序号分页拉取消息
- 优点：解决跨机器切换账号后无法获取消息记录问题，同时也可以实现消息的已读未读功能
- 缺点：服务器需要设置好相应的存储记录的方式
- 适合：对消息具有记录要求的场景以办公类型为主的业务则适合
##### 数据库
存储引擎的选择主要集中在是选择mysql还是mongodb。
- mysql：是关系型数据库，适合处理结构化的数据，有良好的事务处理能力，对数据聚合统计有较好的处理，但在处理大量数据时可能会受到性能瓶颈的限制，分库分表的处理会较为麻烦。
- mongodb：是一种面向文档的NoSQL数据库，适合存储非结构化或半结构化的数据。在处理较大数据量的时候具有更好的可扩展性和高可用性。

### 3 - 通讯机制
#### a. 消息转发
通过websoclet服务端+客户端进行消息转发。
![[Pasted image 20241203142809.png]]
>token的传输验证主要是在用户端发起的第一次http请求的时候进行验证的，在后续是构建好了长连接，此时就不需要再重复的发送token进行鉴权。
#### b. 消息队列



### 4 - 鉴权设计
#### a. 构思 Authorization 接口

>目前已知用户是通过用户服务获取到登入授权的token，而用户服务是基于jwt生成的token，在api服务的验证中是直接通过go-zero的方式进行验证的。因此在自定义的验证机制对象中只需要依据go-zero的实现方式即可做好验证。

>接口的主要功能：
1. 提供验证 token 真伪的功能。
2. 提供根据 token 信息提供 User 身份的功能。
```Go
type Authentication interface {  
    Auth(w http.ResponseWriter, r *http.Request) bool  
    UserId(r *http.Request) string  
}
```
#### b. go-zero 对于 jwt 的支持
在go-zero中因我们在api.api文件中定义了服务server并在其中添加jwt:JwtAuth的操作，在通过命令执行的时候会默认给生成的相关路由添加好jwt的认证中间件。
![[Pasted image 20241203145956.png]]
![[Pasted image 20241203150116.png]]
#### c. 实现思路
先通过token包创建好jwt的auth解析器对象tokenParser，实践的jwt认证是通过调用parseToken方法完成，到此我们基本就以确定了认证方式，而具体的实现则依据authhandler中的定义实现完成即可。
核心代码：
1. 解析 token
2. 翻译成用户 Claim
   ![[Pasted image 20241203150439.png]]
   在Auth方法中主要对用户鉴权，鉴权完成通过tok.Claims获取到用户的id信息，再将id信息存储到request的context中实现上，而在UserId方法中从request.Context中获取到用户的id信息。

```Go
package handler  
  
import (  
    "context"  
    "easy-chat/apps/im/ws/internal/svc"    "easy-chat/pkg/ctxdata"    "github.com/golang-jwt/jwt/v4"    "github.com/zeromicro/go-zero/core/logx"    "github.com/zeromicro/go-zero/rest/token"    "net/http")  
  
type JwtAuth struct {  
    svc    *svc.ServiceContext  
    parser *token.TokenParser  
    logx.Logger  
}  
  
func NewJwtAuth(svc *svc.ServiceContext) *JwtAuth {  
    return &JwtAuth{  
       svc:    svc,  
       parser: token.NewTokenParser(),  
       Logger: logx.WithContext(context.Background()),  
    }}  
func (j *JwtAuth) Auth(w http.ResponseWriter, r *http.Request) bool {  
    if tok := r.Header.Get("sec-websocket-protocol"); tok != "" {  
       r.Header.Set("Authorization", tok)  
    }    tok, err := j.parser.ParseToken(r, j.svc.Config.JwtAuth.AccessSecret, "")  
    if err != nil {  
       return false  
  
    }  
    if !tok.Valid {  
       return false  
    }  
    claims, ok := tok.Claims.(jwt.MapClaims)  
    if !ok {  
       return false  
    }  
  
    *r = *r.WithContext(context.WithValue(r.Context(), ctxdata.IdentityKey, claims[ctxdata.IdentityKey]))  
    return true  
}  
  
func (j *JwtAuth) UserId(r *http.Request) string {  
    return ctxdata.GetUId(r.Context())  
}
```
### 5 - 心跳检测
>实时的聊天过程中，在没有心跳检测的情况下会出现如下问题
>  1. websocket连接后，长时间服务端和客户端没有通讯，服务端会把websocket给断开
>  2. 在网络连接中，连接会因多种因素而导致不可靠，可能是客户端因意外非正常停止运行，也可能是服务端断开连接。
>
>心跳检测的主要作用是如下两个功能点。
>  - 保活
>  - 检测死链
#### grpc 心跳实现分析
在grpc中服务端也有对客户端的连接做心跳的检测,如下：
![[Pasted image 20241205111956.png]]
在 `keepalive()`核心代码中主要有三个核心定时器和三个关键属性值
- 定时器：
    - idleTimer: 检测连接空闲时间定时器
    - ageTimer：连接有效时长连接验证定时器
    - kpTimer：定时检测客户端是否发送连接信息
- 关键属性：
    - idle：记录空闲时间，当有连接进入会将期设置为非0值表示为非空闲状态
    - kq：配置信息记录一个连接最大有效连接时长，以及最大空闲时长等信息
    - lastRead: 记录最后一次通信的时间
##### 1. `idle` 定时器

`idle` 定时器用来检测连接是否在一定时间内处于空闲状态。如果在指定的 `idle` 时间内没有任何活动（例如，数据传输或心跳信号），连接将被视为空闲并可能被关闭，以释放资源。

##### 2. `kp`（Keepalive Ping）定时器

`kp` 定时器控制心跳消息（Keepalive Ping）的发送频率。定期发送心跳消息有助于确保连接的活跃状态，并在网络连接中断时及时发现问题。如果在指定的时间间隔内没有从对端接收到响应，Keepalive Ping 机制可以检测到连接问题。

##### 3. `age` 定时器

`age` 定时器用于检测连接的总寿命。此定时器会在连接达到最大寿命时触发，强制断开并重新建立连接。这有助于防止长期连接可能导致的资源泄漏和性能问题。

#### 工作原理

1. `idle` **定时器**：当连接在指定的 `idle` 时间内没有活动时，触发空闲处理逻辑。例如，关闭连接以节省资源。

2. `kp` **定时器**：定期发送心跳消息，并等待对端的响应。如果在指定时间内没有接收到响应，视为连接断开，并触发重连机制。

3. `age` **定时器**：当连接的寿命达到指定的时间时，强制关闭并重新建立连接，确保系统稳定性和性能。

这三个定时器相互配合，确保 gRPC 连接的稳定性、可靠性和性能。
#### idle定时器：
![[Pasted image 20241205112235.png]]
在流程中会先验证当前是否处于idle状态，若是不在空闲状态则重置定时器进行下一次检测时间，如果存在则往后计算空闲时差是否超过最大约定时间，如果是则发送通知给客户端进行关闭连接。
#### age定时器
![[Pasted image 20241205113213.png]]
当触发ageTimer定时器的时候，它不是立即关闭连接的，而是先向客户端发送关闭连接的消息通知并设置一个延迟关闭的时间，然后再通过选择器等待客户端或者延时关闭时间到自己关闭。

这样的目的是避免直接关闭导致在过程中可能存在任务处理失败的问题。

#### kptime定时器
该定时器主要检测一个连接在一定时间范围内是否有通信
![[Pasted image 20241205113313.png]]
当定时器触发时，先获取最后一次通信的时间，并验证是否已经有新的消息，当存在有消息通信时就会重置定时器进行下一次的检测，如果不存在则会往后执行。

在后续的流程中会期望给客户端发送一个ping的信息，用于验证客户端是否存活，如果存活就应会返回一个ack的确认信息此时会通过上一步的流程中发送的outstandingPing设置为false，如果客户端未发送就端口

即后续的流程为：先验证是否发送过ping信息并且是否超过了最大的等待时间，如果是则关闭连接，如果前面的条件不符合则依据后续发送ping消息，并且重置定时器进行下一次的检测。
#### 实现流程
![[Pasted image 20241205113446.png]]
如上就是实现的流程：在实现的本质就是双方会定时去检测连接是否存在通信，如果不存在则依据响应的规则发送心跳包或者断开连接。

1. 客户端向服务端建立连接
2. 客户端向服务端发送消息，服务的接收到消息后更新相互之间最后一次通信时间
3. 基于定时检测如果当前连接超过多久未通信
    1. 客户端：会发生心跳包，如果有回复则连接正常，如果不存在回复则连接异常重新建立连接
    2. 服务的：会视为客户端异常断开连接。

注意：整个流程中服务的的检测时间间隔应大于客户端的间隔时间，如果服务端的时间间隔小于客户端的时间间隔则会导致频繁的连接建立问题。
### 6 - ACK 序列机制

### 7 - 安全设计

### 8 - 扩展与高可用设计
#### 如何实现配置中心

#### 如何实现热重载
##### 程序中的重启问题

当配置变更后程序理应考虑需重启加载最新配置但是会存在如下问题

1. 中断正在进行的操作：如果在重启期间有用户或客户端正在使用服务器提供的服务，重启将导致操作中断，使这些用户或客户端无法完成它们正在进行的操作，会给他们带来极大不便。
2. 数据丢失：未保存的数据可能会在服务器重启后消失，特别是在没有进行优雅关闭连接的情况下。
3. 系统过载：在服务器启动时，可能需要短时间内进行大量的资源加载和初始化操作，这会导致系统过载和性能下降。
4. 可用性降低：因服务中断导致的可用性降低，可能会让用户对系统的信任度降低，甚至影响到企业的声誉。

另外就是通过如下过程加载配置

1. kill -9
2. 启动服务

这也还是会存在问题

- 未处理完的请求，被迫中断，数据一致性被破坏
- 新服务启动期间，请求无法进来，导致一段时间的服务不可用现象
#####  go-zero中api服务的平滑启动

go-zero中的api服务是基于go标准库中的http实现运行的，因此我们需先看go标准库中的实现。

对 http 服务来说，一般的思路就是关闭对 `fd` 的 `listen` , 确保不会有新的请求进来的情况下处理完已经进入的请求，然后退出。
![[Pasted image 20241204110532.png]]
在处理监听中，`inShutdown` 是一个原子变量，非 0 表示被关闭，那么此时就会停止任务进入到服务中。在过程中会通过closeIdleConns将目前 `Server` 中记录的活跃链接变成变成空闲状态，返回，最后再关闭服务。

**`Shutdown` 可以优雅的终止服务，期间不会中断已经活跃的链接**。

设计关键：

1. 在服务中记录请求数
2. 设置一个参数用于控制是否接收请求
3. 提供一个平滑方法用于重新加载配置属性参数
## 三. 关键元素
### 1 - 高可用、热重载
#### 构建im服务
在im服务中我们需要提供的内容有
1. 对外需考虑提供对消息记录的查询接口
2. 对内需要考虑其他模块的rpc调度推送消息
3. 还需要考虑用户的websocket连接服务
#### 项目结构：
![[Pasted image 20241121104156.png]]


### 2 - 消息序列



### 3 - 长连接支持


### 4 - 安全通讯


