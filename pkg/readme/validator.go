package readme

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"

	pkgerror "github.com/landro/schemadocs/pkg/error"
	"github.com/landro/schemadocs/pkg/generate"
)

func (r *Readme) Validate(schemaPath string, layout string, otherSectionTop bool) error {
	docsFromReadme, err := r.Docs()
	if err != nil {
		return fmt.Errorf("failed to read docs from readme: %w", err)
	}

	docsFromSchema, err := generate.Generate(schemaPath, layout, otherSectionTop)
	if err != nil {
		return fmt.Errorf("failed to generate docs from schema: %w", err)
	}

	diff := cmp.Diff(strings.TrimSpace(docsFromSchema), trimDocs(docsFromReadme, r.startPlaceholder, r.endPlaceholder))
	if diff != "" {
		return fmt.Errorf("documentation from readme %s do not match output generated from %s\n: %w", r.path, schemaPath, pkgerror.ErrInvalidDocs)
	}

	return nil
}

func trimDocs(docs, startPlaceholder, endPlaceholder string) string {
	docs = strings.TrimPrefix(docs, startPlaceholder)
	docs = strings.TrimSuffix(docs, endPlaceholder)
	return strings.TrimSpace(docs)
}
