
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

type TFlowPanelControl struct {
    IObject
    instance unsafe.Pointer
}

// AsFlowPanelControl
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsFlowPanelControl(obj interface{}) *TFlowPanelControl {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TFlowPanelControl{instance: instance}
}

func (f *TFlowPanelControl) _instance() uintptr {
    return uintptr(f.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (f *TFlowPanelControl) Instance() uintptr {
    return f._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (f *TFlowPanelControl) UnsafeAddr() unsafe.Pointer {
    return f.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (f *TFlowPanelControl) IsValid() bool {
    return f.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (f *TFlowPanelControl) Is() TIs {
    return TIs(f._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (f *TFlowPanelControl) As() TAs {
//    return TAs(f._instance())
//}

// TFlowPanelControlClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TFlowPanelControlClass() TClass {
    return FlowPanelControl_StaticClassType()
}

func (f *TFlowPanelControl) AllowAdd() bool {
    return FlowPanelControl_AllowAdd(f._instance())
}

func (f *TFlowPanelControl) AllowDelete() bool {
    return FlowPanelControl_AllowDelete(f._instance())
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (f *TFlowPanelControl) GetNamePath() string {
    return FlowPanelControl_GetNamePath(f._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (f *TFlowPanelControl) Assign(Source IObject) {
    FlowPanelControl_Assign(f._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (f *TFlowPanelControl) ClassType() TClass {
    return FlowPanelControl_ClassType(f._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (f *TFlowPanelControl) ClassName() string {
    return FlowPanelControl_ClassName(f._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (f *TFlowPanelControl) InstanceSize() int32 {
    return FlowPanelControl_InstanceSize(f._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (f *TFlowPanelControl) InheritsFrom(AClass TClass) bool {
    return FlowPanelControl_InheritsFrom(f._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (f *TFlowPanelControl) Equals(Obj IObject) bool {
    return FlowPanelControl_Equals(f._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (f *TFlowPanelControl) GetHashCode() int32 {
    return FlowPanelControl_GetHashCode(f._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (f *TFlowPanelControl) ToString() string {
    return FlowPanelControl_ToString(f._instance())
}

func (f *TFlowPanelControl) Control() *TControl {
    return AsControl(FlowPanelControl_GetControl(f._instance()))
}

func (f *TFlowPanelControl) SetControl(value IControl) {
    FlowPanelControl_SetControl(f._instance(), CheckPtr(value))
}

func (f *TFlowPanelControl) WrapAfter() TWrapAfter {
    return FlowPanelControl_GetWrapAfter(f._instance())
}

func (f *TFlowPanelControl) SetWrapAfter(value TWrapAfter) {
    FlowPanelControl_SetWrapAfter(f._instance(), value)
}

func (f *TFlowPanelControl) Index() int32 {
    return FlowPanelControl_GetIndex(f._instance())
}

func (f *TFlowPanelControl) SetIndex(value int32) {
    FlowPanelControl_SetIndex(f._instance(), value)
}

func (f *TFlowPanelControl) Collection() *TCollection {
    return AsCollection(FlowPanelControl_GetCollection(f._instance()))
}

func (f *TFlowPanelControl) SetCollection(value *TCollection) {
    FlowPanelControl_SetCollection(f._instance(), CheckPtr(value))
}

func (f *TFlowPanelControl) DisplayName() string {
    return FlowPanelControl_GetDisplayName(f._instance())
}

func (f *TFlowPanelControl) SetDisplayName(value string) {
    FlowPanelControl_SetDisplayName(f._instance(), value)
}

