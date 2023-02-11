package services

import (
	"bytes"
	"net/http"

	"Liup.Gateway.Golang/src/models"
	"Liup.Gateway.Golang/src/utils"
)

const APIBaseURL = ""

func SendRequest(user models.User) (interface{}, error) {
	apiURL := APIBaseURL + "/SignIn"

	requestBody, err := utils.EncodeJSON(user)
	dataBody := requestBody.MarshalTo(nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(dataBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var responseObject interface{}
	err = utils.DecodeJSON(response.Body, &responseObject)
	if err != nil {
		return nil, err
	}

	return responseObject, nil
}
