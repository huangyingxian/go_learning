package array_test

import (
	"fmt"
	"log"
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for i, e := range arr3 { // 数组使用 range 返回其中的index 和对应的value
		t.Log(i, e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[:]
	log.Print(fmt.Sprintf("arr3_sec type:%T", arr3_sec))
	t.Log(arr3_sec) // [1 2 3 4 5]

	arr3[0] = 9
	t.Log(arr3_sec) // [9 2 3 4 5]
	t.Log(arr3)     // [9 2 3 4 5]  这里表明，使用 array[:] 得到的数组是原来数组的引用或者说是指向，操作的是同一份数据

}
