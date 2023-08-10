package handlers

import (
	"fmt"
	"log"

	"github.com/ngoduongkha/go-2023/assignment-06/config"
	"github.com/ngoduongkha/go-2023/assignment-06/datatransfers"
	"github.com/ngoduongkha/go-2023/assignment-06/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Handler HandlerFunc

type HandlerFunc interface {
	AuthenticateUser(credentials datatransfers.UserLogin) (token string, err error)
	RegisterUser(credentials datatransfers.UserSignup) (err error)

	RetrieveUser(username string) (user models.User, err error)
	UpdateUser(id uint, user datatransfers.UserUpdate) (err error)

	GetProducts() (products []models.Product, err error)
	GetProduct(id uint) (product models.Product, err error)
	UpdateProduct(id uint, product datatransfers.ProductUpdate) (err error)
	CreateProduct(product datatransfers.ProductCreate) (err error)
}

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn         *gorm.DB
	userOrmer    models.UserOrmer
	productOrmer models.ProductOrmer
}

func InitializeHandler() (err error) {
	// Initialize DB
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
			config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBDatabase,
			config.AppConfig.DBUsername, config.AppConfig.DBPassword),
	), &gorm.Config{})
	if err != nil {
		log.Println("[INIT] failed connecting to PostgreSQL")
		return
	}
	log.Println("[INIT] connected to PostgreSQL")

	// Compose handler modules
	Handler = &module{
		db: &dbEntity{
			conn:         db,
			userOrmer:    models.NewUserOrmer(db),
			productOrmer: models.NewProductOrmer(db),
		},
	}
	return
}
