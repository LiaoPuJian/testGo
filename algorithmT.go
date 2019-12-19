package main

func main() {

}

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
//做法1，设置一个map，将数组的键值对放入map中
func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		x := target - v
		if _, ok := m[x]; ok {
			return []int{k, m[x]}
		}
		m[v] = k
	}
	return []int{-1, -1}
}

func lengthOfLongestSubstring1(s string) int {
	//指针法，记录每个不同的字符串在map中的位置
	m := make(map[byte]int)
	//这个值代表从第几位开始计算长度
	start := 0
	//最大长度
	max := 0
	sArr := []byte(s)

	for k, v := range sArr {
		//首先，判断当前这个值是否在m中，如果当前的值在m中，则代表，上一个v以前的数据都不要了，只能从上一个v+1的地方再开始计算长度
		if i, ok := m[v]; ok && i >= start {
			start = i + 1
		}
		//然后判断，当前这个键，到开始计算长度的start那的长度是多大，是否大于max。如果大于max，则将max置为新的值
		if k-start+1 > max {
			max = k - start + 1
		}
		//然后将这个键值对放入到m中
		m[v] = k
	}

	return max
}

func myAtoi1(str string) int {
	MaxInt32 := 1<<31 - 1
	MinInt32 := -1 << 31

}
