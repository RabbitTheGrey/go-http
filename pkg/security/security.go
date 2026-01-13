package security

import (
	"fmt"
	"go-web/conf"
	"go-web/internal/entity"
	"go-web/internal/repository"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Получение текущего полльзователя.
// Может вернуть ошибку Unauthorized
func CurrentUser(r *http.Request) (entity.User, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return entity.User{}, fmt.Errorf("Authorization header is missing")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return entity.User{}, fmt.Errorf("invalid authorization header format")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	token = strings.TrimSpace(token)

	repository := repository.NewUserRepository()
	user, err := repository.FindByToken(token)

	return user, err
}

// Обновление токена
func RefreshToken(account *entity.User) error {
	claims := Claims{
		UserID:   account.ID,
		Username: account.Login,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(conf.JWTSecret))
	if err != nil {
		return err
	}

	account.Token = &tokenString
	return nil
}

// Хеширование пароля
func HashPassword(plainPassword string) (string, error) {
	bytes := []byte(plainPassword)

	passwordHash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

// Проверка пароля
func ValidatePassword(user entity.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
