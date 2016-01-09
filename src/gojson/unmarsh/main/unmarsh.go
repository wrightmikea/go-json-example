
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
        "os"
)

type Sample1 struct {
	Name string `json:"name"`
	Foo  string `json:"foo"`
	Xid  int    `json:"xid"`
}
type Sample2 struct {
       Name string  `json:"name"`
       Alias string `json:"alias"`
       Xid   int    `json:"xid"`
}
func main() {
        if len(os.Args) != 3 {
              fmt.Println("usage > ./unmarsh src/sample1.json src/sample2.json")
              return
        }
	s1s, _ := ioutil.ReadFile(os.Args[1])
	dec1 := json.NewDecoder(bytes.NewReader(s1s))
	var s1 Sample1
	dec1.Decode(&s1)
	fmt.Printf("unmarshal JSON s1=%v\n", s1)
        s2s, _ := ioutil.ReadFile(os.Args[2])
        dec2 := json.NewDecoder(bytes.NewReader(s2s))
        var s2array []Sample2
        dec2.Decode(&s2array)
        fmt.Printf("unmarshal s2array=%v\n", s2array)
}
