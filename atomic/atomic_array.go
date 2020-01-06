package atomic

import (
	"fmt"
	"github.com/louyuting/go-sync/base"
	"sync/atomic"
	"unsafe"
)

type array struct {
	data []*entry
	// point to the first element in real array
	base unsafe.Pointer
	// length of array, must be const
	len int
}

// An entry is a element in the array.
type entry struct {
	p interface{}
}

func Array(len int) *array {
	ret := &array{
		data: make([]*entry, 0, len),
		base: nil,
		len:  len,
	}
	header := (*base.SliceHeader)(unsafe.Pointer(&ret.data))
	ret.base = unsafe.Pointer((**entry)(header.Data))
	return ret
}

func (a *array) offset(idx int) unsafe.Pointer {
	if idx < 0 || idx >= a.len {
		panic(fmt.Sprintf("The index (%d) is out of bounds, length is %d.\n", idx, a.len))
	}
	return unsafe.Pointer(uintptr(a.base) + uintptr(idx*8))
}

func (a *array) Load(idx int) interface{} {
	e := (*entry)(atomic.LoadPointer((*unsafe.Pointer)(a.offset(idx))))
	return e.p
}

func (a *array) Store(idx int, value interface{}) {
	e := unsafe.Pointer(&entry{p: value})
	atomic.StorePointer((*unsafe.Pointer)(a.offset(idx)), e)
}

func (a *array) CompareAndSwap(idx int, old, new interface{}) bool {
	o := unsafe.Pointer(&entry{p: old})
	n := unsafe.Pointer(&entry{p: new})
	return atomic.CompareAndSwapPointer((*unsafe.Pointer)(a.offset(idx)), o, n)
}
