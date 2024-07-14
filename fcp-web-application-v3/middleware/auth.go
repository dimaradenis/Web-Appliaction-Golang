package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// Mencoba mendapatkan cookie dengan nama "session_token"
		authHeader, err := ctx.Cookie("session_token")
		if err != nil {
			// Jika terjadi error saat mengambil cookie, panggil fungsi handleAuthError
			handleAuthError(ctx)
			return
		}

		// Inisialisasi struktur Claims untuk menyimpan data klaim dari token
		claims := &model.Claims{}
		// Parsing token JWT menggunakan klaim dan kunci rahasia
		token, err := jwt.ParseWithClaims(authHeader, claims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if err != nil {
			// Jika terjadi error saat parsing token, panggil fungsi handleTokenError
			handleTokenError(ctx, err)
			return
		}

		// Memeriksa apakah token yang diberikan valid
		if !token.Valid {
			// Jika token tidak valid, hentikan proses dengan status Unauthorized
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Menyimpan email dari klaim ke dalam konteks permintaan
		ctx.Set("email", claims.Email)
		// Lanjutkan ke middleware berikutnya
		ctx.Next()
	})
}

func handleAuthError(ctx *gin.Context) {
	// Memeriksa jenis konten dari permintaan
	if ctx.Request.Header.Get("Content-Type") == "application/json" {
		// Jika jenis konten adalah JSON, kirim respons JSON dengan status Unauthorized
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("Unauthorized"))
	} else {
		// Jika jenis konten bukan JSON, arahkan pengguna ke halaman login
		ctx.Redirect(http.StatusSeeOther, "/login")
	}
}

func handleTokenError(ctx *gin.Context, err error) {
	// Memeriksa jenis error dari token
	if err == jwt.ErrSignatureInvalid {
		// Jika error karena tanda tangan tidak valid, hentikan proses dengan status Unauthorized
		ctx.AbortWithStatus(http.StatusUnauthorized)
	} else {
		// Jika terjadi error lain, hentikan proses dengan status BadRequest
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
}
