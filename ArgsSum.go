package main
import ("fmt","os","strconv")
func main(){
  var n int
  for _, n := range os.StdArgs[1:]{
    num, _ = strconv.Atoi(n)
    sum += num
  }
  fmt.Println(sum)
}
