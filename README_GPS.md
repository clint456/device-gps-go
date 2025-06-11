# EdgeX GPS设备服务

这是一个基于EdgeX框架的GPS设备服务，用于通过串口读取NMEA GPS数据并集成到EdgeX生态系统中。

## 🚀 功能特性

### 核心功能
- ✅ **串口通信**: 异步读取GPS设备的串口数据
- ✅ **NMEA解析**: 支持多种NMEA语句类型（RMC、GGA、GLL、GSA、GSV、VTG等）
- ✅ **数据缓存**: 智能缓存GPS数据，避免数据丢失
- ✅ **EdgeX集成**: 完全集成到EdgeX设备服务框架
- ✅ **实时监控**: 实时解析和显示GPS位置信息

### 支持的NMEA语句
- **RMC** (Recommended Minimum): 推荐最小定位信息
- **GGA** (Global Positioning System Fix Data): GPS定位数据
- **GLL** (Geographic Position): 地理位置
- **GSA** (GPS DOP and active satellites): GPS精度因子和活跃卫星
- **GSV** (GPS Satellites in view): 可见GPS卫星信息
- **VTG** (Track made good and Ground speed): 航向和地面速度

## 📁 项目结构

```
device-gps-go/
├── run/
│   ├── cmd/device-gps/           # 主程序入口
│   │   ├── main.go              # 主函数
│   │   ├── configuration.yaml   # 服务配置
│   │   ├── device.gps.yaml     # 设备配置文件
│   │   └── devices.yaml        # 设备实例配置
│   ├── driver/                  # GPS驱动实现
│   │   ├── gpsdriver.go        # EdgeX驱动接口实现
│   │   ├── gpsController.go    # GPS设备控制器
│   │   ├── nmea.go            # NMEA协议解析
│   │   ├── serial_port.go     # 串口通信
│   │   └── gps_test.go        # 单元测试
│   └── config/
│       └── configuration.go    # 配置结构定义
├── test_gps_parsing.go         # GPS解析测试程序
└── bin/
    └── device-gps              # 编译后的可执行文件
```

## 🛠️ 技术架构

### 数据流程
```
串口数据 → UartRX_Task → 数据缓冲区 → NMEA解析 → GPS数据结构 → EdgeX命令处理
```

### 关键组件

1. **UartRX_Task**: 异步串口读取任务
   - 使用独立的goroutine处理串口数据
   - 智能缓冲区管理，防止数据溢出
   - 错误处理和自动恢复

2. **NMEA解析器**: 
   - 支持多种NMEA语句格式
   - 校验和验证确保数据完整性
   - 容错处理，跳过无效数据

3. **EdgeX驱动接口**:
   - 实现标准的ProtocolDriver接口
   - 支持读取GPS位置、速度、航向等数据
   - 自动事件推送到EdgeX核心服务

## 📊 支持的数据类型

| 资源名称 | 数据类型 | 单位 | 描述 |
|---------|---------|------|------|
| latitude | Float64 | 度 | GPS纬度（十进制度数） |
| longitude | Float64 | 度 | GPS经度（十进制度数） |
| altitude | Float64 | 米 | 海拔高度 |
| speed | Float64 | km/h | 地面速度 |
| course | Float64 | 度 | 航向角度 |
| utc_time | String | - | UTC时间 |
| fix_quality | Int32 | - | 定位质量（0=无效，1=有效） |
| satellites_used | Int32 | - | 使用的卫星数量 |
| hdop | Float64 | - | 水平精度因子 |
| gps_status | String | - | GPS状态 |

## 🔧 配置说明

### 串口配置
```yaml
GPSCustom:
  SerialPort: "/dev/ttyUSB0"  # 串口设备路径
  BaudRate: 9600              # 波特率
  ReadTimeout: "100ms"        # 读取超时
  Debug: true                 # 调试模式
```

### 设备协议配置
```yaml
protocols:
  UART:
    deviceLocation: "/dev/ttyUSB0"
    baudRate: "9600"
    dataBits: "8"
    stopBits: "1"
    parity: "none"
```

## 🚀 快速开始

### 1. 编译项目
```bash
cd /home/clint/EdgeX/device-gps-go
go build -o bin/device-gps ./run/cmd/device-gps
```

### 2. 运行测试
```bash
# 运行NMEA解析测试
go run test_gps_parsing.go

# 运行单元测试
go test ./run/driver -v
```

### 3. 启动GPS设备服务
```bash
./bin/device-gps
```

## 📝 示例NMEA数据

项目支持解析以下格式的NMEA数据：

```
$GBGGA,055525.000,3044.368753,N,10357.548051,E,1,04,2.40,129.3,M,-32.3,M,,*5A
$GBRMC,055525.000,A,3044.368753,N,10357.548051,E,0.00,000.00,100625,,,A,C*13
$GBGLL,3044.368753,N,10357.548051,E,055525.000,A,A*4B
$GBGSA,A,2,34,21,07,44,,,,,,,,,2.59,2.40,1.00,4*03
$GBGSV,6,1,21,10,80,005,27,34,76,067,33,38,75,161,28,21,57,046,29,1*71
$GBVTG,000.00,T,,M,0.00,N,0.00,K,A*2F
```

## 🔍 API接口

### 读取GPS位置
```
GET /api/v3/device/name/{deviceName}/location
```

### 读取GPS状态
```
GET /api/v3/device/name/{deviceName}/status
```

### 读取所有GPS数据
```
GET /api/v3/device/name/{deviceName}/all_data
```

## 🛡️ 错误处理

- **串口连接失败**: 自动重试连接
- **数据解析错误**: 跳过无效数据，继续处理
- **缓冲区溢出**: 自动清空缓冲区，防止内存泄漏
- **校验和错误**: 丢弃无效的NMEA语句

## 📈 性能优化

- 使用异步goroutine处理串口数据
- 智能缓冲区管理，避免频繁内存分配
- 高效的NMEA解析算法
- 最小化锁竞争，提高并发性能

## 🔧 故障排除

### 常见问题

1. **串口权限问题**
   ```bash
   sudo chmod 666 /dev/ttyUSB0
   ```

2. **GPS设备未连接**
   - 检查USB连接
   - 确认设备路径正确

3. **NMEA数据格式错误**
   - 检查GPS设备输出格式
   - 确认波特率设置正确

## 📄 许可证

本项目基于Apache 2.0许可证开源。

## 🤝 贡献

欢迎提交Issue和Pull Request来改进这个项目！
