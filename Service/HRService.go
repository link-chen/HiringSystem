package Service

import (
	"HiringSystem/DataBaseService"
	"HiringSystem/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddJob(c *gin.Context) {
	//传入Job，但可以通过Job.PostedBy反推HR
	var Job Utils.Job
	c.BindJSON(&Job)
	fmt.Println(Job)
	if &Job != nil {
		ans := DataBaseService.AddJobToDataBase(Job)
		res, _ := createToken(Job.PostedBy, "HR")
		var rw []interface{}
		rw = append(rw, res)
		if ans {
			rw = append(rw, "AddSuccess")
			c.JSON(http.StatusOK, Utils.Response{http.StatusOK, "Success", rw})
		} else {
			rw = append(rw, "AddAddFailed")
			c.JSON(http.StatusOK, Utils.Response{404, "Failed", rw})
		}
	} else {
		c.JSON(404, Utils.Response{404, "Failed", "Can not get job message"})
	}
}

func DeleteJob(c *gin.Context) {
	//传入Job，但可以通过Job.PostedBy反推HR
	var Job Utils.Job
	c.BindJSON(&Job)
	ans := DataBaseService.DeleteJobFromDataBase(Job)
	res, _ := createToken(Job.PostedBy, "HR")
	var rw []interface{}
	rw = append(rw, res)
	if !ans {
		rw = append(rw, "Failed To DeleteJob")
		c.JSON(http.StatusOK, Utils.Response{404, "Failed", rw})
	} else {
		rw = append(rw, "DeleteJob Success")
		c.JSON(http.StatusOK, Utils.Response{200, "Success", rw})
	}
}

func GetPostedJobRequester(c *gin.Context) {
	//传入HR
	var HRUser Utils.HRUser
	c.BindJSON(&HRUser)
	fmt.Println(HRUser)
}

func GetJobsPosted(c *gin.Context) {
	//传入HR
	var HRUser Utils.HRUser
	c.BindJSON(&HRUser)
	ans := DataBaseService.GetJobsPostedByHR(HRUser)
	c.JSON(http.StatusOK, Utils.Response{200, "Success", ans})
}

func GetAppliersByJobId(c *gin.Context) {
	//传入User
	var User Utils.User
	c.BindJSON(&User)
	ans := DataBaseService.FindUsersByJobId(User)
	c.JSON(http.StatusOK, Utils.Response{200, "Success", ans})
}

func GetApplyerResume(c *gin.Context) {
	//传入User
	var User Utils.User
	c.BindJSON(&User)
	fmt.Println(User.Id)
	//res, _ := createToken(123456, "HR")
	//var rw []interface{}
	//rw = append(rw, res)
	c.File(DataBaseService.GetResumeAddress(User))
}

func EmployeeApplyer(c *gin.Context) {
	//传入User
	var User Utils.User
	c.BindJSON(&User)
	Email := DataBaseService.GetEmail(User)
	if Email != "" {
		DataBaseService.CleanUserApply(User)
		Utils.SendSuccessEmail(Email)
	}
}

func NotEmployeeApplyer(c *gin.Context) {
	//传入User
	var User Utils.User
	c.BindJSON(&User)
	Email := DataBaseService.GetEmail(User)
	if Email != "" {
		Utils.SendFailedEmail(Email)
		DataBaseService.CleanUserApply(User)
	}
}

func HRLogin(c *gin.Context) {
	//传入HR
	var user Utils.HRUser
	c.BindJSON(&user)
	ans := DataBaseService.HRLogin(user)
	res, _ := createToken(user.Id, "HR")
	var rw []interface{}
	rw = append(rw, res)
	if ans {
		//中间件，添加数据toekn
		rw = append(rw, "Success")
		c.JSON(http.StatusOK, Utils.Response{200, "Success", rw})
	} else {
		rw = append(rw, "Failed")
		c.JSON(http.StatusOK, Utils.Response{200, "Success", rw})
	}
}
