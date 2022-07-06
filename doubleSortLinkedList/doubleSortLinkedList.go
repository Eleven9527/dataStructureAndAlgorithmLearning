package doubleSortLinkedList

import "fmt"

/*
	双向有序链表demo
*/

//数据节点
type node struct {
	Val      int
	PreNode  *node
	NextNode *node
}

func initHeadNode() *node {
	return new(node)
}

func (l *node) addNode(val int) {
	//检查该数据是否已存在
	if l.checkExist(val) {
		return
	}

	n := node{
		Val: val,
	}

	currentNode := l.NextNode
	for {
		//空链表
		if currentNode == nil {
			l.NextNode = &n
			n.PreNode = l
			return
		}

		//尾节点
		if currentNode.NextNode == nil {
			n.PreNode = currentNode
			currentNode.NextNode = &n
			return
		}

		//中间节点
		if currentNode.Val < val && currentNode.NextNode.Val > val {
			n.PreNode = currentNode
			n.NextNode = currentNode.NextNode
			currentNode.NextNode.PreNode = &n
			currentNode.NextNode = &n
			return
		}

		currentNode = currentNode.NextNode
	}
}

func (l *node) removeNode(val int) {
	if !l.checkExist(val) {
		return
	}

	currentNode := l.NextNode

	for {
		//尾节点
		if currentNode.NextNode == nil {
			currentNode.PreNode.NextNode = nil
			break
		}

		//中间节点
		if currentNode.Val == val {
			currentNode.NextNode.PreNode = currentNode.PreNode
			currentNode.PreNode.NextNode = currentNode.NextNode
			break
		}

		currentNode = currentNode.NextNode
	}
}

func (l *node) list() []int {
	vals := make([]int, 0)

	currentNode := l.NextNode
	for {
		if currentNode == nil {
			return vals
		}

		vals = append(vals, currentNode.Val)

		currentNode = currentNode.NextNode
	}
}

//true=存在
func (l *node) checkExist(val int) bool {
	vals := l.list()

	for _, v := range vals {
		if v == val {
			return true
		}
	}

	return false
}

func DoubleSortLinkedListTest() {
	dataList := []int{1, 3, 5, 2, 4, 6} //待操作的数据

	l := initHeadNode()

	fmt.Println("-------------add node-------------")
	for _, v := range dataList {
		l.addNode(v)
	}
	fmt.Println("链表 = ", l.list())

	fmt.Println("--------------remove node---------")
	for _, v := range dataList {
		l.removeNode(v)
		//fmt.Println("链表 = ", l.list())
	}

}
