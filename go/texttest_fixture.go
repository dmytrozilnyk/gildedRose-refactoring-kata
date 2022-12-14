package main

import (
	"fmt"
	"github.com/dmytrozilnyk/gildedRose-refactoring-kata/go/gildedrose"
	"os"
	"strconv"
)

func main() {
	fmt.Println("OMGHAI!")

	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6}, // <-- :O
	}

	days := 2
	var err error
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	for day := 1; day <= days; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		gildedrose.UpdateQuality(items)
		for _, i := range items {
			fmt.Println(fmt.Sprintf("%s: %d days left, quality is %d", i.Name, i.SellIn, i.Quality))
		}
		fmt.Println("")
	}
}
