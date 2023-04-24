package processor

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/NickGowdy/data-processor/client"
	"github.com/NickGowdy/data-processor/device"
)

type CodeProcessor struct {
	CodeRegistrationLimit int
	MaxConcurrentJobs     int
	BaseUrl               string
	Client                client.Client
	Device                chan device.Device
}

// Worker attempts to register a valid device via external API.
// If successfull, a RegisterDevice struct with it's Identifier and Code will be sent to the work channel.
//
// # Example
//
//	Identifier: 1CEB0080F074F750, Code: 4F750
//
// When an unexpected error occurs, return ctx.Err instead.
func (cp *CodeProcessor) Worker(ctx context.Context, work chan struct{}) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-work:
			saved, registeredDevice := registerDevice(cp.Client, cp.BaseUrl)
			if saved {
				cp.Device <- *registeredDevice
			}
		}
	}
}

func registerDevice(client client.Client, url string) (bool, *device.Device) {
	device := device.NewDevice()

	b := new(bytes.Buffer)
	reqBody := map[string]string{"Deveui": device.Identifier}

	err := json.NewEncoder(b).Encode(&reqBody)
	if err != nil {
		log.Print(err)
	}

	// TODO: call url and handle response
	// resp, err := client.Post(url, "application/json", b)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return true, device
}
