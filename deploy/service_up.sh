#!/bin/bash

#��ʼ��nats jetstream
./jetstream_tool --addrs "nats://iLmz8sCXjkTYuh@127.0.0.1:4222"
cd service
docker-compose down --rmi all
docker-compose up -d


