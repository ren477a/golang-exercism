package hamming

import "errors"

// Distance ...
func Distance(a, b string) (int, error) {
	c := 0
	la := len(a)
	lb := len(b)

	if la != lb {
		return 0, errors.New("Unequal string length")
	}

	for i := 0; i < la; i++ {
		if a[i] != b[i] {
			c++
		}
	}
	return c, nil
}

// func main() {
// 	x, _ := Distance("zxc", "zsa")
// 	fmt.Println(x)
// }
