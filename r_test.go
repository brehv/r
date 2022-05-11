package r

import (
	"reflect"
	"testing"
)

type UnexportedMemberStruct struct {
	notexported int
	Name        string
}

type Person struct {
	Name          string
	StructWithMap StructWithMap
	Slice         []string
}
type StructWithMap struct {
	RegularMap map[string]string
	StructMap  map[string]interface{}
}
type A struct {
	Person Person
}
type B struct {
	A A
}
type Nested struct {
	i int
}

type Aptr struct {
	Person *Person
}
type Bptr struct {
	A *Aptr
}

var pPerson *Person

func TestR(t *testing.T) {
	t.Parallel()
	type args struct {
		subj  interface{}
		fName string
		val   interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Struct will return Struct",
			args: args{
				subj:  Person{Name: "Lar"},
				fName: "Name",
				val:   "Vonny",
			},
			want: Person{Name: "Vonny"},
		},
		{
			name: "Pointer to Struct will return Struct",
			args: args{
				subj:  &Person{Name: "Larry"},
				fName: "Name",
				val:   "Yvonne",
			},
			want: Person{Name: "Yvonne"},
		},
		{
			name: "Will Set Empty Field",
			args: args{
				subj:  Person{},
				fName: "Name",
				val:   "威",
			},
			want: Person{Name: "威"},
		},
		{
			name: "Passing in pointer to empty struct will return original value",
			args: args{
				subj:  pPerson,
				fName: "Name",
				val:   "Irrelevant",
			},
			want: pPerson,
		},
		{
			name: "Two Levels Deep",
			args: args{
				subj:  B{A: A{Person: Person{Name: "Éire"}}},
				fName: "A.Person.Name",
				val:   "Labhrás",
			},
			want: B{A: A{Person: Person{Name: "Labhrás"}}},
		},
		{
			name: "Maps Part One",
			args: args{
				subj:  StructWithMap{StructMap: map[string]interface{}{"oneHundred": Nested{100}}},
				fName: "StructMap",
				val:   map[string]interface{}{"oneHundred": Nested{999}},
			},
			want: StructWithMap{StructMap: map[string]interface{}{"oneHundred": Nested{999}}},
		},
		{
			name: "Maps Part Two",
			args: args{
				subj:  B{A: A{Person: Person{StructWithMap: StructWithMap{StructMap: map[string]interface{}{"oneHundred": Nested{100}}}}}},
				fName: "A.Person.StructWithMap.StructMap",
				val:   map[string]interface{}{"oneHundred": Nested{999}},
			},
			want: B{A: A{Person: Person{StructWithMap: StructWithMap{StructMap: map[string]interface{}{"oneHundred": Nested{999}}}}}},
		},
		{
			name: "Maps Part Three",
			args: args{
				subj:  B{A: A{Person: Person{Name: "Cherry", StructWithMap: StructWithMap{RegularMap: map[string]string{"hello": "hello"}}}}},
				fName: "A.Person.StructWithMap.RegularMap",
				val:   map[string]string{"goodbye": "goodbye"},
			},
			want: B{A: A{Person: Person{Name: "Cherry", StructWithMap: StructWithMap{RegularMap: map[string]string{"goodbye": "goodbye"}}}}},
		},
		{
			name: "Slice",
			args: args{
				subj:  Person{Name: "Milhouse", Slice: []string{"one", "two"}},
				fName: "Slice",
				val:   []string{"three", "legs"},
			},
			want: Person{Name: "Milhouse", Slice: []string{"three", "legs"}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := R(tt.args.subj, tt.args.fName, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("R() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNil(t *testing.T) {
	t.Parallel()
	s := "Hello World"
	var np *string = nil
	var m map[string]string
	var a = [6]int{2, 3, 5, 7, 11, 13}
	ch := make(chan int)
	var nch chan int
	sl := []string{"hello", "world"}
	var nsl []string

	tests := []struct {
		name string
		args interface{}
		want bool
	}{
		{
			name: "Nil is nil",
			args: nil,
			want: true,
		},
		{
			name: "Nil Pointer is Nil",
			args: np,
			want: true,
		},
		{
			name: "Pointer is not Nil",
			args: &s,
			want: false,
		},
		{
			name: "Nil Map is nil",
			args: m,
			want: true,
		},
		{
			name: "Map is not nil",
			args: map[string]string{"hello": "world"},
			want: false,
		},
		{
			name: "Array is not nil",
			args: a,
			want: false,
		},
		{
			name: "Nil Channel is nil",
			args: nch,
			want: true,
		},
		{
			name: "Channel is not nil",
			args: ch,
			want: false,
		},
		{
			name: "Nil Slice is nil",
			args: nsl,
			want: true,
		},
		{
			name: "Slice is not nil",
			args: sl,
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := isNil(tt.args); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
