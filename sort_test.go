package quick

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var (
	strArr = []interface{}{
		"c", "a", "z", "i", "b",
	}
	intArr = []interface{}{8, 5, 7, 6, 1, 3}
	
	floatArr    = []interface{}{8.9, 5.12, 0.314, 3.1415, 1.7058}
	sortedStr   = []string{"z", "i", "c", "b", "a"}
	sortedInt   = []int{8, 7, 6, 5, 3, 1}
	sortedFloat = []float64{0.314, 1.7058, 3.1415, 5.12, 8.9}
)

func TestQuickSort(t *testing.T) {
	
	newStrArr := Sort(strArr, false)
	newIntArr := Sort(intArr, false)
	newFloatArr := Sort(floatArr, true)
	
	for i, v1 := range newStrArr {
		if !reflect.DeepEqual(v1, sortedStr[i]) {
			t.FailNow()
		}
	}
	for i, v1 := range newIntArr {
		if !reflect.DeepEqual(v1, sortedInt[i]) {
			t.FailNow()
		}
	}
	for i, v1 := range newFloatArr {
		if !reflect.DeepEqual(v1, sortedFloat[i]) {
			t.FailNow()
		}
	}
	
}

func BenchmarkQuickSort(b *testing.B) {
	// Fill a random array
	rand.Seed(time.Now().Unix())
	randArr := make([]interface{}, 1000)
	for i := 0; i < 1000; i++ {
		randArr[i] = rand.Float64() + 1
	}
	
	for i := 0; i < b.N; i++ {
		Sort(randArr, true)
	
	}
	
	b.ReportAllocs()
}
