package intakeq

import (
	"errors"
	"io"
	"net/http"
)

// Client implements interface API interface
type Client struct {
	apikey     string
	httpClient *http.Client
}

// API contains the methods for interacting with the intakeq api
type API interface {
	QueryForms(*FormQuery) (*FormSummary, error)
	DownloadPDF(id string) ([]byte, error)
	GetForm(id string) (*IntakeForm, error)
	UpdateQuestions(*IntakeForm) (bool, error)
	QueryClients(*ClientQuery) (ClientSummary, ClientProfile, error)
	NewRequest(method string, endpoint string, body io.Reader) (*http.Request, error)
}

// SetClient initialize the api client with your generated apikey and an httpClient.
// the client struct returned contains the methods for interacting with the intakeq api
func SetClient(apikey string, httpClient *http.Client) (*Client, error) {
	if apikey == "" || httpClient == nil {
		return nil, errors.New("Received zero value for apikey or httpClient")
	}
	return &Client{apikey, httpClient}, nil
}

// NewRequest is used to encapsulate all requests made to the intakeq api and adds the necessary headers.
func (c *Client) NewRequest(method string, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, "https://intakeq.com/api/v1"+endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Auth-Key", c.apikey)
	req.Header.Add("Content-Type", "Application/Json")
	return req, nil
}
