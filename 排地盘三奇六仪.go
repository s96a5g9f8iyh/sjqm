package sjqm

import "fmt"

/*
六仪 戊己庚辛壬癸
三奇 乙丙丁
阳遁顺布六仪逆布三奇
阴遁逆布六仪顺布三奇
甲子遁戊　甲戌遁己　甲申遁庚　甲午遁辛　甲辰遁壬　甲寅遁癸
*/

//排地盘局
func FindSqly(yy int, juN int) (sqly map[int]string) {
	//yy:阴阳数字　n:局数子
	switch yy {
	case 1: //阳
		sqly = SqlyYangDun(juN)
	case 0: //阴
		sqly = SqlyYinDun(juN)
	}
	return
}

//排阴遁地盘三奇六仪
func SqlyYinDun(juN int) map[int]string {
	//juN 定局数字
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	gzarr := []string{"戊", "己", "庚", "辛", "壬", "癸", "丁", "丙", "乙"} //阴遁逆布六仪顺布三奇
	gzmap := make(map[int]string)
	//重新排序
	for j := 0; j < len(x); j++ {
		n := juN - j
		if n < 1 {
			n += 9
		}
		gzmap[n] = gzarr[j]
	}
	//fmt.Printf("阴遁地盘三奇六仪:%v\n", gzmap) //map[1:乙 2:丙 3:丁 4:癸 5:壬 6:辛 7:庚 8:己 9:戊]
	return gzmap
}

//排阳遁地盘三奇六仪
func SqlyYangDun(juN int) map[int]string {
	//juN 定局数字
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	gzarr := []string{"戊", "己", "庚", "辛", "壬", "癸", "丁", "丙", "乙"} //阳遁顺布六仪逆布三奇
	gzmap := make(map[int]string)
	//重新排序
	for j := 0; j <= len(x); j++ {
		if x[j] == juN {
			x1 := x[j:]
			x2 := x[:j]
			x = append(x1, x2...)
			break
		}
	}
	fmt.Println(x)
	for i := 0; i < len(x); i++ {
		gzmap[x[i]] = gzarr[i]
	}

	fmt.Printf("阳遁地盘三奇六仪:%v\n", gzmap)
	return gzmap
}

////////////////////////////////////////////////////////下面弃用
//排地盘局 map[string]int string:六仪 int:六仪对应的九宫位置(数字)
//九宫 二四为肩(左4右2)；六八为足(左8右6) 戴九履一(492 816) 左三右七 五中宫
/*func Find三奇六仪(yy int, n int) (newStarArrMap []map[string]int) {
	//yy:阴阳数字　n:局数子
	switch yy {
	case 1: //阳
		newStarArrMap = 阳遁三奇六仪(n)
	case 0: //阴
		newStarArrMap = 阴遁三奇六仪(n)
	}
	return
}

//阴遁三奇六仪
func 阴遁三奇六仪(number int) []map[string]int {
	//number:局数字
	var n int
	var newStarArrMap []map[string]int
	//甲子遁戊　甲戌遁己　甲申遁庚　甲午遁辛　甲辰遁壬　甲寅遁癸
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
func 阳遁三奇六仪(number int) []map[string]int {
	//number:局数字
	var n int
	var newStarArrMap []map[string]int
	//甲子遁戊　甲戌遁己　甲申遁庚　甲午遁辛　甲辰遁壬　甲寅遁癸
	tmp := []map[string]int{
		{"甲子-戊": 1}, {"甲戌-己": 2}, {"甲申-庚": 3}, {"甲午-辛": 4}, {"甲辰-壬": 5}, {"甲寅-癸": 6}, //六仪
		{"三奇-丁": 7}, {"三奇-丙": 8}, {"三奇-乙": 9}, //三奇
	}
	for i := 0; i < len(tmp); i++ {
		n = number
		n -= i
		if n > 9 {
			n -= 9
		}
		for k := range tmp[i] {
			tmp[i][k] = n
			newStarArrMap = append(newStarArrMap, tmp[i])
			break
		}
	}
	//fmt.Printf("阳遁三奇六仪:%v\n", newStarArrMap)
	return newStarArrMap
}
*/
