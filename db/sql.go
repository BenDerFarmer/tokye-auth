package db

import (
	"log"

	"github.com/ChaotenHG/auth-server/config"
	. "github.com/ChaotenHG/auth-server/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn string

func LoadSQLCredentials(cfg *config.Config) {
	dsn = cfg.Secret.SqlDSN
}

func openDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		log.Panic("failed to connect database")
	}

	return db
}

func InitialMigration() {
	db := openDB()

	// Migrate the schema
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal(err)
	}
}

func FindUser(email string) (User, error) {

	db := openDB()

	var (
		user User
		err  error
	)

	err = db.Model(&User{}).Where("Email = ?", email).First(&user).Error

	return user, err
}

func SaveUser(user *User) error {
	db := openDB()

	return db.Model(User{}).Where("Id = ?", user.ID).Save(*user).Error
}

func FindUserByID(id string) (User, error) {

	db := openDB()

	var (
		result User
		err    error
	)

	err = db.Model(User{ID: id}).First(&result).Error

	return result, err
}

func CreateUser(email string) error {

	db := openDB()

	return db.Create(&User{Email: email}).Error

}
