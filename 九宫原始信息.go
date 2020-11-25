package sjqm

//九宫原始信息 少了中宫....
func JGMap() map[int]JG {
	var jginfo = make(map[int]JG)
	jginfo = map[int]JG{
		/////////////四阳宫
		1: {Name: "坎", Number: 1, ZF: "天蓬", ZS: "休门"}, //坎一宫
		8: {Name: "艮", Number: 8, ZF: "天任", ZS: "生门"}, //艮八宫
		3: {Name: "震", Number: 3, ZF: "天冲", ZS: "伤门"}, //震三宫
		4: {Name: "巽", Number: 4, ZF: "天辅", ZS: "杜门"}, //巽四宫
		//////////////四阴宫
		9: {Name: "离", Number: 9, ZF: "天英", ZS: "景门"}, //离九宫
		2: {Name: "坤", Number: 2, ZF: "天芮", ZS: "死门"}, //坤二宫
		7: {Name: "兑", Number: 7, ZF: "天柱", ZS: "惊门"}, //兑七宫
		6: {Name: "乾", Number: 6, ZF: "天心", ZS: "开门"}, //乾六宫
		5: {Name: "中", Number: 5, ZF: "天禽", ZS: "死门"}, //中宫
	}
	return jginfo
}
