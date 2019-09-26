# quick-sort
An implementation of quick sort in golang
## quick start
This package has a method for sorting which takes a slice of interface and returns sorted slice .
The default sort method is ASC , you would use the asc parameter of method as false for sorting in DESC mode.
### Installation 
get the package using go module :

`go get github.comg/ghoroubi/quick-sort`

import package in your project :

`import "github.com/ghoroubi/quick-sort"`
 
 ###Example
```
 package main
 
 import "github.com/ghoroubi/quick-sort" 
 
func main(){
 arr = []interface{}{
 		"c", "a", "z", "i", "b",
 	}
 }
quick.Sort(arr , true) // sort ascending
quick.Sort(arr,false) //sort Descending

}
 ```
 #####Enjoy using and star if you like
