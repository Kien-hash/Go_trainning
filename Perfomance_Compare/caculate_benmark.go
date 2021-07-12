package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N1 = 500000.0
const N2 = 1000000.0
const N3 = 1000000000.0
const N4 = 10000000000.0
const N5 = 100000000000.0
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	start := time.Now()

	{
		//Do smt here to consume time
	}

	elapsed := time.Since(start)
	fmt.Println("Time comsuming: ", elapsed)
}

/********************************************************************
   Ham tinh tong tu 1 den N
*********************************************************************/
func sumBenmark(N float64) float64 {
	sum := 0.0
	for count := 0.0; count < N; count++ {
		sum += count + 1
	}
	return sum
}

/********************************************************************
   Ham tinh tong cua chuoi phan so
*********************************************************************/
func calBenmark(N float64) float64 {
	sum := 0.0
	sign := -1.0
	var num float64
	for count := 0.0; count < N; count++ {
		num = count + 1
		sum += sign * num / ((num + 1) * (num + 2))
		sign = -sign
		//fmt.Println(sum);
	}
	return sum
}

/********************************************************************
   Ham tinh so thu N trong day Fibonaci
*********************************************************************/
func fibonacci(n int) int64 {
	if n < 3 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

/********************************************************************
   Ham sinh ra chuoi co N ki tu ngau nhien
*********************************************************************/
func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

/********************************************************************
   Ham so sanh 2 chuoi co cung do dai la N
*********************************************************************/
func strCompare(str1 string, str2 string) int {
	count := 0
	for i := range str1 {
		if str1[i] == str2[i] {
			count++
		}
	}
	return count
}
