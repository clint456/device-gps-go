package driver

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/tarm/serial"
)

const BufSize = 2048

type LCX6XZ struct {
	NMEA_RMC    *NMEA_RMC
	NMEA_GGA    *NMEA_GGA
	NMEA_GLL    *NMEA_GLL
	NMEA_VTG    *NMEA_VTG
	NMEA_GSA    *NMEA_GSA
	NMEA_GSV    *NMEA_GSV
	OutputRates map[NMEA_SUB_ID]uint8 // 存储查询到的输出速率
	ResData     []byte
	mutex       sync.Mutex
	uartAckCh   chan struct{}
	uartFd      io.ReadWriteCloser
}

func UartRX_Task(lcx6xz *LCX6XZ) {
	var dataBuffer []byte            // 累积数据缓冲区
	readBuffer := make([]byte, 1024) // 单次读取缓冲区

	for {
		// 读取串口数据
		n, err := lcx6xz.uartFd.Read(readBuffer)

		// 处理读取错误
		if err != nil {
			if err.Error() == "timeout" {
				// 超时是正常的，继续读取
				continue
			}
			fmt.Printf("串口读取错误: %v\n", err)
			time.Sleep(100 * time.Millisecond) // 避免死循环
			continue
		}

		// 没有读取到数据
		if n == 0 {
			time.Sleep(10 * time.Millisecond) // 避免CPU占用过高
			continue
		}

		// 将新数据追加到累积缓冲区
		dataBuffer = append(dataBuffer, readBuffer[:n]...)

		// 处理累积缓冲区中的完整NMEA语句
		dataBuffer = processNMEAData(dataBuffer, lcx6xz)

		// 防止缓冲区无限增长
		if len(dataBuffer) > BufSize {
			fmt.Println("WARNING: 数据缓冲区过大，清空缓冲区")
			dataBuffer = dataBuffer[:0] // 清空缓冲区
		}
	}
}

// processNMEAData 处理NMEA数据并返回剩余的未处理数据
func processNMEAData(data []byte, lcx6xz *LCX6XZ) []byte {
	processed := 0

	for i := 0; i < len(data); {
		// 查找NMEA语句开始标记 '$'
		if data[i] == 0x24 {
			// 查找完整的NMEA语句（以\r\n结尾）
			endPos := findNMEAEnd(data[i:])
			if endPos == -1 {
				// 没有找到完整语句，保留剩余数据
				break
			}

			// 解析NMEA语句
			sentence := data[i : i+endPos]
			err := parseNMEASentence(sentence, lcx6xz)
			if err != nil {
				fmt.Printf("NMEA解析错误: %v\n", err)
			}

			// 移动到下一个位置
			i += endPos
			processed = i

		} else if i < len(data)-1 && data[i] == 0xF1 && data[i+1] == 0xD9 {
			// 处理二进制协议（如果需要）
			skip, err := ParsBM(data[i:], lcx6xz)
			if err != nil {
				fmt.Printf("二进制协议解析错误: %v\n", err)
				i++ // 跳过当前字节
			} else {
				i += skip
				processed = i
			}
		} else {
			// 跳过无效字节
			i++
		}
	}

	// 返回未处理的数据
	if processed > 0 {
		return data[processed:]
	}
	return data
}

// findNMEAEnd 查找NMEA语句的结束位置
func findNMEAEnd(data []byte) int {
	for i := 0; i < len(data)-1; i++ {
		if data[i] == '\r' && data[i+1] == '\n' {
			return i + 2 // 包含\r\n
		}
	}
	return -1 // 没有找到结束标记
}

// parseNMEASentence 解析单个NMEA语句
func parseNMEASentence(sentence []byte, lcx6xz *LCX6XZ) error {
	// 移除\r\n
	sentenceStr := string(sentence)
	if len(sentenceStr) >= 2 && sentenceStr[len(sentenceStr)-2:] == "\r\n" {
		sentenceStr = sentenceStr[:len(sentenceStr)-2]
	}

	// 验证最小长度
	if len(sentenceStr) < 6 {
		return fmt.Errorf("NMEA语句太短: %s", sentenceStr)
	}

	// 解析NMEA语句类型
	nmeaType := ParsNMEAType(sentenceStr, len(sentenceStr))

	// 加锁保护共享数据
	lcx6xz.mutex.Lock()
	defer lcx6xz.mutex.Unlock()

	switch nmeaType {
	case NMEA_RMC_TYPE:
		rmc := ParsNMEARMC(sentenceStr, len(sentenceStr))
		if rmc != nil {
			lcx6xz.NMEA_RMC = rmc
			fmt.Printf("✅ RMC: 时间=%s, 纬度=%s%s, 经度=%s%s, 状态=%s\n",
				trimNullBytes(rmc.UTC[:]), trimNullBytes(rmc.Lat[:]), trimNullBytes(rmc.N_S[:]),
				trimNullBytes(rmc.Lon[:]), trimNullBytes(rmc.E_W[:]), trimNullBytes(rmc.Status[:]))
		}
	case NMEA_GGA_TYPE:
		gga := ParsNMEAGGA(sentenceStr, len(sentenceStr))
		if gga != nil {
			lcx6xz.NMEA_GGA = gga // 存储GGA数据
			fmt.Printf("✅ GGA: 时间=%s, 纬度=%s%s, 经度=%s%s, 质量=%s, 卫星数=%s\n",
				trimNullBytes(gga.UTC[:]), trimNullBytes(gga.Lat[:]), trimNullBytes(gga.N_S[:]),
				trimNullBytes(gga.Lon[:]), trimNullBytes(gga.E_W[:]),
				trimNullBytes(gga.Quality[:]), trimNullBytes(gga.NumSatUsed[:]))
		}
	case NMEA_GLL_TYPE:
		gll := ParsNMEAGLL(sentenceStr, len(sentenceStr))
		if gll != nil {
			lcx6xz.NMEA_GLL = gll // 存储GLL数据
			fmt.Printf("✅ GLL: 纬度=%s%s, 经度=%s%s, 时间=%s, 状态=%s\n",
				trimNullBytes(gll.Lat[:]), trimNullBytes(gll.N_S[:]),
				trimNullBytes(gll.Lon[:]), trimNullBytes(gll.E_W[:]),
				trimNullBytes(gll.UTC[:]), trimNullBytes(gll.Status[:]))
		}
	case NMEA_GSA_TYPE:
		gsa := ParsNMEAGSA(sentenceStr, len(sentenceStr))
		if gsa != nil {
			lcx6xz.NMEA_GSA = gsa // 存储GSA数据
			fmt.Printf("✅ GSA: 模式=%s, 定位模式=%s, PDOP=%s, HDOP=%s, VDOP=%s\n",
				trimNullBytes(gsa.Mode[:]), trimNullBytes(gsa.FixMode[:]),
				trimNullBytes(gsa.PDOP[:]), trimNullBytes(gsa.HDOP[:]), trimNullBytes(gsa.VDOP[:]))
		}
	case NMEA_GSV_TYPE:
		gsv := ParsNMEAGSV(sentenceStr, len(sentenceStr))
		if gsv != nil {
			lcx6xz.NMEA_GSV = gsv // 存储GSV数据
			fmt.Printf("✅ GSV: 总语句数=%s, 语句号=%s, 可视卫星数=%s\n",
				trimNullBytes(gsv.TotalNumSen[:]), trimNullBytes(gsv.SenNum[:]),
				trimNullBytes(gsv.TotalNumSat[:]))
		}
	case NMEA_VTG_TYPE:
		vtg := ParsNMEAVTG(sentenceStr, len(sentenceStr))
		if vtg != nil {
			lcx6xz.NMEA_VTG = vtg // 存储VTG数据
			fmt.Printf("✅ VTG: 航向=%s, 速度(节)=%s, 速度(km/h)=%s\n",
				trimNullBytes(vtg.COGT[:]), trimNullBytes(vtg.SOGN[:]), trimNullBytes(vtg.SOGK[:]))
		}
	default:
		if len(sentenceStr) >= 6 {
			fmt.Printf("⚠️  未知NMEA语句类型: %s\n", sentenceStr[:6])
		}
	}

	return nil
}

// trimNullBytes 移除字节数组中的空字节
func trimNullBytes(data []byte) string {
	for i, b := range data {
		if b == 0 {
			return string(data[:i])
		}
	}
	return string(data)
}

// ParsNMEA 解析NMEA协议语句（保持向后兼容）
func ParsNMEA(buffer []byte, lcx6xz *LCX6XZ) (int, error) {
	// 使用新的解析函数
	err := parseNMEASentence(buffer, lcx6xz)
	if err != nil {
		return 0, err
	}

	// 查找结束位置
	endPos := findNMEAEnd(buffer)
	if endPos == -1 {
		return 0, errors.New("incomplete NMEA sentence")
	}

	return endPos, nil
}

// ParsBM 解析二进制协议语句
func ParsBM(buffer []byte, lcx6xz *LCX6XZ) (int, error) {
	if len(buffer) < 8 {
		return 0, errors.New("二进制消息太短")
	}

	// 检查帧头 0xF1 0xD9
	if buffer[0] != 0xF1 || buffer[1] != 0xD9 {
		return 0, errors.New("无效的二进制消息头")
	}

	groupID := buffer[2]
	subID := buffer[3]
	length := uint16(buffer[4]) | (uint16(buffer[5]) << 8)

	// 检查消息长度
	totalLen := int(6 + length + 2) // 头部6字节 + 载荷 + 校验和2字节
	if len(buffer) < totalLen {
		return 0, errors.New("二进制消息数据不完整")
	}

	// 验证校验和
	checksum := QlCheckQuectel(buffer[:totalLen-2])
	receivedChecksum := uint16(buffer[totalLen-2]) | (uint16(buffer[totalLen-1]) << 8)

	if checksum != receivedChecksum {
		return 0, fmt.Errorf("校验和错误: 期望 %04X, 实际 %04X", checksum, receivedChecksum)
	}

	// 处理不同类型的二进制消息
	switch groupID {
	case 0x05: // BIN_RES_GID - 响应消息
		switch subID {
		case 0x01: // ACK
			fmt.Printf("✅ 收到ACK确认: GroupID=0x%02X, SubID=0x%02X\n",
				buffer[6], buffer[7])
		case 0x00: // NAK
			fmt.Printf("❌ 收到NAK否认: GroupID=0x%02X, SubID=0x%02X\n",
				buffer[6], buffer[7])
		}
	case 0x06: // BIN_CFG_GID - 配置消息响应
		switch subID {
		case 0x01: // MSG配置响应
			if length >= 3 {
				targetGroupID := buffer[6]
				targetSubID := buffer[7]
				outputRate := buffer[8]
				fmt.Printf("📊 输出速率响应: NMEA类型=0x%02X%02X, 速率=%d\n",
					targetGroupID, targetSubID, outputRate)

				// 存储查询结果到设备结构中
				if lcx6xz != nil {
					lcx6xz.mutex.Lock()
					if lcx6xz.OutputRates == nil {
						lcx6xz.OutputRates = make(map[NMEA_SUB_ID]uint8)
					}
					lcx6xz.OutputRates[NMEA_SUB_ID(targetSubID)] = outputRate
					lcx6xz.mutex.Unlock()
				}
			}
		}
	}

	return totalLen, nil
}

// SendBinaryCommand 发送二进制命令到GPS设备
func SendBinaryCommand(lcx6xz *LCX6XZ, data []byte) error {
	if lcx6xz == nil || lcx6xz.uartFd == nil {
		return errors.New("GPS设备未连接")
	}

	fmt.Printf("📤 发送二进制命令: %X\n", data)

	_, err := lcx6xz.uartFd.Write(data)
	if err != nil {
		return fmt.Errorf("发送二进制命令失败: %v", err)
	}

	return nil
}

// SetNMEAOutputRate 设置NMEA消息输出速率
func SetNMEAOutputRate(lcx6xz *LCX6XZ, nmeaType NMEA_SUB_ID, rate uint8) error {
	// 创建设置输出速率的配置消息
	cfgMsg := CfgMsgSetOutRate(NMEA_GID, nmeaType, rate)
	if cfgMsg == nil {
		return errors.New("创建配置消息失败")
	}

	// 发送配置消息
	msgBytes := cfgMsg.ToBytes()
	if msgBytes == nil {
		return errors.New("转换配置消息失败")
	}

	return SendBinaryCommand(lcx6xz, msgBytes)
}

// GetNMEAOutputRate 查询NMEA消息输出速率
func GetNMEAOutputRate(lcx6xz *LCX6XZ, nmeaType NMEA_SUB_ID) error {
	// 创建查询输出速率的配置消息
	cfgMsg := CfgMsgQueOutRate(NMEA_GID, nmeaType)
	if cfgMsg == nil {
		return errors.New("创建查询消息失败")
	}

	// 发送查询消息
	msgBytes := cfgMsg.ToBytes()
	if msgBytes == nil {
		return errors.New("转换查询消息失败")
	}

	return SendBinaryCommand(lcx6xz, msgBytes)
}

// 初始化LCX6XZ
func InitLCX6XZ() (*LCX6XZ, error) {
	lcx6xz := &LCX6XZ{
		OutputRates: make(map[NMEA_SUB_ID]uint8),
		ResData:     make([]byte, 1024),
		uartAckCh:   make(chan struct{}),
	}

	// 配置串口
	config := &serial.Config{
		Name:        "/dev/ttyUSB0",
		Baud:        9600,
		Size:        8,
		Parity:      serial.ParityNone,
		StopBits:    serial.Stop1,
		ReadTimeout: 100 * time.Millisecond,
	}

	port, err := serial.OpenPort(config)
	if err != nil {
		return nil, fmt.Errorf("open uart device: %w", err)
	}

	lcx6xz.uartFd = port

	// 启动接收任务
	go UartRX_Task(lcx6xz)

	return lcx6xz, nil
}
