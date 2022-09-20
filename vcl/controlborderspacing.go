
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

type TControlBorderSpacing struct {
    IObject
    instance unsafe.Pointer
}

// AsControlBorderSpacing
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsControlBorderSpacing(obj interface{}) *TControlBorderSpacing {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TControlBorderSpacing{instance: instance}
}

func (c *TControlBorderSpacing) _instance() uintptr {
    return uintptr(c.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (c *TControlBorderSpacing) Instance() uintptr {
    return c._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (c *TControlBorderSpacing) UnsafeAddr() unsafe.Pointer {
    return c.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (c *TControlBorderSpacing) IsValid() bool {
    return c.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (c *TControlBorderSpacing) Is() TIs {
    return TIs(c._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (c *TControlBorderSpacing) As() TAs {
//    return TAs(c._instance())
//}

// TControlBorderSpacingClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TControlBorderSpacingClass() TClass {
    return ControlBorderSpacing_StaticClassType()
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TControlBorderSpacing) Assign(Source IObject) {
    ControlBorderSpacing_Assign(c._instance(), CheckPtr(Source))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (c *TControlBorderSpacing) GetNamePath() string {
    return ControlBorderSpacing_GetNamePath(c._instance())
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (c *TControlBorderSpacing) ClassType() TClass {
    return ControlBorderSpacing_ClassType(c._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TControlBorderSpacing) ClassName() string {
    return ControlBorderSpacing_ClassName(c._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TControlBorderSpacing) InstanceSize() int32 {
    return ControlBorderSpacing_InstanceSize(c._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TControlBorderSpacing) InheritsFrom(AClass TClass) bool {
    return ControlBorderSpacing_InheritsFrom(c._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (c *TControlBorderSpacing) Equals(Obj IObject) bool {
    return ControlBorderSpacing_Equals(c._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TControlBorderSpacing) GetHashCode() int32 {
    return ControlBorderSpacing_GetHashCode(c._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (c *TControlBorderSpacing) ToString() string {
    return ControlBorderSpacing_ToString(c._instance())
}

func (c *TControlBorderSpacing) Control() *TControl {
    return AsControl(ControlBorderSpacing_GetControl(c._instance()))
}

func (c *TControlBorderSpacing) AroundLeft() int32 {
    return ControlBorderSpacing_GetAroundLeft(c._instance())
}

func (c *TControlBorderSpacing) AroundTop() int32 {
    return ControlBorderSpacing_GetAroundTop(c._instance())
}

func (c *TControlBorderSpacing) AroundRight() int32 {
    return ControlBorderSpacing_GetAroundRight(c._instance())
}

func (c *TControlBorderSpacing) AroundBottom() int32 {
    return ControlBorderSpacing_GetAroundBottom(c._instance())
}

func (c *TControlBorderSpacing) ControlLeft() int32 {
    return ControlBorderSpacing_GetControlLeft(c._instance())
}

func (c *TControlBorderSpacing) ControlTop() int32 {
    return ControlBorderSpacing_GetControlTop(c._instance())
}

func (c *TControlBorderSpacing) ControlWidth() int32 {
    return ControlBorderSpacing_GetControlWidth(c._instance())
}

func (c *TControlBorderSpacing) ControlHeight() int32 {
    return ControlBorderSpacing_GetControlHeight(c._instance())
}

func (c *TControlBorderSpacing) ControlRight() int32 {
    return ControlBorderSpacing_GetControlRight(c._instance())
}

func (c *TControlBorderSpacing) ControlBottom() int32 {
    return ControlBorderSpacing_GetControlBottom(c._instance())
}

// SetOnChange
//
// 设置改变事件。
//
// Set changed event.
func (c *TControlBorderSpacing) SetOnChange(fn TNotifyEvent) {
    ControlBorderSpacing_SetOnChange(c._instance(), fn)
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (c *TControlBorderSpacing) Left() int32 {
    return ControlBorderSpacing_GetLeft(c._instance())
}

// SetLeft
//
// 设置左边位置。
//
// Set Left position.
func (c *TControlBorderSpacing) SetLeft(value int32) {
    ControlBorderSpacing_SetLeft(c._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (c *TControlBorderSpacing) Top() int32 {
    return ControlBorderSpacing_GetTop(c._instance())
}

// SetTop
//
// 设置顶边位置。
//
// Set Top position.
func (c *TControlBorderSpacing) SetTop(value int32) {
    ControlBorderSpacing_SetTop(c._instance(), value)
}

func (c *TControlBorderSpacing) Right() int32 {
    return ControlBorderSpacing_GetRight(c._instance())
}

func (c *TControlBorderSpacing) SetRight(value int32) {
    ControlBorderSpacing_SetRight(c._instance(), value)
}

func (c *TControlBorderSpacing) Bottom() int32 {
    return ControlBorderSpacing_GetBottom(c._instance())
}

func (c *TControlBorderSpacing) SetBottom(value int32) {
    ControlBorderSpacing_SetBottom(c._instance(), value)
}

func (c *TControlBorderSpacing) Around() int32 {
    return ControlBorderSpacing_GetAround(c._instance())
}

func (c *TControlBorderSpacing) SetAround(value int32) {
    ControlBorderSpacing_SetAround(c._instance(), value)
}

func (c *TControlBorderSpacing) InnerBorder() int32 {
    return ControlBorderSpacing_GetInnerBorder(c._instance())
}

func (c *TControlBorderSpacing) SetInnerBorder(value int32) {
    ControlBorderSpacing_SetInnerBorder(c._instance(), value)
}

func (c *TControlBorderSpacing) CellAlignHorizontal() TControlCellAlign {
    return ControlBorderSpacing_GetCellAlignHorizontal(c._instance())
}

func (c *TControlBorderSpacing) SetCellAlignHorizontal(value TControlCellAlign) {
    ControlBorderSpacing_SetCellAlignHorizontal(c._instance(), value)
}

func (c *TControlBorderSpacing) CellAlignVertical() TControlCellAlign {
    return ControlBorderSpacing_GetCellAlignVertical(c._instance())
}

func (c *TControlBorderSpacing) SetCellAlignVertical(value TControlCellAlign) {
    ControlBorderSpacing_SetCellAlignVertical(c._instance(), value)
}

func (c *TControlBorderSpacing) Space(Kind TAnchorKind) int32 {
    return ControlBorderSpacing_GetSpace(c._instance(), Kind)
}

func (c *TControlBorderSpacing) SetSpace(Kind TAnchorKind, value int32) {
    ControlBorderSpacing_SetSpace(c._instance(), Kind, value)
}

