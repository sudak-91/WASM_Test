package htmlelement

type Div struct {
	HtmlElement
	HtmlClass
	parent    Parent
	Id        string
	InnerHtml string
}

//NewDiv create Div's object
func NewDiv(parent Parent, id string) *Div {
	var (
		div Div
	)
	div.Id = id
	div.parent = parent
	parent.SetChild(&div)
	//div.elem = GetDocument().Call("createElement", "div")
	return &div
}

func (d *Div) Render() {
	d.elem = GetDocument().Call("createElement", "div")
	d.SetId(d.Id)
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
