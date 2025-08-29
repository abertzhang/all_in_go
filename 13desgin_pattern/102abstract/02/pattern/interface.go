package pattern
type IPen interface{
	GetPen()string
}
//
type IBook interface{
	GetBook()string
}
//
type IFactory interface{
	CreatePen()IPen
	CreateBook()IBook
}
