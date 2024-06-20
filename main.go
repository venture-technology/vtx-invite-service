package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"github.com/venture-technology/vtx-invites/config"
	"github.com/venture-technology/vtx-invites/internal/controller"
	"github.com/venture-technology/vtx-invites/internal/repository"
	"github.com/venture-technology/vtx-invites/internal/service"
)

func main() {

	router := gin.Default()

	config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = migrate(db, config.Database.Schema)
	if err != nil {
		log.Fatalf("failed to execute migrations: %v", err)
	}

	producer := kafka.NewWriter(kafka.WriterConfig{Brokers: []string{config.Messaging.Brokers}, Topic: config.Messaging.Topic, Balancer: &kafka.LeastBytes{}})

	inviteRepository := repository.NewInviteRepository(db)
	kafkaRepository := repository.NewKafkaRepository(producer)

	inviteService := service.NewInviteService(inviteRepository, kafkaRepository)

	inviteController := controller.NewInviteController(inviteService)

	inviteController.RegisterRoutes(router)

	log.Printf("initing service: %s", config.Name)

	router.Run(fmt.Sprintf(":%d", config.Server.Port))

}

func newPostgres(dbConfig config.Database) string {
	return "user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.Name +
		" host=" + dbConfig.Host +
		" port=" + dbConfig.Port +
		" sslmode=disable"
}

func migrate(db *sql.DB, filepath string) error {
	schema, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}
