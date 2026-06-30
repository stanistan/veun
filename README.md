# veun

> A small server-side HTML template composition library for Go.

Read more: [Building View-Trees in Go](https://www.stanistan.com/writes/building-view-trees-in-go-part-1/) (7-part series) · [Demo](https://veun-http-demo.stanistan.com) · [Demo source](https://github.com/stanistan/veun-http-demo)

## Overview

veun is a composable view library built on top of Go's standard `html/template` package. The core idea: a `View` pairs an HTML renderable with an optional error handler, and views compose into trees.

## Packages

| Package | Purpose |
|---|---|
| `github.com/stanistan/veun` | Core types and `Render` function |
| `github.com/stanistan/veun/template` | `html/template`-based views with slot composition |
| `github.com/stanistan/veun/el` | Functional HTML element construction (codegen'd) |
| `github.com/stanistan/veun/vhttp` | `http.Handler` integration |
| `github.com/stanistan/veun/vhttp/handler` | Handler utilities (checked, path, status) |
| `github.com/stanistan/veun/vhttp/request` | Request handler interface |

## Core Concepts

### AsView

Any type implementing `View(context.Context) (*View, error)` is an `AsView` and can be rendered.

```go
type personView struct{ Name string }

func (v personView) View(_ context.Context) (*veun.View, error) {
    tpl := template.Must(template.New("").Parse(`<div>Hi, {{ .Name }}.</div>`))
    return veun.V(t.Template{Tpl: tpl, Data: v}), nil
}

html, err := veun.Render(ctx, personView{Name: "Stan"})
// html == `<div>Hi, Stan.</div>`
```

### Slot composition

Templates can render child views into named slots:

```go
var containerTpl = t.MustParse("container", `
    <div>
        <header>{{ slot "heading" }}</header>
        <main>{{ slot "body" }}</main>
    </div>
`)

func (v ContainerView) View(_ context.Context) (*veun.View, error) {
    return veun.V(t.Template{
        Tpl:   containerTpl,
        Slots: t.Slots{"heading": v.Heading, "body": v.Body},
    }), nil
}
```

### el — functional element construction

```go
view := el.Div{
    el.Class("container"),
    el.H1{el.Text("Hello")},
    el.P{el.Text("World")},
}
```

### Error handling

Any view can attach an error handler that either replaces failed content or propagates the error:

```go
view := veun.V(myView).WithErrorHandler(veun.ErrorHandlerFunc(
    func(ctx context.Context, err error) (veun.AsView, error) {
        return fallbackView{}, nil
    },
))
```

### HTTP handlers

```go
mux.Handle("/", vhttp.HandlerFunc(func(r *http.Request) (veun.AsView, http.Handler, error) {
    return myPageView{}, nil, nil
}))
```

## Install

```
go get github.com/stanistan/veun
```
