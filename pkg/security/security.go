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

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	tokenString = strings.TrimSpace(tokenString)

	isTokenValid, err := IsTokenValid(tokenString)

	if err != nil {
		return entity.User{}, err
	}

	if !isTokenValid {
		return entity.User{}, fmt.Errorf("Invalid token, please login again.")
	}

	claims := &Claims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.JWTSecret), nil
	})

	repository := repository.NewUserRepository()
	user, err := repository.FindByLogin(claims.Username)

	return user, err
}

// Обновление токена
func RefreshToken(account *entity.User) error {
	tokenExpiresAt := time.Now().Add(time.Hour * 24)

	claims := Claims{
		UserID:   account.ID,
		Username: account.Login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenExpiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
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

// Валидация токена
func IsTokenValid(tokenString string) (bool, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.JWTSecret), nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, nil
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		return false, nil
	}

	return true, nil
}
