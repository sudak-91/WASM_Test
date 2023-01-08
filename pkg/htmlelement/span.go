package htmlelement

type Span struct {
	HtmlElement
	HtmlClass
	Text string
}

func NewSpan(Text string) *Span {
	var (
		span Span
	)
	span.elem = GetDocument().Call("createElement", "span")
	span.Text = Text
	return &span
}

func (s *Span) Render() {
	s.AddClassSliceToClassList(s.GetJs())
	s.Set("innerHTML", s.Text)
}
