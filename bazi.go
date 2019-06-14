package BaziGo

import (
	"log"
	"math"

	. "github.com/warrially/BaziGo/Common"
	"github.com/warrially/BaziGo/DaYun"
	"github.com/warrially/BaziGo/Days"
	"github.com/warrially/BaziGo/JieQi"
	"github.com/warrially/BaziGo/LiChun"
	"github.com/warrially/BaziGo/Lunar"
	"github.com/warrially/BaziGo/SiZhu"
)

// 八字
type TBazi struct {
	SolarDate   TDate       // 新历日期
	LunarDate   TDate       // 农历日期
	BaziDate    TDate       // 八字日期
	PreviousJie TDate       // 上一个节(气)
	NextJie     TDate       // 下一个节(气)
	SiZhu       TSiZhu      // 四柱
	XiYong      TXiYong     // 喜用神
	DaYun       TDaYun      // 大运
	HeHuaChong  THeHuaChong // 合化冲
	WuXing      TWuXingStat // 五行统计
	DiShi       TDiShi      // 十二长生地势
}

// 计算
func calc(bazi *TBazi, nSex int) {
	// 通过立春获取当年的年份
	bazi.BaziDate.Year = LiChun.GetLiChun2(bazi.SolarDate)
	// 通过节气获取当前后的两个节
	bazi.PreviousJie, bazi.NextJie = JieQi.GetJieQi(bazi.SolarDate)
	// 八字所在的节气是上一个的节气
	bazi.BaziDate.JieQi = bazi.PreviousJie.JieQi
	// 节气0 是立春 是1月
	bazi.BaziDate.Month = bazi.BaziDate.JieQi/2 + 1

	// 通过八字年来获取年柱
	bazi.SiZhu.YearZhu = SiZhu.GetZhuFromYear(bazi.BaziDate.Year)
	// 通过年干支和八字月
	bazi.SiZhu.MonthZhu = SiZhu.GetZhuFromMonth(bazi.BaziDate.Month, bazi.SiZhu.YearZhu.Gan.Value)
	// 通过公历 年月日计算日柱
	bazi.SiZhu.DayZhu = SiZhu.GetZhuFromDay(bazi.SolarDate.Year, bazi.SolarDate.Month, bazi.SolarDate.Day)
	//
	bazi.SiZhu.HourZhu = SiZhu.GetZhuFromHour(bazi.SolarDate.Hour, bazi.SiZhu.DayZhu.Gan.Value)

	// 计算十神
	SiZhu.CalcShiShen(&bazi.SiZhu)
	// 计算纳音
	SiZhu.CalcNaYin(&bazi.SiZhu)

	// 检查合化冲
	SiZhu.CheckHeHuaChong(&bazi.SiZhu, &bazi.HeHuaChong)

	// 计算大运
	bazi.DaYun = DaYun.CalcDaYun(&bazi.SiZhu, nSex)

	// 计算起运时间
	DaYun.CalcQiYun(&bazi.DaYun, bazi.PreviousJie, bazi.NextJie, bazi.SolarDate)

	// 计算喜用神
	bazi.XiYong = SiZhu.CalcXiYong(&bazi.SiZhu)

	// 统计五行
	bazi.WuXing = SiZhu.StatWuXing(&bazi.SiZhu)

	// 计算十二长生
	bazi.DiShi = SiZhu.CalcDiShi(&bazi.SiZhu)
}

// 从新历获取八字(年, 月, 日, 时, 分, 秒, 性别男1,女0)
func GetBazi(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex int) TBazi {
	var bazi TBazi

	if !Days.GetDateIsValid(nYear, nMonth, nDay) {
		log.Println("无效的日期", nYear, nMonth, nDay)
		return bazi
	}

	// 新历年
	bazi.SolarDate.Year = nYear
	bazi.SolarDate.Month = nMonth
	bazi.SolarDate.Day = nDay
	bazi.SolarDate.Hour = nHour
	bazi.SolarDate.Minute = nMinute
	bazi.SolarDate.Second = nSecond

	// 转农历
	var nTimeStamp = Days.Get64TimeStamp(nYear, nMonth, nDay, nHour, nMinute, nSecond)
	bazi.LunarDate = Lunar.GetDateFrom64TimeStamp(nTimeStamp)

	// 进行计算
	calc(&bazi, nSex)

	return bazi
}

// 从农历获取八字
func GetBaziFromLunar(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex int, isLeap bool) TBazi {
	nYear, nMonth = Lunar.ChangeLeap(nYear, nMonth, isLeap)

	var bazi TBazi

	if !Lunar.GetDateIsValid(nYear, nMonth, nDay) {
		log.Println("无效的日期", nYear, nMonth, nDay)
		return bazi
	}

	// 农历年
	bazi.LunarDate.Year = nYear
	bazi.LunarDate.Month = nMonth
	bazi.LunarDate.Day = nDay
	bazi.LunarDate.Hour = nHour
	bazi.LunarDate.Minute = nMinute
	bazi.LunarDate.Second = nSecond

	// 转新历
	var nTimeStamp = Lunar.Get64TimeStamp(nYear, nMonth, nDay, nHour, nMinute, nSecond)
	bazi.LunarDate = Days.GetDateFrom64TimeStamp(nTimeStamp)

	// 进行计算
	calc(&bazi, nSex)

	return bazi

}

func (self *TBazi) Print() {
	PrintBazi(self)
}

// 按照特殊格式化输出(未完成)
func PrintBazi(bazi *TBazi) {
	log.Println("======================================================================")
	log.Println("出生日期新历： ", bazi.SolarDate.Year, "年",
		bazi.SolarDate.Month, "月",
		bazi.SolarDate.Day, "日  ",
		bazi.SolarDate.Hour, ":",
		bazi.SolarDate.Minute, ":",
		bazi.SolarDate.Second,
	)
	log.Println("基本八字：",
		bazi.SiZhu.YearZhu.GanZhi.ToString(),
		bazi.SiZhu.MonthZhu.GanZhi.ToString(),
		bazi.SiZhu.DayZhu.GanZhi.ToString(),
		bazi.SiZhu.HourZhu.GanZhi.ToString())

	log.Println("命盘解析：")
	log.Println(
		bazi.SiZhu.YearZhu.Gan.ToString(), "(",
		bazi.SiZhu.YearZhu.Gan.WuXing.ToString(), ")[",
		bazi.SiZhu.YearZhu.Gan.ShiShen.ToString(), "]\t",
		bazi.SiZhu.MonthZhu.Gan.ToString(), "(",
		bazi.SiZhu.MonthZhu.Gan.WuXing.ToString(), ")[",
		bazi.SiZhu.MonthZhu.Gan.ShiShen.ToString(), "]\t",
		bazi.SiZhu.DayZhu.Gan.ToString(), "(",
		bazi.SiZhu.DayZhu.Gan.WuXing.ToString(), ")[日主]\t",
		bazi.SiZhu.HourZhu.Gan.ToString(), "(",
		bazi.SiZhu.HourZhu.Gan.WuXing.ToString(), ")[",
		bazi.SiZhu.HourZhu.Gan.ShiShen.ToString(), "]")
	log.Println(
		bazi.SiZhu.YearZhu.Zhi.ToString(), "(",
		bazi.SiZhu.YearZhu.Zhi.WuXing.ToString(), ")",
		bazi.SiZhu.YearZhu.Zhi.CangGan[0].ToString(),
		bazi.SiZhu.YearZhu.Zhi.CangGan[0].WuXing.ToString(),
		bazi.SiZhu.YearZhu.Zhi.CangGan[1].ToString(),
		bazi.SiZhu.YearZhu.Zhi.CangGan[1].WuXing.ToString(),
		bazi.SiZhu.YearZhu.Zhi.CangGan[2].ToString(),
		bazi.SiZhu.YearZhu.Zhi.CangGan[2].WuXing.ToString(), "[",
		bazi.SiZhu.YearZhu.Zhi.CangGan[0].ShiShen.ToString(),
		bazi.SiZhu.YearZhu.Zhi.CangGan[1].ShiShen.ToString(),
		bazi.SiZhu.YearZhu.Zhi.CangGan[2].ShiShen.ToString(), "]",
		"\t",
		bazi.SiZhu.MonthZhu.Zhi.ToString(), "(",
		bazi.SiZhu.MonthZhu.Zhi.WuXing.ToString(), ")",
		bazi.SiZhu.MonthZhu.Zhi.CangGan[0].ToString(),
		bazi.SiZhu.MonthZhu.Zhi.CangGan[0].WuXing.ToString(),
		bazi.SiZhu.MonthZhu.Zhi.CangGan[1].ToString(),
		bazi.SiZhu.MonthZhu.Zhi.CangGan[1].WuXing.ToString(),
		bazi.SiZhu.MonthZhu.Zhi.CangGan[2].ToString(),
		bazi.SiZhu.MonthZhu.Zhi.CangGan[2].WuXing.ToString(), "[",
		bazi.SiZhu.MonthZhu.Zhi.CangGan[0].ShiShen.ToString(),
		bazi.SiZhu.MonthZhu.Zhi.CangGan[1].ShiShen.ToString(),
		bazi.SiZhu.MonthZhu.Zhi.CangGan[2].ShiShen.ToString(), "]",
		"\t",
		bazi.SiZhu.DayZhu.Zhi.ToString(), "(",
		bazi.SiZhu.DayZhu.Zhi.WuXing.ToString(), ")",
		bazi.SiZhu.DayZhu.Zhi.CangGan[0].ToString(),
		bazi.SiZhu.DayZhu.Zhi.CangGan[0].WuXing.ToString(),
		bazi.SiZhu.DayZhu.Zhi.CangGan[1].ToString(),
		bazi.SiZhu.DayZhu.Zhi.CangGan[1].WuXing.ToString(),
		bazi.SiZhu.DayZhu.Zhi.CangGan[2].ToString(),
		bazi.SiZhu.DayZhu.Zhi.CangGan[2].WuXing.ToString(), "[",
		bazi.SiZhu.DayZhu.Zhi.CangGan[0].ShiShen.ToString(),
		bazi.SiZhu.DayZhu.Zhi.CangGan[1].ShiShen.ToString(),
		bazi.SiZhu.DayZhu.Zhi.CangGan[2].ShiShen.ToString(), "]",
		"\t",
		bazi.SiZhu.HourZhu.Zhi.ToString(), "(",
		bazi.SiZhu.HourZhu.Zhi.WuXing.ToString(), ")",
		bazi.SiZhu.HourZhu.Zhi.CangGan[0].ToString(),
		bazi.SiZhu.HourZhu.Zhi.CangGan[0].WuXing.ToString(),
		bazi.SiZhu.HourZhu.Zhi.CangGan[1].ToString(),
		bazi.SiZhu.HourZhu.Zhi.CangGan[1].WuXing.ToString(),
		bazi.SiZhu.HourZhu.Zhi.CangGan[2].ToString(),
		bazi.SiZhu.HourZhu.Zhi.CangGan[2].WuXing.ToString(), "[",
		bazi.SiZhu.HourZhu.Zhi.CangGan[0].ShiShen.ToString(),
		bazi.SiZhu.HourZhu.Zhi.CangGan[1].ShiShen.ToString(),
		bazi.SiZhu.HourZhu.Zhi.CangGan[2].ShiShen.ToString(), "]")

	// 地势
	log.Println(
		bazi.DiShi.YearChangSheng.ToString(), "\t\t",
		bazi.DiShi.MonthChangSheng.ToString(), "\t\t",
		bazi.DiShi.DayChangSheng.ToString(), "\t\t",
		bazi.DiShi.HourChangSheng.ToString(), "\t\t",
	)

	// 纳音
	log.Println(
		bazi.SiZhu.YearZhu.GanZhi.NaYin.ToString(), "\t\t",
		bazi.SiZhu.MonthZhu.GanZhi.NaYin.ToString(), "\t\t",
		bazi.SiZhu.DayZhu.GanZhi.NaYin.ToString(), "\t\t",
		bazi.SiZhu.HourZhu.GanZhi.NaYin.ToString(), "\t\t",
	)

	// TODO 五行数量
	log.Printf("不计藏干的五行数量: 金(%d)木(%d)水(%d)火(%d)土(%d)", bazi.WuXing.Result[0], bazi.WuXing.Result[1], bazi.WuXing.Result[2], bazi.WuXing.Result[3], bazi.WuXing.Result[4])
	log.Printf("计入藏干后的五行数量: 金(%d)木(%d)水(%d)火(%d)土(%d)", bazi.WuXing.CangGanResult[0], bazi.WuXing.CangGanResult[1], bazi.WuXing.CangGanResult[2], bazi.WuXing.CangGanResult[3], bazi.WuXing.CangGanResult[4])
	// TODO 五行力量
	wuxingSumFun := func() int {
		var sum int
		for _, v := range bazi.XiYong.WuXingWeight {
			sum += v
		}
		return sum
	}
	sum := wuxingSumFun()
	log.Printf("五行力量: 金(%.2f%%)木(%.2f%%)水(%.2f%%)火(%.2f%%)土(%.2f%%)",
		float64(bazi.XiYong.WuXingWeight[0]*100)/float64(sum),
		float64(bazi.XiYong.WuXingWeight[1]*100)/float64(sum),
		float64(bazi.XiYong.WuXingWeight[2]*100)/float64(sum),
		float64(bazi.XiYong.WuXingWeight[3]*100)/float64(sum),
		float64(bazi.XiYong.WuXingWeight[4]*100)/float64(sum),
	)
	// 根据日干五行, 计算出同类和异类
	log.Println("五行同类", bazi.XiYong.Same)
	log.Println("五行异类", bazi.XiYong.Diff)
	if bazi.XiYong.Same >= bazi.XiYong.Diff {
		log.Printf("身强 %d, 同类强度%.2f%%\n", bazi.XiYong.Same-bazi.XiYong.Diff, float64(100*bazi.XiYong.Same)/float64(bazi.XiYong.Diff+bazi.XiYong.Same))
	} else {
		log.Printf("身弱 %d, 同类强度%.2f%%\n", bazi.XiYong.Diff-bazi.XiYong.Same, float64(100*bazi.XiYong.Same)/float64(bazi.XiYong.Diff+bazi.XiYong.Same))
	}
	// TODO 判断旺缺,<7太弱;<=16中;>17太旺
	weight := bazi.XiYong.WuXingWeight[bazi.SiZhu.DayZhu.Gan.WuXing.Value]
	shengwo := GetWuXingShengWo(bazi.SiZhu.DayZhu.Gan.WuXing.Value)
	//wosheng := GetWuXingWoSheng(bazi.SiZhu.DayZhu.Gan.WuXing.Value)
	woke := GetWuXingWoKe(bazi.SiZhu.DayZhu.Gan.WuXing.Value)
	kewo := GetWuXingKeWo(bazi.SiZhu.DayZhu.Gan.WuXing.Value)
	if weight <= 7 || (bazi.XiYong.Diff-bazi.XiYong.Same >= 10 && bazi.XiYong.Diff-bazi.XiYong.Same > 0) {
		log.Printf("日元太弱, 喜%s、%s(%s、%s), 忌%s、%s(%s、%s)",
			bazi.SiZhu.DayZhu.Gan.WuXing.ToString(), GetWuXingFromNumber(shengwo),
			"比劫", GetShiShenFromDayZhuValue(shengwo, bazi.SiZhu.DayZhu.Gan.WuXing.Value),
			GetWuXingFromNumber(woke), GetWuXingFromNumber(kewo),
			GetShiShenFromDayZhuValue(woke, bazi.SiZhu.DayZhu.Gan.WuXing.Value), GetShiShenFromDayZhuValue(kewo, bazi.SiZhu.DayZhu.Gan.WuXing.Value),
		)
	} else if 7 < weight && weight <= 10 || (bazi.XiYong.Diff-bazi.XiYong.Same < 10 && bazi.XiYong.Diff-bazi.XiYong.Same >= 5) {
		log.Printf("日元偏弱, 喜%s、%s(%s、%s), 忌%s、%s(%s、%s)",
			bazi.SiZhu.DayZhu.Gan.WuXing.ToString(), GetWuXingFromNumber(shengwo),
			"比劫", GetShiShenFromDayZhuValue(shengwo, bazi.SiZhu.DayZhu.Gan.WuXing.Value),
			GetWuXingFromNumber(woke), GetWuXingFromNumber(kewo),
			GetShiShenFromDayZhuValue(woke, bazi.SiZhu.DayZhu.Gan.WuXing.Value), GetShiShenFromDayZhuValue(kewo, bazi.SiZhu.DayZhu.Gan.WuXing.Value),
		)
	} else if 10 < weight && weight < 15 && math.Abs(float64(bazi.XiYong.Same-bazi.XiYong.Diff)) < 5 {
		log.Printf("日元中和")
	} else if 15 <= weight && weight <= 17 && bazi.XiYong.Same >= bazi.XiYong.Diff {
		log.Printf("日元偏旺, 喜%s、%s(%s、%s), 忌%s、%s(%s、%s)",
			GetWuXingFromNumber(woke), GetWuXingFromNumber(kewo),
			GetShiShenFromDayZhuValue(woke, bazi.SiZhu.DayZhu.Gan.WuXing.Value), GetShiShenFromDayZhuValue(kewo, bazi.SiZhu.DayZhu.Gan.WuXing.Value),
			bazi.SiZhu.DayZhu.Gan.WuXing.ToString(), GetWuXingFromNumber(shengwo),
			"比劫", GetShiShenFromDayZhuValue(shengwo, bazi.SiZhu.DayZhu.Gan.WuXing.Value),
		)
	} else if 17 < weight && bazi.XiYong.Same >= bazi.XiYong.Diff { // >17
		log.Printf("日元太旺, 喜%s、%s(%s、%s), 忌%s、%s(%s、%s)",
			GetWuXingFromNumber(woke), GetWuXingFromNumber(kewo),
			GetShiShenFromDayZhuValue(woke, bazi.SiZhu.DayZhu.Gan.WuXing.Value), GetShiShenFromDayZhuValue(kewo, bazi.SiZhu.DayZhu.Gan.WuXing.Value),
			bazi.SiZhu.DayZhu.Gan.WuXing.ToString(), GetWuXingFromNumber(shengwo),
			"比劫", GetShiShenFromDayZhuValue(shengwo, bazi.SiZhu.DayZhu.Gan.WuXing.Value),
		)
	} else {
		log.Printf("UnKnown")
	}

	// TODO 十神力量
	shishenSumFun := func() int {
		var sum int
		for _, v := range bazi.XiYong.ShiShenWeight {
			sum += v
		}
		return sum
	}
	sum = shishenSumFun()
	log.Printf("十神力量: 比肩(%.2f%%)劫财(%.2f%%)食神(%.2f%%)伤官(%.2f%%)偏财(%.2f%%)正财(%.2f%%)七杀(%.2f%%)正官(%.2f%%)偏印(%.2f%%)正印(%.2f%%)",
		float64(bazi.XiYong.ShiShenWeight[0]*100)/float64(sum),
		float64(bazi.XiYong.ShiShenWeight[1]*100)/float64(sum),
		float64(bazi.XiYong.ShiShenWeight[2]*100)/float64(sum),
		float64(bazi.XiYong.ShiShenWeight[3]*100)/float64(sum),
		float64(bazi.XiYong.ShiShenWeight[4]*100)/float64(sum),
		float64(bazi.XiYong.ShiShenWeight[5]*100)/float64(sum),
		float64(bazi.XiYong.ShiShenWeight[6]*100)/float64(sum),
		float64(bazi.XiYong.ShiShenWeight[7]*100)/float64(sum),
		float64(bazi.XiYong.ShiShenWeight[8]*100)/float64(sum),
		float64(bazi.XiYong.ShiShenWeight[9]*100)/float64(sum),
	)

	// TODO 八字命格
	// TODO 八字神煞
	// TODO 安命宫 http://www.131.com.tw/word/b3_2_7.htm

	// 天干五合
	log.Println(
		"天干五合：",
		bazi.HeHuaChong.TgWuHe[0].Str,
		bazi.HeHuaChong.TgWuHe[1].Str,
		bazi.HeHuaChong.TgWuHe[2].Str,
		bazi.HeHuaChong.TgWuHe[3].Str,
	)
	// 地支三会
	log.Println(
		"地支三会：",
		bazi.HeHuaChong.DzSanHui[0].Str,
		bazi.HeHuaChong.DzSanHui[1].Str,
	)
	// 地支三合
	log.Println(
		"地支三合：",
		bazi.HeHuaChong.DzSanHe[0].Str,
		bazi.HeHuaChong.DzSanHe[1].Str,
	)

	// 地支六冲
	log.Println(
		"地支六冲：",
		bazi.HeHuaChong.DzLiuChong[0].Str,
		bazi.HeHuaChong.DzLiuChong[1].Str,
		bazi.HeHuaChong.DzLiuChong[2].Str,
		bazi.HeHuaChong.DzLiuChong[3].Str,
	)
	// 地支六害
	log.Println(
		"地支六害：",
		bazi.HeHuaChong.DzLiuHai[0].Str,
		bazi.HeHuaChong.DzLiuHai[1].Str,
		bazi.HeHuaChong.DzLiuHai[2].Str,
		bazi.HeHuaChong.DzLiuHai[3].Str,
	)

	log.Println("----------------------------------------------------------------------")
	log.Println("所属节令：")
	log.Println(GetJieQiFromNumber(bazi.PreviousJie.JieQi), bazi.PreviousJie.Year, "年",
		bazi.PreviousJie.Month, "月",
		bazi.PreviousJie.Day, "日  ",
		bazi.PreviousJie.Hour, ":",
		bazi.PreviousJie.Minute, ":",
		bazi.PreviousJie.Second)
	log.Println(GetJieQiFromNumber(bazi.NextJie.JieQi), bazi.NextJie.Year, "年",
		bazi.NextJie.Month, "月",
		bazi.NextJie.Day, "日  ",
		bazi.NextJie.Hour, ":",
		bazi.NextJie.Minute, ":",
		bazi.NextJie.Second)

	var szDaYun = "大运："
	for i := 0; i < 10; i++ {
		szDaYun = szDaYun + " " + bazi.DaYun.Zhu[i].GanZhi.ToString()
	}
	log.Println(szDaYun)
	log.Println("起运时间", bazi.DaYun.QiYun.Year, "年",
		bazi.DaYun.QiYun.Month, "月",
		bazi.DaYun.QiYun.Day, "日  ",
		bazi.DaYun.QiYun.Hour, ":",
		bazi.DaYun.QiYun.Minute, ":",
		bazi.DaYun.QiYun.Second)
	log.Println("======================================================================")
}
