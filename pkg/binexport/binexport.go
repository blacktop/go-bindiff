package binexport

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/proto"
)

const bindiffExePath = "/Applications/BinDiff/BinDiff.app/Contents/MacOS/bin/bindiff"

type BinExport struct {
	bexp2 BinExport2
}

func runBinExportCmd(path string) (string, error) {
	// create new path of .BinExport file
	//   (NOTE: the assumption is that the BinExport file is in the same directory as the input file)
	binexportPath := filepath.Join(
		filepath.Dir(path),
		strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))+".BinExport")
	// check if .BinExport file already exists
	if _, err := os.Stat(binexportPath); err == nil {
		return binexportPath, nil
	}
	cmd := exec.Command(bindiffExePath, "--export", filepath.Dir(path))
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run bindiff --export: %v: cmd output - '%s'", err, out)
	}
	if _, err := os.Stat(binexportPath); os.IsNotExist(err) {
		return "", fmt.Errorf("failed to create .BinExport file: %s", binexportPath)
	}
	return binexportPath, nil
}

func NewBinExport(path string) (*BinExport, error) {
	var b BinExport
	switch filepath.Ext(path) {
	case ".i64":
		var err error
		path, err = runBinExportCmd(path)
		if err != nil {
			return nil, fmt.Errorf("failed to run bindiff --export: %w", err)
		}
		fallthrough
	case ".BinExport":
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read %s: %w", path, err)
		}
		if err := proto.Unmarshal(data, &b.bexp2); err != nil {
			return nil, fmt.Errorf("failed to unmarshal BinExport2: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported file extension: %s (expected .BinExport or .i64)", filepath.Ext(path))
	}
	return &b, nil
}

func (b *BinExport) ExecutableName() string {
	return b.bexp2.GetMetaInformation().GetExecutableName()
}

// PACIBSP
// STP             X20, X19, [SP,#-0x10+var_10]!
// STP             X29, X30, [SP,#0x10+var_s0]
// ADD             X29, SP, #0x10
// CBZ             s, loc_FFFFFE00072B4038

func (b *BinExport) Dump() {
	fmt.Println(b.bexp2.GetMetaInformation())

	for _, flow := range b.bexp2.GetFlowGraph() {
		bb := b.bexp2.BasicBlock[flow.GetEntryBasicBlockIndex()]
		for _, blockInst := range bb.GetInstructionIndex() {
			var prevAddr uint64
			if blockInst.BeginIndex != nil && blockInst.EndIndex != nil {
				for _, inst := range b.bexp2.Instruction[blockInst.GetBeginIndex():blockInst.GetEndIndex()] {
					if inst.Address != nil {
						mnemonic := b.bexp2.Mnemonic[inst.GetMnemonicIndex()]
						fmt.Printf("%#x: %s\n", inst.GetAddress(), mnemonic.GetName())
						prevAddr = inst.GetAddress()
					} else {
						mnemonic := b.bexp2.Mnemonic[inst.GetMnemonicIndex()]
						var out string
						for _, oidx := range inst.GetOperandIndex() {
							for _, eidx := range b.bexp2.Operand[oidx].GetExpressionIndex() {
								exp := b.bexp2.Expression[eidx]
								fmt.Printf("expression: %d) %s\n", eidx, exp)
								switch exp.GetType() {
								case BinExport2_Expression_SYMBOL:
									out += exp.GetSymbol()
								case BinExport2_Expression_IMMEDIATE_INT:
									out += fmt.Sprintf("%d", int64(exp.GetImmediate()))
								case BinExport2_Expression_IMMEDIATE_FLOAT:
									out += exp.GetSymbol()
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
	}
}
