package comma

import "testing"

func TestComma(t *testing.T) {
	if comma("123456") != "123,456" || comma("234") != "234" || comma("1234") != "1,234" {
		t.Error("TestComma failed")
	}
}
