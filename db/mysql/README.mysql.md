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
