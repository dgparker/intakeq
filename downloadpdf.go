package intakeq

import (
	"errors"
	"io/ioutil"
	"strconv"
)

// DownloadPDF Use this method to download a client’s complete intake package as a PDF file.
func (c *Client) DownloadPDF(id string) ([]byte, error) {
	if id == "" {
		return nil, errors.New("id parameter cannot be empty")
	}

	req, err := c.NewRequest("GET", "/intakes/"+id+"/pdf", nil)
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

	return data, nil
}
