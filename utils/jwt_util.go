// // utils/jwt_util.go
package utils

// import (
// 	"time"

// 	"github.com/golang-jwt/jwt"
// )

// var jwtSecret = []byte("secret")

// // Struktur claim kustom untuk menyimpan data user dalam token
// type CustomClaim struct {
// 	UserID int `json:"user_id"`
// 	jwt.StandardClaims
// }

// // Fungsi untuk membuat token JWT dari data user
// func CreateToken(userID int) (string, error) {
// 	claims := CustomClaim{
// 		UserID: userID,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 1 hari
// 			Issuer:    "YourAppName",                         // Ganti dengan nama aplikasi Anda
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtSecret)

// 	return tokenString, err
// }