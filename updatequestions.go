package intakeq

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
)

// UpdateQuestions this endpoint allows you to update answers to office use questions
func (c *Client) UpdateQuestions(intakeForm *IntakeForm) (IntakeForm, error) {
	payload, err := json.Marshal(intakeForm)
	if err != nil {
		return IntakeForm{}, err
	}

	req, err := c.NewRequest("POST", "/intakes", bytes.NewBuffer(payload))
	if err != nil {
		return IntakeForm{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return IntakeForm{}, err
	}
	if res.StatusCode != 200 {
		return IntakeForm{}, errors.New("Received http status code: " + strconv.Itoa(res.StatusCode))
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return IntakeForm{}, err
	}

	form := IntakeForm{}
	invalid := json.Unmarshal(data, &form)
	if invalid != nil {
		return IntakeForm{}, invalid
	}

	return form, nil
}
