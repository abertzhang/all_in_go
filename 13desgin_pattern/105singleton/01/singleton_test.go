package pattern_test	
import (
	"testing"
	"studyGo/13desgin_pattern/05singleton/01/pattern"
)
	func TestSingleton(t *testing.T){
	s1:=pattern.GetInstance(100)
	s2:=pattern.GetInstance(200)
	if s1!=s2 {
		t.Fatal("error! singleton not working")
	} 
	if s1.GetValue()!=100 || s2.GetValue()!=100 {
		t.Fatal("error! singleton value not working")
	} 
}