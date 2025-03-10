# Service State Exporter

**目录结构**
```
├── collector
│   ├── collector.go      # 监控指标采集器
│   └── serviceState.go   # 执行命令获取监控指标数据
├── go.mod                # 项目依赖管理文件
├── go.sum
└── main.go               # 注册指标并暴露
```


