package observer

import (
	"fmt"
)

type Observer interface {
	Update(temp, humidity, pressure float64)
}

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	notifyObserver()
}

type DisplayElement interface {
	Display()
}

type WeatherData struct {
	observers                map[Observer]struct{}
	temp, humidity, pressure float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		observers: make(map[Observer]struct{}),
	}
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.observers[o] = struct{}{}
}

func (w *WeatherData) RemoveObserver(o Observer) {
	delete(w.observers, o)
}

func (w *WeatherData) notifyObserver() {
	for o := range w.observers {
		o.Update(w.temp, w.humidity, w.pressure)
	}
}

func (w *WeatherData) MeasurementsChanged() {
	w.notifyObserver()
}

func (w *WeatherData) SetMeasurements(temp, humidity, pressure float64) {
	w.temp = temp
	w.humidity = humidity
	w.pressure = pressure
	w.MeasurementsChanged()
}

var _ Subject = (*WeatherData)(nil)

type CurrentConditionDisplay struct {
	temp, humidity, pressure float64
	pressureUp               bool
	weatherData              Subject
}

func NewCurrentConditionDisplay(wd Subject) *CurrentConditionDisplay {
	cur := &CurrentConditionDisplay{weatherData: wd}
	wd.RegisterObserver(cur)
	return cur
}

func (d *CurrentConditionDisplay) Update(temp, humidity, pressure float64) {
	d.temp = temp
	d.humidity = humidity
	d.pressureUp = pressure >= d.pressure
	d.pressure = pressure
	d.Display()
}

func (d CurrentConditionDisplay) Display() {
	fmt.Println("Current conditions:", d.temp, "F degrees and ", d.humidity, "% humidity and pressure upOrNot:", d.pressureUp)
}

var _ Observer = (*CurrentConditionDisplay)(nil)
var _ DisplayElement = (*CurrentConditionDisplay)(nil)

type StatisticDisplay struct {
	count                       int
	minTemp, maxTemp, totalTemp float64
	weatherData                 Subject
}

func NewStatisticDisplay(wd Subject) *StatisticDisplay {
	cur := &StatisticDisplay{
		minTemp:     1000,
		maxTemp:     -1000,
		weatherData: wd,
	}
	wd.RegisterObserver(cur)
	return cur
}

func (d *StatisticDisplay) Update(temp, humidity, pressure float64) {
	if temp < d.minTemp {
		d.minTemp = temp
	}
	if temp > d.maxTemp {
		d.maxTemp = temp
	}
	d.totalTemp += temp
	d.count += 1
	d.Display()
}

func (d StatisticDisplay) Display() {
	fmt.Println("Statistic conditions: minTemp(", d.minTemp, "F degrees) and maxTemp(", d.maxTemp, "F degress) and avgTemp(", d.totalTemp/float64(d.count), ")")
}

var _ Observer = (*StatisticDisplay)(nil)
var _ DisplayElement = (*StatisticDisplay)(nil)
