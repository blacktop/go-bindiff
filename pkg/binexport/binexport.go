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

func (b *BinExport) Run() error {
	data, err := os.ReadFile(b.path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	var exp BinExport2
	if err := proto.Unmarshal(data, &exp); err != nil {
		return fmt.Errorf("failed to unmarshal BinExport2: %w", err)
	}
	_ = exp
	return nil
}
