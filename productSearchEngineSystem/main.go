/*
問題: オンラインストアの商品検索システムの実装

あなたは、オンラインストアの商品検索システムをGo言語で実装する必要があります。
このシステムは、顧客が特定のキーワードで商品を検索できるようにするものです。商品データは既にデータベースに保存されており、以下のような情報を持っています：

商品ID
商品名
商品カテゴリ
価格
在庫数

このシステムの要件は以下の通りです：

①顧客がキーワードを入力すると、商品名と商品カテゴリのどちらでも検索できるようにする。
②価格帯を指定して検索できるようにする。たとえば、1000円未満の商品や10000円以上の商品を検索できる。
システムの拡張性を考慮し、将来的に新しい検索オプションを追加しやすいようにする。
上記の要件を満たすために、デザインパターンを使用してシステムを実装してください。以下は参考としていくつかのデザインパターンの例ですが、これに限らず他のパターンを使用することも可能です：
*/

// 色々な検索オプションが追加される可能性があるからStrategyパターンを使う
// Strategyインターフェイスを実装すれば容易に機能追加できる

package main

import "fmt"

type Product struct {
	ID       int
	Name     string
	Category string
	Price    int
	Stock    int
}

func executeSearch(strategy Strategy, products []Product, keyword string, minPrice, maxPrice int) []Product {
	return strategy.search(products, keyword, minPrice, maxPrice)
}

func main() {
	products := []Product{
		{ID: 1, Name: "ノートパソコン", Category: "電化製品", Price: 80000, Stock: 10},
		{ID: 2, Name: "スマートフォン", Category: "電化製品", Price: 60000, Stock: 5},
		{ID: 3, Name: "スニーカー", Category: "ファッション", Price: 8000, Stock: 20},
	}

	nameCategoryStrategy := &NameCategoryStrategy{}
	results := executeSearch(nameCategoryStrategy, products, "電化製品", 0, 100000)
	fmt.Println("名前とカテゴリで検索結果:")
	for _, p := range results {
		fmt.Printf("ID: %d, Name: %s, Category: %s, Price: %d, Stock: %d\n", p.ID, p.Name, p.Category, p.Price, p.Stock)
	}

	priceStrategy := &PriceStrategy{}
	results = executeSearch(priceStrategy, products, "", 5000, 70000)
	fmt.Println("\n価格で検索結果:")
	for _, p := range results {
		fmt.Printf("ID: %d, Name: %s, Category: %s, Price: %d, Stock: %d\n", p.ID, p.Name, p.Category, p.Price, p.Stock)
	}
}
