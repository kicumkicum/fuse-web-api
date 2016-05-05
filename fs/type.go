package fs

type IContent interface {
	Size() int
}

type Content struct {
	size int
	Id string
	Name string
}

func (c *Content) Size() int {
	return c.size
}

type File struct {
	Content
}

type Directory struct {
	Content
	List []IContent
}


func (d *Directory) Size() int {
	result := 0

	for _, item := range d.List {
		result += item.Size()
	}

	return result
}

