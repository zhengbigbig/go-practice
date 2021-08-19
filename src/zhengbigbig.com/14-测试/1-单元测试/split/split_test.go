package split

import (
	"fmt"
	"reflect"
	"testing"
)

// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行：测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行：测试之后的teardown")
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行：子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行：自测试之后的teardown")
	}
}

func TestSplit(t *testing.T) {
	type testCase struct {
		input string
		sep   string
		want  []string
	}
	testGroup := map[string]testCase{
		"case1": {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"case2": {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"case3": {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"case4": {input: "哈喽world", sep: "喽", want: []string{"哈", "world"}},
		"case5": {input: "abcd", sep: "", want: []string{"a", "b", "c", "d"}},
	}
	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t) 			 // 测试之后执行teardown操作
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
			defer teardownSubTest(t)			// 测试之后执行teardown操作
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%v，got:%v", tc.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("abcd","b")
	}
}

func BenchmarkSplitParallel(b *testing.B) {
	//b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB){
		for pb.Next(){
			Split("abcd","b")
		}
	})
}

func ExampleSplit() {
	fmt.Println(Split("a:b:c",":"))
	fmt.Println(Split("abcd","bc"))
	// Output:
	// [a b c]
	// [a d]
}