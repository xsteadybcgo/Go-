package main
func main() {
  s := []int{0,1,2,3,4,5}
  s = append(s[0:1], s[2:]...) //删除了slice第1个的值
}
