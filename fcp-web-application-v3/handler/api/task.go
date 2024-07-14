package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskAPI interface {
	AddTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTaskByID(c *gin.Context)
	GetTaskList(c *gin.Context)
	GetTaskListByCategory(c *gin.Context)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskRepo service.TaskService) *taskAPI {
	return &taskAPI{taskRepo}
}

func (t *taskAPI) AddTask(c *gin.Context) {
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.taskService.Store(&newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add task success"})
}

// UpdateTask memperbarui task berdasarkan data JSON yang diterima dari request
func (t *taskAPI) UpdateTask(c *gin.Context) {
	var updateTask model.Task
	// Mencoba mengikat data JSON ke struktur Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		// Jika gagal, kirim respons error dengan status BadRequest
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil metode Update dari service dengan ID dan data task yang diperbarui
	if err := t.taskService.Update(updateTask.ID, &updateTask); err != nil {
		// Jika terjadi error saat update, kirim respons error dengan status InternalServerError
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Jika berhasil, kirim respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "update task success"})
}

// DeleteTask menghapus task berdasarkan ID yang diberikan dalam parameter URL
func (t *taskAPI) DeleteTask(c *gin.Context) {
	// Mengonversi ID dari string ke integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Jika konversi gagal, kirim respons error dengan status BadRequest
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	// Memanggil metode Delete dari service dengan ID task
	if err := t.taskService.Delete(id); err != nil {
		// Jika terjadi error saat penghapusan, kirim respons error dengan status InternalServerError
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Jika berhasil, kirim respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "delete task success"})
}

func (t *taskAPI) GetTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	task, err := t.taskService.GetByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// GetTaskList mengambil semua task yang tersedia
func (t *taskAPI) GetTaskList(c *gin.Context) {
	// Memanggil metode GetList dari service task untuk mendapatkan semua task
	taksList, err := t.taskService.GetList()
	if err != nil {
		// Jika terjadi error, kirim response error dengan status InternalServerError
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	// Jika tidak ada error, kirim list task dengan status OK
	c.JSON(http.StatusOK, taksList)
}

// GetTaskListByCategory mengambil semua task berdasarkan kategori yang diberikan
func (t *taskAPI) GetTaskListByCategory(c *gin.Context) {
	// Mengonversi ID kategori dari string ke integer
	GetCategoryByID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Jika konversi gagal, kirim response error dengan status BadRequest
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid category ID"})
		return
	}

	// Memanggil metode GetTaskCategory dari service task dengan ID kategori
	taksList, err := t.taskService.GetTaskCategory(GetCategoryByID)
	if err != nil {
		// Jika terjadi error saat mengambil task berdasarkan kategori, kirim response error
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	// Jika tidak ada error, kirim list task dengan status OK
	c.JSON(http.StatusOK, taksList)
}
