package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/segmentio/kafka-go"
	"github.com/venture-technology/vtx-invites/config"
	"github.com/venture-technology/vtx-invites/internal/repository"
	"github.com/venture-technology/vtx-invites/models"

	_ "github.com/lib/pq"
)

func setupTestDb(t *testing.T) (*sql.DB, *InviteService) {

	t.Helper()

	config, err := config.Load("../../config/config.yaml")
	if err != nil {
		t.Fatalf("falha ao carregar a configuração: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		t.Fatalf("falha ao conectar ao banco de dados: %v", err)
	}

	producer := kafka.NewWriter(kafka.WriterConfig{Brokers: []string{config.Messaging.Brokers}, Topic: config.Messaging.Topic, Balancer: &kafka.LeastBytes{}})

	inviteRepository := repository.NewInviteRepository(db)
	kafkaRepository := repository.NewKafkaRepository(producer)

	inviteService := NewInviteService(inviteRepository, kafkaRepository)

	return db, inviteService

}

func newPostgres(dbConfig config.Database) string {
	return "user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.Name +
		" host=" + dbConfig.Host +
		" port=" + dbConfig.Port +
		" sslmode=disable"
}

func mockDriver() *models.Driver {
	return &models.Driver{
		Name:  "Micael Anderson",
		Email: "MichaelAStephens@rhyta.com ",
		CNH:   "86859349950",
	}
}

func mockSchool() *models.School {
	return &models.School{
		Name:  "E.E Afonso Castellano",
		Email: "covasad938@gawte.com",
		CNPJ:  "33352566000131",
	}
}

func TestInviteDriver(t *testing.T) {

}

func TestReadInvite(t *testing.T) {

}

func TestFindAllInvitesDriverAccount(t *testing.T) {

}

func TestAcceptedInvite(t *testing.T) {

}

func TestDeclineInvite(t *testing.T) {

}

func TestIsEmployee(t *testing.T) {

	_, inviteService := setupTestDb(t)

	inviteMock := models.Invite{
		School: *mockSchool(),
		Driver: *mockDriver(),
		Status: "pending",
	}

	err := inviteService.IsEmployee(context.Background(), &inviteMock)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestCheckInviteEntities(t *testing.T) {

	_, inviteService := setupTestDb(t)

	inviteMock := models.Invite{
		School: *mockSchool(),
		Driver: *mockDriver(),
		Status: "pending",
	}

	err := inviteService.CheckInviteEntities(context.Background(), &inviteMock)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestCreatePartnet(t *testing.T) {

	_, inviteService := setupTestDb(t)

	inviteMock := models.Invite{
		School: *mockSchool(),
		Driver: *mockDriver(),
		Status: "pending",
	}

	err := inviteService.CreatePartner(context.Background(), &inviteMock)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}
