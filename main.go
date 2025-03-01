package main

import (
	"fmt"
	"log"
	"restaurant-management/database"
	"restaurant-management/routes"
)

// func main() {
// 	if len(os.Args) > 1 {
// 		switch os.Args[1] {
// 		case "database":
// 			database.MigrateDB(database.GetDB())
// 			log.Print("\n Database loaded...")
// 			os.Exit(1)
// 		default:
// 			log.Print("\n Starting server....")
// 		}
// 	}

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8000"
// 	}
// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	routes.UserRoutes(router)
// 	router.Use(middleware.Authentication)

// 	routes.FoodRoutes(router)
// 	routes.MenuRoutes(router)
// 	routes.TableRoutes(router)
// 	routes.OrderRoutes(router)
// 	routes.OrderItemRoutes(router)
// 	routes.InvoiceRoutes(router)

// 	router.Run(":" + port)
// }

func main() {
	fmt.Println("Starting server on :8080...")
	database.ConnectDatabase()
	r := routes.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
