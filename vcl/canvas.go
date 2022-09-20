
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

type TCanvas struct {
    IObject
    instance unsafe.Pointer
}

// NewCanvas
//
// 创建一个新的对象。
// 
// Create a new object.
func NewCanvas() *TCanvas {
    c := new(TCanvas)
    c.instance = unsafe.Pointer(Canvas_Create())
    setFinalizer(c, (*TCanvas).Free)
    return c
}

// AsCanvas
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsCanvas(obj interface{}) *TCanvas {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TCanvas{instance: instance}
}

// Free 
//
// 释放对象。
// 
// Free object.
func (c *TCanvas) Free() {
    if c.instance != nullptr {
        Canvas_Free(c._instance())
        c.instance  = nullptr
    }
}

func (c *TCanvas) _instance() uintptr {
    return uintptr(c.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (c *TCanvas) Instance() uintptr {
    return c._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (c *TCanvas) UnsafeAddr() unsafe.Pointer {
    return c.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (c *TCanvas) IsValid() bool {
    return c.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (c *TCanvas) Is() TIs {
    return TIs(c._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (c *TCanvas) As() TAs {
//    return TAs(c._instance())
//}

// TCanvasClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TCanvasClass() TClass {
    return Canvas_StaticClassType()
}

func (c *TCanvas) Arc(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_Arc(c._instance(), X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

func (c *TCanvas) ArcTo(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_ArcTo(c._instance(), X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

func (c *TCanvas) AngleArc(X int32, Y int32, Radius uint32, StartAngle float32, SweepAngle float32) {
    Canvas_AngleArc(c._instance(), X , Y , Radius , StartAngle , SweepAngle)
}

func (c *TCanvas) Chord(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_Chord(c._instance(), X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

func (c *TCanvas) Ellipse(X1 int32, Y1 int32, X2 int32, Y2 int32) {
    Canvas_Ellipse(c._instance(), X1 , Y1 , X2 , Y2)
}

func (c *TCanvas) FloodFill(X int32, Y int32, Color TColor, FillStyle TFillStyle) {
    Canvas_FloodFill(c._instance(), X , Y , Color , FillStyle)
}

// HandleAllocated
//
// 句柄是否已经分配。
//
// Is the handle already allocated.
func (c *TCanvas) HandleAllocated() bool {
    return Canvas_HandleAllocated(c._instance())
}

func (c *TCanvas) LineTo(X int32, Y int32) {
    Canvas_LineTo(c._instance(), X , Y)
}

func (c *TCanvas) MoveTo(X int32, Y int32) {
    Canvas_MoveTo(c._instance(), X , Y)
}

func (c *TCanvas) Pie(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_Pie(c._instance(), X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

func (c *TCanvas) Rectangle(X1 int32, Y1 int32, X2 int32, Y2 int32) {
    Canvas_Rectangle(c._instance(), X1 , Y1 , X2 , Y2)
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (c *TCanvas) Refresh() {
    Canvas_Refresh(c._instance())
}

func (c *TCanvas) RoundRect(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32) {
    Canvas_RoundRect(c._instance(), X1 , Y1 , X2 , Y2 , X3 , Y3)
}

func (c *TCanvas) StretchDraw(Rect TRect, Graphic IGraphic) {
    Canvas_StretchDraw(c._instance(), Rect , CheckPtr(Graphic))
}

func (c *TCanvas) TextExtent(Text string) TSize {
    return Canvas_TextExtent(c._instance(), Text)
}

func (c *TCanvas) TextOut(X int32, Y int32, Text string) {
    Canvas_TextOut(c._instance(), X , Y , Text)
}

func (c *TCanvas) Lock() {
    Canvas_Lock(c._instance())
}

func (c *TCanvas) TextHeight(Text string) int32 {
    return Canvas_TextHeight(c._instance(), Text)
}

func (c *TCanvas) TextWidth(Text string) int32 {
    return Canvas_TextWidth(c._instance(), Text)
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TCanvas) Assign(Source IObject) {
    Canvas_Assign(c._instance(), CheckPtr(Source))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (c *TCanvas) GetNamePath() string {
    return Canvas_GetNamePath(c._instance())
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (c *TCanvas) ClassType() TClass {
    return Canvas_ClassType(c._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TCanvas) ClassName() string {
    return Canvas_ClassName(c._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TCanvas) InstanceSize() int32 {
    return Canvas_InstanceSize(c._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TCanvas) InheritsFrom(AClass TClass) bool {
    return Canvas_InheritsFrom(c._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (c *TCanvas) Equals(Obj IObject) bool {
    return Canvas_Equals(c._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TCanvas) GetHashCode() int32 {
    return Canvas_GetHashCode(c._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (c *TCanvas) ToString() string {
    return Canvas_ToString(c._instance())
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (c *TCanvas) Handle() HDC {
    return Canvas_GetHandle(c._instance())
}

// SetHandle
//
// 设置控件句柄。
//
// Set Control handle.
func (c *TCanvas) SetHandle(value HDC) {
    Canvas_SetHandle(c._instance(), value)
}

// Brush
//
// 获取画刷对象。
//
// Get Brush.
func (c *TCanvas) Brush() *TBrush {
    return AsBrush(Canvas_GetBrush(c._instance()))
}

// SetBrush
//
// 设置画刷对象。
//
// Set Brush.
func (c *TCanvas) SetBrush(value *TBrush) {
    Canvas_SetBrush(c._instance(), CheckPtr(value))
}

func (c *TCanvas) CopyMode() int32 {
    return Canvas_GetCopyMode(c._instance())
}

func (c *TCanvas) SetCopyMode(value int32) {
    Canvas_SetCopyMode(c._instance(), value)
}

// Font
//
// 获取字体。
//
// Get Font.
func (c *TCanvas) Font() *TFont {
    return AsFont(Canvas_GetFont(c._instance()))
}

// SetFont
//
// 设置字体。
//
// Set Font.
func (c *TCanvas) SetFont(value *TFont) {
    Canvas_SetFont(c._instance(), CheckPtr(value))
}

func (c *TCanvas) Pen() *TPen {
    return AsPen(Canvas_GetPen(c._instance()))
}

func (c *TCanvas) SetPen(value *TPen) {
    Canvas_SetPen(c._instance(), CheckPtr(value))
}

// SetOnChange
//
// 设置改变事件。
//
// Set changed event.
func (c *TCanvas) SetOnChange(fn TNotifyEvent) {
    Canvas_SetOnChange(c._instance(), fn)
}

func (c *TCanvas) SetOnChanging(fn TNotifyEvent) {
    Canvas_SetOnChanging(c._instance(), fn)
}

func (c *TCanvas) Pixels(X int32, Y int32) TColor {
    return Canvas_GetPixels(c._instance(), X, Y)
}

func (c *TCanvas) SetPixels(X int32, Y int32, value TColor) {
    Canvas_SetPixels(c._instance(), X, Y, value)
}

