package usecase

import (
	"errors"
	"os"
	"time"

	"github.com/dendik-creation/task-control/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepo   domain.UserRepository
	columnRepo domain.ColumnRepository
}

func NewAuthUsecase(userRepo domain.UserRepository, columnRepo domain.ColumnRepository) domain.UserUsecase {
	return &authUsecase{
		userRepo:   userRepo,
		columnRepo: columnRepo,
	}
}

func (u *authUsecase) Register(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Failed to hash password")
	}
	user.Password = string(hashedPassword)

	err = u.userRepo.Create(user)
	if err != nil {
		return err
	}

	defaultColumns := []domain.Column{
		{UserID: user.ID, Name: "To Do", Color: "#E2E8F0", Position: 0},
		{UserID: user.ID, Name: "In Progress", Color: "#FEF08A", Position: 1},
		{UserID: user.ID, Name: "Done", Color: "#BBF7D0", Position: 2},
	}

	for _, col := range defaultColumns {
		_ = u.columnRepo.Create(&col)
	}

	return nil
}

func (u *authUsecase) Login(email string, password string) (string, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "kikukikuk"
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("gagal membuat token autentikasi")
	}

	return tokenString, nil
}
