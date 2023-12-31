package main

import "fmt"

// 低耦合高内聚编程思想
// 优点：定义好接口，让别的变量和结构体来实现它

// 声明一个接口
type Usb interface{
	// 声明了两个没有实现的方法
	Start()
	Stop()
}


// 声明一个手机结构体
type Phone struct{

}

// 让Phone实现 Usb 接口的方法
func (p Phone)Start(){
	fmt.Println("开始手机连接...")
}

func (p Phone)Stop(){
	fmt.Println("断开手机连接...")
}


// 声明一个相机结构体
type Camera struct{

}

// 让Camer实现 Usb 接口的方法
func (p Camera)Start(){
	fmt.Println("开始相机连接...")
}

func (p Camera)Stop(){
	fmt.Println("断开相机连接...")
}


// 计算机
type Computer struct{

}

// 编写一个方法working 方法. 接收一个Usb接口类型变量
// 只要是实现了 usb接口 （所谓实现 usb 接口，就是指实现了 Usb接口声明所有方法）
// 例如 Usb 接口声明了 strat 和 stop两个实现方法，通过其他结构体就可以调用
func (c Computer) Working(usb Usb){

	// 通过 usb 接口变量来调用 Start 和 Stop 方法
	usb.Start()
	usb.Stop()
}

func main(){

	// 测试
	// 先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	fmt.Println()

	computer.Working(camera)
}