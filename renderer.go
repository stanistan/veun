package veun

import (
	"context"
	"fmt"
	"html/template"
)

func RenderToHTML(ctx context.Context, r Renderable, errHandler any) (template.HTML, error) {
	if r == nil {
		return emptyHTML(), nil
	}

	out, err := r.RenderToHTML(ctx)
	if err != nil {
		return RenderError(&Error{Err: err, ctx: ctx}, errHandler)
	}

	return out, nil
}

func Render(ctx context.Context, v AsRenderable) (template.HTML, error) {
	if v == nil {
		return emptyHTML(), nil
	}

	r, err := v.Renderable(ctx)
	if err != nil {
		return RenderError(&Error{
			Err: fmt.Errorf("%T.Renderable: %w", v, err),
			ctx: ctx,
		}, v)
	}

	out, err := RenderToHTML(ctx, r, v)
	if err != nil {
		return emptyHTML(), err
	}

	return out, nil
}
