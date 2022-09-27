package Interface

func Bar(h Human) {
	switch h.(type) {
	case Person:
		h.(Person).Speak()
	case Agent:
		h.(Agent).Speak()
	}
}
