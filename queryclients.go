package intakeq

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

// ClientQuery available query string parameters for /intakes/clients.
type ClientQuery struct {
	Search           string
	Page             string
	IncludeProfile   string
	DateCreatedStart string
	DateCreatedEnd   string
}

// ClientSummary represents the response data received from /intakes/clients when the includeProfile field is not set, or set to false
type ClientSummary struct {
	Name         string
	Email        string
	Phone        string
	ClientNumber int
}

// ClientProfile represents the response data received from /intakes/clients when the includedProfile field is set to true
type ClientProfile struct {
	ClientID                            string
	Name                                string
	FirstName                           string
	LastName                            string
	MiddleName                          string
	Email                               string
	Phone                               string
	DateOfBirth                         int
	MaritalStatus                       string
	Gender                              string
	Tags                                []string
	Archived                            bool
	HomePhone                           string
	WorkPhone                           string
	MobilePhone                         string
	Address                             string
	UnitNumber                          string
	AdditionalInformation               string
	PrimaryInusranceCompany             string
	PrimaryInsurancePolicyNumber        string
	PrimaryInsuranceGroupNumber         string
	PrimaryInsuranceHolderName          string
	PrimaryInsuranceRelationship        string
	PrimaryInsuranceHolderDateOfBirth   int
	SecondaryInsuranceCompany           string
	SecondaryInsurancePolicyNumber      string
	SecondaryInsuranceGroupNumber       string
	SecondaryInsuranceHolderName        string
	SecondaryInsuranceRelationship      string
	SecondaryInsuranceHolderDateOfBirth int
	DateCreated                         int
	LastActivityDate                    int
	CustomFields                        []*CustomField
}

// CustomField represents the response data contained within includedProfile/customFields
type CustomField struct {
	FieldID string
	Text    string
	Value   string
}

// QueryClients returns a list of clients
func (c *Client) QueryClients(query *ClientQuery) ([]ClientSummary, []ClientProfile, error) {
	queryString := parseClientQuery(query)

	req, err := c.NewRequest("GET", "/clients"+queryString, nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	if res.StatusCode != 200 {
		return nil, nil, errors.New("Received http status code: " + strconv.Itoa(res.StatusCode))
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	if strings.ToLower(query.IncludeProfile) == "true" {
		profile := []ClientProfile{}
		err := json.Unmarshal(data, &profile)
		if err != nil {
			return nil, nil, err
		}
		return nil, profile, nil
	}

	summary := []ClientSummary{}
	invalid := json.Unmarshal(data, &summary)
	if invalid != nil {
		return nil, nil, invalid
	}
	return summary, nil, nil
}

func parseClientQuery(query *ClientQuery) string {
	queryString := "?"
	queryString += "search=" + query.Search
	queryString += "&page=" + query.Page
	queryString += "&includeProfile=" + query.IncludeProfile
	queryString += "&dateCreatedStart=" + query.DateCreatedStart
	queryString += "&dateCreatedEnd=" + query.DateCreatedEnd

	return queryString
}
