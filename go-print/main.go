package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/samber/lo"
)

type a struct {
	b int
	c string
}

type aa []*a

func (v aa) String() string {
	var tmp []string
	lo.ForEach(v, func(val *a, _ int) {
		tmp = append(tmp, fmt.Sprintf("%v", val))
	})
	return strings.Join(tmp, ",")
}

func main() {
	a1 := &a{b: 1, c: "a"}
	a2 := &a{b: 2, c: "b"}
	a3 := &a{b: 3, c: "c"}

	aa1 := aa{a1, a2, a3}
	spew.Sprintf("%v", aa1)

	f := 10
	g := "aaa"
	b := map[string]any{
		"CurrentUserMonqDeckId": f,
		"UserId":                g,
	}
	bb, _ := json.Marshal(b)
	log.Printf("%v", string(bb))

	fmt.Printf("%+v", aa1)
}
