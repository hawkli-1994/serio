package serio

import "go.bug.st/serial"

func ListPorts() ([]string, error) {
    return serial.GetPortsList()
}