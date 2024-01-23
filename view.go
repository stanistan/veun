package veun

import (
	"context"
	"fmt"
	"html/template"
)

// View composes an HTMLRenderable with an ErrorHandler.
type View struct {
	r  HTMLRenderable
	eh ErrorHandler
}

// View conforms View to AsView.
func (r *View) View(_ context.Context) (*View, error) {
	return r, nil
}

// WithErrorHandler creates a new View with the error handler.
func (r *View) WithErrorHandler(eh ErrorHandler) *View {
	return &View{r: r.r, eh: eh}
}

func (r *View) render(ctx context.Context) (template.HTML, error) {
	var empty template.HTML

	if r == nil {
		return empty, nil
	}

	if r.r == nil {
		return empty, nil
	}

	out, err := r.r.AsHTML(ctx)
	if err != nil {
		return RenderError(ctx, r.eh, err)
	}

	return out, nil
}

type viewInvalidError struct {
	Err error
}

func (e viewInvalidError) Error() string {
	return e.Err.Error()
}

func (e viewInvalidError) AsHTML(_ context.Context) (template.HTML, error) {
	return template.HTML(""), e
}

type renderable struct {
	r AsView
}

func (r renderable) AsHTML(ctx context.Context) (template.HTML, error) {
	var empty template.HTML

	if r.r == nil {
		return empty, nil
	}

	v, err := r.r.View(ctx)
	if err != nil {
		return empty, fmt.Errorf("View(%T): %w", r.r, err)
	}

	out, err := v.render(ctx)
	if err != nil {
		return out, fmt.Errorf("Render(%T): %w", r.r, err)
	}

	return out, nil
}
