
package main

import (
	"encoding/json"
	"fmt"
)

type Sample struct {
	Name string `json:"name"`
	Xid  string `json:"xid"`
}

func main() {
	var aSample Sample
	aSample.Name = "foo"
	aSample.Xid = "42"
        s, _ := json.Marshal(aSample)
	fmt.Printf("marshal Sample:%s\n", string(s))
        var sAry = make([]Sample, 2)
        sAry[0].Name = "one"
        sAry[0].Xid = "1"
        sAry[1].Name = "two"
        sAry[1].Xid = "2"
        s, _ = json.Marshal(sAry)
        fmt.Printf("marshal sAry:%s\n", string(s))
}
