package seed

import (
	"github.com/example/go-shopping/models"
	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) {
	var count int64
	db.Model(&models.Product{}).Count(&count)
	if count > 0 {
		return
	}

	products := []models.Product{
		{Name: "スマートフォン", Description: "最新のスマートフォン", Price: 89800, Category: "電化製品"},
		{Name: "ノートパソコン", Description: "ビジネス用ノートPC", Price: 128000, Category: "電化製品"},
		{Name: "コーヒーメーカー", Description: "全自動コーヒーメーカー", Price: 19800, Category: "家電"},
		{Name: "スニーカー", Description: "ランニングシューズ", Price: 8900, Category: "衣類"},
		{Name: "バックパック", Description: "大容量リュックサック", Price: 6800, Category: "バッグ"},
		{Name: "ワイヤレスイヤホン", Description: "ノイズキャンセリング機能付き", Price: 29800, Category: "電化製品"},
		{Name: "デジタルカメラ", Description: "4K動画対応", Price: 78000, Category: "電化製品"},
		{Name: "スマートウォッチ", Description: "心拍数モニター付き", Price: 32800, Category: "電化製品"},
		{Name: "キーボード", Description: "メカニカルキーボード", Price: 15800, Category: "PC周辺機器"},
		{Name: "マウス", Description: "ゲーミングマウス", Price: 8800, Category: "PC周辺機器"},
	}

	db.Create(&products)
}
