package set

type AddCommand[T comparable] struct {
	set *Set[T]
	x   T
}

func (c *AddCommand[T]) Undo() {
	c.set.Delete(c.x)
}

type DeleteCommand[T comparable] struct {
	set *Set[T]
	x   T
}

func (c *DeleteCommand[T]) Undo() {
	c.set.Add(c.x)
}
