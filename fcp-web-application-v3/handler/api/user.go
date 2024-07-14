package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserTaskCategory(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = model.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	recordUser, err := u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	// Membuat variabel user untuk menampung data login
	var user model.UserLogin
	// Mengambil data JSON dari request dan memasukkannya ke variabel user
	if err := c.BindJSON(&user); err != nil {
		// Jika terjadi error saat decode JSON, kirim response error
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	// Memeriksa apakah email atau password kosong
	if user.Email == "" || user.Password == "" {
		// Jika kosong, kirim response error
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("login data is empty"))
		return
	}

	// Membuat record user dari data yang diterima
	recordUser := model.User{
		Email:    user.Email,
		Password: user.Password,
	}

	// Memanggil fungsi Login dari service dengan recordUser sebagai parameter
	token, err := u.userService.Login(&recordUser)
	if err != nil {
		// Jika terjadi error saat login, kirim response error
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	// Mengatur waktu kadaluarsa token
	expirationTime := time.Now().Add(1 * time.Hour)
	// Membuat instance dari Claims
	claims := &model.Claims{}

	// Parsing token dengan claims
	tokenString, err := jwt.ParseWithClaims(*token, claims, func(token *jwt.Token) (interface{}, error) {
		// Mengembalikan kunci JWT
		return model.JwtKey, nil
	})
	if err != nil {
		// Jika terjadi error saat parsing, kirim response error
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	// Memeriksa validitas claims
	if claims, ok := tokenString.Claims.(*model.Claims); ok {
		// Mengatur waktu kadaluarsa dari claims
		claims.ExpiresAt = expirationTime.Unix()
	} else {
		// Jika token tidak valid, kirim response error
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("invalid token"))
		return
	}

	// Mengatur cookie dengan token yang telah diperoleh
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   *token,
		Expires: expirationTime,
	})

	// Mengirim response sukses jika login berhasil
	c.JSON(http.StatusOK, model.NewSuccessResponse("login success"))
}

func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	// Memanggil fungsi GetUserTaskCategory dari service
	cekCategory, err := u.userService.GetUserTaskCategory()
	if err != nil {
		// Jika terjadi error, kirim response error
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	// Mengirim response dengan data kategori tugas pengguna
	c.JSON(http.StatusOK, cekCategory)
}
