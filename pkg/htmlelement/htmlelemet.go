package htmlelement

import "syscall/js"

const Document = js.Global().Get("document")

func QuerySelector(element string) js.Value {
	return Document.Call("querySelector", element)
}

//Renderer it's interface to Render html page
type Renderer interface {
	Render()
}

//JsGeter - interface to get js.Value from struct
type JsGeter interface {
	GetJs() js.Value
}

//HtmlElement struct of default html element
type HtmlElement struct {
	elem         js.Value //Instanse of HTMLElement(div, button, etc..)
	ChildElement []Renderer
}

//AppendChild add element to DOM model
func (h *HtmlElement) AppendChild(child js.Value) {
	h.elem.Call("appendChild", child)
}

//AddChild add child element to ChildElement slice
func (h *HtmlElement) AddChild(child Renderer) {
	h.ChildElement = append(h.ChildElement, child)
}

//GetJs return java script object
func (h *HtmlElement) GetJs() js.Value {
	return h.elem
}

type HtmlClass struct {
	Class []string
}

//AddClass adding class name to string slace of element
func (h *HtmlClass) AddClass(className string) {
	h.Class = append(h.Class, className)
}

//AddClassList adding class to HtmlElement
func (h *HtmlClass) AddClassList(jsElement js.Value, className string) {
	jsElement.Get("classList").Call("add", className)
}

//AddClassSliceToClassList adding strings clise of class to html element
func (h *HtmlClass) AddClassSliceToClassList(jsElement js.Value) {
	for _, v := range h.Class {
		h.AddClassList(jsElement, v)
	}
}
