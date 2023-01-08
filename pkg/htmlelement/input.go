package htmlelement

type Input struct {
	HtmlElement
	HtmlClass
	Type string
}

func NewInput(inputType string) *Input {
	var (
		input Input
	)
	input.elem = GetDocument().Call("createElement", "input")
	input.Type = inputType
	return &input
}

func GetInputValue(inputId string) string {
	elem := GetDocument().Call("getElementById", inputId)
	k := elem.Get("value").String()
	return k
}

func (i *Input) Render() {
	i.AddClassSliceToClassList(i.GetJs())
	i.Set("type", i.Type)
}
