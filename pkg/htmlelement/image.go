package htmlelement

type Image struct {
	HtmlClass
	HtmlAttribute
	HtmlElement
	src string
	Id  string
}

func NewImage(parent Parent, src string, id string) *Image {
	var (
		img Image
	)

	//img.elem = GetDocument().Call("createElement", "img")
	img.src = src
	img.Id = id
	parent.SetChild(&img)
	return &img
}

func (i *Image) Render() {
	i.CreateElement("img")
	i.set("src", i.src)
	i.SetId(i.Id)
	i.AddClassSliceToClassList(i.elem)
}
