package domain

type FieldCollection struct {
	id         int64
	collection Collection
	field      IField
	unique     bool
	editable   bool
	required   bool
	sequence   Sequence
	list       List
}

func NewFieldCollection(collection Collection, field IField) *FieldCollection {
	return &FieldCollection{
		collection: collection,
		field:      field,
	}
}

func (f *FieldCollection) WithSequence(sequence Sequence) *FieldCollection {
	f.sequence = sequence
	return f
}

func (f *FieldCollection) WithList(list List) *FieldCollection {
	f.list = list
	return f
}

func (f *FieldCollection) WithCollection(collection Collection) *FieldCollection {
	f.collection = collection
	return f
}

func (f *FieldCollection) WithField(field IField) *FieldCollection {
	f.field = field
	return f
}

func (f *FieldCollection) WithId(id int64) *FieldCollection {
	f.id = id
	return f
}

func (f *FieldCollection) WithUnique(unique bool) *FieldCollection {
	f.unique = unique
	return f
}

func (f *FieldCollection) WithEditable(editable bool) *FieldCollection {
	f.editable = editable
	return f
}

func (f *FieldCollection) WithRequired(required bool) *FieldCollection {
	f.required = required
	return f
}

func (f *FieldCollection) Id() int64 {
	return f.id
}

func (f *FieldCollection) Unique() bool {
	return f.unique
}

func (f *FieldCollection) Editable() bool {
	return f.editable
}

func (f *FieldCollection) Required() bool {
	return f.required
}

func (f *FieldCollection) Collection() *Collection {
	return &f.collection
}

func (f *FieldCollection) Field() IField {
	return f.field
}

func (f *FieldCollection) List() *List {
	return &f.list
}

func (f *FieldCollection) IdCollection() int64 {
	return f.Collection().id
}

func (f *FieldCollection) IdField() int64 {
	return f.Field().Id()
}

func (f *FieldCollection) Exists() bool {
	return f.IdCollection() > 0 && f.IdField() > 0
}

func (f *FieldCollection) Sequence() *Sequence {
	return &f.sequence
}
