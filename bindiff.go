package bindiff

type BinDiff struct {
	ExportPrev string
	ExportNext string
}

func NewBinDiff(prev, next string) *BinDiff {
	return &BinDiff{
		ExportPrev: prev,
		ExportNext: next,
	}
}

func (b *BinDiff) Run() error {
	return nil
}
