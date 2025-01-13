package binexport

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BinExport2_CallGraph_Vertex_Type int32

const (
	// Regular function with full disassembly.
	BinExport2_CallGraph_Vertex_NORMAL BinExport2_CallGraph_Vertex_Type = 0
	// This function is a well known library function.
	BinExport2_CallGraph_Vertex_LIBRARY BinExport2_CallGraph_Vertex_Type = 1
	// Imported from a dynamic link library (e.g. dll).
	BinExport2_CallGraph_Vertex_IMPORTED BinExport2_CallGraph_Vertex_Type = 2
	// A thunk function, forwarding its work via an unconditional jump.
	BinExport2_CallGraph_Vertex_THUNK BinExport2_CallGraph_Vertex_Type = 3
	// An invalid function (a function that contained invalid code or was
	// considered invalid by some heuristics).
	BinExport2_CallGraph_Vertex_INVALID BinExport2_CallGraph_Vertex_Type = 4
)

// Enum value maps for BinExport2_CallGraph_Vertex_Type.
var (
	BinExport2_CallGraph_Vertex_Type_name = map[int32]string{
		0: "NORMAL",
		1: "LIBRARY",
		2: "IMPORTED",
		3: "THUNK",
		4: "INVALID",
	}
	BinExport2_CallGraph_Vertex_Type_value = map[string]int32{
		"NORMAL":   0,
		"LIBRARY":  1,
		"IMPORTED": 2,
		"THUNK":    3,
		"INVALID":  4,
	}
)

func (x BinExport2_CallGraph_Vertex_Type) Enum() *BinExport2_CallGraph_Vertex_Type {
	p := new(BinExport2_CallGraph_Vertex_Type)
	*p = x
	return p
}

func (x BinExport2_CallGraph_Vertex_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BinExport2_CallGraph_Vertex_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_binexport2_proto_enumTypes[0].Descriptor()
}

func (BinExport2_CallGraph_Vertex_Type) Type() protoreflect.EnumType {
	return &file_binexport2_proto_enumTypes[0]
}

func (x BinExport2_CallGraph_Vertex_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *BinExport2_CallGraph_Vertex_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = BinExport2_CallGraph_Vertex_Type(num)
	return nil
}

// Deprecated: Use BinExport2_CallGraph_Vertex_Type.Descriptor instead.
func (BinExport2_CallGraph_Vertex_Type) EnumDescriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 1, 0, 0}
}

type BinExport2_Expression_Type int32

const (
	BinExport2_Expression_SYMBOL          BinExport2_Expression_Type = 1
	BinExport2_Expression_IMMEDIATE_INT   BinExport2_Expression_Type = 2
	BinExport2_Expression_IMMEDIATE_FLOAT BinExport2_Expression_Type = 3
	BinExport2_Expression_OPERATOR        BinExport2_Expression_Type = 4
	BinExport2_Expression_REGISTER        BinExport2_Expression_Type = 5
	BinExport2_Expression_SIZE_PREFIX     BinExport2_Expression_Type = 6
	BinExport2_Expression_DEREFERENCE     BinExport2_Expression_Type = 7
)

// Enum value maps for BinExport2_Expression_Type.
var (
	BinExport2_Expression_Type_name = map[int32]string{
		1: "SYMBOL",
		2: "IMMEDIATE_INT",
		3: "IMMEDIATE_FLOAT",
		4: "OPERATOR",
		5: "REGISTER",
		6: "SIZE_PREFIX",
		7: "DEREFERENCE",
	}
	BinExport2_Expression_Type_value = map[string]int32{
		"SYMBOL":          1,
		"IMMEDIATE_INT":   2,
		"IMMEDIATE_FLOAT": 3,
		"OPERATOR":        4,
		"REGISTER":        5,
		"SIZE_PREFIX":     6,
		"DEREFERENCE":     7,
	}
)

func (x BinExport2_Expression_Type) Enum() *BinExport2_Expression_Type {
	p := new(BinExport2_Expression_Type)
	*p = x
	return p
}

func (x BinExport2_Expression_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BinExport2_Expression_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_binexport2_proto_enumTypes[1].Descriptor()
}

func (BinExport2_Expression_Type) Type() protoreflect.EnumType {
	return &file_binexport2_proto_enumTypes[1]
}

func (x BinExport2_Expression_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *BinExport2_Expression_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = BinExport2_Expression_Type(num)
	return nil
}

// Deprecated: Use BinExport2_Expression_Type.Descriptor instead.
func (BinExport2_Expression_Type) EnumDescriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 2, 0}
}

type BinExport2_FlowGraph_Edge_Type int32

const (
	BinExport2_FlowGraph_Edge_CONDITION_TRUE  BinExport2_FlowGraph_Edge_Type = 1
	BinExport2_FlowGraph_Edge_CONDITION_FALSE BinExport2_FlowGraph_Edge_Type = 2
	BinExport2_FlowGraph_Edge_UNCONDITIONAL   BinExport2_FlowGraph_Edge_Type = 3
	BinExport2_FlowGraph_Edge_SWITCH          BinExport2_FlowGraph_Edge_Type = 4
)

// Enum value maps for BinExport2_FlowGraph_Edge_Type.
var (
	BinExport2_FlowGraph_Edge_Type_name = map[int32]string{
		1: "CONDITION_TRUE",
		2: "CONDITION_FALSE",
		3: "UNCONDITIONAL",
		4: "SWITCH",
	}
	BinExport2_FlowGraph_Edge_Type_value = map[string]int32{
		"CONDITION_TRUE":  1,
		"CONDITION_FALSE": 2,
		"UNCONDITIONAL":   3,
		"SWITCH":          4,
	}
)

func (x BinExport2_FlowGraph_Edge_Type) Enum() *BinExport2_FlowGraph_Edge_Type {
	p := new(BinExport2_FlowGraph_Edge_Type)
	*p = x
	return p
}

func (x BinExport2_FlowGraph_Edge_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BinExport2_FlowGraph_Edge_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_binexport2_proto_enumTypes[2].Descriptor()
}

func (BinExport2_FlowGraph_Edge_Type) Type() protoreflect.EnumType {
	return &file_binexport2_proto_enumTypes[2]
}

func (x BinExport2_FlowGraph_Edge_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *BinExport2_FlowGraph_Edge_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = BinExport2_FlowGraph_Edge_Type(num)
	return nil
}

// Deprecated: Use BinExport2_FlowGraph_Edge_Type.Descriptor instead.
func (BinExport2_FlowGraph_Edge_Type) EnumDescriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 7, 0, 0}
}

type BinExport2_Comment_Type int32

const (
	// A regular instruction comment. Typically displayed next to the
	// instruction disassembly.
	BinExport2_Comment_DEFAULT BinExport2_Comment_Type = 0
	// A comment line that is typically displayed before (above) the
	// instruction it refers to.
	BinExport2_Comment_ANTERIOR BinExport2_Comment_Type = 1
	// Like ANTERIOR, but a typically displayed after (below).
	BinExport2_Comment_POSTERIOR BinExport2_Comment_Type = 2
	// Similar to an ANTERIOR comment, but applies to the beginning of an
	// identified function. Programs displaying the proto may choose to render
	// these differently (e.g. above an inferred function signature).
	BinExport2_Comment_FUNCTION BinExport2_Comment_Type = 3
	// Named constants, bitfields and similar.
	BinExport2_Comment_ENUM BinExport2_Comment_Type = 4
	// Named locations, usually the target of a jump.
	BinExport2_Comment_LOCATION BinExport2_Comment_Type = 5
	// Data cross references.
	BinExport2_Comment_GLOBAL_REFERENCE BinExport2_Comment_Type = 6
	// Local/stack variables.
	BinExport2_Comment_LOCAL_REFERENCE BinExport2_Comment_Type = 7
)

// Enum value maps for BinExport2_Comment_Type.
var (
	BinExport2_Comment_Type_name = map[int32]string{
		0: "DEFAULT",
		1: "ANTERIOR",
		2: "POSTERIOR",
		3: "FUNCTION",
		4: "ENUM",
		5: "LOCATION",
		6: "GLOBAL_REFERENCE",
		7: "LOCAL_REFERENCE",
	}
	BinExport2_Comment_Type_value = map[string]int32{
		"DEFAULT":          0,
		"ANTERIOR":         1,
		"POSTERIOR":        2,
		"FUNCTION":         3,
		"ENUM":             4,
		"LOCATION":         5,
		"GLOBAL_REFERENCE": 6,
		"LOCAL_REFERENCE":  7,
	}
)

func (x BinExport2_Comment_Type) Enum() *BinExport2_Comment_Type {
	p := new(BinExport2_Comment_Type)
	*p = x
	return p
}

func (x BinExport2_Comment_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BinExport2_Comment_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_binexport2_proto_enumTypes[3].Descriptor()
}

func (BinExport2_Comment_Type) Type() protoreflect.EnumType {
	return &file_binexport2_proto_enumTypes[3]
}

func (x BinExport2_Comment_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *BinExport2_Comment_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = BinExport2_Comment_Type(num)
	return nil
}

// Deprecated: Use BinExport2_Comment_Type.Descriptor instead.
func (BinExport2_Comment_Type) EnumDescriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 10, 0}
}

type BinExport2 struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields

	MetaInformation *BinExport2_Meta          `protobuf:"bytes,1,opt,name=meta_information,json=metaInformation" json:"meta_information,omitempty"`
	Expression      []*BinExport2_Expression  `protobuf:"bytes,2,rep,name=expression" json:"expression,omitempty"`
	Operand         []*BinExport2_Operand     `protobuf:"bytes,3,rep,name=operand" json:"operand,omitempty"`
	Mnemonic        []*BinExport2_Mnemonic    `protobuf:"bytes,4,rep,name=mnemonic" json:"mnemonic,omitempty"`
	Instruction     []*BinExport2_Instruction `protobuf:"bytes,5,rep,name=instruction" json:"instruction,omitempty"`
	BasicBlock      []*BinExport2_BasicBlock  `protobuf:"bytes,6,rep,name=basic_block,json=basicBlock" json:"basic_block,omitempty"`
	FlowGraph       []*BinExport2_FlowGraph   `protobuf:"bytes,7,rep,name=flow_graph,json=flowGraph" json:"flow_graph,omitempty"`
	CallGraph       *BinExport2_CallGraph     `protobuf:"bytes,8,opt,name=call_graph,json=callGraph" json:"call_graph,omitempty"`
	StringTable     []string                  `protobuf:"bytes,9,rep,name=string_table,json=stringTable" json:"string_table,omitempty"`
	// No longer written. This is here so that BinDiff can work with older
	// BinExport files.
	//
	// Deprecated: Marked as deprecated in binexport2.proto.
	AddressComment []*BinExport2_Reference `protobuf:"bytes,10,rep,name=address_comment,json=addressComment" json:"address_comment,omitempty"`
	// Rich comment index used for BinDiff's comment porting.
	Comment                []*BinExport2_Comment       `protobuf:"bytes,17,rep,name=comment" json:"comment,omitempty"`
	StringReference        []*BinExport2_Reference     `protobuf:"bytes,11,rep,name=string_reference,json=stringReference" json:"string_reference,omitempty"`
	ExpressionSubstitution []*BinExport2_Reference     `protobuf:"bytes,12,rep,name=expression_substitution,json=expressionSubstitution" json:"expression_substitution,omitempty"`
	Section                []*BinExport2_Section       `protobuf:"bytes,13,rep,name=section" json:"section,omitempty"`
	Library                []*BinExport2_Library       `protobuf:"bytes,14,rep,name=library" json:"library,omitempty"`
	DataReference          []*BinExport2_DataReference `protobuf:"bytes,15,rep,name=data_reference,json=dataReference" json:"data_reference,omitempty"`
	Module                 []*BinExport2_Module        `protobuf:"bytes,16,rep,name=module" json:"module,omitempty"`
}

func (x *BinExport2) Reset() {
	*x = BinExport2{}
	mi := &file_binexport2_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2) ProtoMessage() {}

func (x *BinExport2) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2.ProtoReflect.Descriptor instead.
func (*BinExport2) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0}
}

func (x *BinExport2) GetMetaInformation() *BinExport2_Meta {
	if x != nil {
		return x.MetaInformation
	}
	return nil
}

func (x *BinExport2) GetExpression() []*BinExport2_Expression {
	if x != nil {
		return x.Expression
	}
	return nil
}

func (x *BinExport2) GetOperand() []*BinExport2_Operand {
	if x != nil {
		return x.Operand
	}
	return nil
}

func (x *BinExport2) GetMnemonic() []*BinExport2_Mnemonic {
	if x != nil {
		return x.Mnemonic
	}
	return nil
}

func (x *BinExport2) GetInstruction() []*BinExport2_Instruction {
	if x != nil {
		return x.Instruction
	}
	return nil
}

func (x *BinExport2) GetBasicBlock() []*BinExport2_BasicBlock {
	if x != nil {
		return x.BasicBlock
	}
	return nil
}

func (x *BinExport2) GetFlowGraph() []*BinExport2_FlowGraph {
	if x != nil {
		return x.FlowGraph
	}
	return nil
}

func (x *BinExport2) GetCallGraph() *BinExport2_CallGraph {
	if x != nil {
		return x.CallGraph
	}
	return nil
}

func (x *BinExport2) GetStringTable() []string {
	if x != nil {
		return x.StringTable
	}
	return nil
}

// Deprecated: Marked as deprecated in binexport2.proto.
func (x *BinExport2) GetAddressComment() []*BinExport2_Reference {
	if x != nil {
		return x.AddressComment
	}
	return nil
}

func (x *BinExport2) GetComment() []*BinExport2_Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

func (x *BinExport2) GetStringReference() []*BinExport2_Reference {
	if x != nil {
		return x.StringReference
	}
	return nil
}

func (x *BinExport2) GetExpressionSubstitution() []*BinExport2_Reference {
	if x != nil {
		return x.ExpressionSubstitution
	}
	return nil
}

func (x *BinExport2) GetSection() []*BinExport2_Section {
	if x != nil {
		return x.Section
	}
	return nil
}

func (x *BinExport2) GetLibrary() []*BinExport2_Library {
	if x != nil {
		return x.Library
	}
	return nil
}

func (x *BinExport2) GetDataReference() []*BinExport2_DataReference {
	if x != nil {
		return x.DataReference
	}
	return nil
}

func (x *BinExport2) GetModule() []*BinExport2_Module {
	if x != nil {
		return x.Module
	}
	return nil
}

type BinExport2_Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Input binary filename including file extension but excluding file path.
	// example: "insider_gcc.exe"
	ExecutableName *string `protobuf:"bytes,1,opt,name=executable_name,json=executableName" json:"executable_name,omitempty"`
	// Application defined executable id. Often the SHA256 hash of the input
	// binary.
	ExecutableId *string `protobuf:"bytes,2,opt,name=executable_id,json=executableId" json:"executable_id,omitempty"`
	// Input architecture name, e.g. x86-32.
	ArchitectureName *string `protobuf:"bytes,3,opt,name=architecture_name,json=architectureName" json:"architecture_name,omitempty"`
	// When did this file get created? Unix time. This may be used for some
	// primitive versioning in case the file format ever changes.
	Timestamp *int64 `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (x *BinExport2_Meta) Reset() {
	*x = BinExport2_Meta{}
	mi := &file_binexport2_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_Meta) ProtoMessage() {}

func (x *BinExport2_Meta) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_Meta.ProtoReflect.Descriptor instead.
func (*BinExport2_Meta) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BinExport2_Meta) GetExecutableName() string {
	if x != nil && x.ExecutableName != nil {
		return *x.ExecutableName
	}
	return ""
}

func (x *BinExport2_Meta) GetExecutableId() string {
	if x != nil && x.ExecutableId != nil {
		return *x.ExecutableId
	}
	return ""
}

func (x *BinExport2_Meta) GetArchitectureName() string {
	if x != nil && x.ArchitectureName != nil {
		return *x.ArchitectureName
	}
	return ""
}

func (x *BinExport2_Meta) GetTimestamp() int64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

type BinExport2_CallGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// vertices == functions in the call graph.
	Vertex []*BinExport2_CallGraph_Vertex `protobuf:"bytes,1,rep,name=vertex" json:"vertex,omitempty"`
	// edges == calls in the call graph.
	Edge []*BinExport2_CallGraph_Edge `protobuf:"bytes,2,rep,name=edge" json:"edge,omitempty"`
}

func (x *BinExport2_CallGraph) Reset() {
	*x = BinExport2_CallGraph{}
	mi := &file_binexport2_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_CallGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_CallGraph) ProtoMessage() {}

func (x *BinExport2_CallGraph) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_CallGraph.ProtoReflect.Descriptor instead.
func (*BinExport2_CallGraph) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 1}
}

func (x *BinExport2_CallGraph) GetVertex() []*BinExport2_CallGraph_Vertex {
	if x != nil {
		return x.Vertex
	}
	return nil
}

func (x *BinExport2_CallGraph) GetEdge() []*BinExport2_CallGraph_Edge {
	if x != nil {
		return x.Edge
	}
	return nil
}

// An operand consists of 1 or more expressions, linked together as a tree.
type BinExport2_Expression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IMMEDIATE_INT is by far the most common type and thus we can save some
	// space by omitting it as the default.
	Type *BinExport2_Expression_Type `protobuf:"varint,1,opt,name=type,enum=BinExport2_Expression_Type,def=2" json:"type,omitempty"`
	// Symbol for this expression. Interpretation depends on type. Examples
	// include: "eax", "[", "+"
	Symbol *string `protobuf:"bytes,2,opt,name=symbol" json:"symbol,omitempty"`
	// If the expression can be interpreted as an integer value (IMMEDIATE_INT)
	// the value is given here.
	Immediate *uint64 `protobuf:"varint,3,opt,name=immediate" json:"immediate,omitempty"`
	// The parent expression. Example expression tree for the second operand of:
	// mov eax, b4 [ebx + 12]
	// "b4" --- "[" --- "+" --- "ebx"
	//
	//	\  "12"
	ParentIndex *int32 `protobuf:"varint,4,opt,name=parent_index,json=parentIndex" json:"parent_index,omitempty"`
	// true if the expression has entry in relocation table
	IsRelocation *bool `protobuf:"varint,5,opt,name=is_relocation,json=isRelocation" json:"is_relocation,omitempty"`
}

// Default values for BinExport2_Expression fields.
const (
	Default_BinExport2_Expression_Type = BinExport2_Expression_IMMEDIATE_INT
)

func (x *BinExport2_Expression) Reset() {
	*x = BinExport2_Expression{}
	mi := &file_binexport2_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_Expression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_Expression) ProtoMessage() {}

func (x *BinExport2_Expression) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_Expression.ProtoReflect.Descriptor instead.
func (*BinExport2_Expression) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 2}
}

func (x *BinExport2_Expression) GetType() BinExport2_Expression_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Default_BinExport2_Expression_Type
}

func (x *BinExport2_Expression) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *BinExport2_Expression) GetImmediate() uint64 {
	if x != nil && x.Immediate != nil {
		return *x.Immediate
	}
	return 0
}

func (x *BinExport2_Expression) GetParentIndex() int32 {
	if x != nil && x.ParentIndex != nil {
		return *x.ParentIndex
	}
	return 0
}

func (x *BinExport2_Expression) GetIsRelocation() bool {
	if x != nil && x.IsRelocation != nil {
		return *x.IsRelocation
	}
	return false
}

// An instruction may have 0 or more operands.
type BinExport2_Operand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains all expressions constituting this operand. All expressions
	// should be linked into a single tree, i.e. there should only be one
	// expression in this list with parent_index == NULL and all others should
	// descend from that. Rendering order for expressions on the same tree level
	// (siblings) is implicitly given by the order they are referenced in this
	// repeated field.
	// Implicit: expression sequence
	ExpressionIndex []int32 `protobuf:"varint,1,rep,name=expression_index,json=expressionIndex" json:"expression_index,omitempty"`
}

func (x *BinExport2_Operand) Reset() {
	*x = BinExport2_Operand{}
	mi := &file_binexport2_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_Operand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_Operand) ProtoMessage() {}

func (x *BinExport2_Operand) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_Operand.ProtoReflect.Descriptor instead.
func (*BinExport2_Operand) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 3}
}

func (x *BinExport2_Operand) GetExpressionIndex() []int32 {
	if x != nil {
		return x.ExpressionIndex
	}
	return nil
}

// An instruction has exactly 1 mnemonic.
type BinExport2_Mnemonic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Literal representation of the mnemonic, e.g.: "mov".
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (x *BinExport2_Mnemonic) Reset() {
	*x = BinExport2_Mnemonic{}
	mi := &file_binexport2_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_Mnemonic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_Mnemonic) ProtoMessage() {}

func (x *BinExport2_Mnemonic) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_Mnemonic.ProtoReflect.Descriptor instead.
func (*BinExport2_Mnemonic) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 4}
}

func (x *BinExport2_Mnemonic) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type BinExport2_Instruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This will only be filled for instructions that do not just flow from the
	// immediately preceding instruction. Regular instructions will have to
	// calculate their own address by adding raw_bytes.size() to the previous
	// instruction's address.
	Address *uint64 `protobuf:"varint,1,opt,name=address" json:"address,omitempty"`
	// If this is a call instruction and call targets could be determined
	// they'll be given here. Note that we may or may not have a flow graph for
	// the target and thus cannot use an index into the flow graph table here.
	// We could potentially use call graph nodes, but linking instructions to
	// the call graph directly does not seem a good choice.
	CallTarget []uint64 `protobuf:"varint,2,rep,name=call_target,json=callTarget" json:"call_target,omitempty"`
	// Index into the mnemonic array of strings. Used for de-duping the data.
	// The default value is used for the most common mnemonic in the executable.
	MnemonicIndex *int32 `protobuf:"varint,3,opt,name=mnemonic_index,json=mnemonicIndex,def=0" json:"mnemonic_index,omitempty"`
	// Indices into the operand tree. On X86 this can be 0, 1 or 2 elements
	// long, 3 elements with VEX/EVEX.
	// Implicit: operand sequence
	OperandIndex []int32 `protobuf:"varint,4,rep,name=operand_index,json=operandIndex" json:"operand_index,omitempty"`
	// The unmodified input bytes corresponding to this instruction.
	RawBytes []byte `protobuf:"bytes,5,opt,name=raw_bytes,json=rawBytes" json:"raw_bytes,omitempty"`
	// Implicit: comment sequence
	CommentIndex []int32 `protobuf:"varint,6,rep,name=comment_index,json=commentIndex" json:"comment_index,omitempty"`
}

// Default values for BinExport2_Instruction fields.
const (
	Default_BinExport2_Instruction_MnemonicIndex = int32(0)
)

func (x *BinExport2_Instruction) Reset() {
	*x = BinExport2_Instruction{}
	mi := &file_binexport2_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_Instruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_Instruction) ProtoMessage() {}

func (x *BinExport2_Instruction) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_Instruction.ProtoReflect.Descriptor instead.
func (*BinExport2_Instruction) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 5}
}

func (x *BinExport2_Instruction) GetAddress() uint64 {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return 0
}

func (x *BinExport2_Instruction) GetCallTarget() []uint64 {
	if x != nil {
		return x.CallTarget
	}
	return nil
}

func (x *BinExport2_Instruction) GetMnemonicIndex() int32 {
	if x != nil && x.MnemonicIndex != nil {
		return *x.MnemonicIndex
	}
	return Default_BinExport2_Instruction_MnemonicIndex
}

func (x *BinExport2_Instruction) GetOperandIndex() []int32 {
	if x != nil {
		return x.OperandIndex
	}
	return nil
}

func (x *BinExport2_Instruction) GetRawBytes() []byte {
	if x != nil {
		return x.RawBytes
	}
	return nil
}

func (x *BinExport2_Instruction) GetCommentIndex() []int32 {
	if x != nil {
		return x.CommentIndex
	}
	return nil
}

type BinExport2_BasicBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Implicit: instruction sequence
	InstructionIndex []*BinExport2_BasicBlock_IndexRange `protobuf:"bytes,1,rep,name=instruction_index,json=instructionIndex" json:"instruction_index,omitempty"`
}

func (x *BinExport2_BasicBlock) Reset() {
	*x = BinExport2_BasicBlock{}
	mi := &file_binexport2_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_BasicBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_BasicBlock) ProtoMessage() {}

func (x *BinExport2_BasicBlock) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_BasicBlock.ProtoReflect.Descriptor instead.
func (*BinExport2_BasicBlock) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 6}
}

func (x *BinExport2_BasicBlock) GetInstructionIndex() []*BinExport2_BasicBlock_IndexRange {
	if x != nil {
		return x.InstructionIndex
	}
	return nil
}

type BinExport2_FlowGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Basic blocks are sorted by address.
	BasicBlockIndex []int32 `protobuf:"varint,1,rep,name=basic_block_index,json=basicBlockIndex" json:"basic_block_index,omitempty"`
	// The flow graph's entry point address is the first instruction of the
	// entry_basic_block.
	EntryBasicBlockIndex *int32                       `protobuf:"varint,3,opt,name=entry_basic_block_index,json=entryBasicBlockIndex" json:"entry_basic_block_index,omitempty"`
	Edge                 []*BinExport2_FlowGraph_Edge `protobuf:"bytes,2,rep,name=edge" json:"edge,omitempty"`
}

func (x *BinExport2_FlowGraph) Reset() {
	*x = BinExport2_FlowGraph{}
	mi := &file_binexport2_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_FlowGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_FlowGraph) ProtoMessage() {}

func (x *BinExport2_FlowGraph) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_FlowGraph.ProtoReflect.Descriptor instead.
func (*BinExport2_FlowGraph) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 7}
}

func (x *BinExport2_FlowGraph) GetBasicBlockIndex() []int32 {
	if x != nil {
		return x.BasicBlockIndex
	}
	return nil
}

func (x *BinExport2_FlowGraph) GetEntryBasicBlockIndex() int32 {
	if x != nil && x.EntryBasicBlockIndex != nil {
		return *x.EntryBasicBlockIndex
	}
	return 0
}

func (x *BinExport2_FlowGraph) GetEdge() []*BinExport2_FlowGraph_Edge {
	if x != nil {
		return x.Edge
	}
	return nil
}

// Generic reference class used for address comments (deprecated), string
// references and expression substitutions. It allows referencing from an
// instruction, operand, expression subtree tuple to a de-duped string in the
// string table.
type BinExport2_Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index into the global instruction table.
	InstructionIndex *int32 `protobuf:"varint,1,opt,name=instruction_index,json=instructionIndex" json:"instruction_index,omitempty"`
	// Index into the operand array local to an instruction.
	InstructionOperandIndex *int32 `protobuf:"varint,2,opt,name=instruction_operand_index,json=instructionOperandIndex,def=0" json:"instruction_operand_index,omitempty"`
	// Index into the expression array local to an operand.
	OperandExpressionIndex *int32 `protobuf:"varint,3,opt,name=operand_expression_index,json=operandExpressionIndex,def=0" json:"operand_expression_index,omitempty"`
	// Index into the global string table.
	StringTableIndex *int32 `protobuf:"varint,4,opt,name=string_table_index,json=stringTableIndex" json:"string_table_index,omitempty"`
}

// Default values for BinExport2_Reference fields.
const (
	Default_BinExport2_Reference_InstructionOperandIndex = int32(0)
	Default_BinExport2_Reference_OperandExpressionIndex  = int32(0)
)

func (x *BinExport2_Reference) Reset() {
	*x = BinExport2_Reference{}
	mi := &file_binexport2_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_Reference) ProtoMessage() {}

func (x *BinExport2_Reference) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_Reference.ProtoReflect.Descriptor instead.
func (*BinExport2_Reference) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 8}
}

func (x *BinExport2_Reference) GetInstructionIndex() int32 {
	if x != nil && x.InstructionIndex != nil {
		return *x.InstructionIndex
	}
	return 0
}

func (x *BinExport2_Reference) GetInstructionOperandIndex() int32 {
	if x != nil && x.InstructionOperandIndex != nil {
		return *x.InstructionOperandIndex
	}
	return Default_BinExport2_Reference_InstructionOperandIndex
}

func (x *BinExport2_Reference) GetOperandExpressionIndex() int32 {
	if x != nil && x.OperandExpressionIndex != nil {
		return *x.OperandExpressionIndex
	}
	return Default_BinExport2_Reference_OperandExpressionIndex
}

func (x *BinExport2_Reference) GetStringTableIndex() int32 {
	if x != nil && x.StringTableIndex != nil {
		return *x.StringTableIndex
	}
	return 0
}

type BinExport2_DataReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index into the global instruction table.
	InstructionIndex *int32 `protobuf:"varint,1,opt,name=instruction_index,json=instructionIndex" json:"instruction_index,omitempty"`
	// Address being referred.
	Address *uint64 `protobuf:"varint,2,opt,name=address" json:"address,omitempty"`
}

func (x *BinExport2_DataReference) Reset() {
	*x = BinExport2_DataReference{}
	mi := &file_binexport2_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_DataReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_DataReference) ProtoMessage() {}

func (x *BinExport2_DataReference) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_DataReference.ProtoReflect.Descriptor instead.
func (*BinExport2_DataReference) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 9}
}

func (x *BinExport2_DataReference) GetInstructionIndex() int32 {
	if x != nil && x.InstructionIndex != nil {
		return *x.InstructionIndex
	}
	return 0
}

func (x *BinExport2_DataReference) GetAddress() uint64 {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return 0
}

type BinExport2_Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index into the global instruction table. This is here to enable
	// comment processing without having to iterate over all instructions.
	// There is an N:M mapping of instructions to comments.
	InstructionIndex *int32 `protobuf:"varint,1,opt,name=instruction_index,json=instructionIndex" json:"instruction_index,omitempty"`
	// Index into the operand array local to an instruction.
	InstructionOperandIndex *int32 `protobuf:"varint,2,opt,name=instruction_operand_index,json=instructionOperandIndex,def=0" json:"instruction_operand_index,omitempty"`
	// Index into the expression array local to an operand, like in Reference.
	// This is not currently used, but allows to implement expression
	// substitutions.
	OperandExpressionIndex *int32 `protobuf:"varint,3,opt,name=operand_expression_index,json=operandExpressionIndex,def=0" json:"operand_expression_index,omitempty"`
	// Index into the global string table.
	StringTableIndex *int32 `protobuf:"varint,4,opt,name=string_table_index,json=stringTableIndex" json:"string_table_index,omitempty"`
	// Comment is propagated to all locations that reference the original
	// location.
	Repeatable *bool                    `protobuf:"varint,5,opt,name=repeatable" json:"repeatable,omitempty"`
	Type       *BinExport2_Comment_Type `protobuf:"varint,6,opt,name=type,enum=BinExport2_Comment_Type,def=0" json:"type,omitempty"`
}

// Default values for BinExport2_Comment fields.
const (
	Default_BinExport2_Comment_InstructionOperandIndex = int32(0)
	Default_BinExport2_Comment_OperandExpressionIndex  = int32(0)
	Default_BinExport2_Comment_Type                    = BinExport2_Comment_DEFAULT
)

func (x *BinExport2_Comment) Reset() {
	*x = BinExport2_Comment{}
	mi := &file_binexport2_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_Comment) ProtoMessage() {}

func (x *BinExport2_Comment) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_Comment.ProtoReflect.Descriptor instead.
func (*BinExport2_Comment) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 10}
}

func (x *BinExport2_Comment) GetInstructionIndex() int32 {
	if x != nil && x.InstructionIndex != nil {
		return *x.InstructionIndex
	}
	return 0
}

func (x *BinExport2_Comment) GetInstructionOperandIndex() int32 {
	if x != nil && x.InstructionOperandIndex != nil {
		return *x.InstructionOperandIndex
	}
	return Default_BinExport2_Comment_InstructionOperandIndex
}

func (x *BinExport2_Comment) GetOperandExpressionIndex() int32 {
	if x != nil && x.OperandExpressionIndex != nil {
		return *x.OperandExpressionIndex
	}
	return Default_BinExport2_Comment_OperandExpressionIndex
}

func (x *BinExport2_Comment) GetStringTableIndex() int32 {
	if x != nil && x.StringTableIndex != nil {
		return *x.StringTableIndex
	}
	return 0
}

func (x *BinExport2_Comment) GetRepeatable() bool {
	if x != nil && x.Repeatable != nil {
		return *x.Repeatable
	}
	return false
}

func (x *BinExport2_Comment) GetType() BinExport2_Comment_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Default_BinExport2_Comment_Type
}

type BinExport2_Section struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Section start address.
	Address *uint64 `protobuf:"varint,1,opt,name=address" json:"address,omitempty"`
	// Section size.
	Size *uint64 `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	// Read flag of the section, True when section is readable.
	FlagR *bool `protobuf:"varint,3,opt,name=flag_r,json=flagR" json:"flag_r,omitempty"`
	// Write flag of the section, True when section is writable.
	FlagW *bool `protobuf:"varint,4,opt,name=flag_w,json=flagW" json:"flag_w,omitempty"`
	// Execute flag of the section, True when section is executable.
	FlagX *bool `protobuf:"varint,5,opt,name=flag_x,json=flagX" json:"flag_x,omitempty"`
}

func (x *BinExport2_Section) Reset() {
	*x = BinExport2_Section{}
	mi := &file_binexport2_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_Section) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_Section) ProtoMessage() {}

func (x *BinExport2_Section) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_Section.ProtoReflect.Descriptor instead.
func (*BinExport2_Section) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 11}
}

func (x *BinExport2_Section) GetAddress() uint64 {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return 0
}

func (x *BinExport2_Section) GetSize() uint64 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *BinExport2_Section) GetFlagR() bool {
	if x != nil && x.FlagR != nil {
		return *x.FlagR
	}
	return false
}

func (x *BinExport2_Section) GetFlagW() bool {
	if x != nil && x.FlagW != nil {
		return *x.FlagW
	}
	return false
}

func (x *BinExport2_Section) GetFlagX() bool {
	if x != nil && x.FlagX != nil {
		return *x.FlagX
	}
	return false
}

type BinExport2_Library struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If this library is statically linked.
	IsStatic *bool `protobuf:"varint,1,opt,name=is_static,json=isStatic" json:"is_static,omitempty"`
	// Address where this library was loaded, 0 if unknown.
	LoadAddress *uint64 `protobuf:"varint,2,opt,name=load_address,json=loadAddress,def=0" json:"load_address,omitempty"`
	// Name of the library (format is platform-dependent).
	Name *string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

// Default values for BinExport2_Library fields.
const (
	Default_BinExport2_Library_LoadAddress = uint64(0)
)

func (x *BinExport2_Library) Reset() {
	*x = BinExport2_Library{}
	mi := &file_binexport2_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_Library) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_Library) ProtoMessage() {}

func (x *BinExport2_Library) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_Library.ProtoReflect.Descriptor instead.
func (*BinExport2_Library) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 12}
}

func (x *BinExport2_Library) GetIsStatic() bool {
	if x != nil && x.IsStatic != nil {
		return *x.IsStatic
	}
	return false
}

func (x *BinExport2_Library) GetLoadAddress() uint64 {
	if x != nil && x.LoadAddress != nil {
		return *x.LoadAddress
	}
	return Default_BinExport2_Library_LoadAddress
}

func (x *BinExport2_Library) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type BinExport2_Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name, such as Java class name. Platform-dependent.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (x *BinExport2_Module) Reset() {
	*x = BinExport2_Module{}
	mi := &file_binexport2_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_Module) ProtoMessage() {}

func (x *BinExport2_Module) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_Module.ProtoReflect.Descriptor instead.
func (*BinExport2_Module) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 13}
}

func (x *BinExport2_Module) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type BinExport2_CallGraph_Vertex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The function's entry point address.
	Address *uint64                           `protobuf:"varint,1,opt,name=address" json:"address,omitempty"`
	Type    *BinExport2_CallGraph_Vertex_Type `protobuf:"varint,2,opt,name=type,enum=BinExport2_CallGraph_Vertex_Type,def=0" json:"type,omitempty"`
	// If the function has a user defined, real name it will be given here.
	// main() is a proper name, sub_BAADF00D is not (auto generated dummy
	// name).
	MangledName *string `protobuf:"bytes,3,opt,name=mangled_name,json=mangledName" json:"mangled_name,omitempty"`
	// Demangled name if the function is a mangled C++ function and we could
	// demangle it.
	DemangledName *string `protobuf:"bytes,4,opt,name=demangled_name,json=demangledName" json:"demangled_name,omitempty"`
	// If this is a library function, what is its index in library arrays.
	LibraryIndex *int32 `protobuf:"varint,5,opt,name=library_index,json=libraryIndex" json:"library_index,omitempty"`
	// If module name, such as class name for DEX files, is present - index in
	// module table.
	ModuleIndex *int32 `protobuf:"varint,6,opt,name=module_index,json=moduleIndex" json:"module_index,omitempty"`
}

// Default values for BinExport2_CallGraph_Vertex fields.
const (
	Default_BinExport2_CallGraph_Vertex_Type = BinExport2_CallGraph_Vertex_NORMAL
)

func (x *BinExport2_CallGraph_Vertex) Reset() {
	*x = BinExport2_CallGraph_Vertex{}
	mi := &file_binexport2_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_CallGraph_Vertex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_CallGraph_Vertex) ProtoMessage() {}

func (x *BinExport2_CallGraph_Vertex) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_CallGraph_Vertex.ProtoReflect.Descriptor instead.
func (*BinExport2_CallGraph_Vertex) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *BinExport2_CallGraph_Vertex) GetAddress() uint64 {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return 0
}

func (x *BinExport2_CallGraph_Vertex) GetType() BinExport2_CallGraph_Vertex_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Default_BinExport2_CallGraph_Vertex_Type
}

func (x *BinExport2_CallGraph_Vertex) GetMangledName() string {
	if x != nil && x.MangledName != nil {
		return *x.MangledName
	}
	return ""
}

func (x *BinExport2_CallGraph_Vertex) GetDemangledName() string {
	if x != nil && x.DemangledName != nil {
		return *x.DemangledName
	}
	return ""
}

func (x *BinExport2_CallGraph_Vertex) GetLibraryIndex() int32 {
	if x != nil && x.LibraryIndex != nil {
		return *x.LibraryIndex
	}
	return 0
}

func (x *BinExport2_CallGraph_Vertex) GetModuleIndex() int32 {
	if x != nil && x.ModuleIndex != nil {
		return *x.ModuleIndex
	}
	return 0
}

type BinExport2_CallGraph_Edge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// source and target index into the vertex repeated field.
	SourceVertexIndex *int32 `protobuf:"varint,1,opt,name=source_vertex_index,json=sourceVertexIndex" json:"source_vertex_index,omitempty"`
	TargetVertexIndex *int32 `protobuf:"varint,2,opt,name=target_vertex_index,json=targetVertexIndex" json:"target_vertex_index,omitempty"`
}

func (x *BinExport2_CallGraph_Edge) Reset() {
	*x = BinExport2_CallGraph_Edge{}
	mi := &file_binexport2_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_CallGraph_Edge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_CallGraph_Edge) ProtoMessage() {}

func (x *BinExport2_CallGraph_Edge) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_CallGraph_Edge.ProtoReflect.Descriptor instead.
func (*BinExport2_CallGraph_Edge) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 1, 1}
}

func (x *BinExport2_CallGraph_Edge) GetSourceVertexIndex() int32 {
	if x != nil && x.SourceVertexIndex != nil {
		return *x.SourceVertexIndex
	}
	return 0
}

func (x *BinExport2_CallGraph_Edge) GetTargetVertexIndex() int32 {
	if x != nil && x.TargetVertexIndex != nil {
		return *x.TargetVertexIndex
	}
	return 0
}

// This is a space optimization. The instructions for an individual basic
// block will usually be in a continuous index range. Thus it is more
// efficient to store the range instead of individual indices. However, this
// does not hold true for all basic blocks, so we need to be able to store
// multiple index ranges per block.
type BinExport2_BasicBlock_IndexRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// These work like begin and end iterators, i.e. the sequence is
	// [begin_index, end_index). If the sequence only contains a single
	// element end_index will be omitted.
	BeginIndex *int32 `protobuf:"varint,1,opt,name=begin_index,json=beginIndex" json:"begin_index,omitempty"`
	EndIndex   *int32 `protobuf:"varint,2,opt,name=end_index,json=endIndex" json:"end_index,omitempty"`
}

func (x *BinExport2_BasicBlock_IndexRange) Reset() {
	*x = BinExport2_BasicBlock_IndexRange{}
	mi := &file_binexport2_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_BasicBlock_IndexRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_BasicBlock_IndexRange) ProtoMessage() {}

func (x *BinExport2_BasicBlock_IndexRange) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_BasicBlock_IndexRange.ProtoReflect.Descriptor instead.
func (*BinExport2_BasicBlock_IndexRange) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 6, 0}
}

func (x *BinExport2_BasicBlock_IndexRange) GetBeginIndex() int32 {
	if x != nil && x.BeginIndex != nil {
		return *x.BeginIndex
	}
	return 0
}

func (x *BinExport2_BasicBlock_IndexRange) GetEndIndex() int32 {
	if x != nil && x.EndIndex != nil {
		return *x.EndIndex
	}
	return 0
}

type BinExport2_FlowGraph_Edge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source instruction will always be the last instruction of the source
	// basic block, target instruction the first instruction of the target
	// basic block.
	SourceBasicBlockIndex *int32                          `protobuf:"varint,1,opt,name=source_basic_block_index,json=sourceBasicBlockIndex" json:"source_basic_block_index,omitempty"`
	TargetBasicBlockIndex *int32                          `protobuf:"varint,2,opt,name=target_basic_block_index,json=targetBasicBlockIndex" json:"target_basic_block_index,omitempty"`
	Type                  *BinExport2_FlowGraph_Edge_Type `protobuf:"varint,3,opt,name=type,enum=BinExport2_FlowGraph_Edge_Type,def=3" json:"type,omitempty"`
	// Indicates whether this is a loop edge as determined by Lengauer-Tarjan.
	IsBackEdge *bool `protobuf:"varint,4,opt,name=is_back_edge,json=isBackEdge,def=0" json:"is_back_edge,omitempty"`
}

// Default values for BinExport2_FlowGraph_Edge fields.
const (
	Default_BinExport2_FlowGraph_Edge_Type       = BinExport2_FlowGraph_Edge_UNCONDITIONAL
	Default_BinExport2_FlowGraph_Edge_IsBackEdge = bool(false)
)

func (x *BinExport2_FlowGraph_Edge) Reset() {
	*x = BinExport2_FlowGraph_Edge{}
	mi := &file_binexport2_proto_msgTypes[18]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinExport2_FlowGraph_Edge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinExport2_FlowGraph_Edge) ProtoMessage() {}

func (x *BinExport2_FlowGraph_Edge) ProtoReflect() protoreflect.Message {
	mi := &file_binexport2_proto_msgTypes[18]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinExport2_FlowGraph_Edge.ProtoReflect.Descriptor instead.
func (*BinExport2_FlowGraph_Edge) Descriptor() ([]byte, []int) {
	return file_binexport2_proto_rawDescGZIP(), []int{0, 7, 0}
}

func (x *BinExport2_FlowGraph_Edge) GetSourceBasicBlockIndex() int32 {
	if x != nil && x.SourceBasicBlockIndex != nil {
		return *x.SourceBasicBlockIndex
	}
	return 0
}

func (x *BinExport2_FlowGraph_Edge) GetTargetBasicBlockIndex() int32 {
	if x != nil && x.TargetBasicBlockIndex != nil {
		return *x.TargetBasicBlockIndex
	}
	return 0
}

func (x *BinExport2_FlowGraph_Edge) GetType() BinExport2_FlowGraph_Edge_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Default_BinExport2_FlowGraph_Edge_Type
}

func (x *BinExport2_FlowGraph_Edge) GetIsBackEdge() bool {
	if x != nil && x.IsBackEdge != nil {
		return *x.IsBackEdge
	}
	return Default_BinExport2_FlowGraph_Edge_IsBackEdge
}

var File_binexport2_proto protoreflect.FileDescriptor

var file_binexport2_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x69, 0x6e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfa, 0x1e, 0x0a, 0x0a, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x32, 0x12, 0x3b, 0x0a, 0x10, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x42, 0x69,
	0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0f, 0x6d,
	0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36,
	0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e,
	0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x32, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x07, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69,
	0x63, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x32, 0x2e, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x52, 0x08, 0x6d,
	0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12, 0x39, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x42,
	0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x32, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x0a, 0x62, 0x61, 0x73, 0x69, 0x63, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x34, 0x0a, 0x0a, 0x66,
	0x6c, 0x6f, 0x77, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x46, 0x6c, 0x6f,
	0x77, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x09, 0x66, 0x6c, 0x6f, 0x77, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x32, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x09, 0x63, 0x61,
	0x6c, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32,
	0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0e,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2d,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a,
	0x10, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x32, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x4e, 0x0a, 0x17, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75,
	0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x16, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2d, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d,
	0x0a, 0x07, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x4c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x52, 0x07, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x40, 0x0a,
	0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x32, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x2a, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x1a, 0xa5, 0x01, 0x0a, 0x04,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4a, 0x04, 0x08,
	0x05, 0x10, 0x06, 0x1a, 0x96, 0x04, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x6c, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x12, 0x34, 0x0a, 0x06, 0x76, 0x65, 0x72, 0x74, 0x65, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x43,
	0x61, 0x6c, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x52,
	0x06, 0x76, 0x65, 0x72, 0x74, 0x65, 0x78, 0x12, 0x2e, 0x0a, 0x04, 0x65, 0x64, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x32, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x45, 0x64, 0x67,
	0x65, 0x52, 0x04, 0x65, 0x64, 0x67, 0x65, 0x1a, 0xba, 0x02, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x74,
	0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x42, 0x69, 0x6e,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x06, 0x4e,
	0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d,
	0x61, 0x6e, 0x67, 0x6c, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6d, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x67, 0x6c, 0x65,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x45, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x49, 0x42, 0x52, 0x41, 0x52, 0x59, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x54, 0x48, 0x55, 0x4e, 0x4b, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x04, 0x1a, 0x66, 0x0a, 0x04, 0x45, 0x64, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x13,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x74, 0x65, 0x78, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2e, 0x0a, 0x13,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x74, 0x65, 0x78, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0xc4, 0x02, 0x0a,
	0x0a, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x42, 0x69, 0x6e, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x0d, 0x49, 0x4d, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54,
	0x45, 0x5f, 0x49, 0x4e, 0x54, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x52,
	0x65, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x78, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59, 0x4d, 0x42, 0x4f, 0x4c, 0x10, 0x01, 0x12, 0x11, 0x0a,
	0x0d, 0x49, 0x4d, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x02,
	0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4d, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x4c,
	0x4f, 0x41, 0x54, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f,
	0x52, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10,
	0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58,
	0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x45, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43,
	0x45, 0x10, 0x07, 0x1a, 0x34, 0x0a, 0x07, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x29,
	0x0a, 0x10, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x1e, 0x0a, 0x08, 0x4d, 0x6e, 0x65,
	0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0xd9, 0x01, 0x0a, 0x0b, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x6c, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x0e, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52,
	0x0d, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x23,
	0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x77, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x61, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x06, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0xa8, 0x01, 0x0a, 0x0a, 0x42, 0x61, 0x73, 0x69, 0x63, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x4e, 0x0a, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x1a, 0x4a, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x1a, 0xd6, 0x03, 0x0a, 0x09, 0x46, 0x6c, 0x6f, 0x77, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x2a,
	0x0a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0f, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x35, 0x0a, 0x17, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x61, 0x73, 0x69, 0x63, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x2e, 0x0a, 0x04, 0x65, 0x64, 0x67, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x46, 0x6c, 0x6f,
	0x77, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x52, 0x04, 0x65, 0x64, 0x67,
	0x65, 0x1a, 0xb5, 0x02, 0x0a, 0x04, 0x45, 0x64, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x42, 0x61, 0x73, 0x69, 0x63, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x37, 0x0a, 0x18, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x42, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x42, 0x69, 0x6e,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x0d, 0x55, 0x4e, 0x43,
	0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x27, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x65, 0x64, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x0a, 0x69,
	0x73, 0x42, 0x61, 0x63, 0x6b, 0x45, 0x64, 0x67, 0x65, 0x22, 0x4e, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x52, 0x55, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x4c, 0x53, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e,
	0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x0a, 0x0a,
	0x06, 0x53, 0x57, 0x49, 0x54, 0x43, 0x48, 0x10, 0x04, 0x1a, 0xe2, 0x01, 0x0a, 0x09, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x3d, 0x0a, 0x19, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x17, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x3b, 0x0a, 0x18, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x16, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e,
	0x64, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x2c, 0x0a, 0x12, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x56,
	0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0xbb, 0x03, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x3d, 0x0a, 0x19, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x17, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x3b,
	0x0a, 0x18, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x3a, 0x01, 0x30, 0x52, 0x16, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2c, 0x0a, 0x12, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x3a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x81, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46,
	0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x4e, 0x54, 0x45, 0x52, 0x49,
	0x4f, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x4f, 0x53, 0x54, 0x45, 0x52, 0x49, 0x4f,
	0x52, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x03, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x4e, 0x55, 0x4d, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x4c,
	0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x4c, 0x4f,
	0x42, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x06, 0x12,
	0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e,
	0x43, 0x45, 0x10, 0x07, 0x1a, 0x7c, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66,
	0x6c, 0x61, 0x67, 0x52, 0x12, 0x15, 0x0a, 0x06, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x77, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x57, 0x12, 0x15, 0x0a, 0x06, 0x66,
	0x6c, 0x61, 0x67, 0x5f, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6c, 0x61,
	0x67, 0x58, 0x1a, 0x60, 0x0a, 0x07, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x12, 0x24, 0x0a, 0x0c, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x3a, 0x01, 0x30, 0x52, 0x0b, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x1c, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x2a, 0x0b, 0x08, 0x80, 0xc2, 0xd7, 0x2f, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x42,
	0x36, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x7a, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x73, 0x42,
	0x09, 0x42, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5a, 0x0b, 0x2e, 0x2f, 0x62, 0x69,
	0x6e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74,
}

var (
	file_binexport2_proto_rawDescOnce sync.Once
	file_binexport2_proto_rawDescData = file_binexport2_proto_rawDesc
)

func file_binexport2_proto_rawDescGZIP() []byte {
	file_binexport2_proto_rawDescOnce.Do(func() {
		file_binexport2_proto_rawDescData = protoimpl.X.CompressGZIP(file_binexport2_proto_rawDescData)
	})
	return file_binexport2_proto_rawDescData
}

var file_binexport2_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_binexport2_proto_msgTypes = make([]protoimpl.MessageInfo, 19)
var file_binexport2_proto_goTypes = []any{
	(BinExport2_CallGraph_Vertex_Type)(0),    // 0: BinExport2.CallGraph.Vertex.Type
	(BinExport2_Expression_Type)(0),          // 1: BinExport2.Expression.Type
	(BinExport2_FlowGraph_Edge_Type)(0),      // 2: BinExport2.FlowGraph.Edge.Type
	(BinExport2_Comment_Type)(0),             // 3: BinExport2.Comment.Type
	(*BinExport2)(nil),                       // 4: BinExport2
	(*BinExport2_Meta)(nil),                  // 5: BinExport2.Meta
	(*BinExport2_CallGraph)(nil),             // 6: BinExport2.CallGraph
	(*BinExport2_Expression)(nil),            // 7: BinExport2.Expression
	(*BinExport2_Operand)(nil),               // 8: BinExport2.Operand
	(*BinExport2_Mnemonic)(nil),              // 9: BinExport2.Mnemonic
	(*BinExport2_Instruction)(nil),           // 10: BinExport2.Instruction
	(*BinExport2_BasicBlock)(nil),            // 11: BinExport2.BasicBlock
	(*BinExport2_FlowGraph)(nil),             // 12: BinExport2.FlowGraph
	(*BinExport2_Reference)(nil),             // 13: BinExport2.Reference
	(*BinExport2_DataReference)(nil),         // 14: BinExport2.DataReference
	(*BinExport2_Comment)(nil),               // 15: BinExport2.Comment
	(*BinExport2_Section)(nil),               // 16: BinExport2.Section
	(*BinExport2_Library)(nil),               // 17: BinExport2.Library
	(*BinExport2_Module)(nil),                // 18: BinExport2.Module
	(*BinExport2_CallGraph_Vertex)(nil),      // 19: BinExport2.CallGraph.Vertex
	(*BinExport2_CallGraph_Edge)(nil),        // 20: BinExport2.CallGraph.Edge
	(*BinExport2_BasicBlock_IndexRange)(nil), // 21: BinExport2.BasicBlock.IndexRange
	(*BinExport2_FlowGraph_Edge)(nil),        // 22: BinExport2.FlowGraph.Edge
}
var file_binexport2_proto_depIdxs = []int32{
	5,  // 0: BinExport2.meta_information:type_name -> BinExport2.Meta
	7,  // 1: BinExport2.expression:type_name -> BinExport2.Expression
	8,  // 2: BinExport2.operand:type_name -> BinExport2.Operand
	9,  // 3: BinExport2.mnemonic:type_name -> BinExport2.Mnemonic
	10, // 4: BinExport2.instruction:type_name -> BinExport2.Instruction
	11, // 5: BinExport2.basic_block:type_name -> BinExport2.BasicBlock
	12, // 6: BinExport2.flow_graph:type_name -> BinExport2.FlowGraph
	6,  // 7: BinExport2.call_graph:type_name -> BinExport2.CallGraph
	13, // 8: BinExport2.address_comment:type_name -> BinExport2.Reference
	15, // 9: BinExport2.comment:type_name -> BinExport2.Comment
	13, // 10: BinExport2.string_reference:type_name -> BinExport2.Reference
	13, // 11: BinExport2.expression_substitution:type_name -> BinExport2.Reference
	16, // 12: BinExport2.section:type_name -> BinExport2.Section
	17, // 13: BinExport2.library:type_name -> BinExport2.Library
	14, // 14: BinExport2.data_reference:type_name -> BinExport2.DataReference
	18, // 15: BinExport2.module:type_name -> BinExport2.Module
	19, // 16: BinExport2.CallGraph.vertex:type_name -> BinExport2.CallGraph.Vertex
	20, // 17: BinExport2.CallGraph.edge:type_name -> BinExport2.CallGraph.Edge
	1,  // 18: BinExport2.Expression.type:type_name -> BinExport2.Expression.Type
	21, // 19: BinExport2.BasicBlock.instruction_index:type_name -> BinExport2.BasicBlock.IndexRange
	22, // 20: BinExport2.FlowGraph.edge:type_name -> BinExport2.FlowGraph.Edge
	3,  // 21: BinExport2.Comment.type:type_name -> BinExport2.Comment.Type
	0,  // 22: BinExport2.CallGraph.Vertex.type:type_name -> BinExport2.CallGraph.Vertex.Type
	2,  // 23: BinExport2.FlowGraph.Edge.type:type_name -> BinExport2.FlowGraph.Edge.Type
	24, // [24:24] is the sub-list for method output_type
	24, // [24:24] is the sub-list for method input_type
	24, // [24:24] is the sub-list for extension type_name
	24, // [24:24] is the sub-list for extension extendee
	0,  // [0:24] is the sub-list for field type_name
}

func init() { file_binexport2_proto_init() }
func file_binexport2_proto_init() {
	if File_binexport2_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_binexport2_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   19,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_binexport2_proto_goTypes,
		DependencyIndexes: file_binexport2_proto_depIdxs,
		EnumInfos:         file_binexport2_proto_enumTypes,
		MessageInfos:      file_binexport2_proto_msgTypes,
	}.Build()
	File_binexport2_proto = out.File
	file_binexport2_proto_rawDesc = nil
	file_binexport2_proto_goTypes = nil
	file_binexport2_proto_depIdxs = nil
}
