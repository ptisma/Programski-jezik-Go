package main

import (
	/* "errors" */
	"errors"
	"fmt"
)

var errNoData = errors.New("no data in list")

type ShopItem struct {
	Name     string
	Price    int
	Callory  int64
	Quantity int
}

func main() {
	shopList := []ShopItem{

		{"Hosomaki Ginger", 54, 250, 1},
		{"Kivi", 10, 150, 6},
		{"Mlijecna cokolada", 8, 600, 3},
		{"Banana", 10, 150, 6},
		{"Lubenica", 10, 50, 1},
	}
	tCal := totalCallory(shopList)
	fmt.Println(tCal) //2950
	best, err := mostHealthy(shopList)
	fmt.Println(best)
	fmt.Printf(err.Error())

}

func totalCallory(shoppingList []ShopItem) (total int) {
	//body

	for _, element := range shoppingList {
		//fmt.Println(element.Callory)
		total += int(element.Callory)
	}
	return
}

func mostHealthy(shoppingList []ShopItem) (items []ShopItem, err error) {
	//tijelofunkcije
	var results []ShopItem

	for _, element := range shoppingList {
		if len(results) == 0 {
			results = append(results, element)
		} else if element.Callory < results[len(results)-1].Callory {
			results = nil
			results = append(results, element)

		} else if element.Callory == results[len(results)-1].Callory {
			results = append(results, element)

		}
	}

	return results, errNoData
}
