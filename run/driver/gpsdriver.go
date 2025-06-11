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

	lat := string(s.gpsDevice.NMEA_RMC.Lat[:])
	ns := string(s.gpsDevice.NMEA_RMC.N_S[:])

	if lat == "" {
		return nil
	}

	// 转换为十进制度数格式
	latValue := s.convertDMSToDecimal(lat, ns)

	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "Float64", latValue)
	return cv
}

// getLongitude 获取经度
func (s *Driver) getLongitude(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil || s.gpsDevice.NMEA_RMC == nil {
		return nil
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	lon := string(s.gpsDevice.NMEA_RMC.Lon[:])
	ew := string(s.gpsDevice.NMEA_RMC.E_W[:])

	if lon == "" {
		return nil
	}

	// 转换为十进制度数格式
	lonValue := s.convertDMSToDecimal(lon, ew)

	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "Float64", lonValue)
	return cv
}

// getAltitude 获取海拔高度
func (s *Driver) getAltitude(req dsModels.CommandRequest) *dsModels.CommandValue {
	// 这里需要从GGA语句获取海拔信息，暂时返回0
	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "Float64", 0.0)
	return cv
}

// getSpeed 获取速度
func (s *Driver) getSpeed(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil || s.gpsDevice.NMEA_RMC == nil {
		return nil
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	sogStr := string(s.gpsDevice.NMEA_RMC.SOG[:])
	if sogStr == "" {
		return nil
	}

	// 转换速度（节）为km/h
	sog := s.parseFloat(sogStr)
	speedKmh := sog * 1.852 // 1节 = 1.852 km/h

	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "Float64", speedKmh)
	return cv
}

// getCourse 获取航向
func (s *Driver) getCourse(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil || s.gpsDevice.NMEA_RMC == nil {
		return nil
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	cogStr := string(s.gpsDevice.NMEA_RMC.COG[:])
	if cogStr == "" {
		return nil
	}

	cog := s.parseFloat(cogStr)
	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "Float64", cog)
	return cv
}

// getUTCTime 获取UTC时间
func (s *Driver) getUTCTime(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil || s.gpsDevice.NMEA_RMC == nil {
		return nil
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	utcStr := string(s.gpsDevice.NMEA_RMC.UTC[:])
	if utcStr == "" {
		return nil
	}

	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "String", utcStr)
	return cv
}

// getFixQuality 获取定位质量
func (s *Driver) getFixQuality(req dsModels.CommandRequest) *dsModels.CommandValue {
	if s.gpsDevice == nil || s.gpsDevice.NMEA_RMC == nil {
		return nil
	}

	s.gpsDevice.mutex.Lock()
	defer s.gpsDevice.mutex.Unlock()

	status := string(s.gpsDevice.NMEA_RMC.Status[:])
	quality := 0
	if status == "A" {
		quality = 1 // 有效定位
	}

	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "Int32", int32(quality))
	return cv
}

// getSatellitesUsed 获取使用的卫星数
func (s *Driver) getSatellitesUsed(req dsModels.CommandRequest) *dsModels.CommandValue {
	// 这里需要从GGA语句获取卫星数信息，暂时返回0
	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "Int32", int32(0))
	return cv
}

// getHDOP 获取水平精度因子
func (s *Driver) getHDOP(req dsModels.CommandRequest) *dsModels.CommandValue {
	// 这里需要从GGA语句获取HDOP信息，暂时返回0
	cv, _ := dsModels.NewCommandValue(req.DeviceResourceName, "Float64", 0.0)
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

	status := string(s.gpsDevice.NMEA_RMC.Status[:])
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

// convertDMSToDecimal 将度分秒格式转换为十进制度数
func (s *Driver) convertDMSToDecimal(dmsStr, direction string) float64 {
	if dmsStr == "" {
		return 0.0
	}

	// 移除空字符
	dmsStr = strings.TrimSpace(dmsStr)
	if len(dmsStr) < 4 {
		return 0.0
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
				return 0.0
			}
			minutes, err = strconv.ParseFloat(dmsStr[dotIndex-2:], 64)
			if err != nil {
				return 0.0
			}
		} else if dotIndex >= 3 {
			// 纬度格式 ddmm.mmmm
			degrees, err = strconv.ParseFloat(dmsStr[:dotIndex-2], 64)
			if err != nil {
				return 0.0
			}
			minutes, err = strconv.ParseFloat(dmsStr[dotIndex-2:], 64)
			if err != nil {
				return 0.0
			}
		}
	}

	decimal := degrees + minutes/60.0

	// 根据方向调整符号
	if direction == "S" || direction == "W" {
		decimal = -decimal
	}

	return decimal
}

// parseFloat 解析浮点数字符串
func (s *Driver) parseFloat(str string) float64 {
	str = strings.TrimSpace(str)
	if str == "" {
		return 0.0
	}

	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}

	return val
}
