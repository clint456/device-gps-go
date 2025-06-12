# 完整的NMEA消息类型支持总结

## 概述

根据您提供的协议文档，GPS设备服务现在支持所有标准的NMEA消息类型的输出速率查询和设置功能。

## 支持的NMEA消息类型

根据协议文档中的消息组ID和消息子ID表：

| 子ID | NMEA类型 | 描述 | 支持状态 |
|------|----------|------|----------|
| 0x00 | GGA | 全球定位系统定位数据 | ✅ 完全支持 |
| 0x01 | GLL | 地理位置—纬度和经度 | ✅ 完全支持 |
| 0x02 | GSA | GNSS精度因子（DOP）与有效卫星 | ✅ 完全支持 |
| 0x03 | GRS | GNSS距离残差 | ✅ 完全支持 |
| 0x04 | GSV | 可视的GNSS卫星 | ✅ 完全支持 |
| 0x05 | RMC | 推荐的最少专有GNSS数据 | ✅ 完全支持 |
| 0x06 | VTG | 相对于地面的实际航向和速度 | ✅ 完全支持 |
| 0x07 | ZDA | 时间与日期（UTC时间，日、月、年） | ✅ 完全支持 |
| 0x08 | GST | GNSS伪距误差统计 | ✅ 完全支持 |

## 功能实现

### 1. 查询功能 (`getOutputRates`)

**支持的查询操作：**
- 一次性查询所有NMEA消息类型的输出速率
- 通过串口发送查询命令到GPS设备
- 解析设备响应并存储结果
- 返回格式化的输出速率信息

**查询顺序：**
```go
nmeaTypes := []struct {
    sid  NMEA_SUB_ID
    name string
}{
    {NMEA_GGA_SID, "GGA"},  // 0x00
    {NMEA_GLL_SID, "GLL"},  // 0x01
    {NMEA_GSA_SID, "GSA"},  // 0x02
    {NMEA_GRS_SID, "GRS"},  // 0x03
    {NMEA_GSV_SID, "GSV"},  // 0x04
    {NMEA_RMC_SID, "RMC"},  // 0x05
    {NMEA_VTG_SID, "VTG"},  // 0x06
    {NMEA_ZDA_SID, "ZDA"},  // 0x07
    {NMEA_GST_SID, "GST"},  // 0x08
}
```

### 2. 设置功能

#### 单个设置 (`setOutputRate`)
- 支持所有9种NMEA消息类型
- 格式：`"NMEA_TYPE:RATE"`
- 示例：`"GGA:1"`, `"GSV:5"`, `"GRS:0"`

#### 批量设置 (`setAllOutputRates`)
- 支持一次性设置多个NMEA消息类型
- 格式：`"TYPE1:RATE1,TYPE2:RATE2,TYPE3:RATE3"`
- 示例：`"GGA:1,GLL:1,GSA:1,GRS:0,GSV:5,RMC:1,VTG:1,ZDA:1,GST:1"`

### 3. 类型转换 (`getNMEASubID`)

完整的字符串到子ID的映射：
```go
switch strings.ToUpper(nmeaType) {
case "GGA": return NMEA_GGA_SID, nil  // 0x00
case "GLL": return NMEA_GLL_SID, nil  // 0x01
case "GSA": return NMEA_GSA_SID, nil  // 0x02
case "GRS": return NMEA_GRS_SID, nil  // 0x03
case "GSV": return NMEA_GSV_SID, nil  // 0x04
case "RMC": return NMEA_RMC_SID, nil  // 0x05
case "VTG": return NMEA_VTG_SID, nil  // 0x06
case "ZDA": return NMEA_ZDA_SID, nil  // 0x07
case "GST": return NMEA_GST_SID, nil  // 0x08
}
```

## API使用示例

### 查询所有输出速率
```bash
curl -X GET "http://localhost:59982/api/v3/device/name/GPS-Device/config"
```

**响应示例：**
```json
{
  "output_rates": "GGA: 1Hz, GLL: 1Hz, GSA: 1Hz, GRS: 禁用, GSV: 5Hz, RMC: 1Hz, VTG: 1Hz, ZDA: 1Hz, GST: 1Hz"
}
```

### 设置单个输出速率
```bash
# 设置GLL消息输出速率为1Hz
curl -X PUT "http://localhost:59982/api/v3/device/name/GPS-Device/set_rate" \
  -H "Content-Type: application/json" \
  -d '{"set_output_rate": "GLL:1"}'

# 禁用GRS消息输出
curl -X PUT "http://localhost:59982/api/v3/device/name/GPS-Device/set_rate" \
  -H "Content-Type: application/json" \
  -d '{"set_output_rate": "GRS:0"}'
```

### 批量设置所有输出速率
```bash
curl -X PUT "http://localhost:59982/api/v3/device/name/GPS-Device/set_all" \
  -H "Content-Type: application/json" \
  -d '{"set_all_rates": "GGA:1,GLL:1,GSA:1,GRS:0,GSV:5,RMC:1,VTG:1,ZDA:1,GST:1"}'
```

## 输出速率说明

| 数值 | 含义 | 描述 |
|------|------|------|
| 0 | 禁用 | 不输出该类型的NMEA消息 |
| 1 | 1Hz | 每秒输出1次 |
| 5 | 5Hz | 每秒输出5次 |
| 10 | 10Hz | 每秒输出10次 |

## 二进制协议支持

### 查询命令格式
```
帧头: 0xF1 0xD9
消息组ID: 0x06 (BIN_CFG_GID)
消息子ID: 0x01 (BM_MSG_SID)
长度: 0x02 0x00
有效载荷: [消息组ID] [消息子ID]
校验码: [CHK1] [CHK2]
```

### 设置命令格式
```
帧头: 0xF1 0xD9
消息组ID: 0x06 (BIN_CFG_GID)
消息子ID: 0x01 (BM_MSG_SID)
长度: 0x03 0x00
有效载荷: [消息组ID] [消息子ID] [输出速率]
校验码: [CHK1] [CHK2]
```

### 响应解析
- 设备会返回ACK/NAK确认消息
- 查询响应包含实际的输出速率值
- 响应数据存储在设备的OutputRates映射中

## 配置文件支持

### 设备资源 (device-gps.yml)
```yaml
- name: "output_rates"
  description: "Get all NMEA message output rates"
  valueType: "String"
  readWrite: "R"

- name: "set_output_rate"
  description: "Set NMEA message output rate"
  valueType: "String"
  readWrite: "W"

- name: "set_all_rates"
  description: "Set all NMEA message output rates"
  valueType: "String"
  readWrite: "W"
```

### 设备命令
```yaml
- name: "config"
  readWrite: "R"
  resourceOperations:
    - { deviceResource: "output_rates" }

- name: "set_rate"
  readWrite: "W"
  resourceOperations:
    - { deviceResource: "set_output_rate" }

- name: "set_all"
  readWrite: "W"
  resourceOperations:
    - { deviceResource: "set_all_rates" }
```

## 错误处理

### 支持的错误检查
- 无效的NMEA类型名称
- 超出范围的输出速率值
- 格式错误的配置字符串
- 设备通信失败
- 校验和验证失败

### 错误示例
```bash
# 无效的NMEA类型
{"set_output_rate": "INVALID:1"}  # 返回错误

# 无效的速率值
{"set_output_rate": "GGA:256"}    # 返回错误

# 格式错误
{"set_output_rate": "GGA"}        # 返回错误
```

## 总结

GPS设备服务现在完全支持协议文档中定义的所有9种NMEA消息类型：

1. **完整覆盖** - 支持所有标准NMEA消息类型
2. **灵活配置** - 支持单个和批量设置
3. **实时查询** - 通过串口实际查询设备状态
4. **错误处理** - 完善的输入验证和错误报告
5. **标准API** - 通过EdgeX标准接口提供服务

这个实现确保了与GPS设备协议的完全兼容性，并提供了用户友好的配置接口。
