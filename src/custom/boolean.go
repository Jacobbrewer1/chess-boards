package custom

import "fmt"

type Bool bool

func (b *Bool) Scan(src any) error {
	*b = fmt.Sprintf("%s", src) == "\x01"
	return nil
}

// Equals is primarily used for passing a default bool variable in. By this is meant of a variable storing a
// bool value as you can compare directly to a true or false value
func (b *Bool) Equals(boolean bool) bool {
	return *b == Bool(boolean)
}

// Boolean returns the built-in bool value of the Bool type. This is only advised to be used within if statements
// where required
func (b *Bool) Boolean() bool {
	return b.Equals(true)
}
