package main
import ("fmt")

func main() {
  var a = make(map[string]string) // The map is empty now
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"
                                 // a is no longer empty
  b := make(map[string]int)
  b["Oslo"] = 1
  b["Bergen"] = 2
  b["Trondheim"] = 3
  b["Stavanger"] = 4

  fmt.Printf("a\t%v\n", a) // a	 map[brand:Ford model:Mustang year:1964]
  fmt.Printf("b\t%v\n", b) // b	 map[Bergen:2 Oslo:1 Stavanger:4 Trondheim:3]
}