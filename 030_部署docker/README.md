# 通过二进制部署docker

# 1. 安装docker 环境。配置docker desktop

> Docker daemon 配置文件
```json5
{
  //加速地址:Preferred Docker registry mirror
  "registry-mirrors": [
    "https://reg-mirror.qiniu.com",
    "https://registry.docker-cn.com",
    "http://hub-mirror.c.163.com",
    "https://3laho3y3.mirror.aliyuncs.com",
    "https://mirror.ccs.tencentyun.com"
  ],
  //Enable insecure registry communication
  "insecure-registries": ["http://harbor.yvjoy.com"],
  "debug": true,
  "experimental": false
}
```

# 2. 拉取镜像

# 3. 进入go项目进行交叉编译
```go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o xxx
```

# 4. 编写Dockerfile
 ```Dockerfile
 FROM busybox

MAINTAINER  Lucas "lucas@test.com"

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
WORKDIR /var/mnt/uwd/
ADD ./uwd ./uwd
ADD ./config/config.toml  ./config/config.toml
ADD ./public/ ./public/
RUN chmod 755 ./uwd
EXPOSE 55372
ENTRYPOINT ["./uwd"]
 ```

# 5. docker build 构建image
```json
docker build -t xxx:v1.0.0 .
```

# 6. docker run 运行容器，进行测试
```json
docker run --rm --name uwd-c -d -p 55372:55372 xxx:v1.0.0
```
# 7. docker push  提交image镜像
```json5
  docker login xxx.com
  输入用户名/密码
  docker tag  2e25d8496557 xxxxx.com/home/uwd-i:v1.0.0
  docker push xxxxx.com/home/uwd-i:v1.0.0
  
  // docker run -v挂载本地卷，挂载本地配置文件
  docker run -d -p 55:55372 --name uwdcs -v /e/zoo/socket/uwd/config/config.toml:/var/mnt/uwd/config/config.toml uwd-i:v2.0.0
```
