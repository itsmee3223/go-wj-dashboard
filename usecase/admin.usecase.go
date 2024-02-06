package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"wj-dashboard/model"
	"wj-dashboard/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/argon2"
)

var (
	argon2Time    = uint32(1)
	argon2Memory  = uint32(64 * 1024) // 64MB
	argon2Threads = uint8(4)
	argon2KeyLen  = uint32(32)
)

type IAdminUsecase interface {
	RegisterAdmin(admin *model.Admin) (model.AdminResponse, error)
	LoginEmail(admin *model.Admin) (string, error)
	LoginUsername(admin *model.Admin) (string, error)
}

type adminUsecase struct {
	ar repository.IAdminRepository
}

func NewAdminUsecase(ar repository.IAdminRepository) IAdminUsecase {
	return &adminUsecase{ar}
}

func GenerateRandomSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func generateJWT(identifier string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"identifier": identifier,
		"exp":        time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (au *adminUsecase) RegisterAdmin(admin *model.Admin) (model.AdminResponse, error) {
	time := uint32(1)
	memory := uint32(64 * 1024)
	threads := uint8(4)
	keyLen := uint32(32)

	saltStr, err := GenerateRandomSalt(16)
	if err != nil {
		return model.AdminResponse{}, err
	}

	salt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		return model.AdminResponse{}, err
	}

	hash := argon2.IDKey([]byte(admin.Password), salt, time, memory, threads, keyLen)
	passwordToStore := fmt.Sprintf("%s:%x", saltStr, hash)

	newAdmin := model.Admin{
		Name:     admin.Name,
		Username: admin.Username,
		Email:    admin.Email,
		Password: passwordToStore,
		RoleID:   admin.RoleID,
	}
	if err := au.ar.CreateAdmin(&newAdmin); err != nil {
		return model.AdminResponse{}, err
	}

	resAdmin := model.AdminResponse{
		ID:       newAdmin.ID,
		Name:     newAdmin.Name,
		Email:    newAdmin.Email,
		Username: newAdmin.Username,
		RoleID:   newAdmin.RoleID,
	}

	return resAdmin, nil
}

func (au *adminUsecase) LoginEmail(admin *model.Admin) (string, error) {
	storedAdmin := model.Admin{}

	if err := au.ar.GetAdminByEmail(&storedAdmin, admin.Email); err != nil {
		return "", err
	}

	parts := strings.Split(storedAdmin.Password, ":")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid password format")
	}
	salt, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return "", err
	}
	storedHash := parts[1]

	hash := argon2.IDKey([]byte(admin.Password), salt, argon2Time, argon2Memory, argon2Threads, argon2KeyLen)
	if fmt.Sprintf("%x", hash) != storedHash {
		return "", fmt.Errorf("invalid password")
	}

	return generateJWT(storedAdmin.Email)
}

func (au *adminUsecase) LoginUsername(admin *model.Admin) (string, error) {
	storedAdmin := model.Admin{}

	if err := au.ar.GetAdminByUsername(&storedAdmin, admin.Username); err != nil {
		return "", err
	}

	parts := strings.Split(storedAdmin.Password, ":")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid password format")
	}
	salt, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return "", err
	}
	storedHash := parts[1]

	hash := argon2.IDKey([]byte(admin.Password), salt, argon2Time, argon2Memory, argon2Threads, argon2KeyLen)
	if fmt.Sprintf("%x", hash) != storedHash {
		return "", fmt.Errorf("invalid password")
	}

	return generateJWT(storedAdmin.Username)
}
