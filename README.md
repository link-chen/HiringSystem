Build by go language

Use mysql 8.0.34 Community

Before you build this project,please instead the email count and email code in ../Config/EmailConfig.go file

You need to write your own email count and code here.

```
const (
	//input your email code here
	code  = ""
	email = ""
)
```

When you first execute this program,use this code.

```
func main() {
	DataBaseService.CreateDataBase()
}
```

This code will help you create table in mysql.

And then,please replace the code with this.

```
func main() {
	DataBaseService.InitalDataBase()
	Utils.CreateDir("Resume")
	r := Router.Router()
	r.Run()
}
```

