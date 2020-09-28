package main

import "testing"

func TestHello(t *testing.T) {

	//断言方法
	assertCorrectMessage := func(t *testing.T, got, want string) {
		//辅助函数，失败报告调用行数，而不是内部行数
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	//子测试
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
}
