package demo

// 压力测试/基准测试
// 命令 go test -bench="."
// 根据长宽获取面积
func GetArea(weight int, height int) int {
	return weight * height
}