package fourthres

import (
	"testing"
)

func TestAdopting(t *testing.T) {
	t.Run("Given hippo, shouldn`t receive err", func(t *testing.T) {
		type testPet struct {
			kind string
			name string
			age  int
		}
		var pet testPet = testPet{
			kind: "hippopotamus",
			name: "Hippo",
			age:  22,
		}
		_, err := adoptPet(pet.kind, pet.name, pet.age)
		if err != nil {
			t.Errorf("got %v error", err)
		}
	})
	t.Run("Given cat, should receive err", func(t *testing.T) {
		type testPet struct {
			kind string
			name string
			age  int
		}
		var pet testPet = testPet{
			kind: "cat",
			name: "Martha",
			age:  3,
		}
		_, err := adoptPet(pet.kind, pet.name, pet.age)
		if err.Error() != "We can`t adopt this kind yet" {
			t.Errorf("should receive err: We can`t adopt this kind yet, got: %v", err)
		}
	})
}
func TestGetInfo(t *testing.T) {

	t.Run("Given hippo, shouldn`t receive err", func(t *testing.T) {
		type testPet struct {
			kind string
			name string
			age  int
		}
		var pet testPet = testPet{
			kind: "hippopotamus",
			name: "Hippo",
			age:  22,
		}
		hippo, _ := adoptPet(pet.kind, pet.name, pet.age)
		var testHippo testPet
		testHippo.name, testHippo.age = zoo.GetInfo(hippo)
		if testHippo.name != pet.name {
			t.Errorf("Wrong name want: %v, got: %v", pet.name, testHippo.name)
		}
		if testHippo.age != pet.age {
			t.Errorf("Wrong age want: %v, got: %v", pet.age, testHippo.age)
		}
	})
}
