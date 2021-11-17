package user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/vcnt72/try-golang-database-lib/entity"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userRepository Repository
	logger         *zap.Logger
}

func NewService(repo Repository, logger *zap.Logger) Usecase {
	return &Service{
		userRepository: repo,
		logger:         logger,
	}
}

func (s *Service) Register(ctx context.Context, createDTO CreateUserDTO) (*entity.User, error) {

	password := []byte(createDTO.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		err := errors.Wrap(err, "error on hashing the password on creating user")
		s.logger.Sugar().Error(err)
		return nil, err
	}

	user := &entity.User{
		Name:     createDTO.Name,
		Email:    createDTO.Email,
		Password: string(hashedPassword),
	}

	user, err = s.userRepository.Store(ctx, user)

	if err != nil {
		err = errors.Wrap(err, "error on storing user to db")

		s.logger.Sugar().Error(err)
		return nil, err
	}

	return user, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := s.userRepository.FindByID(ctx, id)

	if errors.Is(err, entity.ErrNotFound) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := s.userRepository.FindByEmail(ctx, email)

	if errors.Is(err, entity.ErrNotFound) {
		return nil, err
	}

	if err != nil {

		return nil, err
	}

	return user, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (user *entity.User, tokenStr string, err error) {
	user, err = s.userRepository.FindByEmail(ctx, email)

	if errors.Is(err, entity.ErrNotFound) {
		return nil, "", err
	}

	if err != nil {
		return nil, "", err
	}

	secret := []byte(viper.GetString("jwt.secret"))

	claims := &jwt.StandardClaims{
		Id:        user.ID,
		Issuer:    "try-golang-database.com",
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenStr, err = token.SignedString(secret)

	if err != nil {
		return nil, "", err
	}

	return
}
