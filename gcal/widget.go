package gcal

import (
	"fmt"
	"strings"
	"time"

	"github.com/rivo/tview"
	"google.golang.org/api/calendar/v3"
)

func Widget() tview.Primitive {
	events := Fetch()

	widget := tview.NewTextView()
	widget.SetBorder(true)
	widget.SetDynamicColors(true)
	widget.SetTitle(" 🐸 Calendar ")

	data := ""
	for _, item := range events.Items {
		ts, _ := time.Parse(time.RFC3339, item.Start.DateTime)
		timestamp := ts.Format("Mon Jan _2 15:04:05 2006")

		color := colorize(item, ts)

		str := fmt.Sprintf(" [%s]%s[white]\n %s\n\n", color, item.Summary, timestamp)
		data = data + str
	}

	fmt.Fprintf(widget, "%s", data)

	return widget
}

func colorize(item *calendar.Event, ts time.Time) string {
	color := "red"
	if strings.Contains(item.Summary, "1on1") {
		color = "green"
	}

	if ts.Before(time.Now()) {
		color = "grey"
	}

	return color
}