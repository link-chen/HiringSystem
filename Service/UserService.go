package Service

import (
	"HiringSystem/DataBaseService"
	"HiringSystem/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
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
	if ans {
		c.JSON(http.StatusOK, Utils.Response{200, "TrySuccess", "Success"})
	} else {
		c.JSON(http.StatusOK, Utils.Response{400, "TrySuccess", "Failed"})
	}
}

func ApplyJob(c *gin.Context) {
	var User Utils.User
	c.BindJSON(&User)
	exist := DataBaseService.FindUserJobId(User)
	if exist != 0 {
		c.JSON(http.StatusOK, Utils.Response{200, "ApplyFailed", "ApplyExist"})
		return
	}
	ResumeExist := DataBaseService.CheckResumeExist(User)
	if ResumeExist {
		ans := DataBaseService.UserApplyJob(User)
		if ans {
			c.JSON(http.StatusOK, Utils.Response{200, "ApplySuccess", "ApplySuccess"})
		} else {
			c.JSON(http.StatusOK, Utils.Response{200, "ApplyFailed", "ApplyFailed"})
		}
	} else {
		c.JSON(http.StatusOK, Utils.Response{200, "ApplyFailed", "NeedResume"})
	}
}

func FindAllJobs(c *gin.Context) {
	ans := DataBaseService.SearchAllJobs()
	c.JSON(http.StatusOK, Utils.Response{200, "Success", ans})
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
	ans := c.PostForm("id")
	id, _ := strconv.Atoi(ans)
	res := DataBaseService.AddResumeToUser(id, dst)
	if res {
		c.JSON(http.StatusOK, Utils.Response{200, "Success", "AddResumeSuccess"})
	} else {
		c.JSON(http.StatusOK, Utils.Response{200, "Success", "AddResumeFailed"})
	}
}

func SearchApplyedJob(c *gin.Context) {
	var User Utils.User
	c.BindJSON(&User)
	ans, err := DataBaseService.FindUserApplyJob(User)
	if err == nil {
		c.JSON(http.StatusOK, Utils.Response{200, "Ok", ans})
	} else {
		c.JSON(http.StatusOK, "Error")
	}
}
