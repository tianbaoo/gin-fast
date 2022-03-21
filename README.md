# gin-fast

#### 声明
Gin框架通用项目模板  
gin-fast is Free Software and released under the GNU Affero General Public License V3.  

在项目根目录下执行  
```bash
sudo swag init --parseDependency --parseInternal --parseDepth 1 -g main.go
```
在项目根目录下会生成docs目录(每次有新增接口注释后，需要重新执行一次)  
访问http://127.0.0.1:7890/swagger/index.html 即可看到接口界面  

#### 项目目录
    ├─conf              配置文件  
    ├─docs              文档  
    ├─handlers          接口   
    ├─middlewares       中间件  
    ├─models            模型    
    ├─pkg               自定义  
    │  ├─jwt            jwt    
    │  └─util           工具    
    ├─routers           路由  
    ├─serializers       序列化  
    ├─static            静态文件  
    │  ├─css  
    │  ├─img  
    │  └─js  
    ├─templates         模板  
    │ .gitignore        git
    │ go.mod            go mod  
    │ go.sum            go mod
    │ main.go           main入口
    │ README.md         

#### 运行
```bash
go mod tidy
go mod vendor
go run main.go
```

#### docker部署
```bash
sudo docker build -t ginfast:v1 .  # 创建镜像
sudo docker run -it -p 7890:7890 --rm ginfast:v1  # 运行镜像
sudo docker create -p 7890:7890 --name ginfast ginfast:v1  # 创建容器
sudo docker start ginfast  # 启动容器
sudo docker ps -a  # 查看容器ID
sudo docker exec -it 17c9a8cf12df /bin/sh  # 进入容器
curl http://127.0.0.1:7890/api/hello  # 测试接口
```
