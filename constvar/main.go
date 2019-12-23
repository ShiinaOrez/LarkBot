package constvar

var GroupUsersMap = make(map[string][]string)
var WebHooks = make(map[string]map[string][]string)
var NumberToEmoji = map[int]string{
	1: "🤓",
	2: "🤩",
	3: "😝",
	4: "🥰",
	5: "🤪",
}

func init() {
	GroupUsersMap["backend"] = []string{
		"ShiinaOrez",    // 宋汝阳
		"Shadowmaple",   // 章茗超
		"Bowser1704",    // 余鸿奇
		"jiangzc",       // 蒋志成
		"hjm1027",       // 胡嘉旻
		"Chiwency",      // 邓永骏
		"hlyyy",         // 黄凌云
		"jepril",        // 洪欣然
		"MitsuhaOma",    // 王雯坚
		"kocoler",       // 张军洁
		"JacksieCheung", // 张竣淇
	}
	WebHooks["github"] = map[string][]string{
		"push": {
			"https://open.feishu.cn/open-apis/bot/hook/98fc59eb14d2405e880a6ab0fe70d136",
		},
		"go": {
			"https://open.feishu.cn/open-apis/bot/hook/cb973deacb4a4ee699d8d049c51e6908",
		},
		"java": {
			"https://open.feishu.cn/open-apis/bot/hook/0c2f9f5bb48849bda64cd25ebc9f87e1",
		},
		"kotlin": {
			"https://open.feishu.cn/open-apis/bot/hook/0c2f9f5bb48849bda64cd25ebc9f87e1",
		},
	}
}
