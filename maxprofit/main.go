package main

import (
	"fmt"
)

type DiffStock struct {
	BuyIndex int
	BuyVal   int
}

func GetMaxProfitToLast(prices []int, listBuy []DiffStock, indexStart int, indexStop int) int {
	var max int
	for i := indexStart; i <= indexStop; i++ {
		// hitung jika maximal profit
		temp := prices[i] - prices[indexStart]
		if temp < 0 {
			temp = temp * (-1)
		}
		fmt.Println(prices[i], " - ", prices[indexStart], " = ", temp)
		if temp > max {
			max = temp
		}

	}
	fmt.Println("--------------------", max)
	return max
}

func MaxProfit(prices []int, changes int) int {
	max := 0
	// changes := i
	// profit := make(map[int]int)
	var listBuy []DiffStock
	for j, v := range prices {
		if changes != 0 {
			if j+1 <= len(prices)-1 {
				nextVal := prices[j+1]
				// apakah dia lebih murah dari sesudahnya
				if nextVal > v {
					var temp DiffStock
					temp.BuyIndex = j
					temp.BuyVal = v
					listBuy = append(listBuy, temp)
					changes -= 1
				}
			}
		}

	}
	// get profit
	for i, _ := range listBuy {
		if i+1 <= len(listBuy)-1 {
			max = max + GetMaxProfitToLast(prices, listBuy, i, i+1)
		} else {
			max = max + GetMaxProfitToLast(prices, listBuy, i+1, len(prices)-1)
		}
	}
	return max
}

func main() {
	listing := []int{4, 11, 2, 20, 59, 80}
	fmt.Println(MaxProfit(listing, 2))
}
