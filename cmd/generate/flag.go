package generate

import (
	"fmt"

	"github.com/spf13/cobra"

	cmderror "github.com/landro/schemadocs/pkg/error"
)

const (
	flagOutputPath          = "output-path"
	flagDocPlaceholderStart = "doc-placeholder-start"
	flagDocPlaceholderEnd   = "doc-placeholder-end"
	flagLayout              = "layout"
	flagOtherSectionTop     = "other-section-top"
)

type flag struct {
	outputPath          string
	docPlaceholderStart string
	docPlaceholderEnd   string
	layout              string
	otherSectionTop     bool
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.outputPath, flagOutputPath, "o", "", "Path to file to output the generated documentation")
	cmd.Flags().StringVar(&f.docPlaceholderStart, flagDocPlaceholderStart, "", "Placeholder string marking the start of the docs section in the output file")
	cmd.Flags().StringVar(&f.docPlaceholderEnd, flagDocPlaceholderEnd, "", "Placeholder string marking the end of the docs section in the output file")
	cmd.Flags().StringVarP(&f.layout, flagLayout, "l", "default", "Layout of the generated documentation")
	cmd.Flags().BoolVar(&f.otherSectionTop, flagOtherSectionTop, false, "Place the 'Other' section at the top instead of the bottom")
}

func (f *flag) Validate() error {
	if (f.docPlaceholderStart == "") != (f.docPlaceholderEnd == "") {
		return fmt.Errorf("both --%s and --%s flags must be set to non-empty values: %w", flagDocPlaceholderStart, flagDocPlaceholderEnd, cmderror.ErrInvalidFlag)
	}
	return nil
}
