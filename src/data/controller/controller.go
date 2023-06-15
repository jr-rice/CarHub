package controller

import (
	"database/sql"

	"jr.rice/unit5act1-API/entity"
	"jr.rice/unit5act1-API/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type CarSearchController interface {
	ListListed(context *gin.Context) ([]entity.ListedCar, error)
	ListWanted(context *gin.Context) ([]entity.WantedCar, error)
	RequestWanted(context *gin.Context) (entity.WantedCar, error)
}

type carSearchController struct {
	db      *sql.DB
	service service.CarSearchService
}

func New(db *sql.DB, service service.CarSearchService) CarSearchController {
	return &carSearchController{
		db:      db,
		service: service,
	}
}

func (cntrl *carSearchController) ListListed(context *gin.Context) ([]entity.ListedCar, error) {
	var requestData entity.CarRequestData

	if err := context.BindJSON(&requestData); err != nil {
		return []entity.ListedCar{}, nil
	}

	return cntrl.service.ListListed(requestData)
}

func (cntrl *carSearchController) ListWanted(context *gin.Context) ([]entity.WantedCar, error) {
	var requestData entity.CarRequestData

	if err := context.BindJSON(&requestData); err != nil {
		return []entity.WantedCar{}, nil
	}

	return cntrl.service.ListWanted(requestData)
}

func (cntrl *carSearchController) RequestWanted(context *gin.Context) (entity.WantedCar, error) {
	var car entity.WantedCar

	if err := context.BindJSON(&car); err != nil {
		return entity.WantedCar{}, nil
	}

	return cntrl.service.RequestWanted(car)
}
