package stack

import (
	"errors"
	"fmt"
	"strconv"
)

/*
	使用栈实现计算器demo：
		1.准备2个栈，一个存储数字，一个存储运算符
		2.依次扫描公式，数字直接入栈;当运算符c入栈时，如果栈内无元素，则直接入职，否则：
			a.若c优先级大于栈内符号，直接入栈；
			b.否则取出栈内符号和数字栈中前2个数字进行计算,然后结果入栈、运算符丢弃，再把c入栈。
		3.经过第2步，运算符栈中必定加减在前，仅仅栈顶可能会出现一个乘除。
		4.循环出栈1个运算符和2个数字进行计算，计算后数字入栈，运算符丢弃。
			ps：因为出栈是对公式的倒叙计算，需要注意减法的特性：a-b+c=a-(b-c),a-b-c=a-(b+c)，
				把减法后的加减互换是一个正确的选择
		5.当数字栈中只剩下一个元素中，其为要求的值
*/

type calculatorStack struct {
	Top  int
	List []float64
}

func (c *calculatorStack) push(v float64) {
	c.Top++

	if c.Top == len(c.List) {
		c.List = append(c.List, v)
		return
	}

	c.List[c.Top] = v
}

//bool=false表示栈已空
func (c *calculatorStack) pop() (float64, bool) {
	if c.Top == -1 {
		return 0, false
	}
	v := c.List[c.Top]
	c.Top--

	return v, true
}

func (c *calculatorStack) list() []float64 {
	return c.List[:c.Top+1]
}

/*
	返回值：
		bType:1=数字 2=运算符
		bVal=b解析后的值，如果是运算符，则1=+，2=-，3=*，4=/
*/
func parseFormulaItems(b string) (bType int, bVal float64, err error) {
	switch b {
	case "+": // +
		return 2, 1, nil
	case "-": // -
		return 2, 2, nil
	case "*": // *
		return 2, 3, nil
	case "/": // /
		return 2, 4, nil
	default: //数字
		if v, err := strconv.Atoi(b); err == nil {
			return 1, float64(v), nil
		}
		return 0, 0, err
	}
}

/*
	比较运算符优先级，true=a优先级小于等于b
*/
func compareOperator(a, b float64) bool {
	if getPriority(a) <= getPriority(b) {
		return true
	}

	return false
}

/*
	获取优先级：
		+ - 的优先级=1
		* / 的优先级=2
*/
func getPriority(operator float64) int {
	if operator == 1 || operator == 2 {
		return 1
	}

	return 2
}

func calculate(a, b, c float64) (float64, error) {
	switch c {
	case 1: // +
		return b + a, nil
	case 2: // -
		return b - a, nil
	case 3: // *
		return b * a, nil
	case 4: // /
		if a == 0 {
			return 0, errors.New("除数不可以为0！")
		}
		return b / a, nil
	}

	return 0, nil
}

func calculateFormula(num, operator *calculatorStack, formula []string) (float64, error) {
	//fmt.Println("formula = ", formula)
	for _, v := range formula {
		//把formula中的元素解析成数字或运算符
		t, val, err := parseFormulaItems(v)
		if err != nil {
			fmt.Println("解析元素失败：", err.Error())
			return 0, err
		}
		//fmt.Println("t = ", t, ", v = ", val)

		switch t {
		case 1: //数字
			num.push(val)

		case 2: //运算符
			if operator.Top == -1 {
				operator.push(val)
				continue
			}
			if compareOperator(val, operator.List[operator.Top]) {
				a, ok := num.pop()
				if !ok {
					break
				}
				b, ok := num.pop()
				if !ok {
					break
				}
				o, ok := operator.pop()
				if !ok {
					break
				}
				ret, err := calculate(a, b, o)
				if err != nil {
					return 0, err
				}
				//计算后，把新的运算符入栈
				operator.push(val)
				//计算后，把结果入栈
				num.push(ret)
			} else {
				operator.push(val)
			}
		}
	}

	//fmt.Println("num stack = ", num.list())
	//fmt.Println("operator stack = ", operator.list())

	/*
			此时，运算符列表中可能会有多个运算符待处理，因为是栈，所以我们选择先pop再处理。
		但需要注意的是，减法运算没有“乘法结合律”这种概念，如果交换了顺序，则需要变更符号，如：
			a-b-c
		如果先算后面的b和c，则需要变成：a-b+c
	*/
	for {
		if num.Top == 0 {
			v, _ := num.pop()
			return v, nil
		}

		a, _ := num.pop()
		b, _ := num.pop()
		c, _ := operator.pop()
		d, _ := operator.pop()
		if d == 2 {
			if c == 1 {
				c = 2
			}
			if c == 2 {
				c = 1
			}
		}
		operator.push(d) //比较完记得放回去

		ret, err := calculate(a, b, c)
		if err != nil {
			return 0, err
		}

		num.push(ret)
	}
}

/*
	遍历string时，会把多位数当成多个元素输出，例如100，会输出1 0 0 三个字符，
	这是不符合需求的，所以进行修改
*/
func parseFormula(formula string) []string {
	items := make([]string, 0)

	var item string
	for k, v := range formula {
		item += string(v)

		//如果是运算符，直接加入返回值
		if item == "+" || item == "-" || item == "*" || item == "/" {
			items = append(items, item)
			item = ""
			continue
		}

		/*
			数字类型：
			1.如果是最后一个元素，加入返回值
			1.否则检测下一个元素，如果是数字元素，直接拼接字符串；如果是运算发，则停止拼接，加入返回值
		*/
		if k == len(formula)-1 {
			items = append(items, item)
			item = ""
			continue
		}

		nextItem := string(formula[k+1])
		if nextItem == "+" || nextItem == "-" || nextItem == "*" || nextItem == "/" {
			items = append(items, item)
			item = ""
			continue
		}
	}

	return items
}

func calculatorByStack(formula string) (float64, error) {
	numStack := calculatorStack{
		Top:  -1,
		List: make([]float64, 0),
	}

	operatorStack := calculatorStack{
		Top:  -1,
		List: make([]float64, 0),
	}

	return calculateFormula(&numStack, &operatorStack, parseFormula(formula))
}
