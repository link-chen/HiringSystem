package Service

import (
	"HiringSystem/DataBaseService"
	"HiringSystem/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Regist(c *gin.Context) {
	var User Utils.User
	c.BindJSON(&User)
	ans, count := DataBaseService.AddSimpleUser(User)
	if ans {
		c.JSON(http.StatusOK, Utils.Response{200, "CreateSuccess", count})
	} else {
		c.JSON(http.StatusOK, Utils.Response{400, "CreateSuccess", 0})
	}
}

func Login(c *gin.Context) {
	var User Utils.User
	c.BindJSON(&User)
	ans := DataBaseService.SimpleUserLogin(User)
	res, _ := createToken(User.UId, "User")
	var rw []interface{}
	rw = append(rw, res)
	DataBaseService.SetKey(strconv.Itoa(int(User.UId)), res, 10*time.Minute)
	if ans {
		rw = append(rw, "Success")
		c.JSON(http.StatusOK, Utils.Response{200, "Success", rw})
	} else {
		rw = append(rw, "Failed")
		c.JSON(http.StatusOK, Utils.Response{400, "TrySuccess", rw})
	}
}

func ApplyJob(c *gin.Context) {
	var User Utils.User
	c.BindJSON(&User)
	fmt.Println("User.Uid==", User.UId)
	exist := DataBaseService.FindUserJobId(User)
	if exist != 0 {
		c.JSON(http.StatusOK, Utils.Response{200, "ApplyFailed", "ApplyExist"})
		fmt.Println("applyed")
		return
	}
	ResumeExist := DataBaseService.CheckResumeExist(User)
	res, _ := createToken(User.UId, "User")
	var rw []interface{}
	rw = append(rw, res)
	if ResumeExist {
		ans := DataBaseService.UserApplyJob(User)
		if ans {
			rw = append(rw, "ApplySuccess")
			c.JSON(http.StatusOK, Utils.Response{200, "ApplySuccess", rw})
		} else {
			rw = append(rw, "ApplyFailed")
			c.JSON(http.StatusOK, Utils.Response{200, "ApplyFailed", rw})
		}
	} else {
		rw = append(rw, "NeedResume")
		c.JSON(http.StatusOK, Utils.Response{200, "ApplyFailed", rw})
	}
}

func FindAllJobs(c *gin.Context) {
	ans := DataBaseService.SearchAllJobs()
	var User Utils.User
	c.BindJSON(&User)
	fmt.Println("User.id==", User.UId)
	res, _ := createToken(User.UId, "User")
	var rw []interface{}
	rw = append(rw, res)
	rw = append(rw, ans)
	fmt.Println("1111")
	c.JSON(http.StatusOK, Utils.Response{200, "Success", rw})
}

func AddResume(c *gin.Context) {
	file, _ := c.FormFile("file")
	fmt.Println(file.Filename)
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	dst := fmt.Sprintf("%s/%s", currentDir+"/Resume", file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		return
	}
	ans := c.PostForm("uid")
	id, _ := strconv.Atoi(ans)
	res := DataBaseService.AddResumeToUser(id, dst)
	token, _ := createToken(uint(id), "User")
	var rw []interface{}
	rw = append(rw, token)
	if res {
		rw = append(rw, "AddResumeSuccess")
		c.JSON(http.StatusOK, Utils.Response{200, "Success", rw})
	} else {
		rw = append(rw, "AddResumeFailed")
		c.JSON(http.StatusOK, Utils.Response{200, "Success", rw})
	}
}

func SearchApplyedJob(c *gin.Context) {
	var User Utils.User
	c.BindJSON(&User)
	fmt.Println("User.Uid==", User.UId)
	ans, err := DataBaseService.FindUserApplyJob(User)
	res, _ := createToken(User.UId, "User")
	var rw []interface{}
	rw = append(rw, res)
	fmt.Println(res)
	fmt.Println("SearchApplyedJobs")
	DataBaseService.SetKey(strconv.Itoa(int(User.UId)), res, 10*time.Minute)
	if err == nil {
		rw = append(rw, ans)
		c.JSON(http.StatusOK, Utils.Response{200, "Ok", rw})
	} else {
		c.JSON(http.StatusOK, "Error")
	}
}
