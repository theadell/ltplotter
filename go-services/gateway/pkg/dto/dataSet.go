package dto

type ComplexNumber struct {
	Real float64 `json:"real"`
	Imag float64 `json:"imag"`
}

type DataSetType string

const (
	TypeFloat64 DataSetType = "float64"
	TypeComplex DataSetType = "complex"
)

type DataSet[T float64 | ComplexNumber] struct {
	Name      string                    `json:"name"`
	SimType   string                    `json:"simType"`
	Steps     int                       `json:"steps"`
	Variables []string                  `json:"variables"`
	XAxis     string                    `json:"xAxis"`
	Data      map[string]map[string][]T `json:"data"`
	Type      DataSetType               `json:"type"`
}
