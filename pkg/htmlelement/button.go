package htmlelement

import (
	"syscall/js"
)

type Button struct {
	HtmlElement
	HtmlClass
	HtmlAttribute
	Type        string
	Id          string
	Text        string
	OnClickFunc *js.Func
}

func NewButton(parent Parent, id string, btnType string, text string) *Button {
	var (
		btn Button
	)
	btn.Id = id
	btn.Type = btnType
	btn.Text = text
	parent.SetChild(&btn)
	return &btn
}

func (b *Button) AddClickEventListener(jFunc *js.Func) {
	b.OnClickFunc = jFunc
	//b.elem.Call("addEventListener", "click", *jFunc)

}

func (b *Button) Render() {
	b.CreateElement("button")
	b.set("type", b.Type)
	b.SetInnerHtml(b.Text)
	b.SetId(b.Id)
	b.AddClassSliceToClassList(b.elem)
	if b.OnClickFunc != nil {
		b.elem.Call("addEventListener", "click", *b.OnClickFunc)
	}
}
