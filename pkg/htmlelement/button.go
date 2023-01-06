package htmlelement

import "syscall/js"

type Button struct {
	HtmlElement
	HtmlClass
	Type string
}

func NewButton() *Button {
	var (
		btn Button
	)
	btn.elem = GetDocument().Call("createElement", "button")
	return &btn
}

func (b *Button) AddClickEventListener(jFunc *js.Func, args ...string) {
	b.elem.Call("addEventListener", "click", *jFunc)
}

func (b *Button) AddClickGoFunc(funcJ any) {
	b.elem.Call("addEventListener", "click", funcJ)
}
func (b *Button) AddType(btnType string) {
	b.Type = btnType
}

func (b *Button) Render() {
	b.AddClassSliceToClassList(b.elem)

}
