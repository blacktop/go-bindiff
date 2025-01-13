package bindiff

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Metadata struct {
	Version     string
	File1       int
	File2       int
	Description string
	Created     time.Time
	Modified    time.Time
	Similarity  float64
	Confidence  float64
}

type File struct {
	ID              int `gorm:"primaryKey"`
	Filename        string
	ExeFilename     string
	Hash            string
	Functions       int
	LibFunctions    int
	Calls           int
	BasicBlocks     int
	LibBasicBlocks  int
	Edges           int
	LibEdges        int
	Instructions    int
	LibInstructions int
}

type FunctionAlgorithm int

const (
	FAlgorithmNone                             FunctionAlgorithm = iota
	FAlgorithmNameHashMatching                 FunctionAlgorithm = 1
	FAlgorithmHashMatching                     FunctionAlgorithm = 2
	FAlgorithmEdgesFlowgraphMdIndex            FunctionAlgorithm = 3
	FAlgorithmEdgesCallgraphMdIndex            FunctionAlgorithm = 4
	FAlgorithmMdIndexMatchingFlowgraphTopDown  FunctionAlgorithm = 5
	FAlgorithmMdIndexMatchingFlowgraphBottomUp FunctionAlgorithm = 6
	FAlgorithmPrimeSignatureMatching           FunctionAlgorithm = 7
	FAlgorithmMdIndexMatchingCallGraphTopDown  FunctionAlgorithm = 8
	FAlgorithmMdIndexMatchingCallGraphBottomUp FunctionAlgorithm = 9
	FAlgorithmRelaxedMdIndexMatching           FunctionAlgorithm = 10
	FAlgorithmInstructionCount                 FunctionAlgorithm = 11
	FAlgorithmAddressSequence                  FunctionAlgorithm = 12
	FAlgorithmStringReferences                 FunctionAlgorithm = 13
	FAlgorithmLoopCountMatching                FunctionAlgorithm = 14
	FAlgorithmCallSequenceMatchingExact        FunctionAlgorithm = 15
	FAlgorithmCallSequenceMatchingTopology     FunctionAlgorithm = 16
	FAlgorithmCallSequenceMatchingSequence     FunctionAlgorithm = 17
	FAlgorithmCallReferenceMatching            FunctionAlgorithm = 18
	FAlgorithmManual                           FunctionAlgorithm = 19
)

type FunctionMatch struct {
	ID             int             `gorm:"primaryKey"`
	Address1       decimal.Decimal `gorm:"type:BIGINT"`
	Name1          string
	Address2       decimal.Decimal `gorm:"type:BIGINT"`
	Name2          string
	Similarity     float64
	Confidence     float64
	Flags          int
	Algorithm      FunctionAlgorithm
	Evalute        bool
	Commentsported bool
	Basicblocks    int
	Edges          int
	Instructions   int
}

func (fm FunctionMatch) String() string {
	return fmt.Sprintf(
		"FunctionMatch: %s (%#x) -> %s (%#x) similarity=[%.2f] confidence=[%.2f]",
		fm.Name1,
		uint64(fm.Address1.BigInt().Int64()),
		fm.Name2,
		uint64(fm.Address2.BigInt().Int64()),
		fm.Similarity,
		fm.Confidence,
	)
}

type BasicBlockAlgorithm int

const (
	BBAlgorithmNone                         BasicBlockAlgorithm = iota
	BBAlgorithmEdgesPrimeProduct            BasicBlockAlgorithm = 1
	BBAlgorithmHashMatchingFourInstMin      BasicBlockAlgorithm = 2
	BBAlgorithmPrimeMatchingFourInstMin     BasicBlockAlgorithm = 3
	BBAlgorithmCallReferenceMatching        BasicBlockAlgorithm = 4
	BBAlgorithmStringReferencesMatching     BasicBlockAlgorithm = 5
	BBAlgorithmEdgesMdIndexTopDown          BasicBlockAlgorithm = 6
	BBAlgorithmMdIndexMatchingTopDown       BasicBlockAlgorithm = 7
	BBAlgorithmEdgesMdIndexBottomUp         BasicBlockAlgorithm = 8
	BBAlgorithmMdIndexMatchingBottomUp      BasicBlockAlgorithm = 9
	BBAlgorithmRelaxedMdIndexMatching       BasicBlockAlgorithm = 10
	BBAlgorithmPrimeMatchingNoInstMin       BasicBlockAlgorithm = 11
	BBAlgorithmEdgesLengauerTarjanDominated BasicBlockAlgorithm = 12
	BBAlgorithmLoopEntryMatching            BasicBlockAlgorithm = 13
	BBAlgorithmSelfLoopMatching             BasicBlockAlgorithm = 14
	BBAlgorithmEntryPointMatching           BasicBlockAlgorithm = 15
	BBAlgorithmExitPointMatching            BasicBlockAlgorithm = 16
	BBAlgorithmInstructionCountMatching     BasicBlockAlgorithm = 17
	BBAlgorithmJumpSequenceMatching         BasicBlockAlgorithm = 18
	BBAlgorithmPropagationSizeOne           BasicBlockAlgorithm = 19
	BBAlgorithmManual                       BasicBlockAlgorithm = 20
)

type BasicBlockMatch struct {
	ID         int `gorm:"primaryKey"`
	FunctionID int
	Address1   decimal.Decimal `gorm:"type:BIGINT"`
	Address2   decimal.Decimal `gorm:"type:BIGINT"`
	Algorithm  BasicBlockAlgorithm
	Evalute    bool
}

func (bbm BasicBlockMatch) String() string {
	return fmt.Sprintf(
		"BasicBlockMatch: funcID(%d) %#x -> %#x",
		bbm.FunctionID,
		uint64(bbm.Address1.BigInt().Int64()),
		uint64(bbm.Address2.BigInt().Int64()),
	)
}

type Instruction struct {
	ID       int             `gorm:"primaryKey"`
	Address1 decimal.Decimal `gorm:"type:BIGINT"`
	Address2 decimal.Decimal `gorm:"type:BIGINT"`
}

func (i Instruction) String() string {
	return fmt.Sprintf(
		"Instruction: %#x -> %#x",
		uint64(i.Address1.BigInt().Int64()),
		uint64(i.Address2.BigInt().Int64()),
	)
}
