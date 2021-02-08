package math

import "testing"

func TestAnswer(t *testing.T){
	t.Run("返り値に42を返すこと", func(t *testing.T){

		if got := Answer(); got != 42 {
			t.Errorf("万物の答えは42ですが，この関数が返したのは%+vでした．",got)
		}
	})
}