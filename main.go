package main

import (
	"github.com/aiomonitors/blink-reservation/bot"
	"sync"
	"time"
)

func main() {
	botInstance := bot.CreateBot("", 17, 18, 683, "")
	botInstance.Logger.Green("Created bot")
	//fmt.Println(botInstance.DoesReservationExistForDay("2021-04-28T19:40:00"))
	//_, _ = pretty.Println(botInstance.Reservations[len(botInstance.Reservations)-1].StartDateUTC)
	//_, _ = pretty.Println(botInstance.GetValidAvailableSlots(2))
	var wg sync.WaitGroup
	for true {
		wg.Add(1)
		go func() {
			defer wg.Done()
			botInstance.AttemptReservation(1)
			botInstance.AttemptReservation(2)
			botInstance.AttemptReservation(3)
			botInstance.RefreshReservations()
		}()
		wg.Wait()
		time.Sleep(5 * time.Minute)
		continue
	}
}