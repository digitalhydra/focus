package tui

import (
	"github.com/charmbracelet/huh"
)

var WeekDaysSelect = huh.NewMultiSelect[string]().
	Title("Pick the days of the week you want the schedule to run on.").
	Options(
		huh.NewOption("Monday", "mon"),
		huh.NewOption("Tuesday", "tue"),
		huh.NewOption("Wednesday", "wed"),
		huh.NewOption("Thuerday", "thu"),
		huh.NewOption("Friday", "fri"),
		huh.NewOption("Saturday", "sat"),
		huh.NewOption("Sunday", "sun"),
	).
	Limit(7)
