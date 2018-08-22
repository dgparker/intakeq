package intakeq

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
)

// FormSummary represents the response data received from /intakes/summary
type FormSummary struct {
	ID                string
	ClientName        string
	ClientEmail       string
	ClientID          int
	Status            string
	DateCreated       int
	DateSubmitted     int
	QuestionnaireName string
	Practitioner      string
	PractitionerName  string
}

// FormQuery available query string parameters for /intakes/summary
type FormQuery struct {
	Client    string
	StartDate string
	EndDate   string
	Page      string
	All       string
}

// QueryForms Use this method to query client intake form summaries.
// The result set does not contain all the contents of the intake forms,
// but only their basic information (id, status, client info).
func (c *Client) QueryForms(query *FormQuery) ([]FormSummary, error) {
	queryString := parseFormQuery(query)

	req, err := c.NewRequest("GET", "/intakes/summary"+queryString, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("Received http status code: " + strconv.Itoa(res.StatusCode))
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	summary := []FormSummary{}
	invalid := json.Unmarshal(data, &summary)
	if invalid != nil {
		return nil, invalid
	}
	return summary, nil
}

func parseFormQuery(query *FormQuery) string {
	queryString := "?"
	queryString += "client=" + query.Client
	queryString += "&startDate=" + query.StartDate
	queryString += "&endDate=" + query.EndDate
	queryString += "&page=" + query.Page
	queryString += "&all=" + query.All

	return queryString
}
