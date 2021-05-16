package helpers

import (
	"errors"
	"fmt"
	"github.com/aiomonitors/godiscord"
	"github.com/nleeper/goment"
	"net/http"
)

func GetCurrentTime() int64 {
	t, _ := goment.New()
	return t.ToUnix()
}

func GetStartOfDay() int64 {
	t, _ := goment.New()
	return t.StartOf("day").ToUnix()
}

func GetDaysAhead(days int) int64 {
	t, _ := goment.New()
	return t.Add(days, "days").StartOf("day").ToUnix()
}

func GetGomentFromTimestamp(timestamp string) *goment.Goment {
	t, _ := goment.New(timestamp)
	t.Local()
	return t
}

func CheckStatusCode(res *http.Response, statusCodes []int) error {
	for _, b := range statusCodes {
		if res.StatusCode == b {
			return nil
		}
	}
	return errors.New("status code does not match")
}

func GetTodayFormatted() string {
	t, _ := goment.New()
	t.Local()
	return t.Format("MMDDYYYY")
}


func SuccessEmbed(date string, webhook string) error {
	gomentFromDate := GetGomentFromTimestamp(date)
	emb := godiscord.NewEmbed("Reserved Blink slot", fmt.Sprintf("Time: **%v**", gomentFromDate.Format("M/D hh:mm a")), "")
	emb.SetThumbnail("https://media.discordapp.net/attachments/823712467569082399/841066280667971644/ripdullah.png")
	emb.SetColor("#52BE80")
	emb.SetFooter("Blink Reservations", "")
	emb.SendToWebhook(webhook)
	return nil
}