package main

import (
	"log"

	"github.com/alphatechnolog/purplish-memberships/infrastructure/database"
	"github.com/alphatechnolog/purplish-memberships/internal/config"
	"github.com/alphatechnolog/purplish-memberships/internal/di"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const ENV_FILE = ".env"

func main() {
	cfg, err := config.LoadConfig(ENV_FILE)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
		panic(err)
	}

	db := database.MustOpenDB("sqlite3", cfg.DatabaseURL)
	defer db.Close()

	router := gin.Default()
	defer router.Run(":" + cfg.ServerPort)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	membershipGroup := router.Group("/memberships")

	membershipInjector := di.NewMembershipInjector(db)
	membershipInjector.Inject(membershipGroup)
}
