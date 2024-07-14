package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryAPI interface {
	AddCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	GetCategoryList(c *gin.Context)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryRepo service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryRepo}
}

func (ct *categoryAPI) AddCategory(c *gin.Context) {
	var newCategory model.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := ct.categoryService.Store(&newCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add category success"})
}

func (ct *categoryAPI) UpdateCategory(c *gin.Context) {
	// Mengonversi ID kategori dari parameter URL ke integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Jika konversi gagal, kirim respons error dengan status BadRequest
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Category ID"})
		return
	}

	// Membuat variabel category untuk menampung data kategori yang akan diupdate
	var category model.Category
	// Mengikat data JSON yang diterima ke variabel category
	if err := c.ShouldBindJSON(&category); err != nil {
		// Jika terjadi error saat mengikat JSON, kirim respons error dengan status BadRequest
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil metode Update dari service dengan ID dan data kategori yang diupdate
	if err := ct.categoryService.Update(id, category); err != nil {
		// Jika terjadi error saat update, kirim respons error dengan status InternalServerError
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Jika berhasil, kirim respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "category update success"})
}

func (ct *categoryAPI) DeleteCategory(c *gin.Context) {
	// Mengonversi ID kategori dari parameter URL ke integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Jika konversi gagal, kirim respons error dengan status BadRequest
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Category ID"})
		return
	}

	// Memanggil metode Delete dari service dengan ID kategori
	if err := ct.categoryService.Delete(id); err != nil {
		// Jika terjadi error saat penghapusan, kirim respons error dengan status InternalServerError
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Jika berhasil, kirim respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "category delete success"})
}

func (ct *categoryAPI) GetCategoryByID(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid category ID"})
		return
	}

	category, err := ct.categoryService.GetByID(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (ct *categoryAPI) GetCategoryList(c *gin.Context) {
	// Memanggil metode GetList dari service untuk mendapatkan daftar kategori
	getCategory, err := ct.categoryService.GetList()
	if err != nil {
		// Jika terjadi error, kirim respons error dengan status InternalServerError
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	// Jika tidak ada error, kirim daftar kategori dengan status OK
	c.JSON(http.StatusOK, getCategory)
}
