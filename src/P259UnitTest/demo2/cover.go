package demo2

// 覆盖率测试
// 命令 : go test -cover
// 覆盖率测试能知道测试程序总共覆盖了多少业务代码（也就是 demo_test.go 中测试了多少 demo.go 中的代码），可以的话最好是覆盖100%。
func GetArea(weight int, height int) int {
	return weight * height
}
