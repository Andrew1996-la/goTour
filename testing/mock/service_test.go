package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetTitle(id int) (string, error) {
	// Получаем заранее настроенный ответ
	args := m.Called(id)

	// Возвращаем то, что настроили через Return
	return args.String(0), args.Error(1)
}

func TestGetTaskTitle(t *testing.T) {
	// Создаем мок
	repo := new(MockTaskRepository)

	// задаем, какие данные должен вернуть мок, а затем проверяем работу сервиса
	// вернуть "Learn Go" и nil

	/*
	Настоящий репозиторий реализует GetTitle.
	Мок тоже реализует GetTitle.
	Сервису всё равно кого вызывать,
	главное чтобы объект удовлетворял интерфейсу.
	*/
	repo.On("GetTitle", 1).
		Return("Learn Go", nil)

	// Передаем мок в сервис
	service := NewTaskService(repo)

	// Вызываем тестируемый код
	title, err := service.GetTaskTitle(1)

	// Проверяем результат
	require.NoError(t, err)
	assert.Equal(t, "Learn Go", title)

	// Проверяем, что метод действительно вызвался
	repo.AssertExpectations(t)
}
