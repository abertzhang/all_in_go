package pattern
type ICentimeter interface {
	GetLength(float64) float64
}

type IMeter interface {
	GetLength(float64) float64
}

type ILengthAdapter interface {
	GetLength(string,float64) float64
}