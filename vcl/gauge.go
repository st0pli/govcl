
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

type TGauge struct {
    IControl
    instance unsafe.Pointer
}

// NewGauge
//
// 创建一个新的对象。
// 
// Create a new object.
func NewGauge(owner IComponent) *TGauge {
    g := new(TGauge)
    g.instance = unsafe.Pointer(Gauge_Create(CheckPtr(owner)))
    return g
}

// AsGauge
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsGauge(obj interface{}) *TGauge {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TGauge{instance: instance}
}

// Free 
//
// 释放对象。
// 
// Free object.
func (g *TGauge) Free() {
    if g.instance != nullptr {
        Gauge_Free(g._instance())
        g.instance  = nullptr
    }
}

func (g *TGauge) _instance() uintptr {
    return uintptr(g.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (g *TGauge) Instance() uintptr {
    return g._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (g *TGauge) UnsafeAddr() unsafe.Pointer {
    return g.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (g *TGauge) IsValid() bool {
    return g.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (g *TGauge) Is() TIs {
    return TIs(g._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (g *TGauge) As() TAs {
//    return TAs(g._instance())
//}

// TGaugeClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TGaugeClass() TClass {
    return Gauge_StaticClassType()
}

func (g *TGauge) AddProgress(Value int32) {
    Gauge_AddProgress(g._instance(), Value)
}

// BringToFront
//
// 将控件置于最前。
//
// Bring the control to the front.
func (g *TGauge) BringToFront() {
    Gauge_BringToFront(g._instance())
}

// ClientToScreen
//
// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (g *TGauge) ClientToScreen(Point TPoint) TPoint {
    return Gauge_ClientToScreen(g._instance(), Point)
}

// ClientToParent
//
// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (g *TGauge) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Gauge_ClientToParent(g._instance(), Point , CheckPtr(AParent))
}

// Dragging
//
// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (g *TGauge) Dragging() bool {
    return Gauge_Dragging(g._instance())
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (g *TGauge) HasParent() bool {
    return Gauge_HasParent(g._instance())
}

// Hide
//
// 隐藏控件。
//
// Hidden control.
func (g *TGauge) Hide() {
    Gauge_Hide(g._instance())
}

// Invalidate
//
// 要求重绘。
//
// Redraw.
func (g *TGauge) Invalidate() {
    Gauge_Invalidate(g._instance())
}

// Perform
//
// 发送一个消息。
//
// Send a message.
func (g *TGauge) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Gauge_Perform(g._instance(), Msg , WParam , LParam)
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (g *TGauge) Refresh() {
    Gauge_Refresh(g._instance())
}

// Repaint
//
// 重绘。
//
// Repaint.
func (g *TGauge) Repaint() {
    Gauge_Repaint(g._instance())
}

// ScreenToClient
//
// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (g *TGauge) ScreenToClient(Point TPoint) TPoint {
    return Gauge_ScreenToClient(g._instance(), Point)
}

// ParentToClient
//
// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (g *TGauge) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Gauge_ParentToClient(g._instance(), Point , CheckPtr(AParent))
}

// SendToBack
//
// 控件至于最后面。
//
// The control is placed at the end.
func (g *TGauge) SendToBack() {
    Gauge_SendToBack(g._instance())
}

// SetBounds
//
// 设置组件边界。
//
// Set component boundaries.
func (g *TGauge) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Gauge_SetBounds(g._instance(), ALeft , ATop , AWidth , AHeight)
}

// Show
//
// 显示控件。
//
// Show control.
func (g *TGauge) Show() {
    Gauge_Show(g._instance())
}

// Update
//
// 控件更新。
//
// Update.
func (g *TGauge) Update() {
    Gauge_Update(g._instance())
}

// GetTextBuf
//
// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (g *TGauge) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Gauge_GetTextBuf(g._instance(), Buffer , BufSize)
}

// GetTextLen
//
// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (g *TGauge) GetTextLen() int32 {
    return Gauge_GetTextLen(g._instance())
}

// SetTextBuf
//
// 设置控件字符，如果有。
//
// Set control characters, if any.
func (g *TGauge) SetTextBuf(Buffer string) {
    Gauge_SetTextBuf(g._instance(), Buffer)
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (g *TGauge) FindComponent(AName string) *TComponent {
    return AsComponent(Gauge_FindComponent(g._instance(), AName))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (g *TGauge) GetNamePath() string {
    return Gauge_GetNamePath(g._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (g *TGauge) Assign(Source IObject) {
    Gauge_Assign(g._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (g *TGauge) ClassType() TClass {
    return Gauge_ClassType(g._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (g *TGauge) ClassName() string {
    return Gauge_ClassName(g._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (g *TGauge) InstanceSize() int32 {
    return Gauge_InstanceSize(g._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (g *TGauge) InheritsFrom(AClass TClass) bool {
    return Gauge_InheritsFrom(g._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (g *TGauge) Equals(Obj IObject) bool {
    return Gauge_Equals(g._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (g *TGauge) GetHashCode() int32 {
    return Gauge_GetHashCode(g._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (g *TGauge) ToString() string {
    return Gauge_ToString(g._instance())
}

func (g *TGauge) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Gauge_AnchorToNeighbour(g._instance(), ASide , ASpace , CheckPtr(ASibling))
}

func (g *TGauge) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Gauge_AnchorParallel(g._instance(), ASide , ASpace , CheckPtr(ASibling))
}

// AnchorHorizontalCenterTo
//
// 置于指定控件的横向中心。
func (g *TGauge) AnchorHorizontalCenterTo(ASibling IControl) {
    Gauge_AnchorHorizontalCenterTo(g._instance(), CheckPtr(ASibling))
}

// AnchorVerticalCenterTo
//
// 置于指定控件的纵向中心。
func (g *TGauge) AnchorVerticalCenterTo(ASibling IControl) {
    Gauge_AnchorVerticalCenterTo(g._instance(), CheckPtr(ASibling))
}

func (g *TGauge) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    Gauge_AnchorSame(g._instance(), ASide , CheckPtr(ASibling))
}

func (g *TGauge) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    Gauge_AnchorAsAlign(g._instance(), ATheAlign , ASpace)
}

func (g *TGauge) AnchorClient(ASpace int32) {
    Gauge_AnchorClient(g._instance(), ASpace)
}

func (g *TGauge) ScaleDesignToForm(ASize int32) int32 {
    return Gauge_ScaleDesignToForm(g._instance(), ASize)
}

func (g *TGauge) ScaleFormToDesign(ASize int32) int32 {
    return Gauge_ScaleFormToDesign(g._instance(), ASize)
}

func (g *TGauge) Scale96ToForm(ASize int32) int32 {
    return Gauge_Scale96ToForm(g._instance(), ASize)
}

func (g *TGauge) ScaleFormTo96(ASize int32) int32 {
    return Gauge_ScaleFormTo96(g._instance(), ASize)
}

func (g *TGauge) Scale96ToFont(ASize int32) int32 {
    return Gauge_Scale96ToFont(g._instance(), ASize)
}

func (g *TGauge) ScaleFontTo96(ASize int32) int32 {
    return Gauge_ScaleFontTo96(g._instance(), ASize)
}

func (g *TGauge) ScaleScreenToFont(ASize int32) int32 {
    return Gauge_ScaleScreenToFont(g._instance(), ASize)
}

func (g *TGauge) ScaleFontToScreen(ASize int32) int32 {
    return Gauge_ScaleFontToScreen(g._instance(), ASize)
}

func (g *TGauge) Scale96ToScreen(ASize int32) int32 {
    return Gauge_Scale96ToScreen(g._instance(), ASize)
}

func (g *TGauge) ScaleScreenTo96(ASize int32) int32 {
    return Gauge_ScaleScreenTo96(g._instance(), ASize)
}

func (g *TGauge) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    Gauge_AutoAdjustLayout(g._instance(), AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (g *TGauge) FixDesignFontsPPI(ADesignTimePPI int32) {
    Gauge_FixDesignFontsPPI(g._instance(), ADesignTimePPI)
}

func (g *TGauge) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    Gauge_ScaleFontsPPI(g._instance(), AToPPI , AProportion)
}

func (g *TGauge) PercentDone() int32 {
    return Gauge_GetPercentDone(g._instance())
}

// Align
//
// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (g *TGauge) Align() TAlign {
    return Gauge_GetAlign(g._instance())
}

// SetAlign
//
// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (g *TGauge) SetAlign(value TAlign) {
    Gauge_SetAlign(g._instance(), value)
}

// Anchors
//
// 获取四个角位置的锚点。
func (g *TGauge) Anchors() TAnchors {
    return Gauge_GetAnchors(g._instance())
}

// SetAnchors
//
// 设置四个角位置的锚点。
func (g *TGauge) SetAnchors(value TAnchors) {
    Gauge_SetAnchors(g._instance(), value)
}

func (g *TGauge) BackColor() TColor {
    return Gauge_GetBackColor(g._instance())
}

func (g *TGauge) SetBackColor(value TColor) {
    Gauge_SetBackColor(g._instance(), value)
}

// BorderStyle
//
// 获取窗口边框样式。比如：无边框，单一边框等。
func (g *TGauge) BorderStyle() TBorderStyle {
    return Gauge_GetBorderStyle(g._instance())
}

// SetBorderStyle
//
// 设置窗口边框样式。比如：无边框，单一边框等。
func (g *TGauge) SetBorderStyle(value TBorderStyle) {
    Gauge_SetBorderStyle(g._instance(), value)
}

// Color
//
// 获取颜色。
//
// Get color.
func (g *TGauge) Color() TColor {
    return Gauge_GetColor(g._instance())
}

// SetColor
//
// 设置颜色。
//
// Set color.
func (g *TGauge) SetColor(value TColor) {
    Gauge_SetColor(g._instance(), value)
}

// Constraints
//
// 获取约束控件大小。
func (g *TGauge) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Gauge_GetConstraints(g._instance()))
}

// SetConstraints
//
// 设置约束控件大小。
func (g *TGauge) SetConstraints(value *TSizeConstraints) {
    Gauge_SetConstraints(g._instance(), CheckPtr(value))
}

// Enabled
//
// 获取控件启用。
//
// Get the control enabled.
func (g *TGauge) Enabled() bool {
    return Gauge_GetEnabled(g._instance())
}

// SetEnabled
//
// 设置控件启用。
//
// Set the control enabled.
func (g *TGauge) SetEnabled(value bool) {
    Gauge_SetEnabled(g._instance(), value)
}

func (g *TGauge) ForeColor() TColor {
    return Gauge_GetForeColor(g._instance())
}

func (g *TGauge) SetForeColor(value TColor) {
    Gauge_SetForeColor(g._instance(), value)
}

// Font
//
// 获取字体。
//
// Get Font.
func (g *TGauge) Font() *TFont {
    return AsFont(Gauge_GetFont(g._instance()))
}

// SetFont
//
// 设置字体。
//
// Set Font.
func (g *TGauge) SetFont(value *TFont) {
    Gauge_SetFont(g._instance(), CheckPtr(value))
}

func (g *TGauge) Kind() TGaugeKind {
    return Gauge_GetKind(g._instance())
}

func (g *TGauge) SetKind(value TGaugeKind) {
    Gauge_SetKind(g._instance(), value)
}

func (g *TGauge) MinValue() int32 {
    return Gauge_GetMinValue(g._instance())
}

func (g *TGauge) SetMinValue(value int32) {
    Gauge_SetMinValue(g._instance(), value)
}

func (g *TGauge) MaxValue() int32 {
    return Gauge_GetMaxValue(g._instance())
}

func (g *TGauge) SetMaxValue(value int32) {
    Gauge_SetMaxValue(g._instance(), value)
}

// ParentColor
//
// 获取使用父容器颜色。
//
// Get parent color.
func (g *TGauge) ParentColor() bool {
    return Gauge_GetParentColor(g._instance())
}

// SetParentColor
//
// 设置使用父容器颜色。
//
// Set parent color.
func (g *TGauge) SetParentColor(value bool) {
    Gauge_SetParentColor(g._instance(), value)
}

// ParentFont
//
// 获取使用父容器字体。
//
// Get Parent container font.
func (g *TGauge) ParentFont() bool {
    return Gauge_GetParentFont(g._instance())
}

// SetParentFont
//
// 设置使用父容器字体。
//
// Set Parent container font.
func (g *TGauge) SetParentFont(value bool) {
    Gauge_SetParentFont(g._instance(), value)
}

// ParentShowHint
//
// 获取以父容器的ShowHint属性为准。
func (g *TGauge) ParentShowHint() bool {
    return Gauge_GetParentShowHint(g._instance())
}

// SetParentShowHint
//
// 设置以父容器的ShowHint属性为准。
func (g *TGauge) SetParentShowHint(value bool) {
    Gauge_SetParentShowHint(g._instance(), value)
}

// PopupMenu
//
// 获取右键菜单。
//
// Get Right click menu.
func (g *TGauge) PopupMenu() *TPopupMenu {
    return AsPopupMenu(Gauge_GetPopupMenu(g._instance()))
}

// SetPopupMenu
//
// 设置右键菜单。
//
// Set Right click menu.
func (g *TGauge) SetPopupMenu(value IComponent) {
    Gauge_SetPopupMenu(g._instance(), CheckPtr(value))
}

func (g *TGauge) Progress() int32 {
    return Gauge_GetProgress(g._instance())
}

func (g *TGauge) SetProgress(value int32) {
    Gauge_SetProgress(g._instance(), value)
}

// ShowHint
//
// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (g *TGauge) ShowHint() bool {
    return Gauge_GetShowHint(g._instance())
}

// SetShowHint
//
// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (g *TGauge) SetShowHint(value bool) {
    Gauge_SetShowHint(g._instance(), value)
}

func (g *TGauge) ShowText() bool {
    return Gauge_GetShowText(g._instance())
}

func (g *TGauge) SetShowText(value bool) {
    Gauge_SetShowText(g._instance(), value)
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (g *TGauge) Visible() bool {
    return Gauge_GetVisible(g._instance())
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (g *TGauge) SetVisible(value bool) {
    Gauge_SetVisible(g._instance(), value)
}

func (g *TGauge) Action() *TAction {
    return AsAction(Gauge_GetAction(g._instance()))
}

func (g *TGauge) SetAction(value IComponent) {
    Gauge_SetAction(g._instance(), CheckPtr(value))
}

func (g *TGauge) BiDiMode() TBiDiMode {
    return Gauge_GetBiDiMode(g._instance())
}

func (g *TGauge) SetBiDiMode(value TBiDiMode) {
    Gauge_SetBiDiMode(g._instance(), value)
}

func (g *TGauge) BoundsRect() TRect {
    return Gauge_GetBoundsRect(g._instance())
}

func (g *TGauge) SetBoundsRect(value TRect) {
    Gauge_SetBoundsRect(g._instance(), value)
}

// ClientHeight
//
// 获取客户区高度。
//
// Get client height.
func (g *TGauge) ClientHeight() int32 {
    return Gauge_GetClientHeight(g._instance())
}

// SetClientHeight
//
// 设置客户区高度。
//
// Set client height.
func (g *TGauge) SetClientHeight(value int32) {
    Gauge_SetClientHeight(g._instance(), value)
}

func (g *TGauge) ClientOrigin() TPoint {
    return Gauge_GetClientOrigin(g._instance())
}

// ClientRect
//
// 获取客户区矩形。
//
// Get client rectangle.
func (g *TGauge) ClientRect() TRect {
    return Gauge_GetClientRect(g._instance())
}

// ClientWidth
//
// 获取客户区宽度。
//
// Get client width.
func (g *TGauge) ClientWidth() int32 {
    return Gauge_GetClientWidth(g._instance())
}

// SetClientWidth
//
// 设置客户区宽度。
//
// Set client width.
func (g *TGauge) SetClientWidth(value int32) {
    Gauge_SetClientWidth(g._instance(), value)
}

// ControlState
//
// 获取控件状态。
//
// Get control state.
func (g *TGauge) ControlState() TControlState {
    return Gauge_GetControlState(g._instance())
}

// SetControlState
//
// 设置控件状态。
//
// Set control state.
func (g *TGauge) SetControlState(value TControlState) {
    Gauge_SetControlState(g._instance(), value)
}

// ControlStyle
//
// 获取控件样式。
//
// Get control style.
func (g *TGauge) ControlStyle() TControlStyle {
    return Gauge_GetControlStyle(g._instance())
}

// SetControlStyle
//
// 设置控件样式。
//
// Set control style.
func (g *TGauge) SetControlStyle(value TControlStyle) {
    Gauge_SetControlStyle(g._instance(), value)
}

func (g *TGauge) Floating() bool {
    return Gauge_GetFloating(g._instance())
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (g *TGauge) Parent() *TWinControl {
    return AsWinControl(Gauge_GetParent(g._instance()))
}

// SetParent
//
// 设置控件父容器。
//
// Set control parent container.
func (g *TGauge) SetParent(value IWinControl) {
    Gauge_SetParent(g._instance(), CheckPtr(value))
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (g *TGauge) Left() int32 {
    return Gauge_GetLeft(g._instance())
}

// SetLeft
//
// 设置左边位置。
//
// Set Left position.
func (g *TGauge) SetLeft(value int32) {
    Gauge_SetLeft(g._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (g *TGauge) Top() int32 {
    return Gauge_GetTop(g._instance())
}

// SetTop
//
// 设置顶边位置。
//
// Set Top position.
func (g *TGauge) SetTop(value int32) {
    Gauge_SetTop(g._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (g *TGauge) Width() int32 {
    return Gauge_GetWidth(g._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (g *TGauge) SetWidth(value int32) {
    Gauge_SetWidth(g._instance(), value)
}

// Height
//
// 获取高度。
//
// Get height.
func (g *TGauge) Height() int32 {
    return Gauge_GetHeight(g._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (g *TGauge) SetHeight(value int32) {
    Gauge_SetHeight(g._instance(), value)
}

// Cursor
//
// 获取控件光标。
//
// Get control cursor.
func (g *TGauge) Cursor() TCursor {
    return Gauge_GetCursor(g._instance())
}

// SetCursor
//
// 设置控件光标。
//
// Set control cursor.
func (g *TGauge) SetCursor(value TCursor) {
    Gauge_SetCursor(g._instance(), value)
}

// Hint
//
// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (g *TGauge) Hint() string {
    return Gauge_GetHint(g._instance())
}

// SetHint
//
// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (g *TGauge) SetHint(value string) {
    Gauge_SetHint(g._instance(), value)
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (g *TGauge) ComponentCount() int32 {
    return Gauge_GetComponentCount(g._instance())
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (g *TGauge) ComponentIndex() int32 {
    return Gauge_GetComponentIndex(g._instance())
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (g *TGauge) SetComponentIndex(value int32) {
    Gauge_SetComponentIndex(g._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (g *TGauge) Owner() *TComponent {
    return AsComponent(Gauge_GetOwner(g._instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (g *TGauge) Name() string {
    return Gauge_GetName(g._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (g *TGauge) SetName(value string) {
    Gauge_SetName(g._instance(), value)
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (g *TGauge) Tag() int {
    return Gauge_GetTag(g._instance())
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (g *TGauge) SetTag(value int) {
    Gauge_SetTag(g._instance(), value)
}

// AnchorSideLeft
//
// 获取左边锚点。
func (g *TGauge) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSideLeft(g._instance()))
}

// SetAnchorSideLeft
//
// 设置左边锚点。
func (g *TGauge) SetAnchorSideLeft(value *TAnchorSide) {
    Gauge_SetAnchorSideLeft(g._instance(), CheckPtr(value))
}

// AnchorSideTop
//
// 获取顶边锚点。
func (g *TGauge) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSideTop(g._instance()))
}

// SetAnchorSideTop
//
// 设置顶边锚点。
func (g *TGauge) SetAnchorSideTop(value *TAnchorSide) {
    Gauge_SetAnchorSideTop(g._instance(), CheckPtr(value))
}

// AnchorSideRight
//
// 获取右边锚点。
func (g *TGauge) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSideRight(g._instance()))
}

// SetAnchorSideRight
//
// 设置右边锚点。
func (g *TGauge) SetAnchorSideRight(value *TAnchorSide) {
    Gauge_SetAnchorSideRight(g._instance(), CheckPtr(value))
}

// AnchorSideBottom
//
// 获取底边锚点。
func (g *TGauge) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSideBottom(g._instance()))
}

// SetAnchorSideBottom
//
// 设置底边锚点。
func (g *TGauge) SetAnchorSideBottom(value *TAnchorSide) {
    Gauge_SetAnchorSideBottom(g._instance(), CheckPtr(value))
}

// BorderSpacing
//
// 获取边框间距。
func (g *TGauge) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(Gauge_GetBorderSpacing(g._instance()))
}

// SetBorderSpacing
//
// 设置边框间距。
func (g *TGauge) SetBorderSpacing(value *TControlBorderSpacing) {
    Gauge_SetBorderSpacing(g._instance(), CheckPtr(value))
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (g *TGauge) Components(AIndex int32) *TComponent {
    return AsComponent(Gauge_GetComponents(g._instance(), AIndex))
}

// AnchorSide
//
// 获取锚侧面。
func (g *TGauge) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSide(g._instance(), AKind))
}

