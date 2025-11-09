package hax

func If(condition bool, child INode) INode {
	var element INode = nil
	if condition {
		element = child
	}
	return element
}

func Show(condition bool, ifPath, elsePath INode) INode {
	var element INode = nil
	if condition {
		element = ifPath
	} else {
		element = elsePath
	}
	return element
}

/*
func Each[T any](src Gettable[[]T], renderOne func(index int, value T) INode) INode {
}

func EachMap[K comparable, T any](src Gettable[map[K]T], renderOne func(index K, value T) INode) INode {

}
*/
