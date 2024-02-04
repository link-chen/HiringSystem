package Service

import (
	"HiringSystem/DataBaseService"
	"HiringSystem/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"time"
)

func AddJob(c *gin.Context) {
	//传入Job，但可以通过Job.PostedBy反推HR
	var Job Utils.Job
	c.ShouldBindBodyWith(&Job, binding.JSON)
	fmt.Println(Job)
	if &Job != nil {
		ans := DataBaseService.AddJobToDataBase(Job)
		res, _ := createToken(Job.PostedBy, "HR")
		var rw []interface{}
		rw = append(rw, res)
		DataBaseService.SetKey(strconv.Itoa(int(Job.PostedBy)), res, 10*time.Minute)
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
	c.ShouldBindBodyWith(&Job, binding.JSON)
	fmt.Println(Job)
	ans := DataBaseService.DeleteJobFromDataBase(Job)
	res, _ := createToken(Job.PostedBy, "HR")
	var rw []interface{}
	rw = append(rw, res)
	DataBaseService.SetKey(strconv.Itoa(int(Job.PostedBy)), res, 10*time.Minute)
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
	c.ShouldBindBodyWith(&HRUser, binding.JSON)
	ans := DataBaseService.GetJobsPostedByHR(HRUser)
	res, _ := createToken(HRUser.HId, "HR")
	fmt.Println("HRUser.HId==", HRUser.HId)
	DataBaseService.SetKey(strconv.Itoa(int(HRUser.HId)), res, 10*time.Minute)
	var rw []interface{}
	rw = append(rw, res)
	if ans != nil {
		rw = append(rw, ans)
		c.JSON(http.StatusOK, Utils.Response{200, "Success", rw})
	}
}

func GetAppliersByJobId(c *gin.Context) {
	//传入User
	var User Utils.User
	var HRUser Utils.HRUser
	c.ShouldBindBodyWith(&User, binding.JSON)
	c.ShouldBindBodyWith(&HRUser, binding.JSON)
	fmt.Println(User)
	ans := DataBaseService.FindUsersByJobId(User)
	fmt.Println(ans)
	res, _ := createToken(HRUser.HId, "HR")
	DataBaseService.SetKey(strconv.Itoa(int(HRUser.HId)), res, 10*time.Minute)
	var rw []interface{}
	rw = append(rw, res)
	rw = append(rw, ans)
	c.JSON(http.StatusOK, Utils.Response{200, "Success", rw})
}

func GetApplyerResume(c *gin.Context) {
	//传入User
	var User Utils.User
	var HR Utils.HRUser
	c.ShouldBindBodyWith(&User, binding.JSON)
	fmt.Println(User)
	res, _ := createToken(HR.HId, "HR")
	var rw []interface{}
	rw = append(rw, res)
	DataBaseService.SetKey(strconv.Itoa(int(HR.HId)), res, 10*time.Minute)
	file := DataBaseService.GetResumeAddress(User)
	c.File(file)
}

func EmployeeApplyer(c *gin.Context) {
	//传入User
	var User Utils.User
	c.ShouldBindBodyWith(&User, binding.JSON)
	Email := DataBaseService.GetEmail(User)
	if Email != "" {
		DataBaseService.CleanUserApply(User)
		Utils.SendSuccessEmail(Email)
	}
	c.JSON(http.StatusOK, Utils.Response{200, "Success", "Success"})
}

func NotEmployeeApplyer(c *gin.Context) {
	//传入User
	var User Utils.User
	c.ShouldBindBodyWith(&User, binding.JSON)
	Email := DataBaseService.GetEmail(User)
	if Email != "" {
		Utils.SendFailedEmail(Email)
		DataBaseService.CleanUserApply(User)
	}
	c.JSON(http.StatusOK, Utils.Response{200, "Success", "Success"})
}

func HRLogin(c *gin.Context) {
	//传入HR
	var user Utils.HRUser
	c.BindJSON(&user)
	fmt.Println(user)
	ans := DataBaseService.HRLogin(user)
	res, _ := createToken(user.HId, "HR")
	var rw []interface{}
	rw = append(rw, res)
	DataBaseService.SetKey(strconv.Itoa(int(user.HId)), res, 10*time.Minute)
	if ans {
		//中间件，添加数据toekn
		rw = append(rw, "Success")
		fmt.Println("success")
		c.JSON(http.StatusOK, Utils.Response{200, "Success", rw})
	} else {
		rw = append(rw, "Failed")
		c.JSON(http.StatusOK, Utils.Response{200, "Success", rw})
	}
}
