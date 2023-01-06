package htmlelement

import "fmt"

type Body struct {
	HtmlElement
}

func GetBody() *Body {
	var (
		body Body
	)
	body.elem = QuerySelector("body")
	return &body
}

func (b *Body) Render() {
	fmt.Println("Start Body Render")
	for _, v := range b.ChildElement {
		fmt.Printf("Element is %v\n", v)
		render, ok := v.(Renderer)
		fmt.Printf("Renderer status - %t\n", ok)
		if ok {
			render.Render()
		}
		elem, ok := v.(JsGeter)
		fmt.Printf("jsGeter status - %t\n", ok)
		if ok {
			b.AppendChild(elem.GetJs())
		}
	}
}
