package main

const repValue float64 = 0.03333333

func OneRM(weight float64, reps int) float64 {
	if reps > 1 {
		return ((float64(reps) * repValue) + 1) * weight
	}
	return weight
}

func ReverseOneRM(weight float64, reps int) float64 {
	return weight / (1 + (float64(reps) * repValue))
}

func RPE(weight float64, reps int, chiocciola int) float64 {
	buffer := 10 - chiocciola
	return ReverseOneRM(weight, reps + buffer)
}

