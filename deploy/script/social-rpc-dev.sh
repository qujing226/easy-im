#!/bin/bash
reso_addr="crpi-4ll7w1qk44676hc7.cn-hangzhou.personal.cr.aliyuncs.com/go-chat-im/social-rpc-dev"
tag="latest"
container_name="go-chat-im-social-rpc-dev"

# 停止容器
docker stop ${container_name} || true
# 删除容器
docker rm ${container_name} || true
# 检查镜像是否存在，如果存在则删除
if docker images | grep -q "${reso_addr}:${tag}"; then
  docker rmi ${reso_addr}:${tag}
fi
# 拉取最新的镜像
docker pull ${reso_addr}:${tag}
# 运行新的容器
docker run -d --name ${container_name} --network go-im_easy-chat  -p 10001:10001 ${reso_addr}:${tag}
