package service

import "todolist/internal/repo"

type TaskService struct {
	taskRepo *repo.TaskRepo
}

func NewTaskService() *TaskService {
	return &TaskService{
		taskRepo: repo.NewTaskRepo(),
	}
}

func (ts *TaskService) GetInfoTask() string {
	return ts.taskRepo.GetInfoTask()
}
