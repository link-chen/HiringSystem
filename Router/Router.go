package Router

import (
	"HiringSystem/Service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")

		if c.Request.Method == "OPTIONS" {
			// 对 OPTIONS 请求进行特殊处理或者直接返回 200 状态码
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	HRUser := r.Group("/HRService")
	{
		HRUser.POST("/AddJob", Service.AddJob)
		HRUser.POST("/DeleteJob", Service.DeleteJob)
		HRUser.POST("/GetPostedJobs", Service.GetJobsPosted)
		HRUser.POST("/GetRequester", Service.GetPostedJobRequester)
		HRUser.POST("/GetUsersById", Service.GetAppliersByJobId)
		HRUser.POST("/GetResumeById", Service.GetApplyerResume)
		HRUser.POST("/SelectUser", Service.EmployeeApplyer)
		HRUser.POST("/DeleteUser", Service.NotEmployeeApplyer)
		HRUser.POST("/HRLogin", Service.HRLogin)
	}

	User := r.Group("/UserService")
	{
		User.POST("/Regist", Service.Regist)
		User.POST("/Login", Service.Login)
		User.POST("/ApplyJob", Service.ApplyJob)
		User.GET("FindAllJobs", Service.FindAllJobs)
		User.POST("/AddResume", Service.AddResume)
		User.POST("/SearchApplyedJob", Service.SearchApplyedJob)
	}

	return r
}
