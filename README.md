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
go run main.go -port=8001 -mode=release -config=configs/
swag -v
swag init
```
檢視文件: `http://127.0.0.1:8000/swagger/index.html`

### 郵件套件(SMTP) & 介面限流套件 & 檔案監聽套件
```
go get -u gopkg.in/gomail.v2
go get -u github.com/juju/ratelimit@1.0.1
go get -u golang.org/x/sys/...
go get -u github.com/fsnotify/fsnotify
```

### 驗證介面功能是否正常
```
curl -X POST http://127.0.0.1:80000/api/v1/tags -F 'name=Go' -F created_by=eddycjy
```

### 安裝鏈路追蹤的 jaeger 套件(用 Docker)(http://localhost:16686)

```
docker run -d --name jaeger \
-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
-p 5775:5775/udp \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 14268:14268 \
-p 9411:9411 \
jaegertracing/all-in-one:1.16
```

### 安裝 jaeger 的相關協力套件

```
go get -u github.com/opentracing/opentracing-go@v1.1.0
go get -u github.com/uber/jaeger-client-go@v2.22.1
go get -u github.com/eddycjy/opentracing-gorm
```

### 其他

共用元件: 錯誤碼標準化、設定檔管理、資料庫連接、記錄檔規則、回應處理