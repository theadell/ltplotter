package mapper

import (
	"log/slog"
	"ltplotter/gateway/pkg/dto"
	"strconv"

	"github.com/theadell/ltspice"
)

func ToDataSet(data *ltspice.SimData) any {
	vars := make([]string, 0, data.Meta.NoVariables)
	for _, v := range data.GetVariables() {
		vars = append(vars, v.Name)
	}

	if data.Meta.SimType == ltspice.ACAnalysis {

		d := make(map[string]map[string][]dto.ComplexNumber)

		for i := range data.GetSteps() {
			stepKey := strconv.Itoa(i)
			d[stepKey] = make(map[string][]dto.ComplexNumber)

			for _, v := range data.GetVariables() {
				t, err := ltspice.GetTrace[complex128](data, v.Name)
				if err != nil {
					slog.Error("failed to get trace", "trace", v)
					continue
				}
				s := t.GetSignal(i)
				d[stepKey][v.Name] = ConvertComplex128Slice(s)

			}
		}

		return dto.DataSet[dto.ComplexNumber]{
			Name:      data.Meta.Title,
			SimType:   data.GetType().String(),
			Steps:     data.GetSteps(),
			Variables: vars,
			XAxis:     vars[0],
			Data:      d,
			Type:      dto.TypeComplex,
		}
	} else {

		d := make(map[string]map[string][]float64)

		for i := range data.GetSteps() {
			stepKey := strconv.Itoa(i)
			d[stepKey] = make(map[string][]float64)

			for _, v := range data.GetVariables() {
				t, err := ltspice.GetTrace[float64](data, v.Name)
				if err != nil {
					slog.Error("failed to get trace", "trace", v)
					continue
				}
				s := t.GetSignal(i)
				d[stepKey][v.Name] = s

			}
		}
		return dto.DataSet[float64]{
			Name:      data.Meta.Title,
			SimType:   data.GetType().String(),
			Steps:     data.GetSteps(),
			Variables: vars,
			XAxis:     vars[0],
			Data:      d,
			Type:      dto.TypeFloat64,
		}
	}
}

func ConvertComplex128Slice(complexSlice []complex128) []dto.ComplexNumber {
	complexNumberSlice := make([]dto.ComplexNumber, len(complexSlice))
	for i, c := range complexSlice {
		complexNumberSlice[i] = dto.ComplexNumber{
			Real: real(c),
			Imag: imag(c),
		}
	}
	return complexNumberSlice
}
