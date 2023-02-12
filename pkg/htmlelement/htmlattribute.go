package htmlelement

type HtmlAttribute struct {
}

func (h *HtmlAttribute) setAttribute(object JsGeter, key string, value string) {
	object.GetJs().Call("setAttribute", key, value)

}

func (h *HtmlAttribute) SettaBsToggle(object JsGeter) {
	h.setAttribute(object, "data-bs-toggle", "collapse")
}

func (h *HtmlAttribute) SetDataBsTarger(object JsGeter, targetID string) {
	h.setAttribute(object, "data-bs-target", targetID)
}

func (h *HtmlAttribute) SetOnClick(object JsGeter, jsFunction string) {
	h.setAttribute(object, "onclick", jsFunction)
}
