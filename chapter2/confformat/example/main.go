package main

import "cookbook/chapter2/confformat"

func main() {
	if err := confformat.MarshallAll(); err != nil {
		panic(err)
	}

	if err := confformat.UnMarshallAll(); err != nil {
		panic(err)
	}

	if err := confformat.OtherJSONExamples(); err != nil {
		panic(err)
	}
}
