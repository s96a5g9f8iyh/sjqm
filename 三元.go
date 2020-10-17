package sjqm

import (
	"strings"
)

//三元
type SY三元 struct {
	Name   string //节气名称
	Number []int  //三元数字
	BaGong int    //所属宫位数字
}

func NewSY() *SY三元 {
	sy := new(SY三元)
	return sy
}

//当前节气所在的局数字
//name:节气名称 number:局数子　bgn:当前节气的八宫数字
func (sy *SY三元) J局信息(jmc string, yuan string) (name string, number, bgn int) {
	switch jmc {
	case Jmc[0]: //冬至
		name, number, bgn = sy.DS大暑(jmc, yuan)
	case Jmc[1]: //小寒
		name, number, bgn = sy.Xh小寒(jmc, yuan)
	case Jmc[2]: //大寒
		name, number, bgn = sy.Dh大寒(jmc, yuan)
	case Jmc[3]:
		name, number, bgn = sy.Lc立春(jmc, yuan)
	case Jmc[4]:
		name, number, bgn = sy.Ys雨水(jmc, yuan)
	case Jmc[5]:
		name, number, bgn = sy.Jz惊蛰(jmc, yuan)
	case Jmc[6]:
		name, number, bgn = sy.Cf春分(jmc, yuan)
	case Jmc[7]:
		name, number, bgn = sy.Qm清明(jmc, yuan)
	case Jmc[8]:
		name, number, bgn = sy.Gy谷雨(jmc, yuan)
	case Jmc[9]:
		name, number, bgn = sy.Lx立夏(jmc, yuan)
	case Jmc[10]:
		name, number, bgn = sy.Xm小满(jmc, yuan)
	case Jmc[11]:
		name, number, bgn = sy.Mz忙钟(jmc, yuan)
	case Jmc[12]:
		name, number, bgn = sy.Xz夏至(jmc, yuan)
	case Jmc[13]:
		name, number, bgn = sy.Xs小暑(jmc, yuan)
	case Jmc[14]:
		name, number, bgn = sy.DS大暑(jmc, yuan)
	case Jmc[15]:
		name, number, bgn = sy.Lq立秋(jmc, yuan)
	case Jmc[16]:
		name, number, bgn = sy.Cs处暑(jmc, yuan)
	case Jmc[17]:
		name, number, bgn = sy.Bl白露(jmc, yuan)
	case Jmc[18]:
		name, number, bgn = sy.Qf秋分(jmc, yuan)
	case Jmc[19]:
		name, number, bgn = sy.Hl寒露(jmc, yuan)
	case Jmc[20]:
		name, number, bgn = sy.Sj霜降(jmc, yuan)
	case Jmc[21]:
		name, number, bgn = sy.Ld立冬(jmc, yuan)
	case Jmc[22]: //大暑
		name, number, bgn = sy.Xx小雪(jmc, yuan)
	case Jmc[23]:
		name, number, bgn = sy.Dx大雪(jmc, yuan)
	}
	return
}

/////////////////////
//冬至　坎
func (sy *SY三元) Dz冬至(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[0]      //节气名称
	n := []int{1, 7, 4} //节气的上中下三元数字
	bg := 1             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//小寒　坎
func (sy *SY三元) Xh小寒(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[1]      //节气名称
	n := []int{2, 8, 5} //节气的上中下三元数字
	bg := 1             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//大寒　坎
func (sy *SY三元) Dh大寒(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[2]      //节气名称
	n := []int{3, 9, 6} //节气的上中下三元数字
	bg := 1             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

////////////////////////
//立春　艮
func (sy *SY三元) Lc立春(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[3]      //节气名称
	n := []int{8, 5, 2} //节气的上中下三元数字
	bg := 8             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//雨水
func (sy *SY三元) Ys雨水(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[4]      //节气名称
	n := []int{9, 6, 3} //节气的上中下三元数字
	bg := 8             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//惊蛰
func (sy *SY三元) Jz惊蛰(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[5]      //节气名称
	n := []int{1, 7, 4} //节气的上中下三元数字
	bg := 8             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

/////////////////////////
//春分　震
func (sy *SY三元) Cf春分(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[6]      //节气名称
	n := []int{3, 9, 6} //节气的上中下三元数字
	bg := 3             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//清明　震
func (sy *SY三元) Qm清明(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[7]      //节气名称
	n := []int{4, 1, 7} //节气的上中下三元数字
	bg := 3             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//谷雨 震３宫
func (sy *SY三元) Gy谷雨(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[8]      //节气名称
	n := []int{5, 2, 8} //节气的上中下三元数字
	bg := 3             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

////////////
//立夏 巽４宫
func (sy *SY三元) Lx立夏(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[9]      //节气名称
	n := []int{4, 1, 7} //节气的上中下三元数字
	bg := 4             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//小满
func (sy *SY三元) Xm小满(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[10]     //节气名称
	n := []int{5, 2, 8} //节气的上中下三元数字
	bg := 4             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//忙钟
func (sy *SY三元) Mz忙钟(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[11]     //节气名称
	n := []int{6, 3, 9} //节气的上中下三元数字
	bg := 4             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//////////////
//夏至 离９宫
func (sy *SY三元) Xz夏至(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[12]     //节气名称
	n := []int{9, 3, 6} //节气的上中下三元数字
	bg := 9             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//小暑
func (sy *SY三元) Xs小暑(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[13]     //节气名称
	n := []int{8, 2, 5} //节气的上中下三元数字
	bg := 9             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//大暑
func (sy *SY三元) DS大暑(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[14]     //节气名称
	n := []int{7, 1, 4} //节气的上中下三元数字
	bg := 9             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

/////////////////////
//立秋
func (sy *SY三元) Lq立秋(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[15]     //节气名称
	n := []int{2, 5, 8} //节气的上中下三元数字
	bg := 2             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//处暑
func (sy *SY三元) Cs处暑(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[16]     //节气名称
	n := []int{1, 4, 7} //节气的上中下三元数字
	bg := 2             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//白露
func (sy *SY三元) Bl白露(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[17]     //节气名称
	n := []int{9, 3, 6} //节气的上中下三元数字
	bg := 2             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

////////////////////////////
//秋分　兑７宫
func (sy *SY三元) Qf秋分(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[18]     //节气名称
	n := []int{7, 1, 4} //节气的上中下三元数字
	bg := 7             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//寒露
func (sy *SY三元) Hl寒露(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[19]     //节气名称
	n := []int{6, 9, 3} //节气的上中下三元数字
	bg := 7             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//霜降
func (sy *SY三元) Sj霜降(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[20]     //节气名称
	n := []int{5, 8, 2} //节气的上中下三元数字
	bg := 7             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

/////////////////////
//立冬　乾６宫
func (sy *SY三元) Ld立冬(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[21]     //节气名称
	n := []int{6, 9, 3} //节气的上中下三元数字
	bg := 6             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//小雪
func (sy *SY三元) Xx小雪(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[22]     //节气名称
	n := []int{5, 8, 2} //节气的上中下三元数字
	bg := 6             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//大雪
func (sy *SY三元) Dx大雪(jmc string, yuan string) (jn string, number, bgn int) {
	name := Jmc[23]     //节气名称
	n := []int{4, 7, 1} //节气的上中下三元数字
	bg := 6             //8宫数字
	sy = &SY三元{
		Name:   name,
		Number: n,
		BaGong: bg,
	}
	if b := strings.Contains(name, jmc); b == true {
		x := ConvS2N(yuan)
		jn, number, bgn = infoJ(x, n, name)
	}
	return
}

//jn:节气名称　number:局数字　bgn:当前节气对应的八宫数字
func infoJ(x int, n []int, name string) (jn string, number, bgn int) {
	switch x {
	case 0: //上元
		jn = name
		number = n[0]
		bgn = 9
	case 1: //中元
		jn = name
		number = n[1]
		bgn = 9
	case 2: //下元
		jn = name
		number = n[2]
		bgn = 9
	}
	return
}
