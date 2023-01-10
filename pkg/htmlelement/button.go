package htmlelement

import "syscall/js"

type Button struct {
	HtmlElement
	HtmlClass
	HtmlAttribute
	Type        string
	Id          string
	OnClickFunc *js.Func
}

func NewButton(parent Parent, id string, btnType string) *Button {
	var (
		btn Button
	)
	btn.Id = id
	btn.Type = btnType
	parent.SetChild(&btn)
	//btn.elem = GetDocument().Call("createElement", "button")
	return &btn
}

func (b *Button) AddClickEventListener(jFunc *js.Func) {
	b.OnClickFunc = jFunc
	//b.elem.Call("addEventListener", "click", *jFunc)

}

func (b *Button) Render() {
	b.CreateElement("button")
	b.set("type", b.Type)
	b.SetId(b.Id)
	b.AddClassSliceToClassList(b.elem)

}
