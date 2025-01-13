/*
Copyright © 2025 blacktop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/apex/log"
	clihander "github.com/apex/log/handlers/cli"
	"github.com/blacktop/go-bindiff"
	"github.com/spf13/cobra"
)

var (
	verbose bool
	output  string
)

func trim(in string) string {
	if len(in) > 60 {
		return in[:60] + "..."
	}
	return in
}

func init() {
	log.SetHandler(clihander.Default)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "V", false, "Enable verbose logging")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output directory")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "bindiff",
	Short:         "Run bindiff",
	Args:          cobra.ExactArgs(2),
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {

		if verbose {
			log.SetLevel(log.DebugLevel)
		}

		bd, err := bindiff.NewBinDiff(args[0], args[1], output)
		if err != nil {
			log.Fatalf("failed to create binexport: %v", err)
		}
		defer bd.Close()

		log.Info("Running bindiff...")
		if err := bd.Run(); err != nil {
			log.Fatalf("failed to run bindiff: %v", err)
		}

		log.Info("Reading bindiff results...")
		if err := bd.Read(); err != nil {
			log.Fatalf("failed to read bindiff: %v", err)
		}

		log.Info("bindiff complete")
		// w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		for _, f := range bd.FunctionMatches() {
			log.WithFields(log.Fields{
				"similarity": fmt.Sprintf("%.2f", f.Similarity),
				"confidence": fmt.Sprintf("%.2f", f.Confidence),
			}).Info(trim(f.Name2))
			// fmt.Fprintf(w, "%s\t%.2f\t%.2f\n", trim(f.Name2), f.Similarity, f.Confidence)
		}
		// w.Flush()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
