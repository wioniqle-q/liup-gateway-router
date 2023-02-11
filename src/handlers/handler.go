package handlers

import (
	"encoding/json"
	"net/http"

	"Liup.Gateway.Golang/src/models"
	"Liup.Gateway.Golang/src/services"
	"Liup.Gateway.Golang/src/utils"
	"github.com/valyala/fastjson"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	validate := utils.ValidateRequest(r)
	if validate != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request method")
		return
	}

	var user models.User
	err := utils.DecodeJSON(r.Body, &user)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Error decoding request body")
		return
	}

	if user.UserName == "" {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "UserName is a required field")
		return
	}

	response, err := services.SendRequest(user)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Error sending request to .NET Core API")
		return
	}

	var dataBytes []byte
	dataBytes, err = json.Marshal(response)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := fastjson.MustParseBytes(dataBytes)

	utils.WriteJSON(w, http.StatusOK, data)
}
