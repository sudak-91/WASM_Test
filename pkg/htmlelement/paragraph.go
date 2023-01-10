package htmlelement

type Paragraph struct {
	HtmlElement
	HtmlClass
	Text   string
	Id     string
	parent Parent
}

func NewParagaph(parent Parent, id string, text string) *Paragraph {
	var (
		p Paragraph
	)
	p.Text = text
	p.Id = id
	parent.SetChild(&p)
	//p.elem = GetDocument().Call("createElement", "p")
	return &p
}

func (p *Paragraph) Render() {
	p.CreateElement("p")
	p.SetId(p.Id)
	p.AddClassSliceToClassList(p.elem)
	p.SetInnerHtml(p.Text)
}

func (p *Paragraph) ChangeText(text string) {
	p.Text = text
	p.SetInnerHtml(text)
}
