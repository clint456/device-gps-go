// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 Canonical Ltd
// Copyright (C) 2018-2019 IOTech Ltd
// Copyright (C) 2021 Jiangxing Intelligence Ltd
// Copyright (C) 2022 HCL Technologies Ltd
//
// SPDX-License-Identifier: Apache-2.0

// Package driver this package provides an UART implementation of
// ProtocolDriver interface.
//
// CONTRIBUTORS              COMPANY
//===============================================================
// 1. Sathya Durai           HCL Technologies
// 2. Sudhamani Bijivemula   HCL Technologies
// 3. Vediyappan Villali     HCL Technologies
// 4. Vijay Annamalaisamy    HCL Technologies
//
//

package driver

import (
	errorDefault "errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/edgexfoundry/device-sdk-go/v4/pkg/interfaces"
	dsModels "github.com/edgexfoundry/device-sdk-go/v4/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

type Driver struct {
	sdk       interfaces.DeviceServiceSDK
	lc        logger.LoggingClient
	asyncCh   chan<- *dsModels.AsyncValues
	deviceCh  chan<- []dsModels.DiscoveredDevice
	gpsDevice *LCX6XZ // GPS设备实例
}

// Initialize performs protocol-specific initialization for the device
// service.
func (s *Driver) Initialize(sdk interfaces.DeviceServiceSDK) error {
	s.sdk = sdk
	s.lc = sdk.LoggingClient()
	s.asyncCh = sdk.AsyncValuesChannel()
	s.deviceCh = sdk.DiscoveredDeviceChannel()

	s.lc.Info("🚀 初始化GPS设备服务")

	// 初始化GPS设备
	gpsDevice, err := InitLCX6XZ()
	if err != nil {
		s.lc.Errorf("❌ GPS设备初始化失败: %v", err)
		return err
	}

	s.gpsDevice = gpsDevice
	s.lc.Info("✅ GPS设备初始化成功")

	return nil
}

// Start runs device service startup tasks after the SDK has been completely
// initialized. This allows device service to safely use DeviceServiceSDK
// interface features in this function call
func (s *Driver) Start() error {
	return nil
}

// HandleReadCommands triggers a protocol Read operation for the specified device.
func (s *Driver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	s.lc.Debugf("📖 处理设备 %s 的读取命令", deviceName)

	if s.gpsDevice == nil {
		return nil, fmt.Errorf("GPS设备未初始化")
	}

	res = make([]*dsModels.CommandValue, 0, len(reqs))

	for _, req := range reqs {
		s.lc.Debugf("处理资源: %s", req.DeviceResourceName)

		var cv *dsModels.CommandValue

		switch req.DeviceResourceName {
		case "latitude":
			cv = s.getLatitude(req)
			s.lc.Debugf("latitude: %v", cv)
		case "longitude":
			cv = s.getLongitude(req)
		case "altitude":
			cv = s.getAltitude(req)
		case "speed":
			cv = s.getSpeed(req)
		case "course":
			cv = s.getCourse(req)
		case "utc_time":
			cv = s.getUTCTime(req)
		case "fix_quality":
			cv = s.getFixQuality(req)
		case "satellites_used":
			cv = s.getSatellitesUsed(req)
		case "hdop":
			cv = s.getHDOP(req)
		case "gps_status":
			cv = s.getGPSStatus(req)
		default:
			s.lc.Warnf("未知的资源名称: %s", req.DeviceResourceName)
			continue
		}

		if cv != nil {
			res = append(res, cv)
		}
	}

	return res, nil
}

// HandleWriteCommands passes a slice of CommandRequest struct each representing
// a ResourceOperation for a specific device resource.
// Since the commands are actuation commands, params provide parameters for the individual
// command.
func (s *Driver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest,
	params []*dsModels.CommandValue) error {
	s.lc.Debug("✍️ 写操作被触发")
	return nil
}

// Discover triggers protocol specific device discovery, asynchronously writes
// the results to the channel which is passed to the implementation via
// ProtocolDriver.Initialize()
func (s *Driver) Discover() error {
	return fmt.Errorf("Discover function is yet to be implemented!")

}

// ValidateDevice triggers device's protocol properties validation, returns error
// if validation failed and the incoming device will not be added into EdgeX
func (s *Driver) ValidateDevice(device models.Device) error {

	protocol, ok := device.Protocols["UART"]
	if !ok {
		return errorDefault.New("Missing 'UART' protocols")
	}

	deviceLocation, ok := protocol["deviceLocation"]
	if !ok {
		return errorDefault.New("Missing 'deviceLocation' information")
	} else if deviceLocation == "" {
		return errorDefault.New("deviceLocation must not empty")
	}

	baudRate, ok := protocol["baudRate"]
	if !ok {
		return errorDefault.New("Missing 'baudRate' information")
	} else if baudRate == "" {
		return errorDefault.New("baudRate must not empty")
	}

	return nil
}

// Stop the protocol-specific DS code to shutdown gracefully, or
// if the force parameter is 'true', immediately. The driver is responsible
// for closing any in-use channels, including the channel used to send async
// readings (if supported).
func (s *Driver) Stop(force bool) error {
	// Then Logging Client might not be initialized
	if s.lc != nil {
		s.lc.Debugf(fmt.Sprintf("Driver.Stop called: force=%v", force))
	}
	return nil
}

// AddDevice is a callback function that is invoked
// when a new Device associated with this Device Service is added
func (s *Driver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	s.lc.Debugf(fmt.Sprintf("a new Device is added: %s", deviceName))
	return nil
}

// UpdateDevice is a callback function that is invoked
// when a Device associated with this Device Service is updated
func (s *Driver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	s.lc.Debugf(fmt.Sprintf("Device %s is updated", deviceName))
	return nil
}

// RemoveDevice is a callback function that is invoked
// when a Device associated with this Device Service is removed
func (s *Driver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	s.lc.Debugf(fmt.Sprintf("Device %s is removed", deviceName))
	return nil
}

// GPS数据读取辅助方法

// getLatitude 获取纬度
func (s *Driver) getLatitude(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil || s.gpsDevice.NMEA_RMC == nil {
		return nil
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	lat := s.cleanString(string(s.gpsDevice.NMEA_RMC.Lat[:]))
	ns := s.cleanString(string(s.gpsDevice.NMEA_RMC.N_S[:]))

	if lat == "" || ns == "" {
		return nil
	}

	// 转换为十进制度数格式
	latValue, isValid := s.convertDMSToDecimalWithValidation(lat, ns)
	if !isValid {
		return nil
	}

	// 格式化为易读格式
	formattedLat := s.formatCoordinate(latValue, true, ns)
	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "String", formattedLat)
	return cv
}

// getLongitude 获取经度
func (s *Driver) getLongitude(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil || s.gpsDevice.NMEA_RMC == nil {
		return nil
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	lon := s.cleanString(string(s.gpsDevice.NMEA_RMC.Lon[:]))
	ew := s.cleanString(string(s.gpsDevice.NMEA_RMC.E_W[:]))

	if lon == "" || ew == "" {
		return nil
	}

	// 转换为十进制度数格式
	lonValue, isValid := s.convertDMSToDecimalWithValidation(lon, ew)
	if !isValid {
		return nil
	}

	// 格式化为易读格式
	formattedLon := s.formatCoordinate(lonValue, false, ew)
	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "String", formattedLon)
	return cv
}

// getAltitude 获取海拔高度
func (s *Driver) getAltitude(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil || s.gpsDevice.NMEA_GGA == nil {
		return nil
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	altStr := s.cleanString(string(s.gpsDevice.NMEA_GGA.Alt[:]))
	if altStr == "" {
		return nil
	}

	// 解析海拔高度
	altitude := s.parseFloat(altStr)

	// 格式化为易读格式
	formattedAlt := s.formatAltitude(altitude)
	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "String", formattedAlt)
	return cv
}

// getSpeed 获取速度
func (s *Driver) getSpeed(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil {
		return nil
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	var speedKmh float64

	var hasValidData bool

	// 优先从VTG获取km/h速度
	if s.gpsDevice.NMEA_VTG != nil {
		sogkStr := s.cleanString(string(s.gpsDevice.NMEA_VTG.SOGK[:]))
		if sogkStr != "" {
			speedKmh = s.parseFloat(sogkStr)
			hasValidData = true
		}
	}

	// 如果VTG中没有，从RMC获取节速度并转换
	if !hasValidData && s.gpsDevice.NMEA_RMC != nil {
		sogStr := s.cleanString(string(s.gpsDevice.NMEA_RMC.SOG[:]))
		if sogStr != "" {
			sog := s.parseFloat(sogStr)
			speedKmh = sog * 1.852 // 1节 = 1.852 km/h
			hasValidData = true
		}
	}

	// 只有在没有任何有效数据时才返回nil
	if !hasValidData {
		return nil
	}

	// 格式化为易读格式
	formattedSpeed := s.formatSpeed(speedKmh)
	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "String", formattedSpeed)
	return cv
}

// getCourse 获取航向
func (s *Driver) getCourse(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil {
		return nil
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	var course float64
	var hasValidData bool

	// 优先从VTG获取航向
	if s.gpsDevice.NMEA_VTG != nil {
		cogtStr := s.cleanString(string(s.gpsDevice.NMEA_VTG.COGT[:]))
		if cogtStr != "" {
			course = s.parseFloat(cogtStr)
			hasValidData = true
		}
	}

	// 如果VTG中没有，从RMC获取
	if !hasValidData && s.gpsDevice.NMEA_RMC != nil {
		cogStr := s.cleanString(string(s.gpsDevice.NMEA_RMC.COG[:]))
		if cogStr != "" {
			course = s.parseFloat(cogStr)
			hasValidData = true
		}
	}

	// 只有在没有任何有效数据时才返回nil
	if !hasValidData {
		return nil
	}

	// 格式化为易读格式
	formattedCourse := s.formatCourse(course)
	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "String", formattedCourse)
	return cv
}

// getUTCTime 获取UTC时间
func (s *Driver) getUTCTime(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil || s.gpsDevice.NMEA_RMC == nil {
		return nil
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	utcStr := s.cleanString(string(s.gpsDevice.NMEA_RMC.UTC[:]))
	if utcStr == "" {
		return nil
	}

	// 将UTC时间格式化为易读格式
	formattedTime := s.formatUTCTime(utcStr)
	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "String", formattedTime)
	return cv
}

// getFixQuality 获取定位质量
func (s *Driver) getFixQuality(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil {
		return nil
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	var quality int32

	// 优先从GGA获取详细的定位质量
	if s.gpsDevice.NMEA_GGA != nil {
		qualityStr := s.cleanString(string(s.gpsDevice.NMEA_GGA.Quality[:]))
		if qualityStr != "" {
			quality = int32(s.parseFloat(qualityStr))
		}
	}

	// 如果GGA中没有，从RMC状态推断
	if quality == 0 && s.gpsDevice.NMEA_RMC != nil {
		status := s.cleanString(string(s.gpsDevice.NMEA_RMC.Status[:]))
		if status == "A" {
			quality = 1 // 有效定位
		}
	}

	// 格式化为易读格式
	formattedQuality := s.formatFixQuality(quality)
	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "String", formattedQuality)
	return cv
}

// getSatellitesUsed 获取使用的卫星数
func (s *Driver) getSatellitesUsed(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil || s.gpsDevice.NMEA_GGA == nil {
		return nil
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	satStr := s.cleanString(string(s.gpsDevice.NMEA_GGA.NumSatUsed[:]))
	if satStr == "" {
		return nil
	}

	// 解析卫星数
	satCount := int32(s.parseFloat(satStr))

	// 格式化为易读格式
	formattedSatCount := s.formatSatelliteCount(satCount)
	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "String", formattedSatCount)
	return cv
}

// getHDOP 获取水平精度因子
func (s *Driver) getHDOP(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil {
		return nil
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	var hdopStr string

	// 优先从GGA获取HDOP
	if s.gpsDevice.NMEA_GGA != nil {
		hdopStr = s.cleanString(string(s.gpsDevice.NMEA_GGA.HDOP[:]))
	}

	// 如果GGA中没有，尝试从GSA获取
	if hdopStr == "" && s.gpsDevice.NMEA_GSA != nil {
		hdopStr = s.cleanString(string(s.gpsDevice.NMEA_GSA.HDOP[:]))
	}

	if hdopStr == "" {
		return nil
	}

	// 解析HDOP值
	hdop := s.parseFloat(hdopStr)

	// 格式化为易读格式
	formattedHDOP := s.formatHDOP(hdop)
	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "String", formattedHDOP)
	return cv
}

// getGPSStatus 获取GPS状态
func (s *Driver) getGPSStatus(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil || s.gpsDevice.NMEA_RMC == nil {
		cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "String", "DISCONNECTED")
		return cv
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	status := s.cleanString(string(s.gpsDevice.NMEA_RMC.Status[:]))
	var gpsStatus string
	if status == "A" {
		gpsStatus = "ACTIVE"
	} else {
		gpsStatus = "WARNING"
	}

	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "String", gpsStatus)
	return cv
}

// 工具函数

// convertDMSToDecimalWithValidation 将度分秒格式转换为十进制度数，并返回是否有效
func (s *Driver) convertDMSToDecimalWithValidation(dmsStr, direction string) (float64, bool) {
	if dmsStr == "" || direction == "" {
		return 0.0, false
	}

	// 清理字符串
	dmsStr = s.cleanString(dmsStr)
	if len(dmsStr) < 4 {
		return 0.0, false
	}

	// 解析度分格式 (ddmm.mmmm 或 dddmm.mmmm)
	var degrees, minutes float64
	var err error

	if strings.Contains(dmsStr, ".") {
		// 查找小数点位置
		dotIndex := strings.Index(dmsStr, ".")
		if dotIndex >= 4 {
			// 经度格式 dddmm.mmmm
			degrees, err = strconv.ParseFloat(dmsStr[:dotIndex-2], 64)
			if err != nil {
				return 0.0, false
			}
			minutes, err = strconv.ParseFloat(dmsStr[dotIndex-2:], 64)
			if err != nil {
				return 0.0, false
			}
		} else if dotIndex >= 3 {
			// 纬度格式 ddmm.mmmm
			degrees, err = strconv.ParseFloat(dmsStr[:dotIndex-2], 64)
			if err != nil {
				return 0.0, false
			}
			minutes, err = strconv.ParseFloat(dmsStr[dotIndex-2:], 64)
			if err != nil {
				return 0.0, false
			}
		} else {
			return 0.0, false
		}
	} else {
		return 0.0, false
	}

	decimal := degrees + minutes/60.0

	// 根据方向调整符号
	if direction == "S" || direction == "W" {
		decimal = -decimal
	}

	return decimal, true
}

// convertDMSToDecimal 将度分秒格式转换为十进制度数（保持向后兼容）
func (s *Driver) convertDMSToDecimal(dmsStr, direction string) float64 {
	result, _ := s.convertDMSToDecimalWithValidation(dmsStr, direction)
	return result
}

// parseFloat 解析浮点数字符串
func (s *Driver) parseFloat(str string) float64 {
	str = s.cleanString(str)
	if str == "" {
		return 0.0
	}

	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}

	return val
}

// cleanString 清理字符串，移除空字节和多余空格
func (s *Driver) cleanString(str string) string {
	// 移除空字节
	cleaned := strings.ReplaceAll(str, "\x00", "")
	// 移除前后空格
	cleaned = strings.TrimSpace(cleaned)
	return cleaned
}

// formatUTCTime 将UTC时间字符串格式化为易读格式
// 输入格式: HHMMSS.sss (例如: 123456.00)
// 输出格式: HH:MM:SS.sss (例如: 12:34:56.00)
func (s *Driver) formatUTCTime(utcStr string) string {
	if len(utcStr) < 6 {
		return utcStr // 如果格式不正确，返回原始字符串
	}

	// 解析时分秒
	hour := utcStr[0:2]
	minute := utcStr[2:4]
	second := utcStr[4:]

	// 格式化为 HH:MM:SS.sss
	return fmt.Sprintf("%s:%s:%s", hour, minute, second)
}

// formatCoordinate 格式化坐标为易读格式
// 输入: 十进制度数 (例如: 39.969056)
// 输出: 度分秒格式 (例如: 39°58'08.6"N)
func (s *Driver) formatCoordinate(decimal float64, isLatitude bool, direction string) string {
	if decimal == 0.0 {
		return "0°00'00.0\""
	}

	// 取绝对值进行计算
	absDecimal := decimal
	if absDecimal < 0 {
		absDecimal = -absDecimal
	}

	// 计算度分秒
	degrees := int(absDecimal)
	minutes := (absDecimal - float64(degrees)) * 60
	minutesInt := int(minutes)
	seconds := (minutes - float64(minutesInt)) * 60

	// 格式化输出
	return fmt.Sprintf("%d°%02d'%04.1f\"%s", degrees, minutesInt, seconds, direction)
}

// formatSpeed 格式化速度为易读格式
func (s *Driver) formatSpeed(speedKmh float64) string {
	return fmt.Sprintf("%.2f km/h", speedKmh)
}

// formatCourse 格式化航向为易读格式
func (s *Driver) formatCourse(course float64) string {
	// 添加方向描述
	var direction string
	switch {
	case course >= 0 && course < 22.5:
		direction = "北"
	case course >= 22.5 && course < 67.5:
		direction = "东北"
	case course >= 67.5 && course < 112.5:
		direction = "东"
	case course >= 112.5 && course < 157.5:
		direction = "东南"
	case course >= 157.5 && course < 202.5:
		direction = "南"
	case course >= 202.5 && course < 247.5:
		direction = "西南"
	case course >= 247.5 && course < 292.5:
		direction = "西"
	case course >= 292.5 && course < 337.5:
		direction = "西北"
	default:
		direction = "北"
	}

	return fmt.Sprintf("%.1f° (%s)", course, direction)
}

// formatAltitude 格式化海拔高度为易读格式
func (s *Driver) formatAltitude(altitude float64) string {
	return fmt.Sprintf("%.1f 米", altitude)
}

// formatFixQuality 格式化定位质量为易读格式
func (s *Driver) formatFixQuality(quality int32) string {
	switch quality {
	case 0:
		return "无定位"
	case 1:
		return "GPS定位"
	case 2:
		return "差分GPS定位"
	case 3:
		return "PPS定位"
	case 4:
		return "RTK定位"
	case 5:
		return "浮点RTK"
	case 6:
		return "推算定位"
	case 7:
		return "手动输入"
	case 8:
		return "模拟定位"
	default:
		return fmt.Sprintf("未知质量(%d)", quality)
	}
}

// formatSatelliteCount 格式化卫星数量为易读格式
func (s *Driver) formatSatelliteCount(count int32) string {
	return fmt.Sprintf("%d 颗卫星", count)
}

// formatHDOP 格式化水平精度因子为易读格式
func (s *Driver) formatHDOP(hdop float64) string {
	var quality string
	switch {
	case hdop <= 1:
		quality = "优秀"
	case hdop <= 2:
		quality = "良好"
	case hdop <= 5:
		quality = "中等"
	case hdop <= 10:
		quality = "一般"
	case hdop <= 20:
		quality = "较差"
	default:
		quality = "很差"
	}
	return fmt.Sprintf("%.2f (%s)", hdop, quality)
}

// 公共格式化方法，供外部调用

// FormatUTCTime 公共方法：格式化UTC时间
func (s *Driver) FormatUTCTime(utcStr string) string {
	return s.formatUTCTime(utcStr)
}

// FormatCoordinate 公共方法：格式化坐标
func (s *Driver) FormatCoordinate(decimal float64, isLatitude bool, direction string) string {
	return s.formatCoordinate(decimal, isLatitude, direction)
}

// FormatSpeed 公共方法：格式化速度
func (s *Driver) FormatSpeed(speedKmh float64) string {
	return s.formatSpeed(speedKmh)
}

// FormatCourse 公共方法：格式化航向
func (s *Driver) FormatCourse(course float64) string {
	return s.formatCourse(course)
}

// FormatAltitude 公共方法：格式化海拔
func (s *Driver) FormatAltitude(altitude float64) string {
	return s.formatAltitude(altitude)
}

// FormatFixQuality 公共方法：格式化定位质量
func (s *Driver) FormatFixQuality(quality int32) string {
	return s.formatFixQuality(quality)
}

// FormatSatelliteCount 公共方法：格式化卫星数量
func (s *Driver) FormatSatelliteCount(count int32) string {
	return s.formatSatelliteCount(count)
}

// FormatHDOP 公共方法：格式化HDOP
func (s *Driver) FormatHDOP(hdop float64) string {
	return s.formatHDOP(hdop)
}
