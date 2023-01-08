package htmlelement

type Paragraph struct {
	HtmlElement
	HtmlClass
	Text string
}

func NewParagaph(text string) *Paragraph {
	var (
		p Paragraph
	)
	p.Text = text
	p.elem = GetDocument().Call("createElement", "p")
	return &p
}

func (p *Paragraph) Render() {
	p.AddClassSliceToClassList(p.elem)
	p.Set("innerHTML", p.Text)
}
