package source

import (
	"github.com/jedib0t/go-pretty/v6/progress"
	"time"
)

const (
	msgWidth     = 30
	trackerWidth = 10
)

func getProgress() progress.Writer {
	pw := progress.NewWriter()
	pw.SetMessageWidth(msgWidth)
	pw.SetTrackerLength(trackerWidth)
	pw.SetStyle(progress.StyleDefault)
	pw.Style().Colors = progress.StyleColorsExample
	pw.SetTrackerPosition(progress.PositionRight)
	pw.SetUpdateFrequency(time.Millisecond * 100)
	pw.Style().Visibility.ETAOverall = false
	pw.Style().Visibility.Percentage = false
	pw.Style().Options.TimeInProgressPrecision = time.Millisecond

	return pw
}

func appendTracker(pw progress.Writer, message string) *progress.Tracker {
	tracker := &progress.Tracker{
		Message: message,
		Units:   progress.UnitsDefault,
	}
	pw.AppendTracker(tracker)
	return tracker
}
