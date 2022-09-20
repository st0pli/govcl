
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

type Exception struct {
    IObject
    instance unsafe.Pointer
}

// AsException
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsException(obj interface{}) *Exception {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &Exception{instance: instance}
}

func (e *Exception) _instance() uintptr {
    return uintptr(e.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (e *Exception) Instance() uintptr {
    return e._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (e *Exception) UnsafeAddr() unsafe.Pointer {
    return e.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (e *Exception) IsValid() bool {
    return e.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (e *Exception) Is() TIs {
    return TIs(e._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (e *Exception) As() TAs {
//    return TAs(e._instance())
//}

// ExceptionClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func ExceptionClass() TClass {
    return Exception_StaticClassType()
}

// ToString
//
// 文本类信息。
//
// Text information.
func (e *Exception) ToString() string {
    return Exception_ToString(e._instance())
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (e *Exception) ClassType() TClass {
    return Exception_ClassType(e._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (e *Exception) ClassName() string {
    return Exception_ClassName(e._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (e *Exception) InstanceSize() int32 {
    return Exception_InstanceSize(e._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (e *Exception) InheritsFrom(AClass TClass) bool {
    return Exception_InheritsFrom(e._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (e *Exception) Equals(Obj IObject) bool {
    return Exception_Equals(e._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (e *Exception) GetHashCode() int32 {
    return Exception_GetHashCode(e._instance())
}

// Message
//
// 获取异常消息。
func (e *Exception) Message() string {
    return Exception_GetMessage(e._instance())
}

// SetMessage
//
// 设置异常消息。
func (e *Exception) SetMessage(value string) {
    Exception_SetMessage(e._instance(), value)
}

