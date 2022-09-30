//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"sync"
	"unsafe"
)

var (
	// 防止GC的表
	events = sync.Map{}
	// 初始id
	eventNextId uintptr = 0
	//events = list.New()

	// ThreadSync
	threadSyncProc func()

	// sync
	idRefs, threadSyncRef sync.Mutex
)

func GoBool(val uintptr) bool {
	if val != 0 {
		return true
	}
	return false
}

func PascalBool(val bool) uintptr {
	if val {
		return 1
	}
	return 0
}

// typedef struct { void *type; void *value; } GoInterface;
type interfacePtr struct {
	tpy uintptr
	val uintptr
}

func interfaceNotNil(v interface{}) bool {
	ptr := (*interfacePtr)(unsafe.Pointer(&v))
	return ptr != nil && ptr.tpy != 0 && ptr.val != 0
}

//func PtrToElementPtr(v unsafe.Pointer) *list.Element {
//	if v == nil {
//		return nil
//	}
//	return (*list.Element)(v)
//}
//
//func PtrToElementValue(v unsafe.Pointer) interface{} {
//	element := PtrToElementPtr(v)
//	if element != nil {
//		return element.Value
//	}
//	return nil
//}
//
//func RemoveEventElement(v *list.Element) bool {
//	if v != nil {
//		idRefs.Lock()
//		defer idRefs.Unlock()
//		events.Remove(v)
//		return true
//	}
//	return false
//}
//
//func MakeEventDataPtr(fn interface{}) uintptr {
//
//	idRefs.Lock()
//	defer idRefs.Unlock()
//	if interfaceNotNil(fn) {
//		return uintptr(unsafe.Pointer(events.PushBack(fn)))
//	}
//	return 0
//}

func PtrToElementValue(v unsafe.Pointer) interface{} {
	if fn, ok := events.Load(uintptr(v)); ok {
		return fn
	}
	return nil
}

func RemoveEventElement(v unsafe.Pointer) bool {
	idRefs.Lock()
	defer idRefs.Unlock()
	if v != nil {
		events.Delete(uintptr(v))
		return true
	}
	return false
}

// MakeEventDataPtr
//  本想直接返回一个Go的指针，岂料Linux下Go强制不允许这样操作，一定要做就必须先设置环境变量`GODEBUG=cgocheck=0`关闭，但用户肯定不可能每次弄个啊
//  关于不允许传入Go指针到C指针的原因： https://chai2010.cn/advanced-go-programming-book/ch2-cgo/ch2-07-memory.html
func MakeEventDataPtr(fn interface{}) uintptr {
	idRefs.Lock()
	defer idRefs.Unlock()
	if interfaceNotNil(fn) {
		eventNextId++
		events.Store(eventNextId, fn)
		return eventNextId
	}
	return 0
}

func ThreadSyncCallbackFn() func() {
	return threadSyncProc
}
