package strategy

import (
	"fmt"
	"testing"
)

func TestDoubleAverage(t *testing.T) {
	//10个人 抢10000分  也就是10个人抢100块钱
	count, amount := int64(5), int64(1000)
	remain := amount
	sum := int64(0)
	for i := int64(0); i < count; i++ {
		x := DoubleAverage(count-i, remain)
		remain -= x
		sum += x
		fmt.Println(i+1, "=", float64(x)/float64(100), ", ")
	}
	fmt.Printf("总金额是: %v\n", sum)
}
