package hax

import (
	"html"
	"strings"
)

var voidElements = map[string]bool{
	"area": true, "base": true, "br": true, "col": true,
	"embed": true, "hr": true, "img": true, "input": true,
	"link": true, "meta": true, "param": true, "source": true,
	"track": true, "wbr": true,
}

type StringRenderer struct {
	buff strings.Builder
}

func (ssr *StringRenderer) Render(node INode) {
	current := asVNode(node)
	ssr.writeOpenTag(current)
	if !voidElements[current.tag] {
		ssr.writeBody(current)
	}
	ssr.writeCloseTag(current)
}

func (ssr *StringRenderer) writeBody(current *VNode) {
	for _, child := range current.children {
		ssr.Render(child)
	}
	ssr.buff.WriteString(html.EscapeString(current.text))
}

func (ssr *StringRenderer) writeCloseTag(current *VNode) {
	if voidElements[current.tag] {
		return
	}
	ssr.buff.WriteString("</")
	ssr.buff.WriteString(current.tag)
	ssr.buff.WriteString(">")
}

func (ssr *StringRenderer) writeOpenTag(current *VNode) {
	ssr.buff.WriteRune('<')
	ssr.buff.WriteString(current.tag)
	ssr.writeAttributes(current)

	if voidElements[current.tag] {
		ssr.buff.WriteString("/>")
	} else {
		ssr.buff.WriteRune('>')
	}
}

func (ssr *StringRenderer) writeAttributes(current *VNode) {
	if len(current.classes) > 0 {
		ssr.buff.WriteString(` class="`)
		classesLast := len(current.classes) - 1
		i := 0
		for cls := range current.classes {
			ssr.buff.WriteString(html.EscapeString(cls))
			if i < classesLast {
				ssr.buff.WriteRune(' ')
			}
			i++
		}
		ssr.buff.WriteString(`"`)
	}

	if len(current.styles) > 0 {
		ssr.buff.WriteString(` style="`)
		for styleName, styleValue := range current.styles {
			ssr.buff.WriteString(html.EscapeString(styleName))
			ssr.buff.WriteString(": ")
			ssr.buff.WriteString(html.EscapeString(styleValue))
			ssr.buff.WriteRune(';')
		}
		ssr.buff.WriteString(`"`)
	}

	if len(current.attributes) > 0 {
		for attrName, attrValue := range current.attributes {
			ssr.buff.WriteString(" ")
			ssr.buff.WriteString(attrName)
			ssr.buff.WriteString(`="`)
			ssr.buff.WriteString(html.EscapeString(attrValue))
			ssr.buff.WriteString(`"`)
		}
	}

	if len(current.id) > 0 {
		ssr.buff.WriteString(` id="`)
		ssr.buff.WriteString(current.id)
		ssr.buff.WriteString(`"`)
	}
}

func (ssr *StringRenderer) String() string {
	return ssr.buff.String()
}

func (ssr *StringRenderer) Reset() {
	ssr.buff.Reset()
}
