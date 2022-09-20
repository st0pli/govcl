
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

type TRichEdit struct {
    IWinControl
    instance unsafe.Pointer
}

// NewRichEdit
//
// 创建一个新的对象。
// 
// Create a new object.
func NewRichEdit(owner IComponent) *TRichEdit {
    r := new(TRichEdit)
    r.instance = unsafe.Pointer(RichEdit_Create(CheckPtr(owner)))
    return r
}

// AsRichEdit
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsRichEdit(obj interface{}) *TRichEdit {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TRichEdit{instance: instance}
}

// Free 
//
// 释放对象。
// 
// Free object.
func (r *TRichEdit) Free() {
    if r.instance != nullptr {
        RichEdit_Free(r._instance())
        r.instance  = nullptr
    }
}

func (r *TRichEdit) _instance() uintptr {
    return uintptr(r.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (r *TRichEdit) Instance() uintptr {
    return r._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (r *TRichEdit) UnsafeAddr() unsafe.Pointer {
    return r.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (r *TRichEdit) IsValid() bool {
    return r.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (r *TRichEdit) Is() TIs {
    return TIs(r._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (r *TRichEdit) As() TAs {
//    return TAs(r._instance())
//}

// TRichEditClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TRichEditClass() TClass {
    return RichEdit_StaticClassType()
}

// Clear
//
// 清除。
func (r *TRichEdit) Clear() {
    RichEdit_Clear(r._instance())
}

func (r *TRichEdit) FindText(SearchStr string, StartPos int32, Length int32, Options TSearchTypes) int32 {
    return RichEdit_FindText(r._instance(), SearchStr , StartPos , Length , Options)
}

// ClearSelection
//
// 清除选择。
func (r *TRichEdit) ClearSelection() {
    RichEdit_ClearSelection(r._instance())
}

// CopyToClipboard
//
// 复制到粘贴板。
func (r *TRichEdit) CopyToClipboard() {
    RichEdit_CopyToClipboard(r._instance())
}

// CutToClipboard
//
// 剪切到粘贴板。
func (r *TRichEdit) CutToClipboard() {
    RichEdit_CutToClipboard(r._instance())
}

// PasteFromClipboard
//
// 从剪切板粘贴。
func (r *TRichEdit) PasteFromClipboard() {
    RichEdit_PasteFromClipboard(r._instance())
}

// Undo
//
// 撤销上一次操作。
func (r *TRichEdit) Undo() {
    RichEdit_Undo(r._instance())
}

// SelectAll
//
// 全选。
func (r *TRichEdit) SelectAll() {
    RichEdit_SelectAll(r._instance())
}

// CanFocus
//
// 是否可以获得焦点。
func (r *TRichEdit) CanFocus() bool {
    return RichEdit_CanFocus(r._instance())
}

// ContainsControl
//
// 返回是否包含指定控件。
//
// it's contain a specified control.
func (r *TRichEdit) ContainsControl(Control IControl) bool {
    return RichEdit_ContainsControl(r._instance(), CheckPtr(Control))
}

// ControlAtPos
//
// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (r *TRichEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(RichEdit_ControlAtPos(r._instance(), Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
//
// 禁用控件的对齐。
//
// Disable control alignment.
func (r *TRichEdit) DisableAlign() {
    RichEdit_DisableAlign(r._instance())
}

// EnableAlign
//
// 启用控件对齐。
//
// Enabled control alignment.
func (r *TRichEdit) EnableAlign() {
    RichEdit_EnableAlign(r._instance())
}

// FindChildControl
//
// 查找子控件。
//
// Find sub controls.
func (r *TRichEdit) FindChildControl(ControlName string) *TControl {
    return AsControl(RichEdit_FindChildControl(r._instance(), ControlName))
}

func (r *TRichEdit) FlipChildren(AllLevels bool) {
    RichEdit_FlipChildren(r._instance(), AllLevels)
}

// Focused
//
// 返回是否获取焦点。
//
// Return to get focus.
func (r *TRichEdit) Focused() bool {
    return RichEdit_Focused(r._instance())
}

// HandleAllocated
//
// 句柄是否已经分配。
//
// Is the handle already allocated.
func (r *TRichEdit) HandleAllocated() bool {
    return RichEdit_HandleAllocated(r._instance())
}

// InsertControl
//
// 插入一个控件。
//
// Insert a control.
func (r *TRichEdit) InsertControl(AControl IControl) {
    RichEdit_InsertControl(r._instance(), CheckPtr(AControl))
}

// Invalidate
//
// 要求重绘。
//
// Redraw.
func (r *TRichEdit) Invalidate() {
    RichEdit_Invalidate(r._instance())
}

// PaintTo
//
// 绘画至指定DC。
//
// Painting to the specified DC.
func (r *TRichEdit) PaintTo(DC HDC, X int32, Y int32) {
    RichEdit_PaintTo(r._instance(), DC , X , Y)
}

// RemoveControl
//
// 移除一个控件。
//
// Remove a control.
func (r *TRichEdit) RemoveControl(AControl IControl) {
    RichEdit_RemoveControl(r._instance(), CheckPtr(AControl))
}

// Realign
//
// 重新对齐。
//
// Realign.
func (r *TRichEdit) Realign() {
    RichEdit_Realign(r._instance())
}

// Repaint
//
// 重绘。
//
// Repaint.
func (r *TRichEdit) Repaint() {
    RichEdit_Repaint(r._instance())
}

// ScaleBy
//
// 按比例缩放。
//
// Scale by.
func (r *TRichEdit) ScaleBy(M int32, D int32) {
    RichEdit_ScaleBy(r._instance(), M , D)
}

// ScrollBy
//
// 滚动至指定位置。
//
// Scroll by.
func (r *TRichEdit) ScrollBy(DeltaX int32, DeltaY int32) {
    RichEdit_ScrollBy(r._instance(), DeltaX , DeltaY)
}

// SetBounds
//
// 设置组件边界。
//
// Set component boundaries.
func (r *TRichEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    RichEdit_SetBounds(r._instance(), ALeft , ATop , AWidth , AHeight)
}

// SetFocus
//
// 设置控件焦点。
//
// Set control focus.
func (r *TRichEdit) SetFocus() {
    RichEdit_SetFocus(r._instance())
}

// Update
//
// 控件更新。
//
// Update.
func (r *TRichEdit) Update() {
    RichEdit_Update(r._instance())
}

// BringToFront
//
// 将控件置于最前。
//
// Bring the control to the front.
func (r *TRichEdit) BringToFront() {
    RichEdit_BringToFront(r._instance())
}

// ClientToScreen
//
// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (r *TRichEdit) ClientToScreen(Point TPoint) TPoint {
    return RichEdit_ClientToScreen(r._instance(), Point)
}

// ClientToParent
//
// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (r *TRichEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return RichEdit_ClientToParent(r._instance(), Point , CheckPtr(AParent))
}

// Dragging
//
// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (r *TRichEdit) Dragging() bool {
    return RichEdit_Dragging(r._instance())
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (r *TRichEdit) HasParent() bool {
    return RichEdit_HasParent(r._instance())
}

// Hide
//
// 隐藏控件。
//
// Hidden control.
func (r *TRichEdit) Hide() {
    RichEdit_Hide(r._instance())
}

// Perform
//
// 发送一个消息。
//
// Send a message.
func (r *TRichEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return RichEdit_Perform(r._instance(), Msg , WParam , LParam)
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (r *TRichEdit) Refresh() {
    RichEdit_Refresh(r._instance())
}

// ScreenToClient
//
// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (r *TRichEdit) ScreenToClient(Point TPoint) TPoint {
    return RichEdit_ScreenToClient(r._instance(), Point)
}

// ParentToClient
//
// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (r *TRichEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return RichEdit_ParentToClient(r._instance(), Point , CheckPtr(AParent))
}

// SendToBack
//
// 控件至于最后面。
//
// The control is placed at the end.
func (r *TRichEdit) SendToBack() {
    RichEdit_SendToBack(r._instance())
}

// Show
//
// 显示控件。
//
// Show control.
func (r *TRichEdit) Show() {
    RichEdit_Show(r._instance())
}

// GetTextBuf
//
// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (r *TRichEdit) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return RichEdit_GetTextBuf(r._instance(), Buffer , BufSize)
}

// GetTextLen
//
// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (r *TRichEdit) GetTextLen() int32 {
    return RichEdit_GetTextLen(r._instance())
}

// SetTextBuf
//
// 设置控件字符，如果有。
//
// Set control characters, if any.
func (r *TRichEdit) SetTextBuf(Buffer string) {
    RichEdit_SetTextBuf(r._instance(), Buffer)
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (r *TRichEdit) FindComponent(AName string) *TComponent {
    return AsComponent(RichEdit_FindComponent(r._instance(), AName))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (r *TRichEdit) GetNamePath() string {
    return RichEdit_GetNamePath(r._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (r *TRichEdit) Assign(Source IObject) {
    RichEdit_Assign(r._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (r *TRichEdit) ClassType() TClass {
    return RichEdit_ClassType(r._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (r *TRichEdit) ClassName() string {
    return RichEdit_ClassName(r._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (r *TRichEdit) InstanceSize() int32 {
    return RichEdit_InstanceSize(r._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (r *TRichEdit) InheritsFrom(AClass TClass) bool {
    return RichEdit_InheritsFrom(r._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (r *TRichEdit) Equals(Obj IObject) bool {
    return RichEdit_Equals(r._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (r *TRichEdit) GetHashCode() int32 {
    return RichEdit_GetHashCode(r._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (r *TRichEdit) ToString() string {
    return RichEdit_ToString(r._instance())
}

func (r *TRichEdit) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    RichEdit_AnchorToNeighbour(r._instance(), ASide , ASpace , CheckPtr(ASibling))
}

func (r *TRichEdit) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    RichEdit_AnchorParallel(r._instance(), ASide , ASpace , CheckPtr(ASibling))
}

// AnchorHorizontalCenterTo
//
// 置于指定控件的横向中心。
func (r *TRichEdit) AnchorHorizontalCenterTo(ASibling IControl) {
    RichEdit_AnchorHorizontalCenterTo(r._instance(), CheckPtr(ASibling))
}

// AnchorVerticalCenterTo
//
// 置于指定控件的纵向中心。
func (r *TRichEdit) AnchorVerticalCenterTo(ASibling IControl) {
    RichEdit_AnchorVerticalCenterTo(r._instance(), CheckPtr(ASibling))
}

func (r *TRichEdit) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    RichEdit_AnchorSame(r._instance(), ASide , CheckPtr(ASibling))
}

func (r *TRichEdit) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    RichEdit_AnchorAsAlign(r._instance(), ATheAlign , ASpace)
}

func (r *TRichEdit) AnchorClient(ASpace int32) {
    RichEdit_AnchorClient(r._instance(), ASpace)
}

func (r *TRichEdit) ScaleDesignToForm(ASize int32) int32 {
    return RichEdit_ScaleDesignToForm(r._instance(), ASize)
}

func (r *TRichEdit) ScaleFormToDesign(ASize int32) int32 {
    return RichEdit_ScaleFormToDesign(r._instance(), ASize)
}

func (r *TRichEdit) Scale96ToForm(ASize int32) int32 {
    return RichEdit_Scale96ToForm(r._instance(), ASize)
}

func (r *TRichEdit) ScaleFormTo96(ASize int32) int32 {
    return RichEdit_ScaleFormTo96(r._instance(), ASize)
}

func (r *TRichEdit) Scale96ToFont(ASize int32) int32 {
    return RichEdit_Scale96ToFont(r._instance(), ASize)
}

func (r *TRichEdit) ScaleFontTo96(ASize int32) int32 {
    return RichEdit_ScaleFontTo96(r._instance(), ASize)
}

func (r *TRichEdit) ScaleScreenToFont(ASize int32) int32 {
    return RichEdit_ScaleScreenToFont(r._instance(), ASize)
}

func (r *TRichEdit) ScaleFontToScreen(ASize int32) int32 {
    return RichEdit_ScaleFontToScreen(r._instance(), ASize)
}

func (r *TRichEdit) Scale96ToScreen(ASize int32) int32 {
    return RichEdit_Scale96ToScreen(r._instance(), ASize)
}

func (r *TRichEdit) ScaleScreenTo96(ASize int32) int32 {
    return RichEdit_ScaleScreenTo96(r._instance(), ASize)
}

func (r *TRichEdit) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    RichEdit_AutoAdjustLayout(r._instance(), AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (r *TRichEdit) FixDesignFontsPPI(ADesignTimePPI int32) {
    RichEdit_FixDesignFontsPPI(r._instance(), ADesignTimePPI)
}

func (r *TRichEdit) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    RichEdit_ScaleFontsPPI(r._instance(), AToPPI , AProportion)
}

// Align
//
// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (r *TRichEdit) Align() TAlign {
    return RichEdit_GetAlign(r._instance())
}

// SetAlign
//
// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (r *TRichEdit) SetAlign(value TAlign) {
    RichEdit_SetAlign(r._instance(), value)
}

// Alignment
//
// 获取文字对齐。
//
// Get Text alignment.
func (r *TRichEdit) Alignment() TAlignment {
    return RichEdit_GetAlignment(r._instance())
}

// SetAlignment
//
// 设置文字对齐。
//
// Set Text alignment.
func (r *TRichEdit) SetAlignment(value TAlignment) {
    RichEdit_SetAlignment(r._instance(), value)
}

// Anchors
//
// 获取四个角位置的锚点。
func (r *TRichEdit) Anchors() TAnchors {
    return RichEdit_GetAnchors(r._instance())
}

// SetAnchors
//
// 设置四个角位置的锚点。
func (r *TRichEdit) SetAnchors(value TAnchors) {
    RichEdit_SetAnchors(r._instance(), value)
}

func (r *TRichEdit) BiDiMode() TBiDiMode {
    return RichEdit_GetBiDiMode(r._instance())
}

func (r *TRichEdit) SetBiDiMode(value TBiDiMode) {
    RichEdit_SetBiDiMode(r._instance(), value)
}

// BorderStyle
//
// 获取窗口边框样式。比如：无边框，单一边框等。
func (r *TRichEdit) BorderStyle() TBorderStyle {
    return RichEdit_GetBorderStyle(r._instance())
}

// SetBorderStyle
//
// 设置窗口边框样式。比如：无边框，单一边框等。
func (r *TRichEdit) SetBorderStyle(value TBorderStyle) {
    RichEdit_SetBorderStyle(r._instance(), value)
}

// BorderWidth
//
// 获取边框的宽度。
func (r *TRichEdit) BorderWidth() int32 {
    return RichEdit_GetBorderWidth(r._instance())
}

// SetBorderWidth
//
// 设置边框的宽度。
func (r *TRichEdit) SetBorderWidth(value int32) {
    RichEdit_SetBorderWidth(r._instance(), value)
}

// Color
//
// 获取颜色。
//
// Get color.
func (r *TRichEdit) Color() TColor {
    return RichEdit_GetColor(r._instance())
}

// SetColor
//
// 设置颜色。
//
// Set color.
func (r *TRichEdit) SetColor(value TColor) {
    RichEdit_SetColor(r._instance(), value)
}

// DragCursor
//
// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (r *TRichEdit) DragCursor() TCursor {
    return RichEdit_GetDragCursor(r._instance())
}

// SetDragCursor
//
// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (r *TRichEdit) SetDragCursor(value TCursor) {
    RichEdit_SetDragCursor(r._instance(), value)
}

// DragKind
//
// 获取拖拽方式。
//
// Get Drag and drop.
func (r *TRichEdit) DragKind() TDragKind {
    return RichEdit_GetDragKind(r._instance())
}

// SetDragKind
//
// 设置拖拽方式。
//
// Set Drag and drop.
func (r *TRichEdit) SetDragKind(value TDragKind) {
    RichEdit_SetDragKind(r._instance(), value)
}

// DragMode
//
// 获取拖拽模式。
//
// Get Drag mode.
func (r *TRichEdit) DragMode() TDragMode {
    return RichEdit_GetDragMode(r._instance())
}

// SetDragMode
//
// 设置拖拽模式。
//
// Set Drag mode.
func (r *TRichEdit) SetDragMode(value TDragMode) {
    RichEdit_SetDragMode(r._instance(), value)
}

// Enabled
//
// 获取控件启用。
//
// Get the control enabled.
func (r *TRichEdit) Enabled() bool {
    return RichEdit_GetEnabled(r._instance())
}

// SetEnabled
//
// 设置控件启用。
//
// Set the control enabled.
func (r *TRichEdit) SetEnabled(value bool) {
    RichEdit_SetEnabled(r._instance(), value)
}

// Font
//
// 获取字体。
//
// Get Font.
func (r *TRichEdit) Font() *TFont {
    return AsFont(RichEdit_GetFont(r._instance()))
}

// SetFont
//
// 设置字体。
//
// Set Font.
func (r *TRichEdit) SetFont(value *TFont) {
    RichEdit_SetFont(r._instance(), CheckPtr(value))
}

// HideSelection
//
// 获取隐藏选择。
func (r *TRichEdit) HideSelection() bool {
    return RichEdit_GetHideSelection(r._instance())
}

// SetHideSelection
//
// 设置隐藏选择。
func (r *TRichEdit) SetHideSelection(value bool) {
    RichEdit_SetHideSelection(r._instance(), value)
}

// Constraints
//
// 获取约束控件大小。
func (r *TRichEdit) Constraints() *TSizeConstraints {
    return AsSizeConstraints(RichEdit_GetConstraints(r._instance()))
}

// SetConstraints
//
// 设置约束控件大小。
func (r *TRichEdit) SetConstraints(value *TSizeConstraints) {
    RichEdit_SetConstraints(r._instance(), CheckPtr(value))
}

func (r *TRichEdit) Lines() *TStrings {
    return AsStrings(RichEdit_GetLines(r._instance()))
}

func (r *TRichEdit) SetLines(value IStrings) {
    RichEdit_SetLines(r._instance(), CheckPtr(value))
}

// MaxLength
//
// 获取最大长度。
func (r *TRichEdit) MaxLength() int32 {
    return RichEdit_GetMaxLength(r._instance())
}

// SetMaxLength
//
// 设置最大长度。
func (r *TRichEdit) SetMaxLength(value int32) {
    RichEdit_SetMaxLength(r._instance(), value)
}

// ParentColor
//
// 获取使用父容器颜色。
//
// Get parent color.
func (r *TRichEdit) ParentColor() bool {
    return RichEdit_GetParentColor(r._instance())
}

// SetParentColor
//
// 设置使用父容器颜色。
//
// Set parent color.
func (r *TRichEdit) SetParentColor(value bool) {
    RichEdit_SetParentColor(r._instance(), value)
}

// ParentFont
//
// 获取使用父容器字体。
//
// Get Parent container font.
func (r *TRichEdit) ParentFont() bool {
    return RichEdit_GetParentFont(r._instance())
}

// SetParentFont
//
// 设置使用父容器字体。
//
// Set Parent container font.
func (r *TRichEdit) SetParentFont(value bool) {
    RichEdit_SetParentFont(r._instance(), value)
}

// ParentShowHint
//
// 获取以父容器的ShowHint属性为准。
func (r *TRichEdit) ParentShowHint() bool {
    return RichEdit_GetParentShowHint(r._instance())
}

// SetParentShowHint
//
// 设置以父容器的ShowHint属性为准。
func (r *TRichEdit) SetParentShowHint(value bool) {
    RichEdit_SetParentShowHint(r._instance(), value)
}

// PopupMenu
//
// 获取右键菜单。
//
// Get Right click menu.
func (r *TRichEdit) PopupMenu() *TPopupMenu {
    return AsPopupMenu(RichEdit_GetPopupMenu(r._instance()))
}

// SetPopupMenu
//
// 设置右键菜单。
//
// Set Right click menu.
func (r *TRichEdit) SetPopupMenu(value IComponent) {
    RichEdit_SetPopupMenu(r._instance(), CheckPtr(value))
}

// ReadOnly
//
// 获取只读。
func (r *TRichEdit) ReadOnly() bool {
    return RichEdit_GetReadOnly(r._instance())
}

// SetReadOnly
//
// 设置只读。
func (r *TRichEdit) SetReadOnly(value bool) {
    RichEdit_SetReadOnly(r._instance(), value)
}

func (r *TRichEdit) ScrollBars() TScrollStyle {
    return RichEdit_GetScrollBars(r._instance())
}

func (r *TRichEdit) SetScrollBars(value TScrollStyle) {
    RichEdit_SetScrollBars(r._instance(), value)
}

// ShowHint
//
// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (r *TRichEdit) ShowHint() bool {
    return RichEdit_GetShowHint(r._instance())
}

// SetShowHint
//
// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (r *TRichEdit) SetShowHint(value bool) {
    RichEdit_SetShowHint(r._instance(), value)
}

// TabOrder
//
// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (r *TRichEdit) TabOrder() TTabOrder {
    return RichEdit_GetTabOrder(r._instance())
}

// SetTabOrder
//
// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (r *TRichEdit) SetTabOrder(value TTabOrder) {
    RichEdit_SetTabOrder(r._instance(), value)
}

// TabStop
//
// 获取Tab可停留。
//
// Get Tab can stay.
func (r *TRichEdit) TabStop() bool {
    return RichEdit_GetTabStop(r._instance())
}

// SetTabStop
//
// 设置Tab可停留。
//
// Set Tab can stay.
func (r *TRichEdit) SetTabStop(value bool) {
    RichEdit_SetTabStop(r._instance(), value)
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (r *TRichEdit) Visible() bool {
    return RichEdit_GetVisible(r._instance())
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (r *TRichEdit) SetVisible(value bool) {
    RichEdit_SetVisible(r._instance(), value)
}

func (r *TRichEdit) WantTabs() bool {
    return RichEdit_GetWantTabs(r._instance())
}

func (r *TRichEdit) SetWantTabs(value bool) {
    RichEdit_SetWantTabs(r._instance(), value)
}

func (r *TRichEdit) WantReturns() bool {
    return RichEdit_GetWantReturns(r._instance())
}

func (r *TRichEdit) SetWantReturns(value bool) {
    RichEdit_SetWantReturns(r._instance(), value)
}

// WordWrap
//
// 获取自动换行。
//
// Get Automatic line break.
func (r *TRichEdit) WordWrap() bool {
    return RichEdit_GetWordWrap(r._instance())
}

// SetWordWrap
//
// 设置自动换行。
//
// Set Automatic line break.
func (r *TRichEdit) SetWordWrap(value bool) {
    RichEdit_SetWordWrap(r._instance(), value)
}

func (r *TRichEdit) Zoom() int32 {
    return RichEdit_GetZoom(r._instance())
}

func (r *TRichEdit) SetZoom(value int32) {
    RichEdit_SetZoom(r._instance(), value)
}

// SetOnChange
//
// 设置改变事件。
//
// Set changed event.
func (r *TRichEdit) SetOnChange(fn TNotifyEvent) {
    RichEdit_SetOnChange(r._instance(), fn)
}

// SetOnClick
//
// 设置控件单击事件。
//
// Set control click event.
func (r *TRichEdit) SetOnClick(fn TNotifyEvent) {
    RichEdit_SetOnClick(r._instance(), fn)
}

// SetOnContextPopup
//
// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (r *TRichEdit) SetOnContextPopup(fn TContextPopupEvent) {
    RichEdit_SetOnContextPopup(r._instance(), fn)
}

// SetOnDblClick
//
// 设置双击事件。
func (r *TRichEdit) SetOnDblClick(fn TNotifyEvent) {
    RichEdit_SetOnDblClick(r._instance(), fn)
}

// SetOnDragDrop
//
// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (r *TRichEdit) SetOnDragDrop(fn TDragDropEvent) {
    RichEdit_SetOnDragDrop(r._instance(), fn)
}

// SetOnDragOver
//
// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (r *TRichEdit) SetOnDragOver(fn TDragOverEvent) {
    RichEdit_SetOnDragOver(r._instance(), fn)
}

// SetOnEndDrag
//
// 设置拖拽结束。
//
// Set End of drag.
func (r *TRichEdit) SetOnEndDrag(fn TEndDragEvent) {
    RichEdit_SetOnEndDrag(r._instance(), fn)
}

// SetOnEnter
//
// 设置焦点进入。
//
// Set Focus entry.
func (r *TRichEdit) SetOnEnter(fn TNotifyEvent) {
    RichEdit_SetOnEnter(r._instance(), fn)
}

// SetOnExit
//
// 设置焦点退出。
//
// Set Focus exit.
func (r *TRichEdit) SetOnExit(fn TNotifyEvent) {
    RichEdit_SetOnExit(r._instance(), fn)
}

// SetOnKeyDown
//
// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (r *TRichEdit) SetOnKeyDown(fn TKeyEvent) {
    RichEdit_SetOnKeyDown(r._instance(), fn)
}

// SetOnKeyPress
//
// 设置键键下事件。
func (r *TRichEdit) SetOnKeyPress(fn TKeyPressEvent) {
    RichEdit_SetOnKeyPress(r._instance(), fn)
}

// SetOnKeyUp
//
// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (r *TRichEdit) SetOnKeyUp(fn TKeyEvent) {
    RichEdit_SetOnKeyUp(r._instance(), fn)
}

// SetOnMouseDown
//
// 设置鼠标按下事件。
//
// Set Mouse down event.
func (r *TRichEdit) SetOnMouseDown(fn TMouseEvent) {
    RichEdit_SetOnMouseDown(r._instance(), fn)
}

// SetOnMouseEnter
//
// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (r *TRichEdit) SetOnMouseEnter(fn TNotifyEvent) {
    RichEdit_SetOnMouseEnter(r._instance(), fn)
}

// SetOnMouseLeave
//
// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (r *TRichEdit) SetOnMouseLeave(fn TNotifyEvent) {
    RichEdit_SetOnMouseLeave(r._instance(), fn)
}

// SetOnMouseMove
//
// 设置鼠标移动事件。
func (r *TRichEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    RichEdit_SetOnMouseMove(r._instance(), fn)
}

// SetOnMouseUp
//
// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (r *TRichEdit) SetOnMouseUp(fn TMouseEvent) {
    RichEdit_SetOnMouseUp(r._instance(), fn)
}

// SetOnMouseWheel
//
// 设置鼠标滚轮事件。
func (r *TRichEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
    RichEdit_SetOnMouseWheel(r._instance(), fn)
}

// SetOnMouseWheelDown
//
// 设置鼠标滚轮按下事件。
func (r *TRichEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    RichEdit_SetOnMouseWheelDown(r._instance(), fn)
}

// SetOnMouseWheelUp
//
// 设置鼠标滚轮抬起事件。
func (r *TRichEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    RichEdit_SetOnMouseWheelUp(r._instance(), fn)
}

func (r *TRichEdit) DefAttributes() *TTextAttributes {
    return AsTextAttributes(RichEdit_GetDefAttributes(r._instance()))
}

func (r *TRichEdit) SetDefAttributes(value *TTextAttributes) {
    RichEdit_SetDefAttributes(r._instance(), CheckPtr(value))
}

func (r *TRichEdit) SelAttributes() *TTextAttributes {
    return AsTextAttributes(RichEdit_GetSelAttributes(r._instance()))
}

func (r *TRichEdit) SetSelAttributes(value *TTextAttributes) {
    RichEdit_SetSelAttributes(r._instance(), CheckPtr(value))
}

func (r *TRichEdit) Paragraph() *TParaAttributes {
    return AsParaAttributes(RichEdit_GetParagraph(r._instance()))
}

func (r *TRichEdit) CaretPos() TPoint {
    return RichEdit_GetCaretPos(r._instance())
}

func (r *TRichEdit) SetCaretPos(value TPoint) {
    RichEdit_SetCaretPos(r._instance(), value)
}

// CanUndo
//
// 获取能否撤销。
func (r *TRichEdit) CanUndo() bool {
    return RichEdit_GetCanUndo(r._instance())
}

// Modified
//
// 获取修改。
//
// Get modified.
func (r *TRichEdit) Modified() bool {
    return RichEdit_GetModified(r._instance())
}

// SetModified
//
// 设置修改。
//
// Set modified.
func (r *TRichEdit) SetModified(value bool) {
    RichEdit_SetModified(r._instance(), value)
}

// SelLength
//
// 获取选择的长度。
func (r *TRichEdit) SelLength() int32 {
    return RichEdit_GetSelLength(r._instance())
}

// SetSelLength
//
// 设置选择的长度。
func (r *TRichEdit) SetSelLength(value int32) {
    RichEdit_SetSelLength(r._instance(), value)
}

// SelStart
//
// 获取选择的启始位置。
func (r *TRichEdit) SelStart() int32 {
    return RichEdit_GetSelStart(r._instance())
}

// SetSelStart
//
// 设置选择的启始位置。
func (r *TRichEdit) SetSelStart(value int32) {
    RichEdit_SetSelStart(r._instance(), value)
}

// SelText
//
// 获取选择的文本。
func (r *TRichEdit) SelText() string {
    return RichEdit_GetSelText(r._instance())
}

// SetSelText
//
// 设置选择的文本。
func (r *TRichEdit) SetSelText(value string) {
    RichEdit_SetSelText(r._instance(), value)
}

// Text
//
// 获取文本。
func (r *TRichEdit) Text() string {
    return getControlBufferText(r.GetTextLen, r.GetTextBuf)
}

// SetText
//
// 设置文本。
func (r *TRichEdit) SetText(value string) {
    RichEdit_SetText(r._instance(), value)
}

// TextHint
//
// 获取提示文本。
func (r *TRichEdit) TextHint() string {
    return RichEdit_GetTextHint(r._instance())
}

// SetTextHint
//
// 设置提示文本。
func (r *TRichEdit) SetTextHint(value string) {
    RichEdit_SetTextHint(r._instance(), value)
}

// DockClientCount
//
// 获取依靠客户端总数。
func (r *TRichEdit) DockClientCount() int32 {
    return RichEdit_GetDockClientCount(r._instance())
}

// DockSite
//
// 获取停靠站点。
//
// Get Docking site.
func (r *TRichEdit) DockSite() bool {
    return RichEdit_GetDockSite(r._instance())
}

// SetDockSite
//
// 设置停靠站点。
//
// Set Docking site.
func (r *TRichEdit) SetDockSite(value bool) {
    RichEdit_SetDockSite(r._instance(), value)
}

// DoubleBuffered
//
// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (r *TRichEdit) DoubleBuffered() bool {
    return RichEdit_GetDoubleBuffered(r._instance())
}

// SetDoubleBuffered
//
// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (r *TRichEdit) SetDoubleBuffered(value bool) {
    RichEdit_SetDoubleBuffered(r._instance(), value)
}

// MouseInClient
//
// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (r *TRichEdit) MouseInClient() bool {
    return RichEdit_GetMouseInClient(r._instance())
}

// VisibleDockClientCount
//
// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (r *TRichEdit) VisibleDockClientCount() int32 {
    return RichEdit_GetVisibleDockClientCount(r._instance())
}

// Brush
//
// 获取画刷对象。
//
// Get Brush.
func (r *TRichEdit) Brush() *TBrush {
    return AsBrush(RichEdit_GetBrush(r._instance()))
}

// ControlCount
//
// 获取子控件数。
//
// Get Number of child controls.
func (r *TRichEdit) ControlCount() int32 {
    return RichEdit_GetControlCount(r._instance())
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (r *TRichEdit) Handle() HWND {
    return RichEdit_GetHandle(r._instance())
}

// ParentDoubleBuffered
//
// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (r *TRichEdit) ParentDoubleBuffered() bool {
    return RichEdit_GetParentDoubleBuffered(r._instance())
}

// SetParentDoubleBuffered
//
// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (r *TRichEdit) SetParentDoubleBuffered(value bool) {
    RichEdit_SetParentDoubleBuffered(r._instance(), value)
}

// ParentWindow
//
// 获取父容器句柄。
//
// Get Parent container handle.
func (r *TRichEdit) ParentWindow() HWND {
    return RichEdit_GetParentWindow(r._instance())
}

// SetParentWindow
//
// 设置父容器句柄。
//
// Set Parent container handle.
func (r *TRichEdit) SetParentWindow(value HWND) {
    RichEdit_SetParentWindow(r._instance(), value)
}

func (r *TRichEdit) Showing() bool {
    return RichEdit_GetShowing(r._instance())
}

// UseDockManager
//
// 获取使用停靠管理。
func (r *TRichEdit) UseDockManager() bool {
    return RichEdit_GetUseDockManager(r._instance())
}

// SetUseDockManager
//
// 设置使用停靠管理。
func (r *TRichEdit) SetUseDockManager(value bool) {
    RichEdit_SetUseDockManager(r._instance(), value)
}

func (r *TRichEdit) Action() *TAction {
    return AsAction(RichEdit_GetAction(r._instance()))
}

func (r *TRichEdit) SetAction(value IComponent) {
    RichEdit_SetAction(r._instance(), CheckPtr(value))
}

func (r *TRichEdit) BoundsRect() TRect {
    return RichEdit_GetBoundsRect(r._instance())
}

func (r *TRichEdit) SetBoundsRect(value TRect) {
    RichEdit_SetBoundsRect(r._instance(), value)
}

// ClientHeight
//
// 获取客户区高度。
//
// Get client height.
func (r *TRichEdit) ClientHeight() int32 {
    return RichEdit_GetClientHeight(r._instance())
}

// SetClientHeight
//
// 设置客户区高度。
//
// Set client height.
func (r *TRichEdit) SetClientHeight(value int32) {
    RichEdit_SetClientHeight(r._instance(), value)
}

func (r *TRichEdit) ClientOrigin() TPoint {
    return RichEdit_GetClientOrigin(r._instance())
}

// ClientRect
//
// 获取客户区矩形。
//
// Get client rectangle.
func (r *TRichEdit) ClientRect() TRect {
    return RichEdit_GetClientRect(r._instance())
}

// ClientWidth
//
// 获取客户区宽度。
//
// Get client width.
func (r *TRichEdit) ClientWidth() int32 {
    return RichEdit_GetClientWidth(r._instance())
}

// SetClientWidth
//
// 设置客户区宽度。
//
// Set client width.
func (r *TRichEdit) SetClientWidth(value int32) {
    RichEdit_SetClientWidth(r._instance(), value)
}

// ControlState
//
// 获取控件状态。
//
// Get control state.
func (r *TRichEdit) ControlState() TControlState {
    return RichEdit_GetControlState(r._instance())
}

// SetControlState
//
// 设置控件状态。
//
// Set control state.
func (r *TRichEdit) SetControlState(value TControlState) {
    RichEdit_SetControlState(r._instance(), value)
}

// ControlStyle
//
// 获取控件样式。
//
// Get control style.
func (r *TRichEdit) ControlStyle() TControlStyle {
    return RichEdit_GetControlStyle(r._instance())
}

// SetControlStyle
//
// 设置控件样式。
//
// Set control style.
func (r *TRichEdit) SetControlStyle(value TControlStyle) {
    RichEdit_SetControlStyle(r._instance(), value)
}

func (r *TRichEdit) Floating() bool {
    return RichEdit_GetFloating(r._instance())
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (r *TRichEdit) Parent() *TWinControl {
    return AsWinControl(RichEdit_GetParent(r._instance()))
}

// SetParent
//
// 设置控件父容器。
//
// Set control parent container.
func (r *TRichEdit) SetParent(value IWinControl) {
    RichEdit_SetParent(r._instance(), CheckPtr(value))
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (r *TRichEdit) Left() int32 {
    return RichEdit_GetLeft(r._instance())
}

// SetLeft
//
// 设置左边位置。
//
// Set Left position.
func (r *TRichEdit) SetLeft(value int32) {
    RichEdit_SetLeft(r._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (r *TRichEdit) Top() int32 {
    return RichEdit_GetTop(r._instance())
}

// SetTop
//
// 设置顶边位置。
//
// Set Top position.
func (r *TRichEdit) SetTop(value int32) {
    RichEdit_SetTop(r._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (r *TRichEdit) Width() int32 {
    return RichEdit_GetWidth(r._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (r *TRichEdit) SetWidth(value int32) {
    RichEdit_SetWidth(r._instance(), value)
}

// Height
//
// 获取高度。
//
// Get height.
func (r *TRichEdit) Height() int32 {
    return RichEdit_GetHeight(r._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (r *TRichEdit) SetHeight(value int32) {
    RichEdit_SetHeight(r._instance(), value)
}

// Cursor
//
// 获取控件光标。
//
// Get control cursor.
func (r *TRichEdit) Cursor() TCursor {
    return RichEdit_GetCursor(r._instance())
}

// SetCursor
//
// 设置控件光标。
//
// Set control cursor.
func (r *TRichEdit) SetCursor(value TCursor) {
    RichEdit_SetCursor(r._instance(), value)
}

// Hint
//
// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (r *TRichEdit) Hint() string {
    return RichEdit_GetHint(r._instance())
}

// SetHint
//
// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (r *TRichEdit) SetHint(value string) {
    RichEdit_SetHint(r._instance(), value)
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (r *TRichEdit) ComponentCount() int32 {
    return RichEdit_GetComponentCount(r._instance())
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (r *TRichEdit) ComponentIndex() int32 {
    return RichEdit_GetComponentIndex(r._instance())
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (r *TRichEdit) SetComponentIndex(value int32) {
    RichEdit_SetComponentIndex(r._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (r *TRichEdit) Owner() *TComponent {
    return AsComponent(RichEdit_GetOwner(r._instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (r *TRichEdit) Name() string {
    return RichEdit_GetName(r._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (r *TRichEdit) SetName(value string) {
    RichEdit_SetName(r._instance(), value)
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (r *TRichEdit) Tag() int {
    return RichEdit_GetTag(r._instance())
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (r *TRichEdit) SetTag(value int) {
    RichEdit_SetTag(r._instance(), value)
}

// AnchorSideLeft
//
// 获取左边锚点。
func (r *TRichEdit) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(RichEdit_GetAnchorSideLeft(r._instance()))
}

// SetAnchorSideLeft
//
// 设置左边锚点。
func (r *TRichEdit) SetAnchorSideLeft(value *TAnchorSide) {
    RichEdit_SetAnchorSideLeft(r._instance(), CheckPtr(value))
}

// AnchorSideTop
//
// 获取顶边锚点。
func (r *TRichEdit) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(RichEdit_GetAnchorSideTop(r._instance()))
}

// SetAnchorSideTop
//
// 设置顶边锚点。
func (r *TRichEdit) SetAnchorSideTop(value *TAnchorSide) {
    RichEdit_SetAnchorSideTop(r._instance(), CheckPtr(value))
}

// AnchorSideRight
//
// 获取右边锚点。
func (r *TRichEdit) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(RichEdit_GetAnchorSideRight(r._instance()))
}

// SetAnchorSideRight
//
// 设置右边锚点。
func (r *TRichEdit) SetAnchorSideRight(value *TAnchorSide) {
    RichEdit_SetAnchorSideRight(r._instance(), CheckPtr(value))
}

// AnchorSideBottom
//
// 获取底边锚点。
func (r *TRichEdit) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(RichEdit_GetAnchorSideBottom(r._instance()))
}

// SetAnchorSideBottom
//
// 设置底边锚点。
func (r *TRichEdit) SetAnchorSideBottom(value *TAnchorSide) {
    RichEdit_SetAnchorSideBottom(r._instance(), CheckPtr(value))
}

func (r *TRichEdit) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(RichEdit_GetChildSizing(r._instance()))
}

func (r *TRichEdit) SetChildSizing(value *TControlChildSizing) {
    RichEdit_SetChildSizing(r._instance(), CheckPtr(value))
}

// BorderSpacing
//
// 获取边框间距。
func (r *TRichEdit) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(RichEdit_GetBorderSpacing(r._instance()))
}

// SetBorderSpacing
//
// 设置边框间距。
func (r *TRichEdit) SetBorderSpacing(value *TControlBorderSpacing) {
    RichEdit_SetBorderSpacing(r._instance(), CheckPtr(value))
}

// DockClients
//
// 获取指定索引停靠客户端。
func (r *TRichEdit) DockClients(Index int32) *TControl {
    return AsControl(RichEdit_GetDockClients(r._instance(), Index))
}

// Controls
//
// 获取指定索引子控件。
func (r *TRichEdit) Controls(Index int32) *TControl {
    return AsControl(RichEdit_GetControls(r._instance(), Index))
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (r *TRichEdit) Components(AIndex int32) *TComponent {
    return AsComponent(RichEdit_GetComponents(r._instance(), AIndex))
}

// AnchorSide
//
// 获取锚侧面。
func (r *TRichEdit) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(RichEdit_GetAnchorSide(r._instance(), AKind))
}

