package attr

// Attr interface describes the behavior that a serializable attribute
// should have.
//
// Besides retrieving its key and value, it also permits creating a copy of
// the original Attr with a different key or a different value
type Attr interface {
	// Key returns the string key of the attribute Attr
	Key() string
	// Value returns the (any) value of the attribute Attr
	Value() any
	// WithKey returns a copy of this Attr, with key `key`
	WithKey(key string) Attr
	// WithValue returns a copy of this Attr, with value `value`
	//
	// It must be the same type of the original Attr, otherwise returns
	// nil
	WithValue(value any) Attr
}

// New is a generic function to create an Attr
//
// Using a generic approach allows the Attr.WithValue method to be
// scoped with certain constraints for specific applications
func New[T any](key string, value T) Attr {
	if key == "" {
		return nil
	}
	return attr[T]{
		key:   key,
		value: value,
	}
}

type attr[T any] struct {
	key   string
	value T
}

// Key returns the string key of the attribute Attr
func (a attr[T]) Key() string {
	return a.key
}

// Value returns the (any) value of the attribute Attr
func (a attr[T]) Value() any {
	return a.value
}

// WithKey returns a copy of this Attr, with key `key`
func (a attr[T]) WithKey(key string) Attr {
	if key == "" {
		return nil
	}
	return New(key, a.value)
}

// WithValue returns a copy of this Attr, with value `value`
//
// It must be the same type of the original Attr, otherwise returns
// nil
func (a attr[T]) WithValue(value any) Attr {
	if value == nil {
		return nil
	}

	v, ok := (value).(T)
	if !ok {
		return nil
	}
	return New(a.key, v)
}

// NewPtr is a generic function to create an Attr from a pointer value
//
// Using a generic approach allows the Attr.WithValue method to be
// scoped with certain constraints for specific applications
func NewPtr[T any](key string, value *T) Attr {
	if key == "" {
		return nil
	}
	return &ptrAttr[T]{
		key: key,
		ptr: value,
	}
}

type ptrAttr[T any] struct {
	key string
	ptr *T
}

// Key returns the string key of the attribute Attr
func (p *ptrAttr[T]) Key() string {
	return p.key
}

// Value returns the (any) value of the attribute Attr
func (p *ptrAttr[T]) Value() any {
	if p.ptr == nil {
		return nil
	}
	return *p.ptr
}

// WithKey returns a copy of this Attr, with key `key`
func (p *ptrAttr[T]) WithKey(key string) Attr {
	return NewPtr(key, p.ptr)
}

// WithValue returns a copy of this Attr, with value `value`
//
// It must be the same type of the original Attr, otherwise returns
// nil
func (p *ptrAttr[T]) WithValue(value any) Attr {
	if value == nil {
		return nil
	}

	v, ok := (value).(*T)
	if !ok {
		return nil
	}
	return NewPtr(p.key, v)
}
