package r

import (
	"fmt"
	"reflect"
	"testing"
)

func TestG(t *testing.T) {
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
			name: "Simple type",
			args: args{
				subj:  &Person{Name: "Larry"},
				fName: "Name",
			},
			want: "Larry",
		},
		{
			name: "Map type with primitive type",
			args: args{
				subj:  Person{Name: "Vonny", StructWithMap: StructWithMap{RegularMap: map[string]string{"Foobar": "blah"}}},
				fName: "StructWithMap.RegularMap.Foobar",
				val:   "Vonny",
			},
			want: "blah",
		},
		{
			name: "Map type with struct type",
			args: args{
				subj:  Person{Name: "Vonny", StructWithMap: StructWithMap{StructMap: map[string]interface{}{"Foobar": Person{Name: "foobar"}}}},
				fName: "StructWithMap.StructMap.Foobar",
				val:   "Vonny",
			},
			want: Person{Name: "foobar"},
		},
		{
			name: "Map type with struct field",
			args: args{
				subj:  Person{Name: "Vonny", StructWithMap: StructWithMap{StructMap: map[string]interface{}{"Foobar": Person{Name: "foobar"}}}},
				fName: "StructWithMap.StructMap.Foobar.Name",
				val:   "Vonny",
			},
			want: "foobar",
		},
		{
			name: "Map type with pointer type",
			args: args{
				subj:  Person{Name: "Vonny", StructWithMap: StructWithMap{StructMap: map[string]interface{}{"Foobar": &Person{Name: "foobar"}}}},
				fName: "StructWithMap.StructMap.Foobar.Name",
				val:   "Vonny",
			},
			want: "foobar",
		},
		{
			name: "Passing in pointer to empty struct will return original value",
			args: args{
				subj:  pPerson,
				fName: "Name",
				val:   "Irrelevant",
			},
			want: nil,
		},
		{
			name: "Two Levels Deep",
			args: args{
				subj:  B{A: A{Person: Person{Name: "Éire"}}},
				fName: "A.Person.Name",
				val:   "Labhrás",
			},
			want: "Éire",
		},
		{
			name: "Nested type",
			args: args{
				subj:  B{A: A{Person: Person{Name: "foobar"}}},
				fName: "A.Person.Name",
			},
			want: "foobar",
		},
		{
			name: "Nested pointer type",
			args: args{
				subj:  &Bptr{A: &Aptr{Person: &Person{Name: "foobar"}}},
				fName: "A.Person.Name",
			},
			want: "foobar",
		},
		{
			name: "Maps Part Three",
			args: args{
				subj:  B{A: A{Person: Person{Name: "Cherry", StructWithMap: StructWithMap{RegularMap: map[string]string{"hello": "hello"}}}}},
				fName: "A.Person.StructWithMap.RegularMap",
				val:   map[string]string{"goodbye": "goodbye"},
			},
			want: map[string]string{"hello": "hello"},
		},
		{
			name: "Struct with unexported members",
			args: args{
				subj: map[string]interface{}{
					"data": UnexportedMemberStruct{Name: "SomeName", notexported: 1},
				},
				fName: "data.Name",
			},
			want: "SomeName",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := F(tt.args.subj, tt.args.fName); !reflect.DeepEqual(got, tt.want) {
				fmt.Printf("%#v", got)
				t.Errorf("R() = %v, want %v", got, tt.want)
			}
		})
	}
}
