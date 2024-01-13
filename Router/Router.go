package Router

import (
	"HiringSystem/Service"
	"HiringSystem/Utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	secretKey = []byte("your_secret_key") // 替换为实际的密钥
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("进入鉴权中间件")
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			fmt.Println("无token")
			c.JSON(200, Utils.Response{http.StatusUnauthorized, "401", "UnAuthorization"})
			c.Abort()
			return
		}

		// 解析并验证JWT令牌
		_, err := jwt.ParseWithClaims(tokenString, &Utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			fmt.Println("无效token")
			c.JSON(200, Utils.Response{http.StatusUnauthorized, "401", "UnAuthorization"})
			c.Abort()
			return
		}

		c.Next()
	}
}

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

	r.POST("/HRService/HRLogin", Service.HRLogin)
	r.POST("/UserService/Login", Service.Login)
	r.POST("/UserService/Regist", Service.Regist)

	r.Use(authMiddleware())

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
	}

	User := r.Group("/UserService")
	{
		User.POST("/ApplyJob", Service.ApplyJob)
		User.POST("FindAllJobs", Service.FindAllJobs)
		User.POST("/AddResume", Service.AddResume)
		User.POST("/SearchApplyedJob", Service.SearchApplyedJob)
	}

	return r
}
