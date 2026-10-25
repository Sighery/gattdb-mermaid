package bt

// Client usage from
// https://github.com/tinygo-org/bluetooth/blob/v0.16.0/examples/discover/main.go

import (
	"fmt"
	"log"

	bt "tinygo.org/x/bluetooth"
)

var adapter = bt.DefaultAdapter

func init() {
	err := adapter.Enable()
	if err != nil {
		log.Fatalf("Failed to initialise BT adapter: %w", err)
	}
}

func Scan() {
	println("scanning...")
	err := adapter.Scan(func(adapter *bt.Adapter, result bt.ScanResult) {
		fmt.Printf("Found device: %s %d %s\n", result.Address.String(), result.RSSI, result.LocalName())
		for _, data := range result.ManufacturerData() {
			name, ok := ManufacturerName(data.CompanyID)
			if !ok {
				name = "UNKNOWN"
			}
			fmt.Printf("  Manufacturer 0x%04X - %s\n    Data: % X\n", data.CompanyID, name, data.Data)
		}
		fmt.Printf("  Raw advertisment bytes: % X\n", result.Bytes())
	})
	if err != nil {
		log.Fatalf("Failed to BT scan: %w", err)
	}
}

func ParseAddress(val string) (addr bt.Address, err error) {
	mac, err := bt.ParseMAC(val)
	if err != nil {
		return addr, err
	}

	addr.MACAddress = bt.MACAddress{MAC: mac}
	return addr, nil
}

func DeviceGattDb(address bt.Address) string {
	return ""
}
