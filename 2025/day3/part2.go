package main

func Part2(input []string) int {
	var sum int
	for _, line := range input {
		sum += largestJoltage2(line)
	}
	return sum
}

func largestJoltage2(s string) int {
	var result []int

	for i := range s {
		if len(result) >= 12 {
			break
		}

		result = append(result, largest(s[i:i+4]))
		i += 4
	}

	return combine2(result...)
}

func largest(s string) int {
	largest := 0
	for _, r := range s {
		if int(r-'0') > largest {
			largest = int(r - '0')
		}
	}

	return largest
}

func combine2(digits ...int) int {
	var result int
	for _, x := range digits {
		result = result*10 + x
	}

	return result
}
