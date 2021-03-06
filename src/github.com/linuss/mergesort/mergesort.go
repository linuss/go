package main

import(
  "fmt"
  "github.com/linuss/readfiles"
  "os"
  "time"
)

func merge(left []int, right []int) []int{
  result := make([]int, 0)
  for len(left) > 0 || len(right) > 0 {
    switch{
    case len(left) > 0 && len(right) > 0 :
      if left[0] <= right[0] {
        result = append(result, left[0])
        left = left[1:]
      }else{
        result = append(result, right[0])
        right = right[1:]
      }
    case len(left) > 0:
      result = append(result, left[0])
      left = left[1:]
    case len(right) > 0:
      result = append(result, right[0])
      right = right[1:]
    }
  }
  return result
}


func mergesort(nums []int) []int {
  if(len(nums) <= 1){
    return nums
  }
  left := make([]int, 0)
  right := make([]int, 0)
  middle := len(nums)/2

  for i := 0; i< middle; i++{
    left = append(left,nums[i])
  }
  for i := middle; i< len(nums); i++{
    right = append(right,nums[i])
  }
  left  = mergesort(left)
  right = mergesort(right)
  return merge(left,right)
}



func main(){
  args := os.Args
  if(len(args) != 2){
    fmt.Println("Too many arguments!") 
    fmt.Println("Usage: mergesort <filename>");
  }

  start := time.Now()
  nums,err := readfiles.ReadNums(args[1])
  if(err != nil){
    panic(err)
  }
  fmt.Printf("Reading in file took %s \n", time.Since(start).String())

  start = time.Now()
  nums = mergesort(nums)
  fmt.Printf("Sorting took %s \n", time.Since(start).String())
  for _,i := range nums {
    fmt.Printf("%d\n", i)
  }

}


