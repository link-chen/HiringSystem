The HiringSystem is a web server which is programmed by go language.

Gin and gorm are used in this project.

Build by go language.

Use mysql 8.0.34 Community.

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

The front-end will be released soon.

If you want to join us,please send email to 1533842603@qq.com
