package SiZhu

import (
	. "github.com/warrially/BaziGo/Common"
	"github.com/warrially/BaziGo/Days"
)

// TODO 统计五行
func StatWuXing(pSiZhu *TSiZhu) TWuXingStat {
	var stat TWuXingStat
	stat.Result[pSiZhu.YearZhu.Gan.WuXing.Value] += 1
	stat.Result[pSiZhu.YearZhu.Zhi.WuXing.Value] += 1
	stat.Result[pSiZhu.MonthZhu.Gan.WuXing.Value] += 1
	stat.Result[pSiZhu.MonthZhu.Zhi.WuXing.Value] += 1
	stat.Result[pSiZhu.DayZhu.Gan.WuXing.Value] += 1
	stat.Result[pSiZhu.DayZhu.Zhi.WuXing.Value] += 1
	stat.Result[pSiZhu.HourZhu.Gan.WuXing.Value] += 1
	stat.Result[pSiZhu.HourZhu.Zhi.WuXing.Value] += 1

	// 地支藏干
	stat.CangGanResult[pSiZhu.YearZhu.Gan.WuXing.Value] += 1
	stat.CangGanResult[pSiZhu.MonthZhu.Gan.WuXing.Value] += 1
	stat.CangGanResult[pSiZhu.DayZhu.Gan.WuXing.Value] += 1
	stat.CangGanResult[pSiZhu.HourZhu.Gan.WuXing.Value] += 1
	for i := 0; i < 3; i++ {
		// 年
		var nCangGan = pSiZhu.YearZhu.Zhi.CangGan[i].Value
		if nCangGan >= 0 {
			stat.CangGanResult[pSiZhu.YearZhu.Zhi.CangGan[i].WuXing.Value] += 1
		}
		// 月
		nCangGan = pSiZhu.MonthZhu.Zhi.CangGan[i].Value
		if nCangGan >= 0 {
			stat.CangGanResult[pSiZhu.MonthZhu.Zhi.CangGan[i].WuXing.Value] += 1
		}
		// 日
		nCangGan = pSiZhu.DayZhu.Zhi.CangGan[i].Value
		if nCangGan >= 0 {
			stat.CangGanResult[pSiZhu.DayZhu.Zhi.CangGan[i].WuXing.Value] += 1
		}
		// 时
		nCangGan = pSiZhu.HourZhu.Zhi.CangGan[i].Value
		if nCangGan >= 0 {
			stat.CangGanResult[pSiZhu.HourZhu.Zhi.CangGan[i].WuXing.Value] += 1
		}
	}
	return stat
}

// 补充五行
func CalcWuXing(pZhu *TZhu) TZhu {
	// 五行
	Get5XingFromGan2(&pZhu.Gan.WuXing, pZhu.Gan.Value)
	Get5XingFromZhi2(&pZhu.Zhi.WuXing, pZhu.Zhi.Value)
	return *pZhu
}

// 从八字年获得年柱
func GetZhuFromYear(nYear int) TZhu {
	var zhu TZhu
	// 获得八字年的干支，0-59 对应 甲子到癸亥
	GetGanZhiFromYear2(&zhu.GanZhi, nYear)
	// 获得八字年的干0-9 对应 甲到癸
	// 获得八字年的支0-11 对应 子到亥
	ExtractGanZhi2(&zhu.GanZhi, &zhu.Gan, &zhu.Zhi)

	return CalcWuXing(&zhu)
}

// 从八字月 和 年干 获得月柱
func GetZhuFromMonth(nMonth int, nGan int) TZhu {
	nMonth = nMonth % 12
	nGan = nGan % 10
	if nMonth <= 0 {
		nMonth += 12
	}

	if nGan < 0 {
		nGan += 10
	}

	var zhu TZhu
	// 根据口诀从本年干数计算本年首月的干数
	switch nGan {
	case 0, 5:
		// 甲己 丙佐首
		nGan = 2
	case 1, 6:
		// 乙庚 戊为头
		nGan = 4
	case 2, 7:
		// 丙辛 寻庚起
		nGan = 6
	case 3, 8:
		// 丁壬 壬位流
		nGan = 8
	case 4, 9:
		// 戊癸 甲好求
		nGan = 0
	}

	// 计算本月干数
	nGan += ((nMonth - 1) % 10)

	zhu.Gan.Value = nGan % 10
	zhu.Zhi.Value = (nMonth - 1 + 2) % 12

	CombineGanZhi2(&zhu.GanZhi, &zhu.Gan, &zhu.Zhi)

	return CalcWuXing(&zhu)
}

// 从公历天 获得日柱
func GetZhuFromDay(nYear int, nMonth int, nDay int) TZhu {
	var zhu TZhu
	// 通过总天数 计算当前天的一个值
	zhu.GanZhi.Value = Days.GetGanZhiFromDay(Days.GetAllDays(nYear, nMonth, nDay))
	// 获得八字日的干0-9 对应 甲到癸
	// 获得八字日的支0-11 对应 子到亥
	ExtractGanZhi2(&zhu.GanZhi, &zhu.Gan, &zhu.Zhi)

	return CalcWuXing(&zhu)
}

// 从公历小时,  获得日柱天干获取时柱
func GetZhuFromHour(nHour int, nGan int) TZhu {
	var zhu TZhu

	zhu.Gan.Value, zhu.Zhi.Value = Days.GetGanZhiFromHour(nHour, nGan)
	CombineGanZhi2(&zhu.GanZhi, &zhu.Gan, &zhu.Zhi)

	return CalcWuXing(&zhu)
}

// 计算十二长生地势
func CalcDiShi(pSiZhu *TSiZhu) TDiShi {
	var disi TDiShi
	disi.YearChangSheng.Value = SHI_ER_CHANG_SHENG_LIST[pSiZhu.DayZhu.Gan.Value][pSiZhu.YearZhu.Zhi.Value]
	disi.MonthChangSheng.Value = SHI_ER_CHANG_SHENG_LIST[pSiZhu.DayZhu.Gan.Value][pSiZhu.MonthZhu.Zhi.Value]
	disi.DayChangSheng.Value = SHI_ER_CHANG_SHENG_LIST[pSiZhu.DayZhu.Gan.Value][pSiZhu.DayZhu.Zhi.Value]
	disi.HourChangSheng.Value = SHI_ER_CHANG_SHENG_LIST[pSiZhu.DayZhu.Gan.Value][pSiZhu.HourZhu.Zhi.Value]
	return disi
}
