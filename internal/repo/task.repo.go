package repo

type TaskRepo struct{}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{}
}

func (tr *TaskRepo) GetInfoTask() string {
	return "text"
}
