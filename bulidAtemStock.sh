#!/bin/bash
# deploy gobuild first
# docker run -d --name gobuild --restart=always \
#     -v /golang/packages:/go/pkg \
#     -v /golang/source:/go/src \
#     socloudng/gobuilder 
docker exec gobuild sh build.sh atem-stock main
cp application.yml dist -f
cp Dockerfile dist -f
cd ./dist
docker build -t socloudng/atem-stock .
# need docker login before push
# docker push socloudng/atem-stock
