package pattern

import "errors"

type LengthMeter struct{}
func NewLengthMeter() *LengthMeter {
	return &LengthMeter{}
}

func (l *LengthMeter) GetLength(cm float64) float64 {
	return cm / 100
}

type LengthCentimeter struct{}

func NewLengthCentimeter() *LengthCentimeter {
	return &LengthCentimeter{}
}

func (l *LengthCentimeter) GetLength(meters float64) float64 {
	return meters * 100
}

type LengthAdapter struct {
}

func NewLengthAdapter() *LengthAdapter {
	return &LengthAdapter{}
}
func (l *LengthAdapter) GetLength(unit string, length float64) (float64,error) {
	if unit == "m" {
		lengthMeter := NewLengthMeter()
		return lengthMeter.GetLength(length),nil
	} else if unit == "cm" {
		lengthCentimeter := NewLengthCentimeter()
		return lengthCentimeter.GetLength(length),nil
	}
	return 0,errors.New("invalid unit")

}