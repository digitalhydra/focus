package tui

import (
	print "focus/cmd/logger"
	"time"

	"github.com/charmbracelet/huh"
)

var times []string = []string{
	"0000", "0030", "0100", "0130", "0200", "0230", "0300", "0330",
	"0400", "0430", "0500", "0530", "0600", "0630", "0700", "0730",
	"0800", "0830", "0900", "0930", "1000", "1030", "1100", "1200",
	"1230",
}

func GetTimes() {
	startTime, err := time.Parse(time.RFC3339, "2000-01-01T00:00:00Z")
	if err != nil {
		print.Logger().Error("time parse error", "wtf", err)
	}
	print.Logger().Info(startTime.Clock())
	nextTime := startTime
	for nextTime.Hour() < 23 {

		nextTime = nextTime.Add(time.Minute * 30)
		print.Logger().Info(nextTime.Format("03:00PM"))
	}
}

var StartTimeSelect = huh.NewSelect[string]().
	Title("Pick the time of the schedule should start").
	Options(
		huh.NewOption("", ""),
	)
