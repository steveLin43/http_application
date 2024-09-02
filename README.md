# http_application

使用 gin 架構進行開發

### 初始化 & 套件管理 & 檔案管理 & log & validator & DB & JWT

```
go mod init 專案名稱
go get -u github.com/gin-gonic/gin@v1.6.3
go get -u github.com/spf13/viper@v1.4.0
go get -u gopkg.in/natefinch/lumberjack.v2
go get -u github.com/go-playground/validator/v10
go get -u github.com/jinzhu/gorm
go get -u github.com/dgrijalva/jwt-go@v3.2.0
```

### Swagger 套件

```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger@v1.2.0
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template
```

### 驗證服務是否正常 & Swagger 是否安裝成功
如果 swag 安裝失敗，可以執行 `go install github.com/swaggo/swag/cmd/swag@latest`
```
go run main.go
swag -v
swag init
```
檢視文件: `http://127.0.0.1:8000/swagger/index.html`

### 郵件套件(SMTP) & 介面限流套件
```
go get -u gopkg.in/gomail.v2
go get -u github.com/juju/ratelimit@1.0.1
```

### 驗證介面功能是否正常
```
curl -X POST http://127.0.0.1:80000/api/v1/tags -F 'name=Go' -F created_by=eddycjy
```

### 其他

共用元件: 錯誤碼標準化、設定檔管理、資料庫連接、記錄檔規則、回應處理