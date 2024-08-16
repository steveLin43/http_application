# http_application

使用 gin 架構進行開發

### 初始化模塊 & 建立相關套件管理 & 檔案管理套件 & log 元件

```
go mod init 專案名稱
go get -u github.com/gin-gonic/gin@v1.6.3
go get -u github.com/spf13/viper@v1.4.0
go get -u gopkg.in/natefinch/lumberjack.v2
```

### 驗證服務是否正常
```
go run main.go
```

### 其他

共用元件: 錯誤碼標準化、設定檔管理、資料庫連接、記錄檔規則、回應處理