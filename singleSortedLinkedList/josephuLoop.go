package singleSortedLinkedList

/*
	单向链表-约瑟夫环demo
*/

/*
	有10个人依次报数，报到3的出列加入新链表，后续从1继续开始报数，直接全部出列，求新组成的链表
*/

type Node struct {
	Id   int
	Next *Node
}

func (n *Node) addSortedNode(val int) {
	node := Node{Id: val}
	temp := &Node{
		Next: n,
	}

	//add node
	for {
		//tail node
		if temp.Next.Next == nil {
			temp.Next.Next = &node
			break
		}

		//middle node
		if temp.Next.Id < val && temp.Next.Next.Id > val {
			node.Next = temp.Next.Next
			temp.Next.Next = &node
			break
		}

		temp = temp.Next
	}

}

func (n *Node) addNonSortedNode(val int) {
	node := Node{Id: val}
	temp := &Node{
		Next: n,
	}

	//add node
	for {
		//tail node
		if temp.Next.Next == nil {
			temp.Next.Next = &node
			break
		}

		temp = temp.Next
	}
}

func (n *Node) list() []int {
	vals := make([]int, 0)
	vals = append(vals, n.Id)

	temp := n.Next

	for {
		if temp == nil {
			break
		}

		vals = append(vals, temp.Id)

		temp = temp.Next
	}

	return vals
}

func (n *Node) popNode() []*Node {
	num := 0
	nodes := make([]*Node, 0)

	temp := &Node{
		Next: n,
	}
	for {
		//如果只剩最后一个节点了，直接加入列表，并结束循环
		if temp.Next.Next == temp.Next {
			nodes = append(nodes, temp.Next)
			break
		}

		num++

		if num%3 == 0 {
			nodes = append(nodes, temp.Next)
			temp.Next = temp.Next.Next
			num = 0
			continue
		}

		temp = temp.Next
	}

	return nodes
}

func JosephuLoop(vals []int) []int {
	//build old linkedList
	old := &Node{
		Id:   vals[0],
		Next: nil,
	}

	for i := 1; i < len(vals); i++ {
		old.addSortedNode(vals[i])
	}
	//fmt.Println("old = ", old.list())

	//tail link head
	temp := old
	for {
		if temp.Next == nil {
			temp.Next = old
			break
		}

		temp = temp.Next
	}

	//build new linkedList
	nodes := old.popNode()
	newLinkedList := &Node{
		Id: nodes[0].Id,
	}

	for i := 1; i < len(nodes); i++ {
		newLinkedList.addNonSortedNode(nodes[i].Id)
	}

	return newLinkedList.list()
}
