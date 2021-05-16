package bot

import (
	"errors"
	"fmt"
	"github.com/aiomonitors/blink-reservation/handlers"
	"github.com/aiomonitors/blink-reservation/helpers"
	"github.com/aiomonitors/blink-reservation/requests"
	"github.com/aiomonitors/blink-reservation/types"
	"github.com/aiomonitors/gologger"
	"github.com/nleeper/goment"
	"strconv"
)

var Header = map[string]string{
	"Host":"uzkvhe2t35.execute-api.us-west-2.amazonaws.com",
	"x-blink-app-version":"2.26.1",
	"Connection":"keep-alive",
	"x-api-key":"nPFjzhoIsK3GIP52G7Xqx9KCozCg6HmS8VbhSmYD",
	"Accept":"/",
	"User-Agent":"Blink Fitness/2.26.1 (iPhone; iOS 14.4; Scale/3.00)",
	"Accept-Language":"en-US;q=1",
	"Accept-Encoding":"gzip, deflate, br",
}
type Bot struct {
	CurrentToken string
	MinHours int
	MaxHours int
	Reservations []types.Registration
	BusinessUnitCode int
	Logger *gologger.Logger
	Webhook string
}

func CreateBot (token string, minHours int, maxHours int, businessUnitCode int, webhook string) Bot {
	bot := Bot{
		MinHours: minHours,
		MaxHours: maxHours,
		BusinessUnitCode: businessUnitCode,
		Webhook: webhook,
	}
	bot.Logger = gologger.NewLogger("Blink Bot")
	Header["Authorization"] = fmt.Sprintf("Bearer %v", token)

	refreshResponse, refreshError := requests.RefreshToken(token, Header)
	if refreshError != nil {
		panic(refreshError)
	}
	refreshStruct, refreshParseError := handlers.ParseRefreshResponse(refreshResponse)
	if refreshParseError != nil {
		panic(refreshParseError)
	}
	bot.CurrentToken = refreshStruct.Token
	Header["Authorization"] = fmt.Sprintf("Bearer %v", bot.CurrentToken)
	bot.Logger.Green("Got auth token: %s", refreshStruct.Token)

	reservationResponse, reservationResponseError := requests.GetUserReservations(bot.GetHeaders())
	if reservationResponseError != nil {
		panic(reservationResponseError)
	}
	doesStatusCodeMatch := helpers.CheckStatusCode(reservationResponse, []int{200, 201})
	if doesStatusCodeMatch != nil {
		panic(errors.New("token might be incorrect"))
	}
	reservationStruct, reservationParseError := handlers.ParseRegistrationResponse(reservationResponse)
	if reservationParseError != nil {
		panic(reservationParseError)
	}
	bot.Reservations = reservationStruct.Registrations
	bot.Logger.Green("Got user reservations (%v)", len(bot.Reservations))

	return bot
}

func (b *Bot) RefreshReservations() error {
	reservationResponse, reservationResponseError := requests.GetUserReservations(b.GetHeaders())
	if reservationResponseError != nil {
		panic(reservationResponseError)
	}
	doesStatusCodeMatch := helpers.CheckStatusCode(reservationResponse, []int{200, 201})
	if doesStatusCodeMatch != nil {
		defer b.HandleForbidden()
		return reservationResponseError
	}
	reservationStruct, reservationParseError := handlers.ParseRegistrationResponse(reservationResponse)
	if reservationParseError != nil {
		defer b.HandleForbidden()
		b.Logger.Red("Could not refresh reservations")
		return reservationParseError
	}
	b.Reservations = reservationStruct.Registrations
	b.Logger.Green("Got user reservations (%v)", len(b.Reservations))
	return nil
}

// Validators / Checkers
func (b *Bot) GomentMapForReservations() []*goment.Goment {
	reservationTSMap := make([]*goment.Goment, len(b.Reservations))
	for index, obj := range b.Reservations {
		reservationTSMap[index] = helpers.GetGomentFromTimestamp(obj.StartDateUTC)
	}
	return reservationTSMap
}

func (b *Bot) DoesReservationExistExactHour(date string) bool {
	if len(b.Reservations) == 0 {
		return false
	}
	dateObj := helpers.GetGomentFromTimestamp(date)

	reservationTSMap := b.GomentMapForReservations()

	for _, gomentObj := range reservationTSMap {
		if gomentObj.Format("M-D-YY") == dateObj.Format("M-D-YY") {
			if dateObj.Get("hours") == gomentObj.Get("hours") {
				return true
			}
		}
	}
	return false
}

func (b *Bot) DoesReservationExistForDay(date string) bool {
	if len(b.Reservations) == 0 {
		return false
	}
	dateObj := helpers.GetGomentFromTimestamp(date)

	reservationTSMap := b.GomentMapForReservations()

	for _, gomentObj := range reservationTSMap {
		if gomentObj.Format("M-D-YY") == dateObj.Format("M-D-YY") {
			//fmt.Println(gomentObj.Format("M-D-YY"),dateObj.Format("M-D-YY"))
			return true
		}
	}
	return false
}

func (b *Bot) HandleForbidden() error {
	refreshResponse, refreshError := requests.RefreshToken(b.CurrentToken, Header)
	if refreshError != nil {
		return refreshError
	}
	refreshStruct, refreshParseError := handlers.ParseRefreshResponse(refreshResponse)
	if refreshParseError != nil {
		return refreshParseError
	}
	b.CurrentToken = refreshStruct.Token
	b.Logger.Yellow("Refreshed token: %s", refreshStruct.Token)
	return nil
}

func (b *Bot) GetHeaders() map[string]string {
	headers := map[string]string{
		"Host":"uzkvhe2t35.execute-api.us-west-2.amazonaws.com",
		"x-blink-app-version":"2.26.1",
		"Connection":"keep-alive",
		"x-api-key":"nPFjzhoIsK3GIP52G7Xqx9KCozCg6HmS8VbhSmYD",
		"Accept":"/",
		"User-Agent":"Blink Fitness/2.26.1 (iPhone; iOS 14.4; Scale/3.00)",
		"Accept-Language":"en-US;q=1",
		"Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJnZW5kZXIiOnsiQ2hhcmFjdGVyaXN0aWNUeXBlSWQiOjEsIkNoYXJhY3RlcmlzdGljVHlwZU5hbWUiOiJHZW5kZXIiLCJWYWx1ZSI6Ik1hbGUiLCJDaGFyYWN0ZXJpc3RpY1ZhbHVlSWQiOjF9LCJneW1Mb2NhdGlvbiI6IjY4MyIsImhhc0dhdGVkQWNjZXNzIjp7InRyaWFsRW5hYmxlZCI6ZmFsc2UsInRyaWFsRXhwaXJlZCI6ZmFsc2V9LCJtZW1iZXJJZCI6IjMwMDA3MTc4ODMiLCJtZW1iZXJzaGlwTGV2ZWwiOiJCbHVlIiwibWVtYmVyc2hpcFN0YXJ0RGF0ZSI6IjIwMjEtMDMtMTFUMDA6MDA6MDAuMDAwMDAwMC0wNTowMCIsInVzZXJBY2Nlc3MiOiJGdWxsQWNjZXNzIiwiaWF0IjoxNjE5MzczNzExLCJleHAiOjE2MTkzODQ1MTF9.4DO-vpA3JDUt4esuvJTOpmIKOjm8diwhNCFIrw-yTwQ",
		"Accept-Encoding":"gzip, deflate, br",
	}

	headers["Authorization"] = fmt.Sprintf("Bearer %v", b.CurrentToken)
	return headers
}

func (b *Bot) GetAvailableReservationsForDay(days int) ([]types.Slot, error) {
	daysAheadTimestamp := helpers.GetDaysAhead(days)

	slotsResponse, slotsResponseError := requests.RequestSlotsForDay(strconv.Itoa(b.BusinessUnitCode), daysAheadTimestamp, b.GetHeaders())
	if slotsResponseError != nil {
		defer b.HandleForbidden()
		return nil, slotsResponseError
	}

	doesStatusCodeMatch := helpers.CheckStatusCode(slotsResponse, []int{200, 201})
	if doesStatusCodeMatch != nil {
		defer b.HandleForbidden()
		return nil, doesStatusCodeMatch
	}

	parsedResponse, parsedResponseError := handlers.ParseSlotsResponse(slotsResponse)
	if parsedResponseError != nil {
		return nil, parsedResponseError
	}

	slots := make([]types.Slot, 0)
	for _, slot := range parsedResponse.Slots {
		if slot.RemainingSpots > 0 {
			slots = append(slots, slot)
		}
	}

	return slots, nil
}

func (b *Bot) GetValidAvailableSlots(days int) ([]types.Slot, error) {
	totalAvailReservations, resFetchErr := b.GetAvailableReservationsForDay(days)
	if resFetchErr != nil {
		b.Logger.Red("Error fetching available slots, trying again")
		b.HandleForbidden()
		return b.GetAvailableReservationsForDay(days)
	}
	slots := make([]types.Slot, 0)

	for _, slot := range totalAvailReservations {
		dateObj := helpers.GetGomentFromTimestamp(slot.StartDateUTC)
		dateObjHours := dateObj.Get("hours")
		if dateObjHours >= b.MinHours && dateObjHours <= b.MaxHours {
			slots = append(slots, slot)
		}
	}

	return slots, nil
}

func (b *Bot) ReserveSlot(slot types.Slot) error {
	reservationResponse, reservationResponseError := requests.ReserveSlot(slot.EventInstanceID, b.GetHeaders())
	if reservationResponseError != nil {
		defer b.HandleForbidden()
		return reservationResponseError
	}
	parsedResponse, responseParseError := handlers.ParseReservationResponse(reservationResponse)
	if responseParseError != nil {
		defer b.HandleForbidden()
		return errors.New("Could not parse response")
	}
	if parsedResponse != nil && parsedResponse.Success == true {
		helpers.SuccessEmbed( slot.StartDateUTC, b.Webhook)
		return nil
	}
	return errors.New("Could not parse response")
}

func (b *Bot) AttemptReservation(days int) error {
	b.Logger.Blue("Attempting to reserve for %v days ahead", days)

	availableReservations, availReservationsError := b.GetValidAvailableSlots(days)
	if availReservationsError != nil {
		b.Logger.Red("Error getting available reservations for %v days ahead", days)
		defer b.HandleForbidden()
		return availReservationsError
	}

	if len(availableReservations) == 0 {
		b.Logger.Yellow("No reservations available for %v days ahead", days)
		return errors.New(fmt.Sprintf("No reservations available for %v days ahead", days))
	}

	reservationToAttempt := availableReservations[0]
	hasSameDayReservation := b.DoesReservationExistForDay(reservationToAttempt.StartDateUTC)
	if hasSameDayReservation {
		b.Logger.Yellow("Already has reservation for %v days ahead", days)
		return errors.New(fmt.Sprintf("Reservation already made for %v days ahead", days))
	}

	reserveSlotError := b.ReserveSlot(reservationToAttempt)
	if reserveSlotError != nil {
		b.Logger.Red("Error reserving slot %v", reserveSlotError)
		return reserveSlotError
	}
	b.Logger.Green("Reserved slot for %v days ahead!", days)
	defer b.RefreshReservations()
	return nil
}