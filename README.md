The HiringSystem is a web server which is programmed by go language.

HiringSystem is a backend server for a recruitment system based on Go language, with frontend engineering to be released in another repository.

The main part of the backend server is based on gin and gorm. Token component using JWT is under development.In order to get better performance,the Redis and RabbitMQ will be used in this project soon this year.

Build by go language.

Use mysql 8.0.34 Community.

## How to build

Before you start this project,please make sure that mysql and redis were installed in your environment.

You need to create a database named HiringSystem,we default the password for the root user to be 123456. So the  code of open mysql are shown like this:

```
const (
	dsn = "root:123456@tcp(localhost:3306)/HiringSystem?charset=utf8mb4&parseTime=True&loc=Local"
	//default database
)
var db *gorm.DB

func InitalDataBase() {
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if Db != nil {
		db = Db
	}
}
```

You can change the configeration of the database whenever you want,just need to make sure that you can open mysql.

We have provide the  go.mod and go.sum,you don't need to create it again.

Before you build this project,please instead the email count and email code in ./Config/EmailConfig.go file

You need to write your own email count and code here:

```
const (
	//input your email code here
	code  = ""
	email = ""
)
```

if nothing was wroten in EmailConfig.go,the project will work as usual but it can't send email when you want to hiring or employ someone.

## About the front-end

The front-end of this project is based on vue3.

Here is the link to the front-end repository:

```
https://github.com/link-chen/HiringSystem-Front.git
```

If you want to join us,please send email to 1533842603@qq.com

If anyone  who wishes to sponsor us,here is the QR code:

![](./WeChat.png)
