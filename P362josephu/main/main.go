package main

import "fmt"

// 小孩结构体
type Boy struct {
	no   int  // 编号
	next *Boy // 指向下一个小孩的指针
}

// 编写一个函数，构成单向的环形链表
// num: 表示小孩的个数
// *Boy: 返回该环形的链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{}  // 空结点
	curBoy := &Boy{} // 空结点

	// 判断
	if num < 1 {
		fmt.Println("num 的值不对")
		return first
	}
	// 循环的构建这个环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			no: i,
		}
		// 分析构成循环链表，需要一个辅助指针
		// 1. 因为第一个小孩比较特殊
		if i == 1 { // 第一个小孩
			first = boy // 不要动
			curBoy = boy
			curBoy.next = first
		} else {
			curBoy.next = boy
			curBoy = boy
			curBoy.next = first // 构造环形链表
		}
	}

	return first
}

// 显示单向的环形链表
func ShowBoy(first *Boy) {
	// 处理一下如果环形链表为空
	if first.next == nil {
		fmt.Println("链表为空，没有小孩...")
		return
	}

	// 创建一个指针，帮助遍历
	curBoy := first
	for {
		fmt.Printf("小孩编号=%d ->", curBoy.no)
		// 退出的条件
		if curBoy.next == first {
			break
		}

		curBoy = curBoy.next
	}
}

/*
设编号为1, 2, ...n的n 个人围坐一圈，约定编号为k (1 <= k <= n)
的人从 1 开始报数，数到 m 的那个人出列， 它的下一位又从 1 开始报数，
数到 m 的那个人又出列， 依此类推，直到所有人出列为止，由此产生一个出队编号的序列
*/

// 分析思路
// 1. 编写一个函数， PlayGame (first *Boy, startNo int, countNum int)
// 2. 最后我们使用一个算法，按照要求，在环形链表中留下最后一个人

func PlayGame(first *Boy, startNo int, countNum int) {
	// 1. 空的链表我们单独的处理
	if first.next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	// 留一个，判断 startNo <= 小孩的总数
	// 2. 需要定义辅助指针，帮助我们删除小孩
	tail := first
	// 3. 让tail 执行环形链表的最后一个小孩，这个非常重要
	// 因为tail 在删除小孩时需要用到。
	for {
		if tail.next == first { // 说明tail到了最后的小孩
			break
		}
		tail = tail.next
	}

	// 4. 让first 移动到 startNo [后面我们删除小孩，就以firsta 为准]
	for i := 1; i <= startNo-1; i++ {
		first = first.next
		tail = tail.next
	}
	fmt.Println()

	// 5. 开始数 countNum,然后就删除 first 指向的小孩
	for {
		// 开始数 countNum-1 次
		for i := 1; i <= countNum-1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("小孩编号为%d 出圈 \n", first.no)
		// 删除 first 执行的小孩
		first = first.next
		tail.next = first
		// 判断如果 tail == first 圈子中只有一个小孩
		if tail == first {
			break
		}
	}
	fmt.Printf("小孩编号为%d 出圈 ->", first.no)
}

func main() {
	first := AddBoy(5)

	ShowBoy(first)
	PlayGame(first, 4, 3)
}
