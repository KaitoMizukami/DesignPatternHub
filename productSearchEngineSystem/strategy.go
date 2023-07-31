package main

type Strategy interface {
	search([]Product, string, int, int) []Product
}

type NameCategoryStrategy struct{}

func (n *NameCategoryStrategy) search(products []Product, keyword string, minPrice, maxPrice int) []Product {
	var results []Product
	for _, product := range products {
		if product.Category == keyword || product.Name == keyword && product.Price >= minPrice && product.Price <= maxPrice {
			results = append(results, product)
		}
	}
	return results
}

type PriceStrategy struct{}

func (p *PriceStrategy) search(products []Product, keyword string, minPrice, maxPrice int) []Product {
	var results []Product
	for _, product := range products {
		if product.Price >= minPrice && product.Price <= maxPrice {
			results = append(results, product)
		}
	}
	return results
}
