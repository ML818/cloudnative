#!/bin/bash


# official public docker repositories address: https://hub.docker.com/u/hisoka131


# sudo docker build -t mask2live.goserver:1.0 .
# sudo docker tag mask2live.goserver:1.0 hisoka131/mask2live.goserver:1.0
# sudo docker push hisoka131/mask2live.goserver:1.0

sudo docker run -d --rm hisoka131/mask2live.goserver:1.0

pid=$(sudo lsns -t net | grep server | awk 'FNR == 1 {print $4}')

echo
echo "============namespace info of Pid=$pid============"

sudo ls -al /proc/$pid/ns

sudo nsenter -t $pid -n ip a
