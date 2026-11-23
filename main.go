package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/davor-z/vremenske-postaje/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "postaje",
		Usage: "beleženje meritev vremenskih postaj",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Usage: "najnižja dovoljena temperatura meritve v °C",
				Value: -50,
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Usage: "najvišja dovoljena temperatura meritve v °C",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "najmanjše število meritev, potrebnih za izračun statistike",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float64("minTemp")
			maxTemp := cmd.Float64("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			postaje := make(map[string]vreme.Postaja)

			vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
			vreme.DodajPostajo(postaje, "MB01", "Maribor Tabor", "Maribor")
			vreme.DodajPostajo(postaje, "KP01", "Koper Luka", "Koper")
			vreme.DodajPostajo(postaje, "KR01", "Kredarica", "Triglav")

			for _, t := range []float64{18.2, 21.5, 19.8, 16.5} {
				vreme.DodajMeritev(postaje, "LJ01", t, minTemp, maxTemp)
			}
			for _, t := range []float64{12.4} {
				vreme.DodajMeritev(postaje, "MB01", t, minTemp, maxTemp)
			}
			for _, t := range []float64{24.5, 27, 26.1, 28.4} {
				vreme.DodajMeritev(postaje, "KP01", t, minTemp, maxTemp)
			}
			for _, t := range []float64{2, -1.5, 3.5} {
				vreme.DodajMeritev(postaje, "KR01", t, minTemp, maxTemp)
			}

			vreme.IzpisMeritev(postaje)
			fmt.Println()

			vreme.DodajMeritev(postaje, "MB01", 14.1, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ01", -75, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "CE01", 10, minTemp, maxTemp)
			fmt.Println()

			vreme.IzpisMeritev(postaje)
			fmt.Println()

			vreme.IzpisPorocila(postaje, stMeritev)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
