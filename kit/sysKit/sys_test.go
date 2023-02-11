package sysKit

import (
	"fmt"
	"testing"
)

func TestFan(t *testing.T) {
	fmt.Println(GetMacOSFanSpeed())
}

func TestOsInfo(t *testing.T) {
	OsInfo()
}

func TestCpu(t *testing.T) {
	GetMacOSCpu()
}
