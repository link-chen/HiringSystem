package DataBaseService

import (
	"HiringSystem/Utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:123456@tcp(localhost:3306)/HiringSystem?charset=utf8mb4&parseTime=True&loc=Local"
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
	//db.AutoMigrate(&Utils.User{})
	//db.AutoMigrate(&Utils.HRUser{})
	//db.AutoMigrate(&Utils.Job{})
}

func AddJobToDataBase(job Utils.Job) bool {
	ans := db.Create(&job)
	fmt.Println(ans)
	return true
}
func DeleteJobFromDataBase(job Utils.Job) {

	db.Model(&Utils.User{}).Where("job_id = ?", job.Id).Update("job_id", 0)
	db.Delete(&job)
}

func GetJobsPostedByHR(user Utils.HRUser) []uint {
	var Jobs []Utils.Job
	db.Where("posted_by=?", user.Id).Find(&Jobs)
	var ansarray []uint
	for i := 0; i < len(Jobs); i++ {
		ansarray = append(ansarray, Jobs[i].Id)
	}
	return ansarray
}

func AddSimpleUser(user Utils.User) (bool, uint) {
	ans := db.Create(&user)
	if ans != nil {
		return true, user.Id
	}
	return false, 0
}

func SimpleUserLogin(user Utils.User) bool {
	pass := user.Password
	ans := db.Where("id=?", user.Id).Find(&user)
	if ans == nil {
		return false
	}
	return pass == user.Password
}

func UserApplyJob(user Utils.User) bool {
	ans := db.Model(&Utils.User{}).Where("id = ?", user.Id).Update("job_id", user.JobId)
	if ans == nil {
		return false
	}
	return true
}

func SearchAllJobs() []Utils.Job {
	var Jobs []Utils.Job
	db.Find(&Jobs)
	return Jobs
}

func FindUsersByJobId(user Utils.User) []Utils.User {
	var User []Utils.User
	db.Where("job_id=?", user.Id).Find(&User)
	return User
}

func AddResumeToUser(id int, ResumeAddress string) bool {
	ans := db.Model(&Utils.User{}).Where("id = ?", id).Update("resume_address", ResumeAddress)
	if ans == nil {
		return false
	}
	return true
}

func GetResumeAddress(user Utils.User) string {
	db.Where("id=?", user.Id).Find(&user)
	return user.ResumeAddress
}

func GetEmail(user Utils.User) string {
	ans := db.Where("id=?", user.Id).Find(&user)
	if ans == nil {
		fmt.Println("GetFailed")
		return ""
	}
	return user.Email
}

func CleanUserApply(user Utils.User) {
	db.Model(&Utils.User{}).Where("id = ?", user.Id).Update("job_id", 0)
}

func CheckResumeExist(user Utils.User) bool {
	db.Where("id=?", user.Id).Find(&user)
	if user.ResumeAddress == "" {
		return false
	}
	return true
}
