package adapterfacade

import "fmt"

type popper struct{}

func (p popper) on() {
	fmt.Println(`Popcorn Popper on`)
}

func (p popper) pop() {
	fmt.Println(`Popcorn Popper popping popcorn!`)
}

func (p popper) off() {
	fmt.Println(`Popcorn Popper off`)
}

type lights struct{}

func (l lights) dim(d int) {
	fmt.Printf("Theater Ceiling Lights dimming to %d%%\n", d)
}

func (l lights) on() {
	fmt.Println(`Theater Ceiling Lights on`)
}

type screen struct{}

func (s screen) down() {
	fmt.Println(`Theater Screen going down`)
}
func (s screen) up() {
	fmt.Println(`theater Screen going up`)
}

type projector struct{}

func (p projector) on() {
	fmt.Println(`Top-O-Line Projector on`)
}

func (p projector) wideScreenMode() {
	fmt.Println(`Top-O-Line Projector in widescreen mode（16*9 aspect ratio）`)
}

func (p projector) off() {
	fmt.Println(`Top-O-Line Projector off`)
}

type amplifier struct{}

func (amp amplifier) on() {
	fmt.Println(`Top-O-Line Amplifier on`)
}
func (amp amplifier) setDvd() {
	fmt.Println(`Top-O-Line Amplifier setting DVD player to Top-O-Line DVD Player`)
}
func (amp amplifier) setSurroundSound() {
	fmt.Println(`Top-O-Line Amplifier surround sound on（5 speakers, 1 subwoofer）`)
}
func (amp amplifier) setVolume(v int) {
	fmt.Println(`Top-O-Line Amplifier setting volume to`, v)
}
func (amp amplifier) off() {
	fmt.Println(`Top-O-Line Amplifier off`)
}

type dvd struct {
	movie string
}

func (d dvd) on() {
	fmt.Println(`Top-O-Line DVD Player on`)
}
func (d *dvd) play(movie string) {
	d.movie = movie
	fmt.Printf("Top-O-Line DVD Player playing \"%s\"\n", d.movie)
}
func (d dvd) stop() {
	fmt.Printf("Top-O-Line DVD Player stopped \"%s\"\n", d.movie)
}
func (d *dvd) eject() {
	d.movie = ""
	fmt.Println(`Top-O-Line DVD Player eject`)
}
func (d dvd) off() {
	fmt.Println(`Top-O-Line DVD Player off`)
}

type HomeTheaterFacade interface {
	watchMovie(string)
	endMovie()
}

type HomeTheater struct {
	amp   *amplifier
	light *lights
	d     *dvd
	proj  *projector
	scr   *screen
	pp    *popper
}

func NewHomeTheater(amp *amplifier, l *lights, d *dvd, proj *projector, scr *screen, pp *popper) *HomeTheater {
	return &HomeTheater{amp: amp, light: l, d: d, proj: proj, scr: scr, pp: pp}
}

func (h HomeTheater) watchMovie(movie string) {
	fmt.Println(`Get ready to watch a movie...`)
	h.pp.on()
	h.pp.pop()
	h.light.dim(10)
	h.scr.down()
	h.proj.on()
	h.proj.wideScreenMode()
	h.amp.on()
	h.amp.setDvd()
	h.amp.setSurroundSound()
	h.amp.setVolume(5)
	h.d.on()
	h.d.play(movie)
}

func (h HomeTheater) endMovie() {
	fmt.Println(`Shutting movie theater down...`)
	h.pp.off()
	h.light.on()
	h.scr.up()
	h.proj.off()
	h.amp.off()
	h.d.stop()
	h.d.eject()
	h.d.off()
}

var _ HomeTheaterFacade = (*HomeTheater)(nil)
