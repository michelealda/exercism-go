package perfect

import "errors"

//Classification represents the number class
type Classification int

const (
	//ClassificationDeficient class
	ClassificationDeficient Classification = iota
	//ClassificationPerfect class
	ClassificationPerfect
	//ClassificationAbundant class
	ClassificationAbundant
)

//ErrOnlyPositive validate positive input
var ErrOnlyPositive = errors.New("Only positive numbers")

//Classify gives a number its class
func Classify(x int64) (Classification, error) {
	if x <= 0 {
		return -1, ErrOnlyPositive
	}
	if x == 1 {
		return ClassificationDeficient, nil
	}

	a := int64(1)
	for i := int64(2); i <= x/2; i++ {
		if x%i == 0 {
			a += i
		}
	}

	if a > x {
		return ClassificationAbundant, nil
	}
	if a < x {
		return ClassificationDeficient, nil
	}
	return ClassificationPerfect, nil
}
