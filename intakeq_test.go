package intakeq

import (
	"flag"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var tClient struct {
	client Client
}

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tClient.client.apikey = os.Getenv("apikey")
	tClient.client.httpClient = &http.Client{}

	flag.Parse()
	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestSetClient(t *testing.T) {
	client, err := SetClient(tClient.client.apikey, tClient.client.httpClient)
	if err != nil {
		t.Error(client, err)
	}
}

func TestNewRequest(t *testing.T) {
	req, err := tClient.client.NewRequest("GET", "/testing", nil)
	if err != nil {
		t.Error(req, err)
	}
}
