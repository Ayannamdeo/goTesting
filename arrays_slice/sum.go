package arraysslice

func Sum(numbers []int) int {
	res := 0
	for _, n := range numbers {
		res += n
	}
	return res
}

func SumAll(numSlices ...[]int) []int {
	var res []int
	for _, numbers := range numSlices {
		res = append(res, Sum(numbers))
	}
	return res
}

func SumAllTails(numSlices ...[]int) []int {
	// res := make([]int, len(numSlices))
	var res []int
	for _, numslice := range numSlices {
		var tail []int
		if len(numslice) == 0 {
			tail = []int{0}
		} else {
			tail = numslice[1:]
		}
		res = append(res, Sum(tail))
	}
	return res
}

// func TestSumAllTails(t *testing.T) {
// 	got := SumAllTails([]int{1, 2}, []int{0, 9})
// 	want := []int{2, 9}
//
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }
