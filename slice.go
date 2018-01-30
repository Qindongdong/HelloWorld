package	slice

import(
	"hello"
	"testing"
)

type mathTest struct{
	a, b, ret int
}

var addTest = []mathTest{  
    mathTest{4, 6, 10},  
    mathTest{5, 6, 11},  
    mathTest{2, -6, -4},  
}  
  
var maxTest = []mathTest{  
    mathTest{3, 5, 5},  
    mathTest{-3, 5, 5},  
    mathTest{-3, -5, -3},  
}  
  
func TestAdd(t *testing.T) {  
    for _, v := range addTest {  
        ret := hello.Add(v.a, v.b)  
        if ret != v.ret {  
            t.Errorf("%d add %d, want %d, but get %d", v.a, v.b, v.ret, ret)  
        }  
    }  
}  
func TestMax(t *testing.T) {  
    for _, v := range maxTest {  
        ret := hello.Max(v.a, v.b)  
        if ret != v.ret {  
            t.Errorf("the max number between %d and %d is want %d, but get %d", v.a, v.b, v.ret, ret)  
        }  
    }  
}  