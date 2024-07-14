package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(taskID int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	filebased *filebased.Data
}

func NewTaskRepo(filebasedDb *filebased.Data) *taskRepository {
	return &taskRepository{
		filebased: filebasedDb,
	}
}

func (t *taskRepository) Store(task *model.Task) error {
	t.filebased.StoreTask(*task)

	return nil
}

func (t *taskRepository) Update(taskID int, task *model.Task) error {
	t.filebased.UpdateTask(taskID, *task)
	return nil
}

func (t *taskRepository) Delete(id int) error {
	t.filebased.DeleteTask(id)
	return nil
}

// GetByID mengambil task berdasarkan ID yang diberikan.
func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	// Memanggil fungsi GetTaskByID dari filebased untuk mendapatkan task.
	obtainTaskByID, err := t.filebased.GetTaskByID(id)
	if err != nil {
		// Jika terjadi error, kembalikan nil dan error tersebut.
		return nil, err
	}
	// Jika tidak ada error, kembalikan task yang didapatkan.
	return obtainTaskByID, nil
}

// GetList mengambil semua task yang tersimpan.
func (t *taskRepository) GetList() ([]model.Task, error) {
	// Memanggil fungsi GetTasks dari filebased untuk mendapatkan semua task.
	obtainListTask, err := t.filebased.GetTasks()
	if err != nil {
		// Jika terjadi error, kembalikan nil dan error tersebut.
		return nil, err
	}
	// Jika tidak ada error, kembalikan list task yang didapatkan.
	return obtainListTask, nil
}

// GetTaskCategory mengambil semua task dalam kategori tertentu berdasarkan ID kategori.
func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	// Memanggil fungsi GetTaskListByCategory dari filebased untuk mendapatkan task berdasarkan kategori.
	obtainTaskCategory, err := t.filebased.GetTaskListByCategory(id)
	if err != nil {
		// Jika terjadi error, kembalikan nil dan error tersebut.
		return nil, err
	}
	// Jika tidak ada error, kembalikan list task kategori yang didapatkan.
	return obtainTaskCategory, nil
}
