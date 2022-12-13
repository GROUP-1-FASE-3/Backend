package factory

import (
	authDelivery "github.com/GROUP-1-FASE-3/Backend/features/auth/delivery"
	authRepo "github.com/GROUP-1-FASE-3/Backend/features/auth/repository"
	authService "github.com/GROUP-1-FASE-3/Backend/features/auth/service"

	userDelivery "github.com/GROUP-1-FASE-3/Backend/features/user/delivery"
	userRepo "github.com/GROUP-1-FASE-3/Backend/features/user/repository"
	userService "github.com/GROUP-1-FASE-3/Backend/features/user/service"

	villaDelivery "github.com/GROUP-1-FASE-3/Backend/features/villa/delivery"
	villaRepo "github.com/GROUP-1-FASE-3/Backend/features/villa/repository"
	villaService "github.com/GROUP-1-FASE-3/Backend/features/villa/service"

	creditcardDelivery "github.com/GROUP-1-FASE-3/Backend/features/creditcard/delivery"
	creditcardRepo "github.com/GROUP-1-FASE-3/Backend/features/creditcard/repository"
	creditcardService "github.com/GROUP-1-FASE-3/Backend/features/creditcard/service"

	// reservationDelivery "github.com/GROUP-1-FASE-3/Backend/features/reservation/delivery"
	// reservationRepo "github.com/GROUP-1-FASE-3/Backend/features/reservation/repository"
	// reservationService "github.com/GROUP-1-FASE-3/Backend/features/reservation/service"

	ratingDelivery "github.com/GROUP-1-FASE-3/Backend/features/rating/delivery"
	ratingRepo "github.com/GROUP-1-FASE-3/Backend/features/rating/repository"
	ratingService "github.com/GROUP-1-FASE-3/Backend/features/rating/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)

	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	villaRepoFactory := villaRepo.New(db)
	villaServiceFactory := villaService.New(villaRepoFactory)
	villaDelivery.New(villaServiceFactory, e)

	creditcardRepoFactory := creditcardRepo.New(db)
	creditcardServiceFactory := creditcardService.New(creditcardRepoFactory)
	creditcardDelivery.New(creditcardServiceFactory, e)

	// reservationRepoFactory := reservationRepo.New(db)
	// reservationServiceFactory := reservationService.New(reservationRepoFactory)
	// reservationDelivery.New(reservationServiceFactory, e)

	ratingRepoFactory := ratingRepo.New(db)
	ratingServiceFactory := ratingService.New(ratingRepoFactory)
	ratingDelivery.New(ratingServiceFactory, e)
}
