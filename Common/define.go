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

// 获得五行生我关系
func GetWuXingShengWo(nValue int) int {
	if nValue == 0 {
		return 4 // 土生金
	} else if nValue == 1 {
		return 2 // 水生木
	} else if nValue == 2 {
		return 0 // 金生水
	} else if nValue == 3 {
		return 1 // 水生木
	} else {
		return 3 // 火生土
	}
}

// 获得五行我生关系
func GetWuXingWoSheng(nValue int) int {
	if nValue == 0 {
		return 2 // 金生水
	} else if nValue == 1 {
		return 3 // 木生火
	} else if nValue == 2 {
		return 1 // 水生木
	} else if nValue == 3 {
		return 4 // 火生土
	} else {
		return 0 // 土生金
	}
}

// 获得五行克我关系
func GetWuXingKeWo(nValue int) int {
	if nValue == 0 {
		return 3 // 火克金
	} else if nValue == 1 {
		return 0 // 金克木
	} else if nValue == 2 {
		return 4 // 土克水
	} else if nValue == 3 {
		return 2 // 水克火
	} else {
		return 1 // 木克土
	}
}

// 获得五行我克关系
func GetWuXingWoKe(nValue int) int {
	if nValue == 0 {
		return 1 // 金克木
	} else if nValue == 1 {
		return 4 // 木克土
	} else if nValue == 2 {
		return 3 // 水克火
	} else if nValue == 3 {
		return 0 // 火克金
	} else {
		return 2 // 土克水
	}
}

var SI_JI_XI_JI_STR = [5][4]string{
	{"喜有土、火，最忌没有土、金", "必须有水相助，忌木多", "喜有木、火，忌土多", "必须有火、土相助，忌无火、土反而有金、水，忌木多而无火"},                   // 金
	{"必须有火助，有水更好，但忌水太多，也忌土太多", "必须有水相助，忌土太多，也忌木太多", "必须有金相助，但忌金太多，须有土、火才好，但忌水多", "必须有火相助，最好有土、水"}, // 木
	{"必须有土相助，若有火，金，但忌金多", "必须有金相助，忌木多", "必须有金相助，忌土、金、水多，喜木、火", "必须有火相助，喜水多，但忌金多"},                 // 水
	{"此时必为丙火或丁火，大都不错，但忌木多、土多", "必须有水相助，最喜有金", "喜有木，忌水、土多", "必须有木相助，忌有水与金多，喜有土、水、木"},               // 火
	{"喜有火、木，喜有金而少，忌金多、木多", "喜有水、金，忌有木", "喜有火，有木，忌金、水多", "喜有火，更喜有火又有金，喜有土、木"},                      // 土
}

// 五行生克制化
var WU_XING_ZHIHUA_STR = [5][]string{
	// 同类 异类 五行生克制化宜忌	五行之性 四柱五行生克中对应需补脏腑和部位 宜从事行业与方位
	{"金土", "火木水",
		`火旺得水，&nbsp;方成相济。
	<BR>火能生土，&nbsp;土多火晦；&nbsp;强火得土，&nbsp;方止其焰。
	<BR>火能克金，&nbsp;金多火熄；&nbsp;金弱遇火，&nbsp;必见销熔。
	<BR>火赖木生，&nbsp;木多火炽；&nbsp;木能生火，&nbsp;火多木焚。　
	<BR>`,
		`火主礼，&nbsp;其性急，&nbsp;其情恭，&nbsp;其味苦，&nbsp;其色赤。&nbsp;火盛之人头小脚长，&nbsp;上尖下阔，浓眉小耳，&nbsp;精神闪烁，&nbsp;为人谦和恭敬，&nbsp;纯朴急躁。&nbsp;火衰之人则黄瘦尖楞，&nbsp;语言妄诞，&nbsp;诡诈妒毒，&nbsp;做事有始无终。`,
		`肺与大肠互为脏腑表里，&nbsp;又属气管及整个呼吸系统。&nbsp;过旺或过衰，&nbsp;较宜患大肠，&nbsp;肺，&nbsp;脐，&nbsp;咳痰，&nbsp;肝，&nbsp;皮肤，&nbsp;痔疮，&nbsp;鼻气管等方面的疾病。`,
		`宜金者，&nbsp;喜西方。&nbsp;可从事精纤材或金属工具材料，&nbsp;坚硬，&nbsp;决断，&nbsp;武术，&nbsp;鉴定，总管，&nbsp;汽车，&nbsp;交通，&nbsp;金融，&nbsp;工程，&nbsp;种子，&nbsp;开矿，&nbsp;民意代表，&nbsp;伐木，&nbsp;机械等方面的经营和工作。
`}, // 金
	{"木水", "金土火",
		`木旺得金，&nbsp;方成栋梁。
	<BR>木能生火，&nbsp;火多木焚；&nbsp;强木得火，&nbsp;方化其顽。
	<BR>木能克土，&nbsp;土多木折；&nbsp;土弱逢木，&nbsp;必为倾陷。
	<BR>木赖水生，&nbsp;水多木漂；&nbsp;水能生木，&nbsp;木多水缩。`,
		`木主仁，&nbsp;其性直，&nbsp;其情和，&nbsp;其味酸，&nbsp;其色青。&nbsp;木盛的人长得丰姿秀丽，&nbsp;骨骼修长，&nbsp;手足细腻，&nbsp;口尖发美，&nbsp;面色青白。&nbsp;为人有博爱恻隐之心，&nbsp;慈祥恺悌之意，清高慷慨，&nbsp;质朴无伪。&nbsp;木衰之人则个子瘦长，&nbsp;头发稀少，&nbsp;性格偏狭，&nbsp;嫉妒不仁。木气死绝之人则眉眼不正，&nbsp;项长喉结，&nbsp;肌肉干燥，&nbsp;为人鄙下吝啬。`,
		`肝与胆互为脏腑表里，&nbsp;又属筋骨和四肢。&nbsp;过旺或过衰，&nbsp;较宜患肝，&nbsp;胆，头，&nbsp;颈，&nbsp;四肢，&nbsp;关节，&nbsp;筋脉，&nbsp;眼，&nbsp;神经等方面的疾病。 `,
		`宜木者，&nbsp;喜东方。&nbsp;可从事木材，&nbsp;木器，&nbsp;家具，&nbsp;装潢，&nbsp;木成品，&nbsp;纸业，&nbsp;种植，养花，&nbsp;育树苗，&nbsp;敬神物品，&nbsp;香料，&nbsp;植物性素食品等经营和事业。
	`}, // 木
	{"水金", "土火木",
		`水旺得土，&nbsp;方成池沼。
	<BR>水能生木，&nbsp;木多水缩；&nbsp;强水得木，&nbsp;方泄其势。
	<BR>水能克火，&nbsp;火多水干；&nbsp;火弱遇水，&nbsp;必不熄灭。
	<BR>水赖金生，&nbsp;金多水浊；&nbsp;金能生水，&nbsp;水多金沉。`,
		`水主智，&nbsp;其性聪，&nbsp;其情善，&nbsp;其味咸，&nbsp;其色黑。&nbsp;水旺之人面黑有采，&nbsp;语言清和，为人深思熟虑，&nbsp;足智多谋，&nbsp;学识过人。&nbsp;太过则好说是非，&nbsp;飘荡贪淫。&nbsp;不及则人物短小，&nbsp;性情无常，&nbsp;胆小无略，&nbsp;行事反覆。`,
		`肾与膀胱互为脏腑表里，&nbsp;又属脑与泌尿系统。&nbsp;过旺或过衰，&nbsp;较宜患肾，膀胱，&nbsp;胫，&nbsp;足，&nbsp;头，&nbsp;肝，&nbsp;泌尿，&nbsp;阴部，&nbsp;腰部，&nbsp;耳，&nbsp;子宫，&nbsp;疝气等方面的疾病。`,
		`宜水者，&nbsp;喜北方。&nbsp;可从事航海，&nbsp;冷温不燃液体，&nbsp;冰水，&nbsp;鱼类，&nbsp;水产，&nbsp;水利，冷藏，&nbsp;冷冻，&nbsp;打捞，&nbsp;洗洁，&nbsp;扫除，&nbsp;流水，&nbsp;港口，&nbsp;泳池，&nbsp;湖池塘，&nbsp;浴池，&nbsp;冷食物买卖，&nbsp;飘游，&nbsp;奔波，&nbsp;流动，&nbsp;连续性，&nbsp;易变化，&nbsp;属水性质，&nbsp;音响性质，&nbsp;清洁性质，&nbsp;海上作业，&nbsp;迁旅，&nbsp;特技表演，&nbsp;运动，&nbsp;导游，&nbsp;旅行，&nbsp;玩具，&nbsp;魔术，&nbsp;记者，&nbsp;侦探，&nbsp;旅社，灭火器具，&nbsp;钓鱼器具，&nbsp;医疗业，&nbsp;药物经营，&nbsp;医生，&nbsp;护士，&nbsp;占卜等方面的经营和工作。
	`}, // 水
	{"火木", "水金土", `火旺得水，&nbsp;方成相济。
	<BR>火能生土，&nbsp;土多火晦；&nbsp;强火得土，&nbsp;方止其焰。
	<BR>火能克金，&nbsp;金多火熄；&nbsp;金弱遇火，&nbsp;必见销熔。
	<BR>火赖木生，&nbsp;木多火炽；&nbsp;木能生火，&nbsp;火多木焚。　
	<BR>
	`, `火主礼，&nbsp;其性急，&nbsp;其情恭，&nbsp;其味苦，&nbsp;其色赤。&nbsp;火盛之人头小脚长，&nbsp;上尖下阔，浓眉小耳，&nbsp;精神闪烁，&nbsp;为人谦和恭敬，&nbsp;纯朴急躁。&nbsp;火衰之人则黄瘦尖楞，&nbsp;语言妄诞，&nbsp;诡诈妒毒，&nbsp;做事有始无终。
	`, `心脏与小肠互为脏腑表里，&nbsp;又属血脉及整个循环系统。&nbsp;过旺或过衰，&nbsp;较宜患小肠，&nbsp;心脏，&nbsp;肩，&nbsp;血液，&nbsp;经血，&nbsp;脸部，&nbsp;牙齿，&nbsp;腹部，&nbsp;舌部等方面的疾病。`,
		`宜火者，&nbsp;喜南方。&nbsp;可从事放光，&nbsp;照明，&nbsp;光学，&nbsp;高热，&nbsp;易燃，&nbsp;油类，&nbsp;酒精类，热饮食，&nbsp;食品，&nbsp;理发，&nbsp;化妆品，&nbsp;人身装饰品，&nbsp;文艺，&nbsp;文学，&nbsp;文具，&nbsp;文化学生，&nbsp;文人，&nbsp;作家，&nbsp;写作，&nbsp;撰文，&nbsp;教员，&nbsp;校长，&nbsp;秘书，&nbsp;出版，&nbsp;公务，&nbsp;正界等方面的经营和事业。 
	`}, // 火
	{"土火", "木水金", `土旺得水，&nbsp;方能疏通。
	<BR>土能生金，&nbsp;金多土变；&nbsp;强土得金，&nbsp;方制其壅。
	<BR>土能克水，&nbsp;水多土流；&nbsp;水弱逢土，&nbsp;必为淤塞。
	<BR>土赖火生，&nbsp;火多土焦；&nbsp;火能生土，&nbsp;土多火晦。`,
		`土主信，&nbsp;其性重，&nbsp;其情厚，&nbsp;其味甘，&nbsp;其色黄。&nbsp;土盛之人圆腰廓鼻，&nbsp;眉清木秀，口才声重。&nbsp;为人忠孝至诚，&nbsp;度量宽厚，&nbsp;言必行，&nbsp;行必果。&nbsp;土气太过则头脑僵化，愚拙不明，&nbsp;内向好静。&nbsp;不及之人面色忧滞，&nbsp;面扁鼻低，&nbsp;为人狠毒乖戾，&nbsp;不讲信用，不通情理。`,
		`脾与胃互为脏腑表里，&nbsp;又属肠及整个消化系统。&nbsp;过旺或过衰，&nbsp;较宜患脾，&nbsp;胃，&nbsp;肋，&nbsp;背，&nbsp;胸，&nbsp;肺，&nbsp;肚等方面的疾病。`,
		`宜土者，&nbsp;喜中央之地，&nbsp;本地。&nbsp;可从事土产，&nbsp;地产，&nbsp;农村，&nbsp;畜牧，&nbsp;布匹，&nbsp;服装，纺织，&nbsp;石料，&nbsp;石灰，&nbsp;山地，&nbsp;水泥，&nbsp;建筑，&nbsp;房产买卖，&nbsp;雨衣，&nbsp;雨伞，&nbsp;筑堤，&nbsp;容水物品，&nbsp;当铺，&nbsp;古董，&nbsp;中间人，&nbsp;律师，&nbsp;管理，&nbsp;买卖，&nbsp;设计，&nbsp;顾问，&nbsp;丧业，&nbsp;筑墓，&nbsp;墓地管理，&nbsp;僧尼等方面的经营和事业。 
		`}, // 土
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

func GetShiShenFromDayZhuValue(nValue, nDayZhuValue int) string {
	if (nValue >= 0) && (nValue < 5) {
		switch nValue {
		case 0: // 金
			switch nDayZhuValue {
			case 0: // 金
				return GetShiShenFromNumber(0) + GetShiShenFromNumber(1) // 同类:比劫
			case 1: // 木
				return GetShiShenFromNumber(6) + GetShiShenFromNumber(7) // 克我:正官七杀
			case 2: // 水
				return GetShiShenFromNumber(8) + GetShiShenFromNumber(9) // 生我:正偏印
			case 3: // 火
				return GetShiShenFromNumber(4) + GetShiShenFromNumber(5) // 我克:正偏财
			case 4: // 土
				return GetShiShenFromNumber(2) + GetShiShenFromNumber(3) // 我生:食伤
			}
		case 1: // 木
			switch nDayZhuValue {
			case 0: // 金
				return GetShiShenFromNumber(4) + GetShiShenFromNumber(5) // 我克:正偏财
			case 1: // 木
				return GetShiShenFromNumber(0) + GetShiShenFromNumber(1) // 同类:比劫
			case 2: // 水
				return GetShiShenFromNumber(2) + GetShiShenFromNumber(3) // 我生:食伤
			case 3: // 火
				return GetShiShenFromNumber(8) + GetShiShenFromNumber(9) // 生我:正偏印
			case 4: // 土
				return GetShiShenFromNumber(6) + GetShiShenFromNumber(7) // 克我:正官七杀
			}
		case 2: // 水
			switch nDayZhuValue {
			case 0: // 金
				return GetShiShenFromNumber(2) + GetShiShenFromNumber(3) // 我生:食伤
			case 1: // 木
				return GetShiShenFromNumber(8) + GetShiShenFromNumber(9) // 生我:正偏印
			case 2: // 水
				return GetShiShenFromNumber(0) + GetShiShenFromNumber(1) // 同类:比劫
			case 3: // 火
				return GetShiShenFromNumber(6) + GetShiShenFromNumber(7) // 克我:正官七杀
			case 4: // 土
				return GetShiShenFromNumber(4) + GetShiShenFromNumber(5) // 我克:正偏财
			}
		case 3: // 火
			switch nDayZhuValue {
			case 0: // 金
				return GetShiShenFromNumber(6) + GetShiShenFromNumber(7) // 克我:正官七杀
			case 1: // 木
				return GetShiShenFromNumber(2) + GetShiShenFromNumber(3) // 我生:食伤
			case 2: // 水
				return GetShiShenFromNumber(4) + GetShiShenFromNumber(5) // 我克:正偏财
			case 3: // 火
				return GetShiShenFromNumber(0) + GetShiShenFromNumber(1) // 同类:比劫
			case 4: // 土
				return GetShiShenFromNumber(8) + GetShiShenFromNumber(9) // 生我:正偏印
			}
		case 4: // 土
			switch nDayZhuValue {
			case 0: // 金
				return GetShiShenFromNumber(8) + GetShiShenFromNumber(9) // 生我:正偏印
			case 1: // 木
				return GetShiShenFromNumber(4) + GetShiShenFromNumber(5) // 我克:正偏财
			case 2: // 水
				return GetShiShenFromNumber(6) + GetShiShenFromNumber(7) // 克我:正官七杀
			case 3: // 火
				return GetShiShenFromNumber(2) + GetShiShenFromNumber(3) // 我生:食伤
			case 4: // 土
				return GetShiShenFromNumber(0) + GetShiShenFromNumber(1) // 同类:比劫
			}
		}
	}
	return ""
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

// 神煞
type TShenSha struct {
	YearShenSha  []string
	MonthShenSha []string
	DayShenSha   []string
	HourShenSha  []string
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
