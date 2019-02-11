package go_infomap

import (
	// #cgo LDFLAGS: -L${SRCDIR} -lInfomap -fopenmp
	// #include "Infomap.h"
	// #include <stdlib.h>
	"C"
	"fmt"
	"runtime"
	"unsafe"
)

type State struct {
	p *C.struct_Infomap
}

func New(clusterNumber uint8) *State {
	var s = new(State)
	var flags = C.CString("--two-level --silent --preferred-number-of-modules "+fmt.Sprint(clusterNumber))
	defer C.free(unsafe.Pointer(flags))
	s.p = C.NewInfomap(flags)
	runtime.SetFinalizer(s, (*State).Destroy)
	return s
}

func (s *State) Destroy() {
	if s.p != nil {
		C.DestroyInfomap(s.p)
		s.p = nil
	}
}

func (s *State) AddLink(left uint64, right uint64, weight float64) {
	C.InfomapAddLink(s.p, C.uint(left), C.uint(right), C.double(weight))
}

func (s *State) Run() {
	C.InfomapRun(s.p)
}

func (s *State) Iter(f func(node uint64, module uint64, flow float64) bool) {
	var it = C.NewIter(s.p)
	defer C.DestroyIter(it)
	var cont = true
	for cont && bool(C.HaveNext(it)) {
		cont = f(uint64(C.NodeIndex(it)), uint64(C.ModuleIndex(it)), float64(C.Flow(it)))
		C.Next(it)
	}
}