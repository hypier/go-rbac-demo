# 简单用户管理系统

## 功能说明

1. 用户注册、登陆/退出接口
2. 按不同角色权限验证

## 目录结构 
```html
.
├─db    数据库文件
└─src
    ├─controller    表现层
    ├─domain    业务逻辑层
    │  └─entity     实体对象
    ├─repository    持久化层
    └─util      工具类

```

**注：有部分代码暂时没有用，未实现**

## 系统配置
- 数据库配置 db.go
- 服务器端口 main.go 
- 访问地址 **http://localhost:8888/reg**

## 路由地址

### 1. 用户注册 reg
```
curl --location --request POST 'localhost:8888/reg' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--header 'Cookie: admin_name=hypier' \
--data-urlencode 'name=admin' \
--data-urlencode 'password=admin' \
--data-urlencode 'role=Admin'
```

### 2. 用户登陆 login
```
curl --location --request POST 'localhost:8888/login' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--header 'Cookie: admin_name=admin' \
--data-urlencode 'name=admin' \
--data-urlencode 'password=admin'
```

### 3. 管理员查询
```
curl --location --request GET 'localhost:8888/admin'
```

### 4. 普通用户查询
```
curl --location --request GET 'localhost:8888/user'
```

### 5. 用户退出
```
curl --location --request GET 'localhost:8888/logout'
```