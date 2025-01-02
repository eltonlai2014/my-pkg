package mylib2

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello " + name
}
