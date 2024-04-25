package pt_pupuk_indonesia_persero

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Transaction struct {
	Buy  int
	Sell int
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func maxProfit(k int, prices []int) int {
	profit := make([]Transaction, k)
	for i := 0; i < k; i++ {
		profit[i] = Transaction{Buy: -prices[0], Sell: 0}
	}

	for i := 1; i < len(prices); i++ {
		profit[0].Buy = max(profit[0].Buy, -prices[i])
		profit[0].Sell = max(profit[0].Sell, profit[0].Buy+prices[i])

		for j := 1; j < k; j++ {
			profit[j].Buy = max(profit[j].Buy, profit[j-1].Sell-prices[i])
			profit[j].Sell = max(profit[j].Sell, profit[j].Buy+prices[i])
		}
	}

	return profit[k-1].Sell
}

func TestMaxProfit(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4, 10}
	expected := 11
	profit := maxProfit(2, input)

	assert.Equal(t, expected, profit)
}
