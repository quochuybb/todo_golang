package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

type TaskStatus int

const (
	TASK_STATUS_PENDING TaskStatus = iota
	TASK_STATUS_DONE
	TASK_STATUS_DELETED
)

var allTaskStatus = [3]string{"Doing", "Done", "Deleted"}

func (item *TaskStatus) String() string {
	return allTaskStatus[*item]
}
func parseTaskStatus(status string) (TaskStatus, error) {
	for i := range allTaskStatus {
		if allTaskStatus[i] == status {
			return TaskStatus(i), nil
		}
	}
	return TaskStatus(0), errors.New("invalid task status")
}

func (item *TaskStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan data from sql"))
	}
	strValue, err := parseTaskStatus(string(bytes))
	if err != nil {
		return errors.New(fmt.Sprint("Failed to scan data from sql"))
	}
	*item = strValue
	return nil
}
func (item TaskStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

type UserAccount struct {
	Id              string     `json:"id" gorm:"column:id;"`
	Username        string     `json:"username" gorm:"column:username;"`
	Hashed_password string     `json:"hashed_password" gorm:"column:hashed_password;"`
	Is_active       bool       `json:"is_active" gorm:"column:is_active;"`
	Created_at      *time.Time `json:"created_at" gorm:"column:created_at;"`
	Updated_at      *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

type Task struct {
	Account_id  string      `json:"account_id" gorm:"column:account_id;"`
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *TaskStatus `json:"status" gorm:"column:status;"`
	Created_at  *time.Time  `json:"created_at" gorm:"column:created_at;"`
	Updated_at  *time.Time  `json:"updated_at" gorm:"column:updated_at;"`
}

func (Task) TableName() string {
	return "task"
}

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 || p.Limit > 100 {
		p.Limit = 10
	}
}

type TaskCreation struct {
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
}

func (TaskCreation) TableName() string {
	return Task{}.TableName()
}

type TaskUpdate struct {
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
}

func (TaskUpdate) TableName() string {
	return Task{}.TableName()
}
func main() {
	dsn := "root:admin123@tcp(127.0.0.1:3307)/todo-list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		account := v1.Group("/task")
		{
			account.POST("/", CreateTask(db))
			account.GET("/:id", GetTask(db))
			account.GET("/", ListTask(db))
			account.PATCH("/:id", UpdateTask(db))
			account.DELETE("/:id", DeleteTask(db))
		}
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":4000")
}
func CreateTask(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		var data TaskCreation
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Create(&data).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
func GetTask(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		var data Task

		id, err := strconv.Atoi(context.Param("account_id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		if err := db.Where("account_id = ?", id).First(&data).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error,
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}

}
func UpdateTask(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		var data TaskUpdate

		id, err := strconv.Atoi(context.Param("account_id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Where("account_id = ?", id).Updates(&data).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error,
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}
func DeleteTask(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {

		id, err := strconv.Atoi(context.Param("account_id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		if err := db.Table(Task{}.TableName()).Where("account_id = ?", id).Delete(nil).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error,
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}

}
func ListTask(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		var data Paging
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		data.Process()
		var result []Task

		if err := db.Table(Task{}.TableName()).Count(&data.Total).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Offset((data.Page - 1) * data.Limit).Limit(data.Limit).Find(&result).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error,
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}
