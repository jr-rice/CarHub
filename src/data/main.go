package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"jr.rice/unit5act1-API/controller"
	"jr.rice/unit5act1-API/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	username := os.Args[1]
	password := os.Args[2]
	hostname := os.Args[3]
	database := os.Args[4]
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s", username, password, hostname, database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var (
		carSearchService    service.CarSearchService       = service.New(db)
		carSearchController controller.CarSearchController = controller.New(db, carSearchService)
	)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://carhub-u7on.onrender.com/"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/data/search_listed", func(context *gin.Context) {
		cars, err := carSearchController.ListListed(context)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "500 Internal Server Error! Could not fetch data!"})
		}

		context.JSON(http.StatusOK, cars)
	})

	router.POST("/data/search_wanted", func(context *gin.Context) {
		cars, err := carSearchController.ListWanted(context)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "500 Internal Server Error! Could not fetch data!"})
		}

		context.JSON(http.StatusOK, cars)
	})

	router.POST("/data/request_wanted", func(context *gin.Context) {
		resp, err := carSearchController.RequestWanted(context)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "400 Bad Request! Could not prepare SQL statement!"})
		}

		context.JSON(http.StatusOK, resp)
	})

	router.Run(":5000")
}
