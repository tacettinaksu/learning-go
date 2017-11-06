package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

const inputPath = "./input"

func main() {
	for {
		d, _ := os.Open(inputPath)
		files, _ := d.Readdir(-1)
		for _, fi := range files {
			filePath := inputPath + "/" + fi.Name()
			f, _ := os.Open(filePath)
			data, _ := ioutil.ReadAll(f)
			f.Close()
			os.Remove(filePath)

			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))
				records, _ := reader.ReadAll()
				for _, r := range records {
					deathRateData := new(DeathRate)
					deathRateData.Country = r[0]
					deathRateData.RoadTraficDeath = r[1]
					deathRateData.DeathPer100_000, _ = strconv.ParseFloat(r[2], 64)

					fmt.Printf("Country: '%s', \t\t Total Death: '%s', \t\tDeath Per 100.000 person %.2f\n",
						deathRateData.Country, deathRateData.RoadTraficDeath, deathRateData.DeathPer100_000)
				}
			}(string(data))
		}

		d.Close()
		time.Sleep(100 * time.Millisecond)
	}
}

type DeathRate struct {
	Country         string
	RoadTraficDeath string
	DeathPer100_000 float64
}
