package driver

import (
	"encoding/binary"
	"fmt"
)

// GroupID 消息组ID枚举
type GroupID uint8

const (
	NMEA_GID    GroupID = 0xF0 // NMEA标准语句
	BIN_RES_GID GroupID = 0x05 // 二进制协议，响应消息
	BIN_CFG_GID GroupID = 0x06 // 二进制协议，配置消息
)

// NMEA_SUB_ID NMEA子消息ID枚举
type NMEA_SUB_ID uint8

const (
	NMEA_GGA_SID NMEA_SUB_ID = 0x00 // GGA：全球定位系统定位数据
	NMEA_GLL_SID NMEA_SUB_ID = 0x01 // GLL：地理位置—纬度和经度
	NMEA_GSA_SID NMEA_SUB_ID = 0x02 // GSA：GNSS精度因子（DOP）与有效卫星
	NMEA_GRS_SID NMEA_SUB_ID = 0x03 // GRS：GNSS距离残差
	NMEA_GSV_SID NMEA_SUB_ID = 0x04 // GSV：可视的GNSS卫星
	NMEA_RMC_SID NMEA_SUB_ID = 0x05 // RMC：推荐的最少专有GNSS数据
	NMEA_VTG_SID NMEA_SUB_ID = 0x06 // VTG：相对于地面的实际航向和速度
	NMEA_ZDA_SID NMEA_SUB_ID = 0x07 // ZDA：时间与日期。主要用于输出UTC时间，日、月、年；但不支持输出本地时区信息
	NMEA_GST_SID NMEA_SUB_ID = 0x08 // GST：GNSS伪距误差统计
)

// BIN_RES_SID 二进制响应消息子ID枚举
type BIN_RES_SID uint8

const (
	BM_NAK_SID BIN_RES_SID = 0x00 // NAK：否认消息
	BM_ACK_SID BIN_RES_SID = 0x01 // ACK：确认消息
)

// BIN_CFG_SID 二进制配置消息子ID枚举
type BIN_CFG_SID uint8

const (
	BM_PRT_SID       BIN_CFG_SID = 0x00 // PRT：通信接口配置
	BM_MSG_SID       BIN_CFG_SID = 0x01 // MSG：消息输出速率配置
	BM_PPS_SID       BIN_CFG_SID = 0x07 // PPS：PPS配置
	BM_CFG_SID       BIN_CFG_SID = 0x09 // CFG：清除/保存当前配置
	BM_DOP_SID       BIN_CFG_SID = 0x0A // DOP：导航定位DOP阈值配置
	BM_ELEV_SID      BIN_CFG_SID = 0x0B // ELEV：导航定位的卫星仰角阈值配置
	BM_NAVSAT_SID    BIN_CFG_SID = 0x0C // NAVSAT：卫星系统使能状态配置
	BM_SPDHOLD_SID   BIN_CFG_SID = 0x0F // SPDHOLD：静态速度阈值配置
	BM_EPHSAVE_SID   BIN_CFG_SID = 0x10 // EPHSAVE：Flash中的星历存储状态配置
	BM_SIMPLERST_SID BIN_CFG_SID = 0x40 // SIMPLERST：GNSS引擎启动/关闭/复位配置
	BM_SLEEP_SID     BIN_CFG_SID = 0x41 // SLEEP：GNSS模块进入休眠状态配置
	BM_PWRCTL_SID    BIN_CFG_SID = 0x42 // PWRCTL：模块功率控制模式和定位频率配置
)

// CFG_MSG 配置消息结构体
type CFG_MSG struct {
	Data    [32]byte // 数据缓冲区
	DataLen int      // 数据长度
}

// BinaryMessage 二进制消息基础结构
type BinaryMessage struct {
	Header  uint16  // 消息头
	GroupID GroupID // 组ID
	SubID   uint8   // 子ID
	Length  uint16  // 消息长度
	Payload []byte  // 消息载荷
	CRC     uint16  // 校验和
}

// MessageConfig 消息配置结构
type MessageConfig struct {
	GroupID GroupID     // 消息组ID
	SubID   NMEA_SUB_ID // 消息子ID
	OutRate uint8       // 输出速率
}

// QlCheckQuectel 计算Quectel校验和 (与C代码保持一致)
func QlCheckQuectel(data []byte) uint16 {
	if data == nil || len(data) < 4 {
		return 0
	}

	var chk1 uint8 = 0
	var chk2 uint8 = 0

	// 从索引2开始计算校验和
	for i := 2; i < len(data); i++ {
		chk1 = chk1 + data[i]
		chk2 = chk2 + chk1
	}

	chk1 = chk1 & 0xFF
	chk2 = chk2 & 0xFF

	return uint16(chk1) | (uint16(chk2) << 8)
}

// CfgMsgSetOutRate 设置消息输出速率
func CfgMsgSetOutRate(gid GroupID, sid NMEA_SUB_ID, outRate uint8) *CFG_MSG {
	msg := &CFG_MSG{
		DataLen: 0,
	}

	// 构建配置消息 (与C代码保持一致)
	msgData := []byte{
		0xF1, 0xD9, // 帧头
		0x06, 0x01, // 消息组ID 消息子ID
		0x03, 0x00, // 长度（字节）
		0x00, 0x00, 0x00, // 有效载荷:消息组ID 消息子ID 输出速率
		0x00, 0x00, // 校验码:CHK1 CHK2
	}

	// 设置载荷数据
	msgData[6] = uint8(gid) // 目标消息组ID
	msgData[7] = uint8(sid) // 目标消息子ID
	msgData[8] = outRate    // 输出速率

	// 计算校验和 (不包括最后2字节的校验码)
	checksum := QlCheckQuectel(msgData[:len(msgData)-2])
	msgData[9] = uint8(checksum & 0xFF)
	msgData[10] = uint8(checksum >> 8)

	// 复制到CFG_MSG结构
	copy(msg.Data[:], msgData)
	msg.DataLen = len(msgData)
	return msg
}

// CfgMsgQueOutRate 查询消息输出速率
func CfgMsgQueOutRate(gid GroupID, sid NMEA_SUB_ID) *CFG_MSG {
	msg := &CFG_MSG{
		DataLen: 0,
	}

	// 构建查询消息 (与C代码保持一致)
	msgData := []byte{
		0xF1, 0xD9, // 帧头
		0x06, 0x01, // 消息组ID 消息子ID
		0x02, 0x00, // 长度（字节）
		0x00, 0x00, // 有效载荷:消息组ID 消息子ID
		0x00, 0x00, // 校验码:CHK1 CHK2
	}

	// 设置载荷数据
	msgData[6] = uint8(gid) // 目标消息组ID
	msgData[7] = uint8(sid) // 目标消息子ID

	// 计算校验和 (不包括最后2字节的校验码)
	checksum := QlCheckQuectel(msgData[:len(msgData)-2])
	msgData[len(msgData)-2] = uint8(checksum & 0xFF)
	msgData[len(msgData)-1] = uint8(checksum >> 8)

	// 复制到CFG_MSG结构
	copy(msg.Data[:], msgData)
	msg.DataLen = len(msgData)
	return msg
}

// ParseBinaryMessage 解析二进制消息
func ParseBinaryMessage(data []byte) (*BinaryMessage, error) {
	if len(data) < 8 {
		return nil, fmt.Errorf("消息长度不足")
	}

	// 检查消息头
	if data[0] != 0xB5 || data[1] != 0x62 {
		return nil, fmt.Errorf("无效的消息头")
	}

	msg := &BinaryMessage{
		Header:  binary.LittleEndian.Uint16(data[0:2]),
		GroupID: GroupID(data[2]),
		SubID:   data[3],
		Length:  binary.LittleEndian.Uint16(data[4:6]),
	}

	// 检查消息长度
	expectedLen := int(msg.Length) + 8 // 头部6字节 + 载荷 + 校验和2字节
	if len(data) < expectedLen {
		return nil, fmt.Errorf("消息数据不完整")
	}

	// 提取载荷
	if msg.Length > 0 {
		msg.Payload = make([]byte, msg.Length)
		copy(msg.Payload, data[6:6+msg.Length])
	}

	// 提取校验和
	crcOffset := 6 + int(msg.Length)
	msg.CRC = binary.LittleEndian.Uint16(data[crcOffset : crcOffset+2])

	// 验证校验和
	checksumData := data[2:crcOffset]
	expectedCRC := QlCheckQuectel(checksumData)
	if msg.CRC != expectedCRC {
		return nil, fmt.Errorf("校验和错误: 期望 %04X, 实际 %04X", expectedCRC, msg.CRC)
	}

	return msg, nil
}

// ToBytes 将CFG_MSG转换为字节数组
func (msg *CFG_MSG) ToBytes() []byte {
	if msg.DataLen <= 0 || msg.DataLen > len(msg.Data) {
		return nil
	}

	result := make([]byte, msg.DataLen)
	copy(result, msg.Data[:msg.DataLen])
	return result
}

// String 返回CFG_MSG的字符串表示
func (msg *CFG_MSG) String() string {
	if msg.DataLen <= 0 {
		return "CFG_MSG{empty}"
	}

	return fmt.Sprintf("CFG_MSG{DataLen: %d, Data: %X}", msg.DataLen, msg.Data[:msg.DataLen])
}

// GetGroupIDString 获取GroupID的字符串描述
func (gid GroupID) String() string {
	switch gid {
	case NMEA_GID:
		return "NMEA"
	case BIN_RES_GID:
		return "BIN_RES"
	case BIN_CFG_GID:
		return "BIN_CFG"
	default:
		return fmt.Sprintf("UNKNOWN(0x%02X)", uint8(gid))
	}
}

// GetNMEASubIDString 获取NMEA子ID的字符串描述
func (sid NMEA_SUB_ID) String() string {
	switch sid {
	case NMEA_GGA_SID:
		return "GGA"
	case NMEA_GLL_SID:
		return "GLL"
	case NMEA_GSA_SID:
		return "GSA"
	case NMEA_GRS_SID:
		return "GRS"
	case NMEA_GSV_SID:
		return "GSV"
	case NMEA_RMC_SID:
		return "RMC"
	case NMEA_VTG_SID:
		return "VTG"
	case NMEA_ZDA_SID:
		return "ZDA"
	case NMEA_GST_SID:
		return "GST"
	default:
		return fmt.Sprintf("UNKNOWN(0x%02X)", uint8(sid))
	}
}
