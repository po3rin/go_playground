package example

func helloWorld() {
	println("Hello, World")
}

func helloGopher() {
	println("Hello, Gopher")

	{
		helloInnerWorld := func() {
			println("Hello, Inner World")
		}
		helloInnerWorld()
	}
}
