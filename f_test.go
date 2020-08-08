package r

import (
	"reflect"
	"testing"
)

func TestF(t *testing.T) {
	t.Parallel()
	type args struct {
		subj  interface{}
		fName string
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
			},
			want: "Lar",
		},
		{
			name: "Pointer to Struct will return Struct",
			args: args{
				subj:  &Person{Name: "Larry"},
				fName: "Name",
			},
			want: "Larry",
		},
		{
			name: "Will Get Empty Field",
			args: args{
				subj:  Person{},
				fName: "Name",
			},
			want: "",
		},
		{
			name: "Passing in pointer to empty struct will return original value",
			args: args{
				subj:  pPerson,
				fName: "Name",
			},
			want: nil,
		},
		{
			name: "Two Levels Deep",
			args: args{
				subj:  B{A: A{Person: Person{Name: "Éire"}}},
				fName: "A.Person.Name",
			},
			want: "Éire",
		},
		{
			name: "Maps Part One",
			args: args{
				subj:  StructWithMap{StructMap: map[string]interface{}{"oneHundred": Nested{100}}},
				fName: "StructMap",
			},
			want: map[string]interface{}{"oneHundred": Nested{100}},
		},
		{
			name: "Maps Part Two",
			args: args{
				subj:  B{A: A{Person: Person{StructWithMap: StructWithMap{StructMap: map[string]interface{}{"oneHundred": Nested{100}}}}}},
				fName: "A.Person.StructWithMap.StructMap",
			},
			want: map[string]interface{}{"oneHundred": Nested{100}},
		},
		{
			name: "Maps Part Three",
			args: args{
				subj:  B{A: A{Person: Person{Name: "Cherry", StructWithMap: StructWithMap{RegularMap: map[string]string{"hello": "hello"}}}}},
				fName: "A.Person.StructWithMap.RegularMap",
			},
			want: map[string]string{"hello": "hello"},
		},
		{
			name: "Slice",
			args: args{
				subj:  Person{Name: "Milhouse", Slice: []string{"one", "two"}},
				fName: "Slice",
			},
			want: []string{"one", "two"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := F(tt.args.subj, tt.args.fName); !reflect.DeepEqual(got, tt.want) {
				typeofgot := reflect.TypeOf(got)
				typeofwant := reflect.TypeOf(tt.want)
				t.Log(typeofgot, typeofwant)
				t.Errorf("F() = %v, want %v", got, tt.want)
			}
		})
	}
}
