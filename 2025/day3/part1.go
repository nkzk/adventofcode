package main

func Part1(input []string) int {
	var sum int

	for _, line := range input {
		sum += largestJoltage(line)
	}

	return sum
}

func largestJoltage(s string) int {
	largest := 0

	for i, r := range s {
		val := int(r - '0')
		if i+1 < len(s) {
			for _, v := range s[i+1:] {
				val2 := int(v - '0')
				if combine(val, val2) > largest {
					largest = combine(val, val2)
				}
			}
		}
	}

	return largest
}

func combine(a, b int) int {
	return (a * 10) + b
}
