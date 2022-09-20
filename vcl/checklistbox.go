
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

type TCheckListBox struct {
    IWinControl
    instance unsafe.Pointer
}

// NewCheckListBox
//
// 创建一个新的对象。
// 
// Create a new object.
func NewCheckListBox(owner IComponent) *TCheckListBox {
    c := new(TCheckListBox)
    c.instance = unsafe.Pointer(CheckListBox_Create(CheckPtr(owner)))
    return c
}

// AsCheckListBox
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsCheckListBox(obj interface{}) *TCheckListBox {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TCheckListBox{instance: instance}
}

// Free 
//
// 释放对象。
// 
// Free object.
func (c *TCheckListBox) Free() {
    if c.instance != nullptr {
        CheckListBox_Free(c._instance())
        c.instance  = nullptr
    }
}

func (c *TCheckListBox) _instance() uintptr {
    return uintptr(c.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (c *TCheckListBox) Instance() uintptr {
    return c._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (c *TCheckListBox) UnsafeAddr() unsafe.Pointer {
    return c.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (c *TCheckListBox) IsValid() bool {
    return c.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (c *TCheckListBox) Is() TIs {
    return TIs(c._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (c *TCheckListBox) As() TAs {
//    return TAs(c._instance())
//}

// TCheckListBoxClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TCheckListBoxClass() TClass {
    return CheckListBox_StaticClassType()
}

func (c *TCheckListBox) CheckAll(AState TCheckBoxState, AllowGrayed bool, AllowDisabled bool) {
    CheckListBox_CheckAll(c._instance(), AState , AllowGrayed , AllowDisabled)
}

func (c *TCheckListBox) AddItem(Item string, AObject IObject) {
    CheckListBox_AddItem(c._instance(), Item , CheckPtr(AObject))
}

// Clear
//
// 清除。
func (c *TCheckListBox) Clear() {
    CheckListBox_Clear(c._instance())
}

// ClearSelection
//
// 清除选择。
func (c *TCheckListBox) ClearSelection() {
    CheckListBox_ClearSelection(c._instance())
}

// DeleteSelected
//
// 删除选择的。
func (c *TCheckListBox) DeleteSelected() {
    CheckListBox_DeleteSelected(c._instance())
}

func (c *TCheckListBox) ItemAtPos(Pos TPoint, Existing bool) int32 {
    return CheckListBox_ItemAtPos(c._instance(), Pos , Existing)
}

func (c *TCheckListBox) ItemRect(Index int32) TRect {
    return CheckListBox_ItemRect(c._instance(), Index)
}

// SelectAll
//
// 全选。
func (c *TCheckListBox) SelectAll() {
    CheckListBox_SelectAll(c._instance())
}

// CanFocus
//
// 是否可以获得焦点。
func (c *TCheckListBox) CanFocus() bool {
    return CheckListBox_CanFocus(c._instance())
}

// ContainsControl
//
// 返回是否包含指定控件。
//
// it's contain a specified control.
func (c *TCheckListBox) ContainsControl(Control IControl) bool {
    return CheckListBox_ContainsControl(c._instance(), CheckPtr(Control))
}

// ControlAtPos
//
// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (c *TCheckListBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(CheckListBox_ControlAtPos(c._instance(), Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
//
// 禁用控件的对齐。
//
// Disable control alignment.
func (c *TCheckListBox) DisableAlign() {
    CheckListBox_DisableAlign(c._instance())
}

// EnableAlign
//
// 启用控件对齐。
//
// Enabled control alignment.
func (c *TCheckListBox) EnableAlign() {
    CheckListBox_EnableAlign(c._instance())
}

// FindChildControl
//
// 查找子控件。
//
// Find sub controls.
func (c *TCheckListBox) FindChildControl(ControlName string) *TControl {
    return AsControl(CheckListBox_FindChildControl(c._instance(), ControlName))
}

func (c *TCheckListBox) FlipChildren(AllLevels bool) {
    CheckListBox_FlipChildren(c._instance(), AllLevels)
}

// Focused
//
// 返回是否获取焦点。
//
// Return to get focus.
func (c *TCheckListBox) Focused() bool {
    return CheckListBox_Focused(c._instance())
}

// HandleAllocated
//
// 句柄是否已经分配。
//
// Is the handle already allocated.
func (c *TCheckListBox) HandleAllocated() bool {
    return CheckListBox_HandleAllocated(c._instance())
}

// InsertControl
//
// 插入一个控件。
//
// Insert a control.
func (c *TCheckListBox) InsertControl(AControl IControl) {
    CheckListBox_InsertControl(c._instance(), CheckPtr(AControl))
}

// Invalidate
//
// 要求重绘。
//
// Redraw.
func (c *TCheckListBox) Invalidate() {
    CheckListBox_Invalidate(c._instance())
}

// PaintTo
//
// 绘画至指定DC。
//
// Painting to the specified DC.
func (c *TCheckListBox) PaintTo(DC HDC, X int32, Y int32) {
    CheckListBox_PaintTo(c._instance(), DC , X , Y)
}

// RemoveControl
//
// 移除一个控件。
//
// Remove a control.
func (c *TCheckListBox) RemoveControl(AControl IControl) {
    CheckListBox_RemoveControl(c._instance(), CheckPtr(AControl))
}

// Realign
//
// 重新对齐。
//
// Realign.
func (c *TCheckListBox) Realign() {
    CheckListBox_Realign(c._instance())
}

// Repaint
//
// 重绘。
//
// Repaint.
func (c *TCheckListBox) Repaint() {
    CheckListBox_Repaint(c._instance())
}

// ScaleBy
//
// 按比例缩放。
//
// Scale by.
func (c *TCheckListBox) ScaleBy(M int32, D int32) {
    CheckListBox_ScaleBy(c._instance(), M , D)
}

// ScrollBy
//
// 滚动至指定位置。
//
// Scroll by.
func (c *TCheckListBox) ScrollBy(DeltaX int32, DeltaY int32) {
    CheckListBox_ScrollBy(c._instance(), DeltaX , DeltaY)
}

// SetBounds
//
// 设置组件边界。
//
// Set component boundaries.
func (c *TCheckListBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CheckListBox_SetBounds(c._instance(), ALeft , ATop , AWidth , AHeight)
}

// SetFocus
//
// 设置控件焦点。
//
// Set control focus.
func (c *TCheckListBox) SetFocus() {
    CheckListBox_SetFocus(c._instance())
}

// Update
//
// 控件更新。
//
// Update.
func (c *TCheckListBox) Update() {
    CheckListBox_Update(c._instance())
}

// BringToFront
//
// 将控件置于最前。
//
// Bring the control to the front.
func (c *TCheckListBox) BringToFront() {
    CheckListBox_BringToFront(c._instance())
}

// ClientToScreen
//
// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (c *TCheckListBox) ClientToScreen(Point TPoint) TPoint {
    return CheckListBox_ClientToScreen(c._instance(), Point)
}

// ClientToParent
//
// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (c *TCheckListBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return CheckListBox_ClientToParent(c._instance(), Point , CheckPtr(AParent))
}

// Dragging
//
// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (c *TCheckListBox) Dragging() bool {
    return CheckListBox_Dragging(c._instance())
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (c *TCheckListBox) HasParent() bool {
    return CheckListBox_HasParent(c._instance())
}

// Hide
//
// 隐藏控件。
//
// Hidden control.
func (c *TCheckListBox) Hide() {
    CheckListBox_Hide(c._instance())
}

// Perform
//
// 发送一个消息。
//
// Send a message.
func (c *TCheckListBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CheckListBox_Perform(c._instance(), Msg , WParam , LParam)
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (c *TCheckListBox) Refresh() {
    CheckListBox_Refresh(c._instance())
}

// ScreenToClient
//
// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (c *TCheckListBox) ScreenToClient(Point TPoint) TPoint {
    return CheckListBox_ScreenToClient(c._instance(), Point)
}

// ParentToClient
//
// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (c *TCheckListBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return CheckListBox_ParentToClient(c._instance(), Point , CheckPtr(AParent))
}

// SendToBack
//
// 控件至于最后面。
//
// The control is placed at the end.
func (c *TCheckListBox) SendToBack() {
    CheckListBox_SendToBack(c._instance())
}

// Show
//
// 显示控件。
//
// Show control.
func (c *TCheckListBox) Show() {
    CheckListBox_Show(c._instance())
}

// GetTextBuf
//
// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (c *TCheckListBox) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return CheckListBox_GetTextBuf(c._instance(), Buffer , BufSize)
}

// GetTextLen
//
// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (c *TCheckListBox) GetTextLen() int32 {
    return CheckListBox_GetTextLen(c._instance())
}

// SetTextBuf
//
// 设置控件字符，如果有。
//
// Set control characters, if any.
func (c *TCheckListBox) SetTextBuf(Buffer string) {
    CheckListBox_SetTextBuf(c._instance(), Buffer)
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (c *TCheckListBox) FindComponent(AName string) *TComponent {
    return AsComponent(CheckListBox_FindComponent(c._instance(), AName))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (c *TCheckListBox) GetNamePath() string {
    return CheckListBox_GetNamePath(c._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TCheckListBox) Assign(Source IObject) {
    CheckListBox_Assign(c._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (c *TCheckListBox) ClassType() TClass {
    return CheckListBox_ClassType(c._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TCheckListBox) ClassName() string {
    return CheckListBox_ClassName(c._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TCheckListBox) InstanceSize() int32 {
    return CheckListBox_InstanceSize(c._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TCheckListBox) InheritsFrom(AClass TClass) bool {
    return CheckListBox_InheritsFrom(c._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (c *TCheckListBox) Equals(Obj IObject) bool {
    return CheckListBox_Equals(c._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TCheckListBox) GetHashCode() int32 {
    return CheckListBox_GetHashCode(c._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (c *TCheckListBox) ToString() string {
    return CheckListBox_ToString(c._instance())
}

func (c *TCheckListBox) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    CheckListBox_AnchorToNeighbour(c._instance(), ASide , ASpace , CheckPtr(ASibling))
}

func (c *TCheckListBox) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    CheckListBox_AnchorParallel(c._instance(), ASide , ASpace , CheckPtr(ASibling))
}

// AnchorHorizontalCenterTo
//
// 置于指定控件的横向中心。
func (c *TCheckListBox) AnchorHorizontalCenterTo(ASibling IControl) {
    CheckListBox_AnchorHorizontalCenterTo(c._instance(), CheckPtr(ASibling))
}

// AnchorVerticalCenterTo
//
// 置于指定控件的纵向中心。
func (c *TCheckListBox) AnchorVerticalCenterTo(ASibling IControl) {
    CheckListBox_AnchorVerticalCenterTo(c._instance(), CheckPtr(ASibling))
}

func (c *TCheckListBox) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    CheckListBox_AnchorSame(c._instance(), ASide , CheckPtr(ASibling))
}

func (c *TCheckListBox) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    CheckListBox_AnchorAsAlign(c._instance(), ATheAlign , ASpace)
}

func (c *TCheckListBox) AnchorClient(ASpace int32) {
    CheckListBox_AnchorClient(c._instance(), ASpace)
}

func (c *TCheckListBox) ScaleDesignToForm(ASize int32) int32 {
    return CheckListBox_ScaleDesignToForm(c._instance(), ASize)
}

func (c *TCheckListBox) ScaleFormToDesign(ASize int32) int32 {
    return CheckListBox_ScaleFormToDesign(c._instance(), ASize)
}

func (c *TCheckListBox) Scale96ToForm(ASize int32) int32 {
    return CheckListBox_Scale96ToForm(c._instance(), ASize)
}

func (c *TCheckListBox) ScaleFormTo96(ASize int32) int32 {
    return CheckListBox_ScaleFormTo96(c._instance(), ASize)
}

func (c *TCheckListBox) Scale96ToFont(ASize int32) int32 {
    return CheckListBox_Scale96ToFont(c._instance(), ASize)
}

func (c *TCheckListBox) ScaleFontTo96(ASize int32) int32 {
    return CheckListBox_ScaleFontTo96(c._instance(), ASize)
}

func (c *TCheckListBox) ScaleScreenToFont(ASize int32) int32 {
    return CheckListBox_ScaleScreenToFont(c._instance(), ASize)
}

func (c *TCheckListBox) ScaleFontToScreen(ASize int32) int32 {
    return CheckListBox_ScaleFontToScreen(c._instance(), ASize)
}

func (c *TCheckListBox) Scale96ToScreen(ASize int32) int32 {
    return CheckListBox_Scale96ToScreen(c._instance(), ASize)
}

func (c *TCheckListBox) ScaleScreenTo96(ASize int32) int32 {
    return CheckListBox_ScaleScreenTo96(c._instance(), ASize)
}

func (c *TCheckListBox) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    CheckListBox_AutoAdjustLayout(c._instance(), AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (c *TCheckListBox) FixDesignFontsPPI(ADesignTimePPI int32) {
    CheckListBox_FixDesignFontsPPI(c._instance(), ADesignTimePPI)
}

func (c *TCheckListBox) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    CheckListBox_ScaleFontsPPI(c._instance(), AToPPI , AProportion)
}

func (c *TCheckListBox) SetOnClickCheck(fn TNotifyEvent) {
    CheckListBox_SetOnClickCheck(c._instance(), fn)
}

// Align
//
// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (c *TCheckListBox) Align() TAlign {
    return CheckListBox_GetAlign(c._instance())
}

// SetAlign
//
// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (c *TCheckListBox) SetAlign(value TAlign) {
    CheckListBox_SetAlign(c._instance(), value)
}

func (c *TCheckListBox) AllowGrayed() bool {
    return CheckListBox_GetAllowGrayed(c._instance())
}

func (c *TCheckListBox) SetAllowGrayed(value bool) {
    CheckListBox_SetAllowGrayed(c._instance(), value)
}

// Anchors
//
// 获取四个角位置的锚点。
func (c *TCheckListBox) Anchors() TAnchors {
    return CheckListBox_GetAnchors(c._instance())
}

// SetAnchors
//
// 设置四个角位置的锚点。
func (c *TCheckListBox) SetAnchors(value TAnchors) {
    CheckListBox_SetAnchors(c._instance(), value)
}

func (c *TCheckListBox) BiDiMode() TBiDiMode {
    return CheckListBox_GetBiDiMode(c._instance())
}

func (c *TCheckListBox) SetBiDiMode(value TBiDiMode) {
    CheckListBox_SetBiDiMode(c._instance(), value)
}

// BorderStyle
//
// 获取窗口边框样式。比如：无边框，单一边框等。
func (c *TCheckListBox) BorderStyle() TBorderStyle {
    return CheckListBox_GetBorderStyle(c._instance())
}

// SetBorderStyle
//
// 设置窗口边框样式。比如：无边框，单一边框等。
func (c *TCheckListBox) SetBorderStyle(value TBorderStyle) {
    CheckListBox_SetBorderStyle(c._instance(), value)
}

// Color
//
// 获取颜色。
//
// Get color.
func (c *TCheckListBox) Color() TColor {
    return CheckListBox_GetColor(c._instance())
}

// SetColor
//
// 设置颜色。
//
// Set color.
func (c *TCheckListBox) SetColor(value TColor) {
    CheckListBox_SetColor(c._instance(), value)
}

func (c *TCheckListBox) Columns() int32 {
    return CheckListBox_GetColumns(c._instance())
}

func (c *TCheckListBox) SetColumns(value int32) {
    CheckListBox_SetColumns(c._instance(), value)
}

// Constraints
//
// 获取约束控件大小。
func (c *TCheckListBox) Constraints() *TSizeConstraints {
    return AsSizeConstraints(CheckListBox_GetConstraints(c._instance()))
}

// SetConstraints
//
// 设置约束控件大小。
func (c *TCheckListBox) SetConstraints(value *TSizeConstraints) {
    CheckListBox_SetConstraints(c._instance(), CheckPtr(value))
}

// DoubleBuffered
//
// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (c *TCheckListBox) DoubleBuffered() bool {
    return CheckListBox_GetDoubleBuffered(c._instance())
}

// SetDoubleBuffered
//
// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (c *TCheckListBox) SetDoubleBuffered(value bool) {
    CheckListBox_SetDoubleBuffered(c._instance(), value)
}

// DragCursor
//
// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (c *TCheckListBox) DragCursor() TCursor {
    return CheckListBox_GetDragCursor(c._instance())
}

// SetDragCursor
//
// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (c *TCheckListBox) SetDragCursor(value TCursor) {
    CheckListBox_SetDragCursor(c._instance(), value)
}

// DragMode
//
// 获取拖拽模式。
//
// Get Drag mode.
func (c *TCheckListBox) DragMode() TDragMode {
    return CheckListBox_GetDragMode(c._instance())
}

// SetDragMode
//
// 设置拖拽模式。
//
// Set Drag mode.
func (c *TCheckListBox) SetDragMode(value TDragMode) {
    CheckListBox_SetDragMode(c._instance(), value)
}

// Enabled
//
// 获取控件启用。
//
// Get the control enabled.
func (c *TCheckListBox) Enabled() bool {
    return CheckListBox_GetEnabled(c._instance())
}

// SetEnabled
//
// 设置控件启用。
//
// Set the control enabled.
func (c *TCheckListBox) SetEnabled(value bool) {
    CheckListBox_SetEnabled(c._instance(), value)
}

// Font
//
// 获取字体。
//
// Get Font.
func (c *TCheckListBox) Font() *TFont {
    return AsFont(CheckListBox_GetFont(c._instance()))
}

// SetFont
//
// 设置字体。
//
// Set Font.
func (c *TCheckListBox) SetFont(value *TFont) {
    CheckListBox_SetFont(c._instance(), CheckPtr(value))
}

func (c *TCheckListBox) ItemHeight() int32 {
    return CheckListBox_GetItemHeight(c._instance())
}

func (c *TCheckListBox) SetItemHeight(value int32) {
    CheckListBox_SetItemHeight(c._instance(), value)
}

func (c *TCheckListBox) Items() *TStrings {
    return AsStrings(CheckListBox_GetItems(c._instance()))
}

func (c *TCheckListBox) SetItems(value IStrings) {
    CheckListBox_SetItems(c._instance(), CheckPtr(value))
}

// ParentColor
//
// 获取使用父容器颜色。
//
// Get parent color.
func (c *TCheckListBox) ParentColor() bool {
    return CheckListBox_GetParentColor(c._instance())
}

// SetParentColor
//
// 设置使用父容器颜色。
//
// Set parent color.
func (c *TCheckListBox) SetParentColor(value bool) {
    CheckListBox_SetParentColor(c._instance(), value)
}

// ParentDoubleBuffered
//
// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (c *TCheckListBox) ParentDoubleBuffered() bool {
    return CheckListBox_GetParentDoubleBuffered(c._instance())
}

// SetParentDoubleBuffered
//
// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (c *TCheckListBox) SetParentDoubleBuffered(value bool) {
    CheckListBox_SetParentDoubleBuffered(c._instance(), value)
}

// ParentFont
//
// 获取使用父容器字体。
//
// Get Parent container font.
func (c *TCheckListBox) ParentFont() bool {
    return CheckListBox_GetParentFont(c._instance())
}

// SetParentFont
//
// 设置使用父容器字体。
//
// Set Parent container font.
func (c *TCheckListBox) SetParentFont(value bool) {
    CheckListBox_SetParentFont(c._instance(), value)
}

// ParentShowHint
//
// 获取以父容器的ShowHint属性为准。
func (c *TCheckListBox) ParentShowHint() bool {
    return CheckListBox_GetParentShowHint(c._instance())
}

// SetParentShowHint
//
// 设置以父容器的ShowHint属性为准。
func (c *TCheckListBox) SetParentShowHint(value bool) {
    CheckListBox_SetParentShowHint(c._instance(), value)
}

// PopupMenu
//
// 获取右键菜单。
//
// Get Right click menu.
func (c *TCheckListBox) PopupMenu() *TPopupMenu {
    return AsPopupMenu(CheckListBox_GetPopupMenu(c._instance()))
}

// SetPopupMenu
//
// 设置右键菜单。
//
// Set Right click menu.
func (c *TCheckListBox) SetPopupMenu(value IComponent) {
    CheckListBox_SetPopupMenu(c._instance(), CheckPtr(value))
}

// ShowHint
//
// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (c *TCheckListBox) ShowHint() bool {
    return CheckListBox_GetShowHint(c._instance())
}

// SetShowHint
//
// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (c *TCheckListBox) SetShowHint(value bool) {
    CheckListBox_SetShowHint(c._instance(), value)
}

func (c *TCheckListBox) Sorted() bool {
    return CheckListBox_GetSorted(c._instance())
}

func (c *TCheckListBox) SetSorted(value bool) {
    CheckListBox_SetSorted(c._instance(), value)
}

func (c *TCheckListBox) Style() TListBoxStyle {
    return CheckListBox_GetStyle(c._instance())
}

func (c *TCheckListBox) SetStyle(value TListBoxStyle) {
    CheckListBox_SetStyle(c._instance(), value)
}

// TabOrder
//
// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (c *TCheckListBox) TabOrder() TTabOrder {
    return CheckListBox_GetTabOrder(c._instance())
}

// SetTabOrder
//
// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (c *TCheckListBox) SetTabOrder(value TTabOrder) {
    CheckListBox_SetTabOrder(c._instance(), value)
}

// TabStop
//
// 获取Tab可停留。
//
// Get Tab can stay.
func (c *TCheckListBox) TabStop() bool {
    return CheckListBox_GetTabStop(c._instance())
}

// SetTabStop
//
// 设置Tab可停留。
//
// Set Tab can stay.
func (c *TCheckListBox) SetTabStop(value bool) {
    CheckListBox_SetTabStop(c._instance(), value)
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (c *TCheckListBox) Visible() bool {
    return CheckListBox_GetVisible(c._instance())
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (c *TCheckListBox) SetVisible(value bool) {
    CheckListBox_SetVisible(c._instance(), value)
}

// SetOnClick
//
// 设置控件单击事件。
//
// Set control click event.
func (c *TCheckListBox) SetOnClick(fn TNotifyEvent) {
    CheckListBox_SetOnClick(c._instance(), fn)
}

// SetOnContextPopup
//
// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (c *TCheckListBox) SetOnContextPopup(fn TContextPopupEvent) {
    CheckListBox_SetOnContextPopup(c._instance(), fn)
}

// SetOnDblClick
//
// 设置双击事件。
func (c *TCheckListBox) SetOnDblClick(fn TNotifyEvent) {
    CheckListBox_SetOnDblClick(c._instance(), fn)
}

// SetOnDragDrop
//
// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (c *TCheckListBox) SetOnDragDrop(fn TDragDropEvent) {
    CheckListBox_SetOnDragDrop(c._instance(), fn)
}

// SetOnDragOver
//
// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (c *TCheckListBox) SetOnDragOver(fn TDragOverEvent) {
    CheckListBox_SetOnDragOver(c._instance(), fn)
}

// SetOnEndDrag
//
// 设置拖拽结束。
//
// Set End of drag.
func (c *TCheckListBox) SetOnEndDrag(fn TEndDragEvent) {
    CheckListBox_SetOnEndDrag(c._instance(), fn)
}

// SetOnEnter
//
// 设置焦点进入。
//
// Set Focus entry.
func (c *TCheckListBox) SetOnEnter(fn TNotifyEvent) {
    CheckListBox_SetOnEnter(c._instance(), fn)
}

// SetOnExit
//
// 设置焦点退出。
//
// Set Focus exit.
func (c *TCheckListBox) SetOnExit(fn TNotifyEvent) {
    CheckListBox_SetOnExit(c._instance(), fn)
}

// SetOnKeyDown
//
// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (c *TCheckListBox) SetOnKeyDown(fn TKeyEvent) {
    CheckListBox_SetOnKeyDown(c._instance(), fn)
}

// SetOnKeyPress
//
// 设置键键下事件。
func (c *TCheckListBox) SetOnKeyPress(fn TKeyPressEvent) {
    CheckListBox_SetOnKeyPress(c._instance(), fn)
}

// SetOnKeyUp
//
// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (c *TCheckListBox) SetOnKeyUp(fn TKeyEvent) {
    CheckListBox_SetOnKeyUp(c._instance(), fn)
}

func (c *TCheckListBox) SetOnMeasureItem(fn TMeasureItemEvent) {
    CheckListBox_SetOnMeasureItem(c._instance(), fn)
}

// SetOnMouseDown
//
// 设置鼠标按下事件。
//
// Set Mouse down event.
func (c *TCheckListBox) SetOnMouseDown(fn TMouseEvent) {
    CheckListBox_SetOnMouseDown(c._instance(), fn)
}

// SetOnMouseEnter
//
// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (c *TCheckListBox) SetOnMouseEnter(fn TNotifyEvent) {
    CheckListBox_SetOnMouseEnter(c._instance(), fn)
}

// SetOnMouseLeave
//
// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (c *TCheckListBox) SetOnMouseLeave(fn TNotifyEvent) {
    CheckListBox_SetOnMouseLeave(c._instance(), fn)
}

// SetOnMouseMove
//
// 设置鼠标移动事件。
func (c *TCheckListBox) SetOnMouseMove(fn TMouseMoveEvent) {
    CheckListBox_SetOnMouseMove(c._instance(), fn)
}

// SetOnMouseUp
//
// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (c *TCheckListBox) SetOnMouseUp(fn TMouseEvent) {
    CheckListBox_SetOnMouseUp(c._instance(), fn)
}

// Canvas
//
// 获取画布。
func (c *TCheckListBox) Canvas() *TCanvas {
    return AsCanvas(CheckListBox_GetCanvas(c._instance()))
}

func (c *TCheckListBox) Count() int32 {
    return CheckListBox_GetCount(c._instance())
}

func (c *TCheckListBox) TopIndex() int32 {
    return CheckListBox_GetTopIndex(c._instance())
}

func (c *TCheckListBox) SetTopIndex(value int32) {
    CheckListBox_SetTopIndex(c._instance(), value)
}

func (c *TCheckListBox) MultiSelect() bool {
    return CheckListBox_GetMultiSelect(c._instance())
}

func (c *TCheckListBox) SetMultiSelect(value bool) {
    CheckListBox_SetMultiSelect(c._instance(), value)
}

func (c *TCheckListBox) SelCount() int32 {
    return CheckListBox_GetSelCount(c._instance())
}

func (c *TCheckListBox) ItemIndex() int32 {
    return CheckListBox_GetItemIndex(c._instance())
}

func (c *TCheckListBox) SetItemIndex(value int32) {
    CheckListBox_SetItemIndex(c._instance(), value)
}

// DockClientCount
//
// 获取依靠客户端总数。
func (c *TCheckListBox) DockClientCount() int32 {
    return CheckListBox_GetDockClientCount(c._instance())
}

// DockSite
//
// 获取停靠站点。
//
// Get Docking site.
func (c *TCheckListBox) DockSite() bool {
    return CheckListBox_GetDockSite(c._instance())
}

// SetDockSite
//
// 设置停靠站点。
//
// Set Docking site.
func (c *TCheckListBox) SetDockSite(value bool) {
    CheckListBox_SetDockSite(c._instance(), value)
}

// MouseInClient
//
// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (c *TCheckListBox) MouseInClient() bool {
    return CheckListBox_GetMouseInClient(c._instance())
}

// VisibleDockClientCount
//
// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (c *TCheckListBox) VisibleDockClientCount() int32 {
    return CheckListBox_GetVisibleDockClientCount(c._instance())
}

// Brush
//
// 获取画刷对象。
//
// Get Brush.
func (c *TCheckListBox) Brush() *TBrush {
    return AsBrush(CheckListBox_GetBrush(c._instance()))
}

// ControlCount
//
// 获取子控件数。
//
// Get Number of child controls.
func (c *TCheckListBox) ControlCount() int32 {
    return CheckListBox_GetControlCount(c._instance())
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (c *TCheckListBox) Handle() HWND {
    return CheckListBox_GetHandle(c._instance())
}

// ParentWindow
//
// 获取父容器句柄。
//
// Get Parent container handle.
func (c *TCheckListBox) ParentWindow() HWND {
    return CheckListBox_GetParentWindow(c._instance())
}

// SetParentWindow
//
// 设置父容器句柄。
//
// Set Parent container handle.
func (c *TCheckListBox) SetParentWindow(value HWND) {
    CheckListBox_SetParentWindow(c._instance(), value)
}

func (c *TCheckListBox) Showing() bool {
    return CheckListBox_GetShowing(c._instance())
}

// UseDockManager
//
// 获取使用停靠管理。
func (c *TCheckListBox) UseDockManager() bool {
    return CheckListBox_GetUseDockManager(c._instance())
}

// SetUseDockManager
//
// 设置使用停靠管理。
func (c *TCheckListBox) SetUseDockManager(value bool) {
    CheckListBox_SetUseDockManager(c._instance(), value)
}

func (c *TCheckListBox) Action() *TAction {
    return AsAction(CheckListBox_GetAction(c._instance()))
}

func (c *TCheckListBox) SetAction(value IComponent) {
    CheckListBox_SetAction(c._instance(), CheckPtr(value))
}

func (c *TCheckListBox) BoundsRect() TRect {
    return CheckListBox_GetBoundsRect(c._instance())
}

func (c *TCheckListBox) SetBoundsRect(value TRect) {
    CheckListBox_SetBoundsRect(c._instance(), value)
}

// ClientHeight
//
// 获取客户区高度。
//
// Get client height.
func (c *TCheckListBox) ClientHeight() int32 {
    return CheckListBox_GetClientHeight(c._instance())
}

// SetClientHeight
//
// 设置客户区高度。
//
// Set client height.
func (c *TCheckListBox) SetClientHeight(value int32) {
    CheckListBox_SetClientHeight(c._instance(), value)
}

func (c *TCheckListBox) ClientOrigin() TPoint {
    return CheckListBox_GetClientOrigin(c._instance())
}

// ClientRect
//
// 获取客户区矩形。
//
// Get client rectangle.
func (c *TCheckListBox) ClientRect() TRect {
    return CheckListBox_GetClientRect(c._instance())
}

// ClientWidth
//
// 获取客户区宽度。
//
// Get client width.
func (c *TCheckListBox) ClientWidth() int32 {
    return CheckListBox_GetClientWidth(c._instance())
}

// SetClientWidth
//
// 设置客户区宽度。
//
// Set client width.
func (c *TCheckListBox) SetClientWidth(value int32) {
    CheckListBox_SetClientWidth(c._instance(), value)
}

// ControlState
//
// 获取控件状态。
//
// Get control state.
func (c *TCheckListBox) ControlState() TControlState {
    return CheckListBox_GetControlState(c._instance())
}

// SetControlState
//
// 设置控件状态。
//
// Set control state.
func (c *TCheckListBox) SetControlState(value TControlState) {
    CheckListBox_SetControlState(c._instance(), value)
}

// ControlStyle
//
// 获取控件样式。
//
// Get control style.
func (c *TCheckListBox) ControlStyle() TControlStyle {
    return CheckListBox_GetControlStyle(c._instance())
}

// SetControlStyle
//
// 设置控件样式。
//
// Set control style.
func (c *TCheckListBox) SetControlStyle(value TControlStyle) {
    CheckListBox_SetControlStyle(c._instance(), value)
}

func (c *TCheckListBox) Floating() bool {
    return CheckListBox_GetFloating(c._instance())
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (c *TCheckListBox) Parent() *TWinControl {
    return AsWinControl(CheckListBox_GetParent(c._instance()))
}

// SetParent
//
// 设置控件父容器。
//
// Set control parent container.
func (c *TCheckListBox) SetParent(value IWinControl) {
    CheckListBox_SetParent(c._instance(), CheckPtr(value))
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (c *TCheckListBox) Left() int32 {
    return CheckListBox_GetLeft(c._instance())
}

// SetLeft
//
// 设置左边位置。
//
// Set Left position.
func (c *TCheckListBox) SetLeft(value int32) {
    CheckListBox_SetLeft(c._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (c *TCheckListBox) Top() int32 {
    return CheckListBox_GetTop(c._instance())
}

// SetTop
//
// 设置顶边位置。
//
// Set Top position.
func (c *TCheckListBox) SetTop(value int32) {
    CheckListBox_SetTop(c._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (c *TCheckListBox) Width() int32 {
    return CheckListBox_GetWidth(c._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (c *TCheckListBox) SetWidth(value int32) {
    CheckListBox_SetWidth(c._instance(), value)
}

// Height
//
// 获取高度。
//
// Get height.
func (c *TCheckListBox) Height() int32 {
    return CheckListBox_GetHeight(c._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (c *TCheckListBox) SetHeight(value int32) {
    CheckListBox_SetHeight(c._instance(), value)
}

// Cursor
//
// 获取控件光标。
//
// Get control cursor.
func (c *TCheckListBox) Cursor() TCursor {
    return CheckListBox_GetCursor(c._instance())
}

// SetCursor
//
// 设置控件光标。
//
// Set control cursor.
func (c *TCheckListBox) SetCursor(value TCursor) {
    CheckListBox_SetCursor(c._instance(), value)
}

// Hint
//
// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (c *TCheckListBox) Hint() string {
    return CheckListBox_GetHint(c._instance())
}

// SetHint
//
// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (c *TCheckListBox) SetHint(value string) {
    CheckListBox_SetHint(c._instance(), value)
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (c *TCheckListBox) ComponentCount() int32 {
    return CheckListBox_GetComponentCount(c._instance())
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (c *TCheckListBox) ComponentIndex() int32 {
    return CheckListBox_GetComponentIndex(c._instance())
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (c *TCheckListBox) SetComponentIndex(value int32) {
    CheckListBox_SetComponentIndex(c._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (c *TCheckListBox) Owner() *TComponent {
    return AsComponent(CheckListBox_GetOwner(c._instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (c *TCheckListBox) Name() string {
    return CheckListBox_GetName(c._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (c *TCheckListBox) SetName(value string) {
    CheckListBox_SetName(c._instance(), value)
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (c *TCheckListBox) Tag() int {
    return CheckListBox_GetTag(c._instance())
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (c *TCheckListBox) SetTag(value int) {
    CheckListBox_SetTag(c._instance(), value)
}

// AnchorSideLeft
//
// 获取左边锚点。
func (c *TCheckListBox) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(CheckListBox_GetAnchorSideLeft(c._instance()))
}

// SetAnchorSideLeft
//
// 设置左边锚点。
func (c *TCheckListBox) SetAnchorSideLeft(value *TAnchorSide) {
    CheckListBox_SetAnchorSideLeft(c._instance(), CheckPtr(value))
}

// AnchorSideTop
//
// 获取顶边锚点。
func (c *TCheckListBox) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(CheckListBox_GetAnchorSideTop(c._instance()))
}

// SetAnchorSideTop
//
// 设置顶边锚点。
func (c *TCheckListBox) SetAnchorSideTop(value *TAnchorSide) {
    CheckListBox_SetAnchorSideTop(c._instance(), CheckPtr(value))
}

// AnchorSideRight
//
// 获取右边锚点。
func (c *TCheckListBox) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(CheckListBox_GetAnchorSideRight(c._instance()))
}

// SetAnchorSideRight
//
// 设置右边锚点。
func (c *TCheckListBox) SetAnchorSideRight(value *TAnchorSide) {
    CheckListBox_SetAnchorSideRight(c._instance(), CheckPtr(value))
}

// AnchorSideBottom
//
// 获取底边锚点。
func (c *TCheckListBox) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(CheckListBox_GetAnchorSideBottom(c._instance()))
}

// SetAnchorSideBottom
//
// 设置底边锚点。
func (c *TCheckListBox) SetAnchorSideBottom(value *TAnchorSide) {
    CheckListBox_SetAnchorSideBottom(c._instance(), CheckPtr(value))
}

func (c *TCheckListBox) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(CheckListBox_GetChildSizing(c._instance()))
}

func (c *TCheckListBox) SetChildSizing(value *TControlChildSizing) {
    CheckListBox_SetChildSizing(c._instance(), CheckPtr(value))
}

// BorderSpacing
//
// 获取边框间距。
func (c *TCheckListBox) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(CheckListBox_GetBorderSpacing(c._instance()))
}

// SetBorderSpacing
//
// 设置边框间距。
func (c *TCheckListBox) SetBorderSpacing(value *TControlBorderSpacing) {
    CheckListBox_SetBorderSpacing(c._instance(), CheckPtr(value))
}

// Checked
//
// 获取是否选中。
func (c *TCheckListBox) Checked(Index int32) bool {
    return CheckListBox_GetChecked(c._instance(), Index)
}

// SetChecked
//
// 设置是否选中。
func (c *TCheckListBox) SetChecked(Index int32, value bool) {
    CheckListBox_SetChecked(c._instance(), Index, value)
}

func (c *TCheckListBox) ItemEnabled(Index int32) bool {
    return CheckListBox_GetItemEnabled(c._instance(), Index)
}

func (c *TCheckListBox) SetItemEnabled(Index int32, value bool) {
    CheckListBox_SetItemEnabled(c._instance(), Index, value)
}

func (c *TCheckListBox) State(Index int32) TCheckBoxState {
    return CheckListBox_GetState(c._instance(), Index)
}

func (c *TCheckListBox) SetState(Index int32, value TCheckBoxState) {
    CheckListBox_SetState(c._instance(), Index, value)
}

func (c *TCheckListBox) Header(Index int32) bool {
    return CheckListBox_GetHeader(c._instance(), Index)
}

func (c *TCheckListBox) SetHeader(Index int32, value bool) {
    CheckListBox_SetHeader(c._instance(), Index, value)
}

func (c *TCheckListBox) Selected(Index int32) bool {
    return CheckListBox_GetSelected(c._instance(), Index)
}

func (c *TCheckListBox) SetSelected(Index int32, value bool) {
    CheckListBox_SetSelected(c._instance(), Index, value)
}

// DockClients
//
// 获取指定索引停靠客户端。
func (c *TCheckListBox) DockClients(Index int32) *TControl {
    return AsControl(CheckListBox_GetDockClients(c._instance(), Index))
}

// Controls
//
// 获取指定索引子控件。
func (c *TCheckListBox) Controls(Index int32) *TControl {
    return AsControl(CheckListBox_GetControls(c._instance(), Index))
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (c *TCheckListBox) Components(AIndex int32) *TComponent {
    return AsComponent(CheckListBox_GetComponents(c._instance(), AIndex))
}

// AnchorSide
//
// 获取锚侧面。
func (c *TCheckListBox) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(CheckListBox_GetAnchorSide(c._instance(), AKind))
}

