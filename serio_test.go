package serio

import (
	"testing"
	"time"
)

func TestConfigStruct(t *testing.T) {
	config := Config{
		Name:     "/dev/null",
		Baud:     9600,
		DataBits: 8,
		StopBits: 1,
		Parity:   None,
		Timeout:  time.Second,
	}

	if config.Name != "/dev/null" {
		t.Errorf("Expected Name to be '/dev/null', got '%s'", config.Name)
	}

	if config.Baud != 9600 {
		t.Errorf("Expected Baud to be 9600, got %d", config.Baud)
	}

	if config.DataBits != 8 {
		t.Errorf("Expected DataBits to be 8, got %d", config.DataBits)
	}

	if config.StopBits != 1 {
		t.Errorf("Expected StopBits to be 1, got %d", config.StopBits)
	}

	if config.Parity != None {
		t.Errorf("Expected Parity to be None, got %v", config.Parity)
	}

	if config.Timeout != time.Second {
		t.Errorf("Expected Timeout to be 1s, got %v", config.Timeout)
	}
}

func TestPortMethods(t *testing.T) {
	port := &Port{}
	
	// Test SetWriteTimeout
	err := port.SetWriteTimeout(time.Second)
	if err != nil {
		t.Errorf("SetWriteTimeout failed: %v", err)
	}
	
	// Test SetDeadline
	err = port.SetDeadline(time.Now().Add(time.Second))
	if err != nil {
		t.Errorf("SetDeadline failed: %v", err)
	}
}