package r

import "fmt"

type Human struct {
	Name    string
	Parents Parents
}
type Parents struct {
	Mom string
	Dad string
}

func Example_r() {
	p := Human{
		Name: "Milhouse",
		Parents: Parents{
			Mom: "Cherry",
			Dad: "Dennis",
		},
	}
	fmt.Println(p)
	updated := R(p, "Parents.Dad", "Larry")
	fmt.Println(updated)

	// Output:
	//{Milhouse {Cherry Dennis}}
	//{Milhouse {Cherry Larry}}
}

func Example_g() {
	p := Human{
		Name: "Milhouse",
		Parents: Parents{
			Mom: "Cherry",
			Dad: "Dennis",
		},
	}
	fmt.Println(p)
	specificField := F(p, "Parents.Dad")
	fmt.Println(specificField)

	// Output:
	//{Milhouse {Cherry Dennis}}
	//Dennis
}
