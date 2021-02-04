# gin-go-start

## ent generate
- go generate ./app/ent

## run
- go run main.go

## hot load
- go get github.com/pilu/fresh
- fresh

## 目录结构
- controller 控制器
- ent 数据库实体 entgo.io
- middleware 路由中间件，通常用来实现认证功能
- repository 数据访问层 Dao
- router Gin路由器
- service 业务层（业务除使用 service 实现外，可以直接调用 repository ）
- util 系统工具和配置
- vo 价值数据 Value Object（详细用法见 vo/user.go ）