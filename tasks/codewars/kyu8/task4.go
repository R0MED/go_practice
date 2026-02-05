package kyu8

func combat(health, damage float64) float64 {
	if damage >= health {
		return 0
	}
	return health - damage
}

//https://www.codewars.com/kata/586c1cf4b98de0399300001d/train/go
