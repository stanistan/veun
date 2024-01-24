package veun

import "context"

// MustMemo renders a view into another view and will panic
// if there is an unhandled error anywhere in the view tree.
func MustMemo(v AsView) Raw {
	out, err := Render(context.Background(), v)
	if err != nil {
		panic(err)
	}

	return Raw(out)
}
