package el

import (
	"html/template"
	"strings"
)

// Attrs represents an attribute map.
//
// Attrs can be applied to both void and non-void elements,
// as such, attributes implement [Param] and [VoidParam].
//
// Applying this to any element merges/copies this set of
// attributes into the element's.
type Attrs map[string]string

func (a Attrs) render() string {
	var w strings.Builder

	for k, v := range a {
		_, _ = w.WriteString(" ")
		template.HTMLEscape(&w, []byte(k))
		_, _ = w.WriteString(`="`)
		template.HTMLEscape(&w, []byte(v))
		_, _ = w.WriteString(`"`)
	}

	return w.String()
}

func (a Attrs) mergeInto(attrs Attrs) {
	for k, v := range a {
		attrs[k] = v
	}
}

func (a Attrs) applyToElement(e *element[nodeChildren]) { e.attrs(a.mergeInto) }
func (a Attrs) applyToVoidElement(e *element[void])     { e.attrs(a.mergeInto) }

// Attr is a single key value attribute.
//
// It is both a [Param] and a [VoidParam].
type Attr struct {
	Key, Value string
}

func (a Attr) update(attrs Attrs)                      { attrs[a.Key] = a.Value }
func (a Attr) applyToElement(e *element[nodeChildren]) { e.attrs(a.update) }
func (a Attr) applyToVoidElement(e *element[void])     { e.attrs(a.update) }

// AttrWith holds a mutation function for the current
// element's attribute value.
//
// Rhe  update function will receive the old value
// (or an empty string) to merge with the new one.
//
// This is useful to create transformers such as [Class].
//
// AttrWith can be applied to both void and non-void elements.
func AttrWith(name string, update func(string) string) AttrFunc {
	return AttrFunc(func(a Attrs) { a[name] = update(a[name]) })
}

// AttrFunc can update and modify any elements attributes.
//
// It is both a [Param] and a [VoidParam] and can be used
// to make custom modifiers for an element's properties.
type AttrFunc func(Attrs)

func (f AttrFunc) applyToElement(e *element[nodeChildren]) { e.attrs(f) }
func (f AttrFunc) applyToVoidElement(e *element[void])     { e.attrs(f) }

// ClearAttr deletes the key from the element's attributes.
func ClearAttr(name string) AttrFunc {
	return AttrFunc(func(a Attrs) { delete(a, name) })
}

// Class refers to an HTML class property.
//
// This is an attribute that will merge onto the class property
// of the element and is represented as a slice instead of a string.
//
// This is the simplest version of this for now.
func Class(names ...string) AttrFunc {
	return AttrWith("class", func(old string) string {
		if len(old) > 0 {
			return strings.Join(names, " ") + " " + old
		}

		return strings.Join(names, " ")
	})
}

// ID is the ID attribute.
func ID(id string) Attr {
	return Attr{Key: "id", Value: id}
}

// Href creates an href [Attr] with the given value.
func Href(href string) Attr {
	return Attr{Key: "href", Value: href}
}
