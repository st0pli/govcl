
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

type TListBox struct {
    IWinControl
    instance unsafe.Pointer
}

// NewListBox
//
// 创建一个新的对象。
// 
// Create a new object.
func NewListBox(owner IComponent) *TListBox {
    l := new(TListBox)
    l.instance = unsafe.Pointer(ListBox_Create(CheckPtr(owner)))
    return l
}

// AsListBox
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsListBox(obj interface{}) *TListBox {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TListBox{instance: instance}
}

// Free 
//
// 释放对象。
// 
// Free object.
func (l *TListBox) Free() {
    if l.instance != nullptr {
        ListBox_Free(l._instance())
        l.instance  = nullptr
    }
}

func (l *TListBox) _instance() uintptr {
    return uintptr(l.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (l *TListBox) Instance() uintptr {
    return l._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (l *TListBox) UnsafeAddr() unsafe.Pointer {
    return l.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (l *TListBox) IsValid() bool {
    return l.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (l *TListBox) Is() TIs {
    return TIs(l._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (l *TListBox) As() TAs {
//    return TAs(l._instance())
//}

// TListBoxClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TListBoxClass() TClass {
    return ListBox_StaticClassType()
}

func (l *TListBox) AddItem(Item string, AObject IObject) {
    ListBox_AddItem(l._instance(), Item , CheckPtr(AObject))
}

// Clear
//
// 清除。
func (l *TListBox) Clear() {
    ListBox_Clear(l._instance())
}

// ClearSelection
//
// 清除选择。
func (l *TListBox) ClearSelection() {
    ListBox_ClearSelection(l._instance())
}

// DeleteSelected
//
// 删除选择的。
func (l *TListBox) DeleteSelected() {
    ListBox_DeleteSelected(l._instance())
}

func (l *TListBox) ItemAtPos(Pos TPoint, Existing bool) int32 {
    return ListBox_ItemAtPos(l._instance(), Pos , Existing)
}

func (l *TListBox) ItemRect(Index int32) TRect {
    return ListBox_ItemRect(l._instance(), Index)
}

// SelectAll
//
// 全选。
func (l *TListBox) SelectAll() {
    ListBox_SelectAll(l._instance())
}

// CanFocus
//
// 是否可以获得焦点。
func (l *TListBox) CanFocus() bool {
    return ListBox_CanFocus(l._instance())
}

// ContainsControl
//
// 返回是否包含指定控件。
//
// it's contain a specified control.
func (l *TListBox) ContainsControl(Control IControl) bool {
    return ListBox_ContainsControl(l._instance(), CheckPtr(Control))
}

// ControlAtPos
//
// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (l *TListBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(ListBox_ControlAtPos(l._instance(), Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
//
// 禁用控件的对齐。
//
// Disable control alignment.
func (l *TListBox) DisableAlign() {
    ListBox_DisableAlign(l._instance())
}

// EnableAlign
//
// 启用控件对齐。
//
// Enabled control alignment.
func (l *TListBox) EnableAlign() {
    ListBox_EnableAlign(l._instance())
}

// FindChildControl
//
// 查找子控件。
//
// Find sub controls.
func (l *TListBox) FindChildControl(ControlName string) *TControl {
    return AsControl(ListBox_FindChildControl(l._instance(), ControlName))
}

func (l *TListBox) FlipChildren(AllLevels bool) {
    ListBox_FlipChildren(l._instance(), AllLevels)
}

// Focused
//
// 返回是否获取焦点。
//
// Return to get focus.
func (l *TListBox) Focused() bool {
    return ListBox_Focused(l._instance())
}

// HandleAllocated
//
// 句柄是否已经分配。
//
// Is the handle already allocated.
func (l *TListBox) HandleAllocated() bool {
    return ListBox_HandleAllocated(l._instance())
}

// InsertControl
//
// 插入一个控件。
//
// Insert a control.
func (l *TListBox) InsertControl(AControl IControl) {
    ListBox_InsertControl(l._instance(), CheckPtr(AControl))
}

// Invalidate
//
// 要求重绘。
//
// Redraw.
func (l *TListBox) Invalidate() {
    ListBox_Invalidate(l._instance())
}

// PaintTo
//
// 绘画至指定DC。
//
// Painting to the specified DC.
func (l *TListBox) PaintTo(DC HDC, X int32, Y int32) {
    ListBox_PaintTo(l._instance(), DC , X , Y)
}

// RemoveControl
//
// 移除一个控件。
//
// Remove a control.
func (l *TListBox) RemoveControl(AControl IControl) {
    ListBox_RemoveControl(l._instance(), CheckPtr(AControl))
}

// Realign
//
// 重新对齐。
//
// Realign.
func (l *TListBox) Realign() {
    ListBox_Realign(l._instance())
}

// Repaint
//
// 重绘。
//
// Repaint.
func (l *TListBox) Repaint() {
    ListBox_Repaint(l._instance())
}

// ScaleBy
//
// 按比例缩放。
//
// Scale by.
func (l *TListBox) ScaleBy(M int32, D int32) {
    ListBox_ScaleBy(l._instance(), M , D)
}

// ScrollBy
//
// 滚动至指定位置。
//
// Scroll by.
func (l *TListBox) ScrollBy(DeltaX int32, DeltaY int32) {
    ListBox_ScrollBy(l._instance(), DeltaX , DeltaY)
}

// SetBounds
//
// 设置组件边界。
//
// Set component boundaries.
func (l *TListBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ListBox_SetBounds(l._instance(), ALeft , ATop , AWidth , AHeight)
}

// SetFocus
//
// 设置控件焦点。
//
// Set control focus.
func (l *TListBox) SetFocus() {
    ListBox_SetFocus(l._instance())
}

// Update
//
// 控件更新。
//
// Update.
func (l *TListBox) Update() {
    ListBox_Update(l._instance())
}

// BringToFront
//
// 将控件置于最前。
//
// Bring the control to the front.
func (l *TListBox) BringToFront() {
    ListBox_BringToFront(l._instance())
}

// ClientToScreen
//
// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (l *TListBox) ClientToScreen(Point TPoint) TPoint {
    return ListBox_ClientToScreen(l._instance(), Point)
}

// ClientToParent
//
// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (l *TListBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ListBox_ClientToParent(l._instance(), Point , CheckPtr(AParent))
}

// Dragging
//
// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (l *TListBox) Dragging() bool {
    return ListBox_Dragging(l._instance())
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (l *TListBox) HasParent() bool {
    return ListBox_HasParent(l._instance())
}

// Hide
//
// 隐藏控件。
//
// Hidden control.
func (l *TListBox) Hide() {
    ListBox_Hide(l._instance())
}

// Perform
//
// 发送一个消息。
//
// Send a message.
func (l *TListBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ListBox_Perform(l._instance(), Msg , WParam , LParam)
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (l *TListBox) Refresh() {
    ListBox_Refresh(l._instance())
}

// ScreenToClient
//
// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (l *TListBox) ScreenToClient(Point TPoint) TPoint {
    return ListBox_ScreenToClient(l._instance(), Point)
}

// ParentToClient
//
// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (l *TListBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ListBox_ParentToClient(l._instance(), Point , CheckPtr(AParent))
}

// SendToBack
//
// 控件至于最后面。
//
// The control is placed at the end.
func (l *TListBox) SendToBack() {
    ListBox_SendToBack(l._instance())
}

// Show
//
// 显示控件。
//
// Show control.
func (l *TListBox) Show() {
    ListBox_Show(l._instance())
}

// GetTextBuf
//
// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (l *TListBox) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return ListBox_GetTextBuf(l._instance(), Buffer , BufSize)
}

// GetTextLen
//
// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (l *TListBox) GetTextLen() int32 {
    return ListBox_GetTextLen(l._instance())
}

// SetTextBuf
//
// 设置控件字符，如果有。
//
// Set control characters, if any.
func (l *TListBox) SetTextBuf(Buffer string) {
    ListBox_SetTextBuf(l._instance(), Buffer)
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (l *TListBox) FindComponent(AName string) *TComponent {
    return AsComponent(ListBox_FindComponent(l._instance(), AName))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (l *TListBox) GetNamePath() string {
    return ListBox_GetNamePath(l._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (l *TListBox) Assign(Source IObject) {
    ListBox_Assign(l._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (l *TListBox) ClassType() TClass {
    return ListBox_ClassType(l._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (l *TListBox) ClassName() string {
    return ListBox_ClassName(l._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (l *TListBox) InstanceSize() int32 {
    return ListBox_InstanceSize(l._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (l *TListBox) InheritsFrom(AClass TClass) bool {
    return ListBox_InheritsFrom(l._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (l *TListBox) Equals(Obj IObject) bool {
    return ListBox_Equals(l._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (l *TListBox) GetHashCode() int32 {
    return ListBox_GetHashCode(l._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (l *TListBox) ToString() string {
    return ListBox_ToString(l._instance())
}

func (l *TListBox) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ListBox_AnchorToNeighbour(l._instance(), ASide , ASpace , CheckPtr(ASibling))
}

func (l *TListBox) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ListBox_AnchorParallel(l._instance(), ASide , ASpace , CheckPtr(ASibling))
}

// AnchorHorizontalCenterTo
//
// 置于指定控件的横向中心。
func (l *TListBox) AnchorHorizontalCenterTo(ASibling IControl) {
    ListBox_AnchorHorizontalCenterTo(l._instance(), CheckPtr(ASibling))
}

// AnchorVerticalCenterTo
//
// 置于指定控件的纵向中心。
func (l *TListBox) AnchorVerticalCenterTo(ASibling IControl) {
    ListBox_AnchorVerticalCenterTo(l._instance(), CheckPtr(ASibling))
}

func (l *TListBox) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    ListBox_AnchorSame(l._instance(), ASide , CheckPtr(ASibling))
}

func (l *TListBox) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    ListBox_AnchorAsAlign(l._instance(), ATheAlign , ASpace)
}

func (l *TListBox) AnchorClient(ASpace int32) {
    ListBox_AnchorClient(l._instance(), ASpace)
}

func (l *TListBox) ScaleDesignToForm(ASize int32) int32 {
    return ListBox_ScaleDesignToForm(l._instance(), ASize)
}

func (l *TListBox) ScaleFormToDesign(ASize int32) int32 {
    return ListBox_ScaleFormToDesign(l._instance(), ASize)
}

func (l *TListBox) Scale96ToForm(ASize int32) int32 {
    return ListBox_Scale96ToForm(l._instance(), ASize)
}

func (l *TListBox) ScaleFormTo96(ASize int32) int32 {
    return ListBox_ScaleFormTo96(l._instance(), ASize)
}

func (l *TListBox) Scale96ToFont(ASize int32) int32 {
    return ListBox_Scale96ToFont(l._instance(), ASize)
}

func (l *TListBox) ScaleFontTo96(ASize int32) int32 {
    return ListBox_ScaleFontTo96(l._instance(), ASize)
}

func (l *TListBox) ScaleScreenToFont(ASize int32) int32 {
    return ListBox_ScaleScreenToFont(l._instance(), ASize)
}

func (l *TListBox) ScaleFontToScreen(ASize int32) int32 {
    return ListBox_ScaleFontToScreen(l._instance(), ASize)
}

func (l *TListBox) Scale96ToScreen(ASize int32) int32 {
    return ListBox_Scale96ToScreen(l._instance(), ASize)
}

func (l *TListBox) ScaleScreenTo96(ASize int32) int32 {
    return ListBox_ScaleScreenTo96(l._instance(), ASize)
}

func (l *TListBox) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    ListBox_AutoAdjustLayout(l._instance(), AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (l *TListBox) FixDesignFontsPPI(ADesignTimePPI int32) {
    ListBox_FixDesignFontsPPI(l._instance(), ADesignTimePPI)
}

func (l *TListBox) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    ListBox_ScaleFontsPPI(l._instance(), AToPPI , AProportion)
}

func (l *TListBox) ClickOnSelChange() bool {
    return ListBox_GetClickOnSelChange(l._instance())
}

func (l *TListBox) SetClickOnSelChange(value bool) {
    ListBox_SetClickOnSelChange(l._instance(), value)
}

func (l *TListBox) Options() TListBoxOptions {
    return ListBox_GetOptions(l._instance())
}

func (l *TListBox) SetOptions(value TListBoxOptions) {
    ListBox_SetOptions(l._instance(), value)
}

func (l *TListBox) TopIndex() int32 {
    return ListBox_GetTopIndex(l._instance())
}

func (l *TListBox) SetTopIndex(value int32) {
    ListBox_SetTopIndex(l._instance(), value)
}

func (l *TListBox) Style() TListBoxStyle {
    return ListBox_GetStyle(l._instance())
}

func (l *TListBox) SetStyle(value TListBoxStyle) {
    ListBox_SetStyle(l._instance(), value)
}

// Align
//
// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (l *TListBox) Align() TAlign {
    return ListBox_GetAlign(l._instance())
}

// SetAlign
//
// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (l *TListBox) SetAlign(value TAlign) {
    ListBox_SetAlign(l._instance(), value)
}

// Anchors
//
// 获取四个角位置的锚点。
func (l *TListBox) Anchors() TAnchors {
    return ListBox_GetAnchors(l._instance())
}

// SetAnchors
//
// 设置四个角位置的锚点。
func (l *TListBox) SetAnchors(value TAnchors) {
    ListBox_SetAnchors(l._instance(), value)
}

func (l *TListBox) BiDiMode() TBiDiMode {
    return ListBox_GetBiDiMode(l._instance())
}

func (l *TListBox) SetBiDiMode(value TBiDiMode) {
    ListBox_SetBiDiMode(l._instance(), value)
}

// BorderStyle
//
// 获取窗口边框样式。比如：无边框，单一边框等。
func (l *TListBox) BorderStyle() TBorderStyle {
    return ListBox_GetBorderStyle(l._instance())
}

// SetBorderStyle
//
// 设置窗口边框样式。比如：无边框，单一边框等。
func (l *TListBox) SetBorderStyle(value TBorderStyle) {
    ListBox_SetBorderStyle(l._instance(), value)
}

// Color
//
// 获取颜色。
//
// Get color.
func (l *TListBox) Color() TColor {
    return ListBox_GetColor(l._instance())
}

// SetColor
//
// 设置颜色。
//
// Set color.
func (l *TListBox) SetColor(value TColor) {
    ListBox_SetColor(l._instance(), value)
}

func (l *TListBox) Columns() int32 {
    return ListBox_GetColumns(l._instance())
}

func (l *TListBox) SetColumns(value int32) {
    ListBox_SetColumns(l._instance(), value)
}

// Constraints
//
// 获取约束控件大小。
func (l *TListBox) Constraints() *TSizeConstraints {
    return AsSizeConstraints(ListBox_GetConstraints(l._instance()))
}

// SetConstraints
//
// 设置约束控件大小。
func (l *TListBox) SetConstraints(value *TSizeConstraints) {
    ListBox_SetConstraints(l._instance(), CheckPtr(value))
}

// DoubleBuffered
//
// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (l *TListBox) DoubleBuffered() bool {
    return ListBox_GetDoubleBuffered(l._instance())
}

// SetDoubleBuffered
//
// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (l *TListBox) SetDoubleBuffered(value bool) {
    ListBox_SetDoubleBuffered(l._instance(), value)
}

// DragCursor
//
// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (l *TListBox) DragCursor() TCursor {
    return ListBox_GetDragCursor(l._instance())
}

// SetDragCursor
//
// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (l *TListBox) SetDragCursor(value TCursor) {
    ListBox_SetDragCursor(l._instance(), value)
}

// DragKind
//
// 获取拖拽方式。
//
// Get Drag and drop.
func (l *TListBox) DragKind() TDragKind {
    return ListBox_GetDragKind(l._instance())
}

// SetDragKind
//
// 设置拖拽方式。
//
// Set Drag and drop.
func (l *TListBox) SetDragKind(value TDragKind) {
    ListBox_SetDragKind(l._instance(), value)
}

// DragMode
//
// 获取拖拽模式。
//
// Get Drag mode.
func (l *TListBox) DragMode() TDragMode {
    return ListBox_GetDragMode(l._instance())
}

// SetDragMode
//
// 设置拖拽模式。
//
// Set Drag mode.
func (l *TListBox) SetDragMode(value TDragMode) {
    ListBox_SetDragMode(l._instance(), value)
}

// Enabled
//
// 获取控件启用。
//
// Get the control enabled.
func (l *TListBox) Enabled() bool {
    return ListBox_GetEnabled(l._instance())
}

// SetEnabled
//
// 设置控件启用。
//
// Set the control enabled.
func (l *TListBox) SetEnabled(value bool) {
    ListBox_SetEnabled(l._instance(), value)
}

func (l *TListBox) ExtendedSelect() bool {
    return ListBox_GetExtendedSelect(l._instance())
}

func (l *TListBox) SetExtendedSelect(value bool) {
    ListBox_SetExtendedSelect(l._instance(), value)
}

// Font
//
// 获取字体。
//
// Get Font.
func (l *TListBox) Font() *TFont {
    return AsFont(ListBox_GetFont(l._instance()))
}

// SetFont
//
// 设置字体。
//
// Set Font.
func (l *TListBox) SetFont(value *TFont) {
    ListBox_SetFont(l._instance(), CheckPtr(value))
}

func (l *TListBox) ItemHeight() int32 {
    return ListBox_GetItemHeight(l._instance())
}

func (l *TListBox) SetItemHeight(value int32) {
    ListBox_SetItemHeight(l._instance(), value)
}

func (l *TListBox) Items() *TStrings {
    return AsStrings(ListBox_GetItems(l._instance()))
}

func (l *TListBox) SetItems(value IStrings) {
    ListBox_SetItems(l._instance(), CheckPtr(value))
}

func (l *TListBox) MultiSelect() bool {
    return ListBox_GetMultiSelect(l._instance())
}

func (l *TListBox) SetMultiSelect(value bool) {
    ListBox_SetMultiSelect(l._instance(), value)
}

// ParentColor
//
// 获取使用父容器颜色。
//
// Get parent color.
func (l *TListBox) ParentColor() bool {
    return ListBox_GetParentColor(l._instance())
}

// SetParentColor
//
// 设置使用父容器颜色。
//
// Set parent color.
func (l *TListBox) SetParentColor(value bool) {
    ListBox_SetParentColor(l._instance(), value)
}

// ParentDoubleBuffered
//
// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (l *TListBox) ParentDoubleBuffered() bool {
    return ListBox_GetParentDoubleBuffered(l._instance())
}

// SetParentDoubleBuffered
//
// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (l *TListBox) SetParentDoubleBuffered(value bool) {
    ListBox_SetParentDoubleBuffered(l._instance(), value)
}

// ParentFont
//
// 获取使用父容器字体。
//
// Get Parent container font.
func (l *TListBox) ParentFont() bool {
    return ListBox_GetParentFont(l._instance())
}

// SetParentFont
//
// 设置使用父容器字体。
//
// Set Parent container font.
func (l *TListBox) SetParentFont(value bool) {
    ListBox_SetParentFont(l._instance(), value)
}

// ParentShowHint
//
// 获取以父容器的ShowHint属性为准。
func (l *TListBox) ParentShowHint() bool {
    return ListBox_GetParentShowHint(l._instance())
}

// SetParentShowHint
//
// 设置以父容器的ShowHint属性为准。
func (l *TListBox) SetParentShowHint(value bool) {
    ListBox_SetParentShowHint(l._instance(), value)
}

// PopupMenu
//
// 获取右键菜单。
//
// Get Right click menu.
func (l *TListBox) PopupMenu() *TPopupMenu {
    return AsPopupMenu(ListBox_GetPopupMenu(l._instance()))
}

// SetPopupMenu
//
// 设置右键菜单。
//
// Set Right click menu.
func (l *TListBox) SetPopupMenu(value IComponent) {
    ListBox_SetPopupMenu(l._instance(), CheckPtr(value))
}

// ShowHint
//
// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (l *TListBox) ShowHint() bool {
    return ListBox_GetShowHint(l._instance())
}

// SetShowHint
//
// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (l *TListBox) SetShowHint(value bool) {
    ListBox_SetShowHint(l._instance(), value)
}

func (l *TListBox) Sorted() bool {
    return ListBox_GetSorted(l._instance())
}

func (l *TListBox) SetSorted(value bool) {
    ListBox_SetSorted(l._instance(), value)
}

// TabOrder
//
// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (l *TListBox) TabOrder() TTabOrder {
    return ListBox_GetTabOrder(l._instance())
}

// SetTabOrder
//
// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (l *TListBox) SetTabOrder(value TTabOrder) {
    ListBox_SetTabOrder(l._instance(), value)
}

// TabStop
//
// 获取Tab可停留。
//
// Get Tab can stay.
func (l *TListBox) TabStop() bool {
    return ListBox_GetTabStop(l._instance())
}

// SetTabStop
//
// 设置Tab可停留。
//
// Set Tab can stay.
func (l *TListBox) SetTabStop(value bool) {
    ListBox_SetTabStop(l._instance(), value)
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (l *TListBox) Visible() bool {
    return ListBox_GetVisible(l._instance())
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (l *TListBox) SetVisible(value bool) {
    ListBox_SetVisible(l._instance(), value)
}

// SetOnClick
//
// 设置控件单击事件。
//
// Set control click event.
func (l *TListBox) SetOnClick(fn TNotifyEvent) {
    ListBox_SetOnClick(l._instance(), fn)
}

// SetOnContextPopup
//
// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (l *TListBox) SetOnContextPopup(fn TContextPopupEvent) {
    ListBox_SetOnContextPopup(l._instance(), fn)
}

// SetOnDblClick
//
// 设置双击事件。
func (l *TListBox) SetOnDblClick(fn TNotifyEvent) {
    ListBox_SetOnDblClick(l._instance(), fn)
}

// SetOnDragDrop
//
// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (l *TListBox) SetOnDragDrop(fn TDragDropEvent) {
    ListBox_SetOnDragDrop(l._instance(), fn)
}

// SetOnDragOver
//
// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (l *TListBox) SetOnDragOver(fn TDragOverEvent) {
    ListBox_SetOnDragOver(l._instance(), fn)
}

func (l *TListBox) SetOnDrawItem(fn TDrawItemEvent) {
    ListBox_SetOnDrawItem(l._instance(), fn)
}

// SetOnEndDrag
//
// 设置拖拽结束。
//
// Set End of drag.
func (l *TListBox) SetOnEndDrag(fn TEndDragEvent) {
    ListBox_SetOnEndDrag(l._instance(), fn)
}

// SetOnEnter
//
// 设置焦点进入。
//
// Set Focus entry.
func (l *TListBox) SetOnEnter(fn TNotifyEvent) {
    ListBox_SetOnEnter(l._instance(), fn)
}

// SetOnExit
//
// 设置焦点退出。
//
// Set Focus exit.
func (l *TListBox) SetOnExit(fn TNotifyEvent) {
    ListBox_SetOnExit(l._instance(), fn)
}

// SetOnKeyDown
//
// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (l *TListBox) SetOnKeyDown(fn TKeyEvent) {
    ListBox_SetOnKeyDown(l._instance(), fn)
}

// SetOnKeyPress
//
// 设置键键下事件。
func (l *TListBox) SetOnKeyPress(fn TKeyPressEvent) {
    ListBox_SetOnKeyPress(l._instance(), fn)
}

// SetOnKeyUp
//
// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (l *TListBox) SetOnKeyUp(fn TKeyEvent) {
    ListBox_SetOnKeyUp(l._instance(), fn)
}

func (l *TListBox) SetOnMeasureItem(fn TMeasureItemEvent) {
    ListBox_SetOnMeasureItem(l._instance(), fn)
}

// SetOnMouseDown
//
// 设置鼠标按下事件。
//
// Set Mouse down event.
func (l *TListBox) SetOnMouseDown(fn TMouseEvent) {
    ListBox_SetOnMouseDown(l._instance(), fn)
}

// SetOnMouseEnter
//
// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (l *TListBox) SetOnMouseEnter(fn TNotifyEvent) {
    ListBox_SetOnMouseEnter(l._instance(), fn)
}

// SetOnMouseLeave
//
// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (l *TListBox) SetOnMouseLeave(fn TNotifyEvent) {
    ListBox_SetOnMouseLeave(l._instance(), fn)
}

// SetOnMouseMove
//
// 设置鼠标移动事件。
func (l *TListBox) SetOnMouseMove(fn TMouseMoveEvent) {
    ListBox_SetOnMouseMove(l._instance(), fn)
}

// SetOnMouseUp
//
// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (l *TListBox) SetOnMouseUp(fn TMouseEvent) {
    ListBox_SetOnMouseUp(l._instance(), fn)
}

// Canvas
//
// 获取画布。
func (l *TListBox) Canvas() *TCanvas {
    return AsCanvas(ListBox_GetCanvas(l._instance()))
}

func (l *TListBox) Count() int32 {
    return ListBox_GetCount(l._instance())
}

func (l *TListBox) SelCount() int32 {
    return ListBox_GetSelCount(l._instance())
}

func (l *TListBox) ItemIndex() int32 {
    return ListBox_GetItemIndex(l._instance())
}

func (l *TListBox) SetItemIndex(value int32) {
    ListBox_SetItemIndex(l._instance(), value)
}

// DockClientCount
//
// 获取依靠客户端总数。
func (l *TListBox) DockClientCount() int32 {
    return ListBox_GetDockClientCount(l._instance())
}

// DockSite
//
// 获取停靠站点。
//
// Get Docking site.
func (l *TListBox) DockSite() bool {
    return ListBox_GetDockSite(l._instance())
}

// SetDockSite
//
// 设置停靠站点。
//
// Set Docking site.
func (l *TListBox) SetDockSite(value bool) {
    ListBox_SetDockSite(l._instance(), value)
}

// MouseInClient
//
// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (l *TListBox) MouseInClient() bool {
    return ListBox_GetMouseInClient(l._instance())
}

// VisibleDockClientCount
//
// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (l *TListBox) VisibleDockClientCount() int32 {
    return ListBox_GetVisibleDockClientCount(l._instance())
}

// Brush
//
// 获取画刷对象。
//
// Get Brush.
func (l *TListBox) Brush() *TBrush {
    return AsBrush(ListBox_GetBrush(l._instance()))
}

// ControlCount
//
// 获取子控件数。
//
// Get Number of child controls.
func (l *TListBox) ControlCount() int32 {
    return ListBox_GetControlCount(l._instance())
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (l *TListBox) Handle() HWND {
    return ListBox_GetHandle(l._instance())
}

// ParentWindow
//
// 获取父容器句柄。
//
// Get Parent container handle.
func (l *TListBox) ParentWindow() HWND {
    return ListBox_GetParentWindow(l._instance())
}

// SetParentWindow
//
// 设置父容器句柄。
//
// Set Parent container handle.
func (l *TListBox) SetParentWindow(value HWND) {
    ListBox_SetParentWindow(l._instance(), value)
}

func (l *TListBox) Showing() bool {
    return ListBox_GetShowing(l._instance())
}

// UseDockManager
//
// 获取使用停靠管理。
func (l *TListBox) UseDockManager() bool {
    return ListBox_GetUseDockManager(l._instance())
}

// SetUseDockManager
//
// 设置使用停靠管理。
func (l *TListBox) SetUseDockManager(value bool) {
    ListBox_SetUseDockManager(l._instance(), value)
}

func (l *TListBox) Action() *TAction {
    return AsAction(ListBox_GetAction(l._instance()))
}

func (l *TListBox) SetAction(value IComponent) {
    ListBox_SetAction(l._instance(), CheckPtr(value))
}

func (l *TListBox) BoundsRect() TRect {
    return ListBox_GetBoundsRect(l._instance())
}

func (l *TListBox) SetBoundsRect(value TRect) {
    ListBox_SetBoundsRect(l._instance(), value)
}

// ClientHeight
//
// 获取客户区高度。
//
// Get client height.
func (l *TListBox) ClientHeight() int32 {
    return ListBox_GetClientHeight(l._instance())
}

// SetClientHeight
//
// 设置客户区高度。
//
// Set client height.
func (l *TListBox) SetClientHeight(value int32) {
    ListBox_SetClientHeight(l._instance(), value)
}

func (l *TListBox) ClientOrigin() TPoint {
    return ListBox_GetClientOrigin(l._instance())
}

// ClientRect
//
// 获取客户区矩形。
//
// Get client rectangle.
func (l *TListBox) ClientRect() TRect {
    return ListBox_GetClientRect(l._instance())
}

// ClientWidth
//
// 获取客户区宽度。
//
// Get client width.
func (l *TListBox) ClientWidth() int32 {
    return ListBox_GetClientWidth(l._instance())
}

// SetClientWidth
//
// 设置客户区宽度。
//
// Set client width.
func (l *TListBox) SetClientWidth(value int32) {
    ListBox_SetClientWidth(l._instance(), value)
}

// ControlState
//
// 获取控件状态。
//
// Get control state.
func (l *TListBox) ControlState() TControlState {
    return ListBox_GetControlState(l._instance())
}

// SetControlState
//
// 设置控件状态。
//
// Set control state.
func (l *TListBox) SetControlState(value TControlState) {
    ListBox_SetControlState(l._instance(), value)
}

// ControlStyle
//
// 获取控件样式。
//
// Get control style.
func (l *TListBox) ControlStyle() TControlStyle {
    return ListBox_GetControlStyle(l._instance())
}

// SetControlStyle
//
// 设置控件样式。
//
// Set control style.
func (l *TListBox) SetControlStyle(value TControlStyle) {
    ListBox_SetControlStyle(l._instance(), value)
}

func (l *TListBox) Floating() bool {
    return ListBox_GetFloating(l._instance())
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (l *TListBox) Parent() *TWinControl {
    return AsWinControl(ListBox_GetParent(l._instance()))
}

// SetParent
//
// 设置控件父容器。
//
// Set control parent container.
func (l *TListBox) SetParent(value IWinControl) {
    ListBox_SetParent(l._instance(), CheckPtr(value))
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (l *TListBox) Left() int32 {
    return ListBox_GetLeft(l._instance())
}

// SetLeft
//
// 设置左边位置。
//
// Set Left position.
func (l *TListBox) SetLeft(value int32) {
    ListBox_SetLeft(l._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (l *TListBox) Top() int32 {
    return ListBox_GetTop(l._instance())
}

// SetTop
//
// 设置顶边位置。
//
// Set Top position.
func (l *TListBox) SetTop(value int32) {
    ListBox_SetTop(l._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (l *TListBox) Width() int32 {
    return ListBox_GetWidth(l._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (l *TListBox) SetWidth(value int32) {
    ListBox_SetWidth(l._instance(), value)
}

// Height
//
// 获取高度。
//
// Get height.
func (l *TListBox) Height() int32 {
    return ListBox_GetHeight(l._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (l *TListBox) SetHeight(value int32) {
    ListBox_SetHeight(l._instance(), value)
}

// Cursor
//
// 获取控件光标。
//
// Get control cursor.
func (l *TListBox) Cursor() TCursor {
    return ListBox_GetCursor(l._instance())
}

// SetCursor
//
// 设置控件光标。
//
// Set control cursor.
func (l *TListBox) SetCursor(value TCursor) {
    ListBox_SetCursor(l._instance(), value)
}

// Hint
//
// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (l *TListBox) Hint() string {
    return ListBox_GetHint(l._instance())
}

// SetHint
//
// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (l *TListBox) SetHint(value string) {
    ListBox_SetHint(l._instance(), value)
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (l *TListBox) ComponentCount() int32 {
    return ListBox_GetComponentCount(l._instance())
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (l *TListBox) ComponentIndex() int32 {
    return ListBox_GetComponentIndex(l._instance())
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (l *TListBox) SetComponentIndex(value int32) {
    ListBox_SetComponentIndex(l._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (l *TListBox) Owner() *TComponent {
    return AsComponent(ListBox_GetOwner(l._instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (l *TListBox) Name() string {
    return ListBox_GetName(l._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (l *TListBox) SetName(value string) {
    ListBox_SetName(l._instance(), value)
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (l *TListBox) Tag() int {
    return ListBox_GetTag(l._instance())
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (l *TListBox) SetTag(value int) {
    ListBox_SetTag(l._instance(), value)
}

// AnchorSideLeft
//
// 获取左边锚点。
func (l *TListBox) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(ListBox_GetAnchorSideLeft(l._instance()))
}

// SetAnchorSideLeft
//
// 设置左边锚点。
func (l *TListBox) SetAnchorSideLeft(value *TAnchorSide) {
    ListBox_SetAnchorSideLeft(l._instance(), CheckPtr(value))
}

// AnchorSideTop
//
// 获取顶边锚点。
func (l *TListBox) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(ListBox_GetAnchorSideTop(l._instance()))
}

// SetAnchorSideTop
//
// 设置顶边锚点。
func (l *TListBox) SetAnchorSideTop(value *TAnchorSide) {
    ListBox_SetAnchorSideTop(l._instance(), CheckPtr(value))
}

// AnchorSideRight
//
// 获取右边锚点。
func (l *TListBox) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(ListBox_GetAnchorSideRight(l._instance()))
}

// SetAnchorSideRight
//
// 设置右边锚点。
func (l *TListBox) SetAnchorSideRight(value *TAnchorSide) {
    ListBox_SetAnchorSideRight(l._instance(), CheckPtr(value))
}

// AnchorSideBottom
//
// 获取底边锚点。
func (l *TListBox) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(ListBox_GetAnchorSideBottom(l._instance()))
}

// SetAnchorSideBottom
//
// 设置底边锚点。
func (l *TListBox) SetAnchorSideBottom(value *TAnchorSide) {
    ListBox_SetAnchorSideBottom(l._instance(), CheckPtr(value))
}

func (l *TListBox) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(ListBox_GetChildSizing(l._instance()))
}

func (l *TListBox) SetChildSizing(value *TControlChildSizing) {
    ListBox_SetChildSizing(l._instance(), CheckPtr(value))
}

// BorderSpacing
//
// 获取边框间距。
func (l *TListBox) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(ListBox_GetBorderSpacing(l._instance()))
}

// SetBorderSpacing
//
// 设置边框间距。
func (l *TListBox) SetBorderSpacing(value *TControlBorderSpacing) {
    ListBox_SetBorderSpacing(l._instance(), CheckPtr(value))
}

func (l *TListBox) Selected(Index int32) bool {
    return ListBox_GetSelected(l._instance(), Index)
}

func (l *TListBox) SetSelected(Index int32, value bool) {
    ListBox_SetSelected(l._instance(), Index, value)
}

// DockClients
//
// 获取指定索引停靠客户端。
func (l *TListBox) DockClients(Index int32) *TControl {
    return AsControl(ListBox_GetDockClients(l._instance(), Index))
}

// Controls
//
// 获取指定索引子控件。
func (l *TListBox) Controls(Index int32) *TControl {
    return AsControl(ListBox_GetControls(l._instance(), Index))
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (l *TListBox) Components(AIndex int32) *TComponent {
    return AsComponent(ListBox_GetComponents(l._instance(), AIndex))
}

// AnchorSide
//
// 获取锚侧面。
func (l *TListBox) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(ListBox_GetAnchorSide(l._instance(), AKind))
}

