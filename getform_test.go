package intakeq

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetForm(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	httpClient := &http.Client{}
	client, err := SetClient(os.Getenv("apikey"), httpClient)
	if err != nil {
		t.Errorf("client failed to initialize")
	}
	form, err := client.GetForm(os.Getenv("intakeid"))
	if err != nil {
		t.Error(form, err)
	}
}
