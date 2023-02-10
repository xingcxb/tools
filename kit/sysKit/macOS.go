// go:build darwin
package sysKit

import (
	"fmt"
	"github.com/dkorunic/iSMC/smc"
)

const (
	FanCount = "Fan Count" // 风扇个数
)

func OsInfo() {

}

// AllFansInfo 风扇所有信息
type AllFansInfo struct {
	FanCount int       `json:"fanCount"` // 风扇个数
	FanInfos []FanInfo `json:"fanInfos"` // 风扇信息
}

// FanInfo 单个风扇信息
type FanInfo struct {
	CurrentSpeed string `json:"currentSpeed"` // 当前转速
	MinimalSpeed string `json:"minimalSpeed"` // 最小转速
	MaximumSpeed string `json:"maximumSpeed"` // 最大转速
	TargetSpeed  string `json:"targetSpeed"`  // 目标转速
}

// GetMacOSFanSpeed 获取macOS风扇转速
func GetMacOSFanSpeed() AllFansInfo {
	osFans := smc.GetFans()
	allFansInfo := AllFansInfo{}
	// 获取风扇个数
	fansCount := osFans[FanCount].(map[string]interface{})
	allFansInfo.FanCount = int(fansCount["value"].(uint32))
	// 设置默认最低未一个风扇
	fansInfo := make([]FanInfo, 0)
	for fc := 1; fc <= allFansInfo.FanCount; fc++ {
		fanInfo := FanInfo{}
		// 当前转速
		cSpeedMap := osFans[fmt.Sprintf("Fan %v Current Speed", fc)].(map[string]interface{})
		fanInfo.CurrentSpeed = cSpeedMap["value"].(string)
		// 最小转速
		minSpeedMap := osFans[fmt.Sprintf("Fan %v Minimal Speed", fc)].(map[string]interface{})
		fanInfo.MinimalSpeed = minSpeedMap["value"].(string)
		// 最大转速
		maxSpeedMap := osFans[fmt.Sprintf("Fan %v Maximum Speed", fc)].(map[string]interface{})
		fanInfo.MaximumSpeed = maxSpeedMap["value"].(string)
		// 目标转速
		tSpeedMap := osFans[fmt.Sprintf("Fan %v Target Speed", fc)].(map[string]interface{})
		fanInfo.TargetSpeed = tSpeedMap["value"].(string)
		fansInfo = append(fansInfo, fanInfo)
	}
	allFansInfo.FanInfos = fansInfo
	return allFansInfo
}

// GetMacOSCpu 获取macOSCPU信息
func GetMacOSCpu() {
	fmt.Println(smc.GetCurrent())
}
