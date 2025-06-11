package driver

import (
	"strings"
	"testing"
)

func TestParsNMEAType(t *testing.T) {
	tests := []struct {
		sentence string
		expected NMEA_TYPE
	}{
		{"$GBGGA,055525.000,3044.368753,N,10357.548051,E,1,   ,2.40,129.3,M,-32.3,M,,*5A", NMEA_GGA_TYPE},
		{"$GBRMC,055525.000,A,3044.368753,N,10357.548051,E,0.00,000.00,100625,,,A,C*13", NMEA_RMC_TYPE},
		{"$GBGLL,3044.368753,N,10357.548051,E,055525.000,A,A*4B", NMEA_GLL_TYPE},
		{"$GBGSA,A,2,34,21,07,44,,,,,,,,,2.59,2.40,1.00,4*03", NMEA_GSA_TYPE},
		{"$GBGSV,6,1,21,10,80,005,27,34,76,067,33,38,75,161,28,21,57,046,29,1*71", NMEA_GSV_TYPE},
		{"$GBVTG,000.00,T,,M,0.00,N,0.00,K,A*2F", NMEA_VTG_TYPE},
		{"$INVALID", NMEA_UNKONW_TYPE},
	}

	for _, test := range tests {
		result := ParsNMEAType(test.sentence, len(test.sentence))
		if result != test.expected {
			t.Errorf("ParsNMEAType(%s) = %v, expected %v", test.sentence, result, test.expected)
		}
	}
}

func TestValidateNMEAChecksum(t *testing.T) {
	tests := []struct {
		sentence string
		expected bool
	}{
		{"$GBGGA,055525.000,3044.368753,N,10357.548051,E,1,   ,2.40,129.3,M,-32.3,M,,*5A", true},
		{"$GBRMC,055525.000,A,3044.368753,N,10357.548051,E,0.00,000.00,100625,,,A,C*13", true},
		{"$GBGLL,3044.368753,N,10357.548051,E,055525.000,A,A*4B", true},
		{"$GBGGA,055525.000,3044.368753,N,10357.548051,E,1,   ,2.40,129.3,M,-32.3,M,,*FF", false}, // 错误的校验和
		{"$INVALID*00", false}, // 无效的校验和
	}

	for _, test := range tests {
		result := ValidateNMEAChecksum(test.sentence, len(test.sentence))
		if result != test.expected {
			t.Errorf("ValidateNMEAChecksum(%s) = %v, expected %v", test.sentence, result, test.expected)
		}
	}
}

func TestParsNMEARMC(t *testing.T) {
	sentence := "$GBRMC,055525.000,A,3044.368753,N,10357.548051,E,0.00,000.00,100625,,,A,C*13"
	
	rmc := ParsNMEARMC(sentence, len(sentence))
	if rmc == nil {
		t.Fatal("ParsNMEARMC returned nil")
	}

	// 检查解析的字段
	utc := strings.TrimRight(string(rmc.UTC[:]), "\x00")
	if utc != "055525.000" {
		t.Errorf("UTC = %s, expected 055525.000", utc)
	}

	status := strings.TrimRight(string(rmc.Status[:]), "\x00")
	if status != "A" {
		t.Errorf("Status = %s, expected A", status)
	}

	lat := strings.TrimRight(string(rmc.Lat[:]), "\x00")
	if lat != "3044.368753" {
		t.Errorf("Lat = %s, expected 3044.368753", lat)
	}

	ns := strings.TrimRight(string(rmc.N_S[:]), "\x00")
	if ns != "N" {
		t.Errorf("N_S = %s, expected N", ns)
	}

	lon := strings.TrimRight(string(rmc.Lon[:]), "\x00")
	if lon != "10357.548051" {
		t.Errorf("Lon = %s, expected 10357.548051", lon)
	}

	ew := strings.TrimRight(string(rmc.E_W[:]), "\x00")
	if ew != "E" {
		t.Errorf("E_W = %s, expected E", ew)
	}
}

func TestParsNMEAGGA(t *testing.T) {
	sentence := "$GBGGA,055525.000,3044.368753,N,10357.548051,E,1,04,2.40,129.3,M,-32.3,M,,*5A"
	
	gga := ParsNMEAGGA(sentence, len(sentence))
	if gga == nil {
		t.Fatal("ParsNMEAGGA returned nil")
	}

	// 检查解析的字段
	utc := strings.TrimRight(string(gga.UTC[:]), "\x00")
	if utc != "055525.000" {
		t.Errorf("UTC = %s, expected 055525.000", utc)
	}

	quality := strings.TrimRight(string(gga.Quality[:]), "\x00")
	if quality != "1" {
		t.Errorf("Quality = %s, expected 1", quality)
	}

	numSat := strings.TrimRight(string(gga.NumSatUsed[:]), "\x00")
	if numSat != "04" {
		t.Errorf("NumSatUsed = %s, expected 04", numSat)
	}
}

func TestConvertDMSToDecimal(t *testing.T) {
	driver := &Driver{}
	
	tests := []struct {
		dms       string
		direction string
		expected  float64
	}{
		{"3044.368753", "N", 30.739479216666666},
		{"10357.548051", "E", 103.95913418333333},
		{"3044.368753", "S", -30.739479216666666},
		{"10357.548051", "W", -103.95913418333333},
		{"", "N", 0.0},
		{"invalid", "N", 0.0},
	}

	for _, test := range tests {
		result := driver.convertDMSToDecimal(test.dms, test.direction)
		if result != test.expected {
			t.Errorf("convertDMSToDecimal(%s, %s) = %f, expected %f", 
				test.dms, test.direction, result, test.expected)
		}
	}
}
