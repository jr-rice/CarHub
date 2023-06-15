package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type ListedCar struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Stock        int    `json:"stock"`
}

type WantedCar struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
}

func main() {
	username := os.Args[1]
	password := os.Args[2]
	hostname := os.Args[3]
	connStr := fmt.Sprintf("postgres://%s:%s@%s/cars_test?sslmode=disable", username, password, hostname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	router := gin.Default()

	/* Show All Listed Cars */
	router.GET("/get_test/all_listed_cars", func(context *gin.Context) {
		var (
			cars []ListedCar
			car  ListedCar
		)

		rows, err := db.Query("SELECT manufacturer, model, stock FROM listed_cars")
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not connect to SQL Database!"})
			return
		}

		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&car.Manufacturer, &car.Model, &car.Stock)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch table data!"})
				return
			}

			cars = append(cars, car)
		}

		if err := rows.Err(); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch row data!"})
			return
		}

		context.JSON(http.StatusOK, cars)
	})

	/* Show All Wanted Cars */
	router.GET("/get_test/all_wanted_cars", func(context *gin.Context) {
		var (
			cars []WantedCar
			car  WantedCar
		)

		rows, err := db.Query("SELECT manufacturer, model FROM wanted_cars")
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not connect to SQL Database!"})
			return
		}

		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&car.Manufacturer, &car.Model)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch table data!"})
				return
			}

			cars = append(cars, car)
		}

		if err := rows.Err(); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch row data!"})
			return
		}

		context.JSON(200, cars)
	})

	/* Query Listed Cars */
	router.POST("/post_test/query_listed_cars", func(context *gin.Context) {
		var (
			requestData struct {
				Manufacturer string `json:"manufacturer"`
				Model        string `json:"model"`
			}

			args []interface{}
			cars []ListedCar
			car  ListedCar
			rows *sql.Rows
		)

		if err := context.ShouldBindJSON(&requestData); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind request data to JSON!"})
			return
		}

		query := "SELECT manufacturer, model, stock FROM listed_cars"

		if requestData.Manufacturer != "" && requestData.Model != "" {
			query += " WHERE manufacturer ILIKE '%' || $1 || '%' AND model ILIKE '%' || $2 || '%'"
			args = append(args, requestData.Manufacturer, requestData.Model)
		} else if requestData.Manufacturer != "" && requestData.Model == "" {
			query += " WHERE manufacturer ILIKE '%' || $1 || '%'"
			args = append(args, requestData.Manufacturer)
		} else if requestData.Manufacturer == "" && requestData.Model != "" {
			query += " WHERE model ILIKE '%' || $1 || '%'"
			args = append(args, requestData.Model)
		}

		stmt, err := db.Prepare(query)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not prepare SQL statement!"})
			return
		}

		defer stmt.Close()

		if len(args) > 0 {
			rows, err = stmt.Query(args...)
		} else {
			rows, err = stmt.Query()
		}

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not connect to SQL Database!"})
			return
		}

		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&car.Manufacturer, &car.Model, &car.Stock)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch table data!"})
				return
			}

			cars = append(cars, car)
		}

		if err := rows.Err(); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch row data!"})
			return
		}

		context.JSON(http.StatusOK, cars)
	})

	/* Query Wanted Cars */
	router.POST("/post_test/query_wanted_cars", func(context *gin.Context) {
		var (
			requestData struct {
				Manufacturer string `json:"manufacturer"`
				Model        string `json:"model"`
			}

			args []interface{}
			cars []WantedCar
			car  WantedCar
		)

		if err := context.ShouldBindJSON(&requestData); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind request data to JSON!"})
			return
		}

		query := "SELECT manufacturer, model FROM wanted_cars WHERE"

		if requestData.Manufacturer != "" && requestData.Model != "" {
			query += " manufacturer ILIKE '%' || $1 || '%' AND model ILIKE '%' || $2 || '%'"
			args = append(args, requestData.Manufacturer, requestData.Model)
		} else if requestData.Manufacturer != "" && requestData.Model == "" {
			query += " manufacturer ILIKE '%' || $1 || '%'"
			args = append(args, requestData.Manufacturer)
		} else if requestData.Manufacturer == "" && requestData.Model != "" {
			query += " model ILIKE '%' || $1 || '%'"
			args = append(args, requestData.Model)
		}

		stmt, err := db.Prepare(query)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not prepare SQL statement!"})
			return
		}

		defer stmt.Close()

		rows, err := stmt.Query(args...)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not connect to SQL Database!"})
			return
		}

		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&car.Manufacturer, &car.Model)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch table data!"})
				return
			}

			cars = append(cars, car)
		}

		if err := rows.Err(); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch row data!"})
			return
		}

		context.JSON(http.StatusOK, cars)
	})

	/* Add Cars to Wanted */
	router.POST("/post_test/add_wanted_cars", func(context *gin.Context) {
		var car WantedCar

		if err := context.BindJSON(&car); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload!"})
			return
		}

		stmt, err := db.Prepare("INSERT INTO wanted_cars(manufacturer, model) VALUES($1, $2)")
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not prepare SQL statement!"})
			return
		}

		defer stmt.Close()

		_, err = stmt.Exec(car.Manufacturer, car.Model)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into database!"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Data inserted successfully!"})
	})

	router.Run(":5000")
}
