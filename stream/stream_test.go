package stream

import "testing"

type Student struct {
	Name string
	Age  int
}

func TestStudentArr(t *testing.T) {

	students := []Student{
		{"Alice", 20},
		{"Bob", 18},
		{"Charlie", 19},
		{"David", 20},
		{"Eve", 18},
		{"Frank", 19},
		{"Grace", 20},
		{"Heidi", 18},
		{"Ivan", 19},
		{"Judy", 20},
		{"Kevin", 18},
		{"Lily", 19},
		{"Mallory", 20},
		{"Nate", 18},
		{"Oliver", 19},
		{"Peggy", 20},
		// give me some non-adult
		{"Quentin", 17},
		{"Romeo", 17},
		{"Steve", 17},
		{"Trent", 17},
		{"Uma", 17},
		{"Victor", 17},
		{"Walter", 17},
		{"Xavier", 17},
		{"Yvonne", 17},
		{"Zack", 17},
	}

	// filter

	adults := From(students).Filter(func(s Student) bool {
		return s.Age >= 18
	})

	t.Log("adults:", adults)

	// anyMatch

	any20 := adults.AnyMatch(func(s Student) bool {
		return s.Age == 20
	})

	t.Log("any20:", any20)

	// combine

	r := From(students).
		Filter(func(s Student) bool {
			return s.Age == 20
		}).
		AnyMatch(func(s Student) bool {
			return s.Name == "Alice"
		})

	t.Log("20 and Alice:", r)

	// map

	names := MapTo(From(students), func(s Student) string {
		return s.Name
	})

	t.Log("names:", names)


	s := From(students).Filter(func(s Student) bool {
		return s.Age > 17
	})

	m1 := ToMapStream(s, func(s Student)(string, int)  {
		return s.Name, s.Age
	})

	t.Logf("%+v", m1.ToMap())
	
	m2 := ToMapStream(s, func(s Student)(string, Student)  {
		return s.Name, s
	})

	t.Logf("%+v", m2.ToMap())

}

func TestMapStream(t *testing.T) {
	m := map[string]string{
		"1": "a",
		"2": "b",
		"3": "c",
	}

	ms := FromMap(m)
	r := ms.Filter(func(k, v string) bool {
		return k == "1"
	})

	t.Log(r.ToMap())

	r = ms.Map(func(k, v string) (string, string) {
		return k, v + "!"
	})

	t.Log(r.ToMap())

	// map to entries -> Stream[MapEntry[K, V]]
	r2 := ms.Entries().Find(func(e MapEntry[string, string]) bool {
		return e.Key == "1"
	})
	
	t.Logf("%v: %v", r2.Key, r2.Value)
	
}
