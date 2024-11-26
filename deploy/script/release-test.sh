#!/bin/bash

need_start_server_shell=(
  # "rpc.sh"
  "user-rpc-dev.sh"
  "social-rpc-dev.sh"
  "im-ws-test.sh"
  "im-rpc-test.sh"
  # "api.sh"
  "user-api-dev.sh"
  "social-api-dev.sh"
  "im-api-test.sh"

  # task
  "task-mq-test.sh"
)

for i in ${need_start_server_shell[*]}; do
  chmod +x $i
  ./$i
done

docker ps
docker exec -it etcd etcdctl get --prefix ""