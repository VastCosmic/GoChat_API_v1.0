# GoChat Deploy Doc

## Intro

GoChat -- 基于ChatGPT API 的在线对话平台
(This is a release of my private repository, so the commit history is not available.)


后端接口文档：[GochatAPI-Docs (vastcosmic.cn)](https://site.vastcosmic.cn/p/gochatapi-docs/)

### Environment 

#### Back-end

Golang：1.20  &  Mysql：8.0.33

	github.com/gin-contrib/cors v1.4.0
	github.com/gin-gonic/gin v1.9.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	gorm.io/driver/mysql v1.4.7
	gorm.io/gorm v1.24.6
#### Front-end

Vue3 + element-ui

## 部署文档（Ubuntu 22.04 or 20.04）

### 安装并配置后端环境

Golang 1.20、Mysql 8.0

### 下载项目文件

```
git clone https://github.com/VastCosmic/GoChat.git
```

### 配置数据库

1. 使用项目文件中的数据库结构备份文件创建数据库`GoChat/Models/DB/gochat.sql`

2. 将数据库配置文件中的内容修改为要部署的数据库的对应信息`GoChat/Models/DB/dbConfig.json`

   示例：

   ```json
   {
     "username": "go",
     "password": "123456",
     "host": "127.0.0.1",
     "port": 33333,
     "dbname": "gochat",
     "timeout": "10s"
   }
   ```

### 运行后端

1. 编译文件`GoChat/Router/main.go`，请注意修改文件路径。

   ```shell
   go build Router/main.go
   ```

2. 运行编译出来的可执行文件，请注意修改文件路径。

   ```
   ./main
   ```

3. 后端测试
   1. 可以使用Postman等RESTful测试工具进行后端测试。
   2. 后端接口文档：[GochatAPI-Docs (vastcosmic.cn)](https://site.vastcosmic.cn/p/gochatapi-docs/)

### 部署前端

前端代码为前端开发者私有仓库，暂不公开。