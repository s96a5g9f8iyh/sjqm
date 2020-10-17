package sjqm

//节气计算
//精確到小時的二十四節氣

import (
	"time"

	"github.com/nongli/julian"
	"github.com/nongli/utils"
)

//从上年冬至到下年立春
type JQ节气 struct {
	Name []string    //节气名称数组
	Time []time.Time //节气时间数组
}

var Jmc = []string{
	"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰",
	"春分", "清明", "谷雨", "立夏", "小满", "芒种",
	"夏至", "小暑", "大暑", "立秋", "处暑", "白露",
	"秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
}

//夏至时间戳　夏至名称
func (jq *JQ节气) Jq夏至() (name string, time time.Time) {

	name = jq.Name[12]
	time = jq.Time[12] //i=12

	return
}

//冬至时间戳　冬至名称
func (jq *JQ节气) Jq冬至() (name string, time time.Time) {
	i := len(jq.Time) - 4
	name = jq.Name[0]
	time = jq.Time[i] //i=24

	return
}

//上一年冬至到下一年立春的节气
func Jq节气名称(jqt []time.Time) *JQ节气 {

	var jqmc []string     //节气名称
	var jqmct []time.Time //节气时间
	//i=28到下一年立春
	for i := 0; i < 28; i++ {
		x := i
		if x > 23 {
			x = i - 24
		}
		jqmc = append(jqmc, Jmc[x])
		jqmct = append(jqmct, jqt[i])
	}

	return &JQ节气{Name: jqmc, Time: jqmct}
}

//节气时间数组　精确到分钟
func JQT(y int) []time.Time {
	var j24t []time.Time
	data := julian.Data(y)
	for i := 1; i <= 25; i++ { //<=25冬至到冬至
		hs := data[0] + data[i]        //UTC合朔
		hs8 := utils.JdToLocalTime(hs) //CST+8
		j24t = append(j24t, hs8)
	}

	var j24t1 []time.Time
	data1 := julian.Data(y + 1)
	for j := 1; j <= 25; j++ {
		hs := data1[0] + data1[j]      //UTC合朔
		hs8 := utils.JdToLocalTime(hs) //CST+8
		j24t1 = append(j24t1, hs8)
	}

	//去重
	var kx int
I:
	for k := 0; k < len(j24t); k++ {
		xk := j24t[k]
		for k1 := 0; k1 < len(j24t1); k1++ {
			xk1 := j24t1[k1]
			if b := xk.Equal(xk1); b == true {
				kx = k1
				break I
			}
		}
	}
	//上一年冬至(含)开始 0:冬至　1:小寒　2:大寒...
	alljqt := append(j24t, j24t1[kx+1:]...)
	//fmt.Printf("*** allj1:%v\nlen=%d\n", alljqt, len(alljqt))
	return alljqt
}

//根据农历年份　返回本年全年节气时间戳的数组
//UTC+8时区 精确到小时
func Jq节气时间戳(y int) []time.Time {

	var j24t []time.Time
	var hxt time.Time
	data := julian.Data(y)
	for i := 1; i <= 25; i++ { //<=25冬至到冬至
		hs := data[0] + data[i]                                                                             //UTC合朔
		hs8 := utils.ToUTC8Time(hs)                                                                         //UTC+8
		hxt = time.Date(hs8.Year(), hs8.Month(), hs8.Day(), hs8.Hour(), 0, 0, 0, time.UTC /* time.Local */) //精确时间到小时
		j24t = append(j24t, hxt)
	}

	var hxt1 time.Time
	var j24t1 []time.Time
	data1 := julian.Data(y + 1)
	for j := 1; j <= 25; j++ {
		hs := data1[0] + data1[j]   //UTC合朔
		hs8 := utils.ToUTC8Time(hs) //UTC+8
		hxt1 = time.Date(hs8.Year(), hs8.Month(), hs8.Day(), hs8.Hour(), 0, 0, 0, time.UTC /* time.Local */)
		j24t1 = append(j24t1, hxt1)
	}

	//去重
	var kx int
I:
	for k := 0; k < len(j24t); k++ {
		xk := j24t[k]
		for k1 := 0; k1 < len(j24t1); k1++ {
			xk1 := j24t1[k1]
			if b := xk.Equal(xk1); b == true {
				kx = k1
				break I
			}
		}
	}
	//上一年冬至(含)开始 0:冬至　1:小寒　2:大寒...
	alljqt := append(j24t, j24t1[kx+1:]...)

	/* 	if kx == 0 {
	   		alljq := append(j24t, j24t1[kx+1:]...)
	   		fmt.Printf("allj1:%v\nlen=%d\n", alljq, len(alljq))
	   	} else {
	   		alljq := append(j24t, j24t1[kx+1:]...)
	   		fmt.Printf("allj1:%v\nlen=%d\n", alljq, len(alljq))
		   } */
	//fmt.Printf("### allj1:%v\nlen=%d\n", alljqt, len(alljqt))
	return alljqt
}

/* //節氣對應的陽曆時間
//返回值　info:节气信息　jxt:精确到小时的节气时间戳
func J24H(y int, jie string) (info string, jxt time.Time) {
	var _info string
	var _jxt time.Time

	data := julian.Data(y)                //兩冬至間的節氣
	for i := 1; i < len(solar.JMC); i++ { //從小寒開始
		hs := data[0] + data[i]
		hs8 := utils.ToUTC8Time(hs) //转换为UTC+8时区

		if b := strings.Contains(solar.JMC[i-1], jie); b == true {
			//TODO hs8作为一个时间戳返回　计算精确的局数字
			//_jxt = hs8 //精确的节气时间
			_info = fmt.Sprintf("%s: jd:%f 陽曆日期: %d年-%d月-%d日-%d时-%d分-%d秒\n", solar.JMC[i-1],
				hs, hs8.Year(), int(hs8.Month()), hs8.Day(), hs8.Hour(), hs8.Minute(), hs8.Second())
			//精确到时辰的节气　用来计算时辰交换
			ht := time.Date(hs8.Year(), hs8.Month(), hs8.Day(), hs8.Hour(), 0, 0, 0, time.Local)
			_jxt = ht
			//fmt.Printf("J24H:%v\n", ht)
			break
		}
	}
	info = _info
	jxt = _jxt
	return
}
*/
