
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

type TTreeNode struct {
    IObject
    instance unsafe.Pointer
}

// NewTreeNode
//
// 创建一个新的对象。
// 
// Create a new object.
func NewTreeNode(AOwner *TTreeNodes) *TTreeNode {
    t := new(TTreeNode)
    t.instance = unsafe.Pointer(TreeNode_Create(CheckPtr(AOwner)))
    setFinalizer(t, (*TTreeNode).Free)
    return t
}

// AsTreeNode
//
// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsTreeNode(obj interface{}) *TTreeNode {
    instance := getInstance(obj)
    if instance == nullptr { return nil }
    return &TTreeNode{instance: instance}
}

// Free 
//
// 释放对象。
// 
// Free object.
func (t *TTreeNode) Free() {
    if t.instance != nullptr {
        TreeNode_Free(t._instance())
        t.instance  = nullptr
    }
}

func (t *TTreeNode) _instance() uintptr {
    return uintptr(t.instance)
}

// Instance 
//
// 返回对象实例指针。
// 
// Return object instance pointer.
func (t *TTreeNode) Instance() uintptr {
    return t._instance()
}

// UnsafeAddr 
//
// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (t *TTreeNode) UnsafeAddr() unsafe.Pointer {
    return t.instance
}

// IsValid 
//
// 检测地址是否为空。
// 
// Check if the address is empty.
func (t *TTreeNode) IsValid() bool {
    return t.instance != nullptr
}

// Is 
// 
// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (t *TTreeNode) Is() TIs {
    return TIs(t._instance())
}

// As 
//
// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (t *TTreeNode) As() TAs {
//    return TAs(t._instance())
//}

// TTreeNodeClass
//
// 获取类信息指针。
// 
// Get class information pointer.
func TTreeNodeClass() TClass {
    return TreeNode_StaticClassType()
}

func (t *TTreeNode) Bottom() int32 {
    return TreeNode_Bottom(t._instance())
}

func (t *TTreeNode) BottomExpanded() int32 {
    return TreeNode_BottomExpanded(t._instance())
}

func (t *TTreeNode) DefaultTreeViewSort(Node1 *TTreeNode, Node2 *TTreeNode) int32 {
    return TreeNode_DefaultTreeViewSort(t._instance(), CheckPtr(Node1), CheckPtr(Node2))
}

func (t *TTreeNode) DisplayExpandSignLeft() int32 {
    return TreeNode_DisplayExpandSignLeft(t._instance())
}

func (t *TTreeNode) DisplayExpandSignRect() TRect {
    return TreeNode_DisplayExpandSignRect(t._instance())
}

func (t *TTreeNode) DisplayExpandSignRight() int32 {
    return TreeNode_DisplayExpandSignRight(t._instance())
}

func (t *TTreeNode) DisplayIconLeft() int32 {
    return TreeNode_DisplayIconLeft(t._instance())
}

func (t *TTreeNode) DisplayRect(TextOnly bool) TRect {
    return TreeNode_DisplayRect(t._instance(), TextOnly)
}

func (t *TTreeNode) DisplayStateIconLeft() int32 {
    return TreeNode_DisplayStateIconLeft(t._instance())
}

func (t *TTreeNode) DisplayTextLeft() int32 {
    return TreeNode_DisplayTextLeft(t._instance())
}

func (t *TTreeNode) DisplayTextRight() int32 {
    return TreeNode_DisplayTextRight(t._instance())
}

func (t *TTreeNode) EditText() bool {
    return TreeNode_EditText(t._instance())
}

func (t *TTreeNode) FindNode(NodeText string) *TTreeNode {
    return AsTreeNode(TreeNode_FindNode(t._instance(), NodeText))
}

func (t *TTreeNode) GetFirstChild() *TTreeNode {
    return AsTreeNode(TreeNode_GetFirstChild(t._instance()))
}

func (t *TTreeNode) GetFirstVisibleChild() *TTreeNode {
    return AsTreeNode(TreeNode_GetFirstVisibleChild(t._instance()))
}

func (t *TTreeNode) GetLastChild() *TTreeNode {
    return AsTreeNode(TreeNode_GetLastChild(t._instance()))
}

func (t *TTreeNode) GetLastSibling() *TTreeNode {
    return AsTreeNode(TreeNode_GetLastSibling(t._instance()))
}

func (t *TTreeNode) GetLastSubChild() *TTreeNode {
    return AsTreeNode(TreeNode_GetLastSubChild(t._instance()))
}

func (t *TTreeNode) GetLastVisibleChild() *TTreeNode {
    return AsTreeNode(TreeNode_GetLastVisibleChild(t._instance()))
}

func (t *TTreeNode) GetNext() *TTreeNode {
    return AsTreeNode(TreeNode_GetNext(t._instance()))
}

func (t *TTreeNode) GetNextChild(AValue *TTreeNode) *TTreeNode {
    return AsTreeNode(TreeNode_GetNextChild(t._instance(), CheckPtr(AValue)))
}

func (t *TTreeNode) GetNextExpanded() *TTreeNode {
    return AsTreeNode(TreeNode_GetNextExpanded(t._instance()))
}

func (t *TTreeNode) GetNextMultiSelected() *TTreeNode {
    return AsTreeNode(TreeNode_GetNextMultiSelected(t._instance()))
}

func (t *TTreeNode) GetNextSibling() *TTreeNode {
    return AsTreeNode(TreeNode_GetNextSibling(t._instance()))
}

func (t *TTreeNode) GetNextSkipChildren() *TTreeNode {
    return AsTreeNode(TreeNode_GetNextSkipChildren(t._instance()))
}

func (t *TTreeNode) GetNextVisible() *TTreeNode {
    return AsTreeNode(TreeNode_GetNextVisible(t._instance()))
}

func (t *TTreeNode) GetNextVisibleSibling() *TTreeNode {
    return AsTreeNode(TreeNode_GetNextVisibleSibling(t._instance()))
}

func (t *TTreeNode) GetParentNodeOfAbsoluteLevel(TheAbsoluteLevel int32) *TTreeNode {
    return AsTreeNode(TreeNode_GetParentNodeOfAbsoluteLevel(t._instance(), TheAbsoluteLevel))
}

func (t *TTreeNode) GetPrev() *TTreeNode {
    return AsTreeNode(TreeNode_GetPrev(t._instance()))
}

func (t *TTreeNode) GetPrevChild(AValue *TTreeNode) *TTreeNode {
    return AsTreeNode(TreeNode_GetPrevChild(t._instance(), CheckPtr(AValue)))
}

func (t *TTreeNode) GetPrevExpanded() *TTreeNode {
    return AsTreeNode(TreeNode_GetPrevExpanded(t._instance()))
}

func (t *TTreeNode) GetPrevMultiSelected() *TTreeNode {
    return AsTreeNode(TreeNode_GetPrevMultiSelected(t._instance()))
}

func (t *TTreeNode) GetPrevSibling() *TTreeNode {
    return AsTreeNode(TreeNode_GetPrevSibling(t._instance()))
}

func (t *TTreeNode) GetPrevVisible() *TTreeNode {
    return AsTreeNode(TreeNode_GetPrevVisible(t._instance()))
}

func (t *TTreeNode) GetPrevVisibleSibling() *TTreeNode {
    return AsTreeNode(TreeNode_GetPrevVisibleSibling(t._instance()))
}

func (t *TTreeNode) GetTextPath() string {
    return TreeNode_GetTextPath(t._instance())
}

func (t *TTreeNode) HasAsParent(AValue *TTreeNode) bool {
    return TreeNode_HasAsParent(t._instance(), CheckPtr(AValue))
}

func (t *TTreeNode) IndexOf(AValue *TTreeNode) int32 {
    return TreeNode_IndexOf(t._instance(), CheckPtr(AValue))
}

func (t *TTreeNode) IndexOfText(NodeText string) int32 {
    return TreeNode_IndexOfText(t._instance(), NodeText)
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (t *TTreeNode) Assign(Source IObject) {
    TreeNode_Assign(t._instance(), CheckPtr(Source))
}

func (t *TTreeNode) Collapse(Recurse bool) {
    TreeNode_Collapse(t._instance(), Recurse)
}

func (t *TTreeNode) ConsistencyCheck() {
    TreeNode_ConsistencyCheck(t._instance())
}

func (t *TTreeNode) Delete() {
    TreeNode_Delete(t._instance())
}

func (t *TTreeNode) DeleteChildren() {
    TreeNode_DeleteChildren(t._instance())
}

func (t *TTreeNode) EndEdit(Cancel bool) {
    TreeNode_EndEdit(t._instance(), Cancel)
}

func (t *TTreeNode) Expand(Recurse bool) {
    TreeNode_Expand(t._instance(), Recurse)
}

func (t *TTreeNode) ExpandParents() {
    TreeNode_ExpandParents(t._instance())
}

func (t *TTreeNode) FreeAllNodeData() {
    TreeNode_FreeAllNodeData(t._instance())
}

func (t *TTreeNode) MakeVisible() {
    TreeNode_MakeVisible(t._instance())
}

func (t *TTreeNode) MoveTo(Destination *TTreeNode, Mode TNodeAttachMode) {
    TreeNode_MoveTo(t._instance(), CheckPtr(Destination), Mode)
}

func (t *TTreeNode) MultiSelectGroup() {
    TreeNode_MultiSelectGroup(t._instance())
}

// Update
//
// 控件更新。
//
// Update.
func (t *TTreeNode) Update() {
    TreeNode_Update(t._instance())
}

func (t *TTreeNode) WriteDebugReport(Prefix string, Recurse bool) {
    TreeNode_WriteDebugReport(t._instance(), Prefix , Recurse)
}

func (t *TTreeNode) CustomSort(SortProc PFNTVCOMPARE, Data int, ARecurse bool) bool {
    return TreeNode_CustomSort(t._instance(), SortProc , Data , ARecurse)
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (t *TTreeNode) GetNamePath() string {
    return TreeNode_GetNamePath(t._instance())
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (t *TTreeNode) ClassType() TClass {
    return TreeNode_ClassType(t._instance())
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (t *TTreeNode) ClassName() string {
    return TreeNode_ClassName(t._instance())
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (t *TTreeNode) InstanceSize() int32 {
    return TreeNode_InstanceSize(t._instance())
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (t *TTreeNode) InheritsFrom(AClass TClass) bool {
    return TreeNode_InheritsFrom(t._instance(), AClass)
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (t *TTreeNode) Equals(Obj IObject) bool {
    return TreeNode_Equals(t._instance(), CheckPtr(Obj))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (t *TTreeNode) GetHashCode() int32 {
    return TreeNode_GetHashCode(t._instance())
}

// ToString
//
// 文本类信息。
//
// Text information.
func (t *TTreeNode) ToString() string {
    return TreeNode_ToString(t._instance())
}

func (t *TTreeNode) AbsoluteIndex() int32 {
    return TreeNode_GetAbsoluteIndex(t._instance())
}

func (t *TTreeNode) Count() int32 {
    return TreeNode_GetCount(t._instance())
}

func (t *TTreeNode) Cut() bool {
    return TreeNode_GetCut(t._instance())
}

func (t *TTreeNode) SetCut(value bool) {
    TreeNode_SetCut(t._instance(), value)
}

func (t *TTreeNode) Data() unsafe.Pointer {
    return TreeNode_GetData(t._instance())
}

func (t *TTreeNode) SetData(value unsafe.Pointer) {
    TreeNode_SetData(t._instance(), value)
}

func (t *TTreeNode) Deleting() bool {
    return TreeNode_GetDeleting(t._instance())
}

func (t *TTreeNode) DropTarget() bool {
    return TreeNode_GetDropTarget(t._instance())
}

func (t *TTreeNode) SetDropTarget(value bool) {
    TreeNode_SetDropTarget(t._instance(), value)
}

func (t *TTreeNode) Expanded() bool {
    return TreeNode_GetExpanded(t._instance())
}

func (t *TTreeNode) SetExpanded(value bool) {
    TreeNode_SetExpanded(t._instance(), value)
}

// Focused
//
// 获取返回是否获取焦点。
//
// Get Return to get focus.
func (t *TTreeNode) Focused() bool {
    return TreeNode_GetFocused(t._instance())
}

// SetFocused
//
// 设置返回是否获取焦点。
//
// Set Return to get focus.
func (t *TTreeNode) SetFocused(value bool) {
    TreeNode_SetFocused(t._instance(), value)
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (t *TTreeNode) Handle() uintptr {
    return TreeNode_GetHandle(t._instance())
}

func (t *TTreeNode) HasChildren() bool {
    return TreeNode_GetHasChildren(t._instance())
}

func (t *TTreeNode) SetHasChildren(value bool) {
    TreeNode_SetHasChildren(t._instance(), value)
}

// Height
//
// 获取高度。
//
// Get height.
func (t *TTreeNode) Height() int32 {
    return TreeNode_GetHeight(t._instance())
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (t *TTreeNode) SetHeight(value int32) {
    TreeNode_SetHeight(t._instance(), value)
}

// ImageIndex
//
// 获取图像在images中的索引。
func (t *TTreeNode) ImageIndex() int32 {
    return TreeNode_GetImageIndex(t._instance())
}

// SetImageIndex
//
// 设置图像在images中的索引。
func (t *TTreeNode) SetImageIndex(value int32) {
    TreeNode_SetImageIndex(t._instance(), value)
}

func (t *TTreeNode) Index() int32 {
    return TreeNode_GetIndex(t._instance())
}

func (t *TTreeNode) SetIndex(value int32) {
    TreeNode_SetIndex(t._instance(), value)
}

func (t *TTreeNode) IsFullHeightVisible() bool {
    return TreeNode_GetIsFullHeightVisible(t._instance())
}

func (t *TTreeNode) IsVisible() bool {
    return TreeNode_GetIsVisible(t._instance())
}

func (t *TTreeNode) Level() int32 {
    return TreeNode_GetLevel(t._instance())
}

func (t *TTreeNode) MultiSelected() bool {
    return TreeNode_GetMultiSelected(t._instance())
}

func (t *TTreeNode) SetMultiSelected(value bool) {
    TreeNode_SetMultiSelected(t._instance(), value)
}

func (t *TTreeNode) OverlayIndex() int32 {
    return TreeNode_GetOverlayIndex(t._instance())
}

func (t *TTreeNode) SetOverlayIndex(value int32) {
    TreeNode_SetOverlayIndex(t._instance(), value)
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (t *TTreeNode) Owner() *TTreeNodes {
    return AsTreeNodes(TreeNode_GetOwner(t._instance()))
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (t *TTreeNode) Parent() *TTreeNode {
    return AsTreeNode(TreeNode_GetParent(t._instance()))
}

func (t *TTreeNode) Selected() bool {
    return TreeNode_GetSelected(t._instance())
}

func (t *TTreeNode) SetSelected(value bool) {
    TreeNode_SetSelected(t._instance(), value)
}

func (t *TTreeNode) SelectedIndex() int32 {
    return TreeNode_GetSelectedIndex(t._instance())
}

func (t *TTreeNode) SetSelectedIndex(value int32) {
    TreeNode_SetSelectedIndex(t._instance(), value)
}

func (t *TTreeNode) StateIndex() int32 {
    return TreeNode_GetStateIndex(t._instance())
}

func (t *TTreeNode) SetStateIndex(value int32) {
    TreeNode_SetStateIndex(t._instance(), value)
}

func (t *TTreeNode) SubTreeCount() int32 {
    return TreeNode_GetSubTreeCount(t._instance())
}

// Text
//
// 获取文本。
func (t *TTreeNode) Text() string {
    return TreeNode_GetText(t._instance())
}

// SetText
//
// 设置文本。
func (t *TTreeNode) SetText(value string) {
    TreeNode_SetText(t._instance(), value)
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (t *TTreeNode) Top() int32 {
    return TreeNode_GetTop(t._instance())
}

func (t *TTreeNode) TreeNodes() *TTreeNodes {
    return AsTreeNodes(TreeNode_GetTreeNodes(t._instance()))
}

func (t *TTreeNode) TreeView() *TWinControl {
    return AsWinControl(TreeNode_GetTreeView(t._instance()))
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (t *TTreeNode) Visible() bool {
    return TreeNode_GetVisible(t._instance())
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (t *TTreeNode) SetVisible(value bool) {
    TreeNode_SetVisible(t._instance(), value)
}

func (t *TTreeNode) Items(ItemIndex int32) *TTreeNode {
    return AsTreeNode(TreeNode_GetItems(t._instance(), ItemIndex))
}

func (t *TTreeNode) SetItems(ItemIndex int32, value *TTreeNode) {
    TreeNode_SetItems(t._instance(), ItemIndex, CheckPtr(value))
}

func (t *TTreeNode) Item(Index int32) *TTreeNode {
    return AsTreeNode(TreeNode_GetItem(t._instance(), Index))
}

func (t *TTreeNode) SetItem(Index int32, value *TTreeNode) {
    TreeNode_SetItem(t._instance(), Index, CheckPtr(value))
}

