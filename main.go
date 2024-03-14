package main

import (
	"os"
    "log"
	
	"github.com/gin-gonic/gin"
	"github.com/yeliwzinn/EcommerceGO/controllers"
	"github.com/yeliwzinn/EcommerceGO/database"
	"github.com/yeliwzinn/EcommerceGO/middleware"
	"github.com/yeliwzinn/EcommerceGO/routes"
)

// main work of this func is to start the server

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	//app will handle all other routes
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New() //createss a router
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":"+port))

}
