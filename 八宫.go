package sjqm

/*
八门信息
坎一宫　五行水　天蓬星　休门 一白
艮八宫　五行土　天任星　生门 八白
震三宫　五行木　天冲星　伤门 三碧
巽四宫　五行木　天辅星　杜门 四绿
－－－－中宫(五)天禽       五黄
离九宫　五行火　天英星　景门 九紫
坤二宫　五行土　天芮星　死门 二黑
兑七宫　五行金　天柱星　惊门 七赤
乾六宫　五行金　天心星　开门 六白

九星信息
坎卦一宫水天蓬(贪狼)　艮卦八宫土天任(左辅)　 震卦三宫木天冲(禄存)　巽卦四宫木天辅(文曲)
中央土天禽(廉贞)
离卦九宫火天英(右弼)　坤卦二宫土天芮(rui)(巨门)　兑卦七宫金天柱(破军)　乾卦六宫金天心(武曲)
*/
//八门九星顺序为了方便统一为:1834 9276
type BaGong struct {
	Number     int    //宫位对应的数字
	Name       string //宫位名称
	YinYang    string //宫位的阴阳属性(坎艮震巽为阳　离坤兑乾为阴)
	StarName   string //天星 值符
	EightDoors string //八门
	Self       string //五行属性

}

func NewBG() (bg *BaGong) {
	bg = new(BaGong)
	return
}

//一宫　坎
// "冬至", "小寒", "大寒"
func (bg *BaGong) Bg一宫() *BaGong {
	return &BaGong{
		Number:     1,
		Name:       "坎",
		YinYang:    "阳",
		StarName:   "天蓬(贪狼)",
		EightDoors: "休门",
		Self:       "水",
	}
}

//八宫　艮
// "立春", "雨水", "惊蛰"
func (bg *BaGong) Bg八宫() *BaGong {
	return &BaGong{
		Number:     8,
		Name:       "艮",
		YinYang:    "阳",
		StarName:   "天任(左辅)",
		EightDoors: "生门",
		Self:       "土",
	}
}

//三宫　震
//"春分", "清明", "谷雨"
func (bg *BaGong) Bg三宫() *BaGong {

	return &BaGong{
		Number:     3,
		Name:       "震",
		YinYang:    "阳",
		StarName:   "天冲(禄存)",
		EightDoors: "伤门",
		Self:       "木",
	}
}

//四宫　巽
//"立夏", "小满", "芒种",
func (bg *BaGong) Bg四宫() *BaGong {

	return &BaGong{
		Number:     4,
		Name:       "巽",
		YinYang:    "阳",
		StarName:   "天辅(文曲)",
		EightDoors: "杜门",
		Self:       "木",
	}
}

//九宫　离
//"夏至", "小暑", "大暑",
func (bg *BaGong) Bg九宫() *BaGong {

	return &BaGong{
		Number:     9,
		Name:       "离",
		YinYang:    "阴",
		StarName:   "天英(右弼)",
		EightDoors: "景门",
		Self:       "火",
	}
}

//二宫　坤
//"立秋", "处暑", "白露",
func (bg *BaGong) Bg二宫() *BaGong {

	return &BaGong{
		Number:     2,
		Name:       "坤",
		YinYang:    "阴",
		StarName:   "天芮(巨门)",
		EightDoors: "死门",
		Self:       "土",
	}
}

//七宫　兑
//"秋分", "寒露", "霜降",
func (bg *BaGong) Bg七宫() *BaGong {

	return &BaGong{
		Number:     7,
		Name:       "兑",
		YinYang:    "阴",
		StarName:   "天柱(破军)",
		EightDoors: "惊门",
		Self:       "金",
	}
}

//六宫　乾
//"立冬", "小雪", "大雪",
func (bg *BaGong) Bg六宫() *BaGong {

	return &BaGong{
		Number:     6,
		Name:       "乾",
		YinYang:    "阴",
		StarName:   "天心(武曲)",
		EightDoors: "开门",
		Self:       "金",
	}
}

//中五宫 寄坤二宫
func (bg *BaGong) Bg五宫() *BaGong {
	return &BaGong{
		Number:     5,
		Name:       "",
		YinYang:    "",
		StarName:   "天禽",
		EightDoors: "死门",
		Self:       "",
	}
}
