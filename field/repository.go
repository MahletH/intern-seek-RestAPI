package field

import "github.com/abdimussa87/intern-seek-RestAPI/entity"

type FieldRepository interface {
	Field(id uint) (*entity.Field, []error)
	UpdateField(role *entity.Field) (*entity.Field, []error)
	Fields() ([]entity.Field, []error)
	DeleteField(id uint) (*entity.Field, []error)
	StoreField(field *entity.Field) (*entity.Field, []error)
	GetFieldbyName(name string) (*entity.Field, []error)

	FieldInternships(field *entity.Field) ([]entity.Internship, []error)
	FieldInterns(field *entity.Field) ([]entity.Intern, []error)
}
