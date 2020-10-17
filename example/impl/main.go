package main

import (
	"fmt"

	"github.com/sjqm"
)

func main() {
	孤虚法()
}

//日时孤虚法　已知日干支　显示当日孤虚对应的时辰
func 孤虚法() {
	dgz := "庚午" //"丙寅" //"乙丑"
	hgz := "戊辰" //"己巳" // "戊辰" //"己亥" "戊戌"
	info, gx := sjqm.SJQM孤虚法(dgz, hgz)
	fmt.Printf("%s\n", info)
	if _, ok := gx["孤"]; ok {
		fmt.Printf("孤对应地支:%s　虚对应地支:%s\n", gx["孤"], gx["虚"])
	}
}

func 六十甲子() {

}
