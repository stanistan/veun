package veun

import (
	"context"

	"github.com/stanistan/veun/internal/view"
)

// MustMemo renders a view into another view and will panic
// if there is an unhandled error anywhere in the view tree.
func MustMemo(v view.AsView) view.Raw {
	out, err := view.Render(context.Background(), v)
	if err != nil {
		panic(err)
	}

	return view.Raw(out)
}
