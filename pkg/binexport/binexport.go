package binexport

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

type BinExport struct {
	path string
}

func NewBinExport(path string) *BinExport {
	return &BinExport{
		path: path,
	}
}

// PACIBSP
// STP             X20, X19, [SP,#-0x10+var_10]!
// STP             X29, X30, [SP,#0x10+var_s0]
// ADD             X29, SP, #0x10
// CBZ             s, loc_FFFFFE00072B4038

func (b *BinExport) Run() error {
	data, err := os.ReadFile(b.path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var bexp BinExport2
	if err := proto.Unmarshal(data, &bexp); err != nil {
		return fmt.Errorf("failed to unmarshal BinExport2: %w", err)
	}

	fmt.Println(bexp.GetMetaInformation())

	for _, flow := range bexp.GetFlowGraph() {
		bb := bexp.BasicBlock[flow.GetEntryBasicBlockIndex()]
		for _, blockInst := range bb.GetInstructionIndex() {
			var prevAddr uint64
			for _, inst := range bexp.Instruction[blockInst.GetBeginIndex():blockInst.GetEndIndex()] {
				if inst.Address != nil {
					mnemonic := bexp.Mnemonic[inst.GetMnemonicIndex()]
					fmt.Printf("%#x: %s\n", inst.GetAddress(), mnemonic.GetName())
					prevAddr = inst.GetAddress()
				} else {
					mnemonic := bexp.Mnemonic[inst.GetMnemonicIndex()]
					var out string
					for _, oidx := range inst.GetOperandIndex() {
						for _, eidx := range bexp.Operand[oidx].GetExpressionIndex() {
							exp := bexp.Expression[eidx]
							fmt.Printf("expression: %d) %s\n", eidx, exp)
							switch exp.GetType() {
							case BinExport2_Expression_SYMBOL:
								out += exp.GetSymbol()
							case BinExport2_Expression_IMMEDIATE_INT:
								out += fmt.Sprintf("%d", int64(exp.GetImmediate()))
							case BinExport2_Expression_IMMEDIATE_FLOAT:
								out += fmt.Sprintf("%f", exp.GetImmediate())
							case BinExport2_Expression_OPERATOR:
								out += exp.GetSymbol()
							case BinExport2_Expression_REGISTER:
								out += exp.GetSymbol() + ", "
							case BinExport2_Expression_SIZE_PREFIX:
								// out += exp.GetSymbol()
							case BinExport2_Expression_DEREFERENCE:
								out += exp.GetSymbol()
							default:
								out += "unknown"
							}
						}
					}
					fmt.Printf("%#x: %s\t%s\n", prevAddr, mnemonic.GetName(), out)
				}
			}
		}
	}

	return nil
}
