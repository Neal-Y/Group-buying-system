# 團購系統專案

## 專案簡介
這個專案的初衷是為了幫助業主自動化管理 LINE 群組中的團購訂單。在傳統團購中，群組成員經常使用 "+1" 和 "-1" 的方式來統計購買意願，但這種方法容易出錯且耗時。業主必須反覆確認訂單數量，效率低下。

原本，我計劃使用 LINE 提供的 API，針對 LINE 群組中的記事本留言進行 CRUD（新增、讀取、更新、刪除）操作。然而，LINE 官方認為這樣的操作涉及用戶隱私，因此不提供該 API。

最終，我選擇了另一種解決方案，通過建立一個商品管理後台，讓業主可以在群組的``行銷文章``或``記事本``上附上產品下單連結。這樣，群組成員可以直接透過連結進行購買。這個系統不算是一個傳統的購物車系統，因為每個商品都有自己獨立的下單連結。

## 專案架構
``` plaintext
├── builder                                # 各模組的構建工具函數
│   ├── http.go                            # 處理 HTTP 相關邏輯
│   ├── order.go                           # 訂單模組的構建邏輯
│   ├── product.go                         # 產品模組的構建邏輯
│   └── user.go                            # 使用者模組的構建邏輯
├── config                                 # 配置文件，包含應用程式與數據庫設置
│   ├── config.go                          # 應用程序的配置邏輯
│   └── mysql                              # MySQL 配置與初始化 SQL 文件
├── constant                               # 定義應用中的常量，如 LINE API 與 Token 類型
├── database                               # 數據庫相關，包含遷移文件
│   └── migrations                         # 數據庫遷移文件，用於管理表結構
├── docker-compose.yaml                    # Docker 容器編排文件
├── frontend                               # 前端頁面模板，如管理員和用戶登入頁面
├── go.mod                                 # Go 模組的依賴管理
├── go.sum                                 # Go 模組依賴的校驗和版本信息
├── handler                                # HTTP 請求處理層，負責各模組的業務邏輯
│   ├── admin                              # 管理員相關處理邏輯
│   ├── general                            # 通用邏輯處理
│   ├── order                              # 訂單模組的處理邏輯
│   ├── product                            # 產品模組的處理邏輯
│   └── user                               # 使用者模組的處理邏輯
├── infrastructure                         # 基礎設施層，主要處理數據庫連接
├── main.go                                # 主程序入口點，啟動應用
├── middleware                             # 中介軟體層，負責 JWT 認證等功能
├── model                                  # 數據模型，用於數據庫和 DTO（數據傳輸對象）
│   ├── database                           # 資料庫模型定義
│   └── datatransfer                       # 資料傳輸對象定義 (DTO)
├── repository                             # 數據庫訪問層，處理數據的 CRUD 操作
├── route                                  # Gin 框架的路由定義
├── service                                # 服務層，處理具體業務邏輯
└── util                                   # 通用工具類，包含時間處理、JWT、郵件發送等功能

```

## 功能特色
- **商品後台管理**：業主可以通過後台上傳、管理商品，並生成獨立的下單連結。
- **下單頁面**：每個商品都有一個獨立的下單頁面，包含商品圖片、描述、價格等信息，使用者可以選擇購買數量並留言備註。
- **訂單管理**：後台系統會自動匯總使用者的訂單資訊，供業主查看和管理。
- **歷史訂單查詢**：使用者可以查看歷史訂單，並根據需求進行操作，如刪除或查詢過往訂單。

## 專案啟發
專案的靈感來自於團購群組的手動統計方式，這種方式在實際運作中非常麻煩且容易出現錯誤。起初，我計劃利用 LINE 的記事本留言功能來實現團購訂單的自動化統計，但由於涉及用戶隱私問題，LINE 不提供這樣的 API。最終，我選擇了通過打造一個商品管理後台，來取代手動統計的方式，並實現自動化訂單處理。

## 技術棧
- **後端框架**：Golang + Gin
- **資料庫**：MySQL
- **部署平台**：AWS EC2 + RDS
- **其他工具**：ngrok 用於本地開發測試

## 未來改進
- 增加對不同支付方式的支持
- 提供更詳細的訂單分析與報表功能
- 優化前端頁面，提升用戶體驗