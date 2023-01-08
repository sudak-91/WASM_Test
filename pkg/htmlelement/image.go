package htmlelement

type Image struct {
	HtmlClass
	HtmlAttribute
	HtmlElement
	src string
}

func NewImage(src string) *Image {
	var (
		img Image
	)

	img.elem = GetDocument().Call("createElement", "img")
	img.src = src
	return &img
}

func (i *Image) Render() {
	i.Set("src", i.src)
	i.AddClassSliceToClassList(i.elem)
}
