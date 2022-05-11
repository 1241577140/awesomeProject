package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	rwd := []*RewardDetail{{Type: 1, ID: 1, Count: 1}}
	rwdV2 := []RewardDetail{{Type: 1, ID: 1, Count: 1}}
	AddRewards(rwd)
	AddRewardsV2(rwdV2)
	fmt.Println(rwd[0].Count, rwdV2[0].Count)

	slice1 := make([]int, 0, 10)
	slice1 = append(slice1, 1)
	sliceHeader1 := (*reflect.SliceHeader)(unsafe.Pointer(&slice1))
	printHeader(sliceHeader1)
	add(slice1, 2)
	sliceHeader2 := (*reflect.SliceHeader)(unsafe.Pointer(&slice1))
	fmt.Printf("%v\n", sliceHeader2)
	printHeader(sliceHeader2)
	slice2 := slice1[0:10]
	fmt.Printf("slice2:%v len%d cap%d", slice2, len(slice2), cap(slice2))

}

func AddRewards(rewards []*RewardDetail) {
	for _, reward := range rewards {
		reward.Count *= 2
	}
}

func AddRewardsV2(rewards []RewardDetail) {
	for _, reward := range rewards {
		reward.Count *= 2
	}
}

type RewardDetail struct {
	Type  int
	ID    int
	Count int
}

func add(slice []int, n int) {
	slice = append(slice, n)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	printHeader(sliceHeader)
}

func printHeader(header *reflect.SliceHeader) {
	bytes, err := json.Marshal(header)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}
