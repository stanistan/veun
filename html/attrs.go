package html

import (
	"html/template"
	"strings"
)

type Attrs map[string]string

func openingTag(name string, a Attrs) string {
	var sb strings.Builder

	sb.WriteString("<")
	sb.WriteString(name)

	for k, v := range a {
		sb.WriteString(" ")
		template.HTMLEscape(&sb, []byte(k))
		sb.WriteString(`="`)
		template.HTMLEscape(&sb, []byte(v))
		sb.WriteString(`"`)
	}

	sb.WriteString(">")

	return sb.String()
}

func closingTag(name string) string {
	return "</" + name + ">"
}
