package main

import "fmt"

func main() {

  m := make(map[string]int)


  m["k1"] = 7
  m["k2"] = 13

  fmt.Println("map:", m)


  v1 := m["k1"]
  fmt.Println("v1: ", v1) // 7

  fmt.Println("len:", len(m)) // 2


  delete(m, "k2")
  fmt.Println("map:", m) // map[k1:7]

  _, prs := m["k2"]
  fmt.Println("prs:", prs) // false


  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n) // map[foo:1 bar:2]
}

