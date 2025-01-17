package slice_test

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(s2[:cap(s2)])
	t.Log(len(s2), cap(s2))

	s2 = append(s2, 5)
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 6)
	t.Log(len(s2), cap(s2)) // 这里表明，slice 当容量不足时，会自己扩充，扩充依据是以前容量的两倍，即 *2

}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s)) // 这里表明，slice 当容量不足时，会自己扩充，扩充依据是以前容量的两倍，即 *2，这里更加显而易见
	}
}

// 该测试说明，与数组一样，切片同样在使用【：】获取子切片时，是共享同一个对象的
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Logf("%T", summer)
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	// if a == b { //切片只能和nil比较
	// 	t.Log("equal")
	// }
	t.Log(a, b)
	t.Log(fmt.Sprintf("%T", b))
}
