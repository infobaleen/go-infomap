package go_infomap_test

import (
	"github.com/infobaleen/go-infomap"
)
import "testing"

func TestMinimal(t *testing.T) {
	var s = go_infomap.New(2)
	s.AddLink(0,1,1)
	s.AddLink(1,0,1)
	s.AddLink(2,3,1)
	s.AddLink(3,2,1)
	s.Run()
	var cluster [4]byte
	s.Iter(func(node uint64, module uint64, flow float64) bool {
		cluster[node] = byte(module)
		return true
	})
	if !(cluster[0] == cluster[1] && cluster[2] == cluster[3] && cluster[0] != cluster[2]) {
		t.Fatal("unexpected module membership: ", cluster)
	}
}
