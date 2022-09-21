#### 这是一个基于go-zero编写的简单  流程引擎

目录介绍:
BIN 为开发测试

Coommon工具类

DATA docker缓存  内有含mysql包含workflow表

deploy为 docker内环境便于

workflow为  go-zero核心代码

makefile 为linux 编译

openapi.json  API接口  可导入 apifox  或者postman


如何运行   当前目录 docker-compose up -d  查看docker 是否正常

或者可通过  127.0.0.1:9001 进入portainer管理docker  账号为admin  密码12345678

docker 内含  mysql,redis,etcd,nginx,modd,portainer


本地工具连接mysql的话要先进入容器,给root设置下远程连接权限 如果未删除DATA文件可忽略已帮你开启

```shell
$ docker exec -it mysql mysql -uroot -p
##输入密码：PXDNKKK1234
$ use mysql;
$ update user set host='%' where user='root';
$ FLUSH PRIVILEGES;
```

#### 如何使用 

启动完成后  导入OpenApi 到 apifox  或者postman

如果是docker内 请使用  127.0.0.1:9005/api/1.0/ZeroWorkFlow
如果是编译启动  请使用  127.0.0.1:1004


/getToken   param 参数  uid nick-name tenant_id  获取token

然后进行后面的操作  后面操作需要吧token 加到head中 去