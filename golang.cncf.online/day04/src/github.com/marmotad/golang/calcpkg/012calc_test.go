package calcpkg

import "testing"

//单元测试
func TestAdd(t *testing.T) {
	if Add(1, 3) != 4 {
		t.Errorf("test Error")
	}
}

//性能测试
func BenchmarkFact(b *testing.B) {
	for i := 0; i < 10; i++ {
		Fact(i)
	}
}
