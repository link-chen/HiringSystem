package DataBaseService

import (
	"HiringSystem/Utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:123456@tcp(localhost:3306)/HiringSystem?charset=utf8mb4&parseTime=True&loc=Local"
	//default database
	//123456 is the password of root user
	//HiringSystem is the default database
)

var db *gorm.DB

func CreateDataBase() {
	//第一次执行
	db.AutoMigrate(&Utils.User{})
	db.AutoMigrate(&Utils.HRUser{})
	db.AutoMigrate(&Utils.Job{})
}
func InitalDataBase() {
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if Db != nil {
		db = Db
	}
}

func AddJobToDataBase(job Utils.Job) bool {
	ans := db.Create(&job)
	fmt.Println(ans)
	return true
}
func DeleteJobFromDataBase(job Utils.Job) bool {

	db.Model(&Utils.User{}).Where("job_id = ?", job.Id).Update("job_id", 0)
	second := db.Delete(&job)
	if second.RowsAffected != 0 {
		return true
	}
	return false
}

func GetJobsPostedByHR(user Utils.HRUser) []Utils.Job {
	var Jobs []Utils.Job
	db.Where("posted_by=?", user.Id).Find(&Jobs)
	return Jobs
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
	if ans.RowsAffected == 0 {
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
	if ans.RowsAffected == 0 {
		return false
	}
	return true
}

func GetResumeAddress(user Utils.User) string {
	ans := db.Where("id=?", user.Id).Find(&user)
	if ans == nil {
		fmt.Println("Can not Get Resume Address")
	} else {

	}
	fmt.Println(ans.RowsAffected)
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

func FindUserJobId(user Utils.User) uint {
	db.Where("id=?", user.Id).Find(&user)
	return user.JobId
}

func FindUserApplyJob(user Utils.User) ([]interface{}, error) {
	rows, err := db.Debug().Raw(`
    SELECT users.id, jobs.id as job_id, jobs.title, jobs.description, jobs.posted_by, jobs.status
    FROM users
    INNER JOIN jobs ON users.job_id = jobs.id
    WHERE users.id = ?`, user.Id).Rows()

	var ans []interface{}
	if err != nil {
		// 处理错误
		fmt.Println("Error:", err)
		return ans, err
	}
	defer rows.Close()

	// 使用 map 存储查询结果
	result := make(map[string]interface{})
	for rows.Next() {
		err := db.ScanRows(rows, &result)
		if err != nil {
			// 处理错误
			fmt.Println("Error:", err)
			return ans, err
		}
	}

	// 输出查询结果
	ans = append(ans, (result["id"]))
	ans = append(ans, (result["job_id"]))
	ans = append(ans, (result["title"]))
	ans = append(ans, (result["description"]))
	ans = append(ans, (result["posted_by"]))
	ans = append(ans, (result["status"]))

	return ans, nil
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

func HRLogin(user Utils.HRUser) bool {
	pass := user.Password
	ans := db.Where("id=?", user.Id).Find(&user)
	if ans.RowsAffected == 0 {
		return false
	}
	return pass == user.Password
}

func SearchJobStatus(JobId uint) string {
	var Job Utils.Job
	db.Where("id=?", JobId).Find(&Job)
	fmt.Println(Job.Status)
	return Job.Status
}
