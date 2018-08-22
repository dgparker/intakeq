package intakeq

import (
	"testing"
)

func TestQueryForms(t *testing.T) {
	query := &FormQuery{
		"noname",
		"2018-01-01",
		"2018-02-01",
		"1",
		"false",
	}
	sum, err := tClient.client.QueryForms(query)
	if err != nil {
		t.Error(sum, err)
	}
}

func TestParseFormQuery(t *testing.T) {
	query := &FormQuery{
		"noname",
		"2018-01-01",
		"2018-02-01",
		"1",
		"false",
	}
	parseFormQuery(query)
}
