package main
import "fmt"

func fizzBuzz(){
  for i := 1; i < 100; i++  {
    if i % 3  == 0 && i % 5 == 0 {
      fmt.Println("FIZZBUZZ!\n")
    } else if i % 3 == 0 {
      fmt.Println("FIZZ!\n")
    } else if i % 5 == 0 {
      fmt.Println("BUZZ!\n")
    } else {
      fmt.Println(i)
    }
  }
}
