package main

import(
  "fmt"
  "github.com/linuss/readfiles"
  "os"
  "time"
  "strconv"
)


func quicksort(nums []int, ch chan int, level int, threads int)  {
  /*For each level we go deeper, the amount of threads increases
  exponentially in powers of two. Level keeps track of how deep
  we are, allowing us to determine the number of threads ourselves */
  level *= 2;

  //Base case: empty array
  if len(nums)  < 1 {  close(ch); return }

  //Base case: only one element (done sorting)
  if len(nums) == 1 {  ch<- nums[0]; close(ch); return }

  less := make([]int, 0)
  greater := make([]int,0)
  pivot := nums[0]
  //Remove pivot from slice
  nums = nums[1:]

  //Create greater and lesser slices
  for _,i := range nums{
    switch{
    case i <= pivot:
      less = append(less,i)
    case i > pivot:
      greater = append(greater,i)
    }
  }

  ch1 := make(chan int, len(less))
  ch2 := make(chan int, len(greater))

  //Determine whether to create new threads
  if(level <= threads){
    go quicksort(less, ch1, level, threads)
    go quicksort(greater,ch2, level, threads)
  }else{
    quicksort(less,ch1, level, threads)
    quicksort(greater,ch2, level, threads)
  }

  //Concatenate results
  for i := range ch1{
    ch<-i;
  }
  ch<-pivot
  for i := range ch2{
    ch<-i;
  }

  //Close the channel
  close(ch)
  return
}


func main(){
  args := os.Args
  if(len(args) != 3){
    fmt.Println(args)
    fmt.Printf("Usage: quicksort <filename> <threads> \n")
  }

  start := time.Now()
  nums,err := readfiles.ReadNums(args[1])
  if(err != nil){
    panic(err)
  }
  fmt.Printf("Reading in file took %s \n", time.Since(start).String())

  start = time.Now()
  ch := make(chan int)
  threads, err := strconv.Atoi(args[2])
  if(err != nil){
    panic(err)
  }
  
  go quicksort(nums,ch,1,threads)
  <-ch

  fmt.Printf("Sorting took %s \n", time.Since(start).String())
  




}
  