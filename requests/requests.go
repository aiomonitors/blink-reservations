package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aiomonitors/blink-reservation/helpers"
	"net/http"
)

type RefreshBody struct {
	RefreshToken string `json:"refreshToken"`
	ForceUpdate bool `json:"forceUpdate"`
}

type ReservationBody struct {
	EventID int64 `json:"eventInstanceId"`
}

type ReservationResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	EventID int64 `json:"eventInstanceId"`
	RegistrationID int64 `json:"eventInstanceRegistrationId"`
}

func RequestSlotsForDay(businessUnitCode string, timestamp int64, headers map[string]string) (*http.Response, error) {
	formattedURL := fmt.Sprintf("https://uzkvhe2t35.execute-api.us-west-2.amazonaws.com/prod/reservations/slots?businessUnitCode=%v&date=%v", businessUnitCode, timestamp)
	req, reqErr := http.NewRequest("GET", formattedURL, nil)
	if reqErr != nil {
		return nil, reqErr
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, resErr := http.DefaultClient.Do(req)
	return res, resErr
}

func RefreshToken(token string, headers map[string]string) (*http.Response, error) {
	rawBody := RefreshBody{
		RefreshToken: token,
		ForceUpdate: false,
	}

	resBody, _ := json.Marshal(rawBody)
	req, reqErr := http.NewRequest("POST", "https://uzkvhe2t35.execute-api.us-west-2.amazonaws.com/prod/v2/auth/refresh", bytes.NewBuffer(resBody))
	if reqErr != nil {
		return nil, nil
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, resErr := http.DefaultClient.Do(req)
	return res, resErr
}

func GetUserReservations(headers map[string]string) (*http.Response, error) {
	req, reqErr := http.NewRequest("GET", fmt.Sprintf("https://uzkvhe2t35.execute-api.us-west-2.amazonaws.com/prod/reservations/member?fromDate=%v", helpers.GetTodayFormatted()), nil)
	if reqErr != nil {
		return nil, nil
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, resErr := http.DefaultClient.Do(req)
	return res, resErr
}

func ReserveSlot(eventID int64, headers map[string]string) (*http.Response, error) {
	rawBody := ReservationBody{
		eventID,
	}

	reqBody, _ := json.Marshal(rawBody)
	req, reqErr := http.NewRequest("POST", "https://uzkvhe2t35.execute-api.us-west-2.amazonaws.com/prod/reservations/register", bytes.NewBuffer(reqBody))
	if reqErr != nil {
		return nil, nil
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, resErr := http.DefaultClient.Do(req)
	return res, resErr
}