package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":")         // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能直接比较，借助反射包中的方法进行比较
		t.Errorf("excepted:%v,got:%v", want, got) // 测试失败输出错误提示
	}
}

func Test2Split(t *testing.T) {
	got := Split("a:b:c", "")
	want := []string{"a", ":", "b", ":", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v,got:%v", want, got)
	}
}
