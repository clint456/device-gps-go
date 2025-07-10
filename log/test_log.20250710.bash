clint@Clinton:~/EdgeX/device-gps-go/run/cmd/device-gps$ ./device-gps -o -d -cp
level=INFO ts=2025-07-10T14:50:39.701360812+08:00 app=device-gps source=config.go:739 msg="Using Configuration provider (keeper) from: http://localhost:59890 with base path of edgex/v4/core-common-config-bootstrapper/all-services"
level=INFO ts=2025-07-10T14:50:39.712309947+08:00 app=device-gps source=config.go:501 msg="loading the common configuration for service type device-service"
level=INFO ts=2025-07-10T14:50:39.712454419+08:00 app=device-gps source=config.go:739 msg="Using Configuration provider (keeper) from: http://localhost:59890 with base path of edgex/v4/core-common-config-bootstrapper/device-services"
level=INFO ts=2025-07-10T14:50:39.716032423+08:00 app=device-gps source=config.go:274 msg="Common configuration loaded from the Configuration Provider. No overrides applied"
level=INFO ts=2025-07-10T14:50:39.716087524+08:00 app=device-gps source=config.go:739 msg="Using Configuration provider (keeper) from: http://localhost:59890 with base path of edgex/v4/device-gps"
level=INFO ts=2025-07-10T14:50:39.716928758+08:00 app=device-gps source=config.go:752 msg="Loading configuration file from res/configuration.yaml"
level=INFO ts=2025-07-10T14:50:39.717377789+08:00 app=device-gps source=config.go:361 msg="Private configuration loaded from file with 0 overrides applied"
level=INFO ts=2025-07-10T14:50:39.758692582+08:00 app=device-gps source=config.go:330 msg="Private configuration has been pushed to into Configuration Provider with overrides applied"
level=INFO ts=2025-07-10T14:50:39.758756381+08:00 app=device-gps source=config.go:334 msg="listening for private config changes"
level=INFO ts=2025-07-10T14:50:39.758779269+08:00 app=device-gps source=config.go:336 msg="listening for all services common config changes"
level=INFO ts=2025-07-10T14:50:39.758792454+08:00 app=device-gps source=config.go:343 msg="listening for device service common config changes"
level=INFO ts=2025-07-10T14:50:39.758797576+08:00 app=device-gps source=config.go:185 msg="listening for private config changes"
level=INFO ts=2025-07-10T14:50:39.758812866+08:00 app=device-gps source=config.go:187 msg="listening for all services common config changes"
level=INFO ts=2025-07-10T14:50:39.758817656+08:00 app=device-gps source=config.go:194 msg="listening for device service common config changes"
level=INFO ts=2025-07-10T14:50:39.758907948+08:00 app=device-gps source=httpserver.go:152 msg="Web server starting (0.0.0.0:59999)"
level=DEBUG ts=2025-07-10T14:50:39.75901923+08:00 app=device-gps source=ziti.go:187 msg="service security option Mode = "
level=INFO ts=2025-07-10T14:50:39.759039937+08:00 app=device-gps source=ziti.go:231 msg="listening on underlay network. ListenMode '' at 0.0.0.0:59999"
level=INFO ts=2025-07-10T14:50:39.760277064+08:00 app=device-gps source=messaging.go:104 msg="Connected to mqtt Message Bus @ mqtt://localhost:1883 with AuthMode='none'"
level=INFO ts=2025-07-10T14:50:39.760300957+08:00 app=device-gps source=command.go:35 msg="Subscribing to command requests on topic: edgex/device/command/request/device-gps/#"
level=INFO ts=2025-07-10T14:50:39.760316001+08:00 app=device-gps source=command.go:39 msg="Responses to command requests will be published on topic: edgex/response/device-gps/<requestId>"
level=INFO ts=2025-07-10T14:50:39.760669722+08:00 app=device-gps source=callback.go:36 msg="Subscribing to System Events on topics: edgex/system-events/core-metadata/+/+/device-gps/# and edgex/system-events/core-metadata/deviceprofile/delete/#"
level=INFO ts=2025-07-10T14:50:39.761602623+08:00 app=device-gps source=validation.go:30 msg="Subscribing to device validation requests on topic: edgex/device-gps/validate/device"
level=INFO ts=2025-07-10T14:50:39.761635537+08:00 app=device-gps source=validation.go:34 msg="Responses to device validation requests will be published on topic: edgex/response/device-gps/<requestId>"
level=INFO ts=2025-07-10T14:50:39.761973408+08:00 app=device-gps source=manager.go:128 msg="Metrics Manager started with a report interval of 30s"
level=INFO ts=2025-07-10T14:50:39.762007148+08:00 app=device-gps source=clients.go:87 msg="Using REST for 'security-proxy-auth' clients @ http://localhost:59842"
level=INFO ts=2025-07-10T14:50:39.762024013+08:00 app=device-gps source=clients.go:87 msg="Using REST for 'core-metadata' clients @ http://localhost:59881"
level=INFO ts=2025-07-10T14:50:39.762057126+08:00 app=device-gps source=restrouter.go:55 msg="Registering routes..."
level=DEBUG ts=2025-07-10T14:50:39.76208121+08:00 app=device-gps source=init.go:151 msg="Check service 'core-metadata' availability by Ping"
level=INFO ts=2025-07-10T14:50:39.763231356+08:00 app=device-gps source=init.go:170 msg="Check service 'core-metadata' availability succeeded"
level=INFO ts=2025-07-10T14:50:39.765797257+08:00 app=device-gps source=devices.go:75 msg="LastConnected-GPS-Device-01 metric has been registered and will be reported (if enabled)"
level=DEBUG ts=2025-07-10T14:50:39.769194678+08:00 app=device-gps source=service.go:294 msg="trying to find device service device-gps"
level=INFO ts=2025-07-10T14:50:39.770297956+08:00 app=device-gps source=service.go:312 msg="device service device-gps exists, updating it"
level=INFO ts=2025-07-10T14:50:39.774053016+08:00 app=device-gps source=profiles.go:89 msg="Loading pre-defined Device Profiles from /home/clint/EdgeX/device-gps-go/run/cmd/device-gps/res/profiles(1 files found)"
level=INFO ts=2025-07-10T14:50:39.777991277+08:00 app=device-gps source=profiles.go:190 msg="Device Profile GPS-Device exists, using the existing one"
level=INFO ts=2025-07-10T14:50:39.77817223+08:00 app=device-gps source=devices.go:107 msg="Loading pre-defined Devices from /home/clint/EdgeX/device-gps-go/run/cmd/device-gps/res/devices(1 files found)"
level=INFO ts=2025-07-10T14:50:39.778765676+08:00 app=device-gps source=devices.go:187 msg="Device GPS-Device-01 exists, using the existing one"
level=DEBUG ts=2025-07-10T14:50:39.778871467+08:00 app=device-gps source=utils.go:100 msg="EventsSent metric has been registered and will be reported (if enabled)"
level=DEBUG ts=2025-07-10T14:50:39.778892524+08:00 app=device-gps source=utils.go:100 msg="ReadingsSent metric has been registered and will be reported (if enabled)"
level=INFO ts=2025-07-10T14:50:39.778912927+08:00 app=device-gps source=autodiscovery.go:32 msg="AutoDiscovery stopped: disabled by configuration"
level=INFO ts=2025-07-10T14:50:39.778931585+08:00 app=device-gps source=message.go:50 msg="Service dependencies resolved..."
level=INFO ts=2025-07-10T14:50:39.778940748+08:00 app=device-gps source=message.go:51 msg="Starting device-gps 0.0.0 "
level=INFO ts=2025-07-10T14:50:39.77894754+08:00 app=device-gps source=message.go:55 msg="device simple started"
level=INFO ts=2025-07-10T14:50:39.778951808+08:00 app=device-gps source=message.go:58 msg="Service started in: 77.804096ms"
level=INFO ts=2025-07-10T14:50:39.778958837+08:00 app=device-gps source=bootstrap.go:254 msg="SecuritySecretsRequested metric registered and will be reported (if enabled)"
level=INFO ts=2025-07-10T14:50:39.778981203+08:00 app=device-gps source=bootstrap.go:254 msg="SecuritySecretsStored metric registered and will be reported (if enabled)"
level=DEBUG ts=2025-07-10T14:50:39.779000838+08:00 app=device-gps source=gpsdriver.go:74 msg="Driver.HandleReadCommands(): protocol = UART, device location = /dev/ttyUSB0, baud rate = 9600 readTimeout=100 dataBits %!v(MISSING) "
level=INFO ts=2025-07-10T14:50:39.779017172+08:00 app=device-gps source=gpsdriver.go:78 msg="🚀 初始化GPS设备服务"
level=INFO ts=2025-07-10T14:50:39.788743969+08:00 app=device-gps source=gpsdriver.go:88 msg="✅ GPS设备初始化成功"
level=DEBUG ts=2025-07-10T14:50:39.803722964+08:00 app=device-gps source=callback.go:84 msg="System event received on message queue. Topic: edgex/system-events/core-metadata/deviceservice/update/device-gps, Correlation-id: cf2bce94-67c5-41e1-aee2-c9ad25bde6b5"
level=DEBUG ts=2025-07-10T14:50:39.803836088+08:00 app=device-gps source=callback.go:285 msg="device service updated"
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065041.000, 纬度=3044.307697N, 经度=10357.711578E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065042.000, 纬度=3044.307683N, 经度=10357.711563E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307683N, 经度=10357.711563E, 时间=065042.000, 状态=A
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065042.000, 纬度=3044.307683N, 经度=10357.711563E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065043.000, 纬度=3044.307681N, 经度=10357.711561E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307681N, 经度=10357.711561E, 时间=065043.000, 状态=A
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065043.000, 纬度=3044.307681N, 经度=10357.711561E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065044.000, 纬度=3044.307682N, 经度=10357.711559E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307682N, 经度=10357.711559E, 时间=065044.000, 状态=A
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065044.000, 纬度=3044.307682N, 经度=10357.711559E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065045.000, 纬度=3044.307675N, 经度=10357.711548E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307675N, 经度=10357.711548E, 时间=065045.000, 状态=A
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065045.000, 纬度=3044.307675N, 经度=10357.711548E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065046.000, 纬度=3044.307666N, 经度=10357.711533E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307666N, 经度=10357.711533E, 时间=065046.000, 状态=A
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065046.000, 纬度=3044.307666N, 经度=10357.711533E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065047.000, 纬度=3044.307664N, 经度=10357.711526E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307664N, 经度=10357.711526E, 时间=065047.000, 状态=A
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065047.000, 纬度=3044.307664N, 经度=10357.711526E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065048.000, 纬度=3044.307661N, 经度=10357.711517E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307661N, 经度=10357.711517E, 时间=065048.000, 状态=A
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065048.000, 纬度=3044.307661N, 经度=10357.711517E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065049.000, 纬度=3044.307644N, 经度=10357.711514E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307644N, 经度=10357.711514E, 时间=065049.000, 状态=A
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065049.000, 纬度=3044.307644N, 经度=10357.711514E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065050.000, 纬度=3044.307637N, 经度=10357.711505E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307637N, 经度=10357.711505E, 时间=065050.000, 状态=A
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065050.000, 纬度=3044.307637N, 经度=10357.711505E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065051.000, 纬度=3044.307641N, 经度=10357.711502E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307641N, 经度=10357.711502E, 时间=065051.000, 状态=A
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065051.000, 纬度=3044.307641N, 经度=10357.711502E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065052.000, 纬度=3044.307643N, 经度=10357.711491E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307643N, 经度=10357.711491E, 时间=065052.000, 状态=A
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065052.000, 纬度=3044.307643N, 经度=10357.711491E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065053.000, 纬度=3044.307648N, 经度=10357.711481E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307648N, 经度=10357.711481E, 时间=065053.000, 状态=A
level=DEBUG ts=2025-07-10T14:50:52.982286587+08:00 app=device-gps source=gpsdriver.go:151 msg="✍️ 处理设备 GPS-Device-01 的写入命令"
level=DEBUG ts=2025-07-10T14:50:52.982322281+08:00 app=device-gps source=gpsdriver.go:158 msg="处理写入资源: set_all_rates"
level=INFO ts=2025-07-10T14:50:52.982334892+08:00 app=device-gps source=gpsdriver.go:662 msg="开始批量设置输出速率: GGA:1,RMC:1,GSV:1,VTG:1,GSA:1,GLL:5"
📤 发送二进制命令: F1D906010300F00001FB10
level=INFO ts=2025-07-10T14:50:52.982380432+08:00 app=device-gps source=gpsdriver.go:687 msg=成功设置GGA输出速率为1
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=14
📤 发送二进制命令: F1D906010300F00501001A
level=INFO ts=2025-07-10T14:50:53.082989137+08:00 app=device-gps source=gpsdriver.go:687 msg=成功设置RMC输出速率为1
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=14
📤 发送二进制命令: F1D906010300F00401FF18
level=INFO ts=2025-07-10T14:50:53.183329727+08:00 app=device-gps source=gpsdriver.go:687 msg=成功设置GSV输出速率为1
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=14
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=14
📤 发送二进制命令: F1D906010300F00601011C
level=INFO ts=2025-07-10T14:50:53.284014277+08:00 app=device-gps source=gpsdriver.go:687 msg=成功设置VTG输出速率为1
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
📤 发送二进制命令: F1D906010300F00201FD14
level=INFO ts=2025-07-10T14:50:53.384415833+08:00 app=device-gps source=gpsdriver.go:687 msg=成功设置GSA输出速率为1
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
📤 发送二进制命令: F1D906010300F001050016
level=INFO ts=2025-07-10T14:50:53.485166722+08:00 app=device-gps source=gpsdriver.go:687 msg=成功设置GLL输出速率为5
✅ RMC: 时间=065053.000, 纬度=3044.307648N, 经度=10357.711481E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
✅ 收到ACK确认: GroupID=0x06, SubID=0x01
二进制协议解析错误: 二进制消息太短
level=INFO ts=2025-07-10T14:50:53.586071615+08:00 app=device-gps source=gpsdriver.go:697 msg=批量设置输出速率完成
level=DEBUG ts=2025-07-10T14:50:53.586164558+08:00 app=device-gps source=command.go:111 msg="SET Device Command successfully. Device: GPS-Device-01, Source: set_all_rates, X-Correlation-ID: cd2539fd-f07f-41c0-bc3d-024783390e60"
✅ 收到ACK确认: GroupID=0x06, SubID=0x01
✅ 收到ACK确认: GroupID=0x06, SubID=0x01
✅ 收到ACK确认: GroupID=0x06, SubID=0x01
二进制协议解析错误: 二进制消息太短
✅ 收到ACK确认: GroupID=0x06, SubID=0x01
✅ 收到ACK确认: GroupID=0x06, SubID=0x01
串口读取错误: EOF
✅ GGA: 时间=065054.000, 纬度=3044.307630N, 经度=10357.711481E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065054.000, 纬度=3044.307630N, 经度=10357.711481E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065055.000, 纬度=3044.307623N, 经度=10357.711485E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065055.000, 纬度=3044.307623N, 经度=10357.711485E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065056.000, 纬度=3044.307614N, 经度=10357.711485E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065056.000, 纬度=3044.307614N, 经度=10357.711485E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065057.000, 纬度=3044.307607N, 经度=10357.711467E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065057.000, 纬度=3044.307607N, 经度=10357.711467E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
level=DEBUG ts=2025-07-10T14:50:57.526428989+08:00 app=device-gps source=gpsdriver.go:95 msg="📖 处理设备 GPS-Device-01 的读取命令"
level=DEBUG ts=2025-07-10T14:50:57.52646742+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: latitude"
level=DEBUG ts=2025-07-10T14:50:57.526498107+08:00 app=device-gps source=gpsdriver.go:111 msg="latitude: DeviceResource: latitude, String: 30°44'18.5\"N"
level=DEBUG ts=2025-07-10T14:50:57.526511003+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: longitude"
level=DEBUG ts=2025-07-10T14:50:57.52651698+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: altitude"
level=DEBUG ts=2025-07-10T14:50:57.526595963+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: latitude reading: {Id:a8fc0443-7956-483d-b9e3-c75c7486d611 Origin:1752130257526529335 DeviceName:GPS-Device-01 ResourceName:latitude ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:30°44'18.5\"N} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:50:57.526619151+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: longitude reading: {Id:f06b1477-3d95-4edc-b9d6-ea78a1fe64a6 Origin:1752130257526529335 DeviceName:GPS-Device-01 ResourceName:longitude ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:103°57'42.7\"E} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:50:57.526634888+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: altitude reading: {Id:5d2568bd-4185-428f-b63d-97b2f0ba98ff Origin:1752130257526529335 DeviceName:GPS-Device-01 ResourceName:altitude ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:519.1 米} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:50:57.52664365+08:00 app=device-gps source=command.go:72 msg="GET Device Command successfully. Device: GPS-Device-01, Source: location, X-Correlation-ID: e0ed63e9-886f-4ed8-9980-8f344286f3f3"
level=DEBUG ts=2025-07-10T14:50:57.527116623+08:00 app=device-gps source=utils.go:72 msg="Event(profileName: GPS-Device, deviceName: GPS-Device-01, sourceName: location, id: 467bdc22-974b-4e29-881c-ffe5f2d197d9) published to MessageBus on topic: edgex/events/device/device-gps/GPS-Device/GPS-Device-01/location"
✅ GGA: 时间=065058.000, 纬度=3044.307607N, 经度=10357.711453E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307607N, 经度=10357.711453E, 时间=065058.000, 状态=A
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065058.000, 纬度=3044.307607N, 经度=10357.711453E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065059.000, 纬度=3044.307590N, 经度=10357.711453E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065059.000, 纬度=3044.307590N, 经度=10357.711453E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065100.000, 纬度=3044.307586N, 经度=10357.711432E, 质量=1, 卫星数=10
✅ GSA: 模式=A, 定位模式=3, PDOP=1.86, HDOP=0.99, VDOP=1.58
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065100.000, 纬度=3044.307586N, 经度=10357.711432E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
level=DEBUG ts=2025-07-10T14:51:00.260000355+08:00 app=device-gps source=gpsdriver.go:95 msg="📖 处理设备 GPS-Device-01 的读取命令"
level=DEBUG ts=2025-07-10T14:51:00.260041473+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: speed"
level=DEBUG ts=2025-07-10T14:51:00.260054249+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: course"
level=DEBUG ts=2025-07-10T14:51:00.260130918+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: speed reading: {Id:f70bb819-0720-4b5e-83bc-e1902e54dc43 Origin:1752130260260068069 DeviceName:GPS-Device-01 ResourceName:speed ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:0.00 km/h} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:00.260152336+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: course reading: {Id:d0f6dbfc-c934-41f9-98e9-4ec5845e213b Origin:1752130260260068069 DeviceName:GPS-Device-01 ResourceName:course ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:0.0° (北)} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:00.26016774+08:00 app=device-gps source=command.go:72 msg="GET Device Command successfully. Device: GPS-Device-01, Source: motion, X-Correlation-ID: 7f95d49a-1d02-4b6c-be6e-5caf5ca9352c"
level=DEBUG ts=2025-07-10T14:51:00.260386468+08:00 app=device-gps source=utils.go:72 msg="Event(profileName: GPS-Device, deviceName: GPS-Device-01, sourceName: motion, id: cd165cf2-1b7b-496a-91d4-364f0dd9b0b6) published to MessageBus on topic: edgex/events/device/device-gps/GPS-Device/GPS-Device-01/motion"
✅ GGA: 时间=065101.000, 纬度=3044.307575N, 经度=10357.711434E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065101.000, 纬度=3044.307575N, 经度=10357.711434E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065102.000, 纬度=3044.307572N, 经度=10357.711413E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065102.000, 纬度=3044.307572N, 经度=10357.711413E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065103.000, 纬度=3044.307558N, 经度=10357.711404E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307558N, 经度=10357.711404E, 时间=065103.000, 状态=A
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
level=DEBUG ts=2025-07-10T14:51:02.655279927+08:00 app=device-gps source=gpsdriver.go:95 msg="📖 处理设备 GPS-Device-01 的读取命令"
level=DEBUG ts=2025-07-10T14:51:02.655317906+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: fix_quality"
level=DEBUG ts=2025-07-10T14:51:02.655336886+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: satellites_used"
level=DEBUG ts=2025-07-10T14:51:02.655350355+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: hdop"
level=DEBUG ts=2025-07-10T14:51:02.655358445+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: gps_status"
level=DEBUG ts=2025-07-10T14:51:02.65541051+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: fix_quality reading: {Id:63a156e0-2760-4583-8001-ca0d00579c1d Origin:1752130262655369372 DeviceName:GPS-Device-01 ResourceName:fix_quality ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:GPS定位} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:02.655430723+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: satellites_used reading: {Id:3a172d35-d37f-46b3-a89d-1cd5173523e6 Origin:1752130262655369372 DeviceName:GPS-Device-01 ResourceName:satellites_used ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:12 颗卫星} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:02.65545011+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: hdop reading: {Id:740ec51c-9bec-4acc-9e51-ba2a245be90d Origin:1752130262655369372 DeviceName:GPS-Device-01 ResourceName:hdop ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:0.87 (优秀)} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:02.655465998+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: gps_status reading: {Id:e288fbe7-24b7-4e4e-b3a1-736739502a6d Origin:1752130262655369372 DeviceName:GPS-Device-01 ResourceName:gps_status ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:ACTIVE} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:02.6554748+08:00 app=device-gps source=command.go:72 msg="GET Device Command successfully. Device: GPS-Device-01, Source: status, X-Correlation-ID: 0d5aaf0b-08b5-496e-a5b0-559592eee3b7"
level=DEBUG ts=2025-07-10T14:51:02.655687648+08:00 app=device-gps source=utils.go:72 msg="Event(profileName: GPS-Device, deviceName: GPS-Device-01, sourceName: status, id: a732ea33-c984-479a-afe1-c8e16592401d) published to MessageBus on topic: edgex/events/device/device-gps/GPS-Device/GPS-Device-01/status"
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065103.000, 纬度=3044.307558N, 经度=10357.711404E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065104.000, 纬度=3044.307533N, 经度=10357.711387E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065104.000, 纬度=3044.307533N, 经度=10357.711387E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065105.000, 纬度=3044.307513N, 经度=10357.711400E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
level=DEBUG ts=2025-07-10T14:51:04.711479737+08:00 app=device-gps source=gpsdriver.go:95 msg="📖 处理设备 GPS-Device-01 的读取命令"
level=DEBUG ts=2025-07-10T14:51:04.711524081+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: latitude"
level=DEBUG ts=2025-07-10T14:51:04.711542909+08:00 app=device-gps source=gpsdriver.go:111 msg="latitude: DeviceResource: latitude, String: 30°44'18.5\"N"
level=DEBUG ts=2025-07-10T14:51:04.711556662+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: longitude"
level=DEBUG ts=2025-07-10T14:51:04.711569183+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: altitude"
level=DEBUG ts=2025-07-10T14:51:04.711581456+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: speed"
level=DEBUG ts=2025-07-10T14:51:04.711593341+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: course"
level=DEBUG ts=2025-07-10T14:51:04.71160497+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: utc_time"
level=DEBUG ts=2025-07-10T14:51:04.711616352+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: fix_quality"
level=DEBUG ts=2025-07-10T14:51:04.711628067+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: satellites_used"
level=DEBUG ts=2025-07-10T14:51:04.711639752+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: hdop"
level=DEBUG ts=2025-07-10T14:51:04.711651571+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: gps_status"
level=DEBUG ts=2025-07-10T14:51:04.711709705+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: latitude reading: {Id:a4015282-e8ad-4737-99b9-46a12460234c Origin:1752130264711664755 DeviceName:GPS-Device-01 ResourceName:latitude ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:30°44'18.5\"N} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:04.711734481+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: longitude reading: {Id:f95610f5-125b-4f2f-988a-193501d999d1 Origin:1752130264711664755 DeviceName:GPS-Device-01 ResourceName:longitude ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:103°57'42.7\"E} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:04.711751421+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: altitude reading: {Id:6b51bbd1-aeb9-4b39-be2f-4166d465d1f4 Origin:1752130264711664755 DeviceName:GPS-Device-01 ResourceName:altitude ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:519.1 米} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:04.711766939+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: speed reading: {Id:69c8b88a-827f-4d87-a68c-b05ee7f77331 Origin:1752130264711664755 DeviceName:GPS-Device-01 ResourceName:speed ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:0.00 km/h} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:04.711774632+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: course reading: {Id:995567a4-d358-48dc-912a-a4590354f181 Origin:1752130264711664755 DeviceName:GPS-Device-01 ResourceName:course ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:0.0° (北)} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:04.711793137+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: utc_time reading: {Id:85f2c376-acbf-4c12-8870-8e01e6c01159 Origin:1752130264711664755 DeviceName:GPS-Device-01 ResourceName:utc_time ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:06:51:04.000} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:04.71180746+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: fix_quality reading: {Id:9f000a0d-8753-4f62-a207-715af1aa3853 Origin:1752130264711664755 DeviceName:GPS-Device-01 ResourceName:fix_quality ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:GPS定位} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:04.71183344+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: satellites_used reading: {Id:4e4e82c2-be73-4612-8f04-573f46ca4ca2 Origin:1752130264711664755 DeviceName:GPS-Device-01 ResourceName:satellites_used ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:12 颗卫星} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:04.711848853+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: hdop reading: {Id:c42a0950-6696-4c1a-9b2d-03db055fddeb Origin:1752130264711664755 DeviceName:GPS-Device-01 ResourceName:hdop ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:0.87 (优秀)} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:04.711910554+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: gps_status reading: {Id:2ed37aa1-adaf-481e-8b8a-6988d6752902 Origin:1752130264711664755 DeviceName:GPS-Device-01 ResourceName:gps_status ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:ACTIVE} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:04.711926442+08:00 app=device-gps source=command.go:72 msg="GET Device Command successfully. Device: GPS-Device-01, Source: all_data, X-Correlation-ID: d165649e-dbe4-42cb-af40-8e78e5b1d017"
level=DEBUG ts=2025-07-10T14:51:04.712175163+08:00 app=device-gps source=utils.go:72 msg="Event(profileName: GPS-Device, deviceName: GPS-Device-01, sourceName: all_data, id: 1f36eb92-a8d4-4105-ae8c-bdcae429229e) published to MessageBus on topic: edgex/events/device/device-gps/GPS-Device/GPS-Device-01/all_data"
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065105.000, 纬度=3044.307513N, 经度=10357.711400E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065106.000, 纬度=3044.307499N, 经度=10357.711417E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.52
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065106.000, 纬度=3044.307499N, 经度=10357.711417E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065107.000, 纬度=3044.307488N, 经度=10357.711430E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.51
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
level=DEBUG ts=2025-07-10T14:51:06.372600064+08:00 app=device-gps source=gpsdriver.go:95 msg="📖 处理设备 GPS-Device-01 的读取命令"
level=DEBUG ts=2025-07-10T14:51:06.372638024+08:00 app=device-gps source=gpsdriver.go:104 msg="处理资源: get_output_rates"
level=INFO ts=2025-07-10T14:51:06.372643639+08:00 app=device-gps source=gpsdriver.go:552 msg=开始查询所有NMEA消息输出速率
📤 发送二进制命令: F1D906010200F000F911
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
level=DEBUG ts=2025-07-10T14:51:06.572901285+08:00 app=device-gps source=gpsdriver.go:606 msg=已查询GGA输出速率
📤 发送二进制命令: F1D906010200F001FA12
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065107.000, 纬度=3044.307488N, 经度=10357.711430E, 状态=A
level=DEBUG ts=2025-07-10T14:51:06.773156977+08:00 app=device-gps source=gpsdriver.go:606 msg=已查询GLL输出速率
📤 发送二进制命令: F1D906010200F002FB13
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
📊 输出速率响应: NMEA类型=0xF000, 速率=1
📊 输出速率响应: NMEA类型=0xF001, 速率=5
二进制协议解析错误: 二进制消息太短
📊 输出速率响应: NMEA类型=0xF002, 速率=1
串口读取错误: EOF
level=DEBUG ts=2025-07-10T14:51:06.974133017+08:00 app=device-gps source=gpsdriver.go:606 msg=已查询GSA输出速率
📤 发送二进制命令: F1D906010200F003FC14
📊 输出速率响应: NMEA类型=0xF003, 速率=0
✅ GGA: 时间=065108.000, 纬度=3044.307484N, 经度=10357.711429E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307484N, 经度=10357.711429E, 时间=065108.000, 状态=A
level=DEBUG ts=2025-07-10T14:51:07.174846866+08:00 app=device-gps source=gpsdriver.go:606 msg=已查询GRS输出速率
📤 发送二进制命令: F1D906010200F004FD15
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.51
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
level=DEBUG ts=2025-07-10T14:51:07.375799638+08:00 app=device-gps source=gpsdriver.go:606 msg=已查询GSV输出速率
📤 发送二进制命令: F1D906010200F005FE16
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
level=DEBUG ts=2025-07-10T14:51:07.576506627+08:00 app=device-gps source=gpsdriver.go:606 msg=已查询RMC输出速率
📤 发送二进制命令: F1D906010200F006FF17
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065108.000, 纬度=3044.307484N, 经度=10357.711429E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
level=DEBUG ts=2025-07-10T14:51:07.776934048+08:00 app=device-gps source=gpsdriver.go:606 msg=已查询VTG输出速率
📤 发送二进制命令: F1D906010200F0070018
📊 输出速率响应: NMEA类型=0xF004, 速率=1
📊 输出速率响应: NMEA类型=0xF005, 速率=1
📊 输出速率响应: NMEA类型=0xF006, 速率=1
📊 输出速率响应: NMEA类型=0xF007, 速率=0
串口读取错误: EOF
level=DEBUG ts=2025-07-10T14:51:07.978025725+08:00 app=device-gps source=gpsdriver.go:606 msg=已查询ZDA输出速率
📤 发送二进制命令: F1D906010200F0080119
✅ GGA: 时间=065109.000, 纬度=3044.307485N, 经度=10357.711425E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.51
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
level=DEBUG ts=2025-07-10T14:51:08.178839569+08:00 app=device-gps source=gpsdriver.go:606 msg=已查询GST输出速率
level=INFO ts=2025-07-10T14:51:08.178952887+08:00 app=device-gps source=gpsdriver.go:610 msg="查询完成，输出速率: GGA: 未知, GLL: 未知, GSA: 1Hz, GRS: 禁用, GSV: 未知, RMC: 未知, VTG: 未知, ZDA: 禁用, GST: 未知"
level=DEBUG ts=2025-07-10T14:51:08.179112038+08:00 app=device-gps source=transform.go:100 msg="failed to read ResourceOperation: failed to find ResourceOpertaion with DeviceResource get_output_rates in Profile GPS-Device"
level=DEBUG ts=2025-07-10T14:51:08.179240799+08:00 app=device-gps source=transform.go:123 msg="device: GPS-Device-01 DeviceResource: get_output_rates reading: {Id:bd579575-15b1-497d-922b-00b108120be6 Origin:1752130268179052804 DeviceName:GPS-Device-01 ResourceName:get_output_rates ProfileName:GPS-Device ValueType:String Units: Tags:map[] BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:GGA: 未知, GLL: 未知, GSA: 1Hz, GRS: 禁用, GSV: 未知, RMC: 未知, VTG: 未知, ZDA: 禁用, GST: 未知} ObjectReading:{ObjectValue:<nil>} NullReading:{isNull:false}}"
level=DEBUG ts=2025-07-10T14:51:08.179302557+08:00 app=device-gps source=command.go:72 msg="GET Device Command successfully. Device: GPS-Device-01, Source: get_output_rates, X-Correlation-ID: 082fdcec-5940-4eb9-acc5-55a29a3e1a9d"
level=DEBUG ts=2025-07-10T14:51:08.179612122+08:00 app=device-gps source=utils.go:72 msg="Event(profileName: GPS-Device, deviceName: GPS-Device-01, sourceName: get_output_rates, id: f9d9efd6-f423-428c-a835-3731e07bb6d0) published to MessageBus on topic: edgex/events/device/device-gps/GPS-Device/GPS-Device-01/get_output_rates"
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065109.000, 纬度=3044.307485N, 经度=10357.711425E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
📊 输出速率响应: NMEA类型=0xF008, 速率=0
串口读取错误: EOF
✅ GGA: 时间=065110.000, 纬度=3044.307482N, 经度=10357.711420E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.51
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065110.000, 纬度=3044.307482N, 经度=10357.711420E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065111.000, 纬度=3044.307479N, 经度=10357.711420E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.51
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065111.000, 纬度=3044.307479N, 经度=10357.711420E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065112.000, 纬度=3044.307480N, 经度=10357.711420E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.51
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
level=DEBUG ts=2025-07-10T14:51:11.044513159+08:00 app=device-gps source=gpsdriver.go:151 msg="✍️ 处理设备 GPS-Device-01 的写入命令"
level=DEBUG ts=2025-07-10T14:51:11.04455164+08:00 app=device-gps source=gpsdriver.go:158 msg="处理写入资源: set_output_rate"
level=INFO ts=2025-07-10T14:51:11.044559048+08:00 app=device-gps source=gpsdriver.go:646 msg=设置RMC消息输出速率为5
📤 发送二进制命令: F1D906010300F00505041E
level=DEBUG ts=2025-07-10T14:51:11.044609822+08:00 app=device-gps source=command.go:111 msg="SET Device Command successfully. Device: GPS-Device-01, Source: set_output_rate, X-Correlation-ID: 6cd531be-a9e1-425d-a421-c61206ada946"
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ RMC: 时间=065112.000, 纬度=3044.307480N, 经度=10357.711420E, 状态=A
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
✅ 收到ACK确认: GroupID=0x06, SubID=0x01
level=DEBUG ts=2025-07-10T14:51:11.550235857+08:00 app=device-gps source=reporter.go:188 msg="Publish 0 metrics to the 'edgex/telemetry/device-gps' base topic"
level=DEBUG ts=2025-07-10T14:51:11.550290264+08:00 app=device-gps source=manager.go:123 msg="Reported metrics..."
串口读取错误: EOF
✅ GGA: 时间=065113.000, 纬度=3044.307480N, 经度=10357.711419E, 质量=1, 卫星数=12
✅ GLL: 纬度=3044.307480N, 经度=10357.711419E, 时间=065113.000, 状态=A
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.51
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
✅ GGA: 时间=065114.000, 纬度=3044.307476N, 经度=10357.711419E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.51
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=2, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=3, 可视卫星数=15
✅ GSV: 总语句数=4, 语句号=4, 可视卫星数=15
✅ GSV: 总语句数=3, 语句号=1, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=2, 可视卫星数=09
✅ GSV: 总语句数=3, 语句号=3, 可视卫星数=09
✅ VTG: 航向=000.00, 速度(节)=0.00, 速度(km/h)=0.00
串口读取错误: EOF
串口读取错误: EOF
✅ GGA: 时间=065115.000, 纬度=3044.307483N, 经度=10357.711397E, 质量=1, 卫星数=12
✅ GSA: 模式=A, 定位模式=3, PDOP=1.75, HDOP=0.87, VDOP=1.51
✅ GSV: 总语句数=4, 语句号=1, 可视卫星数=15
^Clevel=INFO ts=2025-07-10T14:51:13.92430205+08:00 app=device-gps source=config.go:811 msg="Watching for 'Writable' configuration changes has stopped"
level=INFO ts=2025-07-10T14:51:13.924302069+08:00 app=device-gps source=config.go:873 msg="Watching for 'Writable' configuration changes has stopped"
level=INFO ts=2025-07-10T14:51:13.92435676+08:00 app=device-gps source=manager.go:110 msg="Exited Metrics Manager Run..."
level=INFO ts=2025-07-10T14:51:13.924354702+08:00 app=device-gps source=validation.go:55 msg="Exiting waiting for MessageBus 'edgex/device-gps/validate/device' topic messages"
level=INFO ts=2025-07-10T14:51:13.924378063+08:00 app=device-gps source=config.go:811 msg="Watching for 'Writable' configuration changes has stopped"
level=INFO ts=2025-07-10T14:51:13.924395725+08:00 app=device-gps source=command.go:60 msg="Exiting waiting for MessageBus 'edgex/device/command/request/device-gps/#' topic messages"
level=INFO ts=2025-07-10T14:51:13.924402459+08:00 app=device-gps source=config.go:873 msg="Watching for 'Writable' configuration changes has stopped"
level=INFO ts=2025-07-10T14:51:13.924403626+08:00 app=device-gps source=config.go:873 msg="Watching for 'Writable' configuration changes has stopped"
level=INFO ts=2025-07-10T14:51:13.924403265+08:00 app=device-gps source=config.go:873 msg="Watching for 'Writable' configuration changes has stopped"
level=INFO ts=2025-07-10T14:51:13.924400956+08:00 app=device-gps source=callback.go:79 msg="Exiting waiting for MessageBus 'edgex/system-events/core-metadata/+/+/device-gps/#' topic messages"
level=INFO ts=2025-07-10T14:51:13.924426295+08:00 app=device-gps source=messaging.go:95 msg="Disconnected from MessageBus"
level=INFO ts=2025-07-10T14:51:13.924456847+08:00 app=device-gps source=httpserver.go:178 msg="Web server stopped"
level=INFO ts=2025-07-10T14:51:13.924550152+08:00 app=device-gps source=httpserver.go:149 msg="Web server shut down"
level=DEBUG ts=2025-07-10T14:51:13.924573524+08:00 app=device-gps source=gpsdriver.go:223 msg="Driver.Stop called: force=false"