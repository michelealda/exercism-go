package allergies

var allergenes = [8]string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

//Allergies calculate the list of allergies given a score
func Allergies(score uint) (result []string) {
	for i := range allergenes {
		if score&(1<<uint(i)) > 0 {
			result = append(result, allergenes[i])
		}
	}
	return
}

//AllergicTo calculate if your score match an allergen
func AllergicTo(score uint, substance string) bool {
	for _, a := range Allergies(score) {
		if a == substance {
			return true
		}
	}
	return false
}
