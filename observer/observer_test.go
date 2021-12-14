package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	wd := NewWeatherData()
	d1 := NewCurrentConditionDisplay(wd)
	d2 := NewStatisticDisplay(wd)
	wd.SetMeasurements(27.0, 59.1, 102.2)
	wd.SetMeasurements(28.0, 60.1, 100.2)
	wd.RemoveObserver(d2)
	wd.SetMeasurements(29.0, 61.1, 101.2)
	wd.RemoveObserver(d1)
	wd.SetMeasurements(30.0, 62.1, 104.2)
}
