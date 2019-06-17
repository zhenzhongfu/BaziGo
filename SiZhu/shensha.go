package SiZhu

import (
	"strings"

	. "github.com/warrially/BaziGo/Common"
)

// 神煞排法
// http://www.131.com.tw/word/b3_2_10.htm
// http://piscesdai.pixnet.net/blog/post/370873481-%E5%85%AB%E5%AD%97%E6%8E%92%E6%B3%95%E6%95%99%E5%AD%B8%E3%80%90%E7%A5%9E%E7%85%9E%E6%98%9F%E6%A0%BC%E5%B1%80%E3%80%91

// 神煞的解释
var SHEN_SHA_MEANING_MAP = map[string]string{
	"天上三奇": "命帶三奇者，看起來總是博學多能，一生抱負極大，胸襟寬闊，有功名，有成就，精神超人，人命若得三奇貴，為人聰穎，又帶天月二德貴人者，一般凶災不侵，富貴不淫。威武不屈，是榮華福壽之人。",
	"地上三奇": "命帶三奇者，看起來總是博學多能，一生抱負極大，胸襟寬闊，有功名，有成就，精神超人，人命若得三奇貴，為人聰穎，又帶天月二德貴人者，一般凶災不侵，富貴不淫。威武不屈，是榮華福壽之人。",
	"人中三奇": "命帶三奇者，看起來總是博學多能，一生抱負極大，胸襟寬闊，有功名，有成就，精神超人，人命若得三奇貴，為人聰穎，又帶天月二德貴人者，一般凶災不侵，富貴不淫。威武不屈，是榮華福壽之人。",
	"六秀日":  "丙午、丁末、戊子、戊午、己丑、己未日生者。命帶六秀日主其人秀氣且聰明，有才藝，博學多能，有功名成就。",
	"魁罡":   "戊戌、庚辰、庚戌、壬辰、年月日時有即算。命帶魁罡其人膽大正直，性情剛直，行事果斷，領導力強，博學多能，掌權威，為人嚴謹操守佳，一生命運起伏大，好壞運呈明顯，急性又好勝，積極又勤勞，不怕鬼邪，神鬼避之。男帶魁罡，做官及事業成就大；女帶魁罡，性剛烈。恐婦奪夫權，雖然多半為美女，但一生之命運也多波折。",
	"天赦":   "命帶天赦者可增吉祥，凡事能逢凶不凶，是一顆能為人解災禍之星，天赦入命，若日主得旺，代表可得高官權貴，命中若逢天赦，一生處世無憂，貴人扶助，若日主不得旺時，吉事減半。",
	"將星":   "入命具傑出之領導能力，易掌權柄。將星忌會凶星，會助長凶煞之力。將星坐財，主掌財政。坐官、掌官貴權位。坐殺、掌生殺之權。將星是一顆權力之星，文武兩相宜。",
	"驛馬":   "入命，主奔波忙碌，經常外出，外向，心性較不穩定，適合出外發展，可能常更換職業。適合從事行業為外交工作、外務人員、演員、司機、飲食店、作家、旅遊業、船員、特種營業、推銷商。",
	"亡神":   "又名官符，失去也，入命為喜用時，有權有威，深謀略算計，富心機，凡事不露底，城府較深。入命為忌神時，心胸狹窄急躁，酒色風流。若與凶煞同柱，其凶性更明顯。主官司訴訟，失勢也。",
	"孤辰":   "命帶孤辰者，性較孤獨，面無和氣，較不利六親，同柱坐生旺稍可化解，如坐死絕不好，一生運較生不逢時，孤辰以日柱，時柱為重，六親或夫妻或子女互動較少。 孤辰在年柱時，最好過房，若不過房，有可能會刑剋父母。 男命帶孤辰，正偏財星又逢死絕，表婚緣不佳。女命帶孤寡，如正官七殺逢死絕，也表婚姻不美，宜多溝通。",
	"寡宿":   "命帶孤寡者，性較孤獨，面無和氣，較不利六親，同柱坐生旺稍可化解，如坐死絕不好，一生運較生不逢時，孤寡以日柱，時柱為重，六親或夫妻或子女互動較少。 孤寡在年柱時，最好過房，若不過房，有可能會刑剋父母。 男命帶孤寡，正偏財星又逢死絕，表婚緣不佳。女命帶孤寡，如正官七殺逢死絕，也表婚姻不美，宜多溝通。",
	"華蓋":   "入命主思想獨特，一生具備才藝、技藝、藝術、音樂、設計，具審美之天份，聰明好學喜研究哲理、玄學、宗教等學問，有出世之心。為人寡慾，孤僻。與神佛有緣。一生中可發揮個人潛能，具孤獨之性向，六親較無多大助力。 華蓋喜逢正印正官 ( 同柱最佳 )，則是在官場社會上能有地位及成就。 華蓋落空亡，比較容易走僧道、和尚、尼姑，或孤寡路線。 華蓋落死絕或被沖剋，則幼兒較多災多病，不易扶養，生平亦多是非不順。",
	"隔角":   "命帶隔角一生易犯官司訴訟，或牢獄之災。據經驗顯示命帶隔角者，並非一定主牢獄官訟之災，有的人只是做事或出外，易遭困難阻礙隔絕之情況發生。",
	"日破":   "命帶日破表是個性不太穩定，較不受拘束，也較不安於家中，或配偶身體差，或夫妻較有離異之可能，或夫妻間易爭吵；或與子女不和，或子女健康差。",
	"天乙貴人": "天乙乃尊貴之神也，所至之地一切凶煞自然而避，凡事能逢凶化吉。入命主聰明，有智慧，人緣佳，易得人助，能交際，有功名，運不錯，可得富貴也。",
	"太極貴人": "主其人聪明好学，为人正直，有钻劲，做事有始有终，喜天文、地理、哲学，与佛道、宗教有缘。如命带多个太极贵人，其人天生对命理、玄学、术数感兴趣，并有这方面的天分，如从事这方面的研究，容易有大成就。假如命带太极贵人，又逢生旺或有其他吉星辅助，主其人能福寿双全，不是朝廷做官人，定是人间富贵人。",
	"十惡大敗": "帶有此神煞的人一生花錢如流水，不聚財，生活比較困頓和艱難。",
	"丧门":   "丧门、吊客皆为凶星，如流年逢上此星，多主亲人亡故，哭泣孝服，伤病等事出现，俗说“逢丧门星临门”大主不吉，故流年逢此凶星，须要多加小心。",
	"吊客":   "丧门、吊客皆为凶星，如流年逢上此星，多主亲人亡故，哭泣孝服，伤病等事出现，俗说“逢丧门星临门”大主不吉，故流年逢此凶星，须要多加小心。",
	"金神":   "命带金神即命局中出现金神，命局主人一般性情刚烈，聪明有才，好学不倦，生性固执，破坏力强。",
	"文昌貴人": "主天资聪明，又主逢凶化吉。命带文昌贵人，举止温顺文雅，好学上进，男命注重内涵，女命颇重仪表，乃才德兼备之人。其人天生具有文学艺术才华，能兼顾道德、人情、义理，谈吐不俗，仪表文雅。命格高者，能成为大文学家、艺术家、学者；命格低者，也是多才多艺。",
	"天廚貴人": "其人有食禄之贵，如得国家重用定有优厚的俸禄，并且能学业有成，工作晋升。人命逢之，如不逢刑冲克破，一生不愁吃穿，食禄不虞匮乏。",
	"四廢":   "主身弱多病，做事无成，有始无终。如不遇生扶，又受克害，凶煞制者，主伤残，官司口舌，甚至牢狱之灾，或为僧道之人。",
	"學堂":   "主才華出眾，聰明俊秀。一生中發展性不錯，會讀書，讀書可有成就。 適合公職或教師之職。",
	"金輿":   "得金輿男女命可得良緣或得配偶是有財富者。可蔭六親，可得富貴。人際關係良好，男女長相柔和，親切，言行慢條斯理，生活快樂，親情濃厚。",
	"祿神":   "祿神入命之人，行事較積極也較率性。凡事幹勁十足，行事積極，一生得安逸，財源可順利。祿最喜逢驛馬，祿馬交馳，財源廣進暢通。",
	"文昌":   "入命主可逢凶可化吉，為人聰明靈敏，會變巧。文筆佳有才氣，氣質溫文，重內涵，也注重儀表。",
	"沐浴":   "也是桃花的一種，一生中比較會有男女糾紛之事產生。",
	"紅艷":   "全名為「桃花紅艷煞」。主人見人愛，多情多慾，具有風流浪漫氣質，男女相貌俊秀美麗，結婚者易有婚外情。",
	"飛刃":   "身體易有刑傷，易有血光，易破財，較會有交通意外事故，行為要注意。",
	"墓庫":   "個性內向固執，凡事較會隱藏，才華深藏不露，喜追根究底，做事像是不積極，但具潛能，遇事會有較多的疑問，常有精神不振現象，體質弱。",
	"羊刃":   "屬刀傷。其性剛強勇猛。入命，主個性積極躁進，容易感情用事，行事激烈，較不利六親，做事敢做敢當。適武職、軍警、外科醫師、技術人員之工作。",
	"流霞":   "男命主口舌是非，車禍、災禍、路中或刀下亡。女命主產厄，流產、生產開刀、產厄亡。",
	"天医":   "為掌管疾病之星。天醫入命。有機會可做良醫。學習醫術比平常人更有天份，做事容易事半功倍。亦適合學習玄學、五術、哲學、心理等學科。",
	"天德":   "天德入命，一切皆能逢凶化吉，一生少凶險。能得福氣，常能絕處逢生，會有貴人來幫助。天德為福德之星，品性仁慈，好善喔 !",
	"月德":   "月德入命，一切皆能逢凶化吉，去災禍而能招吉祥，人命得天月二德，主心中慈善，道德涵養高尚，凡事得人心，也為福壽之命。 ",
	"月破":   "乃虛耗之神也。入命主不利身體，易破財，事業易不振。居住之環境常變化，也常會與人有口舌不和之事發生。",
	"六厄":   "一生易遭逢困難，主一生前途較艱辛。如逢吉神同柱相助則會轉吉。",
	"災煞":   "主血光及水火之災難也，也主較會墮落，若有福神相助。多主威武有權。",
	"劫煞":   "命帶劫煞者，個性較急躁，也主一生多是非，破財不順，易患耳、鼻、咽喉、小腸之疾病。若為喜用。其人具競爭力，行事有擔當魄力。若為忌神，主獨裁性暴，易遭意外之災。歲運逢之，主破財，遭朋友之連累，被敲詐等不如意之事。",
	"伏吟":   "書上說：「反吟伏吟、淚吟吟。」( 反吟又叫沖 ) ，伏吟代表痛苦，憂鬱，破財，也主傷害要小心，如能將其合化最好。",
	"元辰":   "歲運逢之，凡事顛三倒四，生活較不得寧息，可能會有外難，凡事較不能結合之意，主其人面有顴骨，鼻低口大，聲音沉濁。",
	"血刃":   "命帶血刃者，生平易發生意外受傷、流血、手術等事，對於尖銳物如：凶刀、凶器等物，比平常人敏感，接觸時要小心。",
	"桃花":   "桃花之好壞，須配合命局喜用五行及吉凶神煞來參考論斷。 在年、月柱為「內桃花」。 在日、時柱為「外桃花」。 若四柱地支有三柱以上的子、午、卯、酉者叫「遍野桃花」。 命帶桃花者，男女多半容貌俊或秀麗，人緣極佳，為人多情，有時較貪慾。 心情不定，行為不檢，好花錢，智慧。",
	"天狗":   "為刑傷之星，入命身體易有損傷、破相，如流年逢之，防損傷、交通事故、小人暗害、疾病產生。 ( 宜常祈福玄天上帝、文昌帝君 ) 自可平安、順利。",
	"白虎":   "為刑傷之星，入命較易犯官司是非，且身體易刑傷、有血光，易破財，較會有交通意外事故。女命帶此星，個性較剛烈，行事果斷。 ( 宜常祈福南斗星君 ) 自可平安、順利。",
	"五鬼":   "主一生中易犯官司、訴訟，易受小人陷害或連累，遇事不講實話或常會說謊，甚至口出狂言，處事要小心為要。 ( 宜常祈福玄天上帝、文昌帝君 ) 自可平安、順利。",
	"勾絞":   "命帶此星是比較容易惹麻煩之星，一生中多是非麻煩，宜防口舌是非、訴訟之事發生，少管閒事就對啦 !",
	"歲破":   "命帶此星主破財，虛耗及多波折，故一生中，如遇歲破之年，若不小心注意防範，會有較大的破財機會，同時亦主一切不安，身體欠佳。 ( 宜常祈福三寶佛、安太歲君 ) 自可平安、順利。",
	"喪門":   "命帶此星，平常儘量避免送葬、探病，以免因沖犯煞氣，導致生病而轉為壞運。流年不佳時凡事會比較不順或有損財之事發生。 ( 宜常祈福三官大帝、觀音大士 ) 自可平安、順利。",
	"天喜":   "命帶此星，男命相貌堂堂，女命面貌清秀，天天都是好日子，可能有桃花運喔 !",
	"紅鸞":   "命帶此星，一般都是男俊女美，在求財方面較能順心如意，也主常會接觸喜事。「紅鸞星動、滿面春風。喜氣盈門、男女喜慶。」是好星。",
	"金匱":   "命帶此星，有機會成為各行業中之傑出人才，在理財及管理之能力方面強，可當幹部主管職位。",
	"福德":   "福德臨入命，諸事大吉，一生可發財，諸事稱心如意。主福祿，具有增強吉祥之作用。 ( 宜常祈福財寶天王、福德正神 ) 自可平安、順利。",
	"龍德":   "命逢龍德貴人之人，萬事稱心如意，凡事平順，凡事可逢凶化吉。",
}

// 三奇
var SHEN_SHA_SAN_QI_STR = [3]string{"天上三奇", "地上三奇", "人中三奇"}
var SHEN_SHA_SAN_QI_LIST = [3][3]string{
	// 年干 月干 日干
	{"甲", "戊", "庚"},
	{"乙", "丙", "丁"},
	{"壬", "癸", "辛"},
}

// 魁罡,看柱,年月日时干支只要出现此组合则是魁罡
var SHEN_SHA_KUI_GANG_STR = "魁罡"

//var SHEN_SHA_KUI_GANG_LIST = [4][2]int{{4, 10}, {5, 10}, {6, 4}, {8, 4}} // 戊戌、庚戌、庚辰、壬辰
var SHEN_SHA_KUI_GANG_LIST = [4]string{"戊戌", "庚戌", "庚辰", "壬辰"}

// 年月日时支查日干表
var SHEN_SHA_DAY_GAN_ALL_ZHI_STR = []string{
	"天乙貴人", "太極貴人", "文昌貴人", "福興貴人", "天廚貴人", "金興", "干祿", "羊刃", "紅艷", "飛刃", "墓庫", "流霞", "學堂", "祿神", "沐浴",
}
var SHEN_SHA_DAY_GAN_ALL_ZHI_LIST = [][10]string{
	{"丑未", "子申", "酉亥", "酉亥", "丑未", "子申", "丑未", "寅午", "卯巳", "卯巳"},
	{"子午", "子午", "酉卯", "酉卯", "丑辰未戌", "丑辰未戌", "寅亥", "寅亥", "巳申", "巳申"},
	{"巳", "午", "申", "酉", "申", "酉", "亥", "子", "寅", "卯"},
	{"寅", "丑亥", "戌子", "酉", "申", "未", "午", "巳", "辰", "丑"},
	{"巳", "午", "巳", "午", "申", "酉", "亥", "子", "寅", "卯"},
	{"辰", "巳", "未", "申", "未", "申", "戌", "亥", "丑", "寅"},
	{"寅", "卯", "巳", "午", "巳", "午", "申", "酉", "亥", "子"},
	{"卯", "寅辰", "午", "巳午", "午", "巳午", "酉", "申戌", "子", "亥丑"},
	{"午", "午", "寅", "未", "辰", "辰", "戌", "酉", "子", "申"},
	/*飛刃*/ {"酉", "戌", "子", "丑", "子", "丑", "卯", "辰", "午", "未"},
	/*墓庫*/ {"未", "戌", "戌", "丑", "戌", "丑", "丑", "辰", "辰", "未"},
	/*流霞*/ {"酉", "戌", "未", "申", "巳", "午", "辰", "卯", "寅", "亥"},
	/*學堂*/ {"亥", "午", "寅", "酉", "寅", "酉", "巳", "子", "申", "卯"},
	/*祿神*/ {"寅", "卯", "巳", "午", "巳", "午", "申", "酉", "亥", "子"},
	/*沐浴*/ {"子", "巳", "卯", "申", "卯", "申", "午", "亥", "酉", "寅"},
}

// 月干、年日时干支查月支表
var SHEN_SHA_MONTH_ZHI_ALL_GANZHI_STR = []string{
	"天德贵人", "天德合", "月德貴人", "月德合", "天医", "天赦", "四廢",
}

// 月干、年日时干支查月支
var SHEN_SHA_MONTH_ZHI_ALL_GANZHI_LIST = [][12]string{
	{"巳", "庚", "丁", "申", "壬", "辛", "亥", "甲", "癸", "寅", "丙", "乙"},
	{"申", "乙", "壬", "巳", "丁", "丙", "寅", "巳", "戊", "亥", "辛", "庚"},
	{"壬", "庚", "丙", "甲", "壬", "庚", "丙", "甲", "壬", "庚", "丙", "甲"},
	{"丁", "乙", "辛", "己", "丁", "乙", "辛", "己", "丁", "乙", "辛", "己"},
	{"亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌"},
	{"甲子", "甲子", "戊寅", "戊寅", "戊寅", "甲午", "甲午", "甲午", "戊申", "戊申", "戊申", "甲子"},
	{"丙午丁巳", "丙午丁巳", "庚申辛酉", "庚申辛酉", "庚申辛酉", "壬子癸亥", "壬子癸亥", "壬子癸亥", "甲寅乙卯", "甲寅乙卯", "甲寅乙卯", "丙午丁巳"},
}

// 年月日时干支查日支表
var SHEN_SHA_DAY_ZHI_ALL_ZHI_STR = []string{
	"驛馬", "華蓋", "將星", "亡神", "劫煞", "孤辰", "寡宿", "咸池", "天喜", "紅鸞",
}
var SHEN_SHA_DAY_ZHI_ALL_ZHI_LIST = [][12]string{
	/*驛馬*/ {"寅", "亥", "申", "巳", "寅", "亥", "申", "巳", "寅", "亥", "申", "巳"},
	/*華蓋*/ {"辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未"},
	/*將星*/ {"子", "酉", "午", "卯", "子", "酉", "午", "卯", "子", "酉", "午", "卯"},
	/*亡神*/ {"亥", "申", "巳", "寅", "亥", "申", "巳", "寅", "亥", "申", "巳", "寅"},
	/*劫煞*/ {"巳", "寅", "亥", "申", "巳", "寅", "亥", "申", "巳", "寅", "亥", "申"},
	/*孤辰*/ {"寅", "寅", "巳", "巳", "巳", "申", "申", "申", "亥", "亥", "亥", "寅"},
	/*寡宿*/ {"戌", "戌", "丑", "丑", "丑", "辰", "辰", "辰", "未", "未", "未", "戌"},
	/*咸池*/ {"酉", "午", "卯", "子", "酉", "午", "卯", "子", "酉", "午", "卯", "子"},
	/*天喜*/ {"酉", "申", "未", "午", "酉", "申", "未", "午", "酉", "申", "未", "午"},
	/*紅鸞*/ {"卯", "寅", "丑", "子", "亥", "戌", "酉", "申", "未", "午", "巳", "辰"},
}

// TODO 日干查日支表
var SHEN_SHA_DAY_ZHI_DAY_GAN_STR = []string{
	"金神", "魁罡", "十惡大敗", "陰差陽錯", "八專", "九醜",
}
var SHEN_SHA_DAY_ZHI_DAY_GAN_LIST = [][12]string{
	/*金神*/ {"", "乙", "", "", "", "", "己", "", "", "", "癸", ""},
	/*魁罡*/ {"", "", "", "", "壬庚", "", "", "", "", "", "庚戊", ""},
	/*十惡大敗*/ {"", "己", "", "", "甲庚", "乙辛", "", "", "丙壬", "", "戊", "丁癸"},
	/*陰差陽錯*/ {"丙", "丁", "戊", "辛", "壬", "癸", "丙", "丁", "戊", "辛", "壬", "癸"},
	/*八專*/ {"", "癸", "甲", "乙", "", "", "", "己丁", "庚", "辛", "戊", ""},
	/*九醜*/ {"戊壬", "", "", "己辛", "", "", "戊壬", "", "", "丁己辛", "", ""},
}

// 金神, 1)查日柱 2)查时柱,日干必须是甲或己
var SHEN_SHA_JIN_SHEN_STR = "金神"
var SHEN_SHA_JIN_SHEN_LIST = [12]string{
	/*金神*/ "", "乙", "", "", "", "己", "", "", "", "癸", "", "",
}

// 年支查日月时支
var SHEN_SHA_YEAR_ZHI_ALL_ZHI_STR = []string{
	"驛馬", "華蓋", "將星", "亡神", "劫煞", "災煞", "孤辰", "寡宿", "咸池", "天喜", "紅鸞", "丧门", "白虎", "吊客", "病符", "血刃", "金匱", "天喜", "德", "福德", "勾絞", "五鬼", "破碎", "大耗", "白虎", "天狗", "桃花", "元辰", "元辰",
}
var SHEN_SHA_YEAR_ZHI_ALL_ZHI_LIST = [][12]string{
	/*驛馬*/ {"寅", "亥", "申", "巳", "寅", "亥", "申", "巳", "寅", "亥", "申", "巳"},
	/*華蓋*/ {"辰", "丑", "戌", "未", "辰", "丑", "戌", "未", "辰", "丑", "戌", "未"},
	/*將星*/ {"子", "酉", "午", "卯", "子", "酉", "午", "卯", "子", "酉", "午", "卯"},
	/*亡神*/ {"亥", "申", "巳", "寅", "亥", "申", "巳", "寅", "亥", "申", "巳", "寅"},
	/*劫煞*/ {"巳", "寅", "亥", "申", "巳", "寅", "亥", "申", "巳", "寅", "亥", "申"},
	/*災煞*/ {"午", "卯", "子", "酉", "午", "卯", "子", "酉", "午", "卯", "子", "酉"},
	/*孤辰*/ {"寅", "寅", "巳", "巳", "巳", "申", "申", "申", "亥", "亥", "亥", "寅"},
	/*寡宿*/ {"戌", "戌", "丑", "丑", "丑", "辰", "辰", "辰", "未", "未", "未", "戌"},
	/*咸池*/ {"酉", "午", "卯", "子", "酉", "午", "卯", "子", "酉", "午", "卯", "子"},
	/*天喜*/ {"酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑", "子", "亥", "戌"},
	/*紅鸞*/ {"卯", "寅", "丑", "子", "亥", "戌", "酉", "申", "未", "午", "巳", "辰"},
	/*喪門*/ {"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"},
	/*白虎*/ {"申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未"},
	/*吊客*/ {"戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉"},
	/*病符*/ {"亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌"},
	/*血刃*/ {"戌", "酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑", "子", "亥"},
	/*金匱*/ {"子", "酉", "午", "卯", "子", "酉", "午", "卯", "子", "酉", "午", "卯"},
	/*天喜*/ {"酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑", "子", "亥", "戌"},
	/*龍德*/ {"未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午"},
	/*福德*/ {"酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申"},
	/*勾絞*/ {"卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅"},
	/*五鬼*/ {"辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯"},
	/*破碎*/ {"午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳"},
	/*大耗*/ {"午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳"},
	/*白虎*/ {"申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未"},
	/*天狗*/ {"戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉"},
	/*桃花*/ {"酉", "午", "卯", "子", "酉", "午", "卯", "子", "酉", "午", "卯", "子"},
	/*+-元辰*/ {"未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午"},
	/*-+元辰*/ {"巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰"},
}

// 六秀
var SHEN_SHA_LIU_XIU_STR = "六秀日"
var SHEN_SHA_LIU_XIU_LIST = []string{
	"丙午", "丁末", "戊子", "戊午", "己丑", "己未",
}

// 计算神煞
func CalcShenSha(pSiZhu *TSiZhu) TShenSha {
	var shensha TShenSha
	// TODO 年月日干判断是否三奇
	for i, v := range SHEN_SHA_SAN_QI_LIST {
		if pSiZhu.YearZhu.Gan.ToString() == v[0] && pSiZhu.MonthZhu.Gan.ToString() == v[1] && pSiZhu.DayZhu.Gan.ToString() == v[2] {
			shensha.YearShenSha = append(shensha.YearShenSha, SHEN_SHA_SAN_QI_STR[i])
		}
	}
	// TODO 单独判断是否魁罡
	for i, v := range SHEN_SHA_KUI_GANG_LIST {
		if pSiZhu.YearZhu.Gan.ToString()+pSiZhu.YearZhu.Zhi.ToString() == v {
			switch i {
			case 0:
				shensha.YearShenSha = append(shensha.YearShenSha, SHEN_SHA_KUI_GANG_STR)
			case 1:
				shensha.MonthShenSha = append(shensha.MonthShenSha, SHEN_SHA_KUI_GANG_STR)
			case 2:
				shensha.DayShenSha = append(shensha.DayShenSha, SHEN_SHA_KUI_GANG_STR)
			case 3:
				shensha.HourShenSha = append(shensha.HourShenSha, SHEN_SHA_KUI_GANG_STR)
			}
		}
	}

	// TODO 年月日时支查日干表
	//nYearGanValue := pSiZhu.YearZhu.Gan.Value
	nMonthGanValue := pSiZhu.MonthZhu.Gan.Value
	nDayGanValue := pSiZhu.DayZhu.Gan.Value
	//nHourGanValue := pSiZhu.HourZhu.Gan.Value

	nYearZhiValue := pSiZhu.YearZhu.Zhi.Value
	nMonthZhiValue := pSiZhu.MonthZhu.Zhi.Value
	nDayZhiValue := pSiZhu.DayZhu.Zhi.Value
	nHourZhiValue := pSiZhu.HourZhu.Zhi.Value
	for i, v := range SHEN_SHA_DAY_GAN_ALL_ZHI_LIST {
		// 年支
		if strings.Contains(v[nDayGanValue], pSiZhu.YearZhu.Zhi.ToString()) {
			shensha.YearShenSha = append(shensha.YearShenSha, SHEN_SHA_DAY_GAN_ALL_ZHI_STR[i])
		}
		// 月支
		if strings.Contains(v[nDayGanValue], pSiZhu.MonthZhu.Zhi.ToString()) {
			shensha.MonthShenSha = append(shensha.MonthShenSha, SHEN_SHA_DAY_GAN_ALL_ZHI_STR[i])
		}
		// 日支
		if strings.Contains(v[nDayGanValue], pSiZhu.DayZhu.Zhi.ToString()) {
			shensha.DayShenSha = append(shensha.DayShenSha, SHEN_SHA_DAY_GAN_ALL_ZHI_STR[i])
		}
		// 时支
		if strings.Contains(v[nDayGanValue], pSiZhu.HourZhu.Zhi.ToString()) {
			shensha.HourShenSha = append(shensha.HourShenSha, SHEN_SHA_DAY_GAN_ALL_ZHI_STR[i])
		}
	}

	// TODO 月干、年日时支查月支表
	for i, v := range SHEN_SHA_MONTH_ZHI_ALL_GANZHI_LIST {
		// 月干
		if strings.Contains(v[nMonthGanValue], pSiZhu.MonthZhu.Gan.ToString()) {
			shensha.MonthShenSha = append(shensha.MonthShenSha, SHEN_SHA_MONTH_ZHI_ALL_GANZHI_STR[i])
		}
		// 年支
		if strings.Contains(v[nMonthZhiValue], pSiZhu.YearZhu.Zhi.ToString()) {
			shensha.YearShenSha = append(shensha.YearShenSha, SHEN_SHA_MONTH_ZHI_ALL_GANZHI_STR[i])
		} // 日支
		if strings.Contains(v[nMonthZhiValue], pSiZhu.DayZhu.Zhi.ToString()) {
			shensha.DayShenSha = append(shensha.DayShenSha, SHEN_SHA_MONTH_ZHI_ALL_GANZHI_STR[i])
		} // 时支
		if strings.Contains(v[nMonthZhiValue], pSiZhu.HourZhu.Zhi.ToString()) {
			shensha.HourShenSha = append(shensha.HourShenSha, SHEN_SHA_MONTH_ZHI_ALL_GANZHI_STR[i])
		}
	}

	// TODO 年月日时干支查日支表
	for i, v := range SHEN_SHA_DAY_ZHI_ALL_ZHI_LIST {
		// 年干
		if strings.Contains(v[nDayZhiValue], pSiZhu.YearZhu.Gan.ToString()) {
			shensha.YearShenSha = append(shensha.YearShenSha, SHEN_SHA_DAY_ZHI_ALL_ZHI_STR[i])
		}
		// 月干
		if strings.Contains(v[nDayZhiValue], pSiZhu.MonthZhu.Gan.ToString()) {
			shensha.MonthShenSha = append(shensha.MonthShenSha, SHEN_SHA_DAY_ZHI_ALL_ZHI_STR[i])
		}
		// 日干
		if strings.Contains(v[nDayZhiValue], pSiZhu.DayZhu.Gan.ToString()) {
			shensha.DayShenSha = append(shensha.DayShenSha, SHEN_SHA_DAY_ZHI_ALL_ZHI_STR[i])
		}
		// 时干
		if strings.Contains(v[nDayZhiValue], pSiZhu.HourZhu.Gan.ToString()) {
			shensha.HourShenSha = append(shensha.HourShenSha, SHEN_SHA_DAY_ZHI_ALL_ZHI_STR[i])
		}
		// 年支
		if strings.Contains(v[nDayZhiValue], pSiZhu.YearZhu.Zhi.ToString()) {
			shensha.YearShenSha = append(shensha.YearShenSha, SHEN_SHA_DAY_ZHI_ALL_ZHI_STR[i])
		}
		// 月支
		if strings.Contains(v[nDayZhiValue], pSiZhu.MonthZhu.Zhi.ToString()) {
			shensha.MonthShenSha = append(shensha.MonthShenSha, SHEN_SHA_DAY_ZHI_ALL_ZHI_STR[i])
		}
		// 时支
		if strings.Contains(v[nDayZhiValue], pSiZhu.HourZhu.Zhi.ToString()) {
			shensha.HourShenSha = append(shensha.HourShenSha, SHEN_SHA_DAY_ZHI_ALL_ZHI_STR[i])
		}
	}

	// TODO 日干查日支表
	for i, v := range SHEN_SHA_DAY_ZHI_DAY_GAN_LIST {
		if strings.Contains(v[nDayZhiValue], pSiZhu.DayZhu.Gan.ToString()) {
			shensha.DayShenSha = append(shensha.DayShenSha, SHEN_SHA_DAY_ZHI_DAY_GAN_STR[i])
		}
	}

	// TODO 金神, 1)查日柱 2)查时柱,日干必须是甲或己
	if strings.Contains(SHEN_SHA_JIN_SHEN_LIST[nDayZhiValue], pSiZhu.DayZhu.Gan.ToString()) ||
		(strings.Contains(SHEN_SHA_JIN_SHEN_LIST[nHourZhiValue], pSiZhu.HourZhu.Gan.ToString()) &&
			(pSiZhu.DayZhu.Gan.ToString() == "甲" || pSiZhu.DayZhu.Gan.ToString() == "己")) {
		shensha.HourShenSha = append(shensha.HourShenSha, SHEN_SHA_JIN_SHEN_STR)
	}

	// TODO 年支查日月时支
	for i, v := range SHEN_SHA_YEAR_ZHI_ALL_ZHI_LIST {
		if strings.Contains(v[nYearZhiValue], pSiZhu.MonthZhu.Zhi.ToString()) {
			shensha.MonthShenSha = append(shensha.MonthShenSha, SHEN_SHA_YEAR_ZHI_ALL_ZHI_STR[i])
		}
		if strings.Contains(v[nYearZhiValue], pSiZhu.DayZhu.Zhi.ToString()) {
			shensha.DayShenSha = append(shensha.DayShenSha, SHEN_SHA_YEAR_ZHI_ALL_ZHI_STR[i])
		}
		if strings.Contains(v[nYearZhiValue], pSiZhu.HourZhu.Zhi.ToString()) {
			shensha.HourShenSha = append(shensha.HourShenSha, SHEN_SHA_YEAR_ZHI_ALL_ZHI_STR[i])
		}
	}

	// TODO 六秀日
	for _, v := range SHEN_SHA_LIU_XIU_LIST {
		if pSiZhu.DayZhu.Gan.ToString()+pSiZhu.DayZhu.Zhi.ToString() == v {
			shensha.DayShenSha = append(shensha.DayShenSha, SHEN_SHA_LIU_XIU_STR)
		}
	}

	return shensha
}
