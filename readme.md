# MyBlog System

使用Go语言结合 Gin 框架和 GORM 库开发一个个人博客系统的后端，实现博客文章的基本管理功能，包括文章的创建、读取、更新和删除（CRUD）操作，同时支持用户认证和简单的评论功能。

## 运行环境
- Go 1.24.5
- MySql 8
- Gin
- Gorm

## 依赖安装
### 安装Gorm-[官网地址](https://gorm.io/zh_CN/docs/)
```shell
   go get -u gorm.io/gorm
```
### 安装Gorm-Mysql 驱动
```shell
   go get -u gorm.io/driver/mysql
```
### 安装Gin - [官网地址](https://gin-gonic.com/zh-cn/docs/)
```shell
   go get -u github.com/gin-gonic/gin
```
### 安装godotenv [gitHub地址](https://github.com/joho/godotenv)
```shell
   go get -u github.com/joho/godotenv
```
### 安装jwt
```shell
   go get -u github.com/golang-jwt/jwt/v5
```
## 项目启动
```shell
 go run main.go
```