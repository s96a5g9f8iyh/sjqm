package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nongli/ccal"
	"github.com/nongli/misc"
	"github.com/nongli/today"
	"github.com/nongli/utils"
	"github.com/sjqm"
)

var h12 int

//当前时间
var Ts = time.Now().Local()

func main() {
	//显示当日对应农历
	expectInfo, err := today.FindLunarMD()

	if err != nil {
		log.Fatal("时间异常:", err)
	}

	leapY := expectInfo.LeapY
	leapM := expectInfo.LeapM
	expectLeapD := expectInfo.ExpectleapD
	leapB := expectInfo.LeapB

	normalY := expectInfo.NormalY
	normalM := expectInfo.NormalM
	expectD := expectInfo.ExpectD
	normalB := expectInfo.NormalB

	T := time.Now().Local()
	h24 := T.Hour()
	h12 = utils.Conv24Hto12H(h24)
	sx := "猴"

	if leapM != 0 && leapB == true {
		err, s, l, g, _ := ccal.Input(leapY, leapM, expectLeapD, h12, sx, leapB)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		fmt.Printf("阳历纪年: %d年-%d月-%d日-周%s-阳历时间约:%d\n", s.SYear, s.SMonth, s.SDay, s.SWeek, s.SHour)
		fmt.Printf("农历纪年: %d年-%d月(%s)-%d日-%d时(%s时) 本年是否有闰月:%t 闰%d月\n",
			l.LYear, l.LMonth, l.LYdxs, l.LDay, l.LHour, l.LaliasHour, l.Leapmb, l.LeapMonth)
		fmt.Printf("干支纪年: %s%s年-%s月-%s%s日-%s时\n\n",
			g.YearGanM, g.YearZhiM, g.MonthGanZhiM, g.DayGanM, g.DayZhiM, g.HourGanZhiM)

		//杨公十三忌
		yginfo := misc.Yg13(l.LMonth, l.LDay)
		fmt.Printf("%s\n", yginfo)
		//八绝日
		jue8 := misc.Jue8(g.DayGanM, g.DayZhiM)
		fmt.Printf("%s\n", jue8)

	} else if normalM != 0 && normalB == false {
		err, s, l, g, _ := ccal.Input(normalY, normalM, expectD, h12, sx, normalB)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		fmt.Printf("阳历纪年: %d年-%d月-%d日-周%s-阳历时间约:%d\n", s.SYear, s.SMonth, s.SDay, s.SWeek, s.SHour)
		fmt.Printf("农历纪年: %d年-%d月(%s)-%d日-%d时(%s时) 本年是否有闰月:%t 闰%d月\n",
			l.LYear, l.LMonth, l.LYdxs, l.LDay, l.LHour, l.LaliasHour, l.Leapmb, l.LeapMonth)
		fmt.Printf("干支纪年: %s%s年-%s月-%s%s日-%s时\n\n",
			g.YearGanM, g.YearZhiM, g.MonthGanZhiM, g.DayGanM, g.DayZhiM, g.HourGanZhiM)

		//杨公十三忌
		yginfo := misc.Yg13(l.LMonth, l.LDay)
		fmt.Printf("%s\n", yginfo)
		//八绝日
		jue8 := misc.Jue8(g.DayGanM, g.DayZhiM)
		fmt.Printf("%s\n", jue8)

		////////////
		jqt := sjqm.JQT(l.LYear)
		jq := sjqm.Jq节气名称(jqt)
		_, dzt := jq.Jq冬至()
		_, xzt := jq.Jq夏至()

		xt := s.SHour //xt:阳历时间　需要把农历时间转换为阳历时间
		now := time.Date(s.SolarDayT.Year(), s.SolarDayT.Month(), s.SolarDayT.Day(), xt, 0, 0, 0, time.Local)

		yy := sjqm.YY阴阳判断(now, dzt, xzt)

		jmc := sjqm.Jqt当前节气(now, jqt)
		//fmt.Printf("当前节气名称:%s\n", jmc)
		yuan, _ := sjqm.FY符头三元(g.DayGan, g.DayZhi)
		//fmt.Printf("%s遁 %s %s 第%d天\n", yy, yuan, jmc, 第几天)

		////
		syinfo := sjqm.NewSY()
		//syinfo.DS大暑(jmc, yuan)
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
}
