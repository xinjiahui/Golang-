package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case_1": testCase{"backfsbsf", "b", []string{"", "ackfs", "sf"}},
		"case_2": testCase{"test:fsf", ":", []string{"test", "fsf"}},
		"case_3": testCase{"abcef", "bc", []string{"a", "ef"}},
		"case_4": testCase{"我是好人", "是", []string{"我", "好人"}},
	}
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}

		})
	}
}

//基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d", ":")
	}
}
