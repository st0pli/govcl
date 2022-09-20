
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TListItems struct {
    IObject
    instance unsafe.Pointer
}

// NewListItems
//
// 创建一个新的对象。
// 
// Create a new object.
func NewListItems(AOwner *TListView) *TListItems {
    l := new(TListItems)
    l.instance = unsafe.Pointer(ListItems_Create(CheckPtr(AOwner)))
    setFinalizer(l, (*TListItems).Free)
    return l
}

// AsListItems
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsListItems(obj interface{}) *TListItems {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TListItems{instance: instance}
}

// Free 
//
// 释放对象。
// 
// Free object.
func (l *TListItems) Free() {
    if l.instance != nullptr {
        ListItems_Free(l._instance())
        l.instance  = nullptr
    }
}

func (l *TListItems) _instance() uintptr {
    return uintptr(l.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (l *TListItems) Instance() uintptr {
    return l._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (l *TListItems) UnsafeAddr() unsafe.Pointer {
    return l.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (l *TListItems) IsValid() bool {
    return l.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (l *TListItems) Is() TIs {
    return TIs(l._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (l *TListItems) As() TAs {
//    return TAs(l._instance())
//}

// TListItemsClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TListItemsClass() TClass {
    return ListItems_StaticClassType()
}

func (l *TListItems) Add() *TListItem {
    return AsListItem(ListItems_Add(l._instance()))
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (l *TListItems) Assign(Source IObject) {
    ListItems_Assign(l._instance(), CheckPtr(Source))
}

func (l *TListItems) BeginUpdate() {
    ListItems_BeginUpdate(l._instance())
}

// Clear
//
// 清除。
func (l *TListItems) Clear() {
    ListItems_Clear(l._instance())
}

func (l *TListItems) Delete(Index int32) {
    ListItems_Delete(l._instance(), Index)
}

func (l *TListItems) EndUpdate() {
    ListItems_EndUpdate(l._instance())
}

func (l *TListItems) IndexOf(Value *TListItem) int32 {
    return ListItems_IndexOf(l._instance(), CheckPtr(Value))
}

func (l *TListItems) Insert(Index int32) *TListItem {
    return AsListItem(ListItems_Insert(l._instance(), Index))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (l *TListItems) GetNamePath() string {
    return ListItems_GetNamePath(l._instance())
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (l *TListItems) ClassType() TClass {
    return ListItems_ClassType(l._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (l *TListItems) ClassName() string {
    return ListItems_ClassName(l._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (l *TListItems) InstanceSize() int32 {
    return ListItems_InstanceSize(l._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (l *TListItems) InheritsFrom(AClass TClass) bool {
    return ListItems_InheritsFrom(l._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (l *TListItems) Equals(Obj IObject) bool {
    return ListItems_Equals(l._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (l *TListItems) GetHashCode() int32 {
    return ListItems_GetHashCode(l._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (l *TListItems) ToString() string {
    return ListItems_ToString(l._instance())
}

func (l *TListItems) Count() int32 {
    return ListItems_GetCount(l._instance())
}

func (l *TListItems) SetCount(value int32) {
    ListItems_SetCount(l._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (l *TListItems) Owner() *TWinControl {
    return AsWinControl(ListItems_GetOwner(l._instance()))
}

func (l *TListItems) Item(Index int32) *TListItem {
    return AsListItem(ListItems_GetItem(l._instance(), Index))
}

func (l *TListItems) SetItem(Index int32, value *TListItem) {
    ListItems_SetItem(l._instance(), Index, CheckPtr(value))
}

