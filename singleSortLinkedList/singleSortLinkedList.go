package singleSortLinkedList

/*
	单向有序链表demo
*/

import (
	"fmt"
	"strconv"
)

//数据节点
type node struct {
	Val      int   //节点存储的值
	NextNode *node //下一个节点的地址
}

//头节点，用来表示整条链表
type headNode struct {
	NodeCount int   //链表中节点个数
	Next      *node //下一个节点的位置
}

//初始化头节点
func initLinkedList() *headNode {
	return new(headNode)
}

func (l *headNode) addNode(val int) {
	//检查该数据是否已存在
	if l.checkExist(val) {
		return
	}

	node := node{Val: val}

	//添加数据
	temp := l.Next
	for {
		//空链表，直接添加数据
		if temp == nil {
			l.Next = &node
			break
		}

		//已到链表尾部，直接添加数据
		if temp.NextNode == nil {
			temp.NextNode = &node
			break
		}

		//逐一比较相邻节点，若待插入数字在它们之间，则在他们之间新增节点
		if temp.Val < val && temp.NextNode.Val > val {
			node.NextNode = temp.NextNode
			temp.NextNode = &node
			break
		}

		temp = temp.NextNode
	}

	//节点个数+1
	l.NodeCount++
}

//true=存在
func (l *headNode) checkExist(val int) bool {
	vals := l.list()

	for _, v := range vals {
		if v == val {
			return true
		}
	}

	return false
}

func (l *headNode) list() []int {
	vals := make([]int, 0)

	//空链表
	//if l.NodeCount == 0 {
	//	return vals
	//}

	temp := l.Next

	for {
		if temp == nil {
			break
		}

		vals = append(vals, temp.Val)

		temp = temp.NextNode
	}

	return vals
}

func (l *headNode) removeNode(val int) {
	//空链表
	if l.NodeCount == 0 {
		return
	}

	//检查链表中是否存在该数据
	if !l.checkExist(val) {
		return
	}

	//remove
	temp := l.Next

	for {
		if temp.Val == val {
			//第一个节点
			if temp == l.Next {
				l.Next = temp.NextNode
				return
			}

			//尾节点
			if temp.NextNode == nil {
				temp = nil
				return
			}
		}

		//中间节点
		if temp.NextNode.Val == val {
			temp.NextNode = temp.NextNode.NextNode
			return
		}

		temp = temp.NextNode
	}
}

func SingleSortLinkedListTest() {
	dataList := []int{1, 3, 5, 2, 4, 6} //待操作的数据

	l := initLinkedList()

	fmt.Println("-------------------add node-----------------")
	for _, v := range dataList {
		l.addNode(v)
	}
	fmt.Println("链表node = ", l.list())

	fmt.Println("-------------------remove node--------------")
	l.removeNode(dataList[0])
	fmt.Println("链表node = ", l.list())
	l.removeNode(dataList[1])
	fmt.Println("链表node = ", l.list())
	l.removeNode(dataList[2])
	fmt.Println("链表node = ", l.list())
	l.removeNode(dataList[3])
	fmt.Println("链表node = ", l.list())
	l.removeNode(dataList[4])
	fmt.Println("链表node = ", l.list())
	l.removeNode(dataList[5])
	fmt.Println("链表node = ", l.list())
}

// InterviewQuestions 面试题
func InterviewQuestions() {
	//求单链表中节点个数
	dataList := []int{1, 3, 5, 2, 4, 6} //待操作的数据

	l := initLinkedList()

	for _, v := range dataList {
		l.addNode(v)
	}
	fmt.Println("链表 = ", l.list())
	fmt.Println("链表节点个数 = ", l.getNodeCount())

	//查找倒数第k个节点
	index := 2
	val := l.getNodeByIndex(index)
	fmt.Println("倒数第"+strconv.Itoa(index)+"个节点 = ", val)

	//单链表的反转
	fmt.Println("反转后的链表 = ", l.reverse().list())
}

func (l *headNode) getNodeCount() int {
	var count int

	temp := l.Next

	for {
		if temp == nil {
			return count
		}

		count++

		temp = temp.NextNode
	}
}

/*
	若某个节点后面第k-1个节点是尾节点，则其为答案
*/
func (l *headNode) getNodeByIndex(k int) int {
	tempTarget := l.Next

	for {
		tempTail := tempTarget

		for i := 0; i < k-1; i++ {
			tempTail = tempTail.NextNode
		}

		if tempTail.NextNode == nil {
			return tempTarget.Val
		}

		tempTarget = tempTarget.NextNode
	}
}

func (l *headNode) reverse() *headNode {
	//依次取出链表中的数据
	vals := make([]int, 0)

	temp := l.Next
	for {
		if temp == nil {
			break
		}
		vals = append(vals, temp.Val)

		temp = temp.NextNode
	}
	//fmt.Println("vals = ", vals)

	//反向插入新的链表
	ll := initLinkedList()
	temp = ll.Next
	for i := len(vals) - 1; i >= 0; i-- {
		n := node{
			Val: vals[i],
		}

		//空链表
		if temp == nil {
			ll.Next = &n
			temp = &n
			continue
		}

		temp.NextNode = &n

		temp = temp.NextNode
	}

	return ll
}
