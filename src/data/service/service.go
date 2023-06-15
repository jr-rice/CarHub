package service

import (
	"database/sql"

	"jr.rice/unit5act1-API/entity"

	_ "github.com/lib/pq"
)

type CarSearchService interface {
	ListListed(carRequest entity.CarRequestData) ([]entity.ListedCar, error)
	ListWanted(carRequest entity.CarRequestData) ([]entity.WantedCar, error)
	RequestWanted(carListing entity.WantedCar) (entity.WantedCar, error)
}

type carSearchService struct {
	db *sql.DB
}

func New(db *sql.DB) CarSearchService {
	return &carSearchService{
		db: db,
	}
}

func (service *carSearchService) ListListed(carRequest entity.CarRequestData) ([]entity.ListedCar, error) {
	var (
		args []interface{}
		cars []entity.ListedCar
		car  entity.ListedCar
		rows *sql.Rows
	)

	query := "SELECT manufacturer, model, stock FROM listed_cars"

	if carRequest.Manufacturer != "" && carRequest.Model != "" {
		query += " WHERE manufacturer ILIKE '%' || $1 || '%' AND model ILIKE '%' || $2 || '%'"
		args = append(args, carRequest.Manufacturer, carRequest.Model)
	} else if carRequest.Manufacturer != "" && carRequest.Model == "" {
		query += " WHERE manufacturer ILIKE '%' || $1 || '%'"
		args = append(args, carRequest.Manufacturer)
	} else if carRequest.Manufacturer == "" && carRequest.Model != "" {
		query += " WHERE model ILIKE '%' || $1 || '%'"
		args = append(args, carRequest.Model)
	}

	stmt, err := service.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	if len(args) > 0 {
		rows, err = stmt.Query(args...)
	} else {
		rows, err = stmt.Query()
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&car.Manufacturer, &car.Model, &car.Stock); err != nil {
			return nil, err
		}

		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}

func (service *carSearchService) ListWanted(carRequest entity.CarRequestData) ([]entity.WantedCar, error) {
	var (
		args []interface{}
		cars []entity.WantedCar
		car  entity.WantedCar
		rows *sql.Rows
	)

	query := "SELECT manufacturer, model FROM wanted_cars"

	if carRequest.Manufacturer != "" && carRequest.Model != "" {
		query += " WHERE manufacturer ILIKE '%' || $1 || '%' AND model ILIKE '%' || $2 || '%'"
		args = append(args, carRequest.Manufacturer, carRequest.Model)
	} else if carRequest.Manufacturer != "" && carRequest.Model == "" {
		query += " WHERE manufacturer ILIKE '%' || $1 || '%'"
		args = append(args, carRequest.Manufacturer)
	} else if carRequest.Manufacturer == "" && carRequest.Model != "" {
		query += " WHERE model ILIKE '%' || $1 || '%'"
		args = append(args, carRequest.Model)
	}

	stmt, err := service.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	if len(args) > 0 {
		rows, err = stmt.Query(args...)
	} else {
		rows, err = stmt.Query()
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&car.Manufacturer, &car.Model); err != nil {
			return nil, err
		}

		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}

func (service *carSearchService) RequestWanted(carListing entity.WantedCar) (entity.WantedCar, error) {
	stmt, err := service.db.Prepare("INSERT INTO wanted_cars(manufacturer, model) VALUES($1, $2)")
	if err != nil {
		return entity.WantedCar{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(carListing.Manufacturer, carListing.Model)
	if err != nil {
		return entity.WantedCar{}, err
	}

	return carListing, nil
}
