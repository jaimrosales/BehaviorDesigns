package main

type numberIterator struct {
	theNumber *Number
	position  int
}

func (nI *numberIterator) First() {
	nI.position = 0
}

func (nI *numberIterator) Next() interface{} {
	if nI.HasMore() {
		nI.position++
		return nI.theNumber.Numbers[nI.	position-1]
	}
	return nil

}

func (nI *numberIterator) HasMore() bool {
	return nI.position < len(nI.theNumber.Numbers)
}
