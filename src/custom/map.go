package custom

type Map[T, V comparable] map[T]V

func (c *Map[T, V]) Set(key T, value V) {
	if (*c) == nil {
		*c = make(Map[T, V])
	}

	(*c)[key] = value
}

func (c *Map[T, V]) Get(key T) *V {
	if c == nil || !c.Has(key) {
		return nil
	}

	foundValue := (*c)[key]
	return &foundValue
}

func (c *Map[T, V]) Has(key T) bool {
	if c == nil {
		return false
	}

	_, ok := (*c)[key]
	return ok
}

func (c *Map[T, V]) Remove(key T) {
	if (*c) == nil {
		return
	}

	delete(*c, key)
}

func (c *Map[T, V]) Values() []V {
	var values []V
	for _, value := range *c {
		values = append(values, value)
	}

	return values
}

func (c *Map[T, V]) Keys() []T {
	var keys []T
	for key, _ := range *c {
		keys = append(keys, key)
	}

	return keys
}
