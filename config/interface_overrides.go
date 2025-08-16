package config

import (
	"go/ast"
	"regexp"
	"strings"
)

var (
	generateRe   = regexp.MustCompile(`mockery_generate: (true|false)`)
	structNameRe = regexp.MustCompile(`mockery_struct_name: ([^\s]+)$`)
	fileNameRe   = regexp.MustCompile(`mockery_file_name: ([^\s]+)$`)
	templateRe   = regexp.MustCompile(`mockery_template: ([^\s]+)$`)
)

type InterfaceOverrides struct {
	// Generate indicates whether the interface should be generated.
	// mockery_generate: <true|false>
	Generate bool

	// StructName is the name of the struct to generate.
	// mockery_struct_name: <struct_name>
	StructName string

	// FileName is the name of the file to generate.
	// mockery_file_name: <file_name>
	FileName string

	// Template is the name of the template to use.
	// mockery_template: <template>
	Template string
}

// ExtractInterfaceOverrides parses the documentation from a declaration
// node and matches comments with mockery's patterns.
func ExtractInterfaceOverrides(decl *ast.GenDecl) *InterfaceOverrides {
	if decl == nil || decl.Doc == nil {
		return nil
	}

	overrides := &InterfaceOverrides{}
	for _, comment := range decl.Doc.List {
		text := strings.TrimSpace(comment.Text)

		if matches := generateRe.FindStringSubmatch(text); len(matches) > 1 {
			overrides.Generate = matches[1] == "true"
		} else if matches := structNameRe.FindStringSubmatch(text); len(matches) > 1 {
			overrides.StructName = strings.TrimSpace(matches[1])
		} else if matches := fileNameRe.FindStringSubmatch(text); len(matches) > 1 {
			overrides.FileName = strings.TrimSpace(matches[1])
		} else if matches := templateRe.FindStringSubmatch(text); len(matches) > 1 {
			overrides.Template = strings.TrimSpace(matches[1])
		}
	}

	return overrides
}

func (i *InterfaceOverrides) Override(cfg *Config) {
	if i == nil {
		return
	}

	if i.StructName != "" {
		cfg.StructName = &i.StructName
	}

	if i.FileName != "" {
		cfg.FileName = &i.FileName
	}

	if i.Template != "" {
		cfg.Template = &i.Template
	}
}

func (i *InterfaceOverrides) ShouldGenerate() bool {
	if i == nil {
		return false
	}

	return i.Generate
}
