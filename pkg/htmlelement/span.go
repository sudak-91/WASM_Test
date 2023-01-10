package htmlelement

type Span struct {
	HtmlElement
	HtmlClass
	Text string
	Id   string
}

func NewSpan(parent Parent, id string, text string) *Span {
	var (
		span Span
	)
	//span.elem = GetDocument().Call("createElement", "span")
	span.Text = text
	span.Id = id
	parent.SetChild(&span)
	return &span
}

func (s *Span) Render() {
	s.CreateElement("span")
	s.SetId(s.Id)
	s.SetInnerHtml(s.Text)
	s.AddClassSliceToClassList(s.GetJs())
}
