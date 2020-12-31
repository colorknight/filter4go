package filter

import (
	"filter4go/metadata"
	"fmt"
	"testing"
)

func TestCommonFilter(t *testing.T) {
	s1 := "'this is a test'"
	meta := ParseFilterMetadata(s1)
	str := metadata.ToString(meta)
	fmt.Printf("%s\n", str)
	s1 = "true"
	meta = ParseFilterMetadata(s1)
	str = metadata.ToString(meta)
	fmt.Printf("%s\n", str)
	s1 = "10.2"
	meta = ParseFilterMetadata(s1)
	str = metadata.ToString(meta)
	fmt.Printf("%s\n", str)
}

func TestCompareFilter(t *testing.T) {
	s1 := "key > 10"
	meta := ParseFilterMetadata(s1)
	str := metadata.ToString(meta)
	fmt.Printf("%s\n", str)
	s1 = "key <= 10"
	meta = ParseFilterMetadata(s1)
	str = metadata.ToString(meta)
	fmt.Printf("%s\n", str)
}

func TestCompare2Filter(t *testing.T) {
	s1 := "key + 3 > 10"
	meta := ParseFilterMetadata(s1)
	str := metadata.ToString(meta)
	fmt.Printf("%s\n", str)
	s1 = "key / 2<= 10"
	meta = ParseFilterMetadata(s1)
	str = metadata.ToString(meta)
	fmt.Printf("%s\n", str)
}

func TestCompare3Filter(t *testing.T) {
	s1 := "key.a + 3 > 10"
	meta := ParseFilterMetadata(s1)
	str := metadata.ToString(meta)
	fmt.Printf("%s\n", str)
	s1 = "key[2] / (2 | 1) <= 10"
	meta = ParseFilterMetadata(s1)
	str = metadata.ToString(meta)
	fmt.Printf("%s\n", str)
}

func TestBetweenFilter(t *testing.T) {
	s1 := "key between 10 and 20"
	meta := ParseFilterMetadata(s1)
	str := metadata.ToString(meta)
	fmt.Printf("%s\n", str)
}

func TestBetweenFilter2(t *testing.T) {
	s1 := "key.a + 10 between 10 and 20"
	meta := ParseFilterMetadata(s1)
	str := metadata.ToString(meta)
	fmt.Printf("%s\n", str)
}

func TestInFilter(t *testing.T) {
	s1 := "key in (1,2,3) "
	meta := ParseFilterMetadata(s1)
	str := metadata.ToString(meta)
	fmt.Printf("%s\n", str)
}

func TestIsFilter(t *testing.T) {
	s1 := "key is not null"
	meta := ParseFilterMetadata(s1)
	str := metadata.ToString(meta)
	fmt.Printf("%s\n", str)
}

func TestLogicFilter(t *testing.T) {
	s1 := "key + 3 > 10 and key.a * 2 = 6 or true"
	meta := ParseFilterMetadata(s1)
	str := metadata.ToString(meta)
	fmt.Printf("%s\n", str)
}

type DataBean struct {
	Name string
	Id  int
	Ary  []string
}

func buildData(count int) []map[string]interface{} {
	dataAry := []map[string]interface{}{}
	for i := 0; i < count; i++ {
		db := DataBean{fmt.Sprintf("a%d", i), i, []string{fmt.Sprint("A", i%10), fmt.Sprint("B", i%20), fmt.Sprint("C", i)}}
		mp := map[string]interface{}{
			"name" : db.Name,
			"bean" : db,
		}
		dataAry = append(dataAry, mp)
	}
	return dataAry
}

func TestCommonFilterOperand(t *testing.T) {
	s1 := "bean.Ary[0] = 'A1' and (bean.Id < 50 or bean.Id > 80)"
	filter, err := ParseFilter(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataAry := buildData(100)
	count := 0
	for _, mp := range dataAry {
		result, err := filter.BooleanOperate(mp)
		if err != nil {
			fmt.Println(err)
			return
		}
		if result {
			count++
		}
	}
	if count != 7 {
		t.Error(s1)
	}
}

func TestBetweenFilterOperand(t *testing.T) {
	s1 := "bean.Ary[0] = 'A1' and bean.Id + 10 between 50 and 80"
	filter, err := ParseFilter(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataAry := buildData(100)
	count := 0
	for _, mp := range dataAry {
		result, err := filter.BooleanOperate(mp)
		if err != nil {
			fmt.Println(err)
			return
		}
		if result {
			count++
		}
	}
	if count != 3 {
		t.Error(s1)
	}
}

func TestInFilterOperand3(t *testing.T) {
	s1 := "bean.Ary[0] in ('A1','A2') and bean.Id < 50"
	filter, err := ParseFilter(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataAry := buildData(100)
	count := 0
	for _, mp := range dataAry {
		result, err := filter.BooleanOperate(mp)
		if err != nil {
			fmt.Println(err)
			return
		}
		if result {
			count++
		}
	}
	if count != 10 {
		t.Error(s1)
	}
}

func TestLikeFilterOperand3(t *testing.T) {
	s1 := "name like 'a4_'"
	filter, err := ParseFilter(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataAry := buildData(100)
	count := 0
	for _, mp := range dataAry {
		result, err := filter.BooleanOperate(mp)
		if err != nil {
			fmt.Println(err)
			return
		}
		if result {
			count++
		}
	}
	if count != 10 {
		t.Error(s1)
	}
}

func TestNotLikeFilterOperand3(t *testing.T) {
	s1 := "not name like 'a4_'"
	filter, err := ParseFilter(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataAry := buildData(100)
	count := 0
	for _, mp := range dataAry {
		result, err := filter.BooleanOperate(mp)
		if err != nil {
			fmt.Println(err)
			return
		}
		if result {
			count++
		}
	}
	if count != 90 {
		t.Error(s1)
	}
}
