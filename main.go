package main

import (
	"amirhossein-shakeri/go-prosperity-game/auth"
	"amirhossein-shakeri/go-prosperity-game/db"
	"amirhossein-shakeri/go-prosperity-game/item"
	"amirhossein-shakeri/go-prosperity-game/level"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var PORT = os.Getenv("PORT")
var ADDRESS = "localhost:" + PORT // drop localhost in production

func main() {
	log.Println("🚀 Starting Prosperity Game Server ...")
	// db := &db.MongoDB{}
	// db.Init()
	// c := db.Client.Database("prosperity-game").Collection("users")
	// c.InsertOne()
	// defer db.Disconnect()
	db.InitMGM()
	gin.ForceConsoleColor()
	mainRouter := gin.Default()
	setupRoutes(mainRouter)
	mainRouter.Run(ADDRESS)
	// https://github.com/gin-gonic/gin#grouping-routes
}

func setupRoutes(router *gin.Engine) {
	router.Any("/health", healthHandler)
	authRouter := router.Group("/auth")
	{
		authRouter.GET("/", auth.AuthorizeJWT(), auth.GetInfo) // get session info
		authRouter.POST("/", auth.Login)                       // login
		authRouter.POST("/signup", auth.Signup)                // signup
	}
	levelRouter := router.Group("/levels", auth.AuthorizeJWT())
	{
		levelRouter.GET("/", level.GetLevels)
		levelRouter.GET("/:levelId", level.GetLevel)
		levelRouter.POST("/", level.PostLevel)
		levelRouter.PATCH("/:levelId", level.UpdateLevel)
		levelRouter.PUT("/:levelId", level.UpdateLevel)
		levelRouter.DELETE("/:levelId", level.DeleteLevel)
	}
	itemRouter := router.Group("/items", auth.AuthorizeJWT())
	{
		itemRouter.GET("/:levelId", item.GetItems)     // Get all level items
		itemRouter.POST("/", level.PostItem)           // Create new item in level
		itemRouter.PATCH("/:itemId")                   // Update an item in a level
		itemRouter.PUT("/:itemId")                     // Update an item in a level
		itemRouter.DELETE("/:itemId", item.DeleteItem) // Delete an item from a level
		// reorder or change the order ...
	}
}

func healthHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "I'm Alive!",
	})
}
