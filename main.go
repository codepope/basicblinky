package main

import (
	"image/color"
	"machine"
	"math/rand"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var neo = machine.Pin(0x08)

var leds [25]color.RGBA

const blinkyspeed = 100

func main() {

	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ws := ws2812.New(8)
	rg := false

	for {
		rg = !rg
		for i := range leds {
			rg = !rg
			if rg {
				// Alpha channel is not supported by WS2812 so we leave it out
				leds[i] = color.RGBA{R: uint8(rand.Intn(64)), G: uint8(rand.Intn(64)), B: uint8(rand.Intn(64))}
			} else {
				leds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
			}
		}

		ws.WriteColors(leds[:])
		time.Sleep(blinkyspeed * time.Millisecond)
	}
}
