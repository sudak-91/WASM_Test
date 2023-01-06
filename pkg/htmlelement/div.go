package htmlelement

type Div struct {
	HtmlElement
	HtmlClass
	Id        string
	InnerHtml string
}

//NewDiv create Div's object
func NewDiv() *Div {
	var (
		div Div
	)
	div.elem = GetDocument().Call("createElement", "div")
	return &div
}

//CreateChildDiv create html div and append to child tree
func (d *Div) CreateChildDiv() *Div {
	div := NewDiv()
	d.AddChild(div)
	return div
}

func (d *Div) Render() {
	d.AddClassSliceToClassList(d.elem)
	for _, v := range d.ChildElement {
		render, ok := v.(Renderer)
		if ok {
			render.Render()
		}
		elem, ok := v.(JsGeter)
		if ok {
			d.AppendChild(elem.GetJs())
		}
	}
}
