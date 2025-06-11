package test

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/edgexfoundry/device-sdk-go/v4/run/driver"
)

func main() {
	fmt.Println("🧪 GPS NMEA解析测试")
	fmt.Println("==================")

	// 测试NMEA数据（来自用户提供的实际GPS数据）
	testData := []string{
		"$GBGGA,055525.000,3044.368753,N,10357.548051,E,1,   ,2.40,129.3,M,-32.3,M,,*5A",
		"$GBGLL,3044.368753,N,10357.548051,E,055525.000,A,A*4B",
		"$GBGSA,A,2,34,21,07,44,,,,,,,,,2.59,2.40,1.00,4*03",
		"$GBGSV,6,1,21,10,80,005,27,34,76,067,33,38,75,161,28,21,57,046,29,1*71",
		"$GBRMC,055525.000,A,3044.368753,N,10357.548051,E,0.00,000.00,100625,,,A,C*13",
		"$GBVTG,000.00,T,,M,0.00,N,0.00,K,A*2F",
	}

	// GPS设备实例将在实际解析中使用

	fmt.Println("\n📡 测试NMEA语句解析:")
	fmt.Println("--------------------")

	for i, sentence := range testData {
		fmt.Printf("\n%d. 测试语句: %s\n", i+1, sentence)

		// 测试语句类型识别
		nmeaType := driver.ParsNMEAType(sentence, len(sentence))
		fmt.Printf("   语句类型: %v\n", nmeaType)

		// 测试校验和验证
		isValid := driver.ValidateNMEAChecksum(sentence, len(sentence))
		fmt.Printf("   校验和验证: %v\n", isValid)

		// 根据类型解析具体数据
		switch nmeaType {
		case driver.NMEA_RMC_TYPE:
			rmc := driver.ParsNMEARMC(sentence, len(sentence))
			if rmc != nil {
				fmt.Printf("   ✅ RMC解析成功:\n")
				fmt.Printf("      时间: %s\n", trimNullBytes(rmc.UTC[:]))
				fmt.Printf("      状态: %s\n", trimNullBytes(rmc.Status[:]))
				fmt.Printf("      纬度: %s%s\n", trimNullBytes(rmc.Lat[:]), trimNullBytes(rmc.N_S[:]))
				fmt.Printf("      经度: %s%s\n", trimNullBytes(rmc.Lon[:]), trimNullBytes(rmc.E_W[:]))
				fmt.Printf("      速度: %s节\n", trimNullBytes(rmc.SOG[:]))
				fmt.Printf("      航向: %s度\n", trimNullBytes(rmc.COG[:]))
			}
		case driver.NMEA_GGA_TYPE:
			gga := driver.ParsNMEAGGA(sentence, len(sentence))
			if gga != nil {
				fmt.Printf("   ✅ GGA解析成功:\n")
				fmt.Printf("      时间: %s\n", trimNullBytes(gga.UTC[:]))
				fmt.Printf("      纬度: %s%s\n", trimNullBytes(gga.Lat[:]), trimNullBytes(gga.N_S[:]))
				fmt.Printf("      经度: %s%s\n", trimNullBytes(gga.Lon[:]), trimNullBytes(gga.E_W[:]))
				fmt.Printf("      定位质量: %s\n", trimNullBytes(gga.Quality[:]))
				fmt.Printf("      卫星数: %s\n", trimNullBytes(gga.NumSatUsed[:]))
				fmt.Printf("      HDOP: %s\n", trimNullBytes(gga.HDOP[:]))
				fmt.Printf("      海拔: %s%s\n", trimNullBytes(gga.Alt[:]), trimNullBytes(gga.AltM[:]))
			}
		default:
			fmt.Printf("   ⚠️  其他类型语句，暂未详细解析\n")
		}
	}

	fmt.Println("\n🔄 测试坐标转换:")
	fmt.Println("------------------")

	// 测试坐标转换
	testCoordinates := []struct {
		dms       string
		direction string
		desc      string
	}{
		{"3044.368753", "N", "北纬"},
		{"10357.548051", "E", "东经"},
		{"3044.368753", "S", "南纬"},
		{"10357.548051", "W", "西经"},
	}

	for _, coord := range testCoordinates {
		decimal := convertDMSToDecimal(coord.dms, coord.direction)
		fmt.Printf("%s %s -> %.6f°\n", coord.desc, coord.dms, decimal)
	}

	fmt.Println("\n✅ 测试完成!")
	fmt.Println("GPS设备服务已准备就绪，可以处理实际的串口数据。")
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

// convertDMSToDecimal 将度分秒格式转换为十进制度数
func convertDMSToDecimal(dmsStr, direction string) float64 {
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
