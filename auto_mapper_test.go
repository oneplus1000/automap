package automap

import (
	"testing"
)

func TestInt(t *testing.T) {
	var src = int(12)
	var dest = int(32)
	bm := new(AutoMapper)
	err := bm.Auto(&src, &dest)

	if err != nil {
		t.Errorf("err := bm.Auto(&src01, &dest01); err=%s", err.Error())
	}

	if dest != src {
		t.Errorf("dest != src ; dest = %d", dest)
	}

}

func TestString(t *testing.T) {
	var src = string("dddddddd")
	var dest = string("")
	bm := new(AutoMapper)
	err := bm.Auto(&src, &dest)

	if err != nil {
		t.Errorf("err := bm.Auto(&src01, &dest01); err=%s", err.Error())
		return
	}

	if dest != src {
		t.Errorf("dest != src ; dest = %s", dest)
		return
	}
}

func TestArrOfInt(t *testing.T) {
	var src = []int{12, 21, 33}
	var dest []int
	bm := new(AutoMapper)
	err := bm.Auto(&src, &dest)

	if err != nil {
		t.Errorf("err := bm.Auto(&src01, &dest01); err=%s", err.Error())
		return
	}

	if len(dest) != len(src) {
		t.Errorf("len(dest) != len(src) ; len(dest) = %d ", len(dest))
		return
	}

	i := 0
	max := len(src)
	for i < max {
		if src[i] != dest[i] {
			t.Errorf("dest[%d] != src[%d] ; dest = %d", i, i, dest)
			return
		}
		i++
	}
}

func TestObj01(t *testing.T) {
	s01 := S01{
		Id:   1,
		Name: "eeeeee eeee",
	}
	var d01 D01
	bm := new(AutoMapper)
	err := bm.Auto(&s01, &d01)
	if err != nil {
		t.Errorf("bm.Auto(&src, &dest); err = %s", err.Error())
		return
	}

	if d01.Id != s01.Id {
		t.Errorf("d01.Id != s01.Id")
		return
	}

	if d01.Name != s01.Name {
		t.Errorf("d01.Name != s01.Name")
		return
	}

	//t.Errorf("%#v", d01)
}

type D01 struct {
	Id   int
	Name string
}

type S01 struct {
	Id   int
	Name string
}

func TestObj02(t *testing.T) {

	src01 := Src01{
		Id: 1,
		ArrInt: []int{
			123,
			456,
		},
		Im: SrcItem01{
			Name: "oneplus",
		},
		Items: []SrcItem01{
			SrcItem01{
				Name: "noi",
			},
			SrcItem01{
				Name: "noi2",
			},
		},
	}
	var dest01 Dest01

	bm := new(AutoMapper)
	err := bm.Auto(&src01, &dest01)
	if err != nil {
		t.Errorf("err := bm.Auto(&src01, &dest01); err=%s", err.Error())
		return
	}

	if src01.Id != dest01.Id {
		t.Errorf("src01.Id != dest01.Id")
		return
	}

	if len(src01.ArrInt) != len(dest01.ArrInt) {
		t.Errorf("len(src01.ArrInt) != len(dest01.ArrInt); len(src01.ArrInt) == %d ; len(dest01.ArrInt == %d ", len(src01.ArrInt), len(dest01.ArrInt))
		return
	} else {
		i := 0
		max := len(src01.ArrInt)
		for i < max {
			if src01.ArrInt[i] != dest01.ArrInt[i] {
				t.Errorf("rc01.ArrInt[%d] != dest01.ArrInt[%d]", i)
				return
			}
			i++
		}

	}

	if src01.Im.Name != dest01.Im.Name {
		t.Errorf("src01.Im.Name != dest01.Im.Name")
	}

	if len(src01.Items) != len(dest01.Items) {
		t.Errorf("len(src01.Items) != len(dest01.Items)")
	}

	t.Errorf("%v", dest01)
}

type Src01 struct {
	Id     int
	ArrInt []int
	Im     SrcItem01
	Items  []SrcItem01
}

type SrcItem01 struct {
	Name string
}

type Dest01 struct {
	Id     int
	ArrInt []int
	Im     DestItem01
	Items  []DestItem01
}

type DestItem01 struct {
	Name string
}
