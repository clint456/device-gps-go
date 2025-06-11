#!/bin/bash

# EdgeX GPS设备服务启动脚本

echo "🚀 启动EdgeX GPS设备服务"
echo "========================"

# 检查可执行文件是否存在
if [ ! -f "./bin/device-gps" ]; then
    echo "📦 编译GPS设备服务..."
    go build -o bin/device-gps ./run/cmd/device-gps
    if [ $? -ne 0 ]; then
        echo "❌ 编译失败！"
        exit 1
    fi
    echo "✅ 编译成功！"
fi

# 检查串口设备
SERIAL_PORT="/dev/ttyUSB0"
if [ ! -e "$SERIAL_PORT" ]; then
    echo "⚠️  警告: 串口设备 $SERIAL_PORT 不存在"
    echo "   请确保GPS设备已连接，或修改配置文件中的串口路径"
    echo "   可用的串口设备："
    ls /dev/tty* | grep -E "(USB|ACM)" || echo "   未找到USB串口设备"
    echo ""
fi

# 检查串口权限
if [ -e "$SERIAL_PORT" ]; then
    if [ ! -r "$SERIAL_PORT" ] || [ ! -w "$SERIAL_PORT" ]; then
        echo "⚠️  警告: 没有串口设备的读写权限"
        echo "   请运行: sudo chmod 666 $SERIAL_PORT"
        echo "   或者将当前用户添加到dialout组: sudo usermod -a -G dialout $USER"
        echo ""
    fi
fi

# 显示配置信息
echo "📋 配置信息:"
echo "   串口设备: $SERIAL_PORT"
echo "   波特率: 9600"
echo "   服务端口: 59999"
echo ""

# 启动服务
echo "🔄 启动GPS设备服务..."
echo "   按 Ctrl+C 停止服务"
echo ""

# 切换到正确的目录
cd run/cmd/device-gps

# 启动服务
../../../bin/device-gps
