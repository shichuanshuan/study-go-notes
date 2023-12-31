package test

import (
	"testing"
)

/*
单元测试日志

Log 	打印日志，同时结束测试
Logf 	格式化打印日志，同时结束测试
Error 	打印错误日志，同时结束测试
Errorf 	格式化打印错误日志，同时结束测试
Fatal 	打印致命日志，同时结束测试
Fatalf 	格式化打印致命日志，同时结束测试
*/

// func TestAddUpper(t *testing.T) {
// 	// 调用
// 	res := addUpper(10)
// 	if res != 55 {
// 		t.Fatalf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
// 	}

// 	// 如果正确，输出日志
// 	t.Logf("AddUpper(10) 执行正确")
// }

// 单元测试
func TestAdd(t *testing.T) {
	res := addUpper(10)
	if res != 45 {
		t.Fatalf("test failed [%v]", res)
	}
}

func TestUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("test [%v]", res)
	}
}