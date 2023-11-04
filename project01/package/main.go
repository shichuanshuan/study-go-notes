package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	fmt.Println(string(jsonBlob[1]))
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}

// func main() {
// 	cmd := exec.Command("cmd.exe", "C", "dir")
// 	cmd.Stdin = strings.NewReader("some input")
// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("in all caps:%q\n", out.String())

// }

/*
func main() {

	// 创建一个新文件，写入内容
	// 1.打开文件
	filePath := "d:/ipconfig.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	// 及时关闭
	defer file.Close()

	cmd := exec.Command("cmd", "/C", "ipconfig")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	// 写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)

	// 把语句写进去
	writer.WriteString(string(out))

	// 因为write是带缓存，因此在调Writestring方法时，其实
	// 内容是先写道缓存的，所有需要调用Flush方法，将缓冲的数据
	// 真正写入到文件中
	writer.Flush()

	file2, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Open file err=", err)
	}

	reader := bufio.NewReader(file2)

	// 循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') // 读到一个就换行

		if err == io.EOF {
			break
		}

		// 输出内容
		fmt.Printf(str)
	}

}
*/

/*
func main() {

	// 使用ioutil.ReadFile一次性将文件读取到位
	file := "d:/test.txt"
	// 文件内容少时，可使用该方法
	content, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Printf("read file err=%v", err)
	}

	// 把读取到的内容显示到终端
	fmt.Printf("%v\n", content)
	fmt.Printf("%v", string(content))

	// const name, age = "Kim", 22
	// s := fmt.Sprintf("%s is %d years old.\n", name, age)

	io.WriteString(os.Stdout, "ipconfig") // Ignoring error for simplicity.

}
*/

/*
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
*/

/*
func test(s string, b byte) int {
	i := len(s)
	for i--; i >= 0; i-- {
		if s[i] == b {
			break
		}
	}
	return i
}

func main() {
	value := test("tcp", ':')
	fmt.Println(value)

	if value < 0 {
		switch "tcp" {
		case "tcp", "tcp4", "tcp6":
		default:
			fmt.Println("TCP")
		}
		fmt.Println("value =", value)
	}

}
*/

/*
func main(){
	// 如果按照map的key的顺序进行排序输出
	// 1. 先将map的key 放入到切片中
	// 2. 对切片排序
	// 3. 遍历切片，然后按照key来输出map的值

	num2 := make(map[int]int)

	num2[8] = 2
	num2[3] = 9
	num2[6] = 1
	num2[1] = 20

	var slice []int

	for index, _ := range num2{
		slice = append(slice, index)
	}

	sort.Ints(slice)
	fmt.Println("after ", slice)
	fmt.Println()

	for _, value := range slice{
		fmt.Printf("num2[%v] = [%v]\n", value, num2[value])
	}

}*/

// func main(){
// 	// 19) 判断字符串是否以指定的字符串开头；
// 	// strings.HasPrefix("ftp://192.168","ftp")   ture
// 	b = strings.HasPrefix("ftp://192.168","ftp")
// 	fmt.Printf("b = %v\n", b)

// 	// 20) 判断字符串是否以指定的字符串结束；
// 	// strings.HasSuffix("ftp://192.168.txt",".txt")   ture

// }

// func main(){
// 	// 7) 找查字串是否在指定的字符串中：strings.Contains("seafood", "foo") // 看字符串(seafood)中是否有 foo,有则为 ture
// 	b := strings.Contains("seafood", "foo")
// 	fmt.Printf("b = %v\n", b)

// 	// 8) 统计一个字符串有几个指定的子串：strings.Count("ceheese", "e")
// 	num := strings.Count("ceheese", "e")
// 	fmt.Printf("num = %v\n", num)
// }

// func (t Time) Hour() int

// func main(){

// 	a := Hour()
// }

// 生成随机数
// func main(){
// 	// 设置一个随机数种子
// 	rand.Seed(time.Now().Unix())

// 	// 随机生成100以内的整数
// 	n := rand.Intn(100) + 1

// 	fmt.Println(n)
// }

// 无中间变量实现两个数互换
// func main(){
// 	var a int = 10
// 	var b int = 20

// 	a = a+b
// 	b = a - b
// 	a = a - b

// 	fmt.Printf("a = %v, b = %v", a, b)
// }

// 单分支控制
// func main(){
// 	var age int

// 	fmt.Printf("输入你的年龄：")
// 	fmt.Scanln(&age)

// 	if age > 18{
// 		fmt.Printf("你要对你的行为负责\n")
// 	}
// }

// 双分支控制
// func main(){
// 	var age int

// 	fmt.Printf("输入你的年龄")
// 	fmt.Scanln(&age)

// 	if age > 18{
// 		fmt.Printf("要对自己的行为负责\n")
// 	}else{
// 		fmt.Printf("这次放过你\n")
// 	}
// }

// func main() {
// 	var i int = 0

// 	for i = 0; i < 10; i++ {
// 		fmt.Printf("shi\n")
// 	}
// }

// 打印金字塔
// 打印 5 行
// 打印 5 列
// func main() {

// 	var i int
// 	for i = 1; i <= 5; i++ {
// 		// 打印空格
// 		for k := 1; k <= 5 - i; k++ {
// 			fmt.Printf(" ")
// 		}

// 		// j表示每层打印多少个*
// 		for j := 1; j <= 2 * i - 1; j++ {
// 			fmt.Printf("*")
// 		}
// 		fmt.Printf("\n")
// 	}
// }
// func main(){
// 	var operator byte

// 	fmt.Printf("请输入字符\n")
// 	fmt.Scanln(&operator)

// 	fmt.Printf("%c\n", operator)
// 	fmt.Printf("字符大小%v\n", unsafe.Sizeof(operator))
// }
