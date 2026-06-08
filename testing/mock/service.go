package mock

type TaskRepository interface {
	GetTitle(id int) (string, error)
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) GetTaskTitle(id int) (string, error) {
	return s.repo.GetTitle(id)
}