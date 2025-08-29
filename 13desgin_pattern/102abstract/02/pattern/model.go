package pattern
type ShopFactory struct {}
func (f *ShopFactory) CreatePen()IPen{
	return &BluePen{}
}
func (f *ShopFactory) CreateBook()IBook{
	return &ComputerBook{}
}
//各类笔
type BluePen struct {}
func (p *BluePen) GetPen() string {
	return "蓝色笔"
}

type RedPen struct {}
func (p *RedPen) GetPen() string {
	return "红色笔"
}
//各类书籍
type ComputerBook struct {}
func (b *ComputerBook) GetBook() string {
	return "计算机书籍"
}
type LiteratureBook struct {}
func (b *LiteratureBook) GetBook() string {
	return "文学书籍"
}
