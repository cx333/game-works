# game-works
```azure
game-works/
├── go.work
├── go.work.sum
├── services/
│   ├── gateway/            # 接入服务
│   ├── player/             # 玩家状态服务
│   ├── match/              # 匹配服务
│   ├── room/               # 房间服务（对局管理）
│   ├── chat/               # 聊天服务
│   ├── monitor/            # 监控服务
│   ├── game-data/          # 游戏数据服务（充值、礼包等）
│   ├── world/              # 世界服务（广播、活动）
│   ├── auth/               # 鉴权服务
│   └── admin-api/          # 后台运营接口
├── pkg/
│   ├── proto/              # protobuf 协议定义
│   ├── logger/             # 通用日志封装
│   ├── natsx/              # NATS 封装
│   ├── config/             # 配置解析模块
│   └── utils/              # 公共工具函数
└── deploy/
    ├── docker-compose.yml  # 可选：容器部署
    ├── local-dev/          # 本地运行脚本
    └── README.md

```