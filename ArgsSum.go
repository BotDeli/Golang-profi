package main
import ("fmt","os")
func main(){
  var n int
  for _, n := range os.StdArgs{
    sum += n
  }
  fmt.Println(sum)
}
