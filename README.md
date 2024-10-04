# http_application

使用 gin 架構進行開發
參考資料：《用 Go 語言完成 6 個大型專案》第二章節 + https://github.com/go-programming-tour-book/blog-service

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

### 編譯相關指令一覽表

子指令
| 指令 | 介紹 |
| -------- | -------- |
| go run | 編譯並執行，只編譯 main 套件下檔案 |
| go build | 編譯(包含依賴項) |
| go install | 編譯並安裝原始檔案，並將項目移到 bin 目錄底下 |

通用參數
| 參數 | 介紹 |
| -------- | -------- |
| -x | 列印編譯過程中的所有執行指令，並執行產生的二進制檔案 |
| -n | 列印編譯過程中的所有執行指令，不執行產生的二進制檔案 |
| -a | 強制重新更新相關依賴 |
| -o | 指定產生的二進制檔案名稱 |
| -p | 指定編譯過程中可平行處理執行程式的數量，預設值為可用的 CPU 數量 |
| -work | 列印臨時工作目錄的完整路徑，在退出時不刪除該目錄 |

編譯時可以使用 upx 套件將可執行檔進行壓縮
另外可以使用 Idflags 套件去設定編譯資訊
```
go build -Idflags "-X main.appName=我的部落格"
```

### 終端快速鍵
| 命令 | 訊號 | 含義 |
| -------- | -------- | -------- |
| ctrl+c | SIGINT | 希望處理程序中斷，處理程序結束 |
| ctrl+z | SIGINT | 工作中斷，處理程序暫停 |
| ctrl+\ | SIGINT | 處理程序結束和 dump core |
