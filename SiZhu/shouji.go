package SiZhu

/*
1、先将手机号码最后四位抽出，作为测字基数。
2、除以80，减去所有整数位(若结果为0，则吉凶数字为80，省略下一步)
3、再将所得出的小数位乘以80，就得出一个吉凶数字。
4、查询手机号码吉凶预测表《八十一数吉凶佩带琥珀守护神八卦吉祥笔画划表》对照结果。
*/

var SHOUJI_STR = []string{
	"大展鸿图，信用得固，名利双收，可获成功|吉|主人性格类型：[做事喜好凭直觉型]，其具体表现为：有特强的第六灵感，性格率直无机心，深得朋辈爱戴。感情路上多采多姿。做事喜好凭个人直觉及预感做决定。",
	"根基不固，摇摇欲坠，一盛一衰，劳而无获|凶|主人性格类型：[独断独行/吸引人型]，其具体表现为：为人独断独行，事事自行作主解决，鲜有求助他人。而这份独立个性，正正就是吸引异性的特质。但其实心底里，却是渴望有人可让他/她依赖。",
	"根深蒂固，蒸蒸日上，如意吉祥，百事顺遂|吉|主人性格类型：[热情/善变梦想家型]，其具体表现为：对人热情无遮掩，时常梦想可以谈一场戏剧性恋爱，亲身体会个中悲欢离合的动人经历，是个大梦想家。但对于感情却易变卦。",
	"前途坎坷，苦难折磨，非有毅力，难望成功|凶|主人性格类型：[自我牺牲/性格被动型]，其具体表现为：惯于无条件付出，从不祈求有回报，有为了成全他人不惜牺牲自己的情操。但讲到本身的爱情观，却流于被动，往往因为内敛而错过大好姻缘。",
	"阴阳和合，生意兴隆，名利双收，后福重重|吉|主人性格类型：[要面包不要爱情型]，其具体表现为：责任心重，尤其对工作充满热诚，是个彻头彻尾工作狂。但往往因为过分专注职务，而忽略身边的家人及朋友，是个宁要面包不需要爱情的理性主义者。",
	"万宝集门，天降幸运，立志奋发，得成大功|吉|主人性格类型：[不善表达/疑心大型]，其具体表现为：在乎身边各人对自己的评价及观感，不善表达个人情感，是个先考虑别人感受，再作出相应配合的内敛一族。对于爱侣，经常存有怀疑之心。",
	"独营生意，和气吉祥，排除万难，必获成功|吉|主人性格类型：[大胆行事冲动派型]，其具体表现为：爱好追寻刺激，有不理后果大胆行事的倾向。崇尚自由奔放的恋爱，会拼尽全力爱一场，是就算明知无结果都在所不惜的冲动派。",
	"努力发达，贯彻志望，不忘进退，可望成功|吉|主人性格类型：[高度戒备难交心型]，其具体表现为：经常处于戒备状态，对任何事都没法放松处理，孤僻性情难以交朋结友。对于爱情，就更加会慎重处理。任何人必须经过你长期观察及通过连番考验，方会减除戒备与你交往。",
	"虽抱奇才，有才无命，独营无力，财力难望|凶|主人性格类型：[好奇心旺求知欲强型]，其具体表现为：好奇心极度旺盛，求知欲又强，有打烂沙盘问到笃的锲而不舍精神。此外，你天生有语言天分，学习外文比一般人更易掌握。",
	"乌云遮月，暗淡无光，空费心力，徒劳无功|凶|主人性格类型：[做事喜好凭直觉型]，其具体表现为：有特强的第六灵感，性格率直无机心，深得朋辈爱戴。感情路上多采多姿。做事喜好凭个人直觉及预感做决定。",
	"草木逢春，枝叶沾露，稳健着实，必得人望|吉|主人性格类型：[独断独行/吸引人型]，其具体表现为：为人独断独行，事事自行作主解决，鲜有求助他人。而这份独立个性，正正就是吸引异性的特质。但其实心底里，却是渴望有人可让他/她依赖。",
	"薄弱无力，孤立无援，外祥内苦，谋事难成|凶|主人性格类型：[热情/善变梦想家型]，其具体表现为：对人热情无遮掩，时常梦想可以谈一场戏剧性恋爱，亲身体会个中悲欢离合的动人经历，是个大梦想家。但对于感情却易变卦。",
	"天赋吉运，能得人望，善用智慧，必获成功|吉|主人性格类型：[自我牺牲/性格被动型]，其具体表现为：惯于无条件付出，从不祈求有回报，有为了成全他人不惜牺牲自己的情操。但讲到本身的爱情观，却流于被动，往往因为内敛而错过大好姻缘。",
	"忍得若难，必有后福，是成是败，惟靠紧毅|凶|主人性格类型：[要面包不要爱情型]，其具体表现为：责任心重，尤其对工作充满热诚，是个彻头彻尾工作狂。但往往因为过分专注职务，而忽略身边的家人及朋友，是个宁要面包不需要爱情的理性主义者。",
	"谦恭做事，外得人和，大事成就，一门兴隆|吉|主人性格类型：[不善表达/疑心大型]，其具体表现为：在乎身边各人对自己的评价及观感，不善表达个人情感，是个先考虑别人感受，再作出相应配合的内敛一族。对于爱侣，经常存有怀疑之心。",
	"能获众望，成就大业，名利双收，盟主四方|吉|主人性格类型：[大胆行事冲动派型]，其具体表现为：爱好追寻刺激，有不理后果大胆行事的倾向。崇尚自由奔放的恋爱，会拼尽全力爱一场，是就算明知无结果都在所不惜的冲动派。",
	"排除万难，有贵人助，把握时机，可得成功|吉|主人性格类型：[高度戒备难交心型]，其具体表现为：经常处于戒备状态，对任何事都没法放松处理，孤僻性情难以交朋结友。对于爱情，就更加会慎重处理。任何人必须经过你长期观察及通过连番考验，方会减除戒备与你交往。",
	"经商做事，顺利昌隆，如能慎始，百事亨通|吉|主人性格类型：[好奇心旺求知欲强型]，其具体表现为：好奇心极度旺盛，求知欲又强，有打烂沙盘问到笃的锲而不舍精神。此外，你天生有语言天分，学习外文比一般人更易掌握。",
	"成功虽早，慎防亏空，内外不合，障碍重重|凶|主人性格类型：[做事喜好凭直觉型]，其具体表现为：有特强的第六灵感，性格率直无机心，深得朋辈爱戴。感情路上多采多姿。做事喜好凭个人直觉及预感做决定。",
	"智商志大，历尽艰难，焦心忧劳，进得两难|凶|主人性格类型：[独断独行/吸引人型]，其具体表现为：为人独断独行，事事自行作主解决，鲜有求助他人。而这份独立个性，正正就是吸引异性的特质。但其实心底里，却是渴望有人可让他/她依赖。",
	"先历困苦，后得幸福，霜雪梅花，春来怒放|吉|主人性格类型：[热情/善变梦想家型]，其具体表现为：对人热情无遮掩，时常梦想可以谈一场戏剧性恋爱，亲身体会个中悲欢离合的动人经历，是个大梦想家。但对于感情却易变卦。",
	"秋草逢霜，怀才不遇，忧愁怨苦，事不如意|凶|主人性格类型：[自我牺牲/性格被动型]，其具体表现为：惯于无条件付出，从不祈求有回报，有为了成全他人不惜牺牲自己的情操。但讲到本身的爱情观，却流于被动，往往因为内敛而错过大好姻缘。",
	"旭日升天，名显四方，渐次进展，终成大业|吉|主人性格类型：[要面包不要爱情型]，其具体表现为：责任心重，尤其对工作充满热诚，是个彻头彻尾工作狂。但往往因为过分专注职务，而忽略身边的家人及朋友，是个宁要面包不需要爱情的理性主义者。",
	"绵绣前程，须靠自力，多用智谋，能奏大功|吉|主人性格类型：[不善表达/疑心大型]，其具体表现为：在乎身边各人对自己的评价及观感，不善表达个人情感，是个先考虑别人感受，再作出相应配合的内敛一族。对于爱侣，经常存有怀疑之心。",
	"天时地利，只欠人和，讲信修睦，即可成功|吉|主人性格类型：[大胆行事冲动派型]，其具体表现为：爱好追寻刺激，有不理后果大胆行事的倾向。崇尚自由奔放的恋爱，会拼尽全力爱一场，是就算明知无结果都在所不惜的冲动派。",
	"波澜起伏，千变万化，凌架万难，必可成功|凶|主人性格类型：[好奇心旺求知欲强型]，其具体表现为：好奇心极度旺盛，求知欲又强，有打烂沙盘问到笃的锲而不舍精神。此外，你天生有语言天分，学习外文比一般人更易掌握。",
	"一成一败，一盛一衰，惟靠谨慎，可守成功|凶带吉|主人性格类型：[做事喜好凭直觉型]，其具体表现为：有特强的第六灵感，性格率直无机心，深得朋辈爱戴。感情路上多采多姿。做事喜好凭个人直觉及预感做决定。",
	"鱼临旱地，难逃恶运，此数大凶，不如更名|凶|主人性格类型：[好奇心旺求知欲强型]，其具体表现为：好奇心极度旺盛，求知欲又强，有打烂沙盘问到笃的锲而不舍精神。此外，你天生有语言天分，学习外文比一般人更易掌握。",
	"如龙得云，青云直上，智谋奋进，才略奏功|吉|主人性格类型：[做事喜好凭直觉型]，其具体表现为：有特强的第六灵感，性格率直无机心，深得朋辈爱戴。感情路上多采多姿。做事喜好凭个人直觉及预感做决定。",
	"吉凶参半，得失相伴，投机取巧，吉凶无常|凶|主人性格类型：[独断独行/吸引人型]，其具体表现为：为人独断独行，事事自行作主解决，鲜有求助他人。而这份独立个性，正正就是吸引异性的特质。但其实心底里，却是渴望有人可让他/她依赖。",
	"此数大吉，名利双收，渐进向上，大业成就|吉|主人性格类型：[热情/善变梦想家型]，其具体表现为：对人热情无遮掩，时常梦想可以谈一场戏剧性恋爱，亲身体会个中悲欢离合的动人经历，是个大梦想家。但对于感情却易变卦。",
	"池中之龙，风云际会，一跃上天，成功可望|吉|主人性格类型：[自我牺牲/性格被动型]，其具体表现为：惯于无条件付出，从不祈求有回报，有为了成全他人不惜牺牲自己的情操。但讲到本身的爱情观，却流于被动，往往因为内敛而错过大好姻缘。",
	"意气用事，人和必失，如能慎始，必可昌隆|吉|主人性格类型：[要面包不要爱情型]，其具体表现为：责任心重，尤其对工作充满热诚，是个彻头彻尾工作狂。但往往因为过分专注职务，而忽略身边的家人及朋友，是个宁要面包不需要爱情的理性主义者。",
	"灾难不绝，难望成功，此数大凶，不如更名|凶|主人性格类型：[不善表达/疑心大型]，其具体表现为：在乎身边各人对自己的评价及观感，不善表达个人情感，是个先考虑别人感受，再作出相应配合的内敛一族。对于爱侣，经常存有怀疑之心。",
	"中吉之数，进退保守，生意安稳，成就普通|吉|主人性格类型：[大胆行事冲动派型]，其具体表现为：爱好追寻刺激，有不理后果大胆行事的倾向。崇尚自由奔放的恋爱，会拼尽全力爱一场，是就算明知无结果都在所不惜的冲动派。",
	"波澜得叠，常陷穷困，动不如静，有才无命|凶|主人性格类型：[好奇心旺求知欲强型]，其具体表现为：好奇心极度旺盛，求知欲又强，有打烂沙盘问到笃的锲而不舍精神。此外，你天生有语言天分，学习外文比一般人更易掌握。",
	"逢凶化吉，吉人天相，风调雨顺，生意兴隆|吉|主人性格类型：[做事喜好凭直觉型]，其具体表现为：有特强的第六灵感，性格率直无机心，深得朋辈爱戴。感情路上多采多姿。做事喜好凭个人直觉及预感做决定。",
	"名虽可得，利则难获，艺界发展，可望成功|凶带吉|主人性格类型：[独断独行/吸引人型]，其具体表现为：为人独断独行，事事自行作主解决，鲜有求助他人。而这份独立个性，正正就是吸引异性的特质。但其实心底里，却是渴望有人可让他/她依赖。",
	"云开见月，虽有劳碌，光明坦途，指日可望|吉|主人性格类型：[热情/善变梦想家型]，其具体表现为：对人热情无遮掩，时常梦想可以谈一场戏剧性恋爱，亲身体会个中悲欢离合的动人经历，是个大梦想家。但对于感情却易变卦。",
	"一成一衰，沉浮不定，知难而退，自获天佑|吉带凶|主人性格类型：[自我牺牲/性格被动型]，其具体表现为：惯于无条件付出，从不祈求有回报，有为了成全他人不惜牺牲自己的情操。但讲到本身的爱情观，却流于被动，往往因为内敛而错过大好姻缘。",
	"天赋吉运，德望兼备，继续努力，前途无限|吉|主人性格类型：[要面包不要爱情型]，其具体表现为：责任心重，尤其对工作充满热诚，是个彻头彻尾工作狂。但往往因为过分专注职务，而忽略身边的家人及朋友，是个宁要面包不需要爱情的理性主义者。",
	"事业不专，十九不成，专心进取，可望成功|吉带凶|主人性格类型：[不善表达/疑心大型]，其具体表现为：在乎身边各人对自己的评价及观感，不善表达个人情感，是个先考虑别人感受，再作出相应配合的内敛一族。对于爱侣，经常存有怀疑之心。",
	"雨夜之花，外祥内苦，忍耐自重，转凶为吉|吉带凶|主人性格类型：[大胆行事冲动派型]，其具体表现为：爱好追寻刺激，有不理后果大胆行事的倾向。崇尚自由奔放的恋爱，会拼尽全力爱一场，是就算明知无结果都在所不惜的冲动派。",
	"虽用心计，事难遂愿，贪功冒进，必招失败|凶|主人性格类型：[高度戒备难交心型]，其具体表现为：经常处于戒备状态，对任何事都没法放松处理，孤僻性情难以交朋结友。对于爱情，就更加会慎重处理。任何人必须经过你长期观察及通过连番考验，方会减除戒备与你交往。",
	"杨柳遇春，绿叶发枝，冲破难关，一举成名|吉|主人性格类型：[热情/善变梦想家型]，其具体表现为：对人热情无遮掩，时常梦想可以谈一场戏剧性恋爱，亲身体会个中悲欢离合的动人经历，是个大梦想家。但对于感情却易变卦。",
	"坎坷不平，艰难重重，若无耐心，难望有成|凶|主人性格类型：[自我牺牲/性格被动型]，其具体表现为：惯于无条件付出，从不祈求有回报，有为了成全他人不惜牺牲自己的情操。但讲到本身的爱情观，却流于被动，往往因为内敛而错过大好姻缘。",
	"有贵人助，可成大业，虽遇不幸，浮沉不定|吉|主人性格类型：[要面包不要爱情型]，其具体表现为：责任心重，尤其对工作充满热诚，是个彻头彻尾工作狂。但往往因为过分专注职务，而忽略身边的家人及朋友，是个宁要面包不需要爱情的理性主义者。",
	"美化丰实，鹤立鸡群，名利俱全，繁荣富贵|吉|主人性格类型：[不善表达/疑心大型]，其具体表现为：在乎身边各人对自己的评价及观感，不善表达个人情感，是个先考虑别人感受，再作出相应配合的内敛一族。对于爱侣，经常存有怀疑之心。",
	"遇吉则吉，遇凶则凶，惟靠谨慎，逢凶化吉|凶|主人性格类型：[大胆行事冲动派型]，其具体表现为：爱好追寻刺激，有不理后果大胆行事的倾向。崇尚自由奔放的恋爱，会拼尽全力爱一场，是就算明知无结果都在所不惜的冲动派。",
	"吉凶互见，一成一败，凶中有吉，吉中有凶|吉带凶|主人性格类型：[高度戒备难交心型]，其具体表现为：经常处于戒备状态，对任何事都没法放松处理，孤僻性情难以交朋结友。对于爱情，就更加会慎重处理。任何人必须经过你长期观察及通过连番考验，方会减除戒备与你交往。",
	"一盛一衰，浮沉不常，自重自处，可保平安|吉带凶|主人性格类型：[好奇心旺求知欲强型]，其具体表现为：好奇心极度旺盛，求知欲又强，有打烂沙盘问到笃的锲而不舍精神。此外，你天生有语言天分，学习外文比一般人更易掌握。",
	"草木逢春，雨过天晴，渡过难关，即获得成功|吉|主人性格类型：[做事喜好凭直觉型]，其具体表现为：有特强的第六灵感，性格率直无机心，深得朋辈爱戴。感情路上多采多姿。做事喜好凭个人直觉及预感做决定。",
	"盛衰参半，外祥内苦，先吉后凶，先凶后吉|吉带凶|主人性格类型：[独断独行/吸引人型]，其具体表现为：为人独断独行，事事自行作主解决，鲜有求助他人。而这份独立个性，正正就是吸引异性的特质。但其实心底里，却是渴望有人可让他/她依赖。",
	"虽倾全力，难望成功，此数大凶，最好改名|凶|主人性格类型：[热情/善变梦想家型]，其具体表现为：对人热情无遮掩，时常梦想可以谈一场戏剧性恋爱，亲身体会个中悲欢离合的动人经历，是个大梦想家。但对于感情却易变卦。",
	"外观昌隆，内隐祸患，克服难关，开出泰运|吉带凶|主人性格类型：[要面包不要爱情型]，其具体表现为：责任心重，尤其对工作充满热诚，是个彻头彻尾工作狂。但往往因为过分专注职务，而忽略身边的家人及朋友，是个宁要面包不需要爱情的理性主义者。",
	"事与愿违，终难成功，欲速不达，有始无终|凶|主人性格类型：[不善表达/疑心大型]，其具体表现为：在乎身边各人对自己的评价及观感，不善表达个人情感，是个先考虑别人感受，再作出相应配合的内敛一族。对于爱侣，经常存有怀疑之心。",
	"虽有困难，时来运转，旷野枯草，春来花开|吉|主人性格类型：[大胆行事冲动派型]，其具体表现为：爱好追寻刺激，有不理后果大胆行事的倾向。崇尚自由奔放的恋爱，会拼尽全力爱一场，是就算明知无结果都在所不惜的冲动派。",
	"半凶半吉，浮沉多端，始凶终吉，能保成功|凶带吉|主人性格类型：[高度戒备难交心型]，其具体表现为：经常处于戒备状态，对任何事都没法放松处理，孤僻性情难以交朋结友。对于爱情，就更加会慎重处理。任何人必须经过你长期观察及通过连番考验，方会减除戒备与你交往。",
	"遇事猜疑，难望成事，大刀阔斧，始可有成|凶|主人性格类型：[好奇心旺求知欲强型]，其具体表现为：好奇心极度旺盛，求知欲又强，有打烂沙盘问到笃的锲而不舍精神。此外，你天生有语言天分，学习外文比一般人更易掌握。",
	"黑暗无光，心迷意乱，出尔反尔，难定方针|凶|主人性格类型：[做事喜好凭直觉型]，其具体表现为：有特强的第六灵感，性格率直无机心，深得朋辈爱戴。感情路上多采多姿。做事喜好凭个人直觉及预感做决定。",
	"运遮半月，内隐风波，应有谨慎，始保平安|吉带凶|主人性格类型：[独断独行/吸引人型]，其具体表现为：为人独断独行，事事自行作主解决，鲜有求助他人。而这份独立个性，正正就是吸引异性的特质。但其实心底里，却是渴望有人可让他/她依赖。",
	"烦闷懊恼，事业难展，自防灾祸，始免困境|凶|主人性格类型：[热情/善变梦想家型]，其具体表现为：对人热情无遮掩，时常梦想可以谈一场戏剧性恋爱，亲身体会个中悲欢离合的动人经历，是个大梦想家。但对于感情却易变卦。",
	"万物化育，繁荣之象，专心一意，必能成功|吉|主人性格类型：[自我牺牲/性格被动型]，其具体表现为：惯于无条件付出，从不祈求有回报，有为了成全他人不惜牺牲自己的情操。但讲到本身的爱情观，却流于被动，往往因为内敛而错过大好姻缘。",
	"见异思迁，十九不成，徒劳无功，不如更名|凶|主人性格类型：[要面包不要爱情型]，其具体表现为：责任心重，尤其对工作充满热诚，是个彻头彻尾工作狂。但往往因为过分专注职务，而忽略身边的家人及朋友，是个宁要面包不需要爱情的理性主义者。",
	"吉运自来，能享盛名，把握时机，必获成功|吉|主人性格类型：[不善表达/疑心大型]，其具体表现为：在乎身边各人对自己的评价及观感，不善表达个人情感，是个先考虑别人感受，再作出相应配合的内敛一族。对于爱侣，经常存有怀疑之心。",
	"黑夜温长，进退维谷，内外不和，信用缺乏|凶|主人性格类型：[大胆行事冲动派型]，其具体表现为：爱好追寻刺激，有不理后果大胆行事的倾向。崇尚自由奔放的恋爱，会拼尽全力爱一场，是就算明知无结果都在所不惜的冲动派。",
	"独营事业，事事如意，功成名就，富贵自来|吉|主人性格类型：[好奇心旺求知欲强型]，其具体表现为：好奇心极度旺盛，求知欲又强，有打烂沙盘问到笃的锲而不舍精神。此外，你天生有语言天分，学习外文比一般人更易掌握。",
	"思虎虑周祥，计书力行，不失先机，可望成功|吉|主人性格类型：[不善表达/疑心大型]，其具体表现为：在乎身边各人对自己的评价及观感，不善表达个人情感，是个先考虑别人感受，再作出相应配合的内敛一族。对于爱侣，经常存有怀疑之心。",
	"动摇不安，常陷逆境，不得时运，难得利润|凶|主人性格类型：[大胆行事冲动派型]，其具体表现为：爱好追寻刺激，有不理后果大胆行事的倾向。崇尚自由奔放的恋爱，会拼尽全力爱一场，是就算明知无结果都在所不惜的冲动派。",
	"惨淡经营，难免贫困，此数不吉，最好改名|凶|主人性格类型：[高度戒备难交心型]，其具体表现为：经常处于戒备状态，对任何事都没法放松处理，孤僻性情难以交朋结友。对于爱情，就更加会慎重处理。任何人必须经过你长期观察及通过连番考验，方会减除戒备与你交往。",
	"吉凶参半，惟赖勇气，贯彻力行，始可成功|吉带凶|主人性格类型：[好奇心旺求知欲强型]，其具体表现为：好奇心极度旺盛，求知欲又强，有打烂沙盘问到笃的锲而不舍精神。此外，你天生有语言天分，学习外文比一般人更易掌握。",
	"利害混集，凶多吉少，得而复失，难以安顺|凶|主人性格类型：[做事喜好凭直觉型]，其具体表现为：有特强的第六灵感，性格率直无机心，深得朋辈爱戴。感情路上多采多姿。做事喜好凭个人直觉及预感做决定。",
	"安乐自来，自然吉祥，力行不懈，终必成功|吉|主人性格类型：[独断独行/吸引人型]，其具体表现为：为人独断独行，事事自行作主解决，鲜有求助他人。而这份独立个性，正正就是吸引异性的特质。但其实心底里，却是渴望有人可让他/她依赖。",
	"利不及费，坐食山空，如无章法，难望成功|凶|主人性格类型：[独断独行/吸引人型]，其具体表现为：为人独断独行，事事自行作主解决，鲜有求助他人。而这份独立个性，正正就是吸引异性的特质。但其实心底里，却是渴望有人可让他/她依赖。",
	"吉中带凶，欲速不达，进不如守，可保安祥|吉带凶|主人性格类型：[热情/善变梦想家型]，其具体表现为：对人热情无遮掩，时常梦想可以谈一场戏剧性恋爱，亲身体会个中悲欢离合的动人经历，是个大梦想家。但对于感情却易变卦。",
	"此数大凶，破产之象，宜速改名，以避厄运|凶|主人性格类型：[自我牺牲/性格被动型]，其具体表现为：惯于无条件付出，从不祈求有回报，有为了成全他人不惜牺牲自己的情操。但讲到本身的爱情观，却流于被动，往往因为内敛而错过大好姻缘。",
	"先苦后甘，先甜后苦，如能守成，不致失败|吉带凶|主人性格类型：[要面包不要爱情型]，其具体表现为：责任心重，尤其对工作充满热诚，是个彻头彻尾工作狂。但往往因为过分专注职务，而忽略身边的家人及朋友，是个宁要面包不需要爱情的理性主义者。",
	"有得有失，华而不实，须防劫财，始保安顺|吉带凶|主人性格类型：[不善表达/疑心大型]，其具体表现为：在乎身边各人对自己的评价及观感，不善表达个人情感，是个先考虑别人感受，再作出相应配合的内敛一族。对于爱侣，经常存有怀疑之心。",
	"如走夜路，前途无光，希望不大，劳而无功|凶|主人性格类型：[大胆行事冲动派型]，其具体表现为：爱好追寻刺激，有不理后果大胆行事的倾向。崇尚自由奔放的恋爱，会拼尽全力爱一场，是就算明知无结果都在所不惜的冲动派。",
	"得而复失，枉费心机，守成无贪，可保安稳|吉带凶|主人性格类型：[好奇心旺求知欲强型]，其具体表现为：好奇心极度旺盛，求知欲又强，有打烂沙盘问到笃的锲而不舍精神。此外，你天生有语言天分，学习外文比一般人更易掌握。",
	"最极之数，还本归元，能得繁荣，发达成功|吉|主人性格类型：[做事喜好凭直觉型]，其具体表现为：有特强的第六灵感，性格率直无机心，深得朋辈爱戴。感情路上多采多姿。做事喜好凭个人直觉及预感做决定。",
}
