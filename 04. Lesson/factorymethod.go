package factorymethod

import "strings"

type DeviceFactory interface {
	createDeviceByName() Device
}

type PhoneFactory struct {
}

func (p *PhoneFactory) createDeviceByName() Device {
	return &Phone{}
}

type WatchFactory struct {
}

func (w WatchFactory) createDeviceByName() Device {
	return &Watch{}
}

type LaptopFactory struct {
}

func (l *LaptopFactory) createDeviceByName() Device {
	return &Laptop{}
}

type ComputerFactory struct {
}

func (c *ComputerFactory) createDeviceByName() Device {
	return &Computer{}
}

type Device interface {
	createDevice() string
}

type Phone struct{}

func (p *Phone) createDevice() string {
	return "iPhone"
}

type Watch struct{}

func (w *Watch) createDevice() string {
	return "AppleWatch"
}

type Laptop struct{}

func (l *Laptop) createDevice() string {
	return "MacBook"
}

type Computer struct{}

func (c *Computer) createDevice() string {
	return "iMac"
}

func createAppleDeviceByName(name string) (DeviceFactory, bool) {
	switch strings.ToLower(name) {
	case "phone":
		return &PhoneFactory{}, true
	case "laptop":
		return &LaptopFactory{}, true
	case "computer":
		return &ComputerFactory{}, true
	case "watch":
		return &WatchFactory{}, true
	default:
		return nil, false
	}
}
