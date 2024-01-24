package template

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
)

// ErrNilTemplate is an error for when there is no
// template provided to render as HTML.
var ErrNilTemplate = errors.New("nil template")

// HTMLTemplate encapsulates basic html template rendering.
type HTMLTemplate struct {
	Tpl  *template.Template
	Data any
}

// AsHTML transforms a HTMLTemplate into html.
//
//nolint:gosec
func (v HTMLTemplate) AsHTML(_ context.Context) (HTML, error) {
	var empty HTML

	if v.Tpl == nil {
		return empty, ErrNilTemplate
	}

	var bs bytes.Buffer
	if err := v.Tpl.Execute(&bs, v.Data); err != nil {
		return empty, fmt.Errorf("execute template: %w", err)
	}

	return HTML(bs.String()), nil
}
