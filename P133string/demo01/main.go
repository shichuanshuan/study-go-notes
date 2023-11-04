package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	str := "hello上海"
	// 1) 字符串遍历，同时处理有中文的问题 r := []rune(str)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符 = %c\n", r[i])
	}

	// 2) 字符串转整数： n, err := strconv.Atio("12") 转换失败返回err
	// 字符串是 字母则打印出错
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("2. 转换错误", err)
	} else {
		fmt.Println("2. 转换结果是", n)
	}

	// 3) 整数转字符串  str = strconv.Itoa(12345)
	str = strconv.Itoa(123)
	fmt.Printf("3 str = %v, str = %T\n", str, str)

	// 4) 字符串转 []byte: var byte = []byte("hello go")
	// 输出的是对应的ascll
	var bytes = []byte("hello")
	fmt.Println("4.byte =", bytes)

	// 5) []byte 转 字符串 var []byte = string([]byte{97, 65})
	// ACSll值转字符
	str = string([]byte{115, 99, 115})
	fmt.Println("5.[]byte 转 字符串 str =", str)

	// 6) 10进制转 2 8 16进制。 str = strconv.FormatInt(123, 2),返回对应的进制数
	str = strconv.FormatInt(123, 2)
	fmt.Println("123对应的二进制是：", str)
	str = strconv.FormatInt(123, 16)
	fmt.Println("123对应的十六进制是：", str)

	// 7) 找查字串是否在指定的字符串中：strings.Contains("seafood", "foo") // 看字符串(seafood)中是否有 foo,有则为 ture
	b := strings.Contains("seafood", "foo")
	fmt.Printf("b = %v\n", b)

	// 8) 统计一个字符串有几个指定的子串：strings.Count("ceheese", "e")
	num := strings.Count("str ceheese str", "str")
	fmt.Printf("8. num = %v\n", num)

	// 9) 不区分大小写的字符比较:strings.EqualFold("abc", "ABC")
	b = strings.EqualFold("abc", "ABC")
	fmt.Printf("b = %v\n", b)

	// 10) 返回子串在字符串第一次出现的index值，如果没有返回-1
	// strings.Index("NLT_abc", "abc")
	index := strings.Index("NLT_abc", "abc")
	fmt.Printf("10. index = %v\n", index)

	// 11) 返回子串在字符串最后一次出现的index值，如果没有返回-1
	// strings.LastIndex("go golang", "go")
	index = strings.LastIndex("go golang", "go")
	fmt.Printf("11. index = %v\n", index)

	// 12) 将指定的子串替换成 另一个子串：strings.Replace("go go hello", "go", "go语言", n)
	// n可以指定你希望替换几个，如果n为 -1 表示全部替换
	str2 := "go go hello"
	str = strings.Replace(str2, "go", "go语言", 1)
	fmt.Printf("str = %v, str2 = %v\n", str, str2)

	// 13) 按照某个字符，为分割标识，将一个字符拆分成字符串数组；
	// strings.Split("hello,wrold,ok", ",")
	strArr := strings.Split("hello,wrold,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("14. str[%v] = %v\n", i, strArr[i])
	}
	fmt.Printf("14 strArr= %v\n", strArr)

	// 14) 将字符串的字母进行大小写转换
	// strings.ToLower("Go") // go
	// strings.ToUpper("Go") // GO
	str = "goLang Hello"
	str = strings.ToLower(str)
	str = strings.ToUpper(str)
	fmt.Println("str = ", str)

	// 15) 将字符串左右两边的空格去掉
	// strings.TrimSpace(" tn a lone gopher ntrn  ")
	str = strings.TrimSpace(" tn a lone gopher ntrn  ")
	fmt.Printf("15 str =%q\n", str)

	// 16) 将字符串左右两边指定的字符去掉
	// strings.Trim("! hello! ", " !")
	str = strings.Trim("hello! ", " !")
	fmt.Printf("16 str = %q\n", str)

	// 17)
	// 18)

	// 19) 判断字符串是否以指定的字符串开头；
	// strings.HasPrefix("ftp://192.168","ftp")   ture
	b = strings.HasPrefix("ftp://192.168", "ftp")
	fmt.Printf("b = %v\n", b)

	// 20) 判断字符串是否以指定的字符串结束；
	// strings.HasSuffix("ftp://192.168","ftp")   ture

	s := []string{"foo", "bar", "baz"}
	fmt.Println("test ", strings.Join(s, ", "))

	fmt.Println(strings.Replace("oink onk oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	var letters = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	fmt.Printf("%v\n", letters)

}
