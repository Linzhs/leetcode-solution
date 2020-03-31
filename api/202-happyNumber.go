package api

type IntSet map[int]struct{}

func (p IntSet) Add(key int) bool {
	if _, ok := p[key]; ok {
		return false
	}

	p[key] = struct{}{}
	return true
}

func (p IntSet) Get(key int) bool {
	_, ok := p[key]
	return ok
}

func isHappy(n int) bool {
	set := make(IntSet)

	for set.Add(n) {
		sum := 0
		for n > 0 {
			remained := n % 10
			sum += remained * remained
			n /= 10
		}

		if sum == 1 {
			return true
		}
		n = sum
	}

	return false
}
