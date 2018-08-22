package intakeq

import (
	"testing"
)

func TestQueryClients(t *testing.T) {
	query := &ClientQuery{
		"noname",
		"1",
		"false",
		"2018-01-01",
		"2018-02-01",
	}
	sum, profile, err := tClient.client.QueryClients(query)
	if err != nil {
		t.Error(sum, profile, err)
	}
}

func TestParseClientQuery(t *testing.T) {
	query := &ClientQuery{
		"noname",
		"1",
		"false",
		"2018-01-01",
		"2018-02-01",
	}
	parseClientQuery(query)
}
