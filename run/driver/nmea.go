package driver

import (
	"strconv"
	"strings"
)

// NMEA_TYPE 定义NMEA语句类型
type NMEA_TYPE int

const (
	NMEA_UNKONW_TYPE NMEA_TYPE = iota
	NMEA_RMC_TYPE
	NMEA_GGA_TYPE
	NMEA_GSV_TYPE
	NMEA_GSA_TYPE
	NMEA_VTG_TYPE
	NMEA_GLL_TYPE
	NMEA_ZDA_TYPE
	NMEA_GRS_TYPE
	NMEA_GST_TYPE
)

// NMEA 结构体定义基本的NMEA信息
type NMEA struct {
	/*
	   ---------------------------------------
	       GNSS星系配置    |   TalkerID
	   ---------------------------------------
	           GPS         |       GP
	   ---------------------------------------
	           GLONASS     |       GL
	   ---------------------------------------
	           Galileo     |       GA
	   ---------------------------------------
	           BDS         |       GB
	   ---------------------------------------
	           QZSS        |       GQ
	   ---------------------------------------
	           组合星系    |       GN
	   ---------------------------------------
	*/
	TalkerID [3]byte
	Type     [4]byte
}

// NMEA_RMC 结构体定义RMC语句信息
type NMEA_RMC struct {
	Nmea NMEA
	/*  含义:   定位的UTC时间。
	    格式:   hhmmss.sss
	            hh：时（00~23）
	            mm：分（00~59）
	            ss：秒（(00~59）
	            sss：秒的十进制小数 */
	UTC [11]byte

	Status [2]byte // A有效，V导航接收警告

	/*  含义:   纬度。
	    格式:   ddmm.mmmmmm
	            dd：度（00~90）
	            mm：分（00~59）
	            mmmmmm：分的十进制小数
	            数据无效时，为空。*/
	Lat [12]byte

	N_S [2]byte // 纬度方向。N=北 S=南。数据无效时，为空

	/*  含义:   经度。
	    格式:   dddmm.mmmmmm
	            ddd：度（00~180）
	            mm：分（00~59）
	            mmmmmm：分的十进制小数
	            数据无效时，为空。*/
	Lon [13]byte

	E_W [2]byte // 经度方向。E=东 W=西。数据无效时，为空

	SOG [11]byte // 对地速度。可变长度。最大长度待定

	COG [7]byte // 对地真航向。可变长度，最大值：359.99。

	/*  含义:   日期。
	    格式:   ddmmyy
	            dd：日
	            mm：月
	            yy：年，真实年份需要+2000*/
	Date [7]byte

	MagVar    [1]byte // MagVar 磁偏角暂不支持
	MagVarDir [1]byte // MagVarDir 磁偏角方向暂不支持

	ModeInd [2]byte /* 模式:A = 自主式模式，卫星系统处于非差分定位模式。
	   D = 差分模式，卫星系统处于差分定位模式；基于地面站或SBAS的修正。
	   N = 无定位，卫星系统没有用于位置定位，或定位无效。 */

	NavStatus [2]byte /* 导航状态:S = 安全
	   C = 警告
	   U = 不安全
	   V = 导航状态无效，设备不提供导航状态指示 */
}

// NMEA_GGA 结构体定义GGA语句信息
type NMEA_GGA struct {
	Nmea NMEA
	/*  含义:   定位的UTC时间。
	    格式:   hhmmss.sss
	            hh：时（00~23）
	            mm：分（00~59）
	            ss：秒（(00~59）
	            sss：秒的十进制小数 */
	UTC [11]byte

	/*  含义:   纬度。
	    格式:   ddmm.mmmmmm
	            dd：度（00~90）
	            mm：分（00~59）
	            mmmmmm：分的十进制小数
	            数据无效时，为空。*/
	Lat [12]byte

	N_S [2]byte // 纬度方向。N=北 S=南。数据无效时，为空

	/*  含义:   经度。
	    格式:   dddmm.mmmmmm
	            ddd：度（00~180）
	            mm：分（00~59）
	            mmmmmm：分的十进制小数
	            数据无效时，为空。*/
	Lon [13]byte

	E_W [2]byte // 经度方向。E=东 W=西。数据无效时，为空

	Quality [2]byte /* GPS 定位模式/状态指示:
	   0 = 定位不可用或无效
	   1 = GPS SPS 模式，定位有效
	   2 = 差分 GPS，SPS 模式，或SBAS，定位有效 */

	NumSatUsed  [3]byte  // 使用的卫星数
	HDOP        [6]byte  // 水平精度因子。数据无效时，此字段为 99.99。
	Alt         [10]byte // 平均海平面以上海拔（大地水准面）。数据无效时，此字段为空。
	AltM        [2]byte  // Alt的单位。M=米。
	Sep         [10]byte // 大地水准面差距（WGS84 基准面与平均海平面之间的差值）。数据无效时，此字段为空。
	SepM        [2]byte  // Sep的单位。M=米。
	DiffAge     [1]byte  // 差分卫星导航系统数据龄期。暂不支持。
	DiffStation [1]byte  // 差分基准站标识号。暂不支持。
}

// SAT_STATUS 结构体定义卫星状态信息
type SAT_STATUS struct {
	/*
	   ----------------------------------------------------------
	    卫星系统 | 系统标识符 | 卫星标识号 | 信号标识号/信号通道
	   ----------------------------------------------------------
	      GPS    | 1 | 1~32:GPS, 33~64:SBAS|       1=L1 C/A
	   ----------------------------------------------------------
	     GLONASS |      2     |    65~96   |         1=L1
	   ----------------------------------------------------------
	     Galileo |      3     |    1~36    |         7=E1
	   ----------------------------------------------------------
	      BDS    |      4     |    1~63    |     1=B1I, 3=B1C
	   ----------------------------------------------------------
	      QZSS   |      5     |    1~10    |       1=L1 C/A
	   ----------------------------------------------------------
	*/
	SatID   [3]byte // 卫星标识号。
	SatElev [3]byte // 仰角。范围：00~90。数据无效时，此字段为空。
	SatAz   [4]byte // 真方位角。范围：000~359。数据无效时，此字段为空。
	SatCN0  [3]byte // 载噪比（C/N0）。范围：00~99。未跟踪时为空。
}

// NMEA_GSV 结构体定义GSV语句信息
type NMEA_GSV struct {
	Nmea        NMEA
	TotalNumSen [2]byte       // 语句总数。范围：1~9。
	SenNum      [2]byte       // 语句号。范围：1~<TotalNumSen>。
	TotalNumSat [3]byte       // 可视的卫星总数。
	SatStatus   [4]SAT_STATUS // 卫星信息
	SignalID    [2]byte       // 信号标识符
}

// NMEA_GSA 结构体定义GSA语句信息
type NMEA_GSA struct {
	Nmea     NMEA
	Mode     [2]byte     // A = 自动，允许2D/3D定位模式自动变换。
	FixMode  [2]byte     // 定位模式: 1=定位不可用 2=2D定位模式 3=3D定位模式
	SatID    [12][3]byte // 解算中用到的卫星标识号。数据无效时，此字段为空。
	PDOP     [6]byte     // 位置精度因子，最大值为99.99。数据无效时，此字段为99.99。
	HDOP     [6]byte     // 水平精度因子，最大值为99.99。数据无效时，此字段为99.99。
	VDOP     [6]byte     // 垂直精度因子，最大值为99.99。数据无效时，此字段为99.99。
	SystemID [2]byte     // GNSS系统标识符。
}

// NMEA_VTG 结构体定义VTG语句信息
type NMEA_VTG struct {
	Nmea    NMEA
	COGT    [7]byte // 对地航向（真北）。数据无效时，此字段为空。
	COGTT   [2]byte // 固定字段：真。
	COGM    [7]byte // 对地航向（磁北）。 暂不支持。
	COGMM   [2]byte // 固定字段：磁场。
	SOGN    [5]byte // 对地速度，以节为单位。数据无效时，此字段为空。
	SOGNN   [2]byte // 固定字段：节。
	SOGK    [5]byte // 对地速度，以千米每小时为单位。数据无效时，此字段为空。
	SOGKK   [2]byte // 固定字段：千米每小时。
	ModeInd [2]byte /* 模式指示:
	   A = 自主式模式，卫星系统处于非差分定位模式
	   D = 差分模式，卫星系统处于差分定位模式；基
	   于地面站或SBAS的修正
	   N = 数据无效
	*/
}

// NMEA_GLL 结构体定义GLL语句信息
type NMEA_GLL struct {
	Nmea NMEA
	/*  含义:   纬度。
	    格式:   ddmm.mmmmmm
	            dd：度（00~90）
	            mm：分（00~59）
	            mmmmmm：分的十进制小数
	            数据无效时，为空。*/
	Lat [12]byte
	N_S [2]byte // 纬度方向。N=北 S=南。数据无效时，为空
	/*  含义:   经度。
	    格式:   dddmm.mmmmmm
	            ddd：度（00~180）
	            mm：分（00~59）
	            mmmmmm：分的十进制小数
	            数据无效时，为空。*/
	Lon [13]byte
	E_W [2]byte // 经度方向。E=东 W=西。数据无效时，为空
	/*  含义:   定位的UTC时间。
	    格式:   hhmmss.sss
	            hh：时（00~23）
	            mm：分（00~59）
	            ss：秒（(00~59）
	            sss：秒的十进制小数 */
	UTC     [11]byte
	Status  [2]byte // A有效，V导航接收警告
	ModeInd [2]byte /* 模式:A = 自主式模式，卫星系统处于非差分定位模式。
	   D = 差分模式，卫星系统处于差分定位模式；基于地面站或SBAS的修正。
	   N = 无定位，卫星系统没有用于位置定位，或定位无效。 */
}

// NMEA_ZDA 结构体定义ZDA语句信息
type NMEA_ZDA struct {
	Nmea NMEA
	/*  含义:   定位的UTC时间。
	    格式:   hhmmss.sss
	            hh：时（00~23）
	            mm：分（00~59）
	            ss：秒（(00~59）
	            sss：秒的十进制小数 */
	UTC       [11]byte
	Day       [3]byte // 日。范围：01~3
	Month     [3]byte // 月。范围：01~12。
	Year      [5]byte // 年。
	LocalHour [3]byte // 本时区小时。暂不支持。
	LocalMin  [3]byte // 本时区分钟。暂不支持。
}

// NMEA_GRS 结构体定义GRS语句信息
type NMEA_GRS struct {
	Nmea NMEA
	/*  含义:   定位的UTC时间。
	    格式:   hhmmss.sss
	            hh：时（00~23）
	            mm：分（00~59）
	            ss：秒（(00~59）
	            sss：秒的十进制小数 */
	UTC      [11]byte
	Mode     [2]byte     // 使用的计算方法。0=残差用于计算匹配GGA句子中给出的位置。1=在计算GGA位置后重新计算残差。
	Resi     [12][6]byte // 导航中使用的可视卫星的距离残差。范围：-999到999。数据无效时，此字段为空。
	SystemID [2]byte     // GNSS系统标识符。
	SignalID [2]byte     // 卫星标识号。
}

// NMEA_GST 结构体定义GST语句信息
type NMEA_GST struct {
	Nmea NMEA
	/*  含义:   定位的UTC时间。
	    格式:   hhmmss.sss
	            hh：时（00~23）
	            mm：分（00~59）
	            ss：秒（(00~59）
	            sss：秒的十进制小数 */
	UTC    [11]byte
	RMS_D  [10]byte // 导航过程中输入的标准偏差范围的RMS值。数据无效时，此字段为空。
	MajorD [10]byte // 误差椭圆半长轴的标准偏差。数据无效时，此字段为空。
	MinorD [10]byte // 误差椭圆半短轴的标准偏差。数据无效时，此字段为空。
	Orient [10]byte // 误差椭圆半长轴方向。数据无效时，此字段为空。
	LatD   [10]byte // 纬度误差的标准差。数据无效时，此字段为空。
	LonD   [10]byte // 经度误差的标准偏差。数据无效时，此字段为空。
	AltD   [10]byte // 高度误差的标准偏差。数据无效时，此字段为空。
}

// QlCheckXOR 计算校验和
func QlCheckXOR(pData []byte, length uint) byte {
	var checksum byte = 0
	for i := uint(0); i < length; i++ {
		checksum ^= pData[i]
	}
	return checksum
}

// Strnstr 在haystack中查找needle，限制搜索长度为len
func Strnstr(haystack string, needle string, length int) string {
	if len(needle) == 0 {
		return haystack
	}

	if len(haystack) == 0 || length <= 0 {
		return ""
	}

	if length > len(haystack) {
		length = len(haystack)
	}

	for i := 0; i <= length-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return haystack[i:]
		}
	}

	return ""
}

// ValidateNMEAChecksum 验证NMEA语句的校验和
func ValidateNMEAChecksum(sentence string, length int) bool {
	if length < 4 {
		return false
	}

	// 查找*号位置
	asteriskPos := strings.LastIndex(sentence, "*")
	if asteriskPos == -1 || asteriskPos >= length-2 {
		return false
	}

	// 提取校验和
	checksumStr := sentence[asteriskPos+1:]
	if len(checksumStr) < 2 {
		return false
	}

	expectedChecksum, err := strconv.ParseUint(checksumStr[:2], 16, 8)
	if err != nil {
		return false
	}

	// 计算实际校验和（从$后到*前的所有字符）
	var actualChecksum byte = 0
	for i := 1; i < asteriskPos; i++ {
		actualChecksum ^= sentence[i]
	}

	return actualChecksum == byte(expectedChecksum)
}

// ParsNMEAType 解析NMEA语句类型
func ParsNMEAType(strNMEA string, length int) NMEA_TYPE {
	if length < 6 {
		return NMEA_UNKONW_TYPE
	}

	// 检查是否以$开头
	if strNMEA[0] != '$' {
		return NMEA_UNKONW_TYPE
	}

	// 提取语句类型（跳过$和前两个字符的TalkerID）
	if length >= 6 {
		msgType := strNMEA[3:6]
		switch msgType {
		case "GGA":
			return NMEA_GGA_TYPE
		case "RMC":
			return NMEA_RMC_TYPE
		case "GLL":
			return NMEA_GLL_TYPE
		case "GSA":
			return NMEA_GSA_TYPE
		case "GSV":
			return NMEA_GSV_TYPE
		case "VTG":
			return NMEA_VTG_TYPE
		case "GST":
			return NMEA_GST_TYPE
		case "GRS":
			return NMEA_GRS_TYPE
		case "ZDA":
			return NMEA_ZDA_TYPE
		}
	}

	return NMEA_UNKONW_TYPE
}

// ParsNMEAGST 解析GST语句
func ParsNMEAGST(strGST string, length int) *NMEA_GST {
	// 实现解析GST语句的逻辑
	return nil
}

// ParsNMEAGRS 解析GRS语句
func ParsNMEAGRS(strGRS string, length int) *NMEA_GRS {
	// 实现解析GRS语句的逻辑
	return nil
}

// ParsNMEAZDA 解析ZDA语句
func ParsNMEAZDA(strZDA string, length int) *NMEA_ZDA {
	// 实现解析ZDA语句的逻辑
	return nil
}

// ParsNMEAGLL 解析GLL语句
func ParsNMEAGLL(strGLL string, length int) *NMEA_GLL {
	// 实现解析GLL语句的逻辑
	return nil
}

// ParsNMEAVTG 解析VTG语句
func ParsNMEAVTG(strVTG string, length int) *NMEA_VTG {
	// 实现解析VTG语句的逻辑
	return nil
}

// ParsNMEAGSA 解析GSA语句
func ParsNMEAGSA(strGSA string, length int) *NMEA_GSA {
	// 实现解析GSA语句的逻辑
	return nil
}

// ParsNMEAGSV 解析GSV语句
func ParsNMEAGSV(strGSV string, length int) *NMEA_GSV {
	// 实现解析GSV语句的逻辑
	return nil
}

// ParsNMEAGGA 解析GGA语句
func ParsNMEAGGA(strGGA string, length int) *NMEA_GGA {
	if length < 10 {
		return nil
	}

	// 验证校验和
	if !ValidateNMEAChecksum(strGGA, length) {
		return nil
	}

	// 分割字段
	fields := strings.Split(strGGA, ",")
	if len(fields) < 15 {
		return nil
	}

	gga := &NMEA_GGA{}

	// 解析TalkerID和Type
	if len(fields[0]) >= 6 {
		copy(gga.Nmea.TalkerID[:], fields[0][1:3])
		copy(gga.Nmea.Type[:], fields[0][3:6])
	}

	// 解析UTC时间
	if len(fields[1]) > 0 && len(fields[1]) <= 10 {
		copy(gga.UTC[:], fields[1])
	}

	// 解析纬度
	if len(fields[2]) > 0 && len(fields[2]) <= 11 {
		copy(gga.Lat[:], fields[2])
	}

	// 解析纬度方向
	if len(fields[3]) > 0 {
		copy(gga.N_S[:], fields[3])
	}

	// 解析经度
	if len(fields[4]) > 0 && len(fields[4]) <= 12 {
		copy(gga.Lon[:], fields[4])
	}

	// 解析经度方向
	if len(fields[5]) > 0 {
		copy(gga.E_W[:], fields[5])
	}

	// 解析定位质量
	if len(fields[6]) > 0 {
		copy(gga.Quality[:], fields[6])
	}

	// 解析使用的卫星数
	if len(fields[7]) > 0 && len(fields[7]) <= 2 {
		copy(gga.NumSatUsed[:], fields[7])
	}

	// 解析水平精度因子
	if len(fields[8]) > 0 && len(fields[8]) <= 5 {
		copy(gga.HDOP[:], fields[8])
	}

	// 解析海拔高度
	if len(fields[9]) > 0 && len(fields[9]) <= 9 {
		copy(gga.Alt[:], fields[9])
	}

	// 解析海拔高度单位
	if len(fields[10]) > 0 {
		copy(gga.AltM[:], fields[10])
	}

	// 解析大地水准面差距
	if len(fields[11]) > 0 && len(fields[11]) <= 9 {
		copy(gga.Sep[:], fields[11])
	}

	// 解析大地水准面差距单位
	if len(fields[12]) > 0 {
		copy(gga.SepM[:], fields[12])
	}

	return gga
}

// ParsNMEARMC 解析RMC语句
func ParsNMEARMC(strRMC string, length int) *NMEA_RMC {
	if length < 10 {
		return nil
	}

	// 验证校验和
	if !ValidateNMEAChecksum(strRMC, length) {
		return nil
	}

	// 分割字段
	fields := strings.Split(strRMC, ",")
	if len(fields) < 12 {
		return nil
	}

	rmc := &NMEA_RMC{}

	// 解析TalkerID和Type
	if len(fields[0]) >= 6 {
		copy(rmc.Nmea.TalkerID[:], fields[0][1:3])
		copy(rmc.Nmea.Type[:], fields[0][3:6])
	}

	// 解析UTC时间
	if len(fields[1]) > 0 && len(fields[1]) <= 10 {
		copy(rmc.UTC[:], fields[1])
	}

	// 解析状态
	if len(fields[2]) > 0 {
		copy(rmc.Status[:], fields[2])
	}

	// 解析纬度
	if len(fields[3]) > 0 && len(fields[3]) <= 11 {
		copy(rmc.Lat[:], fields[3])
	}

	// 解析纬度方向
	if len(fields[4]) > 0 {
		copy(rmc.N_S[:], fields[4])
	}

	// 解析经度
	if len(fields[5]) > 0 && len(fields[5]) <= 12 {
		copy(rmc.Lon[:], fields[5])
	}

	// 解析经度方向
	if len(fields[6]) > 0 {
		copy(rmc.E_W[:], fields[6])
	}

	// 解析对地速度
	if len(fields[7]) > 0 && len(fields[7]) <= 10 {
		copy(rmc.SOG[:], fields[7])
	}

	// 解析对地真航向
	if len(fields[8]) > 0 && len(fields[8]) <= 6 {
		copy(rmc.COG[:], fields[8])
	}

	// 解析日期
	if len(fields[9]) > 0 && len(fields[9]) <= 6 {
		copy(rmc.Date[:], fields[9])
	}

	return rmc
}
