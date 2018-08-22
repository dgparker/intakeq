package intakeq

import (
	"os"
	"testing"
)

func TestUpdateQuestions(t *testing.T) {
	form, err := tClient.client.GetForm(os.Getenv("intakeid"))
	if err != nil {
		t.Error(err)
	}

	newForm, err := tClient.client.UpdateQuestions(&form)
	if err != nil {
		t.Error(newForm, err)
	}
}
