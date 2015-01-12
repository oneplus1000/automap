package automap

import (
	"testing"
)

func _TestInt(t *testing.T) {
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

func _TestString(t *testing.T) {
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
	//t.Errorf("%v", dest)
}

func _Test01(t *testing.T) {

	src01 := Src01{
		Id: 1,
		ArrInt: []int{
			123,
			456,
		},
	}
	var dest01 Dest01

	bm := new(AutoMapper)
	err := bm.Auto(&src01, &dest01)
	if err != nil {
		t.Errorf("err := bm.Auto(&src01, &dest01); err=%s", err.Error())
	}

	if src01.Id != dest01.Id {
		t.Errorf("src01.Id != dest01.Id")
	}

	if len(src01.ArrInt) != len(dest01.ArrInt) {
		t.Errorf("len(src01.ArrInt) != len(dest01.ArrInt); len(src01.ArrInt) == %d ; len(dest01.ArrInt == %d ", len(src01.ArrInt), len(dest01.ArrInt))
	} else {

		i := 0
		max := len(src01.ArrInt)
		for i < max {
			if src01.ArrInt[i] != dest01.ArrInt[i] {
				t.Errorf("rc01.ArrInt[%d] != dest01.ArrInt[%d]", i)
			}
			i++
		}

	}
}

type Src01 struct {
	Id     int
	ArrInt []int
}

type Dest01 struct {
	Id     int
	ArrInt []int
}
