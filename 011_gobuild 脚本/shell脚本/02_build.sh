#!/bin/bash
export LANG="en_US.UTF-8"

echo "开始编译vp-server项目... "
echo "----> 1: linux (64) ; "
echo "----> 2: windows (64) ; "
echo "----> 3: darwin (64) ; "

if read -p "======>请选择您要编译的环境(1/2/3) :" args
then
    echo "------> 你输入的环境是 $args <------"
else
    echo "\n抱歉，你输入超时了。"
fi

echo
projectName="video_server"

if [ $args -eq 1 ]; then
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64    go build -v -a -o ${projectName}_linux_amd64
elif [ $args -eq 2 ]; then
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64    go build -v -a -o ${projectName}_windows_amd64.exe
elif [ $args -eq 3 ]; then
  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64    go build -v -a -o ${projectName}_darwin_amd64
else
    echo "输入的参数$args不正确，执行退出..."
    exit 1
fi
echo

echo "...build success!!! 编译成功..."
