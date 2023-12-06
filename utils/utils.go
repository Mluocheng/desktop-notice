package utils

import (
	"context"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// App struct
type Utils struct {
	ctx context.Context
}

type UtilsInterface interface {
	RandomGrandLotto() [][]int
}

// Greet returns a greeting for the given name
func RandomGrandLotto(count int) [][]int {
	if count <= 0 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())

	var res [][]int
	for i := 0; i < count; i++ {
		item := random5()
		res = append(res, item)
	}
	fmt.Print(res)

	return res
}

// 获得范围内的随机数
func random(max int) int {
	return rand.Intn(max) + 1
}

// 随机一次选择5个数
func random5() []int {
	var res []int
	var res2 []int
	for i := 0; i < 5; i++ {
		for {
			random := random(35)
			if contains(res, random) {
				continue
			} else {
				res = append(res, random)
				break
			}
		}

	}
	sort.Ints(res)

	for i := 0; i < 2; i++ {
		for {
			random := random(12)
			if contains(res, random) {
				continue
			} else {
				res2 = append(res2, random)
				break
			}
		}
	}
	sort.Ints(res2)

	mergedSlice := append(res, res2...)
	return mergedSlice
}

// 判断是否存在
func contains(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}
