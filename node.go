package hax

/*
type Event string

const (
	EventClick  Event = "click"
	EventInput  Event = "input"
	EventChange Event = "change"
)
*/

type Gettable[T any] interface {
	Get() T
}

type Settable[T any] interface {
	Set(T)
}

type INode interface {
	Id(id string) INode

	Text(t string) INode

	Body(childs ...INode) INode
	BodyList(childs []INode) INode

	Attribute(key, value string) INode
	RemoveAttribute(key string) INode
	Style(key, value string) INode
	RemoveStyle(key string) INode
	Class(classes ...string) INode
	RemoveClass(classes ...string) INode

	For(value string) INode

	/*
		On(event Event, handler func(ctx EventContext)) INode
		OnClick(handler func(ctx EventContext)) INode
	*/
}

type VNode struct {
	tag  string
	id   string
	text string

	styles     map[string]string
	classes    map[string]struct{}
	attributes map[string]string
	children   []*VNode
}

func newVNode(tag string) *VNode {
	return &VNode{
		tag:        tag,
		styles:     map[string]string{},
		classes:    map[string]struct{}{},
		attributes: map[string]string{},
		children:   []*VNode{},
	}
}

func (vn *VNode) Id(id string) INode {
	vn.id = id
	return vn
}

func (vn *VNode) Text(t string) INode {
	vn.text = t
	return vn
}

func (vn *VNode) Body(childs ...INode) INode {
	return vn.BodyList(childs)
}

func (vn *VNode) BodyList(childs []INode) INode {
	vn.children = []*VNode{}
	for _, n := range childs {
		vn.children = append(vn.children, asVNode(n))
	}
	return vn
}

func (vn *VNode) Attribute(key, value string) INode {
	vn.attributes[key] = value
	return vn
}

func (vn *VNode) RemoveAttribute(key string) INode {
	delete(vn.attributes, key)
	return vn
}

func (vn *VNode) Style(key, value string) INode {
	vn.styles[key] = value
	return vn
}

func (vn *VNode) RemoveStyle(key string) INode {
	delete(vn.styles, key)
	return vn
}

func (vn *VNode) Class(classes ...string) INode {
	for _, cls := range classes {
		vn.classes[cls] = struct{}{}
	}
	return vn
}

func (vn *VNode) RemoveClass(classes ...string) INode {
	for _, cls := range classes {
		delete(vn.classes, cls)
	}
	return vn
}

func (vn *VNode) For(value string) INode {
	vn.Attribute("for", value)
	return vn
}

type AVNode struct {
	VNode
}

func asA(node *VNode) *AVNode {
	return &AVNode{*node}
}

func (e *AVNode) Href(v string) *AVNode {
	e.Attribute("href", v)
	return e
}

type InputVNode struct {
	VNode
}

func asInput(node *VNode) *InputVNode {
	return &InputVNode{*node}
}

func (element *InputVNode) Value(t string) *InputVNode {
	element.Attribute("value", t)
	return element
}

func (e *InputVNode) Placeholder(v string) *InputVNode {
	e.Attribute("placeholder", v)
	return e
}

func (e *InputVNode) Type(v string) *InputVNode {
	e.Attribute("type", v)
	return e
}

type TextAreaNode struct {
	VNode
}

func asTextArea(node *VNode) *TextAreaNode {
	return &TextAreaNode{
		*node,
	}
}

func (element *TextAreaNode) Value(t string) INode {
	element.Attribute("value", t)
	return element
}

type FormNode struct {
	VNode
}

func asForm(node *VNode) *FormNode {
	return &FormNode{
		*node,
	}
}

func (element *FormNode) Action(endpoint string) *FormNode {
	element.Attribute("action", endpoint)
	return element
}

func (element *FormNode) Method(method string) *FormNode {
	element.Attribute("method", method)
	return element
}

func asVNode(i INode) *VNode {
	switch n := i.(type) {
	case *VNode:
		return n
	case *InputVNode:
		return &n.VNode
	case *AVNode:
		return &n.VNode
	case *TextAreaNode:
		return &n.VNode
	case *FormNode:
		return &n.VNode
	default:
		return nil
	}
}

func H1() *VNode { return newVNode("h1") }
func H2() *VNode { return newVNode("h2") }
func H3() *VNode { return newVNode("h3") }
func H4() *VNode { return newVNode("h4") }
func H5() *VNode { return newVNode("h5") }
func H6() *VNode { return newVNode("h6") }

func Div() *VNode { return newVNode("div") }

func P() *VNode { return newVNode("p") }

func Button() *VNode { return newVNode("button") }

func A() *AVNode              { return asA(newVNode("a")) }
func Span() *VNode            { return newVNode("span") }
func Strong() *VNode          { return newVNode("strong") }
func Em() *VNode              { return newVNode("em") }
func Small() *VNode           { return newVNode("small") }
func Img() *VNode             { return newVNode("img") }
func Input() *InputVNode      { return asInput(newVNode("input")) }
func Label() *VNode           { return newVNode("label") }
func Form() *FormNode         { return asForm(newVNode("form")) }
func Select() *VNode          { return newVNode("select") }
func Option() *VNode          { return newVNode("option") }
func TextArea() *TextAreaNode { return asTextArea(newVNode("textarea")) }
func Output() *VNode          { return newVNode("output") }

func Ul() *VNode { return newVNode("ul") }
func Ol() *VNode { return newVNode("ol") }
func Li() *VNode { return newVNode("li") }

func Table() *VNode { return newVNode("table") }
func THead() *VNode { return newVNode("thead") }
func TBody() *VNode { return newVNode("tbody") }
func TFoot() *VNode { return newVNode("tfoot") }
func Tr() *VNode    { return newVNode("tr") }
func Th() *VNode    { return newVNode("th") }
func Td() *VNode    { return newVNode("td") }

func Nav() *VNode     { return newVNode("nav") }
func Header() *VNode  { return newVNode("header") }
func Footer() *VNode  { return newVNode("footer") }
func Main() *VNode    { return newVNode("main") }
func Section() *VNode { return newVNode("section") }
func Article() *VNode { return newVNode("artible") }
func Aside() *VNode   { return newVNode("aside") }

func Video() *VNode  { return newVNode("video") }
func Audio() *VNode  { return newVNode("audio") }
func Source() *VNode { return newVNode("source") }

func Canvas() *VNode { return newVNode("canvas") }
func Svg() *VNode    { return newVNode("svg") }
func Path() *VNode   { return newVNode("path") }

func Br() *VNode { return newVNode("br") }
func Hr() *VNode { return newVNode("hr") }

func Code() *VNode { return newVNode("code") }
func Pre() *VNode  { return newVNode("pre") }
