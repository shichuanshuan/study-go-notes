// 继承的引出

package main

import "fmt"

// 编写一个学生系统

// 小学生
type Pupil struct{
	name string
	age int
	score int
}

// 显示成绩
func (p *Pupil) ShowInfo(){
	fmt.Printf("name=%v age=%v score=%v\n", p.name, p.age, p.score)
}

// 分数赋予
func (p *Pupil) SetScore(score int){
	p.score = score
}

// 考试状态
func (p *Pupil) testing(){
	fmt.Println("小学生正在考试中.....")
}


// 大学生
type Graduate struct{
	name string
	age int
	score int
}

// 显示成绩
func (p *Graduate) ShowInfo(){
	fmt.Printf("name=%v age=%v score=%v\n", p.name, p.age, p.score)
}

// 分数赋予
func (p *Graduate) SetScore(score int){
	p.score = score
}

// 考试状态
func (p *Graduate) testing(){
	fmt.Println("大学生正在考试中.....")
}

func main(){
	p := &Pupil{
		name : "shi",
		age : 19,
	}

	p.testing()
	p.SetScore(90)
	p.ShowInfo()

	g := &Graduate{
		name : "chuan",
		age : 19,
	}

	g.testing()
	g.SetScore(90)
	g.ShowInfo()
}