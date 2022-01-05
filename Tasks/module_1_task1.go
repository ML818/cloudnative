package main

func Chang2And4(arr []string, second string, forth string) []string {
	for index, v := range arr {
		if index == 2 && v == "stupid" {
			arr[index] = second
		}
		if index == 4 && v == "weak" {
			arr[index] = forth
		}
	}

	return arr
}

func main() {
	str := [5]string{"I", "am", "stupid", "and", "weak"}

	result := Chang2And4(str[:], "smart", "strong")

	for _, v := range result {
		print(v)
		print(" ")
	}
}
