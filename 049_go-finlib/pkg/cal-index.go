package pkg

import (
	"fmt"
	"math"
)

// 计算最大回撤
func calculateMaxDrawdown(data []float64) float64 {
	if len(data) == 0 {
		return 0.0
	}

	maxDrawdown := 0.0
	peak := data[0]

	for i := 1; i < len(data); i++ {
		drawdown := (data[i] - peak) / peak
		fmt.Printf("[%d] drawdown : %.2f\n", i, drawdown)
		if drawdown < maxDrawdown {
			maxDrawdown = drawdown
		}
		if data[i] > peak {
			peak = data[i]
		}
	}

	return math.Abs(maxDrawdown)
}

func CalculateMaxDrawdown_test() {
	// 假设我们有一组投资组合净值的数据
	portfolioNetValue := []float64{1000, 1200, 1500, 1100, 1400, 900, 1300, 1600}

	// 计算最大回撤
	maxDrawdown := calculateMaxDrawdown(portfolioNetValue)

	// 打印结果
	fmt.Printf("最大回撤为: %.2f%%\n", maxDrawdown*100)
}
