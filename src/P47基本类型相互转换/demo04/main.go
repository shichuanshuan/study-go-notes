// 字符串转基本数据类型
package main
import (
	"fmt"
	"strconv"
)

func main(){
	var str string = "123"
	var num int64
	var num2 int
	var num3 float64

	// 数据类型可能只能为64位
	num, _ = strconv.ParseInt(str, 10, 64)
	// 显示转换位int类型
	num2 = int(num)
	fmt.Printf("num = %d\n",num2)

	num3, _ = strconv.ParseFloat(str,64)
	fmt.Printf("num3 = %f", num3)
}