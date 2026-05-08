# go-shopping

Go + Next.js で構築したフルスタックのショッピングアプリケーションです。

## 技術スタック

### バックエンド
- **言語**: Go 1.18
- **フレームワーク**: [Gin](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **データベース**: SQLite
- **セッション管理**: [gorilla/sessions](https://github.com/gorilla/sessions)（Cookie ベース）

### フロントエンド
- **フレームワーク**: [Next.js 14](https://nextjs.org/) (App Router)
- **言語**: TypeScript
- **スタイリング**: Bootstrap（CDN）

## 機能

- 商品一覧表示
- キーワード・価格帯による商品検索・フィルタリング（低価格 / 中価格 / 高価格）
- カテゴリ別商品表示
- カート管理（追加 / 削除 / 合計金額計算）
- 注文（チェックアウト）と注文確認

## プロジェクト構成

```
go-shopping/
├── backend/                  # Go バックエンド
│   ├── main.go               # エントリポイント・ルーティング設定
│   ├── go.mod
│   ├── config/
│   │   └── database.go       # DB 初期化・マイグレーション
│   ├── handlers/
│   │   ├── product_handler.go
│   │   ├── cart_handler.go
│   │   └── order_handler.go
│   ├── models/
│   │   ├── product.go
│   │   ├── customer.go
│   │   ├── order.go
│   │   └── order_detail.go
│   ├── services/
│   │   ├── product_service.go
│   │   └── order_service.go
│   └── seed/
│       └── seed.go           # 初期データ投入
└── frontend/                 # Next.js フロントエンド
    ├── app/
    │   ├── page.tsx          # 商品一覧ページ
    │   ├── cart/page.tsx     # カートページ
    │   ├── checkout/page.tsx # チェックアウトページ
    │   └── order/confirmation/page.tsx  # 注文確認ページ
    └── lib/
        └── api.ts            # バックエンド API クライアント
```

## API エンドポイント

| メソッド | パス | 説明 |
|---------|------|------|
| GET | `/api/products` | 全商品取得 |
| GET | `/api/products/search?keyword=&minPrice=&maxPrice=` | 商品検索 |
| GET | `/api/products/category/:category` | カテゴリ別商品取得 |
| GET | `/api/products/filter/low` | 低価格帯（〜¥10,000） |
| GET | `/api/products/filter/mid` | 中価格帯（¥10,001〜¥50,000） |
| GET | `/api/products/filter/high` | 高価格帯（¥50,001〜） |
| GET | `/api/cart` | カート取得 |
| POST | `/api/cart/add/:id` | カートに商品追加 |
| DELETE | `/api/cart/remove/:id` | カートから商品削除 |
| POST | `/api/orders` | 注文確定 |

## セットアップ・起動方法

### バックエンド

```bash
cd backend
go mod download
go run main.go
```

サーバーは `http://localhost:8080` で起動します。  
初回起動時に SQLite データベース（`shopping.db`）が自動生成され、サンプル商品データが投入されます。

### フロントエンド

```bash
cd frontend
npm install
npm run dev
```

フロントエンドは `http://localhost:3000` で起動します。

## 環境変数

| 変数名 | デフォルト値 | 説明 |
|--------|------------|------|
| `PORT` | `8080` | バックエンドのポート番号 |
| `SESSION_SECRET` | `dev-secret-key-change-in-production` | セッション署名用シークレットキー（本番環境では必ず変更） |

## データモデル

```
Product      商品（ID, Name, Description, Price, Category）
Customer     顧客（ID, Name, Email, Address, PhoneNumber）
Order        注文（ID, CustomerID, OrderDate, TotalAmount）
OrderDetail  注文明細（ID, OrderID, ProductID, Quantity, Price）
```
