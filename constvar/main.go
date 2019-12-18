package constvar

var GroupUsersMap map[string][]string = make(map[string][]string)

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
}
