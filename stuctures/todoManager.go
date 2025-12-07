package stuctures

import "fmt"

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

type TaskList struct {
	Tasks []Todo
}

func (t *TaskList) Add(title string) {
	newTask := Todo{
		ID:        len(t.Tasks) + 1,
		Title:     title,
		Completed: false,
	}

	t.Tasks = append(t.Tasks, newTask)
}

func (t *TaskList) Completed(id int) error {
	for i := range t.Tasks {
		if t.Tasks[i].ID == id {
			t.Tasks[i].Completed = true
			return nil
		}
	}
	return fmt.Errorf("task with id %d was not found", id)
}

func (t *TaskList) Delete(id int) error {
	for i := range t.Tasks {
		if t.Tasks[i].ID == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with id %d was not found", id)
}

func (t TaskList) GetCompleted() {
	for index := range t.Tasks {
		if t.Tasks[index].Completed {
			fmt.Printf("[%d] - %s завершена\n", t.Tasks[index].ID, t.Tasks[index].Title)
		}
	}
}

func (t TaskList) GetPending() {
	for index := range t.Tasks {
		if !t.Tasks[index].Completed {
			fmt.Printf("[%d] - %s в работе\n", t.Tasks[index].ID, t.Tasks[index].Title)

		}
	}
}

func TodoManager() {
	tasks := TaskList{}

	tasks.Add("learn go")
	tasks.Add("drink tea")
	tasks.Add("write simple toDo")

	tasks.Completed(1)
	tasks.Completed(3)

	tasks.GetCompleted()
	tasks.GetPending()
}
