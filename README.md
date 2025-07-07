# EdgeX GPS设备服务

这是一个基于EdgeX框架的GPS设备服务，用于通过串口读取NMEA GPS数据并集成到EdgeX生态系统中。

## 🚀 功能特性

### 核心功能
- ✅ **串口通信**: 异步读取GPS设备的串口数据
- ✅ **NMEA解析**: 支持多种NMEA语句类型（RMC、GGA、GLL、GSA、GSV、VTG等）
- ✅ **数据格式化**: 将原始GPS数据转换为人类易读的格式
- ✅ **智能0值处理**: 正确处理静止、正北等有效的0值状态
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

### 数据格式化输出

所有GPS数据都经过格式化处理，提供人类易读的输出格式：

| 资源名称 | 数据类型 | 格式示例 | 描述 |
|---------|---------|----------|------|
| latitude | String | `39°58'08.6"N` | GPS纬度（度分秒格式） |
| longitude | String | `116°29'30.0"E` | GPS经度（度分秒格式） |
| altitude | String | `123.4 米` | 海拔高度（带单位） |
| speed | String | `23.15 km/h` | 地面速度（带单位） |
| course | String | `45.2° (东北)` | 航向角度（带方向描述） |
| utc_time | String | `12:34:56.00` | UTC时间（HH:MM:SS格式） |
| fix_quality | String | `GPS定位` | 定位质量（描述性文本） |
| satellites_used | String | `8 颗卫星` | 使用的卫星数量（描述性文本） |
| hdop | String | `1.20 (良好)` | 水平精度因子（带质量评估） |
| gps_status | String | `ACTIVE` | GPS状态 |

### 特殊值处理

系统能够正确处理以下有效的0值情况：

- **0°00'00.0"** - 赤道和本初子午线交点坐标
- **0.0 米** - 海平面高度
- **0.00 km/h** - 静止状态（停车、等红灯）
- **0.0° (北)** - 正北方向
- **00:00:00.00** - 午夜时间
- **无定位** - GPS信号丢失状态
- **0 颗卫星** - 室内或信号遮挡环境

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

# 运行特定测试
go test ./run/driver -run TestDataFormatting -v      # 数据格式化测试
go test ./run/driver -run TestZeroValueHandling -v   # 0值处理测试
go test ./run/driver -run TestNMEAParsing -v         # NMEA解析测试
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
**响应示例：**
```json
{
  "latitude": "39°58'08.6\"N",
  "longitude": "116°29'30.0\"E",
  "altitude": "123.4 米"
}
```

### 读取GPS状态
```
GET /api/v3/device/name/{deviceName}/status
```
**响应示例：**
```json
{
  "fix_quality": "GPS定位",
  "satellites_used": "8 颗卫星",
  "hdop": "1.20 (良好)",
  "gps_status": "ACTIVE"
}
```

### 读取运动信息
```
GET /api/v3/device/name/{deviceName}/motion
```
**响应示例：**
```json
{
  "speed": "23.15 km/h",
  "course": "45.2° (东北)"
}
```

### 读取所有GPS数据
```
GET /api/v3/device/name/{deviceName}/all_data
```
**响应示例：**
```json
{
  "latitude": "39°58'08.6\"N",
  "longitude": "116°29'30.0\"E",
  "altitude": "123.4 米",
  "speed": "23.15 km/h",
  "course": "45.2° (东北)",
  "utc_time": "12:34:56.00",
  "fix_quality": "GPS定位",
  "satellites_used": "8 颗卫星",
  "hdop": "1.20 (良好)",
  "gps_status": "ACTIVE"
}
```

## 🛡️ 错误处理与数据完整性

### 错误处理机制
- **串口连接失败**: 自动重试连接
- **数据解析错误**: 跳过无效数据，继续处理
- **缓冲区溢出**: 自动清空缓冲区，防止内存泄漏
- **校验和错误**: 丢弃无效的NMEA语句

### 数据完整性保护
- **0值智能处理**: 区分"无数据"和"有效0值"
- **多数据源验证**: 从多个NMEA语句获取数据，提高可靠性
- **数据有效性检查**: 验证坐标格式和方向标识
- **容错机制**: 单个数据源失败不影响其他数据获取

### 特殊情况处理
- **静止状态**: 正确处理0速度（车辆停止）
- **正北方向**: 正确处理0航向（指向正北）
- **特殊位置**: 正确处理赤道/本初子午线坐标
- **信号丢失**: 优雅处理GPS信号中断情况

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
