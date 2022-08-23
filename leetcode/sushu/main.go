package main

import (
	"fmt"
	"math"
)

//素数个数统计
/*
统计n以内素数个数
素数只能被1和本身整除的数字
重点考察： 埃塞法
*/
func main() {
	fmt.Println(int(math.Sqrt(float64(100))))
	var n int
	fmt.Println("请输入需要计算的范围")
	fmt.Scanln(&n)
	count := bf(n)

	fmt.Printf("暴力破解: 0 - %d内共有 %d个素数\n", n, count)

	count = eratosthenes(n)
	fmt.Printf("埃筛法: 0 - %d内共有 %d个素数\n", n, count)
}

/*
判断素数 其实只需要判断到根号x就可以，假如x是12
x可以 可以拆解为 2*6 3*4 4*3 6*2 前后两部分其实只是顺序不同，其他都一样
所以只需要循环到 根号12 就可以，节省了一半的效率，注意：判断条件需要<=,
i <= int(math.Sqrt(float64(x))) 正统写法
i * i <= x 是简便写法
*/

//判断是否是素数
func isPrime(x int) bool {
	//for i := 2; i < x; i++ { //暴力循环全部
	//for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 { //这里到x就是不包含本身了
			return false
		}
	}
	return true
}

//暴力算法
func bf(n int) int {
	var count int
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

/*
埃筛法： 思路
for j := 2 * i; j < n; j += i {
j += i 其实就是为了 递增乘数
假设i为2
j = 2 * 2  j = 2 *2 +2  其实就是等于3*2  4*2  5*2
而当循环执行的是
2*2 3*2 4*2 5*2 6*2
2*3 3*3 4*3 5*3 6*3 7*3
2*4 3*4 4*4 5*4 6*4 7*4
2*5 3*5 4*5 5*5
可以找到规则 i*i之前的的数字不用进行判断，因为前面已经判断过 所以可以优化为
for j := i * i; j < n; j += i {
这样减少循环次数

		if !isHe[i] { //取反合数就是素数
			count++
这里大家可能有个疑问比如 isHe[4]的时候，是否判断成素数，
不可能发生，因为在内循环的时候 已经把4置为合数，所以会跳过循环
*/

//埃筛法  非素数(合数)  2是素数 3是素数 2*3=6 必定是合数
func eratosthenes(n int) int {
	isHe := make([]bool, n) //false 为合数， true为素数
	var count int
	for i := 2; i < n; i++ {
		if !isHe[i] { //取反合数就是素数
			count++
			for j := i * i; j < n; j += i {
				isHe[j] = true
			}
		}
	}
	return count
}
