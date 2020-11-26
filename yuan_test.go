package sjqm

import (
	"strings"
	"testing"
)

var (
	jqNameArr = []string{
		"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰", "春分", "清明", "谷雨", "立夏", "小满", "芒种",
		"夏至", "小暑", "大暑", "立秋", "处暑", "白露", "秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
		"冬至", "小寒", "大寒", "立春"}
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

//定局测试
func Test定局(t *testing.T) {
	/*
		expy1, expy2, expy3 := 0, 1, 2 //定元数字
		index0 := 0                    //

		name0 := "冬至"
		dz1, dz2, dz3 := 1, 7, 4

		if strings.EqualFold(name0, jqNameArr[index0]) || strings.EqualFold(name0, jqNameArr[24]) {

			n1 := FindJU(expy1, index0, jqNameArr)
			if n1 != dz1 {
				t.Errorf("FindJu(%d %d %s)=%d 预期值:%d\n", expy1, index0, jqNameArr, n1, dz1)
			}

			n2 := FindJU(expy2, index0, jqNameArr)
			if n2 != dz2 {
				t.Errorf("FindJu(%d %d %s)=%d 预期值:%d\n", expy2, index0, jqNameArr, n2, dz2)
			}

			n3 := FindJU(expy3, index0, jqNameArr)
			if n3 != dz3 {
				t.Errorf("FindJu(%d %d %s)=%d 预期值:%d\n", expy3, index0, jqNameArr, n3, dz3)
			}
		}
	*/
	//////////////////
	expy := []int{0, 1, 2} //三元数组
	//index := []int{0, 1, 2} ==>i           //节气索引数组
	name := []string{
		"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰", "春分", "清明", "谷雨", "立夏", "小满", "芒种",
		"夏至", "小暑", "大暑", "立秋", "处暑", "白露", "秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
		"冬至", "小寒", "大寒", "立春"} //节气名称

	//局数字
	type jqn struct {
		n1, n2, n3 int //n1:上元 n2:中元 n3:下元
	}
	//节气三元定局数字
	exptYuan := []jqn{
		{1, 7, 4}, //冬至
		{2, 8, 5}, //小寒
		{3, 9, 6}, //大寒
		{8, 5, 2}, //立春
		{9, 6, 3}, //雨水
		{1, 7, 4}, //惊蛰
		{3, 9, 6}, //春分
		{4, 1, 7}, //清明
		{5, 2, 8}, //谷雨
		{4, 1, 7}, //立夏
		{5, 2, 8}, //小满
		{6, 3, 9}, //芒种

		{9, 3, 6}, //夏至
		{8, 2, 5}, //小暑
		{7, 1, 4}, //大暑
		{2, 5, 8}, //立秋
		{1, 4, 7}, //处暑
		{9, 3, 6}, //白露
		{7, 1, 4}, //秋分
		{6, 9, 3}, //寒露
		{5, 8, 2}, //霜降
		{6, 9, 3}, //立冬
		{5, 8, 2}, //小雪
		{4, 7, 1}, //大雪
	}

	for i := 0; i < len(exptYuan); i++ {
		if !strings.EqualFold(name[i], jqNameArr[i]) {
			t.Errorf("FindJU(_,%d,_) 节气:%s 预期节气:%s\n", i, name[i], jqNameArr[i])
		} else if strings.EqualFold(name[i], jqNameArr[i]) {
			for j := 0; j < len(expy); j++ {
				switch j {
				case 0:
					jx := FindJU(expy[j], i, jqNameArr)
					if jx != exptYuan[i].n1 {
						t.Errorf("FindJu(%d %d %s)=%d 预期值:%d\n", expy[j], i, jqNameArr, jx, exptYuan[i].n1)
					}
				case 1:
					jx := FindJU(expy[j], i, jqNameArr)
					if jx != exptYuan[i].n2 {
						t.Errorf("FindJu(%d %d %s)=%d 预期值:%d\n", expy[j], i, jqNameArr, jx, exptYuan[i].n2)
					}
				case 2:
					jx := FindJU(expy[j], i, jqNameArr)
					if jx != exptYuan[i].n3 {
						t.Errorf("FindJu(%d %d %s)=%d 预期值:%d\n", expy[j], i, jqNameArr, jx, exptYuan[i].n3)
					}
				}
			}
			break
		}
	}
}
