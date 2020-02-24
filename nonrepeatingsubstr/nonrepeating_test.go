package main

import "testing"

/*结果测试和覆盖率测试使用*/
func TestSubStr(t *testing.T) {
	tests := []struct {
		s   string
		num int
	}{
		{"123456", 6},
		{"我是一只小小鸟", 5},
	}

	for _, v := range tests {
		if actual := lengthOfNonRepeatingSubStr(v.s); actual != v.num {
			t.Errorf("lengthOfNonRepeatingSubStr(%s);actual:%d,expect,%d", v.s, actual, v.num)
		}
	}
}

/*测试代码的性能：运行速度*/
func BenchmarkSubStr(b *testing.B) {
	//性能测试结果BenchmarkSubStr-4   	 2688781	       491 ns/op
	s := "我是一只小小鸟"
	num := 6

	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Log(len(s))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubStr(s); actual != num {
			b.Errorf("lengthOfNonRepeatingSubStr(%s);actual:%d,expect,%d", s, actual, num)
		}
	}
}
