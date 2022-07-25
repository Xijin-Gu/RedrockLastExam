/**
* @Author: gxj
* @Data: 2022/7/26-3:20
* @DESC:
**/

package jsone

import (
	"strconv"
	"testing"
)

func TestVtoJ(t *testing.T) {
	var want = [][]byte{
		{92,34,97,92,34},
		{92,34,98,99,92,34},
	}
	var in = []string{
		"a",
		"bc",
	}
	for i,v := range want{
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := VtoJ(in[i])
			if string(got) != string(v) {
				t.Errorf("got:%v",got)
			}
		})
	}
}

func TestJtoV(t *testing.T) {
	var in = [][]byte{
		{92,34,97,92,34},
		{92,34,98,99,92,34},
	}
	var want = []string{
		"a",
		"bc",
	}
	for i,_ := range in{
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := JtoV(in[i])
			if got != "\""+want[i]+"\"" {
				t.Errorf("got:%vwant:%v",got,want[i])
			}
		})
	}
}