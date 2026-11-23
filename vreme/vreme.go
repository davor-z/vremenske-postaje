package vreme

import (
	"fmt"
	"sort"
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	if _, ok := postaje[oznaka]; ok {
		fmt.Printf("Napaka: postaja %s že obstaja.\n", oznaka)
		return
	}
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
}

func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	p, ok := postaje[oznaka]
	if !ok {
		fmt.Printf("Napaka: postaja %s ne obstaja.\n", oznaka)
		return
	}
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Printf("Napaka: meritev %.1f °C za postajo %s ni veljavna.\n", temperatura, oznaka)
		return
	}
	p.meritve = append(p.meritve, temperatura)
	postaje[oznaka] = p
}


func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	p, obstaja := postaje[oznaka]
	if !obstaja || len(p.meritve) == 0 || len(p.meritve) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja, najvisja = p.meritve[0], p.meritve[0]
	vsota := 0.0
	for _, t := range p.meritve {
		if t < najnizja {
			najnizja = t
		}
		if t > najvisja {
			najvisja = t
		}
		vsota += t
	}
	povprecna = vsota / float64(len(p.meritve))
	return najnizja, najvisja, povprecna, true
}


func urejeneOznake(postaje map[string]Postaja) []string {
	oznake := make([]string, 0, len(postaje))
	for o := range postaje {
		oznake = append(oznake, o)
	}
	sort.Strings(oznake)
	return oznake
}


func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for _, o := range urejeneOznake(postaje) {
		p := postaje[o]
		fmt.Printf("%s - %s (%s): %v\n", o, p.ime, p.kraj, p.meritve)
	}
}


func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	fmt.Println("POROČILO:")
	for _, o := range urejeneOznake(postaje) {
		p := postaje[o]
		najnizja, najvisja, povprecna, ok := statistika(postaje, o, stMeritev)
		if !ok {
			fmt.Printf("%s: premalo meritev\n", p.ime)
			continue
		}

		var opis string
		switch {
		case povprecna >= 25:
			opis = "Vroče"
		case povprecna >= 15:
			opis = "Toplo"
		case povprecna >= 5:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}
		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", p.ime, najnizja, najvisja, povprecna, opis)
	}
}
