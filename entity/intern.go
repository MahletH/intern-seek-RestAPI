package entity

type Intern struct {
	// ID            int
	// UserName      string
	// FullName      string
	// Email         string
	// Phone         string
	// password      string
	// Field         string
	// AcademicLevel string

	InternDetail   PersonalDetails
	InternUser     User
	AvailableField []Field
}
