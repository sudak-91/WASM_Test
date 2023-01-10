package htmlelement

type Input struct {
	HtmlElement
	HtmlClass
	Type   string
	Id     string
	parent Parent
}

func NewInput(parent Parent, id string, inputType string) *Input {
	var (
		input Input
	)
	//input.elem = GetDocument().Call("createElement", "input")
	input.Type = inputType
	input.Id = id
	input.parent = parent
	parent.SetChild(&input)
	return &input
}

func GetInputValue(inputId string) string {
	elem := GetDocument().Call("getElementById", inputId)
	k := elem.Get("value").String()
	return k
}

func (i *Input) Render() {
	i.CreateElement("input")
	i.set("type", i.Type)
	i.SetId(i.Id)
	i.AddClassSliceToClassList(i.GetJs())
}
