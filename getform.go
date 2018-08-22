package intakeq

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
)

// IntakeForm represents the response data received from /intakes/:id
type IntakeForm struct {
	ID                string
	ClientName        string
	ClientID          int
	Status            string
	DateCreated       float64
	DateSubmitted     float64
	QuestionnaireNmae string
	Practitioner      string
	PractitionerName  string
	Questions         []Question
	AppointmentID     string
}

// Question represents the IntakeQ question object
type Question struct {
	ID           string
	Text         string
	Answer       string
	QuestionType string
	Attachments  []Attachment
	Rows         []Row
	ColumnNames  []string
	OfficeUse    bool
	OfficeNote   bool
}

// Attachment represents the data associated with the questionType "Attachment"
type Attachment struct {
	ID          string
	URL         string
	ContentType string
	FileName    string
}

// Row represents the text and answers associated with the questionType "Matrix"
type Row struct {
	Text    string
	Answers []string
}

// GetForm Use this method to query client intake form summaries.
// The result set does not contain all the contents of the intake forms,
// but only their basic information (id, status, client info).
func (c *Client) GetForm(id string) (IntakeForm, error) {
	if id == "" {
		return IntakeForm{}, errors.New("received nil or zero value in required arguments")
	}

	req, err := c.NewRequest("GET", "/intakes/"+id, nil)
	if err != nil {
		return IntakeForm{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return IntakeForm{}, err
	}

	if res.StatusCode != 200 {
		return IntakeForm{}, errors.New("Received status code: " + strconv.Itoa(res.StatusCode))
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return IntakeForm{}, err
	}

	form := IntakeForm{}
	invalid := json.Unmarshal(data, &form)
	if invalid != nil {
		return IntakeForm{}, err
	}
	return form, nil
}
