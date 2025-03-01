package controller

import (
	"github.com/gin-gonic/gin"
	"todolist/internal/service"
	"todolist/response"
)

type TaskController struct {
	taskService *service.TaskService
}

func NewTaskController() *TaskController {
	return &TaskController{
		taskService: service.NewTaskService(),
	}
}

func (t *TaskController) AddTask(c *gin.Context) {
	response.SuccessResponse(c, []string{"huy", "vi"}, 20000)
}
