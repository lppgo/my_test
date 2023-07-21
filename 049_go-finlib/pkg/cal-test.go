package pkg

import (
	"fmt"
)

// 计算最大回撤
func CalculateMaxDrawdown_test() {
	// 假设我们有一组投资组合净值的数据
	portfolioNetValue := []float64{1000, 1200, 1500, 1100, 1400, 900, 1300, 1600}

	// 计算最大回撤
	maxDrawdown := calculate_MaxDrawdown(portfolioNetValue)

	// 打印结果
	fmt.Printf("最大回撤为: %.2f%%\n", maxDrawdown*100)
}

// 计算夏普比率
func SharpeRatio_test() {
	// 假设dailys是每日的profit组成的切片，base是初始金额
	dailys := []float64{100, 120, 130, 110, 140}
	base := 1000.0

	sharpeRatio := calculate_SharpeRatio(dailys, base)
	fmt.Printf("夏普比率: %.2f\n", sharpeRatio)
}
