package service

import (
	"crypto/sha1"
	"fmt"
	todo_app "github.com/milya1265/todo-list.git"
	"github.com/milya1265/todo-list.git/pkg/repository"
)

const salt = "qcnc216237gdhoi2kxbct23dui"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo_app.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
