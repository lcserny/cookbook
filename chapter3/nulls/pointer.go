package nulls

import (
	"encoding/json"
	"fmt"
)

// ExamplePointer is the same but uses an *Int
type ExamplePointer struct {
	Age *int `json:"age,omitempty"`
	Name string `json:"name"`
}

// PointerEncoding shows methods
// for dealing with nil / omitted values
func PointerEncoding() error {
	// note that no age = nil age
	pointer := ExamplePointer{}

	if err := json.Unmarshal([]byte(jsonBlob), &pointer); err != nil {
		return err
	}
	fmt.Printf("Pointer unmarshal, with no age: %+v\n", pointer)

	value, err := json.Marshal(&pointer)
	if err != nil {
		return err
	}
	fmt.Println("Pointer marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fullJsonBlob), &pointer); err != nil {
		return err
	}
	fmt.Printf("Pointer unmarshal, with age = 0: %+v\n", pointer)

	value, err = json.Marshal(&pointer)
	if err != nil {
		return err
	}
	fmt.Println("Pointer marshal, with age = 0:", string(value))

	return nil
}