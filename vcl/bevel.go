
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

type TBevel struct {
    IControl
    instance unsafe.Pointer
}

// NewBevel
//
// 创建一个新的对象。
// 
// Create a new object.
func NewBevel(owner IComponent) *TBevel {
    b := new(TBevel)
    b.instance = unsafe.Pointer(Bevel_Create(CheckPtr(owner)))
    return b
}

// AsBevel
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsBevel(obj interface{}) *TBevel {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TBevel{instance: instance}
}

// Free 
//
// 释放对象。
// 
// Free object.
func (b *TBevel) Free() {
    if b.instance != nullptr {
        Bevel_Free(b._instance())
        b.instance  = nullptr
    }
}

func (b *TBevel) _instance() uintptr {
    return uintptr(b.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (b *TBevel) Instance() uintptr {
    return b._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (b *TBevel) UnsafeAddr() unsafe.Pointer {
    return b.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (b *TBevel) IsValid() bool {
    return b.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (b *TBevel) Is() TIs {
    return TIs(b._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (b *TBevel) As() TAs {
//    return TAs(b._instance())
//}

// TBevelClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TBevelClass() TClass {
    return Bevel_StaticClassType()
}

// BringToFront
//
// 将控件置于最前。
//
// Bring the control to the front.
func (b *TBevel) BringToFront() {
    Bevel_BringToFront(b._instance())
}

// ClientToScreen
//
// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (b *TBevel) ClientToScreen(Point TPoint) TPoint {
    return Bevel_ClientToScreen(b._instance(), Point)
}

// ClientToParent
//
// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (b *TBevel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Bevel_ClientToParent(b._instance(), Point , CheckPtr(AParent))
}

// Dragging
//
// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (b *TBevel) Dragging() bool {
    return Bevel_Dragging(b._instance())
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (b *TBevel) HasParent() bool {
    return Bevel_HasParent(b._instance())
}

// Hide
//
// 隐藏控件。
//
// Hidden control.
func (b *TBevel) Hide() {
    Bevel_Hide(b._instance())
}

// Invalidate
//
// 要求重绘。
//
// Redraw.
func (b *TBevel) Invalidate() {
    Bevel_Invalidate(b._instance())
}

// Perform
//
// 发送一个消息。
//
// Send a message.
func (b *TBevel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Bevel_Perform(b._instance(), Msg , WParam , LParam)
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (b *TBevel) Refresh() {
    Bevel_Refresh(b._instance())
}

// Repaint
//
// 重绘。
//
// Repaint.
func (b *TBevel) Repaint() {
    Bevel_Repaint(b._instance())
}

// ScreenToClient
//
// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (b *TBevel) ScreenToClient(Point TPoint) TPoint {
    return Bevel_ScreenToClient(b._instance(), Point)
}

// ParentToClient
//
// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (b *TBevel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Bevel_ParentToClient(b._instance(), Point , CheckPtr(AParent))
}

// SendToBack
//
// 控件至于最后面。
//
// The control is placed at the end.
func (b *TBevel) SendToBack() {
    Bevel_SendToBack(b._instance())
}

// SetBounds
//
// 设置组件边界。
//
// Set component boundaries.
func (b *TBevel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Bevel_SetBounds(b._instance(), ALeft , ATop , AWidth , AHeight)
}

// Show
//
// 显示控件。
//
// Show control.
func (b *TBevel) Show() {
    Bevel_Show(b._instance())
}

// Update
//
// 控件更新。
//
// Update.
func (b *TBevel) Update() {
    Bevel_Update(b._instance())
}

// GetTextBuf
//
// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (b *TBevel) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Bevel_GetTextBuf(b._instance(), Buffer , BufSize)
}

// GetTextLen
//
// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (b *TBevel) GetTextLen() int32 {
    return Bevel_GetTextLen(b._instance())
}

// SetTextBuf
//
// 设置控件字符，如果有。
//
// Set control characters, if any.
func (b *TBevel) SetTextBuf(Buffer string) {
    Bevel_SetTextBuf(b._instance(), Buffer)
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (b *TBevel) FindComponent(AName string) *TComponent {
    return AsComponent(Bevel_FindComponent(b._instance(), AName))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (b *TBevel) GetNamePath() string {
    return Bevel_GetNamePath(b._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (b *TBevel) Assign(Source IObject) {
    Bevel_Assign(b._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (b *TBevel) ClassType() TClass {
    return Bevel_ClassType(b._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (b *TBevel) ClassName() string {
    return Bevel_ClassName(b._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (b *TBevel) InstanceSize() int32 {
    return Bevel_InstanceSize(b._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (b *TBevel) InheritsFrom(AClass TClass) bool {
    return Bevel_InheritsFrom(b._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (b *TBevel) Equals(Obj IObject) bool {
    return Bevel_Equals(b._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (b *TBevel) GetHashCode() int32 {
    return Bevel_GetHashCode(b._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (b *TBevel) ToString() string {
    return Bevel_ToString(b._instance())
}

func (b *TBevel) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Bevel_AnchorToNeighbour(b._instance(), ASide , ASpace , CheckPtr(ASibling))
}

func (b *TBevel) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Bevel_AnchorParallel(b._instance(), ASide , ASpace , CheckPtr(ASibling))
}

// AnchorHorizontalCenterTo
//
// 置于指定控件的横向中心。
func (b *TBevel) AnchorHorizontalCenterTo(ASibling IControl) {
    Bevel_AnchorHorizontalCenterTo(b._instance(), CheckPtr(ASibling))
}

// AnchorVerticalCenterTo
//
// 置于指定控件的纵向中心。
func (b *TBevel) AnchorVerticalCenterTo(ASibling IControl) {
    Bevel_AnchorVerticalCenterTo(b._instance(), CheckPtr(ASibling))
}

func (b *TBevel) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    Bevel_AnchorSame(b._instance(), ASide , CheckPtr(ASibling))
}

func (b *TBevel) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    Bevel_AnchorAsAlign(b._instance(), ATheAlign , ASpace)
}

func (b *TBevel) AnchorClient(ASpace int32) {
    Bevel_AnchorClient(b._instance(), ASpace)
}

func (b *TBevel) ScaleDesignToForm(ASize int32) int32 {
    return Bevel_ScaleDesignToForm(b._instance(), ASize)
}

func (b *TBevel) ScaleFormToDesign(ASize int32) int32 {
    return Bevel_ScaleFormToDesign(b._instance(), ASize)
}

func (b *TBevel) Scale96ToForm(ASize int32) int32 {
    return Bevel_Scale96ToForm(b._instance(), ASize)
}

func (b *TBevel) ScaleFormTo96(ASize int32) int32 {
    return Bevel_ScaleFormTo96(b._instance(), ASize)
}

func (b *TBevel) Scale96ToFont(ASize int32) int32 {
    return Bevel_Scale96ToFont(b._instance(), ASize)
}

func (b *TBevel) ScaleFontTo96(ASize int32) int32 {
    return Bevel_ScaleFontTo96(b._instance(), ASize)
}

func (b *TBevel) ScaleScreenToFont(ASize int32) int32 {
    return Bevel_ScaleScreenToFont(b._instance(), ASize)
}

func (b *TBevel) ScaleFontToScreen(ASize int32) int32 {
    return Bevel_ScaleFontToScreen(b._instance(), ASize)
}

func (b *TBevel) Scale96ToScreen(ASize int32) int32 {
    return Bevel_Scale96ToScreen(b._instance(), ASize)
}

func (b *TBevel) ScaleScreenTo96(ASize int32) int32 {
    return Bevel_ScaleScreenTo96(b._instance(), ASize)
}

func (b *TBevel) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    Bevel_AutoAdjustLayout(b._instance(), AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (b *TBevel) FixDesignFontsPPI(ADesignTimePPI int32) {
    Bevel_FixDesignFontsPPI(b._instance(), ADesignTimePPI)
}

func (b *TBevel) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    Bevel_ScaleFontsPPI(b._instance(), AToPPI , AProportion)
}

// Align
//
// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (b *TBevel) Align() TAlign {
    return Bevel_GetAlign(b._instance())
}

// SetAlign
//
// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (b *TBevel) SetAlign(value TAlign) {
    Bevel_SetAlign(b._instance(), value)
}

// Anchors
//
// 获取四个角位置的锚点。
func (b *TBevel) Anchors() TAnchors {
    return Bevel_GetAnchors(b._instance())
}

// SetAnchors
//
// 设置四个角位置的锚点。
func (b *TBevel) SetAnchors(value TAnchors) {
    Bevel_SetAnchors(b._instance(), value)
}

// Constraints
//
// 获取约束控件大小。
func (b *TBevel) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Bevel_GetConstraints(b._instance()))
}

// SetConstraints
//
// 设置约束控件大小。
func (b *TBevel) SetConstraints(value *TSizeConstraints) {
    Bevel_SetConstraints(b._instance(), CheckPtr(value))
}

// ParentShowHint
//
// 获取以父容器的ShowHint属性为准。
func (b *TBevel) ParentShowHint() bool {
    return Bevel_GetParentShowHint(b._instance())
}

// SetParentShowHint
//
// 设置以父容器的ShowHint属性为准。
func (b *TBevel) SetParentShowHint(value bool) {
    Bevel_SetParentShowHint(b._instance(), value)
}

func (b *TBevel) Shape() TBevelShape {
    return Bevel_GetShape(b._instance())
}

func (b *TBevel) SetShape(value TBevelShape) {
    Bevel_SetShape(b._instance(), value)
}

// ShowHint
//
// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (b *TBevel) ShowHint() bool {
    return Bevel_GetShowHint(b._instance())
}

// SetShowHint
//
// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (b *TBevel) SetShowHint(value bool) {
    Bevel_SetShowHint(b._instance(), value)
}

func (b *TBevel) Style() TBevelStyle {
    return Bevel_GetStyle(b._instance())
}

func (b *TBevel) SetStyle(value TBevelStyle) {
    Bevel_SetStyle(b._instance(), value)
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (b *TBevel) Visible() bool {
    return Bevel_GetVisible(b._instance())
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (b *TBevel) SetVisible(value bool) {
    Bevel_SetVisible(b._instance(), value)
}

// Enabled
//
// 获取控件启用。
//
// Get the control enabled.
func (b *TBevel) Enabled() bool {
    return Bevel_GetEnabled(b._instance())
}

// SetEnabled
//
// 设置控件启用。
//
// Set the control enabled.
func (b *TBevel) SetEnabled(value bool) {
    Bevel_SetEnabled(b._instance(), value)
}

func (b *TBevel) Action() *TAction {
    return AsAction(Bevel_GetAction(b._instance()))
}

func (b *TBevel) SetAction(value IComponent) {
    Bevel_SetAction(b._instance(), CheckPtr(value))
}

func (b *TBevel) BiDiMode() TBiDiMode {
    return Bevel_GetBiDiMode(b._instance())
}

func (b *TBevel) SetBiDiMode(value TBiDiMode) {
    Bevel_SetBiDiMode(b._instance(), value)
}

func (b *TBevel) BoundsRect() TRect {
    return Bevel_GetBoundsRect(b._instance())
}

func (b *TBevel) SetBoundsRect(value TRect) {
    Bevel_SetBoundsRect(b._instance(), value)
}

// ClientHeight
//
// 获取客户区高度。
//
// Get client height.
func (b *TBevel) ClientHeight() int32 {
    return Bevel_GetClientHeight(b._instance())
}

// SetClientHeight
//
// 设置客户区高度。
//
// Set client height.
func (b *TBevel) SetClientHeight(value int32) {
    Bevel_SetClientHeight(b._instance(), value)
}

func (b *TBevel) ClientOrigin() TPoint {
    return Bevel_GetClientOrigin(b._instance())
}

// ClientRect
//
// 获取客户区矩形。
//
// Get client rectangle.
func (b *TBevel) ClientRect() TRect {
    return Bevel_GetClientRect(b._instance())
}

// ClientWidth
//
// 获取客户区宽度。
//
// Get client width.
func (b *TBevel) ClientWidth() int32 {
    return Bevel_GetClientWidth(b._instance())
}

// SetClientWidth
//
// 设置客户区宽度。
//
// Set client width.
func (b *TBevel) SetClientWidth(value int32) {
    Bevel_SetClientWidth(b._instance(), value)
}

// ControlState
//
// 获取控件状态。
//
// Get control state.
func (b *TBevel) ControlState() TControlState {
    return Bevel_GetControlState(b._instance())
}

// SetControlState
//
// 设置控件状态。
//
// Set control state.
func (b *TBevel) SetControlState(value TControlState) {
    Bevel_SetControlState(b._instance(), value)
}

// ControlStyle
//
// 获取控件样式。
//
// Get control style.
func (b *TBevel) ControlStyle() TControlStyle {
    return Bevel_GetControlStyle(b._instance())
}

// SetControlStyle
//
// 设置控件样式。
//
// Set control style.
func (b *TBevel) SetControlStyle(value TControlStyle) {
    Bevel_SetControlStyle(b._instance(), value)
}

func (b *TBevel) Floating() bool {
    return Bevel_GetFloating(b._instance())
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (b *TBevel) Parent() *TWinControl {
    return AsWinControl(Bevel_GetParent(b._instance()))
}

// SetParent
//
// 设置控件父容器。
//
// Set control parent container.
func (b *TBevel) SetParent(value IWinControl) {
    Bevel_SetParent(b._instance(), CheckPtr(value))
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (b *TBevel) Left() int32 {
    return Bevel_GetLeft(b._instance())
}

// SetLeft
//
// 设置左边位置。
//
// Set Left position.
func (b *TBevel) SetLeft(value int32) {
    Bevel_SetLeft(b._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (b *TBevel) Top() int32 {
    return Bevel_GetTop(b._instance())
}

// SetTop
//
// 设置顶边位置。
//
// Set Top position.
func (b *TBevel) SetTop(value int32) {
    Bevel_SetTop(b._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (b *TBevel) Width() int32 {
    return Bevel_GetWidth(b._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (b *TBevel) SetWidth(value int32) {
    Bevel_SetWidth(b._instance(), value)
}

// Height
//
// 获取高度。
//
// Get height.
func (b *TBevel) Height() int32 {
    return Bevel_GetHeight(b._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (b *TBevel) SetHeight(value int32) {
    Bevel_SetHeight(b._instance(), value)
}

// Cursor
//
// 获取控件光标。
//
// Get control cursor.
func (b *TBevel) Cursor() TCursor {
    return Bevel_GetCursor(b._instance())
}

// SetCursor
//
// 设置控件光标。
//
// Set control cursor.
func (b *TBevel) SetCursor(value TCursor) {
    Bevel_SetCursor(b._instance(), value)
}

// Hint
//
// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (b *TBevel) Hint() string {
    return Bevel_GetHint(b._instance())
}

// SetHint
//
// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (b *TBevel) SetHint(value string) {
    Bevel_SetHint(b._instance(), value)
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (b *TBevel) ComponentCount() int32 {
    return Bevel_GetComponentCount(b._instance())
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (b *TBevel) ComponentIndex() int32 {
    return Bevel_GetComponentIndex(b._instance())
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (b *TBevel) SetComponentIndex(value int32) {
    Bevel_SetComponentIndex(b._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (b *TBevel) Owner() *TComponent {
    return AsComponent(Bevel_GetOwner(b._instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (b *TBevel) Name() string {
    return Bevel_GetName(b._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (b *TBevel) SetName(value string) {
    Bevel_SetName(b._instance(), value)
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (b *TBevel) Tag() int {
    return Bevel_GetTag(b._instance())
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (b *TBevel) SetTag(value int) {
    Bevel_SetTag(b._instance(), value)
}

// AnchorSideLeft
//
// 获取左边锚点。
func (b *TBevel) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSideLeft(b._instance()))
}

// SetAnchorSideLeft
//
// 设置左边锚点。
func (b *TBevel) SetAnchorSideLeft(value *TAnchorSide) {
    Bevel_SetAnchorSideLeft(b._instance(), CheckPtr(value))
}

// AnchorSideTop
//
// 获取顶边锚点。
func (b *TBevel) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSideTop(b._instance()))
}

// SetAnchorSideTop
//
// 设置顶边锚点。
func (b *TBevel) SetAnchorSideTop(value *TAnchorSide) {
    Bevel_SetAnchorSideTop(b._instance(), CheckPtr(value))
}

// AnchorSideRight
//
// 获取右边锚点。
func (b *TBevel) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSideRight(b._instance()))
}

// SetAnchorSideRight
//
// 设置右边锚点。
func (b *TBevel) SetAnchorSideRight(value *TAnchorSide) {
    Bevel_SetAnchorSideRight(b._instance(), CheckPtr(value))
}

// AnchorSideBottom
//
// 获取底边锚点。
func (b *TBevel) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSideBottom(b._instance()))
}

// SetAnchorSideBottom
//
// 设置底边锚点。
func (b *TBevel) SetAnchorSideBottom(value *TAnchorSide) {
    Bevel_SetAnchorSideBottom(b._instance(), CheckPtr(value))
}

// BorderSpacing
//
// 获取边框间距。
func (b *TBevel) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(Bevel_GetBorderSpacing(b._instance()))
}

// SetBorderSpacing
//
// 设置边框间距。
func (b *TBevel) SetBorderSpacing(value *TControlBorderSpacing) {
    Bevel_SetBorderSpacing(b._instance(), CheckPtr(value))
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (b *TBevel) Components(AIndex int32) *TComponent {
    return AsComponent(Bevel_GetComponents(b._instance(), AIndex))
}

// AnchorSide
//
// 获取锚侧面。
func (b *TBevel) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(Bevel_GetAnchorSide(b._instance(), AKind))
}

