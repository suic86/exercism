package lasagna

const defaultPrepTime int = 2
const defualtSauceAmount float64 = 0.2
const defaultNoodlesAmount int = 50

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = defaultPrepTime
	}
	return len(layers) * time
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "sauce":
			sauce += defualtSauceAmount
		case "noodles":
			noodles += defaultNoodlesAmount
		}
	}
	return
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, scale int) []float64 {
	scaled := make([]float64, len(quantities))
	fscale := float64(scale) / 2.0
	for i := 0; i < len(quantities); i++ {
		scaled[i] = quantities[i] * fscale
	}
	return scaled
}
