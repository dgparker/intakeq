package intakeq

import (
	"os"
	"testing"
)

func TestDownloadPDF(t *testing.T) {
	_, err := tClient.client.DownloadPDF(os.Getenv("intakeid"))
	if err != nil {
		t.Error(err)
	}
}
