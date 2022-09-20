
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

type TToolBar struct {
    IWinControl
    instance unsafe.Pointer
}

// NewToolBar
//
// 创建一个新的对象。
// 
// Create a new object.
func NewToolBar(owner IComponent) *TToolBar {
    t := new(TToolBar)
    t.instance = unsafe.Pointer(ToolBar_Create(CheckPtr(owner)))
    return t
}

// AsToolBar
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsToolBar(obj interface{}) *TToolBar {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TToolBar{instance: instance}
}

// Free 
//
// 释放对象。
// 
// Free object.
func (t *TToolBar) Free() {
    if t.instance != nullptr {
        ToolBar_Free(t._instance())
        t.instance  = nullptr
    }
}

func (t *TToolBar) _instance() uintptr {
    return uintptr(t.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (t *TToolBar) Instance() uintptr {
    return t._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (t *TToolBar) UnsafeAddr() unsafe.Pointer {
    return t.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (t *TToolBar) IsValid() bool {
    return t.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (t *TToolBar) Is() TIs {
    return TIs(t._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (t *TToolBar) As() TAs {
//    return TAs(t._instance())
//}

// TToolBarClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TToolBarClass() TClass {
    return ToolBar_StaticClassType()
}

func (t *TToolBar) FlipChildren(AllLevels bool) {
    ToolBar_FlipChildren(t._instance(), AllLevels)
}

// CanFocus
//
// 是否可以获得焦点。
func (t *TToolBar) CanFocus() bool {
    return ToolBar_CanFocus(t._instance())
}

// ContainsControl
//
// 返回是否包含指定控件。
//
// it's contain a specified control.
func (t *TToolBar) ContainsControl(Control IControl) bool {
    return ToolBar_ContainsControl(t._instance(), CheckPtr(Control))
}

// ControlAtPos
//
// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (t *TToolBar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(ToolBar_ControlAtPos(t._instance(), Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
//
// 禁用控件的对齐。
//
// Disable control alignment.
func (t *TToolBar) DisableAlign() {
    ToolBar_DisableAlign(t._instance())
}

// EnableAlign
//
// 启用控件对齐。
//
// Enabled control alignment.
func (t *TToolBar) EnableAlign() {
    ToolBar_EnableAlign(t._instance())
}

// FindChildControl
//
// 查找子控件。
//
// Find sub controls.
func (t *TToolBar) FindChildControl(ControlName string) *TControl {
    return AsControl(ToolBar_FindChildControl(t._instance(), ControlName))
}

// Focused
//
// 返回是否获取焦点。
//
// Return to get focus.
func (t *TToolBar) Focused() bool {
    return ToolBar_Focused(t._instance())
}

// HandleAllocated
//
// 句柄是否已经分配。
//
// Is the handle already allocated.
func (t *TToolBar) HandleAllocated() bool {
    return ToolBar_HandleAllocated(t._instance())
}

// InsertControl
//
// 插入一个控件。
//
// Insert a control.
func (t *TToolBar) InsertControl(AControl IControl) {
    ToolBar_InsertControl(t._instance(), CheckPtr(AControl))
}

// Invalidate
//
// 要求重绘。
//
// Redraw.
func (t *TToolBar) Invalidate() {
    ToolBar_Invalidate(t._instance())
}

// PaintTo
//
// 绘画至指定DC。
//
// Painting to the specified DC.
func (t *TToolBar) PaintTo(DC HDC, X int32, Y int32) {
    ToolBar_PaintTo(t._instance(), DC , X , Y)
}

// RemoveControl
//
// 移除一个控件。
//
// Remove a control.
func (t *TToolBar) RemoveControl(AControl IControl) {
    ToolBar_RemoveControl(t._instance(), CheckPtr(AControl))
}

// Realign
//
// 重新对齐。
//
// Realign.
func (t *TToolBar) Realign() {
    ToolBar_Realign(t._instance())
}

// Repaint
//
// 重绘。
//
// Repaint.
func (t *TToolBar) Repaint() {
    ToolBar_Repaint(t._instance())
}

// ScaleBy
//
// 按比例缩放。
//
// Scale by.
func (t *TToolBar) ScaleBy(M int32, D int32) {
    ToolBar_ScaleBy(t._instance(), M , D)
}

// ScrollBy
//
// 滚动至指定位置。
//
// Scroll by.
func (t *TToolBar) ScrollBy(DeltaX int32, DeltaY int32) {
    ToolBar_ScrollBy(t._instance(), DeltaX , DeltaY)
}

// SetBounds
//
// 设置组件边界。
//
// Set component boundaries.
func (t *TToolBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ToolBar_SetBounds(t._instance(), ALeft , ATop , AWidth , AHeight)
}

// SetFocus
//
// 设置控件焦点。
//
// Set control focus.
func (t *TToolBar) SetFocus() {
    ToolBar_SetFocus(t._instance())
}

// Update
//
// 控件更新。
//
// Update.
func (t *TToolBar) Update() {
    ToolBar_Update(t._instance())
}

// BringToFront
//
// 将控件置于最前。
//
// Bring the control to the front.
func (t *TToolBar) BringToFront() {
    ToolBar_BringToFront(t._instance())
}

// ClientToScreen
//
// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (t *TToolBar) ClientToScreen(Point TPoint) TPoint {
    return ToolBar_ClientToScreen(t._instance(), Point)
}

// ClientToParent
//
// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (t *TToolBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ToolBar_ClientToParent(t._instance(), Point , CheckPtr(AParent))
}

// Dragging
//
// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (t *TToolBar) Dragging() bool {
    return ToolBar_Dragging(t._instance())
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (t *TToolBar) HasParent() bool {
    return ToolBar_HasParent(t._instance())
}

// Hide
//
// 隐藏控件。
//
// Hidden control.
func (t *TToolBar) Hide() {
    ToolBar_Hide(t._instance())
}

// Perform
//
// 发送一个消息。
//
// Send a message.
func (t *TToolBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ToolBar_Perform(t._instance(), Msg , WParam , LParam)
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (t *TToolBar) Refresh() {
    ToolBar_Refresh(t._instance())
}

// ScreenToClient
//
// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (t *TToolBar) ScreenToClient(Point TPoint) TPoint {
    return ToolBar_ScreenToClient(t._instance(), Point)
}

// ParentToClient
//
// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (t *TToolBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ToolBar_ParentToClient(t._instance(), Point , CheckPtr(AParent))
}

// SendToBack
//
// 控件至于最后面。
//
// The control is placed at the end.
func (t *TToolBar) SendToBack() {
    ToolBar_SendToBack(t._instance())
}

// Show
//
// 显示控件。
//
// Show control.
func (t *TToolBar) Show() {
    ToolBar_Show(t._instance())
}

// GetTextBuf
//
// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (t *TToolBar) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return ToolBar_GetTextBuf(t._instance(), Buffer , BufSize)
}

// GetTextLen
//
// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (t *TToolBar) GetTextLen() int32 {
    return ToolBar_GetTextLen(t._instance())
}

// SetTextBuf
//
// 设置控件字符，如果有。
//
// Set control characters, if any.
func (t *TToolBar) SetTextBuf(Buffer string) {
    ToolBar_SetTextBuf(t._instance(), Buffer)
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (t *TToolBar) FindComponent(AName string) *TComponent {
    return AsComponent(ToolBar_FindComponent(t._instance(), AName))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (t *TToolBar) GetNamePath() string {
    return ToolBar_GetNamePath(t._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (t *TToolBar) Assign(Source IObject) {
    ToolBar_Assign(t._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (t *TToolBar) ClassType() TClass {
    return ToolBar_ClassType(t._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (t *TToolBar) ClassName() string {
    return ToolBar_ClassName(t._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (t *TToolBar) InstanceSize() int32 {
    return ToolBar_InstanceSize(t._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (t *TToolBar) InheritsFrom(AClass TClass) bool {
    return ToolBar_InheritsFrom(t._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (t *TToolBar) Equals(Obj IObject) bool {
    return ToolBar_Equals(t._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (t *TToolBar) GetHashCode() int32 {
    return ToolBar_GetHashCode(t._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (t *TToolBar) ToString() string {
    return ToolBar_ToString(t._instance())
}

func (t *TToolBar) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ToolBar_AnchorToNeighbour(t._instance(), ASide , ASpace , CheckPtr(ASibling))
}

func (t *TToolBar) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ToolBar_AnchorParallel(t._instance(), ASide , ASpace , CheckPtr(ASibling))
}

// AnchorHorizontalCenterTo
//
// 置于指定控件的横向中心。
func (t *TToolBar) AnchorHorizontalCenterTo(ASibling IControl) {
    ToolBar_AnchorHorizontalCenterTo(t._instance(), CheckPtr(ASibling))
}

// AnchorVerticalCenterTo
//
// 置于指定控件的纵向中心。
func (t *TToolBar) AnchorVerticalCenterTo(ASibling IControl) {
    ToolBar_AnchorVerticalCenterTo(t._instance(), CheckPtr(ASibling))
}

func (t *TToolBar) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    ToolBar_AnchorSame(t._instance(), ASide , CheckPtr(ASibling))
}

func (t *TToolBar) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    ToolBar_AnchorAsAlign(t._instance(), ATheAlign , ASpace)
}

func (t *TToolBar) AnchorClient(ASpace int32) {
    ToolBar_AnchorClient(t._instance(), ASpace)
}

func (t *TToolBar) ScaleDesignToForm(ASize int32) int32 {
    return ToolBar_ScaleDesignToForm(t._instance(), ASize)
}

func (t *TToolBar) ScaleFormToDesign(ASize int32) int32 {
    return ToolBar_ScaleFormToDesign(t._instance(), ASize)
}

func (t *TToolBar) Scale96ToForm(ASize int32) int32 {
    return ToolBar_Scale96ToForm(t._instance(), ASize)
}

func (t *TToolBar) ScaleFormTo96(ASize int32) int32 {
    return ToolBar_ScaleFormTo96(t._instance(), ASize)
}

func (t *TToolBar) Scale96ToFont(ASize int32) int32 {
    return ToolBar_Scale96ToFont(t._instance(), ASize)
}

func (t *TToolBar) ScaleFontTo96(ASize int32) int32 {
    return ToolBar_ScaleFontTo96(t._instance(), ASize)
}

func (t *TToolBar) ScaleScreenToFont(ASize int32) int32 {
    return ToolBar_ScaleScreenToFont(t._instance(), ASize)
}

func (t *TToolBar) ScaleFontToScreen(ASize int32) int32 {
    return ToolBar_ScaleFontToScreen(t._instance(), ASize)
}

func (t *TToolBar) Scale96ToScreen(ASize int32) int32 {
    return ToolBar_Scale96ToScreen(t._instance(), ASize)
}

func (t *TToolBar) ScaleScreenTo96(ASize int32) int32 {
    return ToolBar_ScaleScreenTo96(t._instance(), ASize)
}

func (t *TToolBar) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    ToolBar_AutoAdjustLayout(t._instance(), AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (t *TToolBar) FixDesignFontsPPI(ADesignTimePPI int32) {
    ToolBar_FixDesignFontsPPI(t._instance(), ADesignTimePPI)
}

func (t *TToolBar) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    ToolBar_ScaleFontsPPI(t._instance(), AToPPI , AProportion)
}

func (t *TToolBar) ButtonCount() int32 {
    return ToolBar_GetButtonCount(t._instance())
}

// Canvas
//
// 获取画布。
func (t *TToolBar) Canvas() *TCanvas {
    return AsCanvas(ToolBar_GetCanvas(t._instance()))
}

func (t *TToolBar) RowCount() int32 {
    return ToolBar_GetRowCount(t._instance())
}

// Align
//
// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (t *TToolBar) Align() TAlign {
    return ToolBar_GetAlign(t._instance())
}

// SetAlign
//
// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (t *TToolBar) SetAlign(value TAlign) {
    ToolBar_SetAlign(t._instance(), value)
}

// Anchors
//
// 获取四个角位置的锚点。
func (t *TToolBar) Anchors() TAnchors {
    return ToolBar_GetAnchors(t._instance())
}

// SetAnchors
//
// 设置四个角位置的锚点。
func (t *TToolBar) SetAnchors(value TAnchors) {
    ToolBar_SetAnchors(t._instance(), value)
}

// AutoSize
//
// 获取自动调整大小。
func (t *TToolBar) AutoSize() bool {
    return ToolBar_GetAutoSize(t._instance())
}

// SetAutoSize
//
// 设置自动调整大小。
func (t *TToolBar) SetAutoSize(value bool) {
    ToolBar_SetAutoSize(t._instance(), value)
}

// BorderWidth
//
// 获取边框的宽度。
func (t *TToolBar) BorderWidth() int32 {
    return ToolBar_GetBorderWidth(t._instance())
}

// SetBorderWidth
//
// 设置边框的宽度。
func (t *TToolBar) SetBorderWidth(value int32) {
    ToolBar_SetBorderWidth(t._instance(), value)
}

func (t *TToolBar) ButtonHeight() int32 {
    return ToolBar_GetButtonHeight(t._instance())
}

func (t *TToolBar) SetButtonHeight(value int32) {
    ToolBar_SetButtonHeight(t._instance(), value)
}

func (t *TToolBar) ButtonWidth() int32 {
    return ToolBar_GetButtonWidth(t._instance())
}

func (t *TToolBar) SetButtonWidth(value int32) {
    ToolBar_SetButtonWidth(t._instance(), value)
}

// Caption
//
// 获取控件标题。
//
// Get the control title.
func (t *TToolBar) Caption() string {
    return ToolBar_GetCaption(t._instance())
}

// SetCaption
//
// 设置控件标题。
//
// Set the control title.
func (t *TToolBar) SetCaption(value string) {
    ToolBar_SetCaption(t._instance(), value)
}

// Color
//
// 获取颜色。
//
// Get color.
func (t *TToolBar) Color() TColor {
    return ToolBar_GetColor(t._instance())
}

// SetColor
//
// 设置颜色。
//
// Set color.
func (t *TToolBar) SetColor(value TColor) {
    ToolBar_SetColor(t._instance(), value)
}

// Constraints
//
// 获取约束控件大小。
func (t *TToolBar) Constraints() *TSizeConstraints {
    return AsSizeConstraints(ToolBar_GetConstraints(t._instance()))
}

// SetConstraints
//
// 设置约束控件大小。
func (t *TToolBar) SetConstraints(value *TSizeConstraints) {
    ToolBar_SetConstraints(t._instance(), CheckPtr(value))
}

// DoubleBuffered
//
// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (t *TToolBar) DoubleBuffered() bool {
    return ToolBar_GetDoubleBuffered(t._instance())
}

// SetDoubleBuffered
//
// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (t *TToolBar) SetDoubleBuffered(value bool) {
    ToolBar_SetDoubleBuffered(t._instance(), value)
}

// DockSite
//
// 获取停靠站点。
//
// Get Docking site.
func (t *TToolBar) DockSite() bool {
    return ToolBar_GetDockSite(t._instance())
}

// SetDockSite
//
// 设置停靠站点。
//
// Set Docking site.
func (t *TToolBar) SetDockSite(value bool) {
    ToolBar_SetDockSite(t._instance(), value)
}

// DragCursor
//
// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (t *TToolBar) DragCursor() TCursor {
    return ToolBar_GetDragCursor(t._instance())
}

// SetDragCursor
//
// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (t *TToolBar) SetDragCursor(value TCursor) {
    ToolBar_SetDragCursor(t._instance(), value)
}

// DragKind
//
// 获取拖拽方式。
//
// Get Drag and drop.
func (t *TToolBar) DragKind() TDragKind {
    return ToolBar_GetDragKind(t._instance())
}

// SetDragKind
//
// 设置拖拽方式。
//
// Set Drag and drop.
func (t *TToolBar) SetDragKind(value TDragKind) {
    ToolBar_SetDragKind(t._instance(), value)
}

// DragMode
//
// 获取拖拽模式。
//
// Get Drag mode.
func (t *TToolBar) DragMode() TDragMode {
    return ToolBar_GetDragMode(t._instance())
}

// SetDragMode
//
// 设置拖拽模式。
//
// Set Drag mode.
func (t *TToolBar) SetDragMode(value TDragMode) {
    ToolBar_SetDragMode(t._instance(), value)
}

func (t *TToolBar) EdgeBorders() TEdgeBorders {
    return ToolBar_GetEdgeBorders(t._instance())
}

func (t *TToolBar) SetEdgeBorders(value TEdgeBorders) {
    ToolBar_SetEdgeBorders(t._instance(), value)
}

func (t *TToolBar) EdgeInner() TEdgeStyle {
    return ToolBar_GetEdgeInner(t._instance())
}

func (t *TToolBar) SetEdgeInner(value TEdgeStyle) {
    ToolBar_SetEdgeInner(t._instance(), value)
}

func (t *TToolBar) EdgeOuter() TEdgeStyle {
    return ToolBar_GetEdgeOuter(t._instance())
}

func (t *TToolBar) SetEdgeOuter(value TEdgeStyle) {
    ToolBar_SetEdgeOuter(t._instance(), value)
}

// Enabled
//
// 获取控件启用。
//
// Get the control enabled.
func (t *TToolBar) Enabled() bool {
    return ToolBar_GetEnabled(t._instance())
}

// SetEnabled
//
// 设置控件启用。
//
// Set the control enabled.
func (t *TToolBar) SetEnabled(value bool) {
    ToolBar_SetEnabled(t._instance(), value)
}

// Flat
//
// 获取平面样式。
func (t *TToolBar) Flat() bool {
    return ToolBar_GetFlat(t._instance())
}

// SetFlat
//
// 设置平面样式。
func (t *TToolBar) SetFlat(value bool) {
    ToolBar_SetFlat(t._instance(), value)
}

// Font
//
// 获取字体。
//
// Get Font.
func (t *TToolBar) Font() *TFont {
    return AsFont(ToolBar_GetFont(t._instance()))
}

// SetFont
//
// 设置字体。
//
// Set Font.
func (t *TToolBar) SetFont(value *TFont) {
    ToolBar_SetFont(t._instance(), CheckPtr(value))
}

// Height
//
// 获取高度。
//
// Get height.
func (t *TToolBar) Height() int32 {
    return ToolBar_GetHeight(t._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (t *TToolBar) SetHeight(value int32) {
    ToolBar_SetHeight(t._instance(), value)
}

func (t *TToolBar) HotImages() *TImageList {
    return AsImageList(ToolBar_GetHotImages(t._instance()))
}

func (t *TToolBar) SetHotImages(value IComponent) {
    ToolBar_SetHotImages(t._instance(), CheckPtr(value))
}

// Images
//
// 获取图标索引列表对象。
func (t *TToolBar) Images() *TImageList {
    return AsImageList(ToolBar_GetImages(t._instance()))
}

// SetImages
//
// 设置图标索引列表对象。
func (t *TToolBar) SetImages(value IComponent) {
    ToolBar_SetImages(t._instance(), CheckPtr(value))
}

func (t *TToolBar) Indent() int32 {
    return ToolBar_GetIndent(t._instance())
}

func (t *TToolBar) SetIndent(value int32) {
    ToolBar_SetIndent(t._instance(), value)
}

func (t *TToolBar) List() bool {
    return ToolBar_GetList(t._instance())
}

func (t *TToolBar) SetList(value bool) {
    ToolBar_SetList(t._instance(), value)
}

// ParentColor
//
// 获取使用父容器颜色。
//
// Get parent color.
func (t *TToolBar) ParentColor() bool {
    return ToolBar_GetParentColor(t._instance())
}

// SetParentColor
//
// 设置使用父容器颜色。
//
// Set parent color.
func (t *TToolBar) SetParentColor(value bool) {
    ToolBar_SetParentColor(t._instance(), value)
}

// ParentDoubleBuffered
//
// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (t *TToolBar) ParentDoubleBuffered() bool {
    return ToolBar_GetParentDoubleBuffered(t._instance())
}

// SetParentDoubleBuffered
//
// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (t *TToolBar) SetParentDoubleBuffered(value bool) {
    ToolBar_SetParentDoubleBuffered(t._instance(), value)
}

// ParentFont
//
// 获取使用父容器字体。
//
// Get Parent container font.
func (t *TToolBar) ParentFont() bool {
    return ToolBar_GetParentFont(t._instance())
}

// SetParentFont
//
// 设置使用父容器字体。
//
// Set Parent container font.
func (t *TToolBar) SetParentFont(value bool) {
    ToolBar_SetParentFont(t._instance(), value)
}

// ParentShowHint
//
// 获取以父容器的ShowHint属性为准。
func (t *TToolBar) ParentShowHint() bool {
    return ToolBar_GetParentShowHint(t._instance())
}

// SetParentShowHint
//
// 设置以父容器的ShowHint属性为准。
func (t *TToolBar) SetParentShowHint(value bool) {
    ToolBar_SetParentShowHint(t._instance(), value)
}

// PopupMenu
//
// 获取右键菜单。
//
// Get Right click menu.
func (t *TToolBar) PopupMenu() *TPopupMenu {
    return AsPopupMenu(ToolBar_GetPopupMenu(t._instance()))
}

// SetPopupMenu
//
// 设置右键菜单。
//
// Set Right click menu.
func (t *TToolBar) SetPopupMenu(value IComponent) {
    ToolBar_SetPopupMenu(t._instance(), CheckPtr(value))
}

// ShowCaptions
//
// 获取显示标题。
func (t *TToolBar) ShowCaptions() bool {
    return ToolBar_GetShowCaptions(t._instance())
}

// SetShowCaptions
//
// 设置显示标题。
func (t *TToolBar) SetShowCaptions(value bool) {
    ToolBar_SetShowCaptions(t._instance(), value)
}

// ShowHint
//
// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (t *TToolBar) ShowHint() bool {
    return ToolBar_GetShowHint(t._instance())
}

// SetShowHint
//
// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (t *TToolBar) SetShowHint(value bool) {
    ToolBar_SetShowHint(t._instance(), value)
}

// TabOrder
//
// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (t *TToolBar) TabOrder() TTabOrder {
    return ToolBar_GetTabOrder(t._instance())
}

// SetTabOrder
//
// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (t *TToolBar) SetTabOrder(value TTabOrder) {
    ToolBar_SetTabOrder(t._instance(), value)
}

// TabStop
//
// 获取Tab可停留。
//
// Get Tab can stay.
func (t *TToolBar) TabStop() bool {
    return ToolBar_GetTabStop(t._instance())
}

// SetTabStop
//
// 设置Tab可停留。
//
// Set Tab can stay.
func (t *TToolBar) SetTabStop(value bool) {
    ToolBar_SetTabStop(t._instance(), value)
}

// Transparent
//
// 获取透明。
//
// Get transparent.
func (t *TToolBar) Transparent() bool {
    return ToolBar_GetTransparent(t._instance())
}

// SetTransparent
//
// 设置透明。
//
// Set transparent.
func (t *TToolBar) SetTransparent(value bool) {
    ToolBar_SetTransparent(t._instance(), value)
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (t *TToolBar) Visible() bool {
    return ToolBar_GetVisible(t._instance())
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (t *TToolBar) SetVisible(value bool) {
    ToolBar_SetVisible(t._instance(), value)
}

func (t *TToolBar) Wrapable() bool {
    return ToolBar_GetWrapable(t._instance())
}

func (t *TToolBar) SetWrapable(value bool) {
    ToolBar_SetWrapable(t._instance(), value)
}

// SetOnClick
//
// 设置控件单击事件。
//
// Set control click event.
func (t *TToolBar) SetOnClick(fn TNotifyEvent) {
    ToolBar_SetOnClick(t._instance(), fn)
}

// SetOnContextPopup
//
// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (t *TToolBar) SetOnContextPopup(fn TContextPopupEvent) {
    ToolBar_SetOnContextPopup(t._instance(), fn)
}

// SetOnDblClick
//
// 设置双击事件。
func (t *TToolBar) SetOnDblClick(fn TNotifyEvent) {
    ToolBar_SetOnDblClick(t._instance(), fn)
}

func (t *TToolBar) SetOnDockDrop(fn TDockDropEvent) {
    ToolBar_SetOnDockDrop(t._instance(), fn)
}

// SetOnDragDrop
//
// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (t *TToolBar) SetOnDragDrop(fn TDragDropEvent) {
    ToolBar_SetOnDragDrop(t._instance(), fn)
}

// SetOnDragOver
//
// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (t *TToolBar) SetOnDragOver(fn TDragOverEvent) {
    ToolBar_SetOnDragOver(t._instance(), fn)
}

// SetOnEndDrag
//
// 设置拖拽结束。
//
// Set End of drag.
func (t *TToolBar) SetOnEndDrag(fn TEndDragEvent) {
    ToolBar_SetOnEndDrag(t._instance(), fn)
}

// SetOnEnter
//
// 设置焦点进入。
//
// Set Focus entry.
func (t *TToolBar) SetOnEnter(fn TNotifyEvent) {
    ToolBar_SetOnEnter(t._instance(), fn)
}

// SetOnExit
//
// 设置焦点退出。
//
// Set Focus exit.
func (t *TToolBar) SetOnExit(fn TNotifyEvent) {
    ToolBar_SetOnExit(t._instance(), fn)
}

// SetOnMouseDown
//
// 设置鼠标按下事件。
//
// Set Mouse down event.
func (t *TToolBar) SetOnMouseDown(fn TMouseEvent) {
    ToolBar_SetOnMouseDown(t._instance(), fn)
}

// SetOnMouseEnter
//
// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (t *TToolBar) SetOnMouseEnter(fn TNotifyEvent) {
    ToolBar_SetOnMouseEnter(t._instance(), fn)
}

// SetOnMouseLeave
//
// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (t *TToolBar) SetOnMouseLeave(fn TNotifyEvent) {
    ToolBar_SetOnMouseLeave(t._instance(), fn)
}

// SetOnMouseMove
//
// 设置鼠标移动事件。
func (t *TToolBar) SetOnMouseMove(fn TMouseMoveEvent) {
    ToolBar_SetOnMouseMove(t._instance(), fn)
}

// SetOnMouseUp
//
// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (t *TToolBar) SetOnMouseUp(fn TMouseEvent) {
    ToolBar_SetOnMouseUp(t._instance(), fn)
}

// SetOnResize
//
// 设置大小被改变事件。
func (t *TToolBar) SetOnResize(fn TNotifyEvent) {
    ToolBar_SetOnResize(t._instance(), fn)
}

func (t *TToolBar) SetOnUnDock(fn TUnDockEvent) {
    ToolBar_SetOnUnDock(t._instance(), fn)
}

// DockClientCount
//
// 获取依靠客户端总数。
func (t *TToolBar) DockClientCount() int32 {
    return ToolBar_GetDockClientCount(t._instance())
}

// MouseInClient
//
// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (t *TToolBar) MouseInClient() bool {
    return ToolBar_GetMouseInClient(t._instance())
}

// VisibleDockClientCount
//
// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (t *TToolBar) VisibleDockClientCount() int32 {
    return ToolBar_GetVisibleDockClientCount(t._instance())
}

// Brush
//
// 获取画刷对象。
//
// Get Brush.
func (t *TToolBar) Brush() *TBrush {
    return AsBrush(ToolBar_GetBrush(t._instance()))
}

// ControlCount
//
// 获取子控件数。
//
// Get Number of child controls.
func (t *TToolBar) ControlCount() int32 {
    return ToolBar_GetControlCount(t._instance())
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (t *TToolBar) Handle() HWND {
    return ToolBar_GetHandle(t._instance())
}

// ParentWindow
//
// 获取父容器句柄。
//
// Get Parent container handle.
func (t *TToolBar) ParentWindow() HWND {
    return ToolBar_GetParentWindow(t._instance())
}

// SetParentWindow
//
// 设置父容器句柄。
//
// Set Parent container handle.
func (t *TToolBar) SetParentWindow(value HWND) {
    ToolBar_SetParentWindow(t._instance(), value)
}

func (t *TToolBar) Showing() bool {
    return ToolBar_GetShowing(t._instance())
}

// UseDockManager
//
// 获取使用停靠管理。
func (t *TToolBar) UseDockManager() bool {
    return ToolBar_GetUseDockManager(t._instance())
}

// SetUseDockManager
//
// 设置使用停靠管理。
func (t *TToolBar) SetUseDockManager(value bool) {
    ToolBar_SetUseDockManager(t._instance(), value)
}

func (t *TToolBar) Action() *TAction {
    return AsAction(ToolBar_GetAction(t._instance()))
}

func (t *TToolBar) SetAction(value IComponent) {
    ToolBar_SetAction(t._instance(), CheckPtr(value))
}

func (t *TToolBar) BiDiMode() TBiDiMode {
    return ToolBar_GetBiDiMode(t._instance())
}

func (t *TToolBar) SetBiDiMode(value TBiDiMode) {
    ToolBar_SetBiDiMode(t._instance(), value)
}

func (t *TToolBar) BoundsRect() TRect {
    return ToolBar_GetBoundsRect(t._instance())
}

func (t *TToolBar) SetBoundsRect(value TRect) {
    ToolBar_SetBoundsRect(t._instance(), value)
}

// ClientHeight
//
// 获取客户区高度。
//
// Get client height.
func (t *TToolBar) ClientHeight() int32 {
    return ToolBar_GetClientHeight(t._instance())
}

// SetClientHeight
//
// 设置客户区高度。
//
// Set client height.
func (t *TToolBar) SetClientHeight(value int32) {
    ToolBar_SetClientHeight(t._instance(), value)
}

func (t *TToolBar) ClientOrigin() TPoint {
    return ToolBar_GetClientOrigin(t._instance())
}

// ClientRect
//
// 获取客户区矩形。
//
// Get client rectangle.
func (t *TToolBar) ClientRect() TRect {
    return ToolBar_GetClientRect(t._instance())
}

// ClientWidth
//
// 获取客户区宽度。
//
// Get client width.
func (t *TToolBar) ClientWidth() int32 {
    return ToolBar_GetClientWidth(t._instance())
}

// SetClientWidth
//
// 设置客户区宽度。
//
// Set client width.
func (t *TToolBar) SetClientWidth(value int32) {
    ToolBar_SetClientWidth(t._instance(), value)
}

// ControlState
//
// 获取控件状态。
//
// Get control state.
func (t *TToolBar) ControlState() TControlState {
    return ToolBar_GetControlState(t._instance())
}

// SetControlState
//
// 设置控件状态。
//
// Set control state.
func (t *TToolBar) SetControlState(value TControlState) {
    ToolBar_SetControlState(t._instance(), value)
}

// ControlStyle
//
// 获取控件样式。
//
// Get control style.
func (t *TToolBar) ControlStyle() TControlStyle {
    return ToolBar_GetControlStyle(t._instance())
}

// SetControlStyle
//
// 设置控件样式。
//
// Set control style.
func (t *TToolBar) SetControlStyle(value TControlStyle) {
    ToolBar_SetControlStyle(t._instance(), value)
}

func (t *TToolBar) Floating() bool {
    return ToolBar_GetFloating(t._instance())
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (t *TToolBar) Parent() *TWinControl {
    return AsWinControl(ToolBar_GetParent(t._instance()))
}

// SetParent
//
// 设置控件父容器。
//
// Set control parent container.
func (t *TToolBar) SetParent(value IWinControl) {
    ToolBar_SetParent(t._instance(), CheckPtr(value))
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (t *TToolBar) Left() int32 {
    return ToolBar_GetLeft(t._instance())
}

// SetLeft
//
// 设置左边位置。
//
// Set Left position.
func (t *TToolBar) SetLeft(value int32) {
    ToolBar_SetLeft(t._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (t *TToolBar) Top() int32 {
    return ToolBar_GetTop(t._instance())
}

// SetTop
//
// 设置顶边位置。
//
// Set Top position.
func (t *TToolBar) SetTop(value int32) {
    ToolBar_SetTop(t._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (t *TToolBar) Width() int32 {
    return ToolBar_GetWidth(t._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (t *TToolBar) SetWidth(value int32) {
    ToolBar_SetWidth(t._instance(), value)
}

// Cursor
//
// 获取控件光标。
//
// Get control cursor.
func (t *TToolBar) Cursor() TCursor {
    return ToolBar_GetCursor(t._instance())
}

// SetCursor
//
// 设置控件光标。
//
// Set control cursor.
func (t *TToolBar) SetCursor(value TCursor) {
    ToolBar_SetCursor(t._instance(), value)
}

// Hint
//
// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (t *TToolBar) Hint() string {
    return ToolBar_GetHint(t._instance())
}

// SetHint
//
// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (t *TToolBar) SetHint(value string) {
    ToolBar_SetHint(t._instance(), value)
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (t *TToolBar) ComponentCount() int32 {
    return ToolBar_GetComponentCount(t._instance())
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (t *TToolBar) ComponentIndex() int32 {
    return ToolBar_GetComponentIndex(t._instance())
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (t *TToolBar) SetComponentIndex(value int32) {
    ToolBar_SetComponentIndex(t._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (t *TToolBar) Owner() *TComponent {
    return AsComponent(ToolBar_GetOwner(t._instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (t *TToolBar) Name() string {
    return ToolBar_GetName(t._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (t *TToolBar) SetName(value string) {
    ToolBar_SetName(t._instance(), value)
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (t *TToolBar) Tag() int {
    return ToolBar_GetTag(t._instance())
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (t *TToolBar) SetTag(value int) {
    ToolBar_SetTag(t._instance(), value)
}

// AnchorSideLeft
//
// 获取左边锚点。
func (t *TToolBar) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(ToolBar_GetAnchorSideLeft(t._instance()))
}

// SetAnchorSideLeft
//
// 设置左边锚点。
func (t *TToolBar) SetAnchorSideLeft(value *TAnchorSide) {
    ToolBar_SetAnchorSideLeft(t._instance(), CheckPtr(value))
}

// AnchorSideTop
//
// 获取顶边锚点。
func (t *TToolBar) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(ToolBar_GetAnchorSideTop(t._instance()))
}

// SetAnchorSideTop
//
// 设置顶边锚点。
func (t *TToolBar) SetAnchorSideTop(value *TAnchorSide) {
    ToolBar_SetAnchorSideTop(t._instance(), CheckPtr(value))
}

// AnchorSideRight
//
// 获取右边锚点。
func (t *TToolBar) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(ToolBar_GetAnchorSideRight(t._instance()))
}

// SetAnchorSideRight
//
// 设置右边锚点。
func (t *TToolBar) SetAnchorSideRight(value *TAnchorSide) {
    ToolBar_SetAnchorSideRight(t._instance(), CheckPtr(value))
}

// AnchorSideBottom
//
// 获取底边锚点。
func (t *TToolBar) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(ToolBar_GetAnchorSideBottom(t._instance()))
}

// SetAnchorSideBottom
//
// 设置底边锚点。
func (t *TToolBar) SetAnchorSideBottom(value *TAnchorSide) {
    ToolBar_SetAnchorSideBottom(t._instance(), CheckPtr(value))
}

func (t *TToolBar) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(ToolBar_GetChildSizing(t._instance()))
}

func (t *TToolBar) SetChildSizing(value *TControlChildSizing) {
    ToolBar_SetChildSizing(t._instance(), CheckPtr(value))
}

// BorderSpacing
//
// 获取边框间距。
func (t *TToolBar) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(ToolBar_GetBorderSpacing(t._instance()))
}

// SetBorderSpacing
//
// 设置边框间距。
func (t *TToolBar) SetBorderSpacing(value *TControlBorderSpacing) {
    ToolBar_SetBorderSpacing(t._instance(), CheckPtr(value))
}

func (t *TToolBar) Buttons(Index int32) *TToolButton {
    return AsToolButton(ToolBar_GetButtons(t._instance(), Index))
}

// DockClients
//
// 获取指定索引停靠客户端。
func (t *TToolBar) DockClients(Index int32) *TControl {
    return AsControl(ToolBar_GetDockClients(t._instance(), Index))
}

// Controls
//
// 获取指定索引子控件。
func (t *TToolBar) Controls(Index int32) *TControl {
    return AsControl(ToolBar_GetControls(t._instance(), Index))
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (t *TToolBar) Components(AIndex int32) *TComponent {
    return AsComponent(ToolBar_GetComponents(t._instance(), AIndex))
}

// AnchorSide
//
// 获取锚侧面。
func (t *TToolBar) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(ToolBar_GetAnchorSide(t._instance(), AKind))
}

