package bindiff

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/blacktop/go-bindiff/pkg/binexport"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	bindiffExePath = "/Applications/BinDiff/BinDiff.app/Contents/MacOS/bin/bindiff"
)

type BinDiff struct {
	ExportPrev string
	ExportNext string
	DiffFile   string
	Output     string

	md        Metadata
	files     []File
	fmatches  []FunctionMatch
	bbmatches []BasicBlockMatch
	imatches  []Instruction

	db *gorm.DB
}

func fileNameWithoutExtTrimSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func NewBinDiff(prev, next, output string) (*BinDiff, error) {
	bx1, err := binexport.NewBinExport(prev)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", prev, err)
	}
	bx2, err := binexport.NewBinExport(next)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", next, err)
	}
	return &BinDiff{
		ExportPrev: prev,
		ExportNext: next,
		DiffFile: filepath.Join(output,
			fmt.Sprintf("%s_vs_%s.BinDiff",
				fileNameWithoutExtTrimSuffix(bx1.ExecutableName()),
				fileNameWithoutExtTrimSuffix(bx2.ExecutableName()),
			)),
		Output: output,
	}, nil
}

func (b *BinDiff) Close() error {
	db, err := b.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (b *BinDiff) Run() error {
	cmd := exec.Command(bindiffExePath, "--primary", b.ExportPrev, "--secondary", b.ExportNext, "--output_dir", b.Output)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run bindiff: %v: cmd output - '%s'", err, out)
	}
	return nil
}

func (b *BinDiff) Read() (err error) {
	b.db, err = gorm.Open(sqlite.Open(b.DiffFile), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("failed to connect bindiff sqlite database: %w", err)
	}

	b.db.Table("metadata").First(&b.md)
	b.db.Table("file").Find(&b.files)
	b.db.Table("function").Find(&b.fmatches)
	b.db.Table("basicblock").Find(&b.bbmatches)
	b.db.Table("instruction").Find(&b.imatches)

	return nil
}

func (b *BinDiff) Metadata() Metadata {
	return b.md
}

func (b *BinDiff) Files() []File {
	return b.files
}

func (b *BinDiff) FunctionMatches() []FunctionMatch {
	return b.fmatches
}

func (b *BinDiff) BasicBlockMatches() []BasicBlockMatch {
	return b.bbmatches
}

func (b *BinDiff) Instructions() []Instruction {
	return b.imatches
}
