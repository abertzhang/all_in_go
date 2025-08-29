package pattern
type IDepartment interface{
	Execute(*Do)
	SetNext( IDepartment)
}