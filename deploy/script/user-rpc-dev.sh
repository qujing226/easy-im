# !/bin/bash
reso_addr="crpi-4ll7w1qk44676hc7.cn-hangzhou.personal.cr.aliyuncs.com/go-chat-im/user-rpc-dev"
tag="latest"

container_name = "go-chat-im-user-rpc-dev"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}

docker run -d --name ${container_name} -p 10000:10000 ${reso_addr}:${tag}