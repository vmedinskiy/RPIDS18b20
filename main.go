package main

import (
	"fmt"
	"time"

	"github.com/yryz/ds18b20"
)

func main() {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	fmt.Printf("sensor IDs: %v\n", sensors)

	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			for _, sensor := range sensors {
				t, err := ds18b20.Temperature(sensor)
				if err == nil {
					fmt.Printf("sensor: %s temperature: %.4f°C\n", sensor, t)
				}
			}
		}
	}
}

// env GOOS=linux GOARCH=arm GOARM=5 go build
