package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/joho/godotenv"
)

type ResponseJWT struct {
	Status       int    `json:"status"`
	Message      string `json:"message"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

var key []byte

type Credential struct {
	Email    string
	Password string
}

type Token struct {
	Email string
	jwt.StandardClaims
}

var users = map[string]string{
	"dikisandi2006@gmail.com": "123",
}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key = []byte(os.Getenv("SECRET_KEY"))
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credential
	creds.Email = r.FormValue("Email")
	errEmail := validation.Validate(creds.Email,
		validation.Required,
	)
	if errEmail != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Email harus diisi")
		return
	}

	creds.Password = r.FormValue("Password")
	errPassword := validation.Validate(creds.Password,
		validation.Required,
	)
	if errPassword != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Password harus diisi")
		return
	}
	userPassword, ok := users[creds.Email]

	if !ok || userPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		ReturnError(w, 401, "Email atau Password salah")
		return
	}

	var tokenClaim = Token{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)

	tokenString, err := token.SignedString(key)
	expirationTime := time.Now().Add(168 * time.Hour)
	var refreshTokenClaim = Token{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaim)

	refreshTokenString, err := refreshToken.SignedString(key)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Gagal buat token")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "refreshToken",
		Value:   refreshTokenString,
		Expires: expirationTime,
	})
	var Res ResponseJWT
	Res.Status = 200
	Res.Message = "Success"
	Res.AccessToken = tokenString
	Res.RefreshToken = refreshTokenString
	json.NewEncoder(w).Encode(Res)
}

func ValidateToken(bearerToken string) string {
	tokenString := strings.Split(bearerToken, " ")[1]

	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	// fmt.Printf("%v", token.Claims)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "Key token salah"
		}
		return "Token Expire"
	}
	if !token.Valid {
		return "Token salah"
	}
	return ""
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("refreshToken")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			ReturnError(w, 401, "Token salah")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Token kosong")
		return
	}
	tokenString := c.Value

	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			ReturnError(w, 401, "Key Token salah")
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		ReturnError(w, 401, "Token expire")
		return
	}
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		ReturnError(w, 401, "Token salah")
		return
	}

	claims := token.Claims.(*Token)
	// fmt.Println(claims.Email)
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Token Expire")
		return
	}

	var tokenClaim = Token{
		Email: claims.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)

	tokenString, err = token.SignedString(key)
	expirationTime := time.Now().Add(168 * time.Hour)
	var refreshTokenClaim = Token{
		Email: claims.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaim)

	refreshTokenString, err := refreshToken.SignedString(key)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Gagal buat token")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "refreshToken",
		Value:   refreshTokenString,
		Expires: expirationTime,
	})
	var Res ResponseJWT
	Res.Status = 200
	Res.Message = "Success"
	Res.AccessToken = tokenString
	Res.RefreshToken = refreshTokenString
	json.NewEncoder(w).Encode(Res)
}
