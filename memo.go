package veun

import "context"

func MustMemo(v AsView) Raw {
	out, err := Render(context.Background(), v)
	if err != nil {
		panic(err)
	}

	return Raw(out)
}
