package handlers

import (
	"encoding/json"
	"github.com/aiomonitors/blink-reservation/requests"
	"github.com/aiomonitors/blink-reservation/types"
	"io/ioutil"
	"net/http"
)

func ParseSlotsResponse(res *http.Response) (*types.AvailEvents, error) {
	defer res.Body.Close()

	var resJSON types.AvailEvents
	body, bodyReadErr := ioutil.ReadAll(res.Body)
	if bodyReadErr != nil {
		return nil, nil
	}

	unmErr := json.Unmarshal(body, &resJSON)
	if unmErr != nil {
		return nil, nil
	}
	return &resJSON, nil
}

func ParseRefreshResponse (res *http.Response) (*types.RefreshResponse, error) {
	defer res.Body.Close()

	var resJSON types.RefreshResponse
	body, bodyReadErr := ioutil.ReadAll(res.Body)
	if bodyReadErr != nil {
		return nil, nil
	}
	unmErr := json.Unmarshal(body, &resJSON)
	if unmErr != nil {
		return nil, nil
	}
	return &resJSON, nil
}

func ParseRegistrationResponse (res *http.Response) (*types.MemberReservationResponse, error) {
	defer res.Body.Close()

	var resJSON types.MemberReservationResponse
	body, bodyReadErr := ioutil.ReadAll(res.Body)
	if bodyReadErr != nil {
		return nil, nil
	}

	unmErr := json.Unmarshal(body, &resJSON)
	if unmErr != nil {
		return nil, nil
	}
	return &resJSON, nil
}

func ParseReservationResponse (res *http.Response) (*requests.ReservationResponse, error) {
	defer res.Body.Close()

	var resJSON requests.ReservationResponse
	body, bodyReadErr := ioutil.ReadAll(res.Body)
	if bodyReadErr != nil {
		return nil, bodyReadErr
	}

	unmErr := json.Unmarshal(body, &resJSON)
	if unmErr != nil {
		return nil, unmErr
	}
	return &resJSON, nil
}