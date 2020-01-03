package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//打印链表的方法
func printList(res *ListNode) {
	a := res
	fmt.Print("res:")
	for {
		if a == nil {
			break
		}
		fmt.Print(a.Val)
		a = a.Next
	}
	fmt.Print("\n")
}

func main() {

	/*list1 := &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}}}
	list2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}

	addTwoNumbers(list1, list2)*/

	/*fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{3, 5, 7}, []int{6, 9, 20}))*/

	/*fmt.Println(longestPalindrome("abb"))
	fmt.Println(longestPalindrome("222020221"))
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("qcwerewwq"))
	fmt.Println(longestPalindrome("babaddtattarrattatddetartrateedredividerb"))
	fmt.Println(longestPalindrome("azwdzwmwcqzgcobeeiphemqbjtxzwkhiqpbrprocbppbxrnsxnwgikiaqutwpftbiinlnpyqstkiqzbggcsdzzjbrkfmhgtnbujzszxsycmvipjtktpebaafycngqasbbhxaeawwmkjcziybxowkaibqnndcjbsoehtamhspnidjylyisiaewmypfyiqtwlmejkpzlieolfdjnxntonnzfgcqlcfpoxcwqctalwrgwhvqvtrpwemxhirpgizjffqgntsmvzldpjfijdncexbwtxnmbnoykxshkqbounzrewkpqjxocvaufnhunsmsazgibxedtopnccriwcfzeomsrrangufkjfzipkmwfbmkarnyyrgdsooosgqlkzvorrrsaveuoxjeajvbdpgxlcrtqomliphnlehgrzgwujogxteyulphhuhwyoyvcxqatfkboahfqhjgujcaapoyqtsdqfwnijlkknuralezqmcryvkankszmzpgqutojoyzsnyfwsyeqqzrlhzbc"))
	*/

	//fmt.Println(Zconvert("LEETCODEISHIRINGA", 3))

	//fmt.Println(reverse(120))

	//fmt.Println(isPalindrome(121))

	//fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7, 10}))

	//fmt.Println(intToRoman(2494))

	//fmt.Println(romanToInt("DCXXI"))

	/*fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "doracecar", "docar"}))*/

	//fmt.Println(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))

	//fmt.Println(threeSum([]int{-5, 14, 1, -2, 11, 11, -10, 3, -6, 0, 3, -4, -9, -13, -8, -7, 9, 8, -7, 11, 12, -7, 4, -7, -1, -5, 13, 1, -2, 8, -13, 0, -1, 3, 13, -13, -1, 10, 5, 1, -13, -15, 12, -7, -13, -11, -7, 3, 13, 1, 0, 2, 1, 11, 10, 8, -8, 1, -14, -3, -6, -12, 12, 0, 6, 2, 2, -9, -3, 14, -1, -9, 14, -4, -1, 8, -8, 7, -4, 12, -14, 3, -9, 2, 0, -13, -13, -1, 3, -12, 11, 4, -9, 8, 11, 5, -5, -10, 3, -1, -11, -13, 5, -12, -10, 11, 11, -3, -5, 14, -13, -4, -5, -7, 6, 2, -13, 0, 8, -3, 4, 4, -14, 2}))

	//fmt.Println(threeSum([]int{3, 0, -2, -1, 1, 2}))

	//fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))

	//fmt.Println(letterCombinations("23"))

	//fmt.Println(fourSum([]int{-1, 0, -5, -2, -2, -4, 0, 1, -2}, -9))

	//list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}

	//l := removeNthFromEnd(list, 2)

	//fmt.Println(isValid("([)]"))

	/*list1 := &ListNode{Val: -10, Next: &ListNode{Val: -6, Next: &ListNode{Val: -6, Next: &ListNode{Val: -6, Next: &ListNode{Val: -3, Next: &ListNode{Val: 5, Next: nil}}}}}}
	//list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	fmt.Println(mergeTwoLists(list1, nil))*/

	/*list1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	list3 := &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: nil}}

	l := mergeKLists([]*ListNode{list1, list2, list3})
	printList(l)*/

	printList(swapPairs(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}))
}

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
func twoSum(nums []int, target int) []int {
	//先定义一个map
	m := make(map[int]int)
	for k, v := range nums {
		x := target - v
		if k1, ok := m[x]; ok {
			return []int{k, k1}
		}
		m[v] = k
	}
	return []int{-1, -1}
}

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//1、分别计算出两个链表的值，相加之后，倒过来，再生成链表   (这样不行，测试用例中有很多很长的链表，相加会溢出)

	//2、循环两个链表的每个节点，如果有一个节点为空，则将那个节点的值置为0，将两个节点相加，然后得到的合如果大于10则减去10，小于10不做操作，
	//将这个合放入到输出链表里。如果大于10，则记录进位。
	pre := &ListNode{Val: 0}
	cur := pre

	var carry int

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		//计算当前两个链表的值的和是否进一
		carry = sum / 10
		//计算出生成的链表的值
		sum = sum % 10
		//生成一个节点，放入cur中
		cur.Next = &ListNode{Val: sum}
		//开始处理下个节点
		cur = cur.Next
	}
	//如果循环结束后，此时的carry还为1，证明此时最后一位相加也超出了10，则再往后续一个子节点
	if carry == 1 {
		cur.Next = &ListNode{Val: carry}
	}

	return pre.Next
}

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

/*
	//解法1.暴力破解法，计算出字符串的长度，循环每个字符，往后循环，取出最长的不重复子串，然后比较最终的大小
	max := 1
	l := len(s)
	for i := 0; i < l; i++ {
		//新建一个临时的map
		m := make(map[uint8]int)
		m[s[i]] = 1
		temp := 1
		for j := i + 1; j < l; j++ {
			if _, ok := m[s[j]]; ok {
				break
			} else {
				//将这个值放入m中
				m[s[j]] = 1
				temp++
			}
		}
		if temp > max {
			max = temp
		}
	}
	return max
*/
func getMaxLengthDiffString(str string) (int, string) {
	//定义一个map，用于存储这个字符串的字符和位置
	m := make(map[byte]int)
	start := 0
	maxLength := 0
	returnString := ""
	s := []byte(str)

	for i, ch := range []byte(str) {
		//判断，如果当前这个字节已经存储在m中，且位置大于start，将start置为v+1
		if v, ok := m[ch]; ok && v >= start {
			start = v + 1
		}
		//判断最大长度
		if i-start+1 > maxLength {
			maxLength = i - start + 1
			returnString = string(s[start : i+1])
		}
		m[ch] = i
	}
	return maxLength, returnString
}

//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//你可以假设 nums1 和 nums2 不会同时为空。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//思路1:先将两个数组合并，并从小到大排序。然后判断个数奇偶，然后计算中位数
	/*nums := append(nums1, nums2...)
	//冒泡排序
	nums = func(arr []int) []int {
		lens := len(arr)
		for i := 0; i < lens; i++ {
			for j := i + 1; j < lens; j++ {
				if arr[i] > arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
		return arr
	}(nums)
	x := 0.0
	if len(nums)%2 == 1 {
		//奇数，取中间的值
		x = float64(nums[len(nums)/2])
	} else {
		//偶数，取中间的两个值
		x = float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2
	}
	return x*/

	//思路2：在两个数组中新增两个游标，来分别比较两个游标对应的值的大小，将小的一方放入一个新的数组中，本质上还是排序
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2
	i, j := 0, 0
	var temp []int

	for {
		if i < l1 && j < l2 {
			if nums1[i] <= nums2[j] {
				temp = append(temp, nums1[i])
				i++
			} else {
				temp = append(temp, nums2[j])
				j++
			}
		} else if i >= l1 && j < l2 {
			temp = append(temp, nums2[j])
			j++
		} else if i < l1 && j >= l2 {
			temp = append(temp, nums1[i])
			i++
		}

		if len(temp) == l/2+1 {
			if l%2 == 1 {
				return float64(temp[l/2])
			} else {
				return float64(temp[l/2-1]+temp[l/2]) / 2.0
			}
		}
	}

}

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
func longestPalindrome(s string) string {
	//思路1，暴力破解
	l := len(s)
	if l < 2 {
		return s
	}
	byteArr := []byte(s)
	maxS := string(byteArr[0])
	maxL := 0

	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			//正向字符串
			str := string(byteArr[i:j])
			//反向字符串
			reStr := reverseString(str)
			if str == reStr && len(byteArr[i:j]) > maxL {
				maxL = j - i
				maxS = str
			}
		}
	}
	return maxS

	//思路2，设定一个游标，从第一个值开始，然后往下走，分别比较游标左右的值，如果是回文，则存储下来，再比较左-1和右+1的值
	//将这个字符串用0分割
	/*if len(s) == 0 || len(s) == 1 {
		return s
	}
	s = strings.Join(strings.Split(s, ""), ".")
	s = "." + s + "."
	l := len(s)
	byteArr := []byte(s)
	maxS := string(byteArr[0])
	maxLength := 1

	for i := 0; i < l; i++ {
		if i == 0 && byteArr[0] == byteArr[1] {
			maxS = string(byteArr[0:2])
			maxLength = 2
		}
		if i == l-1 && byteArr[l-1] == byteArr[l-2] && maxLength == 1 {
			maxS = string(byteArr[l-2 : l])
			maxLength = 2
		}
		loop := i
		//判断，当前游标到数组最左边长还是到最右边长
		if l-i-1 < i+1 {
			loop = l - i - 1
		}
		for j := 1; j <= loop; j++ {
			left := byteArr[i-j]
			right := byteArr[i+j]
			if left == right {
				//是回文子串,判断此字符串的长度和maxLength的大小比较
				if maxLength < len(byteArr[i-j:i+j+1]) {
					maxLength = len(byteArr[i-j : i+j+1])
					maxS = string(byteArr[i-j : i+j+1])
				}
			} else if (byteArr[i] == left || byteArr[i] == right) && maxLength == 1 {
				var b []byte
				maxLength = 2
				maxS = string(append(b, byteArr[i], byteArr[i]))
			} else {
				break
			}
		}
	}
	return strings.Replace(maxS, ".", "", 1001)*/
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

/**
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

*/
func Zconvert(s string, numRows int) string {
	//以等差法来计算，以字符串为一个数组，键之间的差值满足 d = 2 * numRows - 2
	if numRows <= 1 || len(s) == 0 {
		return s
	}
	rst := ""
	size := 2*numRows - 2
	//循环行数
	for i := 0; i < numRows; i++ {
		//循环这一行的数据，以size来累计加
		for j := i; j < len(s); j += size {
			rst += string(s[j])
			//加上第一个字符之后，判断这一行如果不是第一行或者最后一行，则中间应该有一个值，这个值的位置是固定的，为j + size - 2*i
			tmp := j + size - 2*i
			if i != 0 && i != numRows-1 && tmp < len(s) {
				rst += string(s[tmp])
			}
		}
	}
	return rst
}

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
func reverse(x int) int {
	MaxInt32 := 1<<31 - 1
	MinInt32 := -1 << 31

	var num, newNum int

	//循环这个值
	for x != 0 {
		//得到当前x的个位的值
		a := x % 10
		//将上一次循环的num*10，和当前x的个位的值相加
		newNum = num*10 + a
		//将这个值付给num
		num = newNum
		//将x减少一个个位，进行下一次循环
		x = x / 10
		if num > MaxInt32 || num < MinInt32 {
			return 0
		}
	}
	return num
}

/**
请你来实现一个 atoi 函数，使其能将字符串转换成整数。
首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
在任何情况下，若函数不能进行有效的转换时，请返回 0。
*/
func myAtoi(str string) int {
	//step1：去无效字符
	str = strings.TrimSpace(str)
	if str == "" || (len(str) == 1 && (str < "0" || str > "9")) {
		return 0
	}
	//step2：规范首字符
	flag := ""
	if string(str[0]) == "-" {
		flag = "-"
		str = str[1:len(str)]
	} else if string(str[0]) == "+" {
		str = str[1:len(str)]
	}
	//step3：遍历检测数字0-9
	resStr := "0"
	for i := 0; i < len(str); i++ {
		if string(str[i]) < "0" || string(str[i]) > "9" {
			break
		}
		resStr += string(str[i])
	}
	resStr = flag + resStr
	//step4：转换
	res, err := strconv.ParseInt(resStr, 10, 32)

	const MaxUint = ^uint32(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	//step5：转换异常处理
	if err != nil {
		if flag == "" {
			return MaxInt
		}
		return MinInt
	}
	return int(res)
}

/**
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/
func isPalindrome(x int) bool {
	//思路1、整数转string，然后通过字符串翻转来做
	/*str := strconv.Itoa(x)
	strR := reverseString(str)
	if str == strR {
		return true
	} else {
		return false
	}*/

	//思路2 不将整数转为字符串
	//1、先判断是否为负数，如果是，则直接返回false
	if x < 0 {
		return false
	}

	//2、翻转这个整数，判断前后是否相等
	x1 := x
	var num, numNew int
	for x1 != 0 {
		//将x1除以10，得到余数
		a := x1 % 10
		numNew = num*10 + a
		num = numNew
		x1 = x1 / 10
	}

	if x == num {
		return true
	} else {
		return false
	}

}

/**
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器，且 n 的值至少为 2。
*/
func maxArea(height []int) int {
	//思路1，暴力计算，循环每一个垂直线之间的差，计算出水的容积
	/*var max int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			h := 0
			if height[j] >= height[i] {
				h = height[i]
			} else {
				h = height[j]
			}
			if max < h*(j-i) {
				max = h * (j - i)
			}
		}
	}
	return max*/

	//思路2，指针移动法。假设有两个指针，一个指向数组的第一个值，第二个指向数组的最后一个值，求出当前的面积，然后存入一个值中
	//然后判断第一个指针对应的值和第二个指针对应的值哪个大，将小的那个指针往中间移动一格，然后重复
	l := 0
	r := len(height) - 1
	max := 0
	for r-l >= 0 {
		//计算面积
		newMax := 0
		if height[l] >= height[r] {
			newMax = height[r] * (r - l)
			r--
		} else {
			newMax = height[l] * (r - l)
			l++
		}
		if max < newMax {
			max = newMax
		}
	}
	return max

}

/**
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。
*/
func intToRoman(num int) string {
	var str string
	var qian, bai, shi, ge int
	qian = num / 1000
	bai = (num - (qian * 1000)) / 100
	shi = (num - (qian*1000 + bai*100)) / 10
	ge = num % 10

	if qian > 0 {
		for i := 0; i < qian; i++ {
			str += "M"
		}
	}

	if bai > 0 {
		if bai <= 3 {
			for i := 0; i < bai; i++ {
				str += "C"
			}
		} else if bai == 4 {
			str += "CD"
		} else if bai == 9 {
			str += "CM"
		} else {
			str += "D"
			if bai >= 6 {
				for i := 1; i <= bai-5; i++ {
					str += "C"
				}
			}
		}
	}

	if shi > 0 {
		if shi <= 3 {
			for i := 0; i < shi; i++ {
				str += "X"
			}
		} else if shi == 4 {
			str += "XL"
		} else if shi == 9 {
			str += "XC"
		} else {
			str += "L"
			if shi >= 6 {
				for i := 1; i <= shi-5; i++ {
					str += "X"
				}
			}
		}
	}

	if ge > 0 {
		if ge <= 3 {
			for i := 0; i < ge; i++ {
				str += "I"
			}
		} else if ge == 4 {
			str += "IV"
		} else if ge == 9 {
			str += "IX"
		} else {
			str += "V"
			if ge >= 6 {
				for i := 1; i <= ge-5; i++ {
					str += "I"
				}
			}
		}
	}
	return str
}

/**
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
*/
func romanToInt(s string) int {
	num := 0
	var prev byte
	for _, v := range []byte(s) {
		if v == 'M' {
			//判断上一个是不是C
			if prev == 'C' {
				num += 800 //这里为什么是800而不是900，因为它前面如果是一个C的话，已经加过一次100了亦如是
			} else {
				num += 1000
			}
		}
		if v == 'D' {
			//判断上一个是不是C
			if prev == 'C' {
				num += 300
			} else {
				num += 500
			}
		}

		if v == 'C' {
			//判断上一个是不是X
			if prev == 'X' {
				num += 80
			} else {
				num += 100
			}
		}

		if v == 'L' {
			//判断上一个是不是X
			if prev == 'X' {
				num += 30
			} else {
				num += 50
			}
		}

		if v == 'X' {
			//判断上一个是不是I
			if prev == 'I' {
				num += 8
			} else {
				num += 10
			}
		}

		if v == 'V' {
			//判断上一个是不是I
			if prev == 'I' {
				num += 3
			} else {
				num += 5
			}
		}

		if v == 'I' {
			num += 1
		}

		prev = v
	}

	return num
}

/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。
示例 1:

输入: ["flower","flow","flight"]
输出: "fl"

示例 2:
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
*/
func longestCommonPrefix(strs []string) string {
	var str []byte

	if len(strs) == 0 {
		return ""
	}

	//获取数组第一个元素的长度
	l := len(strs[0])

	//按照这个第一个元素的长度来循环
	for i := 0; i < l; i++ {
		var x byte
		y := -1
		//判断第一个元素是否都是一样的
		for k, v := range strs {
			//如果是循环的数组的第一个元素，则直接将第一个元素的第i个值赋予x
			if k == 0 {
				x = v[i]
			} else {
				//判断当前这个值是否等于x，如果不等于，则证明第i个元素不是公共元素。 如果当前循环的字符串长度比较短，没有第i个元素，也同理
				if len(v)-1 < i || v[i] != x {
					break
				}
			}
			//如果循环的是最后一个元素，且也通过了，则证明这个i是公共元素，将其放入str中
			if k == len(strs)-1 {
				y = k
			}
		}

		if y == len(strs)-1 {
			str = append(str, x)
		} else {
			//证明不是第i个元素不是每个字符串都通过了，则跳出循环
			break
		}
	}

	return string(str)
}

/**
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/
func threeSum(nums []int) [][]int {

	//1、思路1 暴力破解

	//2、思路2，指针法
	//先将数组从小到大排序
	sort.Sort(IntSlice(nums))

	var res [][]int

	l := len(nums)

	m := make(map[int][]int)

	//指针从第二位开始，到倒数第二位结束
	for i := 1; i <= l-1; i++ {
		first := 0
		last := l - 1
		for last > i && i > first {

			//fmt.Println("转换前：", first, i, last, nums[first], nums[i], nums[last])
			if nums[first]+nums[last]+nums[i] == 0 {

				if v, ok := m[nums[first]]; ok {
					if v[0] == nums[i] || v[0] == nums[last] {
						//fmt.Println("已经存在了，跳过", nums[first], nums[i], nums[last])
						first = first + 1
						last = last - 1
						continue
					}
				}
				//fmt.Println("数据ok，存入", nums[first], nums[i], nums[last])
				res = append(res, []int{nums[first], nums[last], nums[i]})
				m[nums[first]] = []int{nums[i], nums[last]}
				first = first + 1
				last = last - 1

			} else if nums[first]+nums[last]+nums[i] < 0 {
				first = first + 1
			} else {
				last = last - 1
			}

			//fmt.Println("转换后：", first, i, last, nums[first], nums[i], nums[last])

		}
	}

	return res
}

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

/**
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
*/
func threeSumClosest(nums []int, target int) int {

	//1、暴力法    此法会超时
	/*var sum int
	var minCha int
	minCha = 10000

	l := len(nums)
	//循环数组
	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				cha := nums[i] + nums[j] + nums[k] - target
				fmt.Println(nums[i], nums[j], nums[k], cha, minCha)
				if cha < 0 {
					cha = -cha
				}
				if minCha > cha {
					minCha = cha
					sum = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}

	return sum*/

	//2、左右游标法
	var sum int
	minCha := 10000

	l := len(nums)
	//先将数组从小到大排序
	sort.Sort(IntSlice(nums))

	for i := 1; i <= l-1; i++ {
		first := 0
		last := l - 1

		for last-i > 0 && i-first > 0 {
			s := nums[i] + nums[first] + nums[last]
			cha := s - target
			if s-target == 0 {
				return s
			} else if s-target > 0 {
				//太大了，右侧左移一位
				last -= 1
			} else {
				//太小了，左侧右移一位
				first += 1
			}
			//取绝对值
			if cha < 0 {
				cha = -cha
			}
			if cha < minCha {
				minCha = cha
				sum = s
			}
		}
	}

	return sum
}

/**
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/
func letterCombinations(digits string) []string {
	//定义好一个map，用来存放数字对应的字符
	m := make(map[byte][]string)
	m['2'] = []string{"a", "b", "c"}
	m['3'] = []string{"d", "e", "f"}
	m['4'] = []string{"g", "h", "i"}
	m['5'] = []string{"j", "k", "l"}
	m['6'] = []string{"m", "n", "o"}
	m['7'] = []string{"p", "q", "r", "s"}
	m['8'] = []string{"t", "u", "v"}
	m['9'] = []string{"w", "x", "y", "z"}

	result := make([]string, 0)

	//这里为了避免digits为""的情况
	if len(digits) == 0 {
		return result
	}

	f("", digits, &result, m)

	return result
}

/**
回溯算法。输入一个将被拼装的字符串和要回溯的字符串，一个结果集
如果next_digits为空，则证明后续没有需要回溯的字符串了，直接将拼接好的字符串放入结果集中
否则，则获取要回溯的字符串的第一个字符，去m中查到其对应的字符串，然后将这些字符串拼到combination中，并将next_digits从下一位开始，继续调用回溯方法
*/
func f(combination, next_digits string, result *[]string, m map[byte][]string) {
	if len(next_digits) == 0 {
		*result = append(*result, combination)
	} else {
		for _, v := range m[next_digits[0]] {
			f(combination+v, next_digits[1:], result, m)
		}
	}
}

/**
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
注意：
答案中不可以包含重复的四元组。
示例：
给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

*/
func fourSum(nums []int, target int) [][]int {

	var res [][]int

	l := len(nums)

	if l < 4 {
		return res
	}

	//对数组进行排序
	sort.Sort(IntSlice(nums))

	//定义四个指针，k, i, j, h 。k从0开始遍历，i从k+1开始遍历，留下j和h。j指向i+1，h指向数组最后一位
	for k := 0; k < l-3; k++ {
		//当k的值与上一个值相等时，跳过当前循环
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		//获取当前k的最小值，如果最小值直接大于了target，则可以直接跳过循环了
		var min1 = nums[k] + nums[k+1] + nums[k+2] + nums[k+3]
		if min1 > target {
			continue
		}
		//获取当前k的最大值，如果最大值直接小于了target，则也可以直接跳过循环
		var max1 = nums[k] + nums[l-1] + nums[l-2] + nums[l-3]
		if max1 < target {
			continue
		}

		//第二层循环
		for i := k + 1; i < l-2; i++ {
			//当i的值与上一个值相等时，跳过当前循环
			if i > k+1 && nums[i] == nums[i-1] {
				continue
			}
			//定义指针
			var j, h = i + 1, l - 1
			//fmt.Println("最小值：", k, i, j, h, nums[k], nums[i], nums[j], nums[j+1])
			//获取当前i的最小值，如果最小值直接大于了target，则可以直接跳过循环了
			var min2 = nums[k] + nums[i] + nums[j] + nums[j+1]
			if min2 > target {
				continue
			}
			//fmt.Println("最大值：", k, i, j, h, nums[k], nums[i], nums[h-1], nums[h])
			//获取当前k的最大值，如果最大值直接小于了target，则也可以直接跳过循环
			var max2 = nums[k] + nums[i] + nums[h-1] + nums[h]
			if max2 < target {
				continue
			}

			//这里开始操作指针j和h
			for j < h {
				//fmt.Println(k, i, j, h, nums[k], nums[i], nums[j], nums[h])
				sum := nums[k] + nums[i] + nums[j] + nums[h]
				if sum == target {
					x := []int{nums[k], nums[i], nums[j], nums[h]}
					jumpFlag := false
					//如果sum刚好等于目标值，则去重后将其放入到res中
					for _, v := range res {
						if reflect.DeepEqual(v, x) {
							j++
							h--
							jumpFlag = true
							break
						}
					}
					if jumpFlag {
						continue
					}
					//将数组放入res中，
					res = append(res, x)
					//移动指针
					j++
					h--
				} else if sum < target {
					//如果sum小于目标值，则将j往右侧移动一位
					j++
				} else {
					//如果sum大于目标值，则将h往左侧移动一位
					h--
				}
			}
		}
	}

	return res
}

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//递归思想，循环获取需要删除的元素，并将倒数n-1个元素的Next指定为倒数n+1个元素
	//解题思路1、先计算出链表的长度，然后获取倒数第n个值在链表中正向的顺序
	/*l1 := head
	var length = 1
	for l1.Next != nil {
		l1 = l1.Next
		length++
	}

	if length == 1 && n == 1 {
		return nil
	}
	//如果长度和倒数位相等，则证明需要去除第一个值，直接返回head的子元素即可
	if length == n {
		return head.Next
	}

	l := head

	for i := 1; i <= length-n; i++ {
		if i == length-n {
			//将l的Next指向n+1
			l.Next = l.Next.Next
			break
		}
		l = l.Next
	}
	return head*/

	//阶梯思路2、双指针。设定两个指针，一个指向正向顺序为n+1的地方，一个指向正向顺序为0的地方。然后将两个指针一步一步的正向往后移动，一旦第一个指针移动到末尾指向空时
	//第二个指针就指向了要被删除的上一个元素。此时令这个元素的子元素等于其子子元素即可
	dummy := &ListNode{Val: 0, Next: head}
	first := dummy
	second := dummy
	//将第一个指针移动到n+1的位置
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}
	//开始一个单位一个单位的共同正向移动，一旦first为空，则此时的second即为要删除的上一个元素
	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next
}

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

*/
func isValid(s string) bool {
	//可以这么理解，设定一个栈。每读取到一个元素，则判断其情况。如果是左括号，则将其放入栈中，如果是右括号，则检查栈顶的元素是否是跟其匹配的左括号，
	//如果是，则将左括号弹出（pop）。如果没有，则证明表达式无效。如果最后循环完成，栈中仍有数据，则也证明表达式无效
	b := make([]byte, 0)

	for _, v := range []byte(s) {
		switch v {
		case '(', '[', '{':
			b = append(b, v)
		case ')':
			if len(b) >= 1 && b[len(b)-1] == '(' {
				b = b[:len(b)-1]
			} else {
				return false
			}
		case ']':
			if len(b) >= 1 && b[len(b)-1] == '[' {
				b = b[:len(b)-1]
			} else {
				return false
			}
		case '}':
			if len(b) >= 1 && b[len(b)-1] == '{' {
				b = b[:len(b)-1]
			} else {
				return false
			}
		}
		//fmt.Println(string(b))
	}

	if len(b) > 0 {
		return false
	}
	return true
}

/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	var x, y int
	//先循环第一个链表
	for l1 != nil {
		x = l1.Val
		//如果第二个链表不为空，则比较第一个链表和第二个链表的大小，将小的放入新的链表中，并将对应的链表往后移动一位
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
				//如果两个值相等，则往新链表中塞两个节点，并且同时移动l1和l2
				cur.Next = &ListNode{Val: x}
				cur.Next.Next = &ListNode{Val: y}
				cur = cur.Next.Next
				l1 = l1.Next
				l2 = l2.Next
			}
		} else {
			//如果l2为空的话，则直接将l1的值塞入新的链表中，并移动l1
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

func mergeTwoLists1(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 创建最终返回头的前置节点
	pre := &ListNode{}
	node := pre
	for l1 != nil && l2 != nil {
		//判断，如果当前l1的值大于l2的值，则将l2置为node.Next，反之亦然
		if l1.Val > l2.Val {
			node.Next = l2
			l2 = l2.Next
		} else {
			node.Next = l1
			l1 = l1.Next
		}
		//将node往后移动一位
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}

	return pre.Next
}

/**
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

*/
func generateParenthesis(n int) []string {
	var res []string
	//使用回溯方法
	generateF(&res, "", 0, 0, n)
	return res
}

func generateF(res *[]string, cur string, left, right, max int) {
	//判断，如果当前字符串的长度已经达到了n的两倍，则证明其已经达到最大长度，将其放入res中
	if len(cur) == max*2 {
		*res = append(*res, cur)
		return
	}
	//如果左括号的数量小于n，则可以往其中放入左括号
	if left < max {
		generateF(res, cur+"(", left+1, right, max)
	}
	//如果右括号的数量小于左括号，则可以往其中放入右括号
	if right < left {
		generateF(res, cur+")", left, right+1, max)
	}
}

/**
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/
func mergeKLists(lists []*ListNode) *ListNode {
	//获取所有链表的长度，得到一个总长
	/*res := &ListNode{}
	cur := res
	var allL int
	for _, v := range lists {
		v1 := v
		l := 0
		for {
			if v1 == nil {
				break
			}
			l++
			v1 = v1.Next
		}
		allL += l
	}

	if allL == 0 {
		return nil
	}

	//从第一位循环到这个总长
	for i := 1; i <= allL; i++ {
		min := 100000
		key := 0
		for k, v := range lists {
			//比较每个v的值，取其中最小的一个，放入res中，并将其往后移动一位
			if v != nil && v.Val < min {
				min = v.Val
				key = k
				//fmt.Println("当前循环值的键值和最小值：", k, v.Val, min)
			}
		}
		//fmt.Println("获取当前最小值：", min, "最小值的所属值在lists中的键：", key)
		cur.Val = min
		if i != allL {
			cur.Next = &ListNode{}
			cur = cur.Next
		}

		lists[key] = lists[key].Next
	}

	return res*/

	//解法2、递归，将K个链表转换为两个链表的合并问题

	l := len(lists)
	if l == 0 {
		return nil
	}

	if l == 1 {
		return lists[0]
	}
	//这里假设是一个三个元素的数组，则此时lists[:l/2]为第0个元素，lists[l/2:]为第一个和第二个元素
	//此时第一个和第二个元素继续递归，会先合并成一个链表，再跟第0个合成一个链表
	return mergeTwoLists(mergeKLists(lists[:l/2]), mergeKLists(lists[l/2:]))

}

/**
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。


示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.

*/
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	//获取链表的长度
	cur := head
	l := 0
	for cur != nil {
		l++
		cur = cur.Next
	}

	if l == 1 {
		return head
	}

	cur = head
	parent := &ListNode{Val: 0, Next: head}
	p1 := parent
	for i := 1; i <= l; i += 2 {
		//交换当前节点和其子节点
		//将当前节点的父级的子节点指向当前节点的子节点
		printList(p1)
		printList(cur)
		//避免长度为奇数时的报错
		if cur.Next == nil {
			break
		}

		p1.Next = cur.Next
		printList(p1)
		//将当前节点的子节点切换为其孙子节点
		cur.Next = cur.Next.Next
		printList(cur)
		//将当前父节点的孙子节点换为当前节点
		p1.Next.Next = cur
		printList(p1)

		//往后移动两位父节点
		p1 = p1.Next.Next
		//往后移动一位当前节点
		cur = cur.Next
	}

	return parent.Next
}
