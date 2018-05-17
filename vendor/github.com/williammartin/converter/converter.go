package converter

import "github.com/williammartin/temperature"

type Converter struct{}

func New() *Converter {
	return &Converter{}
}

type Temperature = temperature.Temperature

func (c *Converter) Convert(t *Temperature) {

}
