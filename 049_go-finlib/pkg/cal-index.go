package pkg

import (
	"fmt"
	"math"
)

// 计算最大回撤
func calculate_MaxDrawdown(data []float64) float64 {
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

// 计算夏普比率
// 夏普比率 = (预期收益率-无风险收益率)/投资组合标准差
// 夏普比率也叫报酬与波动性比率，可能是最常用的投资组合管理度量标准。它采用的方法是，组合中超过无风险利率的那部分收益要用投资组合的标准差来衡量。
// - 夏普率通常是天为基本单位，如果是短线策略也要先按天计算出净值，就是daily return，然后才能算sharp
// - 无风险收益率用十年期国债收益率，大概是 2.85%，一年有252个交易日
// - 研究出一个量化策略，跑出pnl曲线，然后算一下sharp率，就可以评估策略的可行性

// 夏普比率是一个指标，仅在比较不同投资组合或资产时才有意义。它越高，表示单位风险所获得的超额收益越多，因此更好地衡量了投资的回报与风险之间的平衡。

// base 是初始金额，dailys是每日收益的数组
func calculate_SharpeRatio(dailys []float64, base float64) float64 {
	if len(dailys) <= 0 || base <= 0 {
		return 0
	}

	var rtn []float64 // 用来存放每日收益率
	var sum float64 = 0

	for i := 1; i < len(dailys); i++ {
		// 取每日收益率的对数（logarithm）是为了将每日收益率转换为对数收益率，从而在计算夏普比率时能更好地处理收益率的百分比变化。
		t := math.Log((dailys[i] + base) / (dailys[i-1] + base)) // 计算每日收益率的自然对数
		sum += t
		rtn = append(rtn, t)
	}

	if len(rtn) < 2 {
		return 0
	}

	// 计算收益率切片的平均值（均值）
	mean := sum / float64(len(rtn))

	// 计算收益率切片的方差（标准差的平方）用于后续计算.方差是衡量数据分散程度的一种统计量，表示数据值与均值之间的差异。方差越大，数据的分散程度越大
	var accum float64 = 0
	for _, d := range rtn {
		accum += (d - mean) * (d - mean)
	}

	// 投资组合标准差. 标准差是方差的平方根，用于衡量数据的波动性
	stdev := math.Sqrt(accum / (float64(len(rtn)) - 1)) //

	// 无风险利率 使用的是10年期国债利率 2.85% 的每日收益率（除以252个交易日的天数）
	riskFreeRate := 0.0285 / 252
	return math.Sqrt(252) * (mean - riskFreeRate) / stdev
}
