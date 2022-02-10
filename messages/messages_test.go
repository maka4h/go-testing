package messages

import (
	"testing"
)

func TestGreet(t *testing.T) {
	got := Greet("Gozer")
	expect := "Hello Gozer!\n"

	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got %q\n", got, expect)
	}
}

func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Gozer", expect: "Hello Gozer!\n"},
		{input: "Camil", expect: "Hello Camil!\n"},
		{input: "", expect: "Hello !\n"},
	}

	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("Did not get expected result for input '%v'. "+
				"Expected %q, got %q\n", s.input, got, s.expect)

		}
	}

}

func TestDepart(t *testing.T) {
	got := Depart("Gozer")
	expect := "Goodbye Gozer\n"

	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got %q\n", got, expect)
	}
}

//func TestFailureTypes(t *testing.T) {
//	t.Error("This is a non-blocking erros. The execution go ahead.")
//	t.Fatalf("This is a bloking error! The exection will be stopped.")
//	t.Error("You will never see this error...")
//}
