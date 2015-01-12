# automap
##Example
```go
package main

import (
  "fmt"
  "github.com/oneplus1000/automap"
)

type Dest struct {
	Id   int
	Name string
}

type Src struct {
	Id   int
	Name string
}

func main(){
  var d Dest
  var s Src
  s.Id = 1
  s.Name = "Tony"
  
  am := new(automap.AutoMapper)
  am.Auto(&s,&d)
  fmt.Printf("%v",d)
}



```
