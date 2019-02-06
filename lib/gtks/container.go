package gtks

type Container struct {
	Widget
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
