// Paket main je vstopna točka CLI programa, ki zažene simulacijo vremenskih postaj.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hlyudzyk/weatherstation/vreme"
	"github.com/urfave/cli/v3"
)

// main prebere argumente ukazne vrstice in zažene simulacijo.
func main() {
	cmd := &cli.Command{
		Name:  "weatherStation",
		Usage: "Weather station displays Temperature, Humidity, and Pressure",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Usage: "Minimalna temperatura",
				Value: -50,
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Usage: "Maksimalna temperatura",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "Minimalno stevilo meritev",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float64("minTemp")
			maxTemp := cmd.Float64("maxTemp")
			stMeritev := cmd.Int("stMeritev")
			return runSimilation(minTemp, maxTemp, stMeritev)
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}

// dodajMertive doda več meritev temperature podani postaji.
func dodajMertive(postaje map[string]vreme.Postaja, oznaka string, meritve []float64, minTemp float64, maxTemp float64) {
	for _, meritev := range meritve {
		vreme.DodajMeritev(postaje, oznaka, meritev, minTemp, maxTemp)
	}
}

// runSimilation ustvari vzorčne postaje, doda meritve in izpiše rezultate.
func runSimilation(minTemp float64, maxTemp float64, stMeritev int) error {
	postaje := make(map[string]vreme.Postaja)

	postaje["LJ01"] = *vreme.DodajPostajo(
		"LJ01",
		"Ljubljana Bežigrad",
		"Ljubljana",
	)
	dodajMertive(postaje, "LJ01", []float64{18.2, 21.5, 19.8, 16.5}, minTemp, maxTemp)
	dodajMertive(postaje, "LJ01", []float64{18.2, 21.5, 19.8, 16.5}, minTemp, maxTemp)
	dodajMertive(postaje, "LJ01", []float64{18.2, 21.5, 19.8, 16.5}, minTemp, maxTemp)

	postaje["KP01"] = *vreme.DodajPostajo(
		"KP01",
		"Koper Luka",
		"Koper",
	)
	dodajMertive(postaje, "KP01", []float64{24.2, 27, 26.3, 28.1}, minTemp, maxTemp)

	postaje["MB01"] = *vreme.DodajPostajo(
		"MB01",
		"Maribor Tabor",
		"Maribor",
	)
	dodajMertive(postaje, "MB01", []float64{12.4}, minTemp, maxTemp)

	postaje["KR01"] = *vreme.DodajPostajo(
		"KR01",
		"Kredarica",
		"Triglav",
	)
	dodajMertive(postaje, "KR01", []float64{2, -1.5, 3.4}, minTemp, maxTemp)

	fmt.Println("Dodajanje meritve:")
	vreme.DodajMeritev(postaje, "LJ01", 20.3, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "LJ01", 55, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "CE01", 15.5, minTemp, maxTemp)

	// fmt.Println("\nStatistika postaje LJ01:")

	// najnizja, najvisja, povprecna, ok := statistika(postaje, "LJ01")

	// if ok {
	// 	fmt.Printf(
	// 		"Najnižja: %.1f °C, najvišja: %.1f °C, povprečna: %.1f °C\n", najnizja, najvisja, povprecna,
	// 	)
	// } else {
	// 	fmt.Println("Statistike ni mogoče izračunati.")
	// }

	fmt.Println()
	vreme.IzpisMeritev(postaje)

	fmt.Println()
	vreme.IzpisPorocila(postaje, stMeritev)
	return nil
}
