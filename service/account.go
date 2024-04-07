package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/lw396/WeComCopilot/internal/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) ParseToken(ctx context.Context, tokenString string) (result *jwt.StandardClaims, err error) {
	result = &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, result, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.jwt.Secret), nil
	})
	if err != nil {
		return nil, errors.New(errors.CodeAuthTokenInvalid, err.Error())
	}

	if !token.Valid {
		return nil, errors.New(errors.CodeAuthTokenInvalid, "token is invalid")
	}

	return
}

func (s *Service) AuthenticateAccount(username, password string) (err error) {
	if username == "" || username != s.admin.Username {
		return errors.New(errors.CodeAccountNotExist, "username is wrong")
	}
	if password == "" {
		return errors.New(errors.CodeAuthWrongPassword, "password is empty")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(s.admin.Password), []byte(password)); err != nil {
		return errors.New(errors.CodeAuthWrongPassword, "password is wrong")
	}

	return
}

type Authentication struct {
	Token     string
	ExpiredAt int64
}

func (s *Service) CreateToken(ctx context.Context, username string) (result *Authentication, err error) {
	now := time.Now()
	expiredAt := now.Add(time.Duration(s.jwt.ExpireSecs) * time.Second)
	claims := jwt.StandardClaims{
		Id:        uuid.New().String(),
		Issuer:    "chat-copilot",
		ExpiresAt: expiredAt.Unix(),
		IssuedAt:  now.Unix(),
		Subject:   username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.jwt.Secret))
	if err != nil {
		return
	}

	return &Authentication{
		Token:     t,
		ExpiredAt: expiredAt.Unix(),
	}, nil
}
