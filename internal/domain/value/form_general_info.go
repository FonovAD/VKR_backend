package value

import "errors"

type FormGeneralInfo struct {
	INN         string
	Name        string
	ExistMuseum bool
}

func NewFormGeneralInfo(inn, name string, exist bool) (FormGeneralInfo, error) {
	if inn == "" {
		return FormGeneralInfo{}, errors.New("form general info is empty")
	}
	if name == "" {
		return FormGeneralInfo{}, errors.New("form general info is empty")
	}
	return FormGeneralInfo{INN: inn, Name: name, ExistMuseum: exist}, nil
}
