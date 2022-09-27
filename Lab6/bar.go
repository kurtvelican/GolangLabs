package main

func bar(h human) {
	switch h.(type) {
	case person:
		h.(person).speak()
	case agent:
		h.(agent).speak()
	}
}
