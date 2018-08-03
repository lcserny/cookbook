package confformat

import "fmt"

const (
	exampleToml = `name="Example1" 
age=99`
	exampleJson = `{"name":"Example2", "age":98}`
	exampleYaml = `name: Example3 
age: 97`
)

// UnMarshallAll takes data in various formats
// and converts them into structs
func UnMarshallAll() error {
	t := TOMLData{}
	j := JSONData{}
	y := YAMLData{}

	if _, err := t.Decode([]byte(exampleToml)); err != nil {
		return err
	}
	fmt.Println("TMOL Unmarshal =", t)

	if err := j.Decode([]byte(exampleJson)); err != nil {
		return err
	}
	fmt.Println("JSON Unmarshal =", j)

	if err := y.Decode([]byte(exampleYaml)); err != nil {
		return err
	}
	fmt.Println("YAML Unmarshal =", y)

	return nil
}
