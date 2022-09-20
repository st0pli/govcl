
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

type TPanel struct {
    IWinControl
    instance unsafe.Pointer
}

// NewPanel
//
// 创建一个新的对象。
// 
// Create a new object.
func NewPanel(owner IComponent) *TPanel {
    p := new(TPanel)
    p.instance = unsafe.Pointer(Panel_Create(CheckPtr(owner)))
    return p
}

// AsPanel
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsPanel(obj interface{}) *TPanel {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TPanel{instance: instance}
}

// Free 
//
// 释放对象。
// 
// Free object.
func (p *TPanel) Free() {
    if p.instance != nullptr {
        Panel_Free(p._instance())
        p.instance  = nullptr
    }
}

func (p *TPanel) _instance() uintptr {
    return uintptr(p.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (p *TPanel) Instance() uintptr {
    return p._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (p *TPanel) UnsafeAddr() unsafe.Pointer {
    return p.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (p *TPanel) IsValid() bool {
    return p.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (p *TPanel) Is() TIs {
    return TIs(p._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (p *TPanel) As() TAs {
//    return TAs(p._instance())
//}

// TPanelClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TPanelClass() TClass {
    return Panel_StaticClassType()
}

// CanFocus
//
// 是否可以获得焦点。
func (p *TPanel) CanFocus() bool {
    return Panel_CanFocus(p._instance())
}

// ContainsControl
//
// 返回是否包含指定控件。
//
// it's contain a specified control.
func (p *TPanel) ContainsControl(Control IControl) bool {
    return Panel_ContainsControl(p._instance(), CheckPtr(Control))
}

// ControlAtPos
//
// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (p *TPanel) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(Panel_ControlAtPos(p._instance(), Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
//
// 禁用控件的对齐。
//
// Disable control alignment.
func (p *TPanel) DisableAlign() {
    Panel_DisableAlign(p._instance())
}

// EnableAlign
//
// 启用控件对齐。
//
// Enabled control alignment.
func (p *TPanel) EnableAlign() {
    Panel_EnableAlign(p._instance())
}

// FindChildControl
//
// 查找子控件。
//
// Find sub controls.
func (p *TPanel) FindChildControl(ControlName string) *TControl {
    return AsControl(Panel_FindChildControl(p._instance(), ControlName))
}

func (p *TPanel) FlipChildren(AllLevels bool) {
    Panel_FlipChildren(p._instance(), AllLevels)
}

// Focused
//
// 返回是否获取焦点。
//
// Return to get focus.
func (p *TPanel) Focused() bool {
    return Panel_Focused(p._instance())
}

// HandleAllocated
//
// 句柄是否已经分配。
//
// Is the handle already allocated.
func (p *TPanel) HandleAllocated() bool {
    return Panel_HandleAllocated(p._instance())
}

// InsertControl
//
// 插入一个控件。
//
// Insert a control.
func (p *TPanel) InsertControl(AControl IControl) {
    Panel_InsertControl(p._instance(), CheckPtr(AControl))
}

// Invalidate
//
// 要求重绘。
//
// Redraw.
func (p *TPanel) Invalidate() {
    Panel_Invalidate(p._instance())
}

// PaintTo
//
// 绘画至指定DC。
//
// Painting to the specified DC.
func (p *TPanel) PaintTo(DC HDC, X int32, Y int32) {
    Panel_PaintTo(p._instance(), DC , X , Y)
}

// RemoveControl
//
// 移除一个控件。
//
// Remove a control.
func (p *TPanel) RemoveControl(AControl IControl) {
    Panel_RemoveControl(p._instance(), CheckPtr(AControl))
}

// Realign
//
// 重新对齐。
//
// Realign.
func (p *TPanel) Realign() {
    Panel_Realign(p._instance())
}

// Repaint
//
// 重绘。
//
// Repaint.
func (p *TPanel) Repaint() {
    Panel_Repaint(p._instance())
}

// ScaleBy
//
// 按比例缩放。
//
// Scale by.
func (p *TPanel) ScaleBy(M int32, D int32) {
    Panel_ScaleBy(p._instance(), M , D)
}

// ScrollBy
//
// 滚动至指定位置。
//
// Scroll by.
func (p *TPanel) ScrollBy(DeltaX int32, DeltaY int32) {
    Panel_ScrollBy(p._instance(), DeltaX , DeltaY)
}

// SetBounds
//
// 设置组件边界。
//
// Set component boundaries.
func (p *TPanel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Panel_SetBounds(p._instance(), ALeft , ATop , AWidth , AHeight)
}

// SetFocus
//
// 设置控件焦点。
//
// Set control focus.
func (p *TPanel) SetFocus() {
    Panel_SetFocus(p._instance())
}

// Update
//
// 控件更新。
//
// Update.
func (p *TPanel) Update() {
    Panel_Update(p._instance())
}

// BringToFront
//
// 将控件置于最前。
//
// Bring the control to the front.
func (p *TPanel) BringToFront() {
    Panel_BringToFront(p._instance())
}

// ClientToScreen
//
// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (p *TPanel) ClientToScreen(Point TPoint) TPoint {
    return Panel_ClientToScreen(p._instance(), Point)
}

// ClientToParent
//
// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (p *TPanel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Panel_ClientToParent(p._instance(), Point , CheckPtr(AParent))
}

// Dragging
//
// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (p *TPanel) Dragging() bool {
    return Panel_Dragging(p._instance())
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (p *TPanel) HasParent() bool {
    return Panel_HasParent(p._instance())
}

// Hide
//
// 隐藏控件。
//
// Hidden control.
func (p *TPanel) Hide() {
    Panel_Hide(p._instance())
}

// Perform
//
// 发送一个消息。
//
// Send a message.
func (p *TPanel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Panel_Perform(p._instance(), Msg , WParam , LParam)
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (p *TPanel) Refresh() {
    Panel_Refresh(p._instance())
}

// ScreenToClient
//
// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (p *TPanel) ScreenToClient(Point TPoint) TPoint {
    return Panel_ScreenToClient(p._instance(), Point)
}

// ParentToClient
//
// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (p *TPanel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Panel_ParentToClient(p._instance(), Point , CheckPtr(AParent))
}

// SendToBack
//
// 控件至于最后面。
//
// The control is placed at the end.
func (p *TPanel) SendToBack() {
    Panel_SendToBack(p._instance())
}

// Show
//
// 显示控件。
//
// Show control.
func (p *TPanel) Show() {
    Panel_Show(p._instance())
}

// GetTextBuf
//
// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (p *TPanel) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Panel_GetTextBuf(p._instance(), Buffer , BufSize)
}

// GetTextLen
//
// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (p *TPanel) GetTextLen() int32 {
    return Panel_GetTextLen(p._instance())
}

// SetTextBuf
//
// 设置控件字符，如果有。
//
// Set control characters, if any.
func (p *TPanel) SetTextBuf(Buffer string) {
    Panel_SetTextBuf(p._instance(), Buffer)
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (p *TPanel) FindComponent(AName string) *TComponent {
    return AsComponent(Panel_FindComponent(p._instance(), AName))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (p *TPanel) GetNamePath() string {
    return Panel_GetNamePath(p._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (p *TPanel) Assign(Source IObject) {
    Panel_Assign(p._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (p *TPanel) ClassType() TClass {
    return Panel_ClassType(p._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (p *TPanel) ClassName() string {
    return Panel_ClassName(p._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (p *TPanel) InstanceSize() int32 {
    return Panel_InstanceSize(p._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (p *TPanel) InheritsFrom(AClass TClass) bool {
    return Panel_InheritsFrom(p._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (p *TPanel) Equals(Obj IObject) bool {
    return Panel_Equals(p._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (p *TPanel) GetHashCode() int32 {
    return Panel_GetHashCode(p._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (p *TPanel) ToString() string {
    return Panel_ToString(p._instance())
}

func (p *TPanel) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Panel_AnchorToNeighbour(p._instance(), ASide , ASpace , CheckPtr(ASibling))
}

func (p *TPanel) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Panel_AnchorParallel(p._instance(), ASide , ASpace , CheckPtr(ASibling))
}

// AnchorHorizontalCenterTo
//
// 置于指定控件的横向中心。
func (p *TPanel) AnchorHorizontalCenterTo(ASibling IControl) {
    Panel_AnchorHorizontalCenterTo(p._instance(), CheckPtr(ASibling))
}

// AnchorVerticalCenterTo
//
// 置于指定控件的纵向中心。
func (p *TPanel) AnchorVerticalCenterTo(ASibling IControl) {
    Panel_AnchorVerticalCenterTo(p._instance(), CheckPtr(ASibling))
}

func (p *TPanel) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    Panel_AnchorSame(p._instance(), ASide , CheckPtr(ASibling))
}

func (p *TPanel) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    Panel_AnchorAsAlign(p._instance(), ATheAlign , ASpace)
}

func (p *TPanel) AnchorClient(ASpace int32) {
    Panel_AnchorClient(p._instance(), ASpace)
}

func (p *TPanel) ScaleDesignToForm(ASize int32) int32 {
    return Panel_ScaleDesignToForm(p._instance(), ASize)
}

func (p *TPanel) ScaleFormToDesign(ASize int32) int32 {
    return Panel_ScaleFormToDesign(p._instance(), ASize)
}

func (p *TPanel) Scale96ToForm(ASize int32) int32 {
    return Panel_Scale96ToForm(p._instance(), ASize)
}

func (p *TPanel) ScaleFormTo96(ASize int32) int32 {
    return Panel_ScaleFormTo96(p._instance(), ASize)
}

func (p *TPanel) Scale96ToFont(ASize int32) int32 {
    return Panel_Scale96ToFont(p._instance(), ASize)
}

func (p *TPanel) ScaleFontTo96(ASize int32) int32 {
    return Panel_ScaleFontTo96(p._instance(), ASize)
}

func (p *TPanel) ScaleScreenToFont(ASize int32) int32 {
    return Panel_ScaleScreenToFont(p._instance(), ASize)
}

func (p *TPanel) ScaleFontToScreen(ASize int32) int32 {
    return Panel_ScaleFontToScreen(p._instance(), ASize)
}

func (p *TPanel) Scale96ToScreen(ASize int32) int32 {
    return Panel_Scale96ToScreen(p._instance(), ASize)
}

func (p *TPanel) ScaleScreenTo96(ASize int32) int32 {
    return Panel_ScaleScreenTo96(p._instance(), ASize)
}

func (p *TPanel) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    Panel_AutoAdjustLayout(p._instance(), AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (p *TPanel) FixDesignFontsPPI(ADesignTimePPI int32) {
    Panel_FixDesignFontsPPI(p._instance(), ADesignTimePPI)
}

func (p *TPanel) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    Panel_ScaleFontsPPI(p._instance(), AToPPI , AProportion)
}

// Canvas
//
// 获取画布。
func (p *TPanel) Canvas() *TCanvas {
    return AsCanvas(Panel_GetCanvas(p._instance()))
}

// SetCanvas
//
// 设置画布。
func (p *TPanel) SetCanvas(value *TCanvas) {
    Panel_SetCanvas(p._instance(), CheckPtr(value))
}

// SetOnPaint
//
// 设置绘画事件。
func (p *TPanel) SetOnPaint(fn TNotifyEvent) {
    Panel_SetOnPaint(p._instance(), fn)
}

// Align
//
// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (p *TPanel) Align() TAlign {
    return Panel_GetAlign(p._instance())
}

// SetAlign
//
// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (p *TPanel) SetAlign(value TAlign) {
    Panel_SetAlign(p._instance(), value)
}

// Alignment
//
// 获取文字对齐。
//
// Get Text alignment.
func (p *TPanel) Alignment() TAlignment {
    return Panel_GetAlignment(p._instance())
}

// SetAlignment
//
// 设置文字对齐。
//
// Set Text alignment.
func (p *TPanel) SetAlignment(value TAlignment) {
    Panel_SetAlignment(p._instance(), value)
}

// Anchors
//
// 获取四个角位置的锚点。
func (p *TPanel) Anchors() TAnchors {
    return Panel_GetAnchors(p._instance())
}

// SetAnchors
//
// 设置四个角位置的锚点。
func (p *TPanel) SetAnchors(value TAnchors) {
    Panel_SetAnchors(p._instance(), value)
}

// AutoSize
//
// 获取自动调整大小。
func (p *TPanel) AutoSize() bool {
    return Panel_GetAutoSize(p._instance())
}

// SetAutoSize
//
// 设置自动调整大小。
func (p *TPanel) SetAutoSize(value bool) {
    Panel_SetAutoSize(p._instance(), value)
}

func (p *TPanel) BevelInner() TBevelCut {
    return Panel_GetBevelInner(p._instance())
}

func (p *TPanel) SetBevelInner(value TBevelCut) {
    Panel_SetBevelInner(p._instance(), value)
}

func (p *TPanel) BevelOuter() TBevelCut {
    return Panel_GetBevelOuter(p._instance())
}

func (p *TPanel) SetBevelOuter(value TBevelCut) {
    Panel_SetBevelOuter(p._instance(), value)
}

func (p *TPanel) BiDiMode() TBiDiMode {
    return Panel_GetBiDiMode(p._instance())
}

func (p *TPanel) SetBiDiMode(value TBiDiMode) {
    Panel_SetBiDiMode(p._instance(), value)
}

// BorderWidth
//
// 获取边框的宽度。
func (p *TPanel) BorderWidth() int32 {
    return Panel_GetBorderWidth(p._instance())
}

// SetBorderWidth
//
// 设置边框的宽度。
func (p *TPanel) SetBorderWidth(value int32) {
    Panel_SetBorderWidth(p._instance(), value)
}

// BorderStyle
//
// 获取窗口边框样式。比如：无边框，单一边框等。
func (p *TPanel) BorderStyle() TBorderStyle {
    return Panel_GetBorderStyle(p._instance())
}

// SetBorderStyle
//
// 设置窗口边框样式。比如：无边框，单一边框等。
func (p *TPanel) SetBorderStyle(value TBorderStyle) {
    Panel_SetBorderStyle(p._instance(), value)
}

// Caption
//
// 获取控件标题。
//
// Get the control title.
func (p *TPanel) Caption() string {
    return Panel_GetCaption(p._instance())
}

// SetCaption
//
// 设置控件标题。
//
// Set the control title.
func (p *TPanel) SetCaption(value string) {
    Panel_SetCaption(p._instance(), value)
}

// Color
//
// 获取颜色。
//
// Get color.
func (p *TPanel) Color() TColor {
    return Panel_GetColor(p._instance())
}

// SetColor
//
// 设置颜色。
//
// Set color.
func (p *TPanel) SetColor(value TColor) {
    Panel_SetColor(p._instance(), value)
}

// Constraints
//
// 获取约束控件大小。
func (p *TPanel) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Panel_GetConstraints(p._instance()))
}

// SetConstraints
//
// 设置约束控件大小。
func (p *TPanel) SetConstraints(value *TSizeConstraints) {
    Panel_SetConstraints(p._instance(), CheckPtr(value))
}

// UseDockManager
//
// 获取使用停靠管理。
func (p *TPanel) UseDockManager() bool {
    return Panel_GetUseDockManager(p._instance())
}

// SetUseDockManager
//
// 设置使用停靠管理。
func (p *TPanel) SetUseDockManager(value bool) {
    Panel_SetUseDockManager(p._instance(), value)
}

// DockSite
//
// 获取停靠站点。
//
// Get Docking site.
func (p *TPanel) DockSite() bool {
    return Panel_GetDockSite(p._instance())
}

// SetDockSite
//
// 设置停靠站点。
//
// Set Docking site.
func (p *TPanel) SetDockSite(value bool) {
    Panel_SetDockSite(p._instance(), value)
}

// DoubleBuffered
//
// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (p *TPanel) DoubleBuffered() bool {
    return Panel_GetDoubleBuffered(p._instance())
}

// SetDoubleBuffered
//
// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (p *TPanel) SetDoubleBuffered(value bool) {
    Panel_SetDoubleBuffered(p._instance(), value)
}

// DragCursor
//
// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (p *TPanel) DragCursor() TCursor {
    return Panel_GetDragCursor(p._instance())
}

// SetDragCursor
//
// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (p *TPanel) SetDragCursor(value TCursor) {
    Panel_SetDragCursor(p._instance(), value)
}

// DragKind
//
// 获取拖拽方式。
//
// Get Drag and drop.
func (p *TPanel) DragKind() TDragKind {
    return Panel_GetDragKind(p._instance())
}

// SetDragKind
//
// 设置拖拽方式。
//
// Set Drag and drop.
func (p *TPanel) SetDragKind(value TDragKind) {
    Panel_SetDragKind(p._instance(), value)
}

// DragMode
//
// 获取拖拽模式。
//
// Get Drag mode.
func (p *TPanel) DragMode() TDragMode {
    return Panel_GetDragMode(p._instance())
}

// SetDragMode
//
// 设置拖拽模式。
//
// Set Drag mode.
func (p *TPanel) SetDragMode(value TDragMode) {
    Panel_SetDragMode(p._instance(), value)
}

// Enabled
//
// 获取控件启用。
//
// Get the control enabled.
func (p *TPanel) Enabled() bool {
    return Panel_GetEnabled(p._instance())
}

// SetEnabled
//
// 设置控件启用。
//
// Set the control enabled.
func (p *TPanel) SetEnabled(value bool) {
    Panel_SetEnabled(p._instance(), value)
}

func (p *TPanel) FullRepaint() bool {
    return Panel_GetFullRepaint(p._instance())
}

func (p *TPanel) SetFullRepaint(value bool) {
    Panel_SetFullRepaint(p._instance(), value)
}

// Font
//
// 获取字体。
//
// Get Font.
func (p *TPanel) Font() *TFont {
    return AsFont(Panel_GetFont(p._instance()))
}

// SetFont
//
// 设置字体。
//
// Set Font.
func (p *TPanel) SetFont(value *TFont) {
    Panel_SetFont(p._instance(), CheckPtr(value))
}

func (p *TPanel) ParentBackground() bool {
    return Panel_GetParentBackground(p._instance())
}

func (p *TPanel) SetParentBackground(value bool) {
    Panel_SetParentBackground(p._instance(), value)
}

// ParentColor
//
// 获取使用父容器颜色。
//
// Get parent color.
func (p *TPanel) ParentColor() bool {
    return Panel_GetParentColor(p._instance())
}

// SetParentColor
//
// 设置使用父容器颜色。
//
// Set parent color.
func (p *TPanel) SetParentColor(value bool) {
    Panel_SetParentColor(p._instance(), value)
}

// ParentDoubleBuffered
//
// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (p *TPanel) ParentDoubleBuffered() bool {
    return Panel_GetParentDoubleBuffered(p._instance())
}

// SetParentDoubleBuffered
//
// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (p *TPanel) SetParentDoubleBuffered(value bool) {
    Panel_SetParentDoubleBuffered(p._instance(), value)
}

// ParentFont
//
// 获取使用父容器字体。
//
// Get Parent container font.
func (p *TPanel) ParentFont() bool {
    return Panel_GetParentFont(p._instance())
}

// SetParentFont
//
// 设置使用父容器字体。
//
// Set Parent container font.
func (p *TPanel) SetParentFont(value bool) {
    Panel_SetParentFont(p._instance(), value)
}

// ParentShowHint
//
// 获取以父容器的ShowHint属性为准。
func (p *TPanel) ParentShowHint() bool {
    return Panel_GetParentShowHint(p._instance())
}

// SetParentShowHint
//
// 设置以父容器的ShowHint属性为准。
func (p *TPanel) SetParentShowHint(value bool) {
    Panel_SetParentShowHint(p._instance(), value)
}

// PopupMenu
//
// 获取右键菜单。
//
// Get Right click menu.
func (p *TPanel) PopupMenu() *TPopupMenu {
    return AsPopupMenu(Panel_GetPopupMenu(p._instance()))
}

// SetPopupMenu
//
// 设置右键菜单。
//
// Set Right click menu.
func (p *TPanel) SetPopupMenu(value IComponent) {
    Panel_SetPopupMenu(p._instance(), CheckPtr(value))
}

// ShowHint
//
// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (p *TPanel) ShowHint() bool {
    return Panel_GetShowHint(p._instance())
}

// SetShowHint
//
// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (p *TPanel) SetShowHint(value bool) {
    Panel_SetShowHint(p._instance(), value)
}

// TabOrder
//
// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (p *TPanel) TabOrder() TTabOrder {
    return Panel_GetTabOrder(p._instance())
}

// SetTabOrder
//
// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (p *TPanel) SetTabOrder(value TTabOrder) {
    Panel_SetTabOrder(p._instance(), value)
}

// TabStop
//
// 获取Tab可停留。
//
// Get Tab can stay.
func (p *TPanel) TabStop() bool {
    return Panel_GetTabStop(p._instance())
}

// SetTabStop
//
// 设置Tab可停留。
//
// Set Tab can stay.
func (p *TPanel) SetTabStop(value bool) {
    Panel_SetTabStop(p._instance(), value)
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (p *TPanel) Visible() bool {
    return Panel_GetVisible(p._instance())
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (p *TPanel) SetVisible(value bool) {
    Panel_SetVisible(p._instance(), value)
}

// SetOnAlignPosition
//
// 设置对齐位置事件，当Align为alCustom时Parent会收到这个消息。
func (p *TPanel) SetOnAlignPosition(fn TAlignPositionEvent) {
    Panel_SetOnAlignPosition(p._instance(), fn)
}

// SetOnClick
//
// 设置控件单击事件。
//
// Set control click event.
func (p *TPanel) SetOnClick(fn TNotifyEvent) {
    Panel_SetOnClick(p._instance(), fn)
}

// SetOnContextPopup
//
// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (p *TPanel) SetOnContextPopup(fn TContextPopupEvent) {
    Panel_SetOnContextPopup(p._instance(), fn)
}

func (p *TPanel) SetOnDockDrop(fn TDockDropEvent) {
    Panel_SetOnDockDrop(p._instance(), fn)
}

// SetOnDblClick
//
// 设置双击事件。
func (p *TPanel) SetOnDblClick(fn TNotifyEvent) {
    Panel_SetOnDblClick(p._instance(), fn)
}

// SetOnDragDrop
//
// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (p *TPanel) SetOnDragDrop(fn TDragDropEvent) {
    Panel_SetOnDragDrop(p._instance(), fn)
}

// SetOnDragOver
//
// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (p *TPanel) SetOnDragOver(fn TDragOverEvent) {
    Panel_SetOnDragOver(p._instance(), fn)
}

// SetOnEndDock
//
// 设置停靠结束事件。
//
// Set Dock end event.
func (p *TPanel) SetOnEndDock(fn TEndDragEvent) {
    Panel_SetOnEndDock(p._instance(), fn)
}

// SetOnEndDrag
//
// 设置拖拽结束。
//
// Set End of drag.
func (p *TPanel) SetOnEndDrag(fn TEndDragEvent) {
    Panel_SetOnEndDrag(p._instance(), fn)
}

// SetOnEnter
//
// 设置焦点进入。
//
// Set Focus entry.
func (p *TPanel) SetOnEnter(fn TNotifyEvent) {
    Panel_SetOnEnter(p._instance(), fn)
}

// SetOnExit
//
// 设置焦点退出。
//
// Set Focus exit.
func (p *TPanel) SetOnExit(fn TNotifyEvent) {
    Panel_SetOnExit(p._instance(), fn)
}

func (p *TPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    Panel_SetOnGetSiteInfo(p._instance(), fn)
}

// SetOnMouseDown
//
// 设置鼠标按下事件。
//
// Set Mouse down event.
func (p *TPanel) SetOnMouseDown(fn TMouseEvent) {
    Panel_SetOnMouseDown(p._instance(), fn)
}

// SetOnMouseEnter
//
// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (p *TPanel) SetOnMouseEnter(fn TNotifyEvent) {
    Panel_SetOnMouseEnter(p._instance(), fn)
}

// SetOnMouseLeave
//
// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (p *TPanel) SetOnMouseLeave(fn TNotifyEvent) {
    Panel_SetOnMouseLeave(p._instance(), fn)
}

// SetOnMouseMove
//
// 设置鼠标移动事件。
func (p *TPanel) SetOnMouseMove(fn TMouseMoveEvent) {
    Panel_SetOnMouseMove(p._instance(), fn)
}

// SetOnMouseUp
//
// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (p *TPanel) SetOnMouseUp(fn TMouseEvent) {
    Panel_SetOnMouseUp(p._instance(), fn)
}

// SetOnResize
//
// 设置大小被改变事件。
func (p *TPanel) SetOnResize(fn TNotifyEvent) {
    Panel_SetOnResize(p._instance(), fn)
}

// SetOnStartDock
//
// 设置启动停靠。
func (p *TPanel) SetOnStartDock(fn TStartDockEvent) {
    Panel_SetOnStartDock(p._instance(), fn)
}

func (p *TPanel) SetOnUnDock(fn TUnDockEvent) {
    Panel_SetOnUnDock(p._instance(), fn)
}

// DockClientCount
//
// 获取依靠客户端总数。
func (p *TPanel) DockClientCount() int32 {
    return Panel_GetDockClientCount(p._instance())
}

// MouseInClient
//
// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (p *TPanel) MouseInClient() bool {
    return Panel_GetMouseInClient(p._instance())
}

// VisibleDockClientCount
//
// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (p *TPanel) VisibleDockClientCount() int32 {
    return Panel_GetVisibleDockClientCount(p._instance())
}

// Brush
//
// 获取画刷对象。
//
// Get Brush.
func (p *TPanel) Brush() *TBrush {
    return AsBrush(Panel_GetBrush(p._instance()))
}

// ControlCount
//
// 获取子控件数。
//
// Get Number of child controls.
func (p *TPanel) ControlCount() int32 {
    return Panel_GetControlCount(p._instance())
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (p *TPanel) Handle() HWND {
    return Panel_GetHandle(p._instance())
}

// ParentWindow
//
// 获取父容器句柄。
//
// Get Parent container handle.
func (p *TPanel) ParentWindow() HWND {
    return Panel_GetParentWindow(p._instance())
}

// SetParentWindow
//
// 设置父容器句柄。
//
// Set Parent container handle.
func (p *TPanel) SetParentWindow(value HWND) {
    Panel_SetParentWindow(p._instance(), value)
}

func (p *TPanel) Showing() bool {
    return Panel_GetShowing(p._instance())
}

func (p *TPanel) Action() *TAction {
    return AsAction(Panel_GetAction(p._instance()))
}

func (p *TPanel) SetAction(value IComponent) {
    Panel_SetAction(p._instance(), CheckPtr(value))
}

func (p *TPanel) BoundsRect() TRect {
    return Panel_GetBoundsRect(p._instance())
}

func (p *TPanel) SetBoundsRect(value TRect) {
    Panel_SetBoundsRect(p._instance(), value)
}

// ClientHeight
//
// 获取客户区高度。
//
// Get client height.
func (p *TPanel) ClientHeight() int32 {
    return Panel_GetClientHeight(p._instance())
}

// SetClientHeight
//
// 设置客户区高度。
//
// Set client height.
func (p *TPanel) SetClientHeight(value int32) {
    Panel_SetClientHeight(p._instance(), value)
}

func (p *TPanel) ClientOrigin() TPoint {
    return Panel_GetClientOrigin(p._instance())
}

// ClientRect
//
// 获取客户区矩形。
//
// Get client rectangle.
func (p *TPanel) ClientRect() TRect {
    return Panel_GetClientRect(p._instance())
}

// ClientWidth
//
// 获取客户区宽度。
//
// Get client width.
func (p *TPanel) ClientWidth() int32 {
    return Panel_GetClientWidth(p._instance())
}

// SetClientWidth
//
// 设置客户区宽度。
//
// Set client width.
func (p *TPanel) SetClientWidth(value int32) {
    Panel_SetClientWidth(p._instance(), value)
}

// ControlState
//
// 获取控件状态。
//
// Get control state.
func (p *TPanel) ControlState() TControlState {
    return Panel_GetControlState(p._instance())
}

// SetControlState
//
// 设置控件状态。
//
// Set control state.
func (p *TPanel) SetControlState(value TControlState) {
    Panel_SetControlState(p._instance(), value)
}

// ControlStyle
//
// 获取控件样式。
//
// Get control style.
func (p *TPanel) ControlStyle() TControlStyle {
    return Panel_GetControlStyle(p._instance())
}

// SetControlStyle
//
// 设置控件样式。
//
// Set control style.
func (p *TPanel) SetControlStyle(value TControlStyle) {
    Panel_SetControlStyle(p._instance(), value)
}

func (p *TPanel) Floating() bool {
    return Panel_GetFloating(p._instance())
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (p *TPanel) Parent() *TWinControl {
    return AsWinControl(Panel_GetParent(p._instance()))
}

// SetParent
//
// 设置控件父容器。
//
// Set control parent container.
func (p *TPanel) SetParent(value IWinControl) {
    Panel_SetParent(p._instance(), CheckPtr(value))
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (p *TPanel) Left() int32 {
    return Panel_GetLeft(p._instance())
}

// SetLeft
//
// 设置左边位置。
//
// Set Left position.
func (p *TPanel) SetLeft(value int32) {
    Panel_SetLeft(p._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (p *TPanel) Top() int32 {
    return Panel_GetTop(p._instance())
}

// SetTop
//
// 设置顶边位置。
//
// Set Top position.
func (p *TPanel) SetTop(value int32) {
    Panel_SetTop(p._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (p *TPanel) Width() int32 {
    return Panel_GetWidth(p._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (p *TPanel) SetWidth(value int32) {
    Panel_SetWidth(p._instance(), value)
}

// Height
//
// 获取高度。
//
// Get height.
func (p *TPanel) Height() int32 {
    return Panel_GetHeight(p._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (p *TPanel) SetHeight(value int32) {
    Panel_SetHeight(p._instance(), value)
}

// Cursor
//
// 获取控件光标。
//
// Get control cursor.
func (p *TPanel) Cursor() TCursor {
    return Panel_GetCursor(p._instance())
}

// SetCursor
//
// 设置控件光标。
//
// Set control cursor.
func (p *TPanel) SetCursor(value TCursor) {
    Panel_SetCursor(p._instance(), value)
}

// Hint
//
// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (p *TPanel) Hint() string {
    return Panel_GetHint(p._instance())
}

// SetHint
//
// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (p *TPanel) SetHint(value string) {
    Panel_SetHint(p._instance(), value)
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (p *TPanel) ComponentCount() int32 {
    return Panel_GetComponentCount(p._instance())
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (p *TPanel) ComponentIndex() int32 {
    return Panel_GetComponentIndex(p._instance())
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (p *TPanel) SetComponentIndex(value int32) {
    Panel_SetComponentIndex(p._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (p *TPanel) Owner() *TComponent {
    return AsComponent(Panel_GetOwner(p._instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (p *TPanel) Name() string {
    return Panel_GetName(p._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (p *TPanel) SetName(value string) {
    Panel_SetName(p._instance(), value)
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (p *TPanel) Tag() int {
    return Panel_GetTag(p._instance())
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (p *TPanel) SetTag(value int) {
    Panel_SetTag(p._instance(), value)
}

// AnchorSideLeft
//
// 获取左边锚点。
func (p *TPanel) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(Panel_GetAnchorSideLeft(p._instance()))
}

// SetAnchorSideLeft
//
// 设置左边锚点。
func (p *TPanel) SetAnchorSideLeft(value *TAnchorSide) {
    Panel_SetAnchorSideLeft(p._instance(), CheckPtr(value))
}

// AnchorSideTop
//
// 获取顶边锚点。
func (p *TPanel) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(Panel_GetAnchorSideTop(p._instance()))
}

// SetAnchorSideTop
//
// 设置顶边锚点。
func (p *TPanel) SetAnchorSideTop(value *TAnchorSide) {
    Panel_SetAnchorSideTop(p._instance(), CheckPtr(value))
}

// AnchorSideRight
//
// 获取右边锚点。
func (p *TPanel) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(Panel_GetAnchorSideRight(p._instance()))
}

// SetAnchorSideRight
//
// 设置右边锚点。
func (p *TPanel) SetAnchorSideRight(value *TAnchorSide) {
    Panel_SetAnchorSideRight(p._instance(), CheckPtr(value))
}

// AnchorSideBottom
//
// 获取底边锚点。
func (p *TPanel) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(Panel_GetAnchorSideBottom(p._instance()))
}

// SetAnchorSideBottom
//
// 设置底边锚点。
func (p *TPanel) SetAnchorSideBottom(value *TAnchorSide) {
    Panel_SetAnchorSideBottom(p._instance(), CheckPtr(value))
}

func (p *TPanel) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(Panel_GetChildSizing(p._instance()))
}

func (p *TPanel) SetChildSizing(value *TControlChildSizing) {
    Panel_SetChildSizing(p._instance(), CheckPtr(value))
}

// BorderSpacing
//
// 获取边框间距。
func (p *TPanel) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(Panel_GetBorderSpacing(p._instance()))
}

// SetBorderSpacing
//
// 设置边框间距。
func (p *TPanel) SetBorderSpacing(value *TControlBorderSpacing) {
    Panel_SetBorderSpacing(p._instance(), CheckPtr(value))
}

// DockClients
//
// 获取指定索引停靠客户端。
func (p *TPanel) DockClients(Index int32) *TControl {
    return AsControl(Panel_GetDockClients(p._instance(), Index))
}

// Controls
//
// 获取指定索引子控件。
func (p *TPanel) Controls(Index int32) *TControl {
    return AsControl(Panel_GetControls(p._instance(), Index))
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (p *TPanel) Components(AIndex int32) *TComponent {
    return AsComponent(Panel_GetComponents(p._instance(), AIndex))
}

// AnchorSide
//
// 获取锚侧面。
func (p *TPanel) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(Panel_GetAnchorSide(p._instance(), AKind))
}

