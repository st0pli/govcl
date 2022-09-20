
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

type TColorButton struct {
    IControl
    instance unsafe.Pointer
}

// NewColorButton
//
// 创建一个新的对象。
// 
// Create a new object.
func NewColorButton(owner IComponent) *TColorButton {
    c := new(TColorButton)
    c.instance = unsafe.Pointer(ColorButton_Create(CheckPtr(owner)))
    return c
}

// AsColorButton
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsColorButton(obj interface{}) *TColorButton {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TColorButton{instance: instance}
}

// Free 
//
// 释放对象。
// 
// Free object.
func (c *TColorButton) Free() {
    if c.instance != nullptr {
        ColorButton_Free(c._instance())
        c.instance  = nullptr
    }
}

func (c *TColorButton) _instance() uintptr {
    return uintptr(c.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (c *TColorButton) Instance() uintptr {
    return c._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (c *TColorButton) UnsafeAddr() unsafe.Pointer {
    return c.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (c *TColorButton) IsValid() bool {
    return c.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (c *TColorButton) Is() TIs {
    return TIs(c._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (c *TColorButton) As() TAs {
//    return TAs(c._instance())
//}

// TColorButtonClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TColorButtonClass() TClass {
    return ColorButton_StaticClassType()
}

// Click
//
// 单击。
func (c *TColorButton) Click() {
    ColorButton_Click(c._instance())
}

// BringToFront
//
// 将控件置于最前。
//
// Bring the control to the front.
func (c *TColorButton) BringToFront() {
    ColorButton_BringToFront(c._instance())
}

// ClientToScreen
//
// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (c *TColorButton) ClientToScreen(Point TPoint) TPoint {
    return ColorButton_ClientToScreen(c._instance(), Point)
}

// ClientToParent
//
// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (c *TColorButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ColorButton_ClientToParent(c._instance(), Point , CheckPtr(AParent))
}

// Dragging
//
// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (c *TColorButton) Dragging() bool {
    return ColorButton_Dragging(c._instance())
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (c *TColorButton) HasParent() bool {
    return ColorButton_HasParent(c._instance())
}

// Hide
//
// 隐藏控件。
//
// Hidden control.
func (c *TColorButton) Hide() {
    ColorButton_Hide(c._instance())
}

// Invalidate
//
// 要求重绘。
//
// Redraw.
func (c *TColorButton) Invalidate() {
    ColorButton_Invalidate(c._instance())
}

// Perform
//
// 发送一个消息。
//
// Send a message.
func (c *TColorButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ColorButton_Perform(c._instance(), Msg , WParam , LParam)
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (c *TColorButton) Refresh() {
    ColorButton_Refresh(c._instance())
}

// Repaint
//
// 重绘。
//
// Repaint.
func (c *TColorButton) Repaint() {
    ColorButton_Repaint(c._instance())
}

// ScreenToClient
//
// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (c *TColorButton) ScreenToClient(Point TPoint) TPoint {
    return ColorButton_ScreenToClient(c._instance(), Point)
}

// ParentToClient
//
// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (c *TColorButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ColorButton_ParentToClient(c._instance(), Point , CheckPtr(AParent))
}

// SendToBack
//
// 控件至于最后面。
//
// The control is placed at the end.
func (c *TColorButton) SendToBack() {
    ColorButton_SendToBack(c._instance())
}

// SetBounds
//
// 设置组件边界。
//
// Set component boundaries.
func (c *TColorButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ColorButton_SetBounds(c._instance(), ALeft , ATop , AWidth , AHeight)
}

// Show
//
// 显示控件。
//
// Show control.
func (c *TColorButton) Show() {
    ColorButton_Show(c._instance())
}

// Update
//
// 控件更新。
//
// Update.
func (c *TColorButton) Update() {
    ColorButton_Update(c._instance())
}

// GetTextBuf
//
// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (c *TColorButton) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return ColorButton_GetTextBuf(c._instance(), Buffer , BufSize)
}

// GetTextLen
//
// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (c *TColorButton) GetTextLen() int32 {
    return ColorButton_GetTextLen(c._instance())
}

// SetTextBuf
//
// 设置控件字符，如果有。
//
// Set control characters, if any.
func (c *TColorButton) SetTextBuf(Buffer string) {
    ColorButton_SetTextBuf(c._instance(), Buffer)
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (c *TColorButton) FindComponent(AName string) *TComponent {
    return AsComponent(ColorButton_FindComponent(c._instance(), AName))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (c *TColorButton) GetNamePath() string {
    return ColorButton_GetNamePath(c._instance())
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TColorButton) Assign(Source IObject) {
    ColorButton_Assign(c._instance(), CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (c *TColorButton) ClassType() TClass {
    return ColorButton_ClassType(c._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TColorButton) ClassName() string {
    return ColorButton_ClassName(c._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TColorButton) InstanceSize() int32 {
    return ColorButton_InstanceSize(c._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TColorButton) InheritsFrom(AClass TClass) bool {
    return ColorButton_InheritsFrom(c._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (c *TColorButton) Equals(Obj IObject) bool {
    return ColorButton_Equals(c._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TColorButton) GetHashCode() int32 {
    return ColorButton_GetHashCode(c._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (c *TColorButton) ToString() string {
    return ColorButton_ToString(c._instance())
}

func (c *TColorButton) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ColorButton_AnchorToNeighbour(c._instance(), ASide , ASpace , CheckPtr(ASibling))
}

func (c *TColorButton) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ColorButton_AnchorParallel(c._instance(), ASide , ASpace , CheckPtr(ASibling))
}

// AnchorHorizontalCenterTo
//
// 置于指定控件的横向中心。
func (c *TColorButton) AnchorHorizontalCenterTo(ASibling IControl) {
    ColorButton_AnchorHorizontalCenterTo(c._instance(), CheckPtr(ASibling))
}

// AnchorVerticalCenterTo
//
// 置于指定控件的纵向中心。
func (c *TColorButton) AnchorVerticalCenterTo(ASibling IControl) {
    ColorButton_AnchorVerticalCenterTo(c._instance(), CheckPtr(ASibling))
}

func (c *TColorButton) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    ColorButton_AnchorSame(c._instance(), ASide , CheckPtr(ASibling))
}

func (c *TColorButton) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    ColorButton_AnchorAsAlign(c._instance(), ATheAlign , ASpace)
}

func (c *TColorButton) AnchorClient(ASpace int32) {
    ColorButton_AnchorClient(c._instance(), ASpace)
}

func (c *TColorButton) ScaleDesignToForm(ASize int32) int32 {
    return ColorButton_ScaleDesignToForm(c._instance(), ASize)
}

func (c *TColorButton) ScaleFormToDesign(ASize int32) int32 {
    return ColorButton_ScaleFormToDesign(c._instance(), ASize)
}

func (c *TColorButton) Scale96ToForm(ASize int32) int32 {
    return ColorButton_Scale96ToForm(c._instance(), ASize)
}

func (c *TColorButton) ScaleFormTo96(ASize int32) int32 {
    return ColorButton_ScaleFormTo96(c._instance(), ASize)
}

func (c *TColorButton) Scale96ToFont(ASize int32) int32 {
    return ColorButton_Scale96ToFont(c._instance(), ASize)
}

func (c *TColorButton) ScaleFontTo96(ASize int32) int32 {
    return ColorButton_ScaleFontTo96(c._instance(), ASize)
}

func (c *TColorButton) ScaleScreenToFont(ASize int32) int32 {
    return ColorButton_ScaleScreenToFont(c._instance(), ASize)
}

func (c *TColorButton) ScaleFontToScreen(ASize int32) int32 {
    return ColorButton_ScaleFontToScreen(c._instance(), ASize)
}

func (c *TColorButton) Scale96ToScreen(ASize int32) int32 {
    return ColorButton_Scale96ToScreen(c._instance(), ASize)
}

func (c *TColorButton) ScaleScreenTo96(ASize int32) int32 {
    return ColorButton_ScaleScreenTo96(c._instance(), ASize)
}

func (c *TColorButton) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    ColorButton_AutoAdjustLayout(c._instance(), AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (c *TColorButton) FixDesignFontsPPI(ADesignTimePPI int32) {
    ColorButton_FixDesignFontsPPI(c._instance(), ADesignTimePPI)
}

func (c *TColorButton) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    ColorButton_ScaleFontsPPI(c._instance(), AToPPI , AProportion)
}

func (c *TColorButton) Action() *TAction {
    return AsAction(ColorButton_GetAction(c._instance()))
}

func (c *TColorButton) SetAction(value IComponent) {
    ColorButton_SetAction(c._instance(), CheckPtr(value))
}

// Align
//
// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (c *TColorButton) Align() TAlign {
    return ColorButton_GetAlign(c._instance())
}

// SetAlign
//
// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (c *TColorButton) SetAlign(value TAlign) {
    ColorButton_SetAlign(c._instance(), value)
}

// Anchors
//
// 获取四个角位置的锚点。
func (c *TColorButton) Anchors() TAnchors {
    return ColorButton_GetAnchors(c._instance())
}

// SetAnchors
//
// 设置四个角位置的锚点。
func (c *TColorButton) SetAnchors(value TAnchors) {
    ColorButton_SetAnchors(c._instance(), value)
}

func (c *TColorButton) AllowAllUp() bool {
    return ColorButton_GetAllowAllUp(c._instance())
}

func (c *TColorButton) SetAllowAllUp(value bool) {
    ColorButton_SetAllowAllUp(c._instance(), value)
}

// BorderWidth
//
// 获取边框的宽度。
func (c *TColorButton) BorderWidth() int32 {
    return ColorButton_GetBorderWidth(c._instance())
}

// SetBorderWidth
//
// 设置边框的宽度。
func (c *TColorButton) SetBorderWidth(value int32) {
    ColorButton_SetBorderWidth(c._instance(), value)
}

func (c *TColorButton) ButtonColorAutoSize() bool {
    return ColorButton_GetButtonColorAutoSize(c._instance())
}

func (c *TColorButton) SetButtonColorAutoSize(value bool) {
    ColorButton_SetButtonColorAutoSize(c._instance(), value)
}

func (c *TColorButton) ButtonColorSize() int32 {
    return ColorButton_GetButtonColorSize(c._instance())
}

func (c *TColorButton) SetButtonColorSize(value int32) {
    ColorButton_SetButtonColorSize(c._instance(), value)
}

func (c *TColorButton) ButtonColor() TColor {
    return ColorButton_GetButtonColor(c._instance())
}

func (c *TColorButton) SetButtonColor(value TColor) {
    ColorButton_SetButtonColor(c._instance(), value)
}

func (c *TColorButton) ColorDialog() *TColorDialog {
    return AsColorDialog(ColorButton_GetColorDialog(c._instance()))
}

func (c *TColorButton) SetColorDialog(value IComponent) {
    ColorButton_SetColorDialog(c._instance(), CheckPtr(value))
}

// Constraints
//
// 获取约束控件大小。
func (c *TColorButton) Constraints() *TSizeConstraints {
    return AsSizeConstraints(ColorButton_GetConstraints(c._instance()))
}

// SetConstraints
//
// 设置约束控件大小。
func (c *TColorButton) SetConstraints(value *TSizeConstraints) {
    ColorButton_SetConstraints(c._instance(), CheckPtr(value))
}

// Caption
//
// 获取控件标题。
//
// Get the control title.
func (c *TColorButton) Caption() string {
    return ColorButton_GetCaption(c._instance())
}

// SetCaption
//
// 设置控件标题。
//
// Set the control title.
func (c *TColorButton) SetCaption(value string) {
    ColorButton_SetCaption(c._instance(), value)
}

// Color
//
// 获取颜色。
//
// Get color.
func (c *TColorButton) Color() TColor {
    return ColorButton_GetColor(c._instance())
}

// SetColor
//
// 设置颜色。
//
// Set color.
func (c *TColorButton) SetColor(value TColor) {
    ColorButton_SetColor(c._instance(), value)
}

func (c *TColorButton) Down() bool {
    return ColorButton_GetDown(c._instance())
}

func (c *TColorButton) SetDown(value bool) {
    ColorButton_SetDown(c._instance(), value)
}

// Enabled
//
// 获取控件启用。
//
// Get the control enabled.
func (c *TColorButton) Enabled() bool {
    return ColorButton_GetEnabled(c._instance())
}

// SetEnabled
//
// 设置控件启用。
//
// Set the control enabled.
func (c *TColorButton) SetEnabled(value bool) {
    ColorButton_SetEnabled(c._instance(), value)
}

// Flat
//
// 获取平面样式。
func (c *TColorButton) Flat() bool {
    return ColorButton_GetFlat(c._instance())
}

// SetFlat
//
// 设置平面样式。
func (c *TColorButton) SetFlat(value bool) {
    ColorButton_SetFlat(c._instance(), value)
}

// Font
//
// 获取字体。
//
// Get Font.
func (c *TColorButton) Font() *TFont {
    return AsFont(ColorButton_GetFont(c._instance()))
}

// SetFont
//
// 设置字体。
//
// Set Font.
func (c *TColorButton) SetFont(value *TFont) {
    ColorButton_SetFont(c._instance(), CheckPtr(value))
}

// GroupIndex
//
// 获取团组索引。
func (c *TColorButton) GroupIndex() int32 {
    return ColorButton_GetGroupIndex(c._instance())
}

// SetGroupIndex
//
// 设置团组索引。
func (c *TColorButton) SetGroupIndex(value int32) {
    ColorButton_SetGroupIndex(c._instance(), value)
}

// Hint
//
// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (c *TColorButton) Hint() string {
    return ColorButton_GetHint(c._instance())
}

// SetHint
//
// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (c *TColorButton) SetHint(value string) {
    ColorButton_SetHint(c._instance(), value)
}

func (c *TColorButton) Layout() TButtonLayout {
    return ColorButton_GetLayout(c._instance())
}

func (c *TColorButton) SetLayout(value TButtonLayout) {
    ColorButton_SetLayout(c._instance(), value)
}

func (c *TColorButton) Spacing() int32 {
    return ColorButton_GetSpacing(c._instance())
}

func (c *TColorButton) SetSpacing(value int32) {
    ColorButton_SetSpacing(c._instance(), value)
}

// Transparent
//
// 获取透明。
//
// Get transparent.
func (c *TColorButton) Transparent() bool {
    return ColorButton_GetTransparent(c._instance())
}

// SetTransparent
//
// 设置透明。
//
// Set transparent.
func (c *TColorButton) SetTransparent(value bool) {
    ColorButton_SetTransparent(c._instance(), value)
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (c *TColorButton) Visible() bool {
    return ColorButton_GetVisible(c._instance())
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (c *TColorButton) SetVisible(value bool) {
    ColorButton_SetVisible(c._instance(), value)
}

// SetOnClick
//
// 设置控件单击事件。
//
// Set control click event.
func (c *TColorButton) SetOnClick(fn TNotifyEvent) {
    ColorButton_SetOnClick(c._instance(), fn)
}

func (c *TColorButton) SetOnColorChanged(fn TNotifyEvent) {
    ColorButton_SetOnColorChanged(c._instance(), fn)
}

// SetOnDblClick
//
// 设置双击事件。
func (c *TColorButton) SetOnDblClick(fn TNotifyEvent) {
    ColorButton_SetOnDblClick(c._instance(), fn)
}

// SetOnMouseDown
//
// 设置鼠标按下事件。
//
// Set Mouse down event.
func (c *TColorButton) SetOnMouseDown(fn TMouseEvent) {
    ColorButton_SetOnMouseDown(c._instance(), fn)
}

// SetOnMouseEnter
//
// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (c *TColorButton) SetOnMouseEnter(fn TNotifyEvent) {
    ColorButton_SetOnMouseEnter(c._instance(), fn)
}

// SetOnMouseLeave
//
// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (c *TColorButton) SetOnMouseLeave(fn TNotifyEvent) {
    ColorButton_SetOnMouseLeave(c._instance(), fn)
}

// SetOnMouseMove
//
// 设置鼠标移动事件。
func (c *TColorButton) SetOnMouseMove(fn TMouseMoveEvent) {
    ColorButton_SetOnMouseMove(c._instance(), fn)
}

// SetOnMouseUp
//
// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (c *TColorButton) SetOnMouseUp(fn TMouseEvent) {
    ColorButton_SetOnMouseUp(c._instance(), fn)
}

// SetOnMouseWheel
//
// 设置鼠标滚轮事件。
func (c *TColorButton) SetOnMouseWheel(fn TMouseWheelEvent) {
    ColorButton_SetOnMouseWheel(c._instance(), fn)
}

// SetOnMouseWheelDown
//
// 设置鼠标滚轮按下事件。
func (c *TColorButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    ColorButton_SetOnMouseWheelDown(c._instance(), fn)
}

// SetOnMouseWheelUp
//
// 设置鼠标滚轮抬起事件。
func (c *TColorButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    ColorButton_SetOnMouseWheelUp(c._instance(), fn)
}

// SetOnPaint
//
// 设置绘画事件。
func (c *TColorButton) SetOnPaint(fn TNotifyEvent) {
    ColorButton_SetOnPaint(c._instance(), fn)
}

// SetOnResize
//
// 设置大小被改变事件。
func (c *TColorButton) SetOnResize(fn TNotifyEvent) {
    ColorButton_SetOnResize(c._instance(), fn)
}

// ShowHint
//
// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (c *TColorButton) ShowHint() bool {
    return ColorButton_GetShowHint(c._instance())
}

// SetShowHint
//
// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (c *TColorButton) SetShowHint(value bool) {
    ColorButton_SetShowHint(c._instance(), value)
}

// ParentFont
//
// 获取使用父容器字体。
//
// Get Parent container font.
func (c *TColorButton) ParentFont() bool {
    return ColorButton_GetParentFont(c._instance())
}

// SetParentFont
//
// 设置使用父容器字体。
//
// Set Parent container font.
func (c *TColorButton) SetParentFont(value bool) {
    ColorButton_SetParentFont(c._instance(), value)
}

// ParentShowHint
//
// 获取以父容器的ShowHint属性为准。
func (c *TColorButton) ParentShowHint() bool {
    return ColorButton_GetParentShowHint(c._instance())
}

// SetParentShowHint
//
// 设置以父容器的ShowHint属性为准。
func (c *TColorButton) SetParentShowHint(value bool) {
    ColorButton_SetParentShowHint(c._instance(), value)
}

// PopupMenu
//
// 获取右键菜单。
//
// Get Right click menu.
func (c *TColorButton) PopupMenu() *TPopupMenu {
    return AsPopupMenu(ColorButton_GetPopupMenu(c._instance()))
}

// SetPopupMenu
//
// 设置右键菜单。
//
// Set Right click menu.
func (c *TColorButton) SetPopupMenu(value IComponent) {
    ColorButton_SetPopupMenu(c._instance(), CheckPtr(value))
}

// ImageIndex
//
// 获取图像在images中的索引。
func (c *TColorButton) ImageIndex() int32 {
    return ColorButton_GetImageIndex(c._instance())
}

// SetImageIndex
//
// 设置图像在images中的索引。
func (c *TColorButton) SetImageIndex(value int32) {
    ColorButton_SetImageIndex(c._instance(), value)
}

// Images
//
// 获取图标索引列表对象。
func (c *TColorButton) Images() *TImageList {
    return AsImageList(ColorButton_GetImages(c._instance()))
}

// SetImages
//
// 设置图标索引列表对象。
func (c *TColorButton) SetImages(value IComponent) {
    ColorButton_SetImages(c._instance(), CheckPtr(value))
}

func (c *TColorButton) ImageWidth() int32 {
    return ColorButton_GetImageWidth(c._instance())
}

func (c *TColorButton) SetImageWidth(value int32) {
    ColorButton_SetImageWidth(c._instance(), value)
}

func (c *TColorButton) ShowCaption() bool {
    return ColorButton_GetShowCaption(c._instance())
}

func (c *TColorButton) SetShowCaption(value bool) {
    ColorButton_SetShowCaption(c._instance(), value)
}

func (c *TColorButton) BiDiMode() TBiDiMode {
    return ColorButton_GetBiDiMode(c._instance())
}

func (c *TColorButton) SetBiDiMode(value TBiDiMode) {
    ColorButton_SetBiDiMode(c._instance(), value)
}

func (c *TColorButton) Glyph() *TBitmap {
    return AsBitmap(ColorButton_GetGlyph(c._instance()))
}

func (c *TColorButton) SetGlyph(value *TBitmap) {
    ColorButton_SetGlyph(c._instance(), CheckPtr(value))
}

func (c *TColorButton) NumGlyphs() TNumGlyphs {
    return ColorButton_GetNumGlyphs(c._instance())
}

func (c *TColorButton) SetNumGlyphs(value TNumGlyphs) {
    ColorButton_SetNumGlyphs(c._instance(), value)
}

func (c *TColorButton) BoundsRect() TRect {
    return ColorButton_GetBoundsRect(c._instance())
}

func (c *TColorButton) SetBoundsRect(value TRect) {
    ColorButton_SetBoundsRect(c._instance(), value)
}

// ClientHeight
//
// 获取客户区高度。
//
// Get client height.
func (c *TColorButton) ClientHeight() int32 {
    return ColorButton_GetClientHeight(c._instance())
}

// SetClientHeight
//
// 设置客户区高度。
//
// Set client height.
func (c *TColorButton) SetClientHeight(value int32) {
    ColorButton_SetClientHeight(c._instance(), value)
}

func (c *TColorButton) ClientOrigin() TPoint {
    return ColorButton_GetClientOrigin(c._instance())
}

// ClientRect
//
// 获取客户区矩形。
//
// Get client rectangle.
func (c *TColorButton) ClientRect() TRect {
    return ColorButton_GetClientRect(c._instance())
}

// ClientWidth
//
// 获取客户区宽度。
//
// Get client width.
func (c *TColorButton) ClientWidth() int32 {
    return ColorButton_GetClientWidth(c._instance())
}

// SetClientWidth
//
// 设置客户区宽度。
//
// Set client width.
func (c *TColorButton) SetClientWidth(value int32) {
    ColorButton_SetClientWidth(c._instance(), value)
}

// ControlState
//
// 获取控件状态。
//
// Get control state.
func (c *TColorButton) ControlState() TControlState {
    return ColorButton_GetControlState(c._instance())
}

// SetControlState
//
// 设置控件状态。
//
// Set control state.
func (c *TColorButton) SetControlState(value TControlState) {
    ColorButton_SetControlState(c._instance(), value)
}

// ControlStyle
//
// 获取控件样式。
//
// Get control style.
func (c *TColorButton) ControlStyle() TControlStyle {
    return ColorButton_GetControlStyle(c._instance())
}

// SetControlStyle
//
// 设置控件样式。
//
// Set control style.
func (c *TColorButton) SetControlStyle(value TControlStyle) {
    ColorButton_SetControlStyle(c._instance(), value)
}

func (c *TColorButton) Floating() bool {
    return ColorButton_GetFloating(c._instance())
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (c *TColorButton) Parent() *TWinControl {
    return AsWinControl(ColorButton_GetParent(c._instance()))
}

// SetParent
//
// 设置控件父容器。
//
// Set control parent container.
func (c *TColorButton) SetParent(value IWinControl) {
    ColorButton_SetParent(c._instance(), CheckPtr(value))
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (c *TColorButton) Left() int32 {
    return ColorButton_GetLeft(c._instance())
}

// SetLeft
//
// 设置左边位置。
//
// Set Left position.
func (c *TColorButton) SetLeft(value int32) {
    ColorButton_SetLeft(c._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (c *TColorButton) Top() int32 {
    return ColorButton_GetTop(c._instance())
}

// SetTop
//
// 设置顶边位置。
//
// Set Top position.
func (c *TColorButton) SetTop(value int32) {
    ColorButton_SetTop(c._instance(), value)
}

// Width
//
// 获取宽度。
//
// Get width.
func (c *TColorButton) Width() int32 {
    return ColorButton_GetWidth(c._instance())
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (c *TColorButton) SetWidth(value int32) {
    ColorButton_SetWidth(c._instance(), value)
}

// Height
//
// 获取高度。
//
// Get height.
func (c *TColorButton) Height() int32 {
    return ColorButton_GetHeight(c._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (c *TColorButton) SetHeight(value int32) {
    ColorButton_SetHeight(c._instance(), value)
}

// Cursor
//
// 获取控件光标。
//
// Get control cursor.
func (c *TColorButton) Cursor() TCursor {
    return ColorButton_GetCursor(c._instance())
}

// SetCursor
//
// 设置控件光标。
//
// Set control cursor.
func (c *TColorButton) SetCursor(value TCursor) {
    ColorButton_SetCursor(c._instance(), value)
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (c *TColorButton) ComponentCount() int32 {
    return ColorButton_GetComponentCount(c._instance())
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (c *TColorButton) ComponentIndex() int32 {
    return ColorButton_GetComponentIndex(c._instance())
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (c *TColorButton) SetComponentIndex(value int32) {
    ColorButton_SetComponentIndex(c._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (c *TColorButton) Owner() *TComponent {
    return AsComponent(ColorButton_GetOwner(c._instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (c *TColorButton) Name() string {
    return ColorButton_GetName(c._instance())
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (c *TColorButton) SetName(value string) {
    ColorButton_SetName(c._instance(), value)
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (c *TColorButton) Tag() int {
    return ColorButton_GetTag(c._instance())
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (c *TColorButton) SetTag(value int) {
    ColorButton_SetTag(c._instance(), value)
}

// AnchorSideLeft
//
// 获取左边锚点。
func (c *TColorButton) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(ColorButton_GetAnchorSideLeft(c._instance()))
}

// SetAnchorSideLeft
//
// 设置左边锚点。
func (c *TColorButton) SetAnchorSideLeft(value *TAnchorSide) {
    ColorButton_SetAnchorSideLeft(c._instance(), CheckPtr(value))
}

// AnchorSideTop
//
// 获取顶边锚点。
func (c *TColorButton) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(ColorButton_GetAnchorSideTop(c._instance()))
}

// SetAnchorSideTop
//
// 设置顶边锚点。
func (c *TColorButton) SetAnchorSideTop(value *TAnchorSide) {
    ColorButton_SetAnchorSideTop(c._instance(), CheckPtr(value))
}

// AnchorSideRight
//
// 获取右边锚点。
func (c *TColorButton) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(ColorButton_GetAnchorSideRight(c._instance()))
}

// SetAnchorSideRight
//
// 设置右边锚点。
func (c *TColorButton) SetAnchorSideRight(value *TAnchorSide) {
    ColorButton_SetAnchorSideRight(c._instance(), CheckPtr(value))
}

// AnchorSideBottom
//
// 获取底边锚点。
func (c *TColorButton) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(ColorButton_GetAnchorSideBottom(c._instance()))
}

// SetAnchorSideBottom
//
// 设置底边锚点。
func (c *TColorButton) SetAnchorSideBottom(value *TAnchorSide) {
    ColorButton_SetAnchorSideBottom(c._instance(), CheckPtr(value))
}

// BorderSpacing
//
// 获取边框间距。
func (c *TColorButton) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(ColorButton_GetBorderSpacing(c._instance()))
}

// SetBorderSpacing
//
// 设置边框间距。
func (c *TColorButton) SetBorderSpacing(value *TControlBorderSpacing) {
    ColorButton_SetBorderSpacing(c._instance(), CheckPtr(value))
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (c *TColorButton) Components(AIndex int32) *TComponent {
    return AsComponent(ColorButton_GetComponents(c._instance(), AIndex))
}

// AnchorSide
//
// 获取锚侧面。
func (c *TColorButton) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(ColorButton_GetAnchorSide(c._instance(), AKind))
}

