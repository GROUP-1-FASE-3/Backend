package mysql

import (
	"fmt"
	"log"

	"github.com/GROUP-1-FASE-3/Backend/config"

	creditcardRepo "github.com/GROUP-1-FASE-3/Backend/features/creditcard/repository"
	// ratingRepo "github.com/GROUP-1-FASE-3/Backend/features/rating/repository"
	reservationRepo "github.com/GROUP-1-FASE-3/Backend/features/reservation/repository"
	villaRepo "github.com/GROUP-1-FASE-3/Backend/features/villa/repository"

	userRepo "github.com/GROUP-1-FASE-3/Backend/features/user/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db
}

func DBMigration(db *gorm.DB) {
	db.AutoMigrate(&userRepo.User{})
	db.AutoMigrate(&villaRepo.Villa{})
	db.AutoMigrate(&creditcardRepo.CreditCard{})
	db.AutoMigrate(&reservationRepo.Reservation{})
	// db.AutoMigrate(&ratingRepo.Rating{})

}
