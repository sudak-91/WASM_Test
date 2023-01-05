package htmlelement

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
