package sysKit

import (
	"fmt"
	"testing"
)

func TestFan(t *testing.T) {
	fmt.Println(GetMacOSFanSpeed())
}

func TestCpu(t *testing.T) {
	GetMacOSCpu()
}
