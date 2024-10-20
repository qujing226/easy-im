#!/bin/bash

need_start_server_shell=(
  # "rpc.sh"
  "user-rpc-dev.sh"
  # "api.sh"
  "user-api-dev.sh"
)

for i in ${need_start_server_shell[*]}; do
  chmod +x $i
  ./$i
done

docker ps
docker exec -it etcd etcdctl get --prefix ""