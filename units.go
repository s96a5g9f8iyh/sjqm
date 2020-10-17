package sjqm

//轉換上中下三元字符串爲對應三元結構體的數字
func ConvS2N(yuan string) (n int) {
	switch yuan {
	case "上元":
		n = 0
	case "中元":
		n = 1
	case "下元":
		n = 2
	}
	return
}
