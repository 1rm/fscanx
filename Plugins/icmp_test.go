package Plugins

import (
	"testing"
)

func TestArrayCountValueTop(t *testing.T) {
	ipArray := []string{
		"192.168.1.1",
		"192.168.100.2",
		"192.168.100.3",
	}

	ArrayCountValueTop(ipArray, 10)
}
