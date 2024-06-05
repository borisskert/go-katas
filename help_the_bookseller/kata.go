package kata

import (
	"strconv"
	"strings"
)

// StockList / Help the bookseller !
// https://www.codewars.com/kata/54dc6f5a224c26032800005c/train/go
func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 {
		return ""
	}

	stock := CreateBookStock(listArt)
	categories := CreateCategories(listCat)

	return stock.Find(categories).String()
}

type BookStock struct {
	StockItems map[rune]uint64
}

func CreateBookStock(listArt []string) BookStock {
	items := make(map[rune]uint64)

	for _, art := range listArt {
		item := NewStockItem(art)
		items[item.BookCode] += item.Quantity
	}

	return BookStock{StockItems: items}
}

type StockItem struct {
	BookCode rune
	Quantity uint64
}

func NewStockItem(art string) StockItem {
	words := strings.Split(art, " ")
	categories := words[0]
	quantity, _ := strconv.ParseUint(words[1], 10, 64)

	item := StockItem{}

	item.BookCode = rune(categories[0])
	item.Quantity = quantity

	return item
}

type Categories struct {
	Items []rune
}

func CreateCategories(listCat []string) Categories {
	items := make([]rune, 0)

	for _, cat := range listCat {
		items = append(items, rune(cat[0]))
	}

	return Categories{Items: items}
}

type SearchResult struct {
	Items []StockItem
}

func (ctx BookStock) Find(categories Categories) SearchResult {
	items := make([]StockItem, len(categories.Items))

	for index, cat := range categories.Items {
		if quantity, ok := ctx.StockItems[cat]; ok {
			items[index] = StockItem{BookCode: cat, Quantity: quantity}
		} else {
			items[index] = StockItem{BookCode: cat, Quantity: 0}
		}
	}

	return SearchResult{Items: items}
}

func (ctx SearchResult) String() string {
	var stocksAsStrings = make([]string, len(ctx.Items))

	for key, value := range ctx.Items {
		category := string(value.BookCode)
		amount := strconv.FormatUint(value.Quantity, 10)

		stocksAsStrings[key] = "(" + category + " : " + amount + ")"
	}

	return strings.Join(stocksAsStrings, " - ")
}
