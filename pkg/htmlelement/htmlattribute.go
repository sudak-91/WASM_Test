package htmlelement

type HtmlAttribute struct {
}

func (h *HtmlAttribute) setAttribute(object JsGeter, key string, value string) {
	object.GetJs().Call("setAttribute", key, value)

}

func (h *HtmlAttribute) SetOnClick(object JsGeter, jsFunction string) {
	h.setAttribute(object, "onclick", jsFunction)
}
