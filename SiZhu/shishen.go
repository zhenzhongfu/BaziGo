package SiZhu

import (
	. "github.com/warrially/BaziGo/Common"
)

// 十神
var SHI_SHEN_LIST = [...][10]int{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, // 甲
	{1, 0, 3, 2, 5, 4, 7, 6, 9, 8}, // 乙
	{8, 9, 0, 1, 2, 3, 4, 5, 6, 7}, // 丙
	{9, 8, 1, 0, 3, 2, 5, 4, 7, 6}, // 丁
	{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}, // 戊
	{7, 6, 9, 8, 1, 0, 3, 2, 5, 4}, // 己
	{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, // 庚
	{5, 4, 7, 6, 9, 8, 1, 0, 3, 2}, // 辛
	{2, 3, 4, 5, 6, 7, 8, 9, 0, 1}, // 壬
	{3, 2, 5, 4, 7, 6, 9, 8, 1, 0}} // 癸

// 从日干和目标干获取十神
// nGan1 日干 nGan2 目标干
func GetShiShenFromGan(nGan1 int, nGan2 int) int {
	if nGan1 < 0 || nGan1 >= 10 || nGan2 < 0 || nGan2 >= 10 {
		return -1
	}
	return SHI_SHEN_LIST[nGan1][nGan2]
}

// 从日干和目标干获取十神,带字符串
// nGan1 日干 nGan2 目标干
func GetShiShenFromGan2(pShiShen *TShiShen, nGan1 int, nGan2 int) TShiShen {
	pShiShen.Value = GetShiShenFromGan(nGan1, nGan2)
	return *pShiShen
}

// 从干获取十神
func GetShiShenFromGan3(pGan *TGan, nGan int) TShiShen {
	pGan.ShiShen.Value = GetShiShenFromGan(nGan, pGan.Value)
	return pGan.ShiShen
}

// 地支藏干表
var DI_ZHI_CANG_GAN_LIST = [12][3]int{
	{9, -1, -1}, // 子水 藏干 癸水。
	{5, 9, 7},   // 丑土 藏干 己土、癸水、辛金。
	{0, 2, 4},   // 寅木 藏干 甲木、丙火、戊土。
	{1, -1, -1}, // 卯木 藏干 乙木。
	{4, 1, 9},   // 辰土 藏干 戊土、乙木、癸水。
	{2, 4, 6},   // 巳火 藏干 丙火、戊土、庚金。
	{3, 5, -1},  // 午火 藏干 丁火、己土。
	{5, 1, 3},   // 未土 藏干 己土、乙木、丁火。
	{6, 4, 8},   // 申金 藏干 庚金、戊土、壬水。
	{7, -1, -1}, // 酉金 藏干 辛金。
	{4, 7, 3},   // 戌土 藏干 戊土、辛金、丁火。
	{8, 0, -1}}  // 亥水 藏干 壬水、甲木。

// 通过地支获取藏干
func GetCangGanFromZhi(nZhi int) [3]TGan {
	var gans [3]TGan
	if nZhi < 0 || nZhi >= 12 {
		for i := 0; i < 3; i++ {
			gans[i].Value = -1
		}
		return gans
	}
	for i := 0; i < 3; i++ {
		gans[i].Value = DI_ZHI_CANG_GAN_LIST[nZhi][i]
	}
	return gans
}

// 通过地支获取藏干
func GetCangGanFromZhi2(pZhi *TZhi) {
	var nZhi = pZhi.Value
	if nZhi < 0 || nZhi >= 12 {
		for i := 0; i < 3; i++ {
			pZhi.CangGan[i].Value = -1
		}
	}
	for i := 0; i < 3; i++ {
		pZhi.CangGan[i].Value = DI_ZHI_CANG_GAN_LIST[nZhi][i]
	}
}

// 计算藏干 nGan 日干的值, 用来计算十神用
func CalcCangGan(nGan int, pCangGan *TGan) {
	// 拿到藏干的值
	var nCangGan = pCangGan.Value
	if nCangGan >= 0 {
		GetShiShenFromGan2(&pCangGan.ShiShen, nGan, nCangGan)
		Get5XingFromGan2(&pCangGan.WuXing, nCangGan)
	} else {
		pCangGan.ShiShen.Value = -1
		pCangGan.WuXing.Value = -1
	}
}

// 计算十神
func CalcShiShen(pSiZhu *TSiZhu) {
	// 取出日干十神作为比较
	var nGan = pSiZhu.DayZhu.Gan.Value
	// 天干
	// 年干
	GetShiShenFromGan3(&pSiZhu.YearZhu.Gan, nGan)
	// 月干
	GetShiShenFromGan3(&pSiZhu.MonthZhu.Gan, nGan)
	// 日干
	GetShiShenFromGan3(&pSiZhu.DayZhu.Gan, nGan)
	// 时干
	GetShiShenFromGan3(&pSiZhu.HourZhu.Gan, nGan)

	// 地支藏干
	GetCangGanFromZhi2(&pSiZhu.YearZhu.Zhi)
	GetCangGanFromZhi2(&pSiZhu.MonthZhu.Zhi)
	GetCangGanFromZhi2(&pSiZhu.DayZhu.Zhi)
	GetCangGanFromZhi2(&pSiZhu.HourZhu.Zhi)

	// 地支藏转十神
	for i := 0; i < 3; i++ {
		// 年
		CalcCangGan(nGan, &pSiZhu.YearZhu.Zhi.CangGan[i])
		CalcCangGan(nGan, &pSiZhu.MonthZhu.Zhi.CangGan[i])
		CalcCangGan(nGan, &pSiZhu.DayZhu.Zhi.CangGan[i])
		CalcCangGan(nGan, &pSiZhu.HourZhu.Zhi.CangGan[i])
	}

	// 转字符串
	pSiZhu.DayZhu.Gan.ShiShen.Value = -1 // 日主
}

// 计算藏干在强度表中位置
func CalcCangGanQiangDuIndex(month, gan int) int {
	for i, v := range DI_ZHI_CANG_GAN_LIST[month] {
		if v == gan {
			return i
		}
	}
	return -1
}
