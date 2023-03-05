package oak

type node struct {
	method   string
	path     string
	handler  Handle
	children []*node
}

func (n *node) getValue() Handle {
	if n != nil {
		return n.handler
	}
	return nil
}

func (n *node) getNode(path string) *node {
	for n.path != path {
		if len(n.children) == 0 {
			break
		}
		n = n.children[0]
	}

	if n.path != path {
		return nil
	}

	return n
}
