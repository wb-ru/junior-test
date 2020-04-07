package factorymethod

import (
	"strings"
	"testing"
)

func TestCreateAppleDeviceByName(t *testing.T) {
	currentPhone := "iPhone"
	currentLaptop := "MacBook"
	currentComputer := "iMac"
	currentWatch := "AppleWatch"
	var tests = []struct {
		name string
		want *string
	}{
		{"pHone", &currentPhone},
		{"LaPtOP", &currentLaptop},
		{"computER", &currentComputer},
		{"waTch", &currentWatch},
		{"Car", nil},
		{"microwave", nil},
	}
	for _, test := range tests {
		deviceFactory, ok := createAppleDeviceByName(test.name)
		if ok {
			device := deviceFactory.createDeviceByName()
			got := device.createDevice()
			if got != *test.want {
				t.Errorf("Incorrect %v device was created. Expected : %v, got : %v", strings.ToLower(test.name), *test.want, got)
			}
		} else {
			if test.want != nil {
				t.Errorf("createAppleDeviceByName(%v) didn't create a device, but it exists : %v", test.name, *test.want)
			}
		}
	}
}
