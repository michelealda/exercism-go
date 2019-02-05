package protein

import "errors"

//ErrStop is the stop sequence error
var ErrStop = errors.New("stop")

//ErrInvalidBase is returned when a base is invalid
var ErrInvalidBase = errors.New("invalid base")

//FromCodon returns the amynoacid for the given codon
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

func fromRnaRec(sequence string, state []string) ([]string, error) {
	if len(sequence) < 3 {
		return state, nil
	}
	amynoacid, e := FromCodon(sequence[:3])
	if e == ErrStop {
		return state, nil
	} else if e == ErrInvalidBase {
		return state, e
	} else {
		return fromRnaRec(sequence[3:],
			append(state, amynoacid))
	}
}

//FromRNA return the list of amynoacids for a given sequence
func FromRNA(input string) ([]string, error) {
	return fromRnaRec(input, []string{})
}
