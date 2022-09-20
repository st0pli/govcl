
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

type TLabel struct {
    IControl
    instance unsafe.Pointer
}

// NewLabel
//
// 创建一个新的对象。
// 
// Create a new object.
func NewLabel(owner IComponent) *TLabel {
    l := new(TLabel)
    l.instance = unsafe.Pointer(Label_Create(CheckPtr(owner)))
    return l
}

// AsLabel
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsLabel(obj interface{}) *TLabel {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TLabel{instance: instance}
}

// Free 
//
// 释放对象。
// 
// Free object.
func (l *TLabel) Free() {
    if l.instance != nullptr {
        Label_Free(l._instance())
        l.instance  = nullptr
    }
}

func (l *TLabel) _instance() uintptr {
    return uintptr(l.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (l *TLabel) Instance() uintptr {
    return l._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (l *TLabel) UnsafeAddr() unsafe.Pointer {
    return l.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (l *TLabel) IsValid() bool {
    return l.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (l *TLabel) Is() TIs {
    return TIs(l._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (l *TLabel) As() TAs {
//    return TAs(l._instance())
//}

// TLabelClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TLabelClass() TClass {
    return Label_StaticClassType()
}

// BringToFront
//
// 将控件置于最前。
//
// Bring the control to the front.
func (l *TLabel) BringToFront() {
    Label_BringToFront(l._instance())
}

// ClientToScreen
//
// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (l *TLabel) ClientToScreen(Point TPoint) TPoint {
    return Label_ClientToScreen(l._instance(), Point)
}

// ClientToParent
//
// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (l *TLabel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Label_ClientToParent(l._instance(), Point , CheckPtr(AParent))
}

// Dragging
//
// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (l *TLabel) Dragging() bool {
    return Label_Dragging(l._instance())
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (l *TLabel) HasParent() bool {
    return Label_HasParent(l._instance())
}

// Hide
//
// 隐藏控件。
//
// Hidden control.
func (l *TLabel) Hide() {
    Label_Hide(l._instance())
}

// Invalidate
//
// 要求重绘。
//
// Redraw.
func (l *TLabel) Invalidate() {
    Label_Invalidate(l._instance())
}

// Perform
//
// 发送一个消息。
//
// Send a message.
func (l *TLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Label_Perform(l._instance(), Msg , WParam , LParam)
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (l *TLabel) Refresh() {
    Label_Refresh(l._instance())
}

// Repaint
//
// 重绘。
//
// Repaint.
func (l *TLabel) Repaint() {
    Label_Repaint(l._instance())
}

// ScreenToClient
//
// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (l *TLabel) ScreenToClient(Point TPoint) TPoint {
    return Label_ScreenToClient(l._instance(), Point)
}

// ParentToClient
//
// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (l *TLabel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Label_ParentToClient(l._instance(), Point , CheckPtr(AParent))
}

// SendToBack
//
// 控件至于最后面。
//
// The control is placed at the end.
func (l *TLabel) SendToBack() {
    Label_SendToBack(l._instance())
}

// SetBounds
//
// 设置组件边界。
//
// Set component boundaries.
func (l *TLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Label_SetBounds(l._instance(), ALeft , ATop , AWidth , AHeight)
}

// Show
//
// 显示控件。
//
// Show control.
func (l *TLabel) Show() {
    Label_Show(l._instance())
}

// Update
//
// 控件更新。
//
// Update.
func (l *TLabel) Update() {
    Label_Update(l._instance())
}

// GetTextBuf
//
// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (l *TLabel) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Label_GetTextBuf(l._instance(), Buffer , BufSize)
}

// GetTextLen
//
// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (l *TLabel) GetTextLen() int32 {
    return Label_GetTextLen(l._instance())
}

// SetTextBuf
//
// 设置控件字符，如果有。
//
// Set control characters, if any.
func (l *TLabel) SetTextBuf(Buffer string) {
    Label_SetTextBuf(l._instance(), Buffer)
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (l *TLabel) FindComponent(AName string) *TComponent {
    return AsComponent(Label_FindComponent(l._instance(), AName))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (l *TLabel) GetNamePath() string {
    return Label_GetNamePath(l._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (l *TLabel) Assign(Source IObject) {
    Label_Assign(l._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (l *TLabel) ClassType() TClass {
    return Label_ClassType(l._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (l *TLabel) ClassName() string {
    return Label_ClassName(l._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (l *TLabel) InstanceSize() int32 {
    return Label_InstanceSize(l._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (l *TLabel) InheritsFrom(AClass TClass) bool {
    return Label_InheritsFrom(l._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (l *TLabel) Equals(Obj IObject) bool {
    return Label_Equals(l._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (l *TLabel) GetHashCode() int32 {
    return Label_GetHashCode(l._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (l *TLabel) ToString() string {
    return Label_ToString(l._instance())
}

func (l *TLabel) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Label_AnchorToNeighbour(l._instance(), ASide , ASpace , CheckPtr(ASibling))
}

func (l *TLabel) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Label_AnchorParallel(l._instance(), ASide , ASpace , CheckPtr(ASibling))
}

// AnchorHorizontalCenterTo
//
// 置于指定控件的横向中心。
func (l *TLabel) AnchorHorizontalCenterTo(ASibling IControl) {
    Label_AnchorHorizontalCenterTo(l._instance(), CheckPtr(ASibling))
}

// AnchorVerticalCenterTo
//
// 置于指定控件的纵向中心。
func (l *TLabel) AnchorVerticalCenterTo(ASibling IControl) {
    Label_AnchorVerticalCenterTo(l._instance(), CheckPtr(ASibling))
}

func (l *TLabel) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    Label_AnchorSame(l._instance(), ASide , CheckPtr(ASibling))
}

func (l *TLabel) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    Label_AnchorAsAlign(l._instance(), ATheAlign , ASpace)
}

func (l *TLabel) AnchorClient(ASpace int32) {
    Label_AnchorClient(l._instance(), ASpace)
}

func (l *TLabel) ScaleDesignToForm(ASize int32) int32 {
    return Label_ScaleDesignToForm(l._instance(), ASize)
}

func (l *TLabel) ScaleFormToDesign(ASize int32) int32 {
    return Label_ScaleFormToDesign(l._instance(), ASize)
}

func (l *TLabel) Scale96ToForm(ASize int32) int32 {
    return Label_Scale96ToForm(l._instance(), ASize)
}

func (l *TLabel) ScaleFormTo96(ASize int32) int32 {
    return Label_ScaleFormTo96(l._instance(), ASize)
}

func (l *TLabel) Scale96ToFont(ASize int32) int32 {
    return Label_Scale96ToFont(l._instance(), ASize)
}

func (l *TLabel) ScaleFontTo96(ASize int32) int32 {
    return Label_ScaleFontTo96(l._instance(), ASize)
}

func (l *TLabel) ScaleScreenToFont(ASize int32) int32 {
    return Label_ScaleScreenToFont(l._instance(), ASize)
}

func (l *TLabel) ScaleFontToScreen(ASize int32) int32 {
    return Label_ScaleFontToScreen(l._instance(), ASize)
}

func (l *TLabel) Scale96ToScreen(ASize int32) int32 {
    return Label_Scale96ToScreen(l._instance(), ASize)
}

func (l *TLabel) ScaleScreenTo96(ASize int32) int32 {
    return Label_ScaleScreenTo96(l._instance(), ASize)
}

func (l *TLabel) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    Label_AutoAdjustLayout(l._instance(), AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (l *TLabel) FixDesignFontsPPI(ADesignTimePPI int32) {
    Label_FixDesignFontsPPI(l._instance(), ADesignTimePPI)
}

func (l *TLabel) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    Label_ScaleFontsPPI(l._instance(), AToPPI , AProportion)
}

func (l *TLabel) OptimalFill() bool {
    return Label_GetOptimalFill(l._instance())
}

func (l *TLabel) SetOptimalFill(value bool) {
    Label_SetOptimalFill(l._instance(), value)
}

// Align
//
// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (l *TLabel) Align() TAlign {
    return Label_GetAlign(l._instance())
}

// SetAlign
//
// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (l *TLabel) SetAlign(value TAlign) {
    Label_SetAlign(l._instance(), value)
}

// Alignment
//
// 获取文字对齐。
//
// Get Text alignment.
func (l *TLabel) Alignment() TAlignment {
    return Label_GetAlignment(l._instance())
}

// SetAlignment
//
// 设置文字对齐。
//
// Set Text alignment.
func (l *TLabel) SetAlignment(value TAlignment) {
    Label_SetAlignment(l._instance(), value)
}

// Anchors
//
// 获取四个角位置的锚点。
func (l *TLabel) Anchors() TAnchors {
    return Label_GetAnchors(l._instance())
}

// SetAnchors
//
// 设置四个角位置的锚点。
func (l *TLabel) SetAnchors(value TAnchors) {
    Label_SetAnchors(l._instance(), value)
}

// AutoSize
//
// 获取自动调整大小。
func (l *TLabel) AutoSize() bool {
    return Label_GetAutoSize(l._instance())
}

// SetAutoSize
//
// 设置自动调整大小。
func (l *TLabel) SetAutoSize(value bool) {
    Label_SetAutoSize(l._instance(), value)
}

func (l *TLabel) BiDiMode() TBiDiMode {
    return Label_GetBiDiMode(l._instance())
}

func (l *TLabel) SetBiDiMode(value TBiDiMode) {
    Label_SetBiDiMode(l._instance(), value)
}

// Caption
//
// 获取控件标题。
//
// Get the control title.
func (l *TLabel) Caption() string {
    return Label_GetCaption(l._instance())
}

// SetCaption
//
// 设置控件标题。
//
// Set the control title.
func (l *TLabel) SetCaption(value string) {
    Label_SetCaption(l._instance(), value)
}

// Color
//
// 获取颜色。
//
// Get color.
func (l *TLabel) Color() TColor {
    return Label_GetColor(l._instance())
}

// SetColor
//
// 设置颜色。
//
// Set color.
func (l *TLabel) SetColor(value TColor) {
    Label_SetColor(l._instance(), value)
}

// Constraints
//
// 获取约束控件大小。
func (l *TLabel) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Label_GetConstraints(l._instance()))
}

// SetConstraints
//
// 设置约束控件大小。
func (l *TLabel) SetConstraints(value *TSizeConstraints) {
    Label_SetConstraints(l._instance(), CheckPtr(value))
}

// DragCursor
//
// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (l *TLabel) DragCursor() TCursor {
    return Label_GetDragCursor(l._instance())
}

// SetDragCursor
//
// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (l *TLabel) SetDragCursor(value TCursor) {
    Label_SetDragCursor(l._instance(), value)
}

// DragKind
//
// 获取拖拽方式。
//
// Get Drag and drop.
func (l *TLabel) DragKind() TDragKind {
    return Label_GetDragKind(l._instance())
}

// SetDragKind
//
// 设置拖拽方式。
//
// Set Drag and drop.
func (l *TLabel) SetDragKind(value TDragKind) {
    Label_SetDragKind(l._instance(), value)
}

// DragMode
//
// 获取拖拽模式。
//
// Get Drag mode.
func (l *TLabel) DragMode() TDragMode {
    return Label_GetDragMode(l._instance())
}

// SetDragMode
//
// 设置拖拽模式。
//
// Set Drag mode.
func (l *TLabel) SetDragMode(value TDragMode) {
    Label_SetDragMode(l._instance(), value)
}

// Enabled
//
// 获取控件启用。
//
// Get the control enabled.
func (l *TLabel) Enabled() bool {
    return Label_GetEnabled(l._instance())
}

// SetEnabled
//
// 设置控件启用。
//
// Set the control enabled.
func (l *TLabel) SetEnabled(value bool) {
    Label_SetEnabled(l._instance(), value)
}

// FocusControl
//
// 获取当按下Alt+Key时跳转到指定的控件。
func (l *TLabel) FocusControl() *TWinControl {
    return AsWinControl(Label_GetFocusControl(l._instance()))
}

// SetFocusControl
//
// 设置当按下Alt+Key时跳转到指定的控件。
func (l *TLabel) SetFocusControl(value IWinControl) {
    Label_SetFocusControl(l._instance(), CheckPtr(value))
}

// Font
//
// 获取字体。
//
// Get Font.
func (l *TLabel) Font() *TFont {
    return AsFont(Label_GetFont(l._instance()))
}

// SetFont
//
// 设置字体。
//
// Set Font.
func (l *TLabel) SetFont(value *TFont) {
    Label_SetFont(l._instance(), CheckPtr(value))
}

// ParentColor
//
// 获取使用父容器颜色。
//
// Get parent color.
func (l *TLabel) ParentColor() bool {
    return Label_GetParentColor(l._instance())
}

// SetParentColor
//
// 设置使用父容器颜色。
//
// Set parent color.
func (l *TLabel) SetParentColor(value bool) {
    Label_SetParentColor(l._instance(), value)
}

// ParentFont
//
// 获取使用父容器字体。
//
// Get Parent container font.
func (l *TLabel) ParentFont() bool {
    return Label_GetParentFont(l._instance())
}

// SetParentFont
//
// 设置使用父容器字体。
//
// Set Parent container font.
func (l *TLabel) SetParentFont(value bool) {
    Label_SetParentFont(l._instance(), value)
}

// ParentShowHint
//
// 获取以父容器的ShowHint属性为准。
func (l *TLabel) ParentShowHint() bool {
    return Label_GetParentShowHint(l._instance())
}

// SetParentShowHint
//
// 设置以父容器的ShowHint属性为准。
func (l *TLabel) SetParentShowHint(value bool) {
    Label_SetParentShowHint(l._instance(), value)
}

// PopupMenu
//
// 获取右键菜单。
//
// Get Right click menu.
func (l *TLabel) PopupMenu() *TPopupMenu {
    return AsPopupMenu(Label_GetPopupMenu(l._instance()))
}

// SetPopupMenu
//
// 设置右键菜单。
//
// Set Right click menu.
func (l *TLabel) SetPopupMenu(value IComponent) {
    Label_SetPopupMenu(l._instance(), CheckPtr(value))
}

func (l *TLabel) ShowAccelChar() bool {
    return Label_GetShowAccelChar(l._instance())
}

func (l *TLabel) SetShowAccelChar(value bool) {
    Label_SetShowAccelChar(l._instance(), value)
}

// ShowHint
//
// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (l *TLabel) ShowHint() bool {
    return Label_GetShowHint(l._instance())
}

// SetShowHint
//
// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (l *TLabel) SetShowHint(value bool) {
    Label_SetShowHint(l._instance(), value)
}

// Transparent
//
// 获取透明。
//
// Get transparent.
func (l *TLabel) Transparent() bool {
    return Label_GetTransparent(l._instance())
}

// SetTransparent
//
// 设置透明。
//
// Set transparent.
func (l *TLabel) SetTransparent(value bool) {
    Label_SetTransparent(l._instance(), value)
}

func (l *TLabel) Layout() TTextLayout {
    return Label_GetLayout(l._instance())
}

func (l *TLabel) SetLayout(value TTextLayout) {
    Label_SetLayout(l._instance(), value)
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (l *TLabel) Visible() bool {
    return Label_GetVisible(l._instance())
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (l *TLabel) SetVisible(value bool) {
    Label_SetVisible(l._instance(), value)
}

// WordWrap
//
// 获取自动换行。
//
// Get Automatic line break.
func (l *TLabel) WordWrap() bool {
    return Label_GetWordWrap(l._instance())
}

// SetWordWrap
//
// 设置自动换行。
//
// Set Automatic line break.
func (l *TLabel) SetWordWrap(value bool) {
    Label_SetWordWrap(l._instance(), value)
}

// SetOnClick
//
// 设置控件单击事件。
//
// Set control click event.
func (l *TLabel) SetOnClick(fn TNotifyEvent) {
    Label_SetOnClick(l._instance(), fn)
}

// SetOnContextPopup
//
// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (l *TLabel) SetOnContextPopup(fn TContextPopupEvent) {
    Label_SetOnContextPopup(l._instance(), fn)
}

// SetOnDblClick
//
// 设置双击事件。
func (l *TLabel) SetOnDblClick(fn TNotifyEvent) {
    Label_SetOnDblClick(l._instance(), fn)
}

// SetOnDragDrop
//
// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (l *TLabel) SetOnDragDrop(fn TDragDropEvent) {
    Label_SetOnDragDrop(l._instance(), fn)
}

// SetOnDragOver
//
// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (l *TLabel) SetOnDragOver(fn TDragOverEvent) {
    Label_SetOnDragOver(l._instance(), fn)
}

// SetOnEndDrag
//
// 设置拖拽结束。
//
// Set End of drag.
func (l *TLabel) SetOnEndDrag(fn TEndDragEvent) {
    Label_SetOnEndDrag(l._instance(), fn)
}

// SetOnMouseDown
//
// 设置鼠标按下事件。
//
// Set Mouse down event.
func (l *TLabel) SetOnMouseDown(fn TMouseEvent) {
    Label_SetOnMouseDown(l._instance(), fn)
}

// SetOnMouseMove
//
// 设置鼠标移动事件。
func (l *TLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    Label_SetOnMouseMove(l._instance(), fn)
}

// SetOnMouseUp
//
// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (l *TLabel) SetOnMouseUp(fn TMouseEvent) {
    Label_SetOnMouseUp(l._instance(), fn)
}

// SetOnMouseEnter
//
// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (l *TLabel) SetOnMouseEnter(fn TNotifyEvent) {
    Label_SetOnMouseEnter(l._instance(), fn)
}

// SetOnMouseLeave
//
// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (l *TLabel) SetOnMouseLeave(fn TNotifyEvent) {
    Label_SetOnMouseLeave(l._instance(), fn)
}

// Canvas
//
// 获取画布。
func (l *TLabel) Canvas() *TCanvas {
    return AsCanvas(Label_GetCanvas(l._instance()))
}

func (l *TLabel) Action() *TAction {
    return AsAction(Label_GetAction(l._instance()))
}

func (l *TLabel) SetAction(value IComponent) {
    Label_SetAction(l._instance(), CheckPtr(value))
}

func (l *TLabel) BoundsRect() TRect {
    return Label_GetBoundsRect(l._instance())
}

func (l *TLabel) SetBoundsRect(value TRect) {
    Label_SetBoundsRect(l._instance(), value)
}

// ClientHeight
//
// 获取客户区高度。
//
// Get client height.
func (l *TLabel) ClientHeight() int32 {
    return Label_GetClientHeight(l._instance())
}

// SetClientHeight
//
// 设置客户区高度。
//
// Set client height.
func (l *TLabel) SetClientHeight(value int32) {
    Label_SetClientHeight(l._instance(), value)
}

func (l *TLabel) ClientOrigin() TPoint {
    return Label_GetClientOrigin(l._instance())
}

// ClientRect
//
// 获取客户区矩形。
//
// Get client rectangle.
func (l *TLabel) ClientRect() TRect {
    return Label_GetClientRect(l._instance())
}

// ClientWidth
//
// 获取客户区宽度。
//
// Get client width.
func (l *TLabel) ClientWidth() int32 {
    return Label_GetClientWidth(l._instance())
}

// SetClientWidth
//
// 设置客户区宽度。
//
// Set client width.
func (l *TLabel) SetClientWidth(value int32) {
    Label_SetClientWidth(l._instance(), value)
}

// ControlState
//
// 获取控件状态。
//
// Get control state.
func (l *TLabel) ControlState() TControlState {
    return Label_GetControlState(l._instance())
}

// SetControlState
//
// 设置控件状态。
//
// Set control state.
func (l *TLabel) SetControlState(value TControlState) {
    Label_SetControlState(l._instance(), value)
}

// ControlStyle
//
// 获取控件样式。
//
// Get control style.
func (l *TLabel) ControlStyle() TControlStyle {
    return Label_GetControlStyle(l._instance())
}

// SetControlStyle
//
// 设置控件样式。
//
// Set control style.
func (l *TLabel) SetControlStyle(value TControlStyle) {
    Label_SetControlStyle(l._instance(), value)
}

func (l *TLabel) Floating() bool {
    return Label_GetFloating(l._instance())
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (l *TLabel) Parent() *TWinControl {
    return AsWinControl(Label_GetParent(l._instance()))
}

// SetParent
//
// 设置控件父容器。
//
// Set control parent container.
func (l *TLabel) SetParent(value IWinControl) {
    Label_SetParent(l._instance(), CheckPtr(value))
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (l *TLabel) Left() int32 {
    return Label_GetLeft(l._instance())
}

// SetLeft
//
// 设置左边位置。
//
// Set Left position.
func (l *TLabel) SetLeft(value int32) {
    Label_SetLeft(l._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (l *TLabel) Top() int32 {
    return Label_GetTop(l._instance())
}

// SetTop
//
// 设置顶边位置。
//
// Set Top position.
func (l *TLabel) SetTop(value int32) {
    Label_SetTop(l._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (l *TLabel) Width() int32 {
    return Label_GetWidth(l._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (l *TLabel) SetWidth(value int32) {
    Label_SetWidth(l._instance(), value)
}

// Height
//
// 获取高度。
//
// Get height.
func (l *TLabel) Height() int32 {
    return Label_GetHeight(l._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (l *TLabel) SetHeight(value int32) {
    Label_SetHeight(l._instance(), value)
}

// Cursor
//
// 获取控件光标。
//
// Get control cursor.
func (l *TLabel) Cursor() TCursor {
    return Label_GetCursor(l._instance())
}

// SetCursor
//
// 设置控件光标。
//
// Set control cursor.
func (l *TLabel) SetCursor(value TCursor) {
    Label_SetCursor(l._instance(), value)
}

// Hint
//
// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (l *TLabel) Hint() string {
    return Label_GetHint(l._instance())
}

// SetHint
//
// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (l *TLabel) SetHint(value string) {
    Label_SetHint(l._instance(), value)
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (l *TLabel) ComponentCount() int32 {
    return Label_GetComponentCount(l._instance())
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (l *TLabel) ComponentIndex() int32 {
    return Label_GetComponentIndex(l._instance())
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (l *TLabel) SetComponentIndex(value int32) {
    Label_SetComponentIndex(l._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (l *TLabel) Owner() *TComponent {
    return AsComponent(Label_GetOwner(l._instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (l *TLabel) Name() string {
    return Label_GetName(l._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (l *TLabel) SetName(value string) {
    Label_SetName(l._instance(), value)
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (l *TLabel) Tag() int {
    return Label_GetTag(l._instance())
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (l *TLabel) SetTag(value int) {
    Label_SetTag(l._instance(), value)
}

// AnchorSideLeft
//
// 获取左边锚点。
func (l *TLabel) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(Label_GetAnchorSideLeft(l._instance()))
}

// SetAnchorSideLeft
//
// 设置左边锚点。
func (l *TLabel) SetAnchorSideLeft(value *TAnchorSide) {
    Label_SetAnchorSideLeft(l._instance(), CheckPtr(value))
}

// AnchorSideTop
//
// 获取顶边锚点。
func (l *TLabel) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(Label_GetAnchorSideTop(l._instance()))
}

// SetAnchorSideTop
//
// 设置顶边锚点。
func (l *TLabel) SetAnchorSideTop(value *TAnchorSide) {
    Label_SetAnchorSideTop(l._instance(), CheckPtr(value))
}

// AnchorSideRight
//
// 获取右边锚点。
func (l *TLabel) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(Label_GetAnchorSideRight(l._instance()))
}

// SetAnchorSideRight
//
// 设置右边锚点。
func (l *TLabel) SetAnchorSideRight(value *TAnchorSide) {
    Label_SetAnchorSideRight(l._instance(), CheckPtr(value))
}

// AnchorSideBottom
//
// 获取底边锚点。
func (l *TLabel) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(Label_GetAnchorSideBottom(l._instance()))
}

// SetAnchorSideBottom
//
// 设置底边锚点。
func (l *TLabel) SetAnchorSideBottom(value *TAnchorSide) {
    Label_SetAnchorSideBottom(l._instance(), CheckPtr(value))
}

// BorderSpacing
//
// 获取边框间距。
func (l *TLabel) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(Label_GetBorderSpacing(l._instance()))
}

// SetBorderSpacing
//
// 设置边框间距。
func (l *TLabel) SetBorderSpacing(value *TControlBorderSpacing) {
    Label_SetBorderSpacing(l._instance(), CheckPtr(value))
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (l *TLabel) Components(AIndex int32) *TComponent {
    return AsComponent(Label_GetComponents(l._instance(), AIndex))
}

// AnchorSide
//
// 获取锚侧面。
func (l *TLabel) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(Label_GetAnchorSide(l._instance(), AKind))
}

