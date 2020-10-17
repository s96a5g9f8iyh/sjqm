package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nongli/ccal"
	"github.com/sjqm"
	"github.com/sjqm/cmd"
)

func main() {
	infos := cmd.GetInPut()
	y := infos.Y
	m := infos.M
	d := infos.D
	h := infos.H
	b := infos.B
	sx := "猴"

	//生成日期信息
	err, s, l, g, _ := ccal.Input(y, m, d, h, sx, b)
	if err != nil {
		log.Fatal("ccal:", err)
	}

	fmt.Printf("阳历纪年: %d年-%d月-%d日-周%s-阳历时间范围:%d\n", s.SYear, s.SMonth, s.SDay, s.SWeek, s.SHour)
	fmt.Printf("农历纪年: %d年-%d月(%s)-%d日-%d时(%s时) 本年是否有闰月:%t 闰%d月\n",
		l.LYear, l.LMonth, l.LYdxs, l.LDay, l.LHour, l.LaliasHour, l.Leapmb, l.LeapMonth)
	fmt.Printf("干支纪年: %s%s年-%s月-%s%s日-%s时\n\n",
		g.YearGanM, g.YearZhiM, g.MonthGanZhiM, g.DayGanM, g.DayZhiM, g.HourGanZhiM)

	//yg := g.YearGanM
	//yz := g.YearZhiM
	//ygz := fmt.Sprintf("%s%s", yg, yz) //年干支
	//dg := g.DayGanM
	//dz := g.DayZhiM
	//dgz := fmt.Sprintf("%s%s", dg, dz) //日干支

	////////////
	jqt := sjqm.JQT(l.LYear)
	jq := sjqm.Jq节气名称(jqt)
	_, dzt := jq.Jq冬至()
	_, xzt := jq.Jq夏至()

	xt := h*2 - 1 //9 //xt:阳历时间　需要把农历时间转换为阳历时间
	now := time.Date(s.SolarDayT.Year(), s.SolarDayT.Month(), s.SolarDayT.Day(), xt, 0, 0, 0, time.Local)
	yy := sjqm.YY阴阳判断(now, dzt, xzt)

	jmc := sjqm.Jqt当前节气(now, jqt)
	//fmt.Printf("当前节气名称:%s\n", jmc)
	yuan, _ := sjqm.FY符头三元(g.DayGan, g.DayZhi)
	//fmt.Printf("%s遁 %s %s 第%d天\n", yy, yuan, jmc, 第几天)

	////
	syinfo := sjqm.NewSY()
	syinfo.DS大暑(jmc, yuan)
	_, number, _ := syinfo.J局信息(jmc, yuan)
	//fmt.Printf("节气名称:%s 局数字:%d 八宫数字:%d\n", name, number, bgn)
	fmt.Printf("节气:%s %s遁 %s %d局\n", jmc, yy, yuan, number)

	///旬首
	xunshou := sjqm.XS旬首(g.HourGanZhiM)
	fmt.Printf("旬首:%s\n", xunshou)
	////三奇六仪
	sqly := sjqm.New三奇六仪(yy, number)
	////值符数字
	_, 旬首数字 := sjqm.ZF数字(xunshou, sqly)
	//fmt.Printf("旬首名称:%s 旬首数字:%d\n", name, 旬首数字)

	////值符　八门
	bginfo := sjqm.NewBG()
	info := bginfo.Zf值符(旬首数字)
	fmt.Printf("值符:%s 八门:%s\n", info.StarName, info.EightDoors)

}
