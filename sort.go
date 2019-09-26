package quick

import (
	"reflect"
	"strings"
)

// Default sorting order
var asc = true

// package entry point for external usage
func Sort(arr []interface{}, ascending bool) []interface{} {
	asc = ascending
	cloned := make([]interface{}, len(arr))
	
	for i, v := range arr {
		cloned[i] = v
	}
	
	sort(cloned, 0, len(arr)-1)
	
	return cloned
}

// return true if i<j
func less(i, j interface{}, asc bool) bool {
	switch i.(type) {
	case string:
		
		if asc {
			return strings.Compare(reflect.ValueOf(i).String(), reflect.ValueOf(j).String()) == -1
		} else {
			return strings.Compare(reflect.ValueOf(i).String(), reflect.ValueOf(j).String()) == 1
		}
	case float64:
		if asc {
			return reflect.ValueOf(i).Float() < reflect.ValueOf(j).Float()
		} else {
			return reflect.ValueOf(i).Float() > reflect.ValueOf(j).Float()
		}
	case float32:
		if asc {
			return reflect.ValueOf(i).Float() < reflect.ValueOf(j).Float()
		} else {
			return reflect.ValueOf(i).Float() > reflect.ValueOf(j).Float()
		}
	case int:
		if asc {
			return reflect.ValueOf(i).Int() < reflect.ValueOf(j).Int()
		} else {
			return reflect.ValueOf(i).Int() > reflect.ValueOf(j).Int()
		}
	case int8:
		if asc {
			return reflect.ValueOf(i).Int() < reflect.ValueOf(j).Int()
		} else {
			return reflect.ValueOf(i).Int() > reflect.ValueOf(j).Int()
		}
	case int16:
		if asc {
			return reflect.ValueOf(i).Int() < reflect.ValueOf(j).Int()
		} else {
			return reflect.ValueOf(i).Int() > reflect.ValueOf(j).Int()
		}
	case int32:
		if asc {
			return reflect.ValueOf(i).Int() < reflect.ValueOf(j).Int()
		} else {
			return reflect.ValueOf(i).Int() > reflect.ValueOf(j).Int()
		}
	case int64:
		if asc {
			return reflect.ValueOf(i).Int() < reflect.ValueOf(j).Int()
		} else {
			return reflect.ValueOf(i).Int() > reflect.ValueOf(j).Int()
		}
		
	}
	return false
}

// sort
// main logic of quickSort
func sort(arr []interface{}, start, end int) {
	if (end - start) < 1 {
		return
	}
	
	pivot := arr[end]
	splitter := start
	
	for i := start; i < end; i++ {
		if less(arr[i], pivot, asc) {
			// swap
			arr[i], arr[splitter] = arr[splitter], arr[i]
			
			/*temp := arr[splitter]
			arr[splitter] = arr[i]
			arr[i] = temp*/
			
			splitter++
		}
	}
	
	arr[end] = arr[splitter]
	arr[splitter] = pivot
	
	sort(arr, start, splitter-1)
	sort(arr, splitter+1, end)
}
