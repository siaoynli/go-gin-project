//Email: 120235331@qq.com
//Github: http：//www.github.com/siaoynli
//Date: 2020/5/14 9:37

package main

import "fmt"

func bin_search(num []int,search_data int) int{
    min:=0
    max:=len(num)-1
    for min<=max {
		mid:=(min+max)/2
		if  mid > search_data {
			max=mid
		}else if mid < search_data {
			min= mid
		}else{
			return  mid
		}
	}
	return -1
}



func main() {
	  var a = make([]int,1024*1024*1024,1024*1024*1024)
	  for i:=0;i<1024*1024*1024;i++ {
	  	   a[i]=i+1
	  }

	  index:= bin_search(a,10235623543245)

	  fmt.Println("查找到",index, a[index])
}
