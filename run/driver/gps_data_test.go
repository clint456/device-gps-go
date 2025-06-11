package driver

import (
	"testing"

	dsModels "github.com/edgexfoundry/device-sdk-go/v4/pkg/models"
)

// TestGPSDataRetrieval 测试GPS数据获取功能
func TestGPSDataRetrieval(t *testing.T) {
	// 创建模拟的GPS设备
	gpsDevice := &LCX6XZ{}

	// 模拟RMC数据
	rmc := &NMEA_RMC{}
	copy(rmc.UTC[:], "123456.00")
	copy(rmc.Lat[:], "3958.1234")
	copy(rmc.N_S[:], "N")
	copy(rmc.Lon[:], "11629.5678")
	copy(rmc.E_W[:], "E")
	copy(rmc.SOG[:], "12.5")
	copy(rmc.COG[:], "45.2")
	copy(rmc.Status[:], "A")
	gpsDevice.NMEA_RMC = rmc

	// 模拟GGA数据
	gga := &NMEA_GGA{}
	copy(gga.UTC[:], "123456.00")
	copy(gga.Lat[:], "3958.1234")
	copy(gga.N_S[:], "N")
	copy(gga.Lon[:], "11629.5678")
	copy(gga.E_W[:], "E")
	copy(gga.Quality[:], "1")
	copy(gga.NumSatUsed[:], "8")
	copy(gga.HDOP[:], "1.2")
	copy(gga.Alt[:], "123.4")
	gpsDevice.NMEA_GGA = gga

	// 模拟VTG数据
	vtg := &NMEA_VTG{}
	copy(vtg.COGT[:], "45.2")
	copy(vtg.SOGN[:], "12.5")
	copy(vtg.SOGK[:], "23.15")
	gpsDevice.NMEA_VTG = vtg

	// 模拟GSA数据
	gsa := &NMEA_GSA{}
	copy(gsa.Mode[:], "A")
	copy(gsa.FixMode[:], "3")
	copy(gsa.HDOP[:], "1.2")
	copy(gsa.PDOP[:], "1.5")
	copy(gsa.VDOP[:], "0.8")
	gpsDevice.NMEA_GSA = gsa

	// 创建Driver实例
	driver := &Driver{
		gpsDevice: gpsDevice,
	}

	// 测试各种数据读取
	testCases := []struct {
		resourceName string
		expectedType string
	}{
		{"latitude", "String"},        // 格式化的坐标字符串 (度分秒格式)
		{"longitude", "String"},       // 格式化的坐标字符串 (度分秒格式)
		{"altitude", "String"},        // 格式化的海拔字符串 (带单位)
		{"speed", "String"},           // 格式化的速度字符串 (带单位)
		{"course", "String"},          // 格式化的航向字符串 (带方向)
		{"utc_time", "String"},        // 格式化的时间字符串 (HH:MM:SS.sss)
		{"fix_quality", "String"},     // 格式化的定位质量字符串
		{"satellites_used", "String"}, // 格式化的卫星数量字符串
		{"hdop", "String"},            // 格式化的精度因子字符串 (带质量评估)
		{"gps_status", "String"},      // GPS状态字符串
	}

	for _, tc := range testCases {
		t.Run(tc.resourceName, func(t *testing.T) {
			req := dsModels.CommandRequest{
				DeviceResourceName: tc.resourceName,
			}

			var cv *dsModels.CommandValue
			switch tc.resourceName {
			case "latitude":
				cv = driver.getLatitude(req)
			case "longitude":
				cv = driver.getLongitude(req)
			case "altitude":
				cv = driver.getAltitude(req)
			case "speed":
				cv = driver.getSpeed(req)
			case "course":
				cv = driver.getCourse(req)
			case "utc_time":
				cv = driver.getUTCTime(req)
			case "fix_quality":
				cv = driver.getFixQuality(req)
			case "satellites_used":
				cv = driver.getSatellitesUsed(req)
			case "hdop":
				cv = driver.getHDOP(req)
			case "gps_status":
				cv = driver.getGPSStatus(req)
			}

			if cv == nil {
				t.Errorf("Expected non-nil CommandValue for %s", tc.resourceName)
				return
			}

			if cv.Type != tc.expectedType {
				t.Errorf("Expected type %s for %s, got %s", tc.expectedType, tc.resourceName, cv.Type)
			}

			t.Logf("%s: %v (type: %s)", tc.resourceName, cv.Value, cv.Type)
		})
	}
}

// TestNMEAParsing 测试NMEA解析功能
func TestNMEAParsing(t *testing.T) {
	testCases := []struct {
		name     string
		sentence string
		nmeaType NMEA_TYPE
	}{
		{
			name:     "RMC",
			sentence: "$GPRMC,123456.00,A,3958.1234,N,11629.5678,E,12.5,45.2,010123,,,A*5E",
			nmeaType: NMEA_RMC_TYPE,
		},
		{
			name:     "GGA",
			sentence: "$GPGGA,123456.00,3958.1234,N,11629.5678,E,1,8,1.2,123.4,M,45.6,M,,*7A",
			nmeaType: NMEA_GGA_TYPE,
		},
		{
			name:     "GLL",
			sentence: "$GPGLL,3958.1234,N,11629.5678,E,123456.00,A,A*5C",
			nmeaType: NMEA_GLL_TYPE,
		},
		{
			name:     "VTG",
			sentence: "$GPVTG,45.2,T,,M,12.5,N,23.15,K,A*3F",
			nmeaType: NMEA_VTG_TYPE,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nmeaType := ParsNMEAType(tc.sentence, len(tc.sentence))
			if nmeaType != tc.nmeaType {
				t.Errorf("Expected NMEA type %d for %s, got %d", tc.nmeaType, tc.name, nmeaType)
			}

			// 测试具体的解析函数
			switch tc.nmeaType {
			case NMEA_RMC_TYPE:
				rmc := ParsNMEARMC(tc.sentence, len(tc.sentence))
				if rmc == nil {
					t.Errorf("Failed to parse RMC sentence")
				} else {
					t.Logf("RMC parsed: UTC=%s, Lat=%s, Lon=%s",
						string(rmc.UTC[:]), string(rmc.Lat[:]), string(rmc.Lon[:]))
				}
			case NMEA_GGA_TYPE:
				gga := ParsNMEAGGA(tc.sentence, len(tc.sentence))
				if gga == nil {
					t.Errorf("Failed to parse GGA sentence")
				} else {
					t.Logf("GGA parsed: UTC=%s, Lat=%s, Lon=%s, Quality=%s",
						string(gga.UTC[:]), string(gga.Lat[:]), string(gga.Lon[:]), string(gga.Quality[:]))
				}
			case NMEA_GLL_TYPE:
				gll := ParsNMEAGLL(tc.sentence, len(tc.sentence))
				if gll == nil {
					t.Errorf("Failed to parse GLL sentence")
				} else {
					t.Logf("GLL parsed: Lat=%s, Lon=%s, UTC=%s",
						string(gll.Lat[:]), string(gll.Lon[:]), string(gll.UTC[:]))
				}
			case NMEA_VTG_TYPE:
				vtg := ParsNMEAVTG(tc.sentence, len(tc.sentence))
				if vtg == nil {
					t.Errorf("Failed to parse VTG sentence")
				} else {
					t.Logf("VTG parsed: Course=%s, Speed(N)=%s, Speed(K)=%s",
						string(vtg.COGT[:]), string(vtg.SOGN[:]), string(vtg.SOGK[:]))
				}
			}
		})
	}
}

// TestStringCleaning 测试字符串清理功能
func TestStringCleaning(t *testing.T) {
	driver := &Driver{}

	testCases := []struct {
		input    string
		expected string
	}{
		{"123.45", "123.45"},
		{"123.45\x00\x00", "123.45"},
		{" 123.45 ", "123.45"},
		{" 123.45\x00 ", "123.45"},
		{"", ""},
		{"\x00\x00", ""},
	}

	for _, tc := range testCases {
		result := driver.cleanString(tc.input)
		if result != tc.expected {
			t.Errorf("cleanString(%q) = %q, expected %q", tc.input, result, tc.expected)
		}
	}
}

// TestDataFormatting 测试数据格式化功能
func TestDataFormatting(t *testing.T) {
	driver := &Driver{}

	t.Run("UTC时间格式化", func(t *testing.T) {
		testCases := []struct {
			input    string
			expected string
		}{
			{"123456.00", "12:34:56.00"},
			{"235959.999", "23:59:59.999"},
			{"000000", "00:00:00"},
			{"12345", "12:34:5"}, // 不完整的输入
		}

		for _, tc := range testCases {
			result := driver.formatUTCTime(tc.input)
			if result != tc.expected {
				t.Errorf("formatUTCTime(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		}
	})

	t.Run("坐标格式化", func(t *testing.T) {
		testCases := []struct {
			decimal   float64
			direction string
			expected  string
		}{
			{39.969056, "N", "39°58'08.6\"N"},
			{116.491667, "E", "116°29'30.0\"E"},
			{0.0, "N", "0°00'00.0\""},
		}

		for _, tc := range testCases {
			result := driver.formatCoordinate(tc.decimal, true, tc.direction)
			if result != tc.expected {
				t.Errorf("formatCoordinate(%.6f, %s) = %q, expected %q",
					tc.decimal, tc.direction, result, tc.expected)
			}
		}
	})

	t.Run("速度格式化", func(t *testing.T) {
		testCases := []struct {
			speed    float64
			expected string
		}{
			{23.15, "23.15 km/h"},
			{0.0, "0.00 km/h"},
			{120.5, "120.50 km/h"},
		}

		for _, tc := range testCases {
			result := driver.formatSpeed(tc.speed)
			if result != tc.expected {
				t.Errorf("formatSpeed(%.2f) = %q, expected %q", tc.speed, result, tc.expected)
			}
		}
	})

	t.Run("航向格式化", func(t *testing.T) {
		testCases := []struct {
			course   float64
			expected string
		}{
			{0.0, "0.0° (北)"},
			{45.0, "45.0° (东北)"},
			{90.0, "90.0° (东)"},
			{180.0, "180.0° (南)"},
			{270.0, "270.0° (西)"},
		}

		for _, tc := range testCases {
			result := driver.formatCourse(tc.course)
			if result != tc.expected {
				t.Errorf("formatCourse(%.1f) = %q, expected %q", tc.course, result, tc.expected)
			}
		}
	})

	t.Run("定位质量格式化", func(t *testing.T) {
		testCases := []struct {
			quality  int32
			expected string
		}{
			{0, "无定位"},
			{1, "GPS定位"},
			{2, "差分GPS定位"},
			{99, "未知质量(99)"},
		}

		for _, tc := range testCases {
			result := driver.formatFixQuality(tc.quality)
			if result != tc.expected {
				t.Errorf("formatFixQuality(%d) = %q, expected %q", tc.quality, result, tc.expected)
			}
		}
	})

	t.Run("HDOP格式化", func(t *testing.T) {
		testCases := []struct {
			hdop     float64
			expected string
		}{
			{0.8, "0.80 (优秀)"},
			{1.5, "1.50 (良好)"},
			{3.0, "3.00 (中等)"},
			{8.0, "8.00 (一般)"},
			{15.0, "15.00 (较差)"},
			{25.0, "25.00 (很差)"},
		}

		for _, tc := range testCases {
			result := driver.formatHDOP(tc.hdop)
			if result != tc.expected {
				t.Errorf("formatHDOP(%.2f) = %q, expected %q", tc.hdop, result, tc.expected)
			}
		}
	})
}

// TestZeroValueHandling 测试0值处理
func TestZeroValueHandling(t *testing.T) {
	// 创建模拟的GPS设备，包含0值数据
	gpsDevice := &LCX6XZ{}

	// 模拟RMC数据，包含0值
	rmc := &NMEA_RMC{}
	copy(rmc.UTC[:], "000000.00") // 0时0分0秒
	copy(rmc.Lat[:], "0000.0000") // 0度纬度（赤道）
	copy(rmc.N_S[:], "N")
	copy(rmc.Lon[:], "00000.0000") // 0度经度（本初子午线）
	copy(rmc.E_W[:], "E")
	copy(rmc.SOG[:], "0.0")  // 0速度（静止）
	copy(rmc.COG[:], "0.0")  // 0航向（正北）
	copy(rmc.Status[:], "A") // 有效状态
	gpsDevice.NMEA_RMC = rmc

	// 模拟GGA数据，包含0值
	gga := &NMEA_GGA{}
	copy(gga.UTC[:], "000000.00")
	copy(gga.Lat[:], "0000.0000")
	copy(gga.N_S[:], "N")
	copy(gga.Lon[:], "00000.0000")
	copy(gga.E_W[:], "E")
	copy(gga.Quality[:], "0")    // 0质量（无定位）
	copy(gga.NumSatUsed[:], "0") // 0颗卫星
	copy(gga.HDOP[:], "0.0")     // 0 HDOP
	copy(gga.Alt[:], "0.0")      // 0海拔（海平面）
	gpsDevice.NMEA_GGA = gga

	// 模拟VTG数据，包含0值
	vtg := &NMEA_VTG{}
	copy(vtg.COGT[:], "0.0") // 0航向
	copy(vtg.SOGN[:], "0.0") // 0速度（节）
	copy(vtg.SOGK[:], "0.0") // 0速度（km/h）
	gpsDevice.NMEA_VTG = vtg

	// 创建Driver实例
	driver := &Driver{
		gpsDevice: gpsDevice,
	}

	// 测试0值数据读取
	testCases := []struct {
		resourceName    string
		shouldHaveValue bool
		description     string
	}{
		{"latitude", true, "0度纬度应该是有效值"},
		{"longitude", true, "0度经度应该是有效值"},
		{"altitude", true, "0米海拔应该是有效值"},
		{"speed", true, "0速度应该是有效值（静止状态）"},
		{"course", true, "0航向应该是有效值（正北方向）"},
		{"utc_time", true, "00:00:00时间应该是有效值"},
		{"fix_quality", true, "0定位质量应该是有效值（无定位状态）"},
		{"satellites_used", true, "0颗卫星应该是有效值"},
		{"hdop", true, "0 HDOP应该是有效值"},
		{"gps_status", true, "GPS状态应该总是有值"},
	}

	for _, tc := range testCases {
		t.Run(tc.resourceName, func(t *testing.T) {
			req := dsModels.CommandRequest{
				DeviceResourceName: tc.resourceName,
			}

			var cv *dsModels.CommandValue
			switch tc.resourceName {
			case "latitude":
				cv = driver.getLatitude(req)
			case "longitude":
				cv = driver.getLongitude(req)
			case "altitude":
				cv = driver.getAltitude(req)
			case "speed":
				cv = driver.getSpeed(req)
			case "course":
				cv = driver.getCourse(req)
			case "utc_time":
				cv = driver.getUTCTime(req)
			case "fix_quality":
				cv = driver.getFixQuality(req)
			case "satellites_used":
				cv = driver.getSatellitesUsed(req)
			case "hdop":
				cv = driver.getHDOP(req)
			case "gps_status":
				cv = driver.getGPSStatus(req)
			}

			if tc.shouldHaveValue {
				if cv == nil {
					t.Errorf("%s: %s，但得到了nil", tc.resourceName, tc.description)
				} else {
					t.Logf("%s: %v ✓", tc.resourceName, cv.Value)
				}
			} else {
				if cv != nil {
					t.Errorf("%s: 应该返回nil，但得到了值: %v", tc.resourceName, cv.Value)
				}
			}
		})
	}
}
