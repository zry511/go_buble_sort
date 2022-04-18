package main

import (
	"fmt"
)

func main() {
  array := [] int{2,5,1,9,6,4,888,333,45,67,34,999,0,8,2,33,7};
  res := bubleSort(array)
	fmt.Println(res)

}

func bubleSort(array []int) []int {
    length :=len(array)
    for i:=0;i<length;i++ {
        for j:=0;j<length-i-1;j++ {
            if array[j] > array[j+1] {
                array[j+1],array[j] = array[j],array[j+1]
            }
        }
    }
    return array;
}