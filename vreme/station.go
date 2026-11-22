// Paket vreme vsebuje tipe in funkcije za beleženje in poročanje o temperaturah vremenskih postaj.
package vreme

import (
	"fmt"
	"sort"
)

// Postaja predstavlja vremensko postajo z oznako, imenom, krajem in meritvami temperature.
type Postaja struct {
	oznaka  string
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo ustvari novo vremensko postajo.
func DodajPostajo(oznaka string, ime string, kraj string) *Postaja {
	return &Postaja{oznaka: oznaka, ime: ime, kraj: kraj}
}

// DodajMeritev doda meritev temperature postaji, če ta obstaja in je vrednost v dovoljenem območju.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, minTemp float64, maxTemp float64) {
	postaja, obstaja := postaje[oznaka]

	if !obstaja {
		fmt.Printf("Napaka: postaja %s ne obstaja.\n", oznaka)
		return
	}

	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Printf(
			"Napaka: temperatura %.1f °C ni v dovoljenem območju (%.1f °C do %.1f °C).\n",
			temperatura,
			minTemp,
			maxTemp,
		)
		return
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja

	fmt.Printf(
		"Meritev %.1f °C je bila dodana postaji %s.\n",
		temperatura,
		oznaka,
	)
}

// statistika vrne najnižjo, najvišjo in povprečno temperaturo postaje z zadostnim številom meritev.
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, obstaja := postaje[oznaka]

	if !obstaja || len(postaja.meritve) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja = postaja.meritve[0]
	najvisja = postaja.meritve[0]
	var vsota float64

	for _, temperatura := range postaja.meritve {
		if temperatura < najnizja {
			najnizja = temperatura
		}

		if temperatura > najvisja {
			najvisja = temperatura
		}

		vsota += temperatura
	}

	povprecna = vsota / float64(len(postaja.meritve))

	return najnizja, najvisja, povprecna, true
}

// IzpisMeritev izpiše vse postaje in njihove meritve, razvrščene po oznaki.
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")

	oznake := make([]string, 0, len(postaje))

	for oznaka := range postaje {
		oznake = append(oznake, oznaka)
	}

	sort.Strings(oznake)

	for _, oznaka := range oznake {
		postaja := postaje[oznaka]

		fmt.Printf(
			"%s - %s (%s): %v\n", oznaka, postaja.ime, postaja.kraj, postaja.meritve,
		)
	}
}

// IzpisPorocila izpiše povzetek temperatur in opis vremena za vsako postajo.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	oznake := make([]string, 0, len(postaje))

	for oznaka := range postaje {
		oznake = append(oznake, oznaka)
	}

	sort.Strings(oznake)

	for _, oznaka := range oznake {
		postaja := postaje[oznaka]

		najnizja, najvisja, povprecna, ok := statistika(postaje, oznaka, stMeritev)

		if !ok {
			fmt.Printf("%s: premalo meritev\n", postaja.ime)
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

		fmt.Printf(
			"%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n",
			postaja.ime, najnizja, najvisja, povprecna, opis,
		)
	}
}
