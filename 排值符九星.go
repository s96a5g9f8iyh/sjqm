package sjqm

import (
	"strings"
)

//六甲旬遁: 甲子遁戊　甲戌遁己　甲申遁庚　甲午遁辛　甲辰遁壬　甲寅遁癸
//时辰旬首查值符值使
//值符
func FindZiFu(xunshou string, sqly map[int]string, jgmap map[int]JG) (zf, zs string) {
	//sqly:定地盘局之后的三奇六仪
	//根据旬首和地盘三奇六仪 找出六甲旬遁所在的宫位
	xunDun := map[string]string{"甲子": "戊", "甲戌": "己", "甲申": "庚", "甲午": "辛", "甲辰": "壬", "甲寅": "癸"}
	for liuJia, dun := range xunDun {
		if strings.EqualFold(xunshou, liuJia) { //找旬首对应的六甲
			for N, name := range sqly { //根据旬首对应的六甲找到遁
				if strings.EqualFold(name, dun) {
					//fmt.Printf("六甲旬遁: %s遁: %s 所在宫位数字为:%d\n", liuJia, dun, N)
					//看当前宫原始宫位的信息 查值符(九星) 值使(八门)
					for k, v := range jgmap {
						if N == k { //遁宫数字==原始宫位数字 这里没有中宫
							zf = v.ZF //值符
							zs = v.ZS //值使
							break
						}
					}
					break
				}
			}
			break
		}
	}
	return
}

//排值符九星
//值符随时干 从时辰天干所在局数字顺排九星 同时带九星所在地盘的三奇六仪
//九星: 天蓬 天任 天冲 天辅 天英 天芮 天柱 天心 天禽(中宫)寄坤二宫即天芮星
func ZhiFu(zhifu, hgz string, sqly map[int]string) (zhifN int, gstar, gdun map[int]string) {
	//九星排序
	//jiu := []string{"天蓬", "天任", "天冲", "天辅", "天英", "天芮", "天柱", "天心", "天禽"}
	jiux := []string{"天蓬", "天任", "天冲", "天辅", "天英", "天芮", "天柱", "天心"}
	//如果值符是天禽 天禽寄2宫
	for a := 0; a < len(jiux); a++ {
		if strings.EqualFold(zhifu, "天禽") {
			if strings.EqualFold("天芮", jiux[a]) {
				s1 := jiux[:a]
				s2 := jiux[a:]
				jiux = append(s2, s1...)
			}
		}
		if strings.EqualFold(zhifu, jiux[a]) && !strings.EqualFold(zhifu, "天禽") {
			s1 := jiux[:a]
			s2 := jiux[a:]
			jiux = append(s2, s1...)
		}
	}
	jiu := append(jiux, "天禽") //放到最后 寄坤二宫?

	//九星原始宫位配旬遁
	mapjiu := map[int]string{1: "天蓬", 8: "天任", 3: "天冲", 4: "天辅", 9: "天英", 2: "天芮", 7: "天柱", 6: "天心", 5: "天禽"} //九宫原始信息
	var jiudun = make(map[string]string)
	for number, dun := range sqly {
		for k, v := range mapjiu {
			if k == number {
				jiudun[v] = dun
				continue
			}
		}
	}
	//fmt.Printf("天星-九遁:%v\n", jiudun)

	var nstar = make(map[int]string) //九宫配九星
	var nxd = make(map[int]string)   //九宫配旬遁
	//返回原始宫位对应的九星旬遁信息
	x := []int{1, 8, 3, 4, 9, 2, 7, 6} //九宫原始位顺序
	for n, gan := range sqly {         //n:三奇六仪对应的宫数字
		if strings.ContainsAny(hgz, gan) { //找到时辰天干所在宫位数字 如果是天禽 这里就不对了
			//fmt.Printf("%s时 六仪天干:%s %d宫\n", hgz, gan, n)
			//重排九宫
			for i := 0; i < len(x); i++ {
				if x[i] == n { //找到遁数字n在x数组的索引
					x1 := x[:i]
					x2 := x[i:]
					x = append(x2, x1...) //从时干所在宫位开始顺排九星
					x = append(x, 5)      //寄坤二宫
					//开始排九星...
					for xn := 0; xn < len(x); xn++ { //
						for jn := xn; jn < len(jiu); jn++ {
							//	fmt.Printf("顺排九数:%d宫-九星%s\n", x[xn], jiu[jn])
							//旬遁匹配(原始宫位的九星对应的旬遁)
							for star, dun := range jiudun {
								if strings.EqualFold(star, jiu[jn]) { //最终结果
									//	fmt.Printf("原始宫位:%d 九星:%s 该宫位旬遁:%s\n", x[xn], star, dun)
									//返回结果
									nstar[x[xn]] = star //九宫配九星
									nxd[x[xn]] = dun    //九宫配旬遁
									//返回值符的宫位 排八神用
									if strings.EqualFold(star, zhifu) { //值符随时干
										//fmt.Printf("值符:%s 宫位:%d\n", star, n)
										zhifN = n
									}
								}
							}
							break
						}
					}
					break
				}

			}
			break
		}
	}
	gstar = nstar
	gdun = nxd

	return
}
