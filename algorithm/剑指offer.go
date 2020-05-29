package algorithm

import (
	"math"
)

func FindRepeatNumber(nums []int) int {
	//思路1：看代码
	/*m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		} else {
			m[v] = 1
		}
	}
	return 0*/

	//思路2：原地置换。假设这个数组中没有重复的元素，那么从小到大排列的话，下标跟值应该是相等的。如果不相等，则跟对应的值调换
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			temp := nums[i]
			nums[i] = nums[temp]
			nums[temp] = temp
		}
	}
	return 0
}

/**
二维数组中的查找
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

*/
func FindNumberIn2DArray(matrix [][]int, target int) bool {
	//从二维数组的第一行最右侧开始找。如果找到则返回，如果当前值比目标值大，则查找列往左侧移动一列。
	//如果当前值比目标值小，则往下移动一行。
	lr := len(matrix)
	if lr == 0 {
		return false
	}
	lc := len(matrix[0])
	if lc == 0 {
		return false
	}
	i := 0
	j := lc - 1
	for i < lr && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

/**
替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
*/
func ReplaceSpace(s string) string {
	//思路1，从前往后扫描。这种由于每次遇到一个空格需要将后面的元素整体往后移动三位，所以时间复杂度不太好
	//思路2，从后往前扫描，获取到字符串中空格的数量，并计算出替换后的字符串长度。设定两个指针p1,p2，分别指向
	//替换前的字符串末尾和替换后的字符串末尾。依次移动p1，如果当前的字符不是空格，则将当前字符放入p2中。
	//如果是空格，那么p1往前移动1位，p2放入%20并往前移动三位。
	ls := len(s)
	ln := ls
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ln += 2
		}
	}
	p1 := ls - 1
	p2 := ln - 1
	//如果此时两个长度相等，那么证明其中没空格，直接返回s
	if p1 == p2 {
		return s
	}
	s2 := make([]byte, ln)
	for p1 >= 0 {
		if s[p1] == ' ' {
			s2[p2] = '0'
			s2[p2-1] = '2'
			s2[p2-2] = '%'
			p1--
			p2 -= 3
		} else {
			//将s[p1]放到第二个字符串的末尾，并将p1和p2分别-1
			s2[p2] = s[p1]
			p1--
			p2--
		}
	}
	return string(s2)
}

/**
从尾到头打印链表
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
输入：head = [1,3,2]
输出：[2,3,1]
*/
func reversePrint(head *ListNode) []int {
	//思路1，依次遍历链表，然后将数组反向输出
	cur := make([]int, 0)
	res := make([]int, 0)
	for head != nil {
		cur = append(cur, head.Val)
		head = head.Next
	}
	l := len(cur)
	if l > 0 {
		for i := l - 1; i >= 0; i-- {
			res = append(res, cur[i])
		}
	}
	return res
}

/**
重建二叉树
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7


*/
func BuildTree3(preorder []int, inorder []int) *TreeNode {
	//使用递归。
	//前序遍历的第一个值必为当前数的根节点。找到这个值在中序遍历的位置，那么中序遍历中这个值的左侧即为它的左子树，右侧即为它的右子树
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	var mid int
	for k, v := range inorder {
		if v == preorder[0] {
			mid = k
			break
		}
	}
	root.Left = BuildTree3(preorder[1:mid+1], inorder[0:mid])
	root.Right = BuildTree3(preorder[mid+1:], inorder[mid+1:])
	return root
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

/**
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
*/
type CQueue struct {
	Stack1 []int
	Stack2 []int
}

func ConstructorCQueue() CQueue {
	stack1 := make([]int, 0)
	stack2 := make([]int, 0)
	return CQueue{Stack1: stack1, Stack2: stack2}
}

func (this *CQueue) AppendTail(value int) {
	//用栈模拟队列，首先考虑插入，插入的话，就直接往第一个栈里放入元素。
	this.Stack1 = append(this.Stack1, value)
}

func (this *CQueue) DeleteHead() int {
	//出队列，判断第二个栈中是否有元素，如果有，直接出栈弹出。如果没有，则将stack1中的元素出栈依次压入stack2中然后再弹出
	if len(this.Stack2) == 0 {
		for len(this.Stack1) > 0 {
			l := len(this.Stack1)
			last := this.Stack1[l-1]
			this.Stack1 = this.Stack1[:l-1]
			this.Stack2 = append(this.Stack2, last)
		}
	}
	l2 := len(this.Stack2)
	if l2 == 0 {
		return -1
	}
	res := this.Stack2[l2-1]
	this.Stack2 = this.Stack2[:l2-1]
	return res
}

/**
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下：
*/
func Fib(n int) int {
	//递归有重复计算，直接使用循环
	if n == 0 || n == 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}

/**
青蛙跳台阶问题
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
*/
func numWays(n int) int {
	//青蛙跳台阶本质上还是斐波那契数列 青蛙跳上1级台阶只有一种方式，跳上2级台阶有两种（1+1 2）
	//跳上3级台阶有3种（3*1 2+1 1+2） 4级台阶有5种(4*1 2*2 1+1+2 2+1+1 1+2+1) 依次类推
	if n == 0 { //这一步很奇怪，但是leecode在n为0的时候就是为1的。
		return 1
	}
	if n <= 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}

/**
旋转组的最小数字

把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
*/
func minArray(numbers []int) int {
	//思路1 遍历数组，保存将当前值和上一个值对比。如果当前值小于上一个值，那么其就为最小值
	/*if len(numbers) == 0 {
		return 0
	}
	var last = numbers[0]
	var res = numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < last {
			res = numbers[i]
			break
		} else {
			last = numbers[i]
		}
	}
	return res*/

	//思路2 二分查找
	var i, j = 0, len(numbers) - 1
	for i < j {
		m := (i + j) / 2
		if numbers[m] > numbers[j] {
			i = m + 1
		} else if numbers[m] < numbers[j] {
			j = m
		} else {
			j--
		}
	}
	return numbers[i]
}

/**
矩阵中的路径

请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，
每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。

*/
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i int, j int, word string, k int) bool {
	if board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	temp := board[i][j]
	board[i][j] = byte(' ')
	if 0 <= i-1 && dfs(board, i-1, j, word, k+1) {
		return true
	}
	if i+1 < len(board) && dfs(board, i+1, j, word, k+1) {
		return true
	}
	if 0 <= j-1 && dfs(board, i, j-1, word, k+1) {
		return true
	}
	if j+1 < len(board[0]) && dfs(board, i, j+1, word, k+1) {
		return true
	}
	board[i][j] = temp
	return false
}

/**
机器人的运动范围
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），
也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，
因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
*/
//此为记录当前节点是否走过的二维数组
var visited [][]bool

func movingCountDfs(m int, n int, k int) int {
	//dfs 深度优先。
	visited = make([][]bool, m)
	for k, _ := range visited {
		visited[k] = make([]bool, n)
	}
	return dfsForMovingCount(0, 0, 0, 0, m, n, k)
}

/**
si为左侧坐标的数位和
sj为右侧坐标的数位和
*/
func dfsForMovingCount(i, j, si, sj, m, n, k int) int {
	if i >= m || j >= n || k < si+sj || visited[i][j] {
		return 0
	}
	visited[i][j] = true
	//计算往下移动时，下一个横坐标的数位和
	var nsi int
	if (i+1)%10 != 0 {
		nsi = si + 1
	} else {
		nsi = si - 8
	}
	//计算往右移动时，下一个纵坐标的数位和
	var nsj int
	if (j+1)%10 != 0 {
		nsj = sj + 1
	} else {
		nsj = sj - 8
	}
	return 1 + dfsForMovingCount(i+1, j, nsi, sj, m, n, k) + dfsForMovingCount(i, j+1, si, nsj, m, n, k)
}

func movingCountBfs(m int, n int, k int) int {
	//bfs 广度优先。
	visited = make([][]bool, m)
	for k, _ := range visited {
		visited[k] = make([]bool, n)
	}
	var res int
	var queue = make([][]int, 0)
	queue = append(queue, []int{0, 0, 0, 0})
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		//判断当前节点是否为可以走的节点。
		i := cur[0]
		j := cur[1]
		si := cur[2]
		sj := cur[3]

		if i >= m || j >= n || k < si+sj || visited[i][j] {
			continue
		}
		visited[i][j] = true
		res++
		//计算往下移动时，下一个横坐标的数位和
		var nsi int
		if (i+1)%10 != 0 {
			nsi = si + 1
		} else {
			nsi = si - 8
		}
		//计算往右移动时，下一个纵坐标的数位和
		var nsj int
		if (j+1)%10 != 0 {
			nsj = sj + 1
		} else {
			nsj = sj - 8
		}
		queue = append(queue, []int{i + 1, j, nsi, sj})
		queue = append(queue, []int{i, j + 1, si, nsj})
	}
	return res
}

/**
剪绳子
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1）
每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
*/
func CuttingRope(n int) int {
	//思路1 动态规划
	/*if n < 4 {
		return n - 1
	}
	var max = func(x, y int) int {
		if x >= y {
			return x
		} else {
			return y
		}
	}

	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	dp[3] = 2

	for i := 4; i <= n; i++ {
		for j := 1; j < i; j++ {
			//(i-j)*j表示在j处剪一刀，然后两边相乘的值
			//j*dp[i-j]表示在j处剪一刀之后，跟i-j处的最大值相乘
			dp[i] = max(dp[i], max((i-j)*j, j*dp[i-j]))
		}
	}
	return dp[n]*/

	//思路2 贪心算法，尽可能分解出多的3，最后分不出的情况下，一定会余0或者1或者2
	//如果是余0，那么不管，结果是pwd(3, n)，如果是余1，那么将最后一个3取出，整体乘积变成pow(3, n-1) * 3 * 1 由于3和1在一起一定可以变成2*2
	//那么最大值一定是pow(3, n-1) * 2 * 2   如果最后是余2，那么直接pow(3, n) * 2 即可
	if n < 4 {
		return n - 1
	}
	i := n / 3
	j := n % 3
	if j == 0 {
		return int(math.Pow(float64(3), float64(i)))
	}
	if j == 1 {
		return int(math.Pow(float64(3), float64(i-1))) * 4
	}
	if j == 2 {
		return int(math.Pow(float64(3), float64(i))) * 2
	}
	return 0
}

//上一题的大数版本
func cuttingRopeForBig(n int) int {
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	var pow3 = func(n int) int {
		res := 1
		for i := 0; i < n; i++ {
			res = (res * 3) % 1000000007
		}
		return res
	}

	if n%3 == 0 {
		return pow3(n / 3)
	} else if n%3 == 1 {
		return pow3((n-3)/3) * 4 % 1000000007
	} else if n%3 == 2 {
		return pow3((n-2)/3) * 2 % 1000000007
	}
	return 0
}

/**
二进制中1的个数
请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数。例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
示例 1：

输入：00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。
示例 2：

输入：00000000000000000000000010000000
输出：1
解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。

*/
func hammingWeight(num uint32) int {
	//如果n&1为1，则证明n的最后一位是1。如果n&1为0，证明最后一位是0
	var res uint32
	for num > 0 {
		res += num & 1
		num = num >> 1
	}
	return int(res)
}

/**
数值的整数次方
实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。

示例 1:
输入: 2.00000, 10
输出: 1024.00000
示例 2:
输入: 2.10000, 3
输出: 9.26100

-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [pwd(-2, 31), pow(2, 31)-1] 。
*/
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	power := 1.0
	if n > 0 {
		//例如此刻计算1999，那么将1999分为2个999和1个1，那么1999最终即为pow(x,999) * pow(x,999) * pow(x,1)
		power = myPow(x, n/2)
		return power * power * myPow(x, n%2)
	}
	if n < 0 {
		m := -n
		power = myPow(x, m/2)
		power = power * power * myPow(x, m%2)
		return 1 / power
	}
	return power
}

/**
打印从1到最大的n位数
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:
输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]


*/
func PrintNumbers(n int) []int {
	if n == 0 {
		return []int{}
	}
	max := int(math.Pow(10, float64(n)))
	res := make([]int, max-1)
	for i := 0; i < max-1; i++ {
		res[i] = i + 1
	}
	return res
}

/**
删除链表的节点
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
返回删除后的链表的头节点。
注意：此题对比原题有改动

示例 1:
输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
示例 2:
输入: head = [4,5,1,9], val = 1
输出: [4,5,9]
解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

*/
func DeleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	root := &ListNode{Next: head}
	preNode := root
	for preNode != nil {
		if preNode.Next != nil && preNode.Next.Val == val {
			if preNode.Next.Next == nil {
				preNode.Next = nil
			} else {
				preNode.Next = preNode.Next.Next
			}
		} else {
			preNode = preNode.Next
		}
	}
	return root.Next
}

/**
调整数组顺序使奇数位于偶数前面
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
示例：
输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。
*/
func Exchange(nums []int) []int {
	//思路1，遍历数组，将偶数往数组的后方放
	//思路2，双指针p1p2，分别指向数组的头部跟尾部。首先判断p1是否为偶数，如果是，则移动p2，直到p2找到第一个奇数位置，交换两者的位置
	p1 := 0
	p2 := len(nums) - 1
	for p2 > p1 {
		if nums[p1]%2 == 0 {
			//如果p1为偶数，那么看下p2是否为奇数
			if nums[p2]%2 == 1 {
				//交换两者的值
				nums[p1], nums[p2] = nums[p2], nums[p1]
				p2--
				p1++
			} else {
				//移动p2
				p2--
			}
		} else {
			p1++
		}
	}
	return nums
}

/**
链表中倒数第k个节点
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。
*/
func GetKthFromEnd(head *ListNode, k int) *ListNode {
	//思路：双指针p1p2，都指向链表头部。p1先往后移动k个值，然后p1p2一起往后移动
	//等到p1移动到底部时，p2即为倒数第k个节点
	p1, p2 := head, head
	for i := 1; i <= k; i++ {
		if p1 == nil {
			return nil
		}
		p1 = p1.Next
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}

/**
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。



示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

*/
func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var p *ListNode
	var tmp *ListNode
	/**
			1-2-3-4-5
	    ->  1  2-3-4-5
	    ->  2-1 3-4-5
	    ->  3-2-1 4-5
	    ->  4-3-2-1 5
	    ->  5-4-3-2-1
	*/
	for head != nil {
		//将当前节点的后面部分全部放入一个临时变量汇总
		tmp = head.Next
		//将当前节点的Next指向p
		head.Next = p
		//将p置为当前节点的链表
		p = head
		//遍历下一个节点
		head = tmp
	}
	return p
}

/**
合并两个排序的链表
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
示例1：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：
0 <= 链表长度 <= 1000

*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	var x, y int

	for l1 != nil {
		x = l1.Val
		if l2 != nil {
			y = l2.Val
			if x > y {
				cur.Next = &ListNode{Val: y}
				l2 = l2.Next
				cur = cur.Next
			} else if x < y {
				cur.Next = &ListNode{Val: x}
				l1 = l1.Next
				cur = cur.Next
			} else {
				cur.Next = &ListNode{Val: x}
				cur.Next.Next = &ListNode{Val: y}
				cur = cur.Next.Next
				l1 = l1.Next
				l2 = l2.Next
			}
		} else {
			cur.Next = &ListNode{Val: x}
			l1 = l1.Next
			cur = cur.Next
		}
	}
	//如果第一个循环完了第二个还有节点，则直接将剩下的节点拼到后面
	for l2 != nil {
		cur.Next = &ListNode{Val: l2.Val}
		cur = cur.Next
		l2 = l2.Next
	}

	return pre.Next
}

/**
树的子结构
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:
     3
    / \
   4   5
  / \
 1   2
给定的树 B：
   4
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：
输入：A = [1,2,3], B = [3,1]
输出：false
示例 2：
输入：A = [3,4,5,1,2], B = [4,1]
输出：true

*/
func IsSubStructure(A *TreeNode, B *TreeNode) bool {
	//思路：递归
	return (A != nil && B != nil) && (recur(A, B) || IsSubStructure(A.Left, B) || IsSubStructure(A.Right, B))
}

//看B树是否为A树的子树（根节点相同的情况下）
func recur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}

/**
二叉树的镜像
请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：
     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1

示例 1：
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]

*/
func mirrorTree(root *TreeNode) *TreeNode {
	//思路1 递归
	//前序遍历这个树，将左右两侧的节点交换
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		mirrorTree(root.Left)
		mirrorTree(root.Right)
		return root
	}
	return nil
}

/**
对称的二叉树
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3
*/
func IsSymmetric1(root *TreeNode) bool {
	//思路，bfs，将root的值都装入到队列中（包括nil节点）装一次之后，获取当前队列的所有值（这个树的一排）用双指针，判断当前的元素是否相等
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		cur := queue
		queue = make([]*TreeNode, 0)
		if len(cur)%2 == 1 {
			return false
		} else {
			p1, p2 := 0, len(cur)-1
			for p2 > p1 {
				if (cur[p1] == nil && cur[p2] != nil) || (cur[p1] != nil && cur[p2] == nil) {
					return false
				}
				if cur[p1] == nil && cur[p2] == nil {
					p1++
					p2--
				} else {
					if cur[p1].Val != cur[p2].Val {
						return false
					} else {
						p1++
						p2--
					}
				}
			}
		}
		//如果这一层是对称的，那么将下一层放入队列汇总
		for _, v := range cur {
			if v != nil {
				queue = append(queue, v.Left, v.Right)
			}
		}
	}
	return true
}

/**
顺时针打印矩阵
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]
*/
func SpiralOrder(matrix [][]int) []int {
	lr := len(matrix)
	if lr == 0 {
		return []int{}
	}
	lc := len(matrix[0])
	if lc == 0 {
		return []int{}
	}
	//思路 递归。先打印外面一圈，然后把里面的矩阵递归处理，然后结果拼接
	return spiralOrderHelp(matrix, 0, 0, lr-1, lc-1)
}

func spiralOrderHelp(matrix [][]int, x, y, m, n int) []int {
	res := make([]int, 0)
	//先打印第一排
	for i := x; i <= n; i++ {
		res = append(res, matrix[x][i])
	}
	//再打印最后一列
	for i := x + 1; i <= m; i++ {
		res = append(res, matrix[i][n])
	}
	//判断是否还需要打印下面的
	if m > x {
		//再倒叙打印最后一排
		for i := n - 1; i >= y; i-- {
			res = append(res, matrix[m][i])
		}
	}
	if n > y {
		//再倒叙打印最左边一列
		for i := m - 1; i >= x+1; i-- {
			res = append(res, matrix[i][y])
		}
	}
	//再将内侧的值递归打印
	if m-1 >= x+1 && n-1 >= y+1 {
		res = append(res, spiralOrderHelp(matrix, x+1, y+1, m-1, n-1)...)
	}
	return res
}

/**
包含min函数的栈
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
*/
type MinStack struct {
	stack    []int
	minStack []int
	curMin   int
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
		curMin:   1<<31 - 1,
	}
}

func (this *MinStack) Push(x int) {
	//在push的时候，将元素放入栈中，同时判断当前元素跟curMin的大小，如果比curMin大，则将curMin放入辅助栈，否则将当前元素放入辅助栈，并更新curMin
	this.stack = append(this.stack, x)
	if x < this.curMin {
		this.curMin = x
	}
	this.minStack = append(this.minStack, this.curMin)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

/**
栈的压入、弹出序列
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。
例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，
但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
示例 1：
输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
示例 2：
输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。

*/
func validateStackSequences(pushed []int, popped []int) bool {
	//思路，使用一个辅助栈来做。依次将pushed中的元素放入到辅助栈。同时判断该元素是否为popped中的第一个值，如果是，将其弹出
	stack := make([]int, 0)
	i := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return len(stack) == 0
}
