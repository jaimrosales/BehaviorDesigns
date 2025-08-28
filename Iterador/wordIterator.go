package main

type wordIterator struct {
	theWord  *Word
	position int
}

func (wI *wordIterator) First() {
	wI.position = 0
}
func (wI *wordIterator) Next() interface{} {
	if wI.HasMore() {
		var r string
		switch wI.position {
		case 0:
			r = wI.theWord.Word1
		case 1:
			r = wI.theWord.Word2
		case 2:
			r = wI.theWord.Word3
		case 3:
			r = wI.theWord.Word4
		default:
			r = "no more"
		}
		wI.position++
		return r
	}
	return nil
}
func (wI *wordIterator) HasMore() bool {
	return wI.position < 4
}
