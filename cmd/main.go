package main

import (
	"fmt"
	"github.com/zapling/ubiquiti-firmware-notify/ubiquiti"
)

func main() {
	firmwares, err := ubiquiti.GetAvailableFirmware("edgerouter-x-sfp")
	if err != nil {
		fmt.Println(err)
		return
	}

	var items []ubiquiti.Firmware
	for _, firm := range firmwares {
		if firm.CategroySlug == "firmware" {
			items = append(items, firm)
		}
	}

	for _, item := range items {
		fmt.Println(item.DatePublished, item.Name)
	}

	// fmt.Printf("%+v", items[0])
}
