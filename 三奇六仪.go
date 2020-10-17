package sjqm

//阳遁顺布六仪逆布三奇　 戊己庚辛壬癸丁丙乙
//阴遁逆布六仪顺布三奇
//甲子遁戊　甲戌遁己　甲申遁庚　甲午遁辛　甲辰遁壬　甲寅遁癸

//当前局三奇六仪
//yy:阴阳遁　n:局数子
func New三奇六仪(yy string, n int) (newStarArrMap []map[string]int) {
	switch yy {
	case "阳":
		newStarArrMap = S阳遁三奇六仪(n)
	case "阴":
		newStarArrMap = S阴遁三奇六仪(n)
	}
	return
}

//阴遁三奇六仪
//number:局数字
func S阴遁三奇六仪(number int) []map[string]int {
	var n int
	var newStarArrMap []map[string]int
	//Go不能把map数组定义为常量　这里定义局部变量tmp
	tmp := []map[string]int{
		{"甲子-戊": 1}, {"甲戌-己": 2}, {"甲申-庚": 3}, {"甲午-辛": 4}, {"甲辰-壬": 5}, {"甲寅-癸": 6}, //六仪
		{"三奇-丁": 7}, {"三奇-丙": 8}, {"三奇-乙": 9}, //三奇
	}
	for i := 0; i < len(tmp); i++ {
		n = number
		n -= i
		if n < 1 {
			n = 9 + n
		}
		for k := range tmp[i] {
			tmp[i][k] = n
			newStarArrMap = append(newStarArrMap, tmp[i])
			break
		}
	}
	//fmt.Printf("阴遁三奇六仪:%v\n", newStarArrMap)
	return newStarArrMap
}

//阳遁三奇六仪
//number:局数字
func S阳遁三奇六仪(number int) []map[string]int {

	var n int
	var newStarArrMap []map[string]int
	//Go不能把map数组定义为常量　这里定义局部变量tmp
	tmp := []map[string]int{
		{"甲子-戊": 1}, {"甲戌-己": 2}, {"甲申-庚": 3}, {"甲午-辛": 4}, {"甲辰-壬": 5}, {"甲寅-癸": 6}, //六仪
		{"三奇-丁": 7}, {"三奇-丙": 8}, {"三奇-乙": 9}, //三奇
	}
	for i := 0; i < len(tmp); i++ {
		n = number
		n += i
		if n > 9 {
			n -= 9
		}

		for k := range tmp[i] {
			tmp[i][k] = n
			newStarArrMap = append(newStarArrMap, tmp[i] /* 六甲[i] */)
			break
		}
	}
	//fmt.Printf("阳遁三奇六仪:%v\n", newStarArrMap)
	return newStarArrMap
}
