package sjqm

import (
	"time"
)

//根据当前阳历时间判断阴阳遁
//时间戳精确到时
//now:当前阳历时间戳　要匹配冬至和夏至时间戳　精确到小时
//dzt:冬至时间戳　xzt:夏至时间戳
func YY阴阳判断(now, dzt, xzt time.Time) (yy string) {

	//dzb := now.After(dzt) //当前时间在冬至之后
	xzb := now.After(xzt)
	eq := now.Equal(xzt)

	if xzb == true || eq == true {
		yy = "阴"
	} else {
		yy = "阳"
	}

	return
}
