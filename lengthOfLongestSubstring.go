package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[byte]int)// map存储最近一次该字符出现的位置
	start := 0	// 无重复字符起始位置
	maxLen := 0  // 记录最大长度
	for i, v := range []byte(s) {
		if lastIndex, ok := lastOccurred[v]; ok && lastIndex >= start { // sa a asab 
			start = lastIndex + 1
			fmt.Printf("start:%d,index:%d\n",start,i)
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1  
		}
		lastOccurred[v]=i 
	}
	return maxLen

}

func main() {
	fmt.Println(lengthOfLongestSubstring ("saaasab"))
}
