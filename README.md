# gin_gorm_restful
使用Gin与GORM构建的RESTfulAPI，连接MySQL

## API设计

User路由 /users

具体实现在`/web/user-controller/user_controller.go`

|method|path|description|
|-|-|-|
| GET | /:id | 根据ID查询用户信息 |
| GET | / | 查询用户列表 |
| POST | / | 新增用户 |
| PUT  | /:id | 更新用户信息 |
| DELETE | /:id | 删除用户 |

## 目录结构

```
├── model                      Model层，数据库连接
|   ├── user-dao               用户Model相关
|   |   └── user_dao.go        用户struct及表操作
|   └── db.go                  GORM的数据库连接及相关配置
├── util                       工具
|   └── string_util.go         String类相关工具
├── web                        Web相关
|   ├── result                 Result类相关
|   |   ├── result_func.go     获取Result的函数
|   |   └── result.go          Result的Struct以及相关函数
|   ├── user-controller        用户Web相关
|   |   └── user_controller.go 用户相关操作路由控制器
|   └── web.go                 Gin的初始化与设置
└── main.go                    启动类
```