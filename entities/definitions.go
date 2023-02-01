package entities

type QualityIndex interface {
	GetValue() int
}

type CyclomaticComplexity int

func (c CyclomaticComplexity) GetValue() int {
	return int(c)
}

type MaintainabilityIndex int

func (m MaintainabilityIndex) GetValue() int {
	return int(m)
}
