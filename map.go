package main

import (
	"fmt"
	//"sort"
)

func main() {
	var m map[int]string
	m = make(map[int]string)
	m[1] = "ok"
	m[2] = "yes"
	delete(m, 1)
	fmt.Println(m)

	m2 := make(map[int]map[int]int)
	a, ok := m2[2][1]
	if !ok {
		m2[2] = make(map[int]int)
	}
	m2[2][1] = 10
	a, ok = m2[2][1]
	fmt.Println(a, ok)

	sm := make([]map[int]string, 10)
	for k := range sm {
		sm[k] = make(map[int]string)
		sm[k][1] = "x"
		fmt.Println(sm[k])
	}
	fmt.Println(sm)

	/*	mmm := map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"}
		smm := make([]int, len(mmm))
		i := 0
		for k,_ := range mmm{
			smm[i] = k
			i++
		}
		sort.Ints(smm)
		fmt.Println(smm)*/

	/*	m1 := map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"}
		m2 := make(map[string]int)

		for k,v := range m1 {
			m2[v] = k
		}
		fmt.Println(m1)
		fmt.Println(m2)*/

}
