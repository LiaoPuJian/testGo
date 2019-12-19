package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
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

	fmt.Println(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))
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
				num += 800
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
	if len(strs) == 0 {
		return ""
	}

	common := ""
	str := strs[0]

	for i := 0; i < len(str); i++ {
		s := str[0 : i+1]
		for _, vv := range strs[1:] {
			if strings.Index(vv, s) != 0 {
				return common
			}
		}
		common = s
	}

	return common
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
	sli := [][]int{}
	l := len(nums)
	if l < 3 {
		return sli
	}
	//排序数组
	sort.Ints(nums)
	//以第二个元素作为中间轴，定义两个指针，一个在数组头，一个在数组尾巴
	for i := 1; i < l-1; i++ {
		first := 0
		last := l - 1
	}

	return sli
}
