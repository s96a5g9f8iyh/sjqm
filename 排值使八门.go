package sjqm

import (
	"strings"
)

//值使(八门)随时辰地支 从时辰地支宫顺时针排
func ZhiShi(hgz string, anGZ map[int]string, zhis string) (zhiShimap map[int]string) {
	//anGZ:暗干支 hgz:时辰干支 zhis:值使(八门)
	x := []int{1, 8, 3, 4, 9, 2, 7, 6} //5最后加
	bm := []string{"休", "生", "伤", "杜", "景", "死", "惊", "开"}

	var zsmap = make(map[int]string)
	for n, name := range anGZ {
		if strings.EqualFold(hgz, name) {
			//宫位排序
			for i := 0; i < len(x); i++ {
				if n == x[i] {
					x1 := x[:i]
					x2 := x[i:]
					x = append(x2, x1...)
					x = append(x, 5) //中宫排最后
					break
				}
			}
			//八门排序
			for j := 0; j < len(bm); j++ {
				if strings.ContainsAny(bm[j], zhis) {
					bm1 := bm[:j]
					bm2 := bm[j:]
					bm = append(bm2, bm1...)
					break
				}
			}
			//八门配宫
			for ix := 0; ix < len(x); ix++ {
				for mx := ix; mx < len(bm); mx++ {
					zsmap[x[ix]] = bm[mx]
					break
				}
			}
			break
		}
	}
	zhiShimap = zsmap
	return
}
