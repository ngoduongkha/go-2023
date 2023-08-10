package handlers

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ngoduongkha/go-2023/assignment-06/config"
	"github.com/ngoduongkha/go-2023/assignment-06/constants"
	"github.com/ngoduongkha/go-2023/assignment-06/datatransfers"
	"github.com/ngoduongkha/go-2023/assignment-06/models"
	"golang.org/x/crypto/bcrypt"
)

func (m *module) AuthenticateUser(credentials datatransfers.UserLogin) (token string, err error) {
	var user models.User
	if user, err = m.db.userOrmer.GetOneByUsername(credentials.Username); err != nil {
		return "", errors.New("incorrect credentials")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return "", errors.New("incorrect credentials")
	}

	return generateToken(user)
}

func generateToken(user models.User) (string, error) {
	now := time.Now()
	expiry := time.Now().Add(constants.AuthenticationTimeout)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.FormatUint(uint64(user.ID), 10),
		ExpiresAt: jwt.NewNumericDate(expiry),
		IssuedAt:  jwt.NewNumericDate(now),
	})
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

func (m *module) RegisterUser(credentials datatransfers.UserSignup) (err error) {
	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost); err != nil {
		return errors.New("failed hashing password")
	}
	if _, err = m.db.userOrmer.InsertUser(models.User{
		Username: credentials.Username,
		Email:    credentials.Email,
		Password: string(hashedPassword),
		Bio:      credentials.Bio,
	}); err != nil {
		log.Print(err)
		return fmt.Errorf("error inserting user. %v", err)
	}
	return
}
