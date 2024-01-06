package Service

import (
	"HiringSystem/DataBaseService"
	"HiringSystem/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddJob(c *gin.Context) {
	var Job Utils.Job
	c.BindJSON(&Job)
	fmt.Println(Job)
	if &Job != nil {
		ans := DataBaseService.AddJobToDataBase(Job)
		if ans {
			c.JSON(http.StatusOK, Utils.Response{http.StatusOK, "Success", "AddSuccess"})
		} else {
			c.JSON(http.StatusOK, Utils.Response{404, "Failed", "AddFailed"})
		}
	} else {
		c.JSON(404, Utils.Response{404, "Failed", "Can not get job message"})
	}
}

func DeleteJob(c *gin.Context) {
	var Job Utils.Job
	c.BindJSON(&Job)
	ans := DataBaseService.DeleteJobFromDataBase(Job)
	fmt.Println(ans)
	if !ans {
		c.JSON(http.StatusOK, Utils.Response{404, "Failed", "Failed To DeleteJob"})
	} else {
		c.JSON(http.StatusOK, Utils.Response{200, "Failed", "DeleteJob Success"})
	}
}

func GetPostedJobRequester(c *gin.Context) {
	var HRUser Utils.HRUser
	c.BindJSON(&HRUser)
	fmt.Println(HRUser)
}

func GetJobsPosted(c *gin.Context) {
	var HRUser Utils.HRUser
	c.BindJSON(&HRUser)
	ans := DataBaseService.GetJobsPostedByHR(HRUser)
	c.JSON(http.StatusOK, Utils.Response{200, "Success", ans})
}

func GetAppliersByJobId(c *gin.Context) {
	var User Utils.User
	c.BindJSON(&User)
	ans := DataBaseService.FindUsersByJobId(User)
	c.JSON(http.StatusOK, Utils.Response{200, "Success", ans})
}

func GetApplyerResume(c *gin.Context) {
	var User Utils.User
	c.BindJSON(&User)
	fmt.Println(User.Id)
	c.File(DataBaseService.GetResumeAddress(User))
}

func EmployeeApplyer(c *gin.Context) {
	var User Utils.User
	c.BindJSON(&User)
	Email := DataBaseService.GetEmail(User)
	if Email != "" {
		DataBaseService.CleanUserApply(User)
		Utils.SendSuccessEmail(Email)
	}
}

func NotEmployeeApplyer(c *gin.Context) {
	var User Utils.User
	c.BindJSON(&User)
	Email := DataBaseService.GetEmail(User)
	if Email != "" {
		Utils.SendFailedEmail(Email)
		DataBaseService.CleanUserApply(User)
	}
}

func HRLogin(c *gin.Context) {
	var user Utils.HRUser
	c.BindJSON(&user)
	ans := DataBaseService.HRLogin(user)
	if ans {
		c.JSON(http.StatusOK, Utils.Response{200, "Success", "Success"})
	} else {
		c.JSON(http.StatusOK, Utils.Response{200, "Success", "Failed"})
	}
}
