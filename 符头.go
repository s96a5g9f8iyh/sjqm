package sjqm

//符頭所屬那一元
//dg:當日天干數字 dz:當日地支數字
func FY符头三元(dg, dz int) (yuan string, xday int) {
	fg, offg := FG符头天干(dg)
	fz := FZ符头地支(dz, offg)

	switch fg {
	case 1: //符頭甲
		yuan = yuanx(fz)
	case 6: //符頭己
		yuan = yuanx(fz)
	}
	xday = offg + 1
	return
}

//符頭天干數字
//offg:天干偏移數字 天干偏移數字用來計算地支　同時計算三元定局本日和符頭的偏移天數(offg+1)
//g:天干数子 甲1 乙2 丙3 丁4...
func FG符头天干(g int) (fg, offg int) {
	if g <= 5 {
		fg = 1 //符頭甲
		offg = g - 1
	} else {
		fg = 6 //符頭己
		offg = g - 6
	}
	return
}

//符頭地支數字
//dz:日支數字　offg：FG符头天干函數返回的差值
func FZ符头地支(dz, offg int) (fz int) {
	x := 12 - offg + dz
	if x > 12 {
		fz = x - 12
	} else {
		fz = x
	}
	return
}

//符头三元
func yuanx(fz int) (yuanx string) {
	switch fz {
	case 1, 7, 4, 10: //子1 午7 卯4 酉10　上元
		yuanx = "上元"
	case 3, 9, 6, 12: //寅3 申9 巳6 亥12　中元
		yuanx = "中元"
	case 5, 11, 2, 8: //辰5 戌11 醜2 未8　下元
		yuanx = "下元"
	}
	return
}
