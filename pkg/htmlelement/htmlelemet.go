package htmlelement

import "syscall/js"

func GetDocument() js.Value {
	return js.Global().Get("document")
}

func QuerySelector(element string) js.Value {
	return GetDocument().Call("querySelector", element)
}

//Renderer it's interface to Render html page
type Renderer interface {
	Render()
}

//JsGeter - interface to get js.Value from struct
type JsGeter interface {
	GetJs() js.Value
}

//Parent - interface to eject depend
type Parent interface {
	SetChild(child any)
}

//HtmlElement struct of default html element
type HtmlElement struct {
	elem         js.Value //Instanse of HTMLElement(div, button, etc..)
	ChildElement []interface{}
}

//AppendChild add element to DOM model
func (h *HtmlElement) AppendChild(child js.Value) {
	h.elem.Call("appendChild", child)
}

func (h *HtmlElement) set(setType string, value string) {
	h.elem.Set(setType, value)
}
func (h *HtmlElement) CreateElement(element string) {
	h.elem = GetDocument().Call("createElement", element)
}

func (h *HtmlElement) SetInnerHtml(value string) {
	h.set("innerHTML", value)
}
func (h *HtmlElement) SetId(id string) {
	h.set("id", id)
}

//GetJs return java script object
func (h *HtmlElement) GetJs() js.Value {
	return h.elem
}
func (h *HtmlElement) SetChild(child any) {
	h.ChildElement = append(h.ChildElement, child)
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
