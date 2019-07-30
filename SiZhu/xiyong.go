package SiZhu

// 喜用神
// 喜用神是中国传统八字命理学上的术语， 喜用神是喜神与用神的合称。
// 八字，即把人出生的年、月、日、时分作四柱，每柱配有一天干和地支，合共八字。
// 八字不同的排列，包含不同的阴阳五行信息，构成各种不同的八字命局。
// 命局中有“不及”和“太过”等情况，称作“病”，而“用神”正是针对不同的“病”所下的“药”。
// “喜神”则是对“用神”能够起到生扶作用的阴阳五行元素。
// 四柱命局以用神为核心， 用神健全有力与否，影响人一生的命；
// 一生补救与否, 影响人一生的运。凡用神之力不足，四柱中有生助用神者，
// 或四柱刑冲克害用神而能化凶神，制凶神者，就是喜神。 四柱没有用神，就得靠行运流年来补。
// 对于命局五行较为平衡，用神不太紧缺的四柱，其一生较为平顺，无大起大落。

// 天干十神生克作用关系是以日主即日干为我而与其他七字的生克关系而确定的。其基本规贝。是:
// 生我者为印绶(正印、偏印)；
// 我生者为食伤(食神、伤官)；
// 克我者为官煞(正官、七杀)；
// 我克者为妻财(正财、偏财)；
// 同类者为比劫(比肩、劫财)。

import (
	"log"
	"strings"

	. "github.com/warrially/BaziGo/Common"
)

// 天干地支强度测试
// 三种算法
// 1) 天干地支算10分，藏干算1分, 这里采用这种算法，http://freehoro.net/BaZi/NatalXiJi/NatalXiJi.php?NUrl=&CNT=1920&User=TSwxOTg0LTEwLTA5IDE1OjU0OjAwLDgsMCw4LDAsMixOLE4sMCwx&SG1=MSwxLDEsMTEsMywxLDMsOQ&JQ1=MTk4NC0xMC0wOCAxMDo0MTo1NiwxOTU&JQ2=MTk4NC0xMS0wNyAxMzo0NTozOSwyMjU
// 2) 按表1
// 3) 按表2

// fix 天干强度表，原表似乎结果不准确, 按照http://www.131.com.tw/word/b3_2_14.htm的算法修正
var TIAN_GAN_QIANG_DU_LIST = [10]int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5} // 甲乙丙丁戊己庚辛壬癸
var DI_ZHI_QIANG_DU_LIST = [36]int{
	// 子 子 子  丑   丑   丑   寅   寅  寅  卯   卯  卯 辰   辰   辰   巳  巳  巳   午   午  午 未   未   未   申  申  申   酉   酉 酉  戌   戌   戌   亥   亥   亥
	8, 0, 0, 5, 2, 1, 5, 2, 1, 8, 0, 0, 5, 2, 1, 5, 2, 1, 5, 3, 0, 5, 2, 1, 5, 2, 1, 8, 0, 0, 5, 2, 1, 5, 3, 0}

// 五行旺度表
// TODO 判断旺61，相45，囚21，休8，死1
// 春月：木旺、火相、水休、金囚、土死；
// 夏月：火旺、土相、木休、水囚、金死；
// 秋月：金旺、水相、土休、火囚、木死；
// 冬月：水旺、木相、金休、土囚、火死；
// 四季：土旺、金相、火休、木囚、水死。
// 这里的四季，是指辰戌丑未四季月最后十八天。春月，是寅卯月，以及辰月的前十二天；夏月，是指巳午月，以及未月的前十二天；秋月，是指申酉月，以及戌月的前十二天；冬月是指亥子月以及丑月的前十二天。春、夏、秋、冬、四季各占72天。

/*
	　　　　　木　　火　　　土　　　金　　　水
	　　寅月　旺　　相　　　死　　　囚　　　休
	　　卯月　旺　　相　　　死　　　囚　　　休
	　　辰月　余气　休　　　旺　　　相　　　死或相
	　　巳月　休　　旺　　　相　　　死　　　囚
	　　午月　休　　旺　　　相　　　死　　　囚
	　　未月　囚　　余气　　旺　　　死　　　死
	　　申月　死　　囚　　　休　　　旺　　　相
	　　酉月　死　　囚　　　休　　　旺　　　相
	　　戌月　囚　　休或相　旺　　　相或死　死
	　　亥月　相　　死　　　囚　　　休　　　旺
	　　子月　相　　死　　　囚　　　休　　　旺
	　　丑月　囚　　死　　　旺或相　相　　　余气

	　　特殊情况：
	　　木生于辰月——木不以囚论，而以余气论，因为辰月为春末，木有余气；
	　　水生于辰月――当命局中水较旺时以相论，反之以死论；
	　　火生于未月——火不以休论，而以余气论，因为未月为夏末，火还有余热；
	　　金生于未月——金以死论不以相论，燥土不生金；
	　　火生于戌月——当燥土党众或原局很干燥时以相论，反之以休论；
	　　金生于戌月――当燥土党众或原局很干燥时以死论，反之以相论；
	　　水生于丑月——水不以死论而以余气论，因为丑月为冬末，水还有余气；
	　　土生于丑月——未戌土生于丑月不以旺论而以相论，辰丑月生于丑月以旺论。
	　　“旺相休囚死”五行旺度系数：
	　　旺：　2
	　　余气：1.6
	　　相：　1.5
	　　休：　0.8
	　　囚：　0.7
	　　死：　0.5
*/

var WUXING_WANGDU_LIST = [12][5]string{
	// 金木水火土
	{"囚", "旺", "休", "相", "死"},
	{"囚", "旺", "休", "相", "死"},
	{"相", "余气", "死,相", "休", "旺"},
	{"死", "休", "囚", "旺", "相"},
	{"死", "休", "囚", "旺", "相"},
	{"死", "囚", "死", "余气", "旺"},
	{"旺", "死", "相", "囚", "休"},
	{"旺", "死", "相", "囚", "休"},
	{"相,死", "囚", "死", "休,相", "旺"},
	{"休", "相", "旺", "死", "囚"},
	{"休", "相", "旺", "死", "囚"},
	{"相", "囚", "余气", "死", "旺,相"},
}
var WUXING_WANGDU_VALUE = map[string]int{
	"旺":  20,
	"余气": 16,
	"相":  15,
	"休":  8,
	"囚":  7,
	"死":  5,
}

// 天干强度表
var TIAN_GAN_QIANG_DU_LIST2 = [12][10]int{
	//甲   乙    丙    丁    戊    己    庚    辛    壬    癸
	{1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000, 1200, 1200}, //子月
	{1060, 1060, 1000, 1000, 1100, 1100, 1140, 1140, 1100, 1100}, //丑月
	{1140, 1140, 1200, 1200, 1060, 1060, 1000, 1000, 1000, 1000}, //寅月
	{1200, 1200, 1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000}, //卯月
	{1100, 1100, 1060, 1060, 1100, 1100, 1100, 1100, 1040, 1040}, //辰月
	{1000, 1000, 1140, 1140, 1140, 1140, 1060, 1060, 1060, 1060}, //巳月
	{1000, 1000, 1200, 1200, 1200, 1200, 1000, 1000, 1000, 1000}, //午月
	{1040, 1040, 1100, 1100, 1160, 1160, 1100, 1100, 1000, 1000}, //未月
	{1060, 1060, 1000, 1000, 1000, 1000, 1140, 1140, 1200, 1200}, //申月
	{1000, 1000, 1000, 1000, 1000, 1000, 1200, 1200, 1200, 1200}, //酉月
	{1000, 1000, 1040, 1040, 1140, 1140, 1160, 1160, 1060, 1060}, //戌月
	{1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000, 1140, 1140}} //亥月

// 地支强度表
var DI_ZHI_QIANG_DU_LIST2 = [12][36]int{
	// 子 子 子  丑   丑   丑   寅   寅  寅  卯   卯  卯 辰   辰   辰   巳  巳  巳   午   午  午 未   未   未   申  申  申   酉   酉 酉  戌   戌   戌   亥   亥   亥
	// 癸        己   癸   辛   甲   丙      乙          戊   乙   癸   丙  戊  庚   丁   己     己   乙   丁   庚      壬   辛          戊   辛   丁   壬   甲
	{1000, 0, 0, 530, 300, 200, 798, 360, 0, 1140, 0, 0, 530, 342, 200, 840, 0, 300, 1200, 0, 0, 530, 228, 360, 700, 0, 300, 1000, 0, 0, 530, 300, 240, 700, 342, 0}, // 寅月
	{1000, 0, 0, 500, 300, 200, 840, 360, 0, 1200, 0, 0, 500, 360, 200, 840, 0, 300, 1200, 0, 0, 500, 240, 360, 700, 0, 300, 1000, 0, 0, 500, 300, 240, 700, 360, 0}, // 卯月
	{1040, 0, 0, 550, 312, 230, 770, 318, 0, 1100, 0, 0, 550, 330, 208, 742, 0, 330, 1060, 0, 0, 550, 220, 318, 770, 0, 312, 1100, 0, 0, 550, 330, 212, 728, 330, 0}, // 辰月
	{1060, 0, 0, 570, 318, 212, 700, 342, 0, 1000, 0, 0, 600, 300, 200, 840, 0, 300, 1140, 0, 0, 570, 200, 342, 742, 0, 318, 1060, 0, 0, 570, 318, 228, 742, 300, 0}, // 巳月
	{1000, 0, 0, 600, 300, 200, 700, 360, 0, 1000, 0, 0, 600, 300, 200, 840, 0, 300, 1200, 0, 0, 600, 200, 360, 700, 0, 300, 1000, 0, 0, 600, 300, 240, 700, 300, 0}, // 午月
	{1000, 0, 0, 580, 300, 220, 728, 330, 0, 1040, 0, 0, 580, 312, 200, 798, 0, 330, 1100, 0, 0, 580, 208, 330, 770, 0, 300, 1100, 0, 0, 580, 330, 220, 700, 312, 0}, // 未月
	{1200, 0, 0, 500, 360, 228, 742, 300, 0, 1060, 0, 0, 500, 318, 240, 700, 0, 342, 1000, 0, 0, 500, 212, 300, 798, 0, 360, 1140, 0, 0, 500, 342, 200, 840, 318, 0}, // 申月
	{1200, 0, 0, 500, 360, 248, 700, 300, 0, 1000, 0, 0, 500, 300, 240, 700, 0, 360, 1000, 0, 0, 500, 200, 300, 840, 0, 360, 1200, 0, 0, 500, 360, 200, 840, 300, 0}, // 酉月
	{1060, 0, 0, 570, 318, 232, 700, 342, 0, 1000, 0, 0, 570, 300, 212, 728, 0, 348, 1040, 0, 0, 570, 200, 312, 812, 0, 318, 1160, 0, 0, 570, 348, 208, 724, 300, 0}, // 戌月
	{1140, 0, 0, 500, 342, 200, 840, 318, 0, 1200, 0, 0, 500, 360, 228, 742, 0, 300, 1060, 0, 0, 500, 240, 318, 700, 0, 342, 1000, 0, 0, 500, 300, 212, 798, 360, 0}, // 亥月
	{1200, 0, 0, 500, 360, 200, 840, 300, 0, 1200, 0, 0, 500, 360, 240, 700, 0, 300, 1000, 0, 0, 500, 240, 300, 700, 0, 360, 1000, 0, 0, 500, 300, 200, 840, 360, 0}, // 子月
	{1100, 0, 0, 550, 330, 228, 742, 300, 0, 1060, 0, 0, 550, 318, 220, 700, 0, 342, 1000, 0, 0, 550, 212, 300, 798, 0, 330, 1140, 0, 0, 550, 342, 200, 770, 318, 0}} // 丑月

// 计算喜用神
func CalcXiYong(month int, pSiZhu *TSiZhu, pWuXing *TWuXingStat) TXiYong {
	var xiyong TXiYong

	// 1. 通过四柱计算天干地支强度,以及十神强度
	var wuxing = [5]int{0, 0, 0, 0, 0} // 金木水火土
	var wuxingCount = pWuXing.CangGanResult
	var shishen = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // 比劫食伤才财杀官卩印

	// 2. 拿到四柱的月支
	var nMonthZhi = pSiZhu.MonthZhu.Zhi.Value
	// log.Println("月支是", nMonthZhi, pSiZhu.MonthZhu.Zhi.Str)

	// 3. 计算旺度值
	for wx, count := range wuxingCount {
		wuxing[wx] = getWangDuValue(month, wx) * count
	}
	log.Println("计算完毕天干后的五行权值是:", wuxing)

	// 4. 根据四柱天干, 换算强度
	shishen[pSiZhu.YearZhu.Gan.ShiShen.Value] += 10
	shishen[pSiZhu.MonthZhu.Gan.ShiShen.Value] += 10
	shishen[GetShiShenFromGan(pSiZhu.DayZhu.Gan.Value, pSiZhu.DayZhu.Gan.Value)] += 10
	shishen[pSiZhu.HourZhu.Gan.ShiShen.Value] += 10

	// 4. 根据四柱地支, 换算强度, 计算十神权值
	for i := 0; i < 3; i++ {
		// 年
		var nCangGan = pSiZhu.YearZhu.Zhi.CangGan[i].Value
		if nCangGan >= 0 {
			//wuxing[pSiZhu.YearZhu.Zhi.CangGan[i].WuXing.Value] += 1
			shishen[pSiZhu.YearZhu.Zhi.CangGan[i].ShiShen.Value] += 1
		}

		// 月
		nCangGan = pSiZhu.MonthZhu.Zhi.CangGan[i].Value
		if nCangGan >= 0 {
			//wuxing[pSiZhu.MonthZhu.Zhi.CangGan[i].WuXing.Value] += 1
			shishen[pSiZhu.MonthZhu.Zhi.CangGan[i].ShiShen.Value] += 1
		}

		// 日
		nCangGan = pSiZhu.DayZhu.Zhi.CangGan[i].Value
		if nCangGan >= 0 {
			//wuxing[pSiZhu.DayZhu.Zhi.CangGan[i].WuXing.Value] += 1
			shishen[pSiZhu.DayZhu.Zhi.CangGan[i].ShiShen.Value] += 1
		}

		// 时
		nCangGan = pSiZhu.HourZhu.Zhi.CangGan[i].Value
		if nCangGan >= 0 {
			//wuxing[pSiZhu.HourZhu.Zhi.CangGan[i].WuXing.Value] += 1
			shishen[pSiZhu.HourZhu.Zhi.CangGan[i].ShiShen.Value] += 1
		}
	}

	// 5. 根据日干五行, 计算出同类和异类
	var nDayWuXing = pSiZhu.DayZhu.Gan.WuXing.Value
	xiyong.Same, xiyong.Diff, xiyong.SameList, xiyong.DiffList = CalcWuXingQiangRuo(nDayWuXing, wuxing)
	// 月支
	xiyong.MonthZhi = nMonthZhi
	// 日五行
	xiyong.DayWuXing = nDayWuXing
	// 五行权值
	xiyong.WuXingWeight = wuxing
	// 十神权值
	xiyong.ShiShenWeight = shishen

	return xiyong
}

// 计算五行强弱
func CalcWuXingQiangRuo(nDayWuXing int, wuxing [5]int) (int, int, []int, []int) {
	var nSame = 0 // 同类
	var nDiff = 0 // 异类
	var aSame = []int{}
	var aDiff = []int{}

	// 自己
	nSame += wuxing[nDayWuXing]

	switch nDayWuXing {
	case 0: // 金 同类土
		nSame += wuxing[4]
		nDiff += wuxing[1] + wuxing[2] + wuxing[3]
		aSame = append(aSame, []int{0, 4}...)
		aDiff = append(aDiff, []int{1, 2, 3}...)
	case 1: // 木 同类水
		nSame += wuxing[2]
		nDiff += wuxing[0] + wuxing[3] + wuxing[4]
		aSame = append(aSame, []int{1, 2}...)
		aDiff = append(aDiff, []int{0, 3, 4}...)
	case 2: // 水 同类金
		nSame += wuxing[0]
		nDiff += wuxing[1] + wuxing[3] + wuxing[4]
		aSame = append(aSame, []int{0, 2}...)
		aDiff = append(aDiff, []int{1, 3, 4}...)
	case 3: // 火 同类木
		nSame += wuxing[1]
		nDiff += wuxing[0] + wuxing[2] + wuxing[4]
		aSame = append(aSame, []int{1, 3}...)
		aDiff = append(aDiff, []int{0, 2, 4}...)
	case 4: // 土 同类火
		nSame += wuxing[3]
		nDiff += wuxing[0] + wuxing[1] + wuxing[2]
		aSame = append(aSame, []int{4, 3}...)
		aDiff = append(aDiff, []int{0, 2, 1}...)
	}
	return nSame, nDiff, aSame, aDiff
}

func getWangDuValue(month, wuxing int) int {
	month -= 1
	wangdu := WUXING_WANGDU_LIST[month][wuxing]
	list := strings.Split(wangdu, ",")
	// TODO 暂时取第一个
	log.Println("---- ", month, wuxing, list[0])
	return WUXING_WANGDU_VALUE[list[0]]
}
