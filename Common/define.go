package Common

// 日期
type TDate struct {
	Year   int // 年
	Month  int // 月
	Day    int // 日
	Hour   int // 时
	Minute int // 分
	Second int // 秒
	JieQi  int // 节气
}

// 五行属性
type TWuXing struct {
	Value int // 五行
}

// 五行统计
type TWuXingStat struct {
	Result        [5]int
	CangGanResult [5]int
}

// 十二长生
type TChangSheng struct {
	Value int
}
type TDiShi struct {
	YearChangSheng  TChangSheng
	MonthChangSheng TChangSheng
	DayChangSheng   TChangSheng
	HourChangSheng  TChangSheng
}

// {* 五行字符串，以通常的金木水火土为顺序 }
// 这里没用五行相生或者相克来排列
var WU_XING_STR = [5]string{
	"金", "木", "水", "火", "土"}

// 从数字获得五行名, 0-4
func GetWuXingFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 5) {
		return WU_XING_STR[nValue]
	}
	return ""
}

func (self *TWuXing) ToString() string {
	return GetWuXingFromNumber(self.Value)
}

// 十神属性
type TShiShen struct {
	Value int // 十神
}

// 十神字符串
var SHI_SHEN_STR = [10]string{
	"比", "劫", "食", "伤", "才",
	"财", "杀", "官", "卩", "印"}

func GetShiShenFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 10) {
		return SHI_SHEN_STR[nValue]
	}
	return ""
}

func (self *TShiShen) ToString() string {
	return GetShiShenFromNumber(self.Value)
}

// 纳音
type TNaYin struct {
	Value int // 纳音五行
}

//  {* 纳音五行，与相邻一对六十干支对应}
// 甲子乙丑海中金丙寅丁卯炉中火戊辰己巳大林木
// 庚午辛未路旁土壬申癸酉剑锋金甲戌乙亥山头火
// 丙子丁丑涧下水戊寅己卯城头土庚辰辛巳白蜡金
// 壬午癸未杨柳木甲申乙酉井泉水丙戌丁亥屋上土
// 戊子己丑霹雳火庚寅辛卯松柏木壬辰癸巳长流水
// 甲午乙未砂中金丙申丁酉山下火戊戌己亥平地木
// 庚子辛丑壁上土壬寅癸卯金箔金甲辰乙巳覆灯火
// 丙午丁未天河水戊申己酉大驿土庚戌辛亥钗钏金
// 壬子癸丑桑柘木甲寅乙卯大溪水丙辰丁巳砂中土
// 戊午己未天上火庚申辛酉石榴木壬戌癸亥大海水
var NA_YIN_STR = [30]string{
	"海中金", "炉中火", "大林木",
	"路旁土", "剑锋金", "山头火",

	"涧下水", "城墙土", "白蜡金",
	"杨柳木", "泉中水", "屋上土",

	"霹雷火", "松柏木", "长流水",
	"沙中金", "山下火", "平地木",

	"壁上土", "金箔金", "佛灯火",
	"天河水", "大驿土", "钗钏金",

	"桑柘木", "大溪水", "沙中土",
	"天上火", "石榴木", "大海水"}

func GetNaYinFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 30) {
		return NA_YIN_STR[nValue]
	}
	return ""
}

func (self *TNaYin) ToString() string {
	return GetNaYinFromNumber(self.Value)
}

// 干支属性
type TGanZhi struct {
	Value int    // 干支0-59 对应 甲子到癸亥
	NaYin TNaYin // 纳音
}

// 60 干支
var GAN_ZHI_STR = [60]string{
	"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉",
	"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未",
	"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳",
	"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯",
	"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑",
	"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥"}

// 从数字获得天干地支名, 0-59
func GetGanZhiFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 60) {
		return GAN_ZHI_STR[nValue]
	}
	return "未知"
}

func (self *TGanZhi) ToString() string {
	return GetGanZhiFromNumber(self.Value)
}

// 干属性
type TGan struct {
	Value   int      // 天干
	WuXing  TWuXing  // 天干五行
	ShiShen TShiShen // 天干十神
}

// {* 天干字符串，Heavenly Stems}
var TIAN_GAN_STR = [10]string{
	"甲", "乙", "丙", "丁", "戊",
	"己", "庚", "辛", "壬", "癸"}

// 从数字获得天干名, 0-9
func GetTianGanFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 10) {
		return TIAN_GAN_STR[nValue]
	}
	return ""
}

func (self *TGan) ToString() string {
	return GetTianGanFromNumber(self.Value)
}

// 支属性
type TZhi struct {
	Value   int     // 地支
	WuXing  TWuXing // 地支五行
	CangGan [3]TGan // 藏干
}

// {* 地支字符串，Earthly Branches}
var DI_ZHI_STR = [12]string{
	"子", "丑", "寅", "卯",
	"辰", "巳", "午", "未",
	"申", "酉", "戌", "亥"}

// 从数字获得地支名, 0-11
func GetDiZhiFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 12) {
		return DI_ZHI_STR[nValue]
	}
	return ""
}

func (self *TZhi) ToString() string {
	return GetDiZhiFromNumber(self.Value)
}

// 柱子
type TZhu struct {
	GanZhi TGanZhi // 干支
	Gan    TGan    // 天干
	Zhi    TZhi    // 地支
}

// 四柱
type TSiZhu struct {
	YearZhu  TZhu // 年柱
	MonthZhu TZhu // 月柱
	DayZhu   TZhu // 日柱
	HourZhu  TZhu // 时柱
}

// 喜用神
type TXiYong struct {
	MonthZhi      int     // 月支
	DayWuXing     int     // 日干五行
	Same          int     // 同类
	Diff          int     // 异类
	WuXingWeight  [5]int  // 五行权值
	ShiShenWeight [10]int // 十神权值
}

// 大运
type TDaYun struct {
	Zhu    [12]TZhu // 柱
	QiYun  TDate    // XX年XX月开始起运
	ShunNi bool     // 顺转还是逆转(true 顺,  false 逆)
}

// 天干五合
type TTgWuHe struct {
	Gan1 int    // 干1
	Gan2 int    // 干2
	He   int    // 和的结果
	Str  string // 描述
}

// 地支三会
type TDzSanHui struct {
	Zhi1 int    // 支1
	Zhi2 int    // 支2
	Zhi3 int    // 支3
	Hui  int    // 三会结果
	Str  string // 描述
}

// 地支三合
type TDzSanHe struct {
	Zhi1 int    // 支1
	Zhi2 int    // 支2
	Zhi3 int    // 支3
	He   int    // 三合结果
	Str  string // 描述
}

// 地支六冲
type TDzLiuChong struct {
	Zhi1 int    // 支1
	Zhi2 int    // 支2
	Str  string // 描述
}

// 地支六害
type TDzLiuHai struct {
	Zhi1 int    // 支1
	Zhi2 int    // 支2
	Str  string // 描述
}

// 合化冲
type THeHuaChong struct {
	TgWuHe     [4]TTgWuHe     // 天干五合
	DzSanHui   [2]TDzSanHui   // 地支三会
	DzSanHe    [2]TDzSanHe    // 地支三合
	DzLiuChong [4]TDzLiuChong // 地支六冲
	DzLiuHai   [4]TDzLiuHai   // 地支六害
}

var JIE_QI_STR = [24]string{
	"立春", // 节气  Beginning of Spring   0
	"雨水", // 中气  Rain Water            1
	"惊蛰", // 节气  Waking of Insects     2
	"春分", // 中气  March Equinox         3
	"清明", // 节气  Pure Brightness       4
	"谷雨", // 中气  Grain Rain            5
	"立夏", // 节气  Beginning of Summer   6
	"小满", // 中气  Grain Full            7
	"芒种", // 节气  Grain in Ear          8
	"夏至", // 中气  Summer Solstice       9
	"小暑", // 节气  Slight Heat           10
	"大暑", // 中气  Great Heat            11
	"立秋", // 节气  Beginning of Autumn   12
	"处暑", // 中气  Limit of Heat         13
	"白露", // 节气  White Dew             14
	"秋分", // 中气  September Equinox     15
	"寒露", // 节气  Cold Dew              16
	"霜降", // 中气  Descent of Frost      17
	"立冬", // 节气  Beginning of Winter   18
	"小雪", // 中气  Slight Snow           19
	"大雪", // 节气  Great Snow            20
	"冬至", // 中气  Winter Solstice       21
	"小寒", // 节气  Slight Cold           22，这是一公历年中的第一个节气
	"大寒"} // 中气  Great Cold            23

// 从数字获得节气名, 0-23
func GetJieQiFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 24) {
		return JIE_QI_STR[nValue]
	}
	return "未知"
}

// "阴", "阳"
var YIN_YANG_STR = [2]string{
	"阴", "阳"}

func GetYinYangFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 2) {
		return YIN_YANG_STR[nValue]
	}
	return ""
}

var LUNAR_MONTH_STR = [12]string{
	"一月", "二月", "三月", "四月", "五月", "六月",
	"七月", "八月", "九月", "十月", "冬月", "腊月"}

func GetLunarMonthFromNumber(nValue int) string {
	if (nValue >= 1) && (nValue <= 12) {
		return LUNAR_MONTH_STR[nValue-1]
	}

	return ""
}

var LUNAR_DAY_STR = [30]string{
	"初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十",
	"十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十",
	"廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十"}

func GetLunarDayFromNumber(nValue int) string {
	if (nValue >= 1) && (nValue <= 30) {
		return LUNAR_DAY_STR[nValue-1]
	}

	return ""
}

// 十二长生表，以日干对应四柱查表
//https://sites.google.com/site/laughing8word/shi-er-zhang-sheng-cha-biao-fa
var SHI_ER_CHANG_SHENG_STR = [12]string{
	"长生", "沐浴", "冠带", "临官", "帝旺", "衰", "病", "死", "墓", "绝", "胎", "养"}

var SHI_ER_CHANG_SHENG_LIST = [10][12]int{
	// 长生 沐浴 冠带 临官 帝旺 衰 病 死 墓 绝 胎 养
	{11, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, // 甲
	{6, 5, 4, 3, 2, 1, 0, 11, 10, 9, 8, 7}, // 乙
	{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0, 1}, // 丙
	{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 11, 10}, // 丁
	{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0, 1}, // 戊
	{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 11, 10}, // 己
	{5, 6, 7, 8, 9, 10, 11, 0, 1, 2, 3, 4}, // 庚
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, // 辛
	{8, 9, 10, 11, 0, 1, 2, 3, 4, 5, 6, 7}, // 壬
	{3, 2, 1, 0, 11, 10, 9, 8, 7, 6, 5, 4}, // 癸
}

func (self *TChangSheng) ToString() string {
	return GetChangShengFromNumber(self.Value)
}

// 从数字获得十二长生运名字
func GetChangShengFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 12) {
		return SHI_ER_CHANG_SHENG_STR[nValue]
	}
	return ""
}
