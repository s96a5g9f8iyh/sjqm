package sjqm

import (
	"strings"
	"testing"
)

//定元 测试
func Test定元(t *testing.T) {
	//上元 二十天
	s上元 := []string{"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己卯", "庚辰", "辛巳", "壬午", "癸未", "甲午", "乙未", "丙申", "丁酉", "戊戌", "己酉", "庚戌", "辛亥", "壬子", "癸丑"}
	//中元 二十天
	s中元 := []string{"己巳", "庚午", "辛未", "壬申", "癸酉", "甲申", "乙酉", "丙戌", "丁亥", "戊子", "己亥", "庚子", "辛丑", "壬寅", "癸卯", "甲寅", "乙卯", "丙辰", "丁巳", "戊午"}
	//下元 二十天
	s下元 := []string{"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己丑", "庚寅", "辛卯", "壬辰", "癸巳", "甲辰", "乙巳", "丙午", "丁未", "戊申", "己未", "庚申", "辛酉", "壬戌", "癸亥"}

	//预期值
	expected1 := "上元"
	expected2 := "中元"
	expected3 := "下元"

	n1, n2, n3 := 0, 1, 2

	//上元预期值测试
	for i := 0; i < len(s上元); i++ {
		actual, nx1 := YuanN(s上元[i])
		if !strings.EqualFold(actual, expected1) {
			t.Errorf("YuanN(%s)=%s 预期结果:%s", s上元[i], actual, expected1)
		}
		if nx1 != n1 {
			t.Errorf("YuanN(%s)=%d 预期结果:%d\n", s上元[i], nx1, n1)
		}
	}

	//中元预期值测试
	for j := 0; j < len(s中元); j++ {
		actual, nx2 := YuanN(s中元[j])
		if !strings.EqualFold(actual, expected2) {
			t.Errorf("YuanN(%s)=%s 预期结果:%s", s中元[j], actual, expected2)
		}
		if nx2 != n2 {
			t.Errorf("YuanN(%s)=%d 预期结果:%d\n", s中元[j], nx2, n2)
		}
	}

	//下元预期值测试
	for k := 0; k < len(s下元); k++ {
		actual, nx3 := YuanN(s下元[k])
		if !strings.EqualFold(actual, expected3) {
			t.Errorf("YuanN(%s)=%s 预期结果:%s\n", s下元[k], actual, expected3)
		}

		if nx3 != n3 {
			t.Errorf("YuanN(%s)=%d 预期结果:%d\n", s上元[k], nx3, n3)
		}
	}
}
