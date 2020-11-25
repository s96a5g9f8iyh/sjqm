package sjqm

import (
	"strings"
)

//值符原始宫位起时辰旬首 阳顺阴逆排
func AnGZ(zhifu, xunshou string, yinyang int) (agz map[int]string) {
	//九星原始宫位
	mapjiu := map[int]string{1: "天蓬", 8: "天任", 3: "天冲", 4: "天辅", 9: "天英", 2: "天芮", 7: "天柱", 6: "天心", 5: "天禽"} //九宫原始信息
	var n int
	for value, name := range mapjiu {
		if strings.EqualFold(zhifu, name) {
			n = value
			break
		}
	}
	//旬首干支数组
	var gzarr []string
	for i := 0; i < len(jz60); i++ {
		if strings.EqualFold(xunshou, jz60[i]) {
			gzarr = jz60[i : i+10]
			break
		}
	}

	//原始宫位
	x := []int{9, 8, 7, 6, 5, 4, 3, 2, 1} //阴遁顺序
	for j := 0; j < len(x); j++ {
		if x[j] == n {
			x1 := x[:j]
			x2 := x[j:]
			x = append(x2, x1...) //排序后宫位
		}
	}

	var aGZmap = make(map[int]string) //暗干支
	switch yinyang {
	case 0: //阴 逆
		x := []int{9, 8, 7, 6, 5, 4, 3, 2, 1} //阴遁顺序
		for j := 0; j < len(x); j++ {
			if x[j] == n {
				x1 := x[:j]
				x2 := x[j:]
				x = append(x2, x1...)
			}
		}
		///干支配值符原始宫
		for gn := 0; gn < len(gzarr); gn++ {
			for xn := gn; xn < len(x); xn++ {
				aGZmap[x[xn]] = gzarr[gn]
				break
			}
		}
		aGZmap[x[0]] = gzarr[0] + gzarr[9] //最后一个暗干支加到起始宫位上(值符原始宫位)
		agz = aGZmap
	case 1: //阳 顺
		x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //阳遁顺序
		for j := 0; j < len(x); j++ {
			if x[j] == n {
				x1 := x[:j]
				x2 := x[j:]
				x = append(x2, x1...)
			}
		}
		///干支配值符原始宫
		for gn := 0; gn < len(gzarr); gn++ {
			for xn := gn; xn < len(x); xn++ {
				aGZmap[x[xn]] = gzarr[gn]
				break
			}
		}
		aGZmap[x[0]] = gzarr[0] + gzarr[9]
		agz = aGZmap
	}
	return
}
