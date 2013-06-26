package main

import(
  "fmt"
  "github.com/linuss/readfiles"
  "os"
  "time"
)

func concat(old1 []int, pivot int, old2 []int) []int{
  result := make([]int, 0)
  for i := range old1{
    result = append(result,old1[i])
  }
  result = append(result,pivot)
  for i := range old2{
    result = append(result,old2[i])
  }
  return result
}

func quicksort(nums []int) ([] int){
  if len(nums) <= 1 { return nums }
  less := make([]int, 0)
  greater := make([]int,0)
  pivot := nums[0]
  nums = nums[1:]

  for _,i := range nums{
    switch{
    case i <= pivot:
      less = append(less,i)
    case i > pivot:
      greater = append(greater,i)
    }
  }
  return concat(quicksort(less),pivot, quicksort(greater))
}


func main(){
  args := os.Args
  if(len(args) > 2){
    fmt.Println(args)
    fmt.Printf("Too many arguments!\nUsage: quicksort <filename>\n")
  }

  start := time.Now()
  nums,err := readfiles.ReadNums(args[1])
  if(err != nil){
    panic(err)
  }
  fmt.Printf("Reading in file took %s \n", time.Since(start).String())

  start = time.Now()
  quicksort(nums)

  fmt.Printf("Sorting took %s \n", time.Since(start).String())
  




}
  
