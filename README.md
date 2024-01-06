The HiringSystem is a web server which is programmed by go language.

HiringSystem is a backend server for a recruitment system based on Go language and Vue, with frontend engineering to be released in another repository.

Gin and gorm are used in this project.

Build by go language.

Use mysql 8.0.34 Community.

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

Before you build this project,please instead the email count and email code in ../Config/EmailConfig.go file

You need to write your own email count and code here:

```
const (
	//input your email code here
	code  = ""
	email = ""
)
```

When you first execute this program,use this code:

```
func main() {
	DataBaseService.InitalDataBase()
	DataBaseService.CreateDataBase()
}
```

This code will help you create table in mysql.

And then,please replace the code with this:

```
func main() {
	DataBaseService.InitalDataBase()
	Utils.CreateDir("Resume")
	r := Router.Router()
	r.Run()
}
```

Here is the link to the front-end repository:

```
https://github.com/link-chen/HiringSystem-Front.git
```

If you want to join us,please send email to 1533842603@qq.com

If anyone wishes to sponsor us,here is the QR code:

![](./WeChat.png)
