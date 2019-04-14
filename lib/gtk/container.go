package gtk

import sugar "github.com/jopbrown/gtk-sugar"

type Container struct {
	Widget
}

func NewContainerFromCandy(candy sugar.Candy, id string) *Container {
	v := Container{}
	v.CandyWrapper = candy.NewWrapper(id)
	return &v
}

// FUNCTION_NAME = gtk_container_add, NONE, NONE, 2, WIDGET, WIDGET
func (c *Container) Add(widget IWidget) {
	c.Candy().Guify("gtk_container_add", c, widget)
}

// FUNCTION_NAME = gtk_container_remove, NONE, NONE, 2, WIDGET, WIDGET
func (c *Container) Remove(widget IWidget) {
	c.Candy().Guify("gtk_container_remove", c, widget)
}

// FUNCTION_NAME = gtk_container_set_border_width, NONE, NONE, 2, WIDGET, INT
func (c *Container) SetBorderWidth(width int) {
	c.Candy().Guify("gtk_container_set_border_width", c, width)
}
