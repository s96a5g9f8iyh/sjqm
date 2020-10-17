package sjqm

import (
	"strings"
)

//值符值使
//number:定局之后　旬首对应的数字
func (bg *BaGong) Zf值符(number int) (info *BaGong) {

	switch number {
	//////////阳宫
	case 1: //坎
		info = bg.Bg一宫()
	case 8: //艮
		info = bg.Bg八宫()
	case 3: //震
		info = bg.Bg三宫()
	case 4: //巽
		info = bg.Bg四宫()
	//////////阴宫
	case 9: //离
		info = bg.Bg九宫()
	case 2: //坤
		info = bg.Bg二宫()
	case 7: //兑
		info = bg.Bg七宫()
	case 6: //乾
		info = bg.Bg六宫()
	///////////中五宫
	case 5: //中5宫
		info = bg.Bg五宫()
	}
	return
}

//旬首:　时干支对应旬首　sqly:定局之后的三奇六仪
func ZF数字(旬首 string, sqly []map[string]int) (name string, number int) {
	for i := 0; i < len(sqly); i++ {
		for k, v := range sqly[i] {
			if b := strings.HasPrefix(k, 旬首); b == true {
				name = k
				number = v
			}
		}
	}
	return
}
