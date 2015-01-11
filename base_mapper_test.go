package basemapper

import (
	"testing"
)

func Test01(t *testing.T) {

	src01 := Src01{
		Id: 1,
		ArrInt: []int{
			123,
			456,
		},
	}
	var dest01 Dest01

	bm := new(BaseMapper)
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
