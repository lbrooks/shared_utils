package doublekey

import (
	"testing"
	"strconv"
)

func TestNew(t *testing.T) {
    dk := NewStringString()
    if len(dk.Data) != 0 {
        t.Errorf("New DK Has Length: %v", dk)
    }
}

const dataLength = 1000
const keyLength = 100

func BenchmarkAddSlice(b *testing.B) {
	data := make([]string, dataLength)
	for i := 0; i < dataLength; i++ {
		data[i] = strconv.Itoa(i)
	}
	keys := make([]string, keyLength)
	for i := 0; i < keyLength; i++ {
		keys[i] = strconv.Itoa(i)
	}

	b.ResetTimer()

    for i := 0; i < b.N; i++ {
        s := NewStringString()
		for _, key := range keys {
			s.AddSlice(key, data)	
		}
    }
}

func BenchmarkSizedAddSlice(b *testing.B) {
	data := make([]string, dataLength)
	for i := 0; i < dataLength; i++ {
		data[i] = strconv.Itoa(i)
	}
	keys := make([]string, keyLength)
	for i := 0; i < keyLength; i++ {
		keys[i] = strconv.Itoa(i)
	}

	b.ResetTimer()

    for i := 0; i < b.N; i++ {
        s := NewStringStringOfSize(dataLength * 2)
		for _, key := range keys {
			s.AddSlice(key, data)	
		}
    }
}

func TestAdd(t *testing.T) {
    dk := NewStringString()
    if len(dk.Data) != 0 {
        t.Errorf("New DK Has Length: %v", dk)
    }

	dk.Add("a", "b")

	if v1, o1 := dk.Data["a"]; o1 {
		if _, o2 := v1["b"]; !o2 {
			t.Errorf("Missing Key 2: %v", dk)
		}
	} else {
		t.Errorf("Missing Key 1: %v", dk)
	}
}

func TestAddSlice(t *testing.T) {
    dk := NewStringString()
    if len(dk.Data) != 0 {
        t.Errorf("New DK Has Length: %v", dk)
    }

	dk.AddSlice("a", []string{"b", "c"})

	if v1, o1 := dk.Data["a"]; o1 {
		if _, o2 := v1["b"]; !o2 {
			t.Errorf("Missing Key 2: %v", dk)
		}

		if _, o2 := v1["c"]; !o2 {
			t.Errorf("Missing Key 2: %v", dk)
		}
	} else {
		t.Errorf("Missing Key 1: %v", dk)
	}
}

func TestRemovePartial(t *testing.T) {
    dk := NewStringString()
    if len(dk.Data) != 0 {
        t.Errorf("New DK Has Length: %v", dk)
    }

	dk.AddSlice("a", []string{"b", "c"})

	if len(dk.Data) != 1 {
        t.Errorf("Added DK Has Wrong Length: %v", dk)
    }
	if len(dk.Data["a"]) != 2 {
        t.Errorf("Added DK Has Wrong Length: %v", dk)
    }

	dk.Remove("a", "b")

	if len(dk.Data) != 1 {
        t.Errorf("Added DK Has Wrong Length: %v", dk)
    }
	if len(dk.Data["a"]) != 1 {
        t.Errorf("Added DK Has Wrong Length: %v", dk)
    }
}

func TestRemoveFull(t *testing.T) {
    dk := NewStringString()
    if len(dk.Data) != 0 {
        t.Errorf("New DK Has Length: %v", dk)
    }

	dk.AddSlice("a", []string{"b"})

	if len(dk.Data) != 1 {
        t.Errorf("Added DK Has Wrong Length: %v", dk)
    }
	if len(dk.Data["a"]) != 1 {
        t.Errorf("Added DK Has Wrong Length: %v", dk)
    }

	dk.Remove("a", "b")

	if len(dk.Data) != 0 {
        t.Errorf("Added DK Has Wrong Length: %v", dk)
    }
}