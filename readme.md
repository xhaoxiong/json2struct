#### 这是基于三方API写的一个简要的接口

<a href="https://json2struct.xhxblog.cn">项目地址</a>

```
项目支持go1.11及其以上
主要包括json和yaml格式的字符串转golang结构体

项目主要使用react+iris 适合新手入门

项目运行：由于已经将前端react打包了所以直接克隆之后配置好数据库可以运行

1、在conf文件中将数据库账号密码及其数据库弄好 

 mysql:
   username: "root"
   password: "123456(自定义)"
   addr: "127.0.0.1:3306"
   name: "json2struct" //使用utf8_mb4

2、go build

3、./json2struct

```