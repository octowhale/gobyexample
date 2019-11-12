# MySQL


```yaml
# docker-compose
version: '3.1'

services:
  db:
    image: mysql:5.7
    # command: --default-authentication-plugin=mysql_native_password
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: V69Gz4wsysY
    ports:
      - 3306:3306
    network_mode: default
```


## 下载数据库driver
```go
go get -u github.com/go-sql-driver/mysql
```

## 简单用法

+ init_db: [main.go](main.go)
+ query: query.go
+ insert/delete/update: insert.go


## MySQL 预处理
1. 先编译sql, 在传入变量。 类似 `regex.Compile`。
