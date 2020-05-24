package leetcode

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	area1 := (C - A) * (D - B)
	area2 := (G - E) * (H - F)

	left := max(A, E)
	right := min(C, G)
	top := min(D, H)
	bottom := max(F, B)

	overlapArea := 0
	if right > left && top > bottom {
		overlapArea = (right - left) * (top - bottom)
	}

	return area1 + area2 - overlapArea
}
