
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

type TButton struct {
    IWinControl
    instance unsafe.Pointer
}

// NewButton
//
// 创建一个新的对象。
// 
// Create a new object.
func NewButton(owner IComponent) *TButton {
    b := new(TButton)
    b.instance = unsafe.Pointer(Button_Create(CheckPtr(owner)))
    return b
}

// AsButton
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsButton(obj interface{}) *TButton {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TButton{instance: instance}
}

// Free 
//
// 释放对象。
// 
// Free object.
func (b *TButton) Free() {
    if b.instance != nullptr {
        Button_Free(b._instance())
        b.instance  = nullptr
    }
}

func (b *TButton) _instance() uintptr {
    return uintptr(b.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (b *TButton) Instance() uintptr {
    return b._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (b *TButton) UnsafeAddr() unsafe.Pointer {
    return b.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (b *TButton) IsValid() bool {
    return b.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (b *TButton) Is() TIs {
    return TIs(b._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (b *TButton) As() TAs {
//    return TAs(b._instance())
//}

// TButtonClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TButtonClass() TClass {
    return Button_StaticClassType()
}

// Click
//
// 单击。
func (b *TButton) Click() {
    Button_Click(b._instance())
}

// CanFocus
//
// 是否可以获得焦点。
func (b *TButton) CanFocus() bool {
    return Button_CanFocus(b._instance())
}

// ContainsControl
//
// 返回是否包含指定控件。
//
// it's contain a specified control.
func (b *TButton) ContainsControl(Control IControl) bool {
    return Button_ContainsControl(b._instance(), CheckPtr(Control))
}

// ControlAtPos
//
// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (b *TButton) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(Button_ControlAtPos(b._instance(), Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
//
// 禁用控件的对齐。
//
// Disable control alignment.
func (b *TButton) DisableAlign() {
    Button_DisableAlign(b._instance())
}

// EnableAlign
//
// 启用控件对齐。
//
// Enabled control alignment.
func (b *TButton) EnableAlign() {
    Button_EnableAlign(b._instance())
}

// FindChildControl
//
// 查找子控件。
//
// Find sub controls.
func (b *TButton) FindChildControl(ControlName string) *TControl {
    return AsControl(Button_FindChildControl(b._instance(), ControlName))
}

func (b *TButton) FlipChildren(AllLevels bool) {
    Button_FlipChildren(b._instance(), AllLevels)
}

// Focused
//
// 返回是否获取焦点。
//
// Return to get focus.
func (b *TButton) Focused() bool {
    return Button_Focused(b._instance())
}

// HandleAllocated
//
// 句柄是否已经分配。
//
// Is the handle already allocated.
func (b *TButton) HandleAllocated() bool {
    return Button_HandleAllocated(b._instance())
}

// InsertControl
//
// 插入一个控件。
//
// Insert a control.
func (b *TButton) InsertControl(AControl IControl) {
    Button_InsertControl(b._instance(), CheckPtr(AControl))
}

// Invalidate
//
// 要求重绘。
//
// Redraw.
func (b *TButton) Invalidate() {
    Button_Invalidate(b._instance())
}

// PaintTo
//
// 绘画至指定DC。
//
// Painting to the specified DC.
func (b *TButton) PaintTo(DC HDC, X int32, Y int32) {
    Button_PaintTo(b._instance(), DC , X , Y)
}

// RemoveControl
//
// 移除一个控件。
//
// Remove a control.
func (b *TButton) RemoveControl(AControl IControl) {
    Button_RemoveControl(b._instance(), CheckPtr(AControl))
}

// Realign
//
// 重新对齐。
//
// Realign.
func (b *TButton) Realign() {
    Button_Realign(b._instance())
}

// Repaint
//
// 重绘。
//
// Repaint.
func (b *TButton) Repaint() {
    Button_Repaint(b._instance())
}

// ScaleBy
//
// 按比例缩放。
//
// Scale by.
func (b *TButton) ScaleBy(M int32, D int32) {
    Button_ScaleBy(b._instance(), M , D)
}

// ScrollBy
//
// 滚动至指定位置。
//
// Scroll by.
func (b *TButton) ScrollBy(DeltaX int32, DeltaY int32) {
    Button_ScrollBy(b._instance(), DeltaX , DeltaY)
}

// SetBounds
//
// 设置组件边界。
//
// Set component boundaries.
func (b *TButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Button_SetBounds(b._instance(), ALeft , ATop , AWidth , AHeight)
}

// SetFocus
//
// 设置控件焦点。
//
// Set control focus.
func (b *TButton) SetFocus() {
    Button_SetFocus(b._instance())
}

// Update
//
// 控件更新。
//
// Update.
func (b *TButton) Update() {
    Button_Update(b._instance())
}

// BringToFront
//
// 将控件置于最前。
//
// Bring the control to the front.
func (b *TButton) BringToFront() {
    Button_BringToFront(b._instance())
}

// ClientToScreen
//
// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (b *TButton) ClientToScreen(Point TPoint) TPoint {
    return Button_ClientToScreen(b._instance(), Point)
}

// ClientToParent
//
// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (b *TButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Button_ClientToParent(b._instance(), Point , CheckPtr(AParent))
}

// Dragging
//
// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (b *TButton) Dragging() bool {
    return Button_Dragging(b._instance())
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (b *TButton) HasParent() bool {
    return Button_HasParent(b._instance())
}

// Hide
//
// 隐藏控件。
//
// Hidden control.
func (b *TButton) Hide() {
    Button_Hide(b._instance())
}

// Perform
//
// 发送一个消息。
//
// Send a message.
func (b *TButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Button_Perform(b._instance(), Msg , WParam , LParam)
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (b *TButton) Refresh() {
    Button_Refresh(b._instance())
}

// ScreenToClient
//
// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (b *TButton) ScreenToClient(Point TPoint) TPoint {
    return Button_ScreenToClient(b._instance(), Point)
}

// ParentToClient
//
// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (b *TButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Button_ParentToClient(b._instance(), Point , CheckPtr(AParent))
}

// SendToBack
//
// 控件至于最后面。
//
// The control is placed at the end.
func (b *TButton) SendToBack() {
    Button_SendToBack(b._instance())
}

// Show
//
// 显示控件。
//
// Show control.
func (b *TButton) Show() {
    Button_Show(b._instance())
}

// GetTextBuf
//
// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (b *TButton) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Button_GetTextBuf(b._instance(), Buffer , BufSize)
}

// GetTextLen
//
// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (b *TButton) GetTextLen() int32 {
    return Button_GetTextLen(b._instance())
}

// SetTextBuf
//
// 设置控件字符，如果有。
//
// Set control characters, if any.
func (b *TButton) SetTextBuf(Buffer string) {
    Button_SetTextBuf(b._instance(), Buffer)
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (b *TButton) FindComponent(AName string) *TComponent {
    return AsComponent(Button_FindComponent(b._instance(), AName))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (b *TButton) GetNamePath() string {
    return Button_GetNamePath(b._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (b *TButton) Assign(Source IObject) {
    Button_Assign(b._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (b *TButton) ClassType() TClass {
    return Button_ClassType(b._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (b *TButton) ClassName() string {
    return Button_ClassName(b._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (b *TButton) InstanceSize() int32 {
    return Button_InstanceSize(b._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (b *TButton) InheritsFrom(AClass TClass) bool {
    return Button_InheritsFrom(b._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (b *TButton) Equals(Obj IObject) bool {
    return Button_Equals(b._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (b *TButton) GetHashCode() int32 {
    return Button_GetHashCode(b._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (b *TButton) ToString() string {
    return Button_ToString(b._instance())
}

func (b *TButton) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Button_AnchorToNeighbour(b._instance(), ASide , ASpace , CheckPtr(ASibling))
}

func (b *TButton) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Button_AnchorParallel(b._instance(), ASide , ASpace , CheckPtr(ASibling))
}

// AnchorHorizontalCenterTo
//
// 置于指定控件的横向中心。
func (b *TButton) AnchorHorizontalCenterTo(ASibling IControl) {
    Button_AnchorHorizontalCenterTo(b._instance(), CheckPtr(ASibling))
}

// AnchorVerticalCenterTo
//
// 置于指定控件的纵向中心。
func (b *TButton) AnchorVerticalCenterTo(ASibling IControl) {
    Button_AnchorVerticalCenterTo(b._instance(), CheckPtr(ASibling))
}

func (b *TButton) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    Button_AnchorSame(b._instance(), ASide , CheckPtr(ASibling))
}

func (b *TButton) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    Button_AnchorAsAlign(b._instance(), ATheAlign , ASpace)
}

func (b *TButton) AnchorClient(ASpace int32) {
    Button_AnchorClient(b._instance(), ASpace)
}

func (b *TButton) ScaleDesignToForm(ASize int32) int32 {
    return Button_ScaleDesignToForm(b._instance(), ASize)
}

func (b *TButton) ScaleFormToDesign(ASize int32) int32 {
    return Button_ScaleFormToDesign(b._instance(), ASize)
}

func (b *TButton) Scale96ToForm(ASize int32) int32 {
    return Button_Scale96ToForm(b._instance(), ASize)
}

func (b *TButton) ScaleFormTo96(ASize int32) int32 {
    return Button_ScaleFormTo96(b._instance(), ASize)
}

func (b *TButton) Scale96ToFont(ASize int32) int32 {
    return Button_Scale96ToFont(b._instance(), ASize)
}

func (b *TButton) ScaleFontTo96(ASize int32) int32 {
    return Button_ScaleFontTo96(b._instance(), ASize)
}

func (b *TButton) ScaleScreenToFont(ASize int32) int32 {
    return Button_ScaleScreenToFont(b._instance(), ASize)
}

func (b *TButton) ScaleFontToScreen(ASize int32) int32 {
    return Button_ScaleFontToScreen(b._instance(), ASize)
}

func (b *TButton) Scale96ToScreen(ASize int32) int32 {
    return Button_Scale96ToScreen(b._instance(), ASize)
}

func (b *TButton) ScaleScreenTo96(ASize int32) int32 {
    return Button_ScaleScreenTo96(b._instance(), ASize)
}

func (b *TButton) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    Button_AutoAdjustLayout(b._instance(), AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (b *TButton) FixDesignFontsPPI(ADesignTimePPI int32) {
    Button_FixDesignFontsPPI(b._instance(), ADesignTimePPI)
}

func (b *TButton) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    Button_ScaleFontsPPI(b._instance(), AToPPI , AProportion)
}

func (b *TButton) Action() *TAction {
    return AsAction(Button_GetAction(b._instance()))
}

func (b *TButton) SetAction(value IComponent) {
    Button_SetAction(b._instance(), CheckPtr(value))
}

// Align
//
// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (b *TButton) Align() TAlign {
    return Button_GetAlign(b._instance())
}

// SetAlign
//
// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (b *TButton) SetAlign(value TAlign) {
    Button_SetAlign(b._instance(), value)
}

// Anchors
//
// 获取四个角位置的锚点。
func (b *TButton) Anchors() TAnchors {
    return Button_GetAnchors(b._instance())
}

// SetAnchors
//
// 设置四个角位置的锚点。
func (b *TButton) SetAnchors(value TAnchors) {
    Button_SetAnchors(b._instance(), value)
}

func (b *TButton) BiDiMode() TBiDiMode {
    return Button_GetBiDiMode(b._instance())
}

func (b *TButton) SetBiDiMode(value TBiDiMode) {
    Button_SetBiDiMode(b._instance(), value)
}

func (b *TButton) Cancel() bool {
    return Button_GetCancel(b._instance())
}

func (b *TButton) SetCancel(value bool) {
    Button_SetCancel(b._instance(), value)
}

// Caption
//
// 获取控件标题。
//
// Get the control title.
func (b *TButton) Caption() string {
    return Button_GetCaption(b._instance())
}

// SetCaption
//
// 设置控件标题。
//
// Set the control title.
func (b *TButton) SetCaption(value string) {
    Button_SetCaption(b._instance(), value)
}

// Constraints
//
// 获取约束控件大小。
func (b *TButton) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Button_GetConstraints(b._instance()))
}

// SetConstraints
//
// 设置约束控件大小。
func (b *TButton) SetConstraints(value *TSizeConstraints) {
    Button_SetConstraints(b._instance(), CheckPtr(value))
}

func (b *TButton) Default() bool {
    return Button_GetDefault(b._instance())
}

func (b *TButton) SetDefault(value bool) {
    Button_SetDefault(b._instance(), value)
}

// DoubleBuffered
//
// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (b *TButton) DoubleBuffered() bool {
    return Button_GetDoubleBuffered(b._instance())
}

// SetDoubleBuffered
//
// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (b *TButton) SetDoubleBuffered(value bool) {
    Button_SetDoubleBuffered(b._instance(), value)
}

// DragCursor
//
// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (b *TButton) DragCursor() TCursor {
    return Button_GetDragCursor(b._instance())
}

// SetDragCursor
//
// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (b *TButton) SetDragCursor(value TCursor) {
    Button_SetDragCursor(b._instance(), value)
}

// DragKind
//
// 获取拖拽方式。
//
// Get Drag and drop.
func (b *TButton) DragKind() TDragKind {
    return Button_GetDragKind(b._instance())
}

// SetDragKind
//
// 设置拖拽方式。
//
// Set Drag and drop.
func (b *TButton) SetDragKind(value TDragKind) {
    Button_SetDragKind(b._instance(), value)
}

// DragMode
//
// 获取拖拽模式。
//
// Get Drag mode.
func (b *TButton) DragMode() TDragMode {
    return Button_GetDragMode(b._instance())
}

// SetDragMode
//
// 设置拖拽模式。
//
// Set Drag mode.
func (b *TButton) SetDragMode(value TDragMode) {
    Button_SetDragMode(b._instance(), value)
}

// Enabled
//
// 获取控件启用。
//
// Get the control enabled.
func (b *TButton) Enabled() bool {
    return Button_GetEnabled(b._instance())
}

// SetEnabled
//
// 设置控件启用。
//
// Set the control enabled.
func (b *TButton) SetEnabled(value bool) {
    Button_SetEnabled(b._instance(), value)
}

// Font
//
// 获取字体。
//
// Get Font.
func (b *TButton) Font() *TFont {
    return AsFont(Button_GetFont(b._instance()))
}

// SetFont
//
// 设置字体。
//
// Set Font.
func (b *TButton) SetFont(value *TFont) {
    Button_SetFont(b._instance(), CheckPtr(value))
}

// ModalResult
//
// 获取模态对话框显示结果。
func (b *TButton) ModalResult() TModalResult {
    return Button_GetModalResult(b._instance())
}

// SetModalResult
//
// 设置模态对话框显示结果。
func (b *TButton) SetModalResult(value TModalResult) {
    Button_SetModalResult(b._instance(), value)
}

// ParentDoubleBuffered
//
// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (b *TButton) ParentDoubleBuffered() bool {
    return Button_GetParentDoubleBuffered(b._instance())
}

// SetParentDoubleBuffered
//
// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (b *TButton) SetParentDoubleBuffered(value bool) {
    Button_SetParentDoubleBuffered(b._instance(), value)
}

// ParentFont
//
// 获取使用父容器字体。
//
// Get Parent container font.
func (b *TButton) ParentFont() bool {
    return Button_GetParentFont(b._instance())
}

// SetParentFont
//
// 设置使用父容器字体。
//
// Set Parent container font.
func (b *TButton) SetParentFont(value bool) {
    Button_SetParentFont(b._instance(), value)
}

// ParentShowHint
//
// 获取以父容器的ShowHint属性为准。
func (b *TButton) ParentShowHint() bool {
    return Button_GetParentShowHint(b._instance())
}

// SetParentShowHint
//
// 设置以父容器的ShowHint属性为准。
func (b *TButton) SetParentShowHint(value bool) {
    Button_SetParentShowHint(b._instance(), value)
}

// PopupMenu
//
// 获取右键菜单。
//
// Get Right click menu.
func (b *TButton) PopupMenu() *TPopupMenu {
    return AsPopupMenu(Button_GetPopupMenu(b._instance()))
}

// SetPopupMenu
//
// 设置右键菜单。
//
// Set Right click menu.
func (b *TButton) SetPopupMenu(value IComponent) {
    Button_SetPopupMenu(b._instance(), CheckPtr(value))
}

// ShowHint
//
// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (b *TButton) ShowHint() bool {
    return Button_GetShowHint(b._instance())
}

// SetShowHint
//
// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (b *TButton) SetShowHint(value bool) {
    Button_SetShowHint(b._instance(), value)
}

// TabOrder
//
// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (b *TButton) TabOrder() TTabOrder {
    return Button_GetTabOrder(b._instance())
}

// SetTabOrder
//
// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (b *TButton) SetTabOrder(value TTabOrder) {
    Button_SetTabOrder(b._instance(), value)
}

// TabStop
//
// 获取Tab可停留。
//
// Get Tab can stay.
func (b *TButton) TabStop() bool {
    return Button_GetTabStop(b._instance())
}

// SetTabStop
//
// 设置Tab可停留。
//
// Set Tab can stay.
func (b *TButton) SetTabStop(value bool) {
    Button_SetTabStop(b._instance(), value)
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (b *TButton) Visible() bool {
    return Button_GetVisible(b._instance())
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (b *TButton) SetVisible(value bool) {
    Button_SetVisible(b._instance(), value)
}

// SetOnClick
//
// 设置控件单击事件。
//
// Set control click event.
func (b *TButton) SetOnClick(fn TNotifyEvent) {
    Button_SetOnClick(b._instance(), fn)
}

// SetOnContextPopup
//
// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (b *TButton) SetOnContextPopup(fn TContextPopupEvent) {
    Button_SetOnContextPopup(b._instance(), fn)
}

// SetOnDragDrop
//
// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (b *TButton) SetOnDragDrop(fn TDragDropEvent) {
    Button_SetOnDragDrop(b._instance(), fn)
}

// SetOnDragOver
//
// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (b *TButton) SetOnDragOver(fn TDragOverEvent) {
    Button_SetOnDragOver(b._instance(), fn)
}

// SetOnEndDrag
//
// 设置拖拽结束。
//
// Set End of drag.
func (b *TButton) SetOnEndDrag(fn TEndDragEvent) {
    Button_SetOnEndDrag(b._instance(), fn)
}

// SetOnEnter
//
// 设置焦点进入。
//
// Set Focus entry.
func (b *TButton) SetOnEnter(fn TNotifyEvent) {
    Button_SetOnEnter(b._instance(), fn)
}

// SetOnExit
//
// 设置焦点退出。
//
// Set Focus exit.
func (b *TButton) SetOnExit(fn TNotifyEvent) {
    Button_SetOnExit(b._instance(), fn)
}

// SetOnKeyDown
//
// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (b *TButton) SetOnKeyDown(fn TKeyEvent) {
    Button_SetOnKeyDown(b._instance(), fn)
}

// SetOnKeyPress
//
// 设置键键下事件。
func (b *TButton) SetOnKeyPress(fn TKeyPressEvent) {
    Button_SetOnKeyPress(b._instance(), fn)
}

// SetOnKeyUp
//
// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (b *TButton) SetOnKeyUp(fn TKeyEvent) {
    Button_SetOnKeyUp(b._instance(), fn)
}

// SetOnMouseDown
//
// 设置鼠标按下事件。
//
// Set Mouse down event.
func (b *TButton) SetOnMouseDown(fn TMouseEvent) {
    Button_SetOnMouseDown(b._instance(), fn)
}

// SetOnMouseEnter
//
// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (b *TButton) SetOnMouseEnter(fn TNotifyEvent) {
    Button_SetOnMouseEnter(b._instance(), fn)
}

// SetOnMouseLeave
//
// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (b *TButton) SetOnMouseLeave(fn TNotifyEvent) {
    Button_SetOnMouseLeave(b._instance(), fn)
}

// SetOnMouseMove
//
// 设置鼠标移动事件。
func (b *TButton) SetOnMouseMove(fn TMouseMoveEvent) {
    Button_SetOnMouseMove(b._instance(), fn)
}

// SetOnMouseUp
//
// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (b *TButton) SetOnMouseUp(fn TMouseEvent) {
    Button_SetOnMouseUp(b._instance(), fn)
}

// DockClientCount
//
// 获取依靠客户端总数。
func (b *TButton) DockClientCount() int32 {
    return Button_GetDockClientCount(b._instance())
}

// DockSite
//
// 获取停靠站点。
//
// Get Docking site.
func (b *TButton) DockSite() bool {
    return Button_GetDockSite(b._instance())
}

// SetDockSite
//
// 设置停靠站点。
//
// Set Docking site.
func (b *TButton) SetDockSite(value bool) {
    Button_SetDockSite(b._instance(), value)
}

// MouseInClient
//
// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (b *TButton) MouseInClient() bool {
    return Button_GetMouseInClient(b._instance())
}

// VisibleDockClientCount
//
// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (b *TButton) VisibleDockClientCount() int32 {
    return Button_GetVisibleDockClientCount(b._instance())
}

// Brush
//
// 获取画刷对象。
//
// Get Brush.
func (b *TButton) Brush() *TBrush {
    return AsBrush(Button_GetBrush(b._instance()))
}

// ControlCount
//
// 获取子控件数。
//
// Get Number of child controls.
func (b *TButton) ControlCount() int32 {
    return Button_GetControlCount(b._instance())
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (b *TButton) Handle() HWND {
    return Button_GetHandle(b._instance())
}

// ParentWindow
//
// 获取父容器句柄。
//
// Get Parent container handle.
func (b *TButton) ParentWindow() HWND {
    return Button_GetParentWindow(b._instance())
}

// SetParentWindow
//
// 设置父容器句柄。
//
// Set Parent container handle.
func (b *TButton) SetParentWindow(value HWND) {
    Button_SetParentWindow(b._instance(), value)
}

func (b *TButton) Showing() bool {
    return Button_GetShowing(b._instance())
}

// UseDockManager
//
// 获取使用停靠管理。
func (b *TButton) UseDockManager() bool {
    return Button_GetUseDockManager(b._instance())
}

// SetUseDockManager
//
// 设置使用停靠管理。
func (b *TButton) SetUseDockManager(value bool) {
    Button_SetUseDockManager(b._instance(), value)
}

func (b *TButton) BoundsRect() TRect {
    return Button_GetBoundsRect(b._instance())
}

func (b *TButton) SetBoundsRect(value TRect) {
    Button_SetBoundsRect(b._instance(), value)
}

// ClientHeight
//
// 获取客户区高度。
//
// Get client height.
func (b *TButton) ClientHeight() int32 {
    return Button_GetClientHeight(b._instance())
}

// SetClientHeight
//
// 设置客户区高度。
//
// Set client height.
func (b *TButton) SetClientHeight(value int32) {
    Button_SetClientHeight(b._instance(), value)
}

func (b *TButton) ClientOrigin() TPoint {
    return Button_GetClientOrigin(b._instance())
}

// ClientRect
//
// 获取客户区矩形。
//
// Get client rectangle.
func (b *TButton) ClientRect() TRect {
    return Button_GetClientRect(b._instance())
}

// ClientWidth
//
// 获取客户区宽度。
//
// Get client width.
func (b *TButton) ClientWidth() int32 {
    return Button_GetClientWidth(b._instance())
}

// SetClientWidth
//
// 设置客户区宽度。
//
// Set client width.
func (b *TButton) SetClientWidth(value int32) {
    Button_SetClientWidth(b._instance(), value)
}

// ControlState
//
// 获取控件状态。
//
// Get control state.
func (b *TButton) ControlState() TControlState {
    return Button_GetControlState(b._instance())
}

// SetControlState
//
// 设置控件状态。
//
// Set control state.
func (b *TButton) SetControlState(value TControlState) {
    Button_SetControlState(b._instance(), value)
}

// ControlStyle
//
// 获取控件样式。
//
// Get control style.
func (b *TButton) ControlStyle() TControlStyle {
    return Button_GetControlStyle(b._instance())
}

// SetControlStyle
//
// 设置控件样式。
//
// Set control style.
func (b *TButton) SetControlStyle(value TControlStyle) {
    Button_SetControlStyle(b._instance(), value)
}

func (b *TButton) Floating() bool {
    return Button_GetFloating(b._instance())
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (b *TButton) Parent() *TWinControl {
    return AsWinControl(Button_GetParent(b._instance()))
}

// SetParent
//
// 设置控件父容器。
//
// Set control parent container.
func (b *TButton) SetParent(value IWinControl) {
    Button_SetParent(b._instance(), CheckPtr(value))
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (b *TButton) Left() int32 {
    return Button_GetLeft(b._instance())
}

// SetLeft
//
// 设置左边位置。
//
// Set Left position.
func (b *TButton) SetLeft(value int32) {
    Button_SetLeft(b._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (b *TButton) Top() int32 {
    return Button_GetTop(b._instance())
}

// SetTop
//
// 设置顶边位置。
//
// Set Top position.
func (b *TButton) SetTop(value int32) {
    Button_SetTop(b._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (b *TButton) Width() int32 {
    return Button_GetWidth(b._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (b *TButton) SetWidth(value int32) {
    Button_SetWidth(b._instance(), value)
}

// Height
//
// 获取高度。
//
// Get height.
func (b *TButton) Height() int32 {
    return Button_GetHeight(b._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (b *TButton) SetHeight(value int32) {
    Button_SetHeight(b._instance(), value)
}

// Cursor
//
// 获取控件光标。
//
// Get control cursor.
func (b *TButton) Cursor() TCursor {
    return Button_GetCursor(b._instance())
}

// SetCursor
//
// 设置控件光标。
//
// Set control cursor.
func (b *TButton) SetCursor(value TCursor) {
    Button_SetCursor(b._instance(), value)
}

// Hint
//
// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (b *TButton) Hint() string {
    return Button_GetHint(b._instance())
}

// SetHint
//
// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (b *TButton) SetHint(value string) {
    Button_SetHint(b._instance(), value)
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (b *TButton) ComponentCount() int32 {
    return Button_GetComponentCount(b._instance())
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (b *TButton) ComponentIndex() int32 {
    return Button_GetComponentIndex(b._instance())
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (b *TButton) SetComponentIndex(value int32) {
    Button_SetComponentIndex(b._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (b *TButton) Owner() *TComponent {
    return AsComponent(Button_GetOwner(b._instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (b *TButton) Name() string {
    return Button_GetName(b._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (b *TButton) SetName(value string) {
    Button_SetName(b._instance(), value)
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (b *TButton) Tag() int {
    return Button_GetTag(b._instance())
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (b *TButton) SetTag(value int) {
    Button_SetTag(b._instance(), value)
}

// AnchorSideLeft
//
// 获取左边锚点。
func (b *TButton) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(Button_GetAnchorSideLeft(b._instance()))
}

// SetAnchorSideLeft
//
// 设置左边锚点。
func (b *TButton) SetAnchorSideLeft(value *TAnchorSide) {
    Button_SetAnchorSideLeft(b._instance(), CheckPtr(value))
}

// AnchorSideTop
//
// 获取顶边锚点。
func (b *TButton) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(Button_GetAnchorSideTop(b._instance()))
}

// SetAnchorSideTop
//
// 设置顶边锚点。
func (b *TButton) SetAnchorSideTop(value *TAnchorSide) {
    Button_SetAnchorSideTop(b._instance(), CheckPtr(value))
}

// AnchorSideRight
//
// 获取右边锚点。
func (b *TButton) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(Button_GetAnchorSideRight(b._instance()))
}

// SetAnchorSideRight
//
// 设置右边锚点。
func (b *TButton) SetAnchorSideRight(value *TAnchorSide) {
    Button_SetAnchorSideRight(b._instance(), CheckPtr(value))
}

// AnchorSideBottom
//
// 获取底边锚点。
func (b *TButton) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(Button_GetAnchorSideBottom(b._instance()))
}

// SetAnchorSideBottom
//
// 设置底边锚点。
func (b *TButton) SetAnchorSideBottom(value *TAnchorSide) {
    Button_SetAnchorSideBottom(b._instance(), CheckPtr(value))
}

func (b *TButton) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(Button_GetChildSizing(b._instance()))
}

func (b *TButton) SetChildSizing(value *TControlChildSizing) {
    Button_SetChildSizing(b._instance(), CheckPtr(value))
}

// BorderSpacing
//
// 获取边框间距。
func (b *TButton) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(Button_GetBorderSpacing(b._instance()))
}

// SetBorderSpacing
//
// 设置边框间距。
func (b *TButton) SetBorderSpacing(value *TControlBorderSpacing) {
    Button_SetBorderSpacing(b._instance(), CheckPtr(value))
}

// DockClients
//
// 获取指定索引停靠客户端。
func (b *TButton) DockClients(Index int32) *TControl {
    return AsControl(Button_GetDockClients(b._instance(), Index))
}

// Controls
//
// 获取指定索引子控件。
func (b *TButton) Controls(Index int32) *TControl {
    return AsControl(Button_GetControls(b._instance(), Index))
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (b *TButton) Components(AIndex int32) *TComponent {
    return AsComponent(Button_GetComponents(b._instance(), AIndex))
}

// AnchorSide
//
// 获取锚侧面。
func (b *TButton) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(Button_GetAnchorSide(b._instance(), AKind))
}

