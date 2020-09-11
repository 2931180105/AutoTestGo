package comptroller

import "testing"

var compRun = NewCmpControllReader()

func TestAdmin(t *testing.T) {
	//compRun.Wingsupplystate()
	compRun.Admin(compRun.Cfg.Comptroller)
}

//func TestMinAllowed(t *testing.T) {
//	//compRun.Wingsupplystate()
//	compRun.Mintallowed()
//}
