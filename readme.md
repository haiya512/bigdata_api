# Beego Api

### beego restful api, jwt

## 流程说明
> 主要就两大块，一个是添加message，另外一个是返回message
> 返回message这里又分为两种方式，一种是发送pushover，一种是直接通过API请求，
> 设计一个API，可以直接把所有20分钟的新消息都请求出来，然后去过一遍，这个可以试试。

> 其余的都是一些细枝末节，查看消息的后台，方便编辑消息，或者内容


## 测试路由
### 注册
* POST /v1/user/reg

### 登录
* POST /v1/user/login

### 认证测试
* GET /v1/user/auth
