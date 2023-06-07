package validator

type fieldType int

type Structure interface {
	Fields() Fields
}

type Fields []*Field
type Field struct {
	name           string
	fieldType      fieldType
	validateFunc   ValidateFunc
	structure      Structure
	slice          []Structure
	isJoinedStruct bool
}

const (
	primitiveType fieldType = iota
	structureType
	sliceType
)

func NewFields(fields ...*Field) Fields {
	return fields
}

func (fs Fields) Append(field *Field) Fields {
	fs = append(fs, field)
	return fs
}

func (fs Fields) Join(structure Structure) Fields {
	fs = append(fs, NewJoinedStruct(structure))
	return fs
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

func NewJoinedStruct(structure Structure) *Field {
	return &Field{
		fieldType:      structureType,
		structure:      structure,
		isJoinedStruct: true,
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
