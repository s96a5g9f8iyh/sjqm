package sjqm

import (
	"time"
)

//傳入當日陽曆的時間戳st  jqt:年份節氣时间戳
func Jqt当前节气(st time.Time, jqt []time.Time) (jmc string) {

	st = time.Date(st.Year(), st.Month(), st.Day(), st.Hour(), 0, 0, 0, time.UTC) //精确到小时

	for i := 0; i < len(jqt); i++ {
		//當日時間到最近節氣的天數
		t1 := jqt[i]
		t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), 0, 0, 0, time.UTC)
		t := t1.Sub(st)
		x := int(t.Hours())

		//x=8当前时间即是节气时间
		if x == 8 {
			n := i
			//fmt.Printf("i8:%d x8:%d\n %v %v n=%d\n", i, x, st, jqt[i], n)
			jmc = Jmc[n]
			break
		}
		//x>24 ==> i-1=对应24节气的索引数字 当前时间在索引数字对应的节气之后
		if x > 24 && x != 8 {
			n := i - 1
			//fmt.Printf("i:%d x:%d %v %v n=%d\n", i, x, st, jqt[n], n)
			jmc = Jmc[n]
			break
		}
		if (x > 0 && x < 24) && x != 8 {
			n := i //- 1
			//fmt.Printf("i:%d x:%d %v %v n=%d\n", i, x, st, jqt[n], n)
			jmc = Jmc[n]
			break
		}
	}
	return
}

/*
//返回預測日的節氣名稱　精确到日
//傳入當日陽曆的時間戳st  jqt:年份節氣时间戳
func JTx当前节气时间(st time.Time, jqt []time.Time) (jmc string) {
	//先找出距離st最近的一個節氣時間戳
	//再看st是在節氣時間戳的前面還是後面
	//然後確定用局
	//st = st.AddDate(0, 0, 16)
	var n int
	var tx int
	for i := 0; i < len(jqt); i++ {

		//當日時間到最近節氣的天數
		//tx>=0当前时间在最近的节气时间之后　或者当日等于节气时间
		//tx<0说明当前时间没有超过下一个节气时间(还在最近的节气时间之后)
		tx = int(st.Sub(jqt[i])) / (24 * 3600000000000) //精确到天

		//一氣三元十五天　節氣的時間間隔15　當日到最近的節日時間必然小於15
		if tx < 15 {
			n = i
			break
		}
	}

	//距离当前日期最近的节气　这个节气时间可能在当日之前也可能在当日之后
	//fmt.Printf("\tn=%d\n", n)
	if n > 24 {
		n -= 24
	}
	b := st.After(jqt[n])

	//2020-7-6 小暑
	switch b {
	case true: //當前時間在節氣時間之後
		fmt.Printf("\t當前時間在節氣之後\n")
		fmt.Printf("\t節氣名稱:%s\n", Jmc[n]) //這裏n+1和julian.JqHs裏面的名稱對應
		jmc = Jmc[n]                      //節氣名稱
	case false: //當前時間在節氣之前或者當日即是節氣時間
		//當日和節氣同一天
		if eb := st.Equal(jqt[n]); eb == true {
			fmt.Printf("\t當前時間＝節氣時間\n")
			fmt.Printf("\t節氣名稱:%s\n", Jmc[n]) //這裏n+1和julian.JqHs裏面的名稱對應
			jmc = Jmc[n]
			//如果当前时间在当日节气时间戳之前得按照上一个节气计算那一元

		} else if tx < 0 { //当前时间在節氣之後並没有超过下一个节气时间　依然按照上一节气时间计算符头
			fmt.Printf("\t當前時間在節氣之前\n")
			fmt.Printf("\t節氣名稱:%s\n", Jmc[n]) //這裏n+1和julian.JqHs裏面的名稱對應
			if nb := st.Before(jqt[n]); nb == true {
				jmc = julian.JqHs[n] //当前时间在節氣之後並没有超过下一个节气时间　依然按照上一节气时间计算符头
			} else { //炒接置潤法的話上面的if語句可以不要？只要下面這個？
				jmc = Jmc[n] //下一個節氣名稱
			}
		}
	}
	return
}
*/
