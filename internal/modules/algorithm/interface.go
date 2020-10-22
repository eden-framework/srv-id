package algorithm

import "github.com/eden-framework/srv-id/internal/constants/enums"

type GeneratorAlgorithm interface {
	GetAlgorithmID() enums.GenerateAlgorithm
	InitGenerator() error
	GenerateUniqueID() (uint64, error)
}

type GeneratorContainer map[enums.GenerateAlgorithm]GeneratorAlgorithm

func (a GeneratorContainer) RegisterAlgorithm(v GeneratorAlgorithm) {
	v.InitGenerator()
	a[v.GetAlgorithmID()] = v
}

func (a GeneratorContainer) GetAlgorithm(algorithm enums.GenerateAlgorithm) GeneratorAlgorithm {
	if v, ok := a[algorithm]; ok {
		return v
	}
	return nil
}

var GeneratorContainerInstance GeneratorContainer

func init() {
	GeneratorContainerInstance = GeneratorContainer{}
}
