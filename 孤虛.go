package sjqm

import (
	"fmt"
	"strings"
)

//孤虛法(旬中孤虚法)等同于时孤虚
//返回当日干支的孤虚时辰　如果当前时辰正好是孤虚时辰　那么返回当前时辰孤虚对应的地支
//dgz:日干支 hgz:時辰干支　也适用月日(dgz:月　hgz:日)
func SJQM孤虚法(dgz, hgz string) (info string, gx map[string]string) {

	//六十甲子
	lsjz := []string{
		"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉",
		"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未",
		"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳",
		"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯",
		"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑",
		"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥",
	}

	for i := 0; i < len(lsjz); i++ {
		if strings.EqualFold(dgz, lsjz[i]) { //尋找日干支在六十甲子中的索引號
			switch i {
			case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9: //甲子旬
				//日干支对应时辰
				info = fmt.Sprintf("\"日干支\":%s 甲子旬　孤:戌亥　虚:辰巳", dgz)
				//当前时辰正好为为孤虚时辰
				if strings.ContainsAny(hgz, "戌") {
					hgu := make(map[string]string)
					hgu = map[string]string{"孤": "戌", "虚": "辰"}
					gx = hgu
				} else if strings.ContainsAny(hgz, "亥") {
					hgu := make(map[string]string)
					hgu = map[string]string{"孤": "亥", "虚": "巳"}
					gx = hgu
				}
			case 10, 11, 12, 13, 14, 15, 16, 17, 18, 19: //甲戌旬
				//日干支对应时辰
				info = fmt.Sprintf("\"日干支\":%s 甲戌旬　孤:戌亥　虚:辰巳", dgz)
				//当前时辰正好为为孤虚时辰
				if strings.ContainsAny(hgz, "申") {
					hgu := make(map[string]string)
					hgu = map[string]string{"孤": "申", "虚": "寅"}
					gx = hgu
				} else if strings.ContainsAny(hgz, "酉") {
					hgu := make(map[string]string)
					hgu = map[string]string{"孤": "酉", "虚": "卯"}
					gx = hgu
				}
			case 20, 21, 22, 23, 24, 25, 26, 27, 28, 29: //甲申
				//日干支对应时辰
				info = fmt.Sprintf("\"日干支\":%s 甲申旬　孤:戌亥　虚:辰巳", dgz)
				//当前时辰正好为为孤虚时辰
				if strings.ContainsAny(hgz, "午") {
					hgu := make(map[string]string)
					hgu = map[string]string{"孤": "午", "虚": "子"}
					gx = hgu
				} else if strings.ContainsAny(hgz, "未") {
					hgu := make(map[string]string)
					hgu = map[string]string{"孤": "未", "虚": "丑"}
					gx = hgu
				}
			case 30, 31, 32, 33, 34, 35, 36, 37, 38, 39: //甲午
				//日干支对应时辰
				info = fmt.Sprintf("\"日干支\":%s 甲午旬　孤:戌亥　虚:辰巳", dgz)
				//当前时辰正好为为孤虚时辰
				if strings.ContainsAny(hgz, "辰") {
					hgu := make(map[string]string)
					hgu = map[string]string{"孤": "辰", "虚": "戌"}
					gx = hgu
				} else if strings.ContainsAny(hgz, "巳") {
					hgu := make(map[string]string)
					hgu = map[string]string{"孤": "巳", "虚": "亥"}
					gx = hgu
				}
			case 40, 41, 42, 43, 44, 45, 46, 47, 48, 49: //甲辰
				//日干支对应时辰
				info = fmt.Sprintf("\"日干支\":%s 甲辰旬　孤:戌亥　虚:辰巳", dgz)
				//当前时辰正好为为孤虚时辰
				if strings.ContainsAny(hgz, "寅") {
					hgu := make(map[string]string)
					hgu = map[string]string{"孤": "寅", "虚": "申"}
					gx = hgu
				} else if strings.ContainsAny(hgz, "卯") {
					hgu := make(map[string]string)
					hgu = map[string]string{"孤": "卯", "虚": "酉"}
					gx = hgu
				}
			case 50, 51, 52, 53, 54, 55, 56, 57, 58, 59: //甲寅
				//日干支对应时辰
				info = fmt.Sprintf("\"日干支\":%s 甲寅旬　孤:戌亥　虚:辰巳", dgz)
				//当前时辰正好为为孤虚时辰
				if strings.ContainsAny(hgz, "子") {
					hgu := make(map[string]string)
					hgu = map[string]string{"孤": "子", "虚": "午"}
					gx = hgu
				} else if strings.ContainsAny(hgz, "丑") {
					hgu := make(map[string]string)
					hgu = map[string]string{"孤": "丑", "虚": "未"}
					gx = hgu
				}
			}
		}
	}
	return
}
