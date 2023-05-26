package validator

type fieldType int

type Structure interface {
	Fields() []*Field
}

const (
	primitiveType fieldType = iota
	structureType
	sliceType
)

type Field struct {
	name           string
	fieldType      fieldType
	validateFunc   ValidateFunc
	structure      Structure
	slice          []Structure
	isInlineStruct bool
}

func NewField(name string, validateFunc ValidateFunc) *Field {
	return &Field{
		name:         name,
		fieldType:    primitiveType,
		validateFunc: validateFunc,
	}
}

func NewStruct(name string, structure Structure) *Field {
	return &Field{
		name:      name,
		fieldType: structureType,
		structure: structure,
	}
}

func NewInlineStruct(structure Structure) *Field {
	return &Field{
		fieldType:      structureType,
		structure:      structure,
		isInlineStruct: true,
	}
}

func NewSlice(name string, slice []Structure) *Field {
	return &Field{
		name:      name,
		fieldType: sliceType,
		slice:     slice,
	}
}

func SliceOfStruct[T Structure](structures []T) (result []Structure) {
	for _, s := range structures {
		result = append(result, s)
	}
	return
}
