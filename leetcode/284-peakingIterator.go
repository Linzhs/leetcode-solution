package leetcode

type Iterator struct {
}

func (p *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return false
}

func (p *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	iterator   *Iterator
	currIntVal int
	hasCurr    bool
}

func ConstructorPeekingIterator(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iterator: iter}
}

func (p *PeekingIterator) hasNext() bool {
	return p.hasCurr || p.iterator.hasNext()
}

func (p *PeekingIterator) next() int {
	if p.hasCurr == false {
		return p.iterator.next()
	}

	p.hasCurr = false
	res := p.currIntVal
	p.currIntVal = 0
	return res
}

func (p *PeekingIterator) peek() int {
	if p.hasCurr {
		return p.currIntVal
	}

	if p.iterator.hasNext() {
		p.hasCurr = true
		p.currIntVal = p.iterator.next()
	}

	return p.currIntVal
}
