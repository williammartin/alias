package converter

import "github.com/williammartin/temperature"

type Converter struct{}

func New() *Converter {
	return &Converter{}
}

func (c *Converter) Convert(t *temperature.Temperature) {

}
