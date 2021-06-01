package set

import (
	"fmt"
	"testing"
)

func TestSetCreate(t *testing.T) {
	fmt.Println("Testing: CreateEmptySet")

	i := 0
	for i < 10 {
		test := CreateEmptySet()
		if test.size_ != 0 {
			t.Fail()
		}
		i++
	}
}

func TestAdd(t *testing.T) {
	fmt.Println("Testing: add")
	testSet := CreateEmptySet()

	//testing basic bulk add
	for i := 0; i < 100; i++ {
		testSet.add(i)

		if val := testSet.container[i]; val == false {
			t.Fail()
		}
	}

	//testing adding an existing value
	result := testSet.add(0)
	if result {
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	fmt.Println("Testing: contains")

	testSet := CreateEmptySet()
	for i := 0; i < 1000; i++ {
		testSet.add(i)
	}

	for i := 300; i <= 700; i++ {
		result := testSet.contains(i)
		if !result {
			t.Fail()
		}
	}

	result := testSet.contains(1001)
	if result {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	fmt.Println("Testing: remove")

	testSet := CreateEmptySet()
	for i := 0; i > -1000; i-- {
		testSet.add(i)
	}

	for i := 0; i > -1000; i-- {
		result := testSet.remove(i)
		if !result || testSet.contains(i) {
			t.Fail()
		}
	}

	if result := testSet.remove(10); result {
		t.Fail()
	}
}

func TestSize(t *testing.T) {
	fmt.Println("Testing: size")

	test := CreateEmptySet()

	for i := 0; i < 100; i++ {
		if test.size() != i {
			t.Fail()
		}
		test.add(i)
	}

	for i := 99; i >= 0; i-- {
		test.remove(i)
		if test.size() != i {
			t.Fail()
		}
	}
}

func TestUnion(t *testing.T) {
	fmt.Println("Testing: union")
	set1 := CreateEmptySet()
	set2 := CreateEmptySet()

	i := 1
	for i <= 50 {
		set1.add(i)
		i++
	}
	for i <= 100 {
		set2.add(i)
		i++
	}

	var set3 Set = set2
	var union Set = set1.union(&set3)
	i -= 1

	switch ty := union.(type) {
	case StaticIntSet:
		for i >= 1 {
			if !ty.contains(i) {
				t.Fail()
			}
			i--
		}
	default:
		fmt.Printf("Cant test on type: %T", ty)
	}

}

func TestDifference(t *testing.T) {
	fmt.Println("Testing: difference")

	set1 := CreateEmptySet()
	set2 := CreateEmptySet()

	for i := 1; i <= 51; i++ {
		set1.add(i)
	}
	for i := 51; i <= 100; i++ {
		set2.add(i)
	}

	set3 := Set(set2)
	diff := set1.difference(&set3)

	switch ty := diff.(type) {
	case StaticIntSet:
		if ty.size() != 50 {
			t.Fail()
		}
	}

	set1 = CreateEmptySet()
	set2 = CreateEmptySet()

	for i := 1; i <= 100; i++ {
		set1.add(i)
	}
	for i := 1; i <= 100; i++ {
		set2.add(i)
	}

	set3 = Set(set2)
	diff = set1.difference(&set3)

	switch ty := diff.(type) {
	case StaticIntSet:
		if ty.size() != 0 {
			t.Fail()
		}
	}
}

func TestIntersection(t *testing.T) {
	fmt.Println("Testing: intersection")
	set1 := CreateEmptySet()
	set2 := CreateEmptySet()

	for i := 1; i <= 51; i++ {
		set1.add(i)
	}
	for i := 51; i <= 100; i++ {
		set2.add(i)
	}

	set3 := Set(set2)
	diff := set1.intersection(&set3)

	switch ty := diff.(type) {
	case StaticIntSet:
		if ty.size() != 1 {
			t.Fail()
		}
	}

	set1 = CreateEmptySet()
	set2 = CreateEmptySet()

	for i := 1; i <= 100; i++ {
		set1.add(i)
	}
	for i := 1; i <= 100; i++ {
		set2.add(i)
	}

	set3 = Set(set2)
	diff = set1.intersection(&set3)

	switch ty := diff.(type) {
	case StaticIntSet:
		if ty.size() != 100 {
			t.Fail()
		}
	}
}

func TestSubset(t *testing.T) {
	fmt.Println("Testing: subset")
	set1 := CreateEmptySet()
	set2 := CreateEmptySet()

	for i := 1; i <= 51; i++ {
		set1.add(i)
	}
	for i := 51; i <= 100; i++ {
		set2.add(i)
	}

	set3 := Set(set2)

	if set1.subset(&set3) {
		t.Fail()
	}

	set1 = CreateEmptySet()
	set2 = CreateEmptySet()

	for i := 1; i <= 50; i++ {
		set1.add(i)
	}
	for i := 2; i <= 48; i++ {
		set2.add(i)
	}

	set3 = Set(set2)

	if !set1.subset(&set3) {
		t.Fail()
	}

}
