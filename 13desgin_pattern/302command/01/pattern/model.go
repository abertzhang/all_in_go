package pattern
import "fmt"

type Button struct {
	Command ICommand
}
func (b *Button) Press()  {
	b.Command.Execute()
}
//实现ICommond接口
type OnCommand struct {
	Device IDevice
}
func (o *OnCommand) Execute() {
	o.Device.On()
}
//实现ICommond接口
type OffCommand struct {
	Device IDevice
}
func (o *OffCommand) Execute() {
	o.Device.Off()
}
//举例TV
//实现IDevice接口
type TV struct {
}

func (t *TV) On() {
	fmt.Println("TV is ON")
}

func (t *TV) Off() {
	fmt.Println("TV is OFF")
}
