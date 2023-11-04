// 快速入门
package main

import "fmt"

// 编写一个学生系统

// 继承部分
type Student struct{
	name string
	age int
	score int
}

// 显示成绩
func (p *Student) ShowInfo(){
	fmt.Printf("name=%v age=%v score=%v\n", p.name, p.age, p.score)
}

// 分数赋予
func (p *Student) SetScore(score int){
	p.score = score
}


// 小学生
type Pupil struct{
	Student
}

// 考试状态
func (p *Pupil) testing(){
	fmt.Println("小学生正在考试中.....")
}


// 大学生
type Graduate struct{
	Student
}

// 考试状态
func (p *Graduate) testing(){
	fmt.Println("大学生正在考试中.....")
}

func main(){

	p := &Pupil{}
	(*p).Student.name = "shi"
	(*p).Student.age = 20
	(*p).testing()
	(*p).Student.SetScore(80)
	(*p).Student.ShowInfo()

	g := &Graduate{}
	(*g).Student.name = "chuan"
	(*g).Student.age = 20
	(*g).testing()
	(*g).Student.SetScore(81)
	(*g).Student.ShowInfo()
}