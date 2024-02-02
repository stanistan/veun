// Code generated by ./generate-elements in el.go with source: elements.txt; DO NOT EDIT.
package el

import (
	"context"

	"github.com/stanistan/veun"
)

// A is struct representation of the *a* HTML elmenent.
//
// This struct is autogenerated.
type A []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [A].
//
// This function is autogenerated.
func (v A) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "a"}.apply(v).View(ctx)
}

// applyToElement appends A to the given [*element].
//
// This function is autogenerated.
func (v A) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Area is struct representation of the *area* HTML elmenent.
//
// This struct is autogenerated.
type Area []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Area].
//
// This function is autogenerated.
func (v Area) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "area"}.apply(v).View(ctx)
}

// applyToElement appends Area to the given [*element].
//
// This function is autogenerated.
func (v Area) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Article is struct representation of the *article* HTML elmenent.
//
// This struct is autogenerated.
type Article []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Article].
//
// This function is autogenerated.
func (v Article) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "article"}.apply(v).View(ctx)
}

// applyToElement appends Article to the given [*element].
//
// This function is autogenerated.
func (v Article) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Aside is struct representation of the *aside* HTML elmenent.
//
// This struct is autogenerated.
type Aside []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Aside].
//
// This function is autogenerated.
func (v Aside) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "aside"}.apply(v).View(ctx)
}

// applyToElement appends Aside to the given [*element].
//
// This function is autogenerated.
func (v Aside) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Base is struct representation of the *base* HTML elmenent.
//
// This struct is autogenerated.
type Base []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Base].
//
// This function is autogenerated.
func (v Base) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "base"}.apply(v).View(ctx)
}

// applyToElement appends Base to the given [*element].
//
// This function is autogenerated.
func (v Base) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Body is struct representation of the *body* HTML elmenent.
//
// This struct is autogenerated.
type Body []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Body].
//
// This function is autogenerated.
func (v Body) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "body"}.apply(v).View(ctx)
}

// applyToElement appends Body to the given [*element].
//
// This function is autogenerated.
func (v Body) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Br is struct representation of the *br* HTML elmenent.
//
// This struct is autogenerated.
type Br []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Br].
//
// This function is autogenerated.
func (v Br) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "br"}.apply(v).View(ctx)
}

// applyToElement appends Br to the given [*element].
//
// This function is autogenerated.
func (v Br) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Caption is struct representation of the *caption* HTML elmenent.
//
// This struct is autogenerated.
type Caption []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Caption].
//
// This function is autogenerated.
func (v Caption) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "caption"}.apply(v).View(ctx)
}

// applyToElement appends Caption to the given [*element].
//
// This function is autogenerated.
func (v Caption) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Code is struct representation of the *code* HTML elmenent.
//
// This struct is autogenerated.
type Code []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Code].
//
// This function is autogenerated.
func (v Code) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "code"}.apply(v).View(ctx)
}

// applyToElement appends Code to the given [*element].
//
// This function is autogenerated.
func (v Code) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Col is struct representation of the *col* HTML elmenent.
//
// This struct is autogenerated.
type Col []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Col].
//
// This function is autogenerated.
func (v Col) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "col"}.apply(v).View(ctx)
}

// applyToElement appends Col to the given [*element].
//
// This function is autogenerated.
func (v Col) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Command is struct representation of the *command* HTML elmenent.
//
// This struct is autogenerated.
type Command []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Command].
//
// This function is autogenerated.
func (v Command) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "command"}.apply(v).View(ctx)
}

// applyToElement appends Command to the given [*element].
//
// This function is autogenerated.
func (v Command) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Dd is struct representation of the *dd* HTML elmenent.
//
// This struct is autogenerated.
type Dd []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Dd].
//
// This function is autogenerated.
func (v Dd) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "dd"}.apply(v).View(ctx)
}

// applyToElement appends Dd to the given [*element].
//
// This function is autogenerated.
func (v Dd) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Details is struct representation of the *details* HTML elmenent.
//
// This struct is autogenerated.
type Details []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Details].
//
// This function is autogenerated.
func (v Details) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "details"}.apply(v).View(ctx)
}

// applyToElement appends Details to the given [*element].
//
// This function is autogenerated.
func (v Details) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Div is struct representation of the *div* HTML elmenent.
//
// This struct is autogenerated.
type Div []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Div].
//
// This function is autogenerated.
func (v Div) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "div"}.apply(v).View(ctx)
}

// applyToElement appends Div to the given [*element].
//
// This function is autogenerated.
func (v Div) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Dl is struct representation of the *dl* HTML elmenent.
//
// This struct is autogenerated.
type Dl []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Dl].
//
// This function is autogenerated.
func (v Dl) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "dl"}.apply(v).View(ctx)
}

// applyToElement appends Dl to the given [*element].
//
// This function is autogenerated.
func (v Dl) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Dt is struct representation of the *dt* HTML elmenent.
//
// This struct is autogenerated.
type Dt []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Dt].
//
// This function is autogenerated.
func (v Dt) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "dt"}.apply(v).View(ctx)
}

// applyToElement appends Dt to the given [*element].
//
// This function is autogenerated.
func (v Dt) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Em is struct representation of the *em* HTML elmenent.
//
// This struct is autogenerated.
type Em []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Em].
//
// This function is autogenerated.
func (v Em) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "em"}.apply(v).View(ctx)
}

// applyToElement appends Em to the given [*element].
//
// This function is autogenerated.
func (v Em) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Embed is struct representation of the *embed* HTML elmenent.
//
// This struct is autogenerated.
type Embed []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Embed].
//
// This function is autogenerated.
func (v Embed) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "embed"}.apply(v).View(ctx)
}

// applyToElement appends Embed to the given [*element].
//
// This function is autogenerated.
func (v Embed) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// H1 is struct representation of the *h1* HTML elmenent.
//
// This struct is autogenerated.
type H1 []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [H1].
//
// This function is autogenerated.
func (v H1) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "h1"}.apply(v).View(ctx)
}

// applyToElement appends H1 to the given [*element].
//
// This function is autogenerated.
func (v H1) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// H2 is struct representation of the *h2* HTML elmenent.
//
// This struct is autogenerated.
type H2 []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [H2].
//
// This function is autogenerated.
func (v H2) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "h2"}.apply(v).View(ctx)
}

// applyToElement appends H2 to the given [*element].
//
// This function is autogenerated.
func (v H2) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// H3 is struct representation of the *h3* HTML elmenent.
//
// This struct is autogenerated.
type H3 []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [H3].
//
// This function is autogenerated.
func (v H3) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "h3"}.apply(v).View(ctx)
}

// applyToElement appends H3 to the given [*element].
//
// This function is autogenerated.
func (v H3) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// H4 is struct representation of the *h4* HTML elmenent.
//
// This struct is autogenerated.
type H4 []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [H4].
//
// This function is autogenerated.
func (v H4) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "h4"}.apply(v).View(ctx)
}

// applyToElement appends H4 to the given [*element].
//
// This function is autogenerated.
func (v H4) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// H5 is struct representation of the *h5* HTML elmenent.
//
// This struct is autogenerated.
type H5 []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [H5].
//
// This function is autogenerated.
func (v H5) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "h5"}.apply(v).View(ctx)
}

// applyToElement appends H5 to the given [*element].
//
// This function is autogenerated.
func (v H5) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// H6 is struct representation of the *h6* HTML elmenent.
//
// This struct is autogenerated.
type H6 []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [H6].
//
// This function is autogenerated.
func (v H6) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "h6"}.apply(v).View(ctx)
}

// applyToElement appends H6 to the given [*element].
//
// This function is autogenerated.
func (v H6) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// HTML is struct representation of the *html* HTML elmenent.
//
// This struct is autogenerated.
type HTML []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [HTML].
//
// This function is autogenerated.
func (v HTML) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "html"}.apply(v).View(ctx)
}

// applyToElement appends HTML to the given [*element].
//
// This function is autogenerated.
func (v HTML) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Head is struct representation of the *head* HTML elmenent.
//
// This struct is autogenerated.
type Head []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Head].
//
// This function is autogenerated.
func (v Head) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "head"}.apply(v).View(ctx)
}

// applyToElement appends Head to the given [*element].
//
// This function is autogenerated.
func (v Head) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Hr is struct representation of the *hr* HTML elmenent.
//
// This struct is autogenerated.
type Hr []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Hr].
//
// This function is autogenerated.
func (v Hr) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "hr"}.apply(v).View(ctx)
}

// applyToElement appends Hr to the given [*element].
//
// This function is autogenerated.
func (v Hr) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Img is struct representation of the *img* HTML elmenent.
//
// This struct is autogenerated.
type Img []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Img].
//
// This function is autogenerated.
func (v Img) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "img"}.apply(v).View(ctx)
}

// applyToElement appends Img to the given [*element].
//
// This function is autogenerated.
func (v Img) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Input is struct representation of the *input* HTML elmenent.
//
// This struct is autogenerated.
type Input []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Input].
//
// This function is autogenerated.
func (v Input) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "input"}.apply(v).View(ctx)
}

// applyToElement appends Input to the given [*element].
//
// This function is autogenerated.
func (v Input) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Keygen is struct representation of the *keygen* HTML elmenent.
//
// This struct is autogenerated.
type Keygen []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Keygen].
//
// This function is autogenerated.
func (v Keygen) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "keygen"}.apply(v).View(ctx)
}

// applyToElement appends Keygen to the given [*element].
//
// This function is autogenerated.
func (v Keygen) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Li is struct representation of the *li* HTML elmenent.
//
// This struct is autogenerated.
type Li []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Li].
//
// This function is autogenerated.
func (v Li) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "li"}.apply(v).View(ctx)
}

// applyToElement appends Li to the given [*element].
//
// This function is autogenerated.
func (v Li) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Link is struct representation of the *link* HTML elmenent.
//
// This struct is autogenerated.
type Link []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Link].
//
// This function is autogenerated.
func (v Link) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "link"}.apply(v).View(ctx)
}

// applyToElement appends Link to the given [*element].
//
// This function is autogenerated.
func (v Link) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Meta is struct representation of the *meta* HTML elmenent.
//
// This struct is autogenerated.
type Meta []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Meta].
//
// This function is autogenerated.
func (v Meta) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "meta"}.apply(v).View(ctx)
}

// applyToElement appends Meta to the given [*element].
//
// This function is autogenerated.
func (v Meta) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Nav is struct representation of the *nav* HTML elmenent.
//
// This struct is autogenerated.
type Nav []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Nav].
//
// This function is autogenerated.
func (v Nav) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "nav"}.apply(v).View(ctx)
}

// applyToElement appends Nav to the given [*element].
//
// This function is autogenerated.
func (v Nav) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Ol is struct representation of the *ol* HTML elmenent.
//
// This struct is autogenerated.
type Ol []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Ol].
//
// This function is autogenerated.
func (v Ol) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "ol"}.apply(v).View(ctx)
}

// applyToElement appends Ol to the given [*element].
//
// This function is autogenerated.
func (v Ol) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// P is struct representation of the *p* HTML elmenent.
//
// This struct is autogenerated.
type P []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [P].
//
// This function is autogenerated.
func (v P) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "p"}.apply(v).View(ctx)
}

// applyToElement appends P to the given [*element].
//
// This function is autogenerated.
func (v P) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Picture is struct representation of the *picture* HTML elmenent.
//
// This struct is autogenerated.
type Picture []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Picture].
//
// This function is autogenerated.
func (v Picture) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "picture"}.apply(v).View(ctx)
}

// applyToElement appends Picture to the given [*element].
//
// This function is autogenerated.
func (v Picture) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Pre is struct representation of the *pre* HTML elmenent.
//
// This struct is autogenerated.
type Pre []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Pre].
//
// This function is autogenerated.
func (v Pre) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "pre"}.apply(v).View(ctx)
}

// applyToElement appends Pre to the given [*element].
//
// This function is autogenerated.
func (v Pre) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Script is struct representation of the *script* HTML elmenent.
//
// This struct is autogenerated.
type Script []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Script].
//
// This function is autogenerated.
func (v Script) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "script"}.apply(v).View(ctx)
}

// applyToElement appends Script to the given [*element].
//
// This function is autogenerated.
func (v Script) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Section is struct representation of the *section* HTML elmenent.
//
// This struct is autogenerated.
type Section []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Section].
//
// This function is autogenerated.
func (v Section) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "section"}.apply(v).View(ctx)
}

// applyToElement appends Section to the given [*element].
//
// This function is autogenerated.
func (v Section) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Source is struct representation of the *source* HTML elmenent.
//
// This struct is autogenerated.
type Source []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Source].
//
// This function is autogenerated.
func (v Source) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "source"}.apply(v).View(ctx)
}

// applyToElement appends Source to the given [*element].
//
// This function is autogenerated.
func (v Source) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Span is struct representation of the *span* HTML elmenent.
//
// This struct is autogenerated.
type Span []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Span].
//
// This function is autogenerated.
func (v Span) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "span"}.apply(v).View(ctx)
}

// applyToElement appends Span to the given [*element].
//
// This function is autogenerated.
func (v Span) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Strong is struct representation of the *strong* HTML elmenent.
//
// This struct is autogenerated.
type Strong []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Strong].
//
// This function is autogenerated.
func (v Strong) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "strong"}.apply(v).View(ctx)
}

// applyToElement appends Strong to the given [*element].
//
// This function is autogenerated.
func (v Strong) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Style is struct representation of the *style* HTML elmenent.
//
// This struct is autogenerated.
type Style []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Style].
//
// This function is autogenerated.
func (v Style) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "style"}.apply(v).View(ctx)
}

// applyToElement appends Style to the given [*element].
//
// This function is autogenerated.
func (v Style) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Summary is struct representation of the *summary* HTML elmenent.
//
// This struct is autogenerated.
type Summary []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Summary].
//
// This function is autogenerated.
func (v Summary) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "summary"}.apply(v).View(ctx)
}

// applyToElement appends Summary to the given [*element].
//
// This function is autogenerated.
func (v Summary) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// TBody is struct representation of the *tbody* HTML elmenent.
//
// This struct is autogenerated.
type TBody []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [TBody].
//
// This function is autogenerated.
func (v TBody) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "tbody"}.apply(v).View(ctx)
}

// applyToElement appends TBody to the given [*element].
//
// This function is autogenerated.
func (v TBody) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// TFoot is struct representation of the *tfoot* HTML elmenent.
//
// This struct is autogenerated.
type TFoot []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [TFoot].
//
// This function is autogenerated.
func (v TFoot) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "tfoot"}.apply(v).View(ctx)
}

// applyToElement appends TFoot to the given [*element].
//
// This function is autogenerated.
func (v TFoot) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// THead is struct representation of the *thead* HTML elmenent.
//
// This struct is autogenerated.
type THead []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [THead].
//
// This function is autogenerated.
func (v THead) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "thead"}.apply(v).View(ctx)
}

// applyToElement appends THead to the given [*element].
//
// This function is autogenerated.
func (v THead) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Table is struct representation of the *table* HTML elmenent.
//
// This struct is autogenerated.
type Table []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Table].
//
// This function is autogenerated.
func (v Table) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "table"}.apply(v).View(ctx)
}

// applyToElement appends Table to the given [*element].
//
// This function is autogenerated.
func (v Table) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Td is struct representation of the *td* HTML elmenent.
//
// This struct is autogenerated.
type Td []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Td].
//
// This function is autogenerated.
func (v Td) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "td"}.apply(v).View(ctx)
}

// applyToElement appends Td to the given [*element].
//
// This function is autogenerated.
func (v Td) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Th is struct representation of the *th* HTML elmenent.
//
// This struct is autogenerated.
type Th []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Th].
//
// This function is autogenerated.
func (v Th) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "th"}.apply(v).View(ctx)
}

// applyToElement appends Th to the given [*element].
//
// This function is autogenerated.
func (v Th) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Tr is struct representation of the *tr* HTML elmenent.
//
// This struct is autogenerated.
type Tr []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Tr].
//
// This function is autogenerated.
func (v Tr) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "tr"}.apply(v).View(ctx)
}

// applyToElement appends Tr to the given [*element].
//
// This function is autogenerated.
func (v Tr) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Track is struct representation of the *track* HTML elmenent.
//
// This struct is autogenerated.
type Track []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Track].
//
// This function is autogenerated.
func (v Track) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "track"}.apply(v).View(ctx)
}

// applyToElement appends Track to the given [*element].
//
// This function is autogenerated.
func (v Track) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Ul is struct representation of the *ul* HTML elmenent.
//
// This struct is autogenerated.
type Ul []Param

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Ul].
//
// This function is autogenerated.
func (v Ul) View(ctx context.Context) (*veun.View, error) {
	return element{tag: "ul"}.apply(v).View(ctx)
}

// applyToElement appends Ul to the given [*element].
//
// This function is autogenerated.
func (v Ul) applyToElement(e *element) {
	e.children = append(e.children, v)
}

// Wbr is struct representation of the *wbr* HTML elmenent.
//
// This struct is autogenerated.
type Wbr []VoidParam

// View fulfills the [veun.AsView] interface. It applies all of
// the options for [Wbr].
//
// This function is autogenerated.
func (v Wbr) View(ctx context.Context) (*veun.View, error) {
	return voidElement{tag: "wbr"}.apply(v).View(ctx)
}

// applyToElement appends Wbr to the given [*element].
//
// This function is autogenerated.
func (v Wbr) applyToElement(e *element) {
	e.children = append(e.children, v)
}