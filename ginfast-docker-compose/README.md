# gin-fast

#### 第一步
构建镜像(本地已有镜像的话则跳过此步)
```bash
sudo docker build -t ginfast:v1 .
```

#### 第二步
拷贝ginfast-docker-compose目录到任意想要存放的位置后
开启终端，进入当前目录，运行下面指令
```bash
docker-compose up -d
```

#### 说明
此项是为了将容器中的日志、配置文件映射到外部的数据卷，方便排查问题