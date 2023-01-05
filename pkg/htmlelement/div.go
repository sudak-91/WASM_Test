package htmlelement

import "syscall/js"

type Div struct {
	HtmlElement
	HtmlClass
	Id        string
	InnerHtml string
	Items     []js.Value
}

//NewDiv create Div's object
func NewDiv() *Div {
	var (
		div Div
	)
	div.elem = Document.Call("createElement", "div")
	return &div
}

//CreateChildDiv create html div and append to child tree
func (d *Div) CreateChildDiv() *Div {
	div := NewDiv()
	d.AddChild(div.GetJs())
	return div
}

func (d *Div) Render() {
	d.AddClassSliceToClassList(d.elem)
	for _, v := range d.ChildElement {
		v.Render()
		d.AppendChild(v)
	}
	for _, v := range d.Items {
		d.AppendChild(v)
	}

}
