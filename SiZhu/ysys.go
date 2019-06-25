package SiZhu

// 一生运势简批

var YSYS_MONTH_STR = map[string]string{
	"正月": `此月生人，前年四月受胎，立春节后出生。为人忠厚，富有侠义心，仁德待人，同情心深，牺牲自己，成人之美。但带有神经质。利官近贵，可富贵增荣。然大事小成，凡事仔细，若不失机，能招四方之财。幼年平常，中年运开，晚年荣富。无刑克之命。<br/>
	<BR>诗曰：相貌端正前缘，早年衣禄自安然。贵人接引鸿运路，夫妇团圆过百年。<br/>
	<BR>`,
	"二月": `此月生人，前年五月受胎，惊蛰节后出生。性情温良，为人诚实。出言无毒，善作阴德，诸事谦尊，成功如登梯，贪急则失败，切记随机行事，方万事安然。初限辛苦，中年发达，四十兴荣后，可终年利路享通，乃虎假半真之命。<br/>
	<BR>诗曰：平生良善自有持，衣禄增荣盛有余。钱财家业中年好，贵人提拔上云梯。　<br/>
	<BR>`,
	`三月`: `此月生人，前年六月受胎，清明节后出生。志气不凡，内心强固，交际巧妙，脑智明晰。宽宏大量，事多忍耐，忍耐不如奋斗，莫要错失良机。应修身谨慎，有重色情前程，为情色害良缘之虞。三十岁前逢盛运，财来财往，有虚元实；四十岁后自安然，凡事顺遂，不可焦急。福禄永在。&nbsp;<BR>诗曰：为人心情自宽怀，平生积得四方财，一旦时来当发福，犹如枯木遇春天。　<BR>`,
	`四月`: `此月生人，前年七月受胎，立夏节后出生。广交朋友，多才巧智。学艺能成功，奈欲望过大，心性难定。宜不懈奋斗，终保发达。义侠心强，肯牺牲自己，败强助弱，后逢贵提拔，声扬名震，卅七岁后发展，为头目之人。<br/>
	<BR>诗曰，一年喜乐一年忧，无须怨恨匆忧愁，最宜持济行方便，夫妇齐眉乐到头。　<br/>
	<BR>`,
	`五月`: `此月生，前年八月受胎，芒种节后出生，为人善良，性情温和伶俐。行而正道，自能成功，恐对事虎头蛇尾，无忍耐性，遇良机不能得。离祖成家。夫妇半途，婚迁为吉，卅一岁或卅五岁后，方能大得利益。<br/>
	<BR>诗曰：自出常遇见横财，上人接引笑颜开。田园产业家豪富，荣华富贵步金阶。<br/>
	<BR>`,
	`六月`: `此月生人，前年九月受胎，小暑节后出生。为人远达，心巧伶俐，思虑致密，艺术多能。具努力有坚定心者，终为大发达。若贪小利则失大败，重色情必破家庭。白手成家，难得祖业，初限难为，中年平顺未运富贵。<br/>
	<BR>诗曰：一生衣禄人安康，为人显达有文光。三春快东蓄家富，夫妻同居松柏长。<br/>
	<BR>`,
	`七月`: `此月生人，前年十月受胎，立秋节后出生。心地慈善，作风仔细，为人亲切，外刚内柔，意志坚固，作事始终。少年辛苦，初限破财。对子妇烦恼，小心阴人。中年后开泰，晚年家庭园满，财源大旺。<br/>
	<BR>诗曰：为人一生不须忧，少小定心有根由。家宅由园宜主管，方知福禄不待人。<br/>
	<BR>`,
	`八月`: `此月生人，前年十一月受胎，白露节后出生，文章显达，记忆机敏，四方多艺创造巧妙。但独立不宜，须合夥成事，正直无私。初限幸福，中年离乱，于不意中失败，晚年福禄济美。<br/>
	<BR>诗曰：为人端正貌堂堂，皆因前世性温良。今生宜多行善事，自然福禄寿绵长。<br/>
	<BR>`,
	`九月`: `此月生人，前年十二女受胎，寒露节后出生。智慧锐敏，招四方之财。恐怕聪明自误，有失仁和，宜养温柔之心，自然贵人扶持。常省过去，奋发向前，趁时乘利，自得权柄。至四十而大发，子孙兴隆，多多顺利，晚景幸福。<br/>
	<BR>诗曰：此人生后大得财，钱财用尽又送来。八字好星家豪富衣禄儿职称心怀。　<br/>
	<BR>`,
	`十月`: `此月生人，前年正月受胎，冬至节后出生。为人声音刚强，心灵艺巧，然男女多相克，夫妇难合睦，心情易变动，常为失败之因。初限不理想，晚景安逸。<br/>
	<BR>诗曰：为人生来庆半余，免得灾殃祸其身，更宜持济行善事，一生衣禄睦三春。<br/>
	<BR>`,
	`十一月`: `此月生人，前年二月受胎，大雪节后出生。伶俐却性急，近贵却多计较，易招障害。务得人和，作事努力，名位自得。初限难为，中年灾涉色情，晚运大好，享子孙福。<br/>
	<BR>诗曰：自觉早年成立家，生平衣禄有荣华，亲戚兄弟全无靠，交接好友胜其他。<br/>
	<BR>`,
	`腊月`: `此月生人，前年三月受胎，小寒节后出生。为人心直口快，兄弟难靠。出外风光好。衣禄有足余。广交朋友，爱管闲事；抱负甚大，志气恢宏，心贪易遭失败。初限有福，中年辛劳，晚景大吉。<br/>
	<BR>诗曰：初限勤劳受苦辛，自然未后下求人。好运来时禄至，夫妇团圆寿百春。<br/>
	<BR>`,
}

var YSYS_DAY_STR = map[string]string{
	"初一": "此日生人，福禄难全，财星拱照，受人引进，事业发达，大有良机，初年平常，中年运到，利路亨通，晚景荣幸，发福之命。",
	"初二": "此日生人，性格善良，与人和睦，身体健康，家族缘薄，离祖成家，青年辛苦，兄弟难靠，独立生汁，中年运天，财源广进，男主清奇，女主聪明，成立之命。",
	"初三": "此日生人，夫妻和睦，不能偕老，子息克乏，须修身布德，初年多乖，三十有庆，受人提拔，四十盛运，左作右中，环境良好，荣华之命。",
	"初四": "此日生人，为人多学，才知出众，少年不宜，中运财好，在家多是非，出外逢贵人，夫妻和顺，家庭圆满，活泼，快乐之命。",
	"初五": "末运大旺，与人亲睦，贵人提拔，发达成功，父兄无靠。白手成家。",
	"初六": "末运大旺，与人亲睦，贵人提拔，发达成功，父兄无靠。白手成家。",
	"初七": "此日生人，性格多变易动，事业浮沉末定，半生波澜风霜，卅五后来佳运，事事如意，女命福禄，守身平和，乃健全长寿之命。",
	"初八": "此日生人，性格伶俐，一生安中年成功，父母无缘，离祖成家，出外逢贵，乃富贵荣华之命。",
	"初九": "此日生人，身体健全，性格清朗，受人敬受，须事事勉励，勤俭行善，德被乡党，中年平顺，晚景千钟，福分无量，名利长存，慈悲富贵之命。",
	"初十": "此日生人，为人伶俐，忠诚待人，家族缘薄，离祖成家，缘和四海，少年辛苦，中年开发，晚年大兴，事业通达，艺术成功，安乐之命。",
	"十一": "此日生人富有智力，意志坚固，享有决断，至中年虽有横财，不能料事，空放幸运，宜要谨慎，财源循来，福分之命。",
	"十二": "此日生人，为人温柔，刻苦耐劳，善好勤俭，多积蓄物，少年不宜，中年大吉，将见名扬，福禄双至，晚年馀庆，家门隆兴，福禄之命。",
	"十三": "此日生人，金运可达，福禄有馀，遵守道德，受人敬爱，贵人提拔，命运通达，大有成功，获得幸福，福禄双收，女命富贵，金运之命。",
	"十四": "此日生人，环境良好，为人厚重，沉静不动，男人清秀，女人聪明，贵人得助，青年平常，努力前程，中年运开，事得顺调，晚年发达，厚分之命。",
	"十五": "此日生人，夫妻敬重，子孙刑克，强争好斗，破害前程，卅五之后方来馀庆，男者离租，他乡发展，女人克夫，必配硬命，平常之命。",
	"十六": "此日生人，为人聪明，艺术超群，琴棋达人，书诗出分。青年勤功，中年艺精，成功发达，祖业不宜，身闲心劳，忧闷之命。",
	"十七": "此日生人，为人聪明，智力平凡，忍耐力强，少年多障害，难关重重来，善理时时机，对兄弟情薄，故六亲无告，自力更生，中年大发，发达之命。",
	"十八": "此日生人，智能可畏，自作聪明，不容他人，性格过刚，与人不和，独立自好，父兄无缘，晚景大运，中年平平，普通之命。",
	"十九": "此日生人，名利双收，成功运强，为人出众，色情心重，桃唇小身有暗病，苦恼自叹，应防患未然，中年平平，未运福禄，荣华之命。",
	"二十": "此日生人，刑克上少成多劳，心身多烦，波澜重见，男人离祖，租业难富，亲朋无靠，山外得财，有贵人助，晚景大幸，昌盛之命。",
	"廿一": "此日生人，内助得力，衣禄和顺，受人提拔，待有金运，喜好投机，艺高胆大，不服他人，中年平平，事无称心，未境发达，晚福之命。",
	"廿二": "此日生人，一生聪明，情义或嘉，作享无虚，先难后易，少年多难，苦中得甘，廿五运到，良好前程，加添努力，晚景大兴，名利之命。",
	"廿三": "此日生人，心境常换，作事不定，多变多变，易与争斗勤励事业，矫正缺点，衣禄有馀，中年平平未运福到，荣华富贵，平安之命。",
	"廿四": "此日生人，为人伶俐，作事专业，有头有尾，对人亲切，四处友朋，受人敬爱，好积财宝，出外逢贵，刑克妻子，盛昌之命。",
	"廿五": "此日生人，为人忠实，大有器才，喜好公益，巧理家庭，专业经营，事业发达，多管他事，易生敌对，心性未定，可得内助，晚年发达，大旺之命。",
	"廿六": "此日生人，为人仁德，慈悲心重，受人爱慕，长上提拔，立身处世，先苦后甘，作事无虚，勤俭积蓄，金运满载，幸福无疆，福贵之命。",
	"廿七": "此日生人，为人巧奇，金运缘薄，多收多出，变动无常，居所未定，出洋成功，努力奋斗，前程有馀，身体强壮，勤俭励业，成功之命。",
	"廿八": "此日生人，青年薄运，幼年病多，独立意志，认真作事，中年运到，积蓄金钱，自得良缘，亲朋难靠，早婚刑克，晚婚平静，求得温和，金运之命。",
	"廿九": "此日生人，为人忠厚，肯作肯劳，重义信用，与人豪杰，慷慨待人，广结社会，积财有馀。初限平顺，中年运到，晚年馀庆，幸福之命。",
	"三十": "此日生人，性情刚果。不眼他人，造成怨城，受人攻击，与人和柔，幸运循来，少年薄运，多劳不乐。愁眉不展，至中年，运气一新，兴隆之命。",
}

var YSYS_HOUR_STR = map[string]string{
	"子": `（午夜十一点起上午一点止，为鼠时辰）
	<BR>性急刚富于勤俭，有谋欠勇，是非多端，父母得力，妻子相助，早年发运，白手成家。&nbsp;
	<BR>
	<BR>适业：艺术、政治、建筑、电气、属金水事业可。忌士类。&nbsp;
	<BR>
	<BR>凶年：十一岁，十八岁，三六岁，四十五岁，五十八岁，八十八岁寿终。&nbsp;
	<BR>
	<BR>子时头生，时头生人克母命，十成九败多进退，医术僧道终皆吉，诸事难招六亲冷。&nbsp;
	<BR>
	<BR>子时中生，时中生人无克破，一生作事硬占强，七倒八&nbsp;起后兴旺，离祖成家得清闲。
	<BR>
	<BR>子时未生，时未生人先克父，六亲无靠劳禄夫，有财无库宜积蓄，子息两硬逢过房。
	<BR>`,
	"丑": `（上午一点起上午三点止为牛时辰）&nbsp;
	<BR>家族缘薄，离乡成功，上官近贵，性情暴躁，二十运开，四五兴旺，晚年齐福。&nbsp;
	<BR>
	<BR>适业：商业、技师、教员、官吏、学者、饮食、加工、忌木类。&nbsp;
	<BR>
	<BR>凶年：十八岁，甘三岁，卅一岁，四十六岁，六十二岁寿终。
	<BR>
	<BR>丑时头生：时头生人无克破，一身衣禄多有馀，荣华富贵兼安闲，子孙繁荣得显见。&nbsp;
	<BR>
	<BR>丑时中生：时中生人克父命，为人诚实进田庄，夫妻和命于得力，老景财帛庆馀丰。&nbsp;
	<BR>
	<BR>丑时末生，时末生人先克母，为人孤独心中善，先苦后甘终得福，晚年交来贵人扶。　
	<BR>`,
	"寅": `（上午三点起上午五点止为虎时辰）&nbsp;
	<BR>六亲相克，多劳，受苦，难守祖业，流浪异乡，十八财害，四十发福，晚景逢贵。&nbsp;
	<BR>
	<BR>适业：医师、音乐师、美术家、艺人、流浪职业。忌类金。&nbsp;
	<BR>
	<BR>凶年：廿六岁，廿九岁、卅三岁、卅九岁、四十九岁、六十六岁寿终。&nbsp;
	<BR>
	<BR>寅时生：时头生人先克义，为人聪明兼伶俐，一生良禄自可求，时常变动心未休。&nbsp;
	<BR>
	<BR>寅时中生：时中生人父母全，为人作事有权柄，六亲有分衣禄全，离祖扬名进宅田。&nbsp;
	<BR>
	<BR>寅时末生：时末生人先克母，六亲无力诸事难，兄弟不和多苦劳，卅六岁运宜发财。　
	<BR>`,
	"卯": `（上午五点起上午七点止为兔时辰）&nbsp;
	<BR>福多劳少，父母难常，兄弟难靠，出外经营，利路亨通，夫妻相克，先难后易，立身成功。&nbsp;
	<BR>
	<BR>适业：机械、演戏、文学、美术、宗教家、职员。忌人类。&nbsp;
	<BR>
	<BR>凶年：十六岁、付岁、五十五岁、七十二岁寿终。&nbsp;
	<BR>
	<BR>卯时头生，时头生人先克母，为人作事无始终，兄弟少力六亲难，离祖过房最安宁。&nbsp;
	<BR>
	<BR>卯时中生：时中生人父母全，男子居宦女有福，一生富贵庆有馀，六亲兄皆如意。&nbsp;
	<BR>
	<BR>卯时未生：时末生人先克父，十成九败运不通，初限不达苦劳禄，老年枯木再生花。　
	<BR>`,
	"辰": `（上午七点起上午九点止为龙时辰）&nbsp;
	<BR>聪明伶俐，意志强固，目中元人，衣禄光辉，为人孤独，自信过重，宜之戒心。&nbsp;
	<BR>
	<BR>适业：实业家、政治家、中介人、教员、矿业。忌木类。凶年：十九岁、廿七岁、卅六岁。卅九岁、六十六岁寿终。&nbsp;
	<BR>
	<BR>辰时头生：时头生人父母在，六亲无力兄弟休，心慈心善多手艺，四十二岁后渐渐好。&nbsp;
	<BR>
	<BR>辰时中生：时中生人先克父，为人公道性子急，兄弟六亲情疏远，离祖成家贵人扶。&nbsp;
	<BR>
	<BR>辰时末生：时末生人母先卒，为人聪明财禄聚，兄弟有顾六亲旺，近官益夫福禄全。　
	<BR>`,
	"巳": `（上午九点起上午十一点止为蛇时辰）&nbsp;
	<BR>智能非几，自成业，六亲无缘，离祖成家；快乐待人，女人华难得良缘，饮酒成癖。
	<BR>
	<BR>适业：评论家、刺绣、矿业、加工业、忌水类。&nbsp;
	<BR>
	<BR>凶年：卅一岁、卅六岁。四十七岁、四十九岁、八十九岁寿终。&nbsp;
	<BR>
	<BR>己时头生：时头生人先克母，兄弟难靠六亲疏，早婚下宜害妻子，末限衣禄女益夫。&nbsp;
	<BR>
	<BR>己时中生：时中生人父母全，一生近贵财禄全，兄弟子息多和顺，宜作长者有福人。&nbsp;
	<BR>
	<BR>己时未生：时末生人先克父，作事一为一败，兄弟妻子六亲冷，早年奔波苦劳碌。　
	<BR>`,
	"午": `（上午十一点起下午一点止为马时晨）
	<BR>伶俐敏捷，不守祖业，自作聪明，女人娇娇，极端主人，浪费奢侈。
	<BR>
	<BR>适业：医师、护士、政治。明星、技艺、料理店、油业。忌金类。&nbsp;
	<BR>
	<BR>凶年：六岁、十二岁、什四岁、卅三岁、四十五岁、八十五岁寿终。&nbsp;
	<BR>
	<BR>午时头生：时头生人父母在，为人利害近贵人，兄弟六亲皆有靠，子息三五衣禄归。&nbsp;
	<BR>
	<BR>午时中生：时中生人先克父，福禄平平苦奔波，三十岁后渐渐好，先难后益得安然。&nbsp;
	<BR>
	<BR>午时末生：时末生人先克母，聪明伶俐兼性急，兄弟六亲难可靠，衣食不周多奔波。　
	<BR>`,
	"未": `（下午一点起下午三点为羊时辰）
	<BR>父母难依，刑克妻子，兄弟无靠，中限惊恐，女子多智，副业有为，多变易动，心身难定。&nbsp;
	<BR>
	<BR>适业：土木业、电气商，建筑业、木器商、酒类商。忌水类。&nbsp;
	<BR>
	<BR>凶年：十九岁、甘六岁、五十六岁、七十岁寿终。&nbsp;
	<BR>
	<BR>未时头生：时头生人父母全，一生安乐六亲荣，男有头目女益夫，知识轻重有操持。&nbsp;
	<BR>
	<BR>未时中生：时中生人先克父，为人性宽衣禄平，兄弟有情克首妻，六亲少靠子息难。&nbsp;
	<BR>
	<BR>未时末生：时末生人母先卒，三胜三败自安排，六亲少力衣禄平，兄弟和睦于息强。　
	<BR>`,
	"申": `（下午三点起下午五点止为猴时辰）&nbsp;
	<BR>败来败去，难守祖业，父母无靠，夫妻和谐，女人破婚，宜养操忌木类。&nbsp;
	<BR>
	<BR>凶年：十九岁、廿二岁、甘八岁、卅岁、四十二岁、五十四岁、七十二岁寿终。&nbsp;
	<BR>
	<BR>申时头生：时头生人父母全，为人聪明近贵人，能文能武志气大，六亲有禄进田园。&nbsp;
	<BR>
	<BR>申时中生，时中生人先克父，六亲不和兄弟疏，一忧一喜离祖吉，早妻刑克总多劳。&nbsp;
	<BR>
	<BR>申时末生：时末生人先克母，六亲兄弟多冷淡，早辛苦身多病，三十岁平四十益。　
	<BR>`,
	"酉": `（下午五点起下午七点止为鸡时辰）&nbsp;
	<BR>幼时辛苦，兄弟分离，父母无缘，乏子宜养，女人多情，善多机密，自尊心强，暴发好争。&nbsp;
	<BR>
	<BR>适业：仳学、采访、文艺、新业事业、加工业。忌土类。&nbsp;
	<BR>
	<BR>凶年：十九岁、甘五岁、卅二岁、四十九岁、七十八岁寿终。&nbsp;
	<BR>
	<BR>酉时头生：时头生人父母全，文武皆知近官贵，六亲有靠兄弟财禄胜前子孙稀。&nbsp;
	<BR>
	<BR>酉时中生：时中生人先克父，兄弟不利离祖居，夫妻刑克子又迟，早年不遂末限好。&nbsp;
	<BR>
	<BR>酉时末生：时末和人先克母，兄弟少力衣禄平，夫妻刑克子又逆，女人淫乱宜守德。
	<BR>`,
	"戌": `（下午七点起下午九点止为狗时辰） 
	<BR>
	<BR>富具阻力，独行好斗，财禄有进，歌乐一生，女人虚荣，性情暴燥，无忍耐性，不重金钱。 
	<BR>
	<BR>适业：诗人、作家、投机、五金、农林、米谷商、机械。忌火类。 
	<BR>
	<BR>凶年：十六岁、廿六岁、卅五岁、四十四岁，四十九岁、五十七岁、七十八岁寿终。 
	<BR>
	<BR>戌时头生：时头生人先克母，为人性急心慈悲，手足难靠六亲平，三十八发财有权柄。 
	<BR>
	<BR>戌时中生：时中生人父先卒，六亲兄少帮助，夫妻长子带刑克，四十二岁后好安排。 
	<BR>
　　<BR>戌时末生：时末生人父母全，文武皆性焦急，六亲兄弟多有靠，手艺精通夫妻和 
	<BR>`,
	"亥": `（下午九点起下午十一点止为猪时辰）&nbsp;
	<BR>意志坚强，沉着热心，不文际人，手艺特精，女人性刚，易努解，天生勤劳。财帛大旺。&nbsp;
	<BR>
	<BR>适业：外科医，僧侣、旅馆、支配人，艺术、古董、五金。忌人类。&nbsp;
	<BR>
	<BR>凶年：二一岁、廿六岁、卅六岁、卅九岁、四十九岁、五十六岁、七十八寿终。&nbsp;
	<BR>
	<BR>亥时头生：时头生人先克母，性宽手足情疏远，六亲少力初年苦，子息二三衣禄归。&nbsp;
	<BR>
	<BR>亥时中生：时中生人父母全，为人聪明性情急，亲戚兄弟多有分，女秉家权末限生。&nbsp;
	<BR>
	<BR>亥时末生：时末生人先克父，性燥心慈六亲疏，兄亲难为早年劳，男娶双妻女克夫。
	<BR>`,
}
