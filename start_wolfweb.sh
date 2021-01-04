#!/bin/bash

NAME="wolfweb"
CID=$(docker ps -a | awk '{if($NF=="'${NAME}'")print $1}')
if [[ $CID"x" != "x" ]];then
    docker rm --force $CID
fi

docker run -d --restart always \
--network host \
-v /:/hostfs:ro \
-e HOST_ETC=/hostfs/etc \
-e HOST_PROC=/hostfs/proc \
-e HOST_SYS=/hostfs/sys \
-e HOST_VAR=/hostfs/var \
-e HOST_RUN=/hostfs/run \
--name $NAME \
dockerhub.deepglint.com/deepface/wolfweb:0.0.1
