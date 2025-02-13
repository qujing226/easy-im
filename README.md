# go-im

#### 介绍
使用go语言创建的即时通讯项目

#### 软件架构
``` bash
./
├── apps
│   └── user
│       ├── api
│       ├── exec.sh
│       └── rpc
│           ├── etc
│           │   └── dev
│           │       └── user.yaml
│           ├── internal
│           │   ├── config
│           │   │   └── config.go
│           │   ├── logic
│           │   │   ├── finduserlogic.go
│           │   │   ├── getuserinfologic.go
│           │   │   ├── loginlogic.go
│           │   │   ├── pinglogic.go
│           │   │   └── registerlogic.go
│           │   ├── server
│           │   │   └── userserver.go
│           │   └── svc
│           │       └── servicecontext.go
│           ├── user
│           │   ├── user_grpc.pb.go
│           │   └── user.pb.go
│           ├── userclient
│           │   └── user.go
│           ├── user.go
│           └── user.proto
├── bin
│   └── user-rpc
├── deploy
│   ├── dockerfile
│   │   └── Dockerfile_user_rpc_dev
│   ├── mk
│   │   └── user-rpc.mk
│   └── script
│       ├── release-test.sh
│       └── user-rpc-dev.sh
├── docker-compose.yaml
├── go.mod
├── go.sum
├── Makefile
├── pkg
├── README.md
└── structure.txt


```
#### 测试用例
``` text
nickname    Qwyk        peninsula   admin       xiaoming    xiaohong    xiaowang
phone       17309710356 18239655309 17344995006 11122223333 22233334444 33344445555
password    yining2024  yujie2024   admin       123456      123456      123456

ID 
admin  1843306294396588032
xiaming 1843306302982328320

```

#### 群测试
``` 
1843311241150337024   Go语言开发者(4)
1843306294396588032     admin  (群主)
1843306302982328320     xiaoming
1843306311148638208     xiaohong
1843306319776321536     xiaowang
```
#### 安装教程

1.  sudo make user-rpc-dev
2.  sudo make install-server
3.  sudo docker compose up -d

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx



#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
