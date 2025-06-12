# GPS设备Profile文件更新总结

## 概述

为了支持GPS设备的二进制协议配置功能，我们对设备配置文件 `device-gps.yml` 进行了扩展，添加了NMEA消息输出速率的查询和设置功能。

## 新增设备资源 (deviceResources)

### 1. 查询类资源

#### `output_rates`
- **描述**: 获取所有NMEA消息输出速率的人类易读格式
- **类型**: String (只读)
- **用途**: 一次性查询所有支持的NMEA消息类型及其当前输出速率
- **返回示例**: `"GGA: 1Hz, RMC: 1Hz, GSV: 5Hz, VTG: 1Hz"`

#### `nmea_output_rate`
- **描述**: 获取指定NMEA消息的输出速率
- **类型**: String (只读)
- **属性**: `nmea_type: "RMC"` (默认查询RMC类型)
- **用途**: 查询特定NMEA消息类型的输出速率
- **返回示例**: `"RMC: 1Hz"`

### 2. 设置类资源

#### `set_output_rate`
- **描述**: 设置单个NMEA消息的输出速率
- **类型**: Uint8 (只写)
- **范围**: 0-10 (0表示禁用，1-10表示Hz频率)
- **属性**: `nmea_type: "RMC"` (默认设置RMC类型)
- **用途**: 通过数值方式设置特定NMEA消息的输出频率

#### `set_nmea_output_rate`
- **描述**: 使用字符串格式设置NMEA消息输出速率
- **类型**: String (只写)
- **格式**: `"NMEA_TYPE:RATE"` 或 `"NMEA_TYPE=RATE"`
- **示例**: `"GGA:1"`, `"RMC=5"`
- **用途**: 通过字符串格式灵活设置单个NMEA消息的输出速率

#### `set_multiple_output_rates`
- **描述**: 批量设置多个NMEA消息的输出速率
- **类型**: String (只写)
- **格式**: `"TYPE1:RATE1,TYPE2:RATE2,TYPE3:RATE3"`
- **示例**: `"GGA:1,RMC:1,GSV:5,VTG:1"`
- **用途**: 一次性配置多个NMEA消息类型的输出速率

## 新增设备命令 (deviceCommands)

### 1. `config` - 配置查询命令
- **类型**: 只读 (R)
- **资源**: `output_rates`
- **用途**: 查询当前所有NMEA消息的输出速率配置
- **API端点**: `GET /api/v3/device/name/{deviceName}/config`

### 2. `get_nmea_rate` - 单个速率查询命令
- **类型**: 只读 (R)
- **资源**: `nmea_output_rate`
- **用途**: 查询特定NMEA消息类型的输出速率
- **API端点**: `GET /api/v3/device/name/{deviceName}/get_nmea_rate`

### 3. `set_single_rate` - 单个速率设置命令
- **类型**: 只写 (W)
- **资源**: `set_output_rate`
- **用途**: 设置单个NMEA消息的输出速率（数值方式）
- **API端点**: `PUT /api/v3/device/name/{deviceName}/set_single_rate`

### 4. `set_nmea_rate` - 字符串格式设置命令
- **类型**: 只写 (W)
- **资源**: `set_nmea_output_rate`
- **用途**: 使用字符串格式设置单个NMEA消息的输出速率
- **API端点**: `PUT /api/v3/device/name/{deviceName}/set_nmea_rate`

### 5. `set_batch_rates` - 批量设置命令
- **类型**: 只写 (W)
- **资源**: `set_multiple_output_rates`
- **用途**: 批量设置多个NMEA消息的输出速率
- **API端点**: `PUT /api/v3/device/name/{deviceName}/set_batch_rates`

## 支持的NMEA消息类型

| 类型 | 描述 | 子ID |
|------|------|------|
| GGA | 全球定位系统定位数据 | 0x00 |
| GLL | 地理位置—纬度和经度 | 0x01 |
| GSA | GNSS精度因子（DOP）与有效卫星 | 0x02 |
| GRS | GNSS距离残差 | 0x03 |
| GSV | 可视的GNSS卫星 | 0x04 |
| RMC | 推荐的最少专有GNSS数据 | 0x05 |
| VTG | 相对于地面的实际航向和速度 | 0x06 |
| ZDA | 时间与日期 | 0x07 |
| GST | GNSS伪距误差统计 | 0x08 |

## 输出速率说明

| 数值 | 含义 | 描述 |
|------|------|------|
| 0 | 禁用 | 不输出该类型的NMEA消息 |
| 1 | 1Hz | 每秒输出1次 |
| 5 | 5Hz | 每秒输出5次 |
| 10 | 10Hz | 每秒输出10次 |

## API使用示例

### 查询配置
```bash
# 查询所有NMEA消息输出速率
curl -X GET "http://localhost:59982/api/v3/device/name/GPS-Device/config"

# 响应示例
{
  "apiVersion": "v3",
  "statusCode": 200,
  "event": {
    "readings": [
      {
        "resourceName": "output_rates",
        "value": "GGA: 1Hz, RMC: 1Hz, GSV: 5Hz, VTG: 1Hz"
      }
    ]
  }
}
```

### 设置单个速率
```bash
# 设置GGA消息输出速率为1Hz
curl -X PUT "http://localhost:59982/api/v3/device/name/GPS-Device/set_nmea_rate" \
  -H "Content-Type: application/json" \
  -d '{"set_nmea_output_rate": "GGA:1"}'
```

### 批量设置速率
```bash
# 批量设置多个NMEA消息输出速率
curl -X PUT "http://localhost:59982/api/v3/device/name/GPS-Device/set_batch_rates" \
  -H "Content-Type: application/json" \
  -d '{"set_multiple_output_rates": "GGA:1,RMC:1,GSV:5,VTG:1,GSA:1"}'
```

## 配置属性说明

### primaryTable属性
- **LOCATION**: 位置相关数据（纬度、经度、海拔）
- **MOTION**: 运动相关数据（速度、航向）
- **TIME**: 时间相关数据（UTC时间）
- **STATUS**: 状态相关数据（定位质量、卫星数、精度因子、GPS状态）
- **CONFIG**: 配置相关数据（输出速率设置）

### nmea_type属性
用于指定默认的NMEA消息类型，主要用于单个消息类型的查询和设置操作。

## 向后兼容性

- 所有原有的GPS数据读取功能保持不变
- 原有的API端点继续正常工作
- 新增的配置功能是可选的，不影响基本GPS功能

## 部署注意事项

1. **重启服务**: 更新profile文件后需要重启GPS设备服务
2. **设备重新注册**: 可能需要重新注册设备以加载新的profile
3. **权限检查**: 确保EdgeX Core服务有权限访问更新后的profile文件
4. **测试验证**: 部署后应测试所有新增的配置功能

## 总结

通过这次profile文件更新，GPS设备服务现在支持：

1. **完整的配置查询功能** - 可以查询所有或特定NMEA消息的输出速率
2. **灵活的配置设置功能** - 支持单个和批量设置，支持数值和字符串格式
3. **标准化的API接口** - 通过EdgeX标准API提供配置功能
4. **用户友好的格式** - 支持易读的字符串格式配置
5. **完整的文档支持** - 详细的描述和使用示例

这些改进大大增强了GPS设备的可配置性和易用性，用户可以根据实际需求灵活调整NMEA消息的输出频率，优化数据传输和处理性能。
