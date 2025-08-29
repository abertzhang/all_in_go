package pattern

import "fmt"

var colorFactory *colorFlyweightFactory
//单例
type colorFlyweightFactory struct{
	maps map[string]*ColorFlyweight
}
func (f *colorFlyweightFactory) Get(fileName string) *ColorFlyweight {
	color := f.maps[fileName]
	if color == nil {
		color = newColorFlyweight(fileName)
		f.maps[fileName] = color
	}
	return color
}
//获取单例
func getColorFlyweightFactory() *colorFlyweightFactory {
	if colorFactory == nil {
		colorFactory = &colorFlyweightFactory{
			maps: make(map[string]*ColorFlyweight),
		}
	}
	return colorFactory
}

//颜色内容,KV中的value值
type ColorFlyweight struct{
	Data string
}
func newColorFlyweight(fileName string) *ColorFlyweight {
	data:=fmt.Sprintf("color data %s",fileName)
	return &ColorFlyweight{Data: data}
}

type colorViewer struct{
	*ColorFlyweight
}

func NewColorViewer(name string) *colorViewer {
	color:=getColorFlyweightFactory().Get(name)
	return &colorViewer{ColorFlyweight: color}
}