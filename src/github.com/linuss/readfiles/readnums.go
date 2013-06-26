package readfiles


import(
  "io/ioutil"
  "strings"
  "strconv"
)

//ReadNums reads numbers from a file and returns a slice of integers
func ReadNums(fname string) (nums []int, err error) {
  buf, err := ioutil.ReadFile(fname)
  if err != nil { return nil,err}

  lines := strings.Split(string(buf), "\n")

  nums = make([]int, 0, len(lines))

  for _,l := range(lines){
    if len(l) == 0 { continue }
    n, err := strconv.Atoi(l)
    if err != nil {return nil, err }
    nums = append(nums,n)
  }

  return nums, nil
}
