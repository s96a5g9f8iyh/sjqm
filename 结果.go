package sjqm

import (
	"time"

	"liangzi.local/nongli/solar"
)

func Result(y int, dgz, hgz string, st time.Time) *G {
	jqt := solar.JQT(y)
	jq := solar.NewJQ(jqt)
	jqnames := jq.Name //节气名称数组
	jqarr := jq.Time
	_, dzt := jq.Q冬至()
	_, xzt := jq.Q夏至()
	st = time.Date(st.Year(), st.Month(), st.Day(), st.Hour(), 0, 0, 0, time.Local)
	//定节气
	index, jmc := FindJQ(st, jqarr, jqnames)
	//定阴阳遁
	yinyangN := YY阴阳判断(st, dzt, xzt)
	var yy string
	if yinyangN == 0 {
		yy = "阴遁"
	}
	if yinyangN == 1 {
		yy = "阳遁"
	}
	//定元
	yuanx, yn := YuanN(dgz)
	//定局
	juN := FindJU(yn, index, jqnames)
	sqly := FindSqly(yinyangN, juN)
	//fmt.Printf("地盘三奇六仪:%v\n", sqly)
	//旬首
	xunshou := XunShou(hgz)
	//fmt.Printf("旬首:%s\n", xunshou)
	//找值符九星
	jgmap := JGMap()
	zhifu, zhishi := FindZiFu(xunshou, sqly, jgmap)
	//fmt.Printf("值符:%s 值使:%s\n", zhifu, zhishi)
	zhifN, starmap, dunmap := ZhiFu(zhifu, hgz, sqly)
	//fmt.Printf("九宫配九星:%v\n九宫配旬遁:%v\n", starmap, dunmap)
	//暗干支
	agzmap := AnGZ(zhifu, xunshou, yinyangN)
	//fmt.Printf("九宫配暗干支:%v\n", agzmap)
	//值使八门
	zhishimap := ZhiShi(hgz, agzmap, zhishi)
	//fmt.Printf("九宫配值使八门:%v\n", zhishimap)
	//八神
	bsmap := BaShen(zhifN, yinyangN)
	//fmt.Printf("九宫配八神:%v\n", bsmap)

	////////////////////////////////////////////////结果
	//值符(九星) 值使(八门) 暗干支 六甲旬遁 八神 六仪三奇
	var arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8, arr9 []string
	arr1 = append(arr1, starmap[1], zhishimap[1], agzmap[1], dunmap[1], bsmap[1], sqly[1]) //一宫信息
	arr2 = append(arr2, starmap[2], zhishimap[2], agzmap[2], dunmap[2], bsmap[2], sqly[2])
	arr3 = append(arr3, starmap[3], zhishimap[3], agzmap[3], dunmap[3], bsmap[3], sqly[3])
	arr4 = append(arr4, starmap[4], zhishimap[4], agzmap[4], dunmap[4], bsmap[4], sqly[4])
	arr5 = append(arr5, starmap[5], zhishimap[5], agzmap[5], dunmap[5], bsmap[5], sqly[5])
	arr6 = append(arr6, starmap[6], zhishimap[6], agzmap[6], dunmap[6], bsmap[6], sqly[6])
	arr7 = append(arr7, starmap[7], zhishimap[7], agzmap[7], dunmap[7], bsmap[7], sqly[7])
	arr8 = append(arr8, starmap[8], zhishimap[8], agzmap[8], dunmap[8], bsmap[8], sqly[8])
	arr9 = append(arr9, starmap[9], zhishimap[9], agzmap[9], dunmap[9], bsmap[9], sqly[9]) //九宫信息

	g := new(G)
	g = &G{
		JieQi:   jmc,
		YinYang: yy,
		N:       juN,
		YUAN:    yuanx,
		XS:      xunshou,
		ZHIFU:   zhifu,
		ZHISHI:  zhishi,
		G1:      arr1,
		G2:      arr2,
		G3:      arr3,
		G4:      arr4,
		G5:      arr5,
		G6:      arr6,
		G7:      arr7,
		G8:      arr8,
		G9:      arr9,
	}
	return g
}
