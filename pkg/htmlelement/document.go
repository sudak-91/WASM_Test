package htmlelement

import "syscall/js"

type Document struct {
	elem js.Value
}

func (d *Document) GetJS() js.Value {
	return js.Global().Get("document")
}

func (d Document) CreateDiv() Div {
	var (
		div Div
	)
	div.SetJsElement(d.GetJS().Call("createElement", "div"))
	return div

}

func (d Document) CreateImg() js.Value {
	return d.GetJS().Call("createElement", "img")
}
