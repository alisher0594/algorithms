package recursion


func Factorial(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}

func Sum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	} else {
		return arr[0] + sum(arr[1:])
	}
}

