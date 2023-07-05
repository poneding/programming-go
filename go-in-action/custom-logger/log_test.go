package cuzlog_test

import (
	"cuzlog"
	"testing"
)

// go test -v cuzlog
func TestLog(t *testing.T) {
	cuzlog.SetFile("test.log")
	cuzlog.SetLevel(cuzlog.WARN)

	cuzlog.Tracef("Hello %s", "Jay Chou")
	cuzlog.Debugf("Hello %s", "Jay Chou")
	cuzlog.Infof("Hello %s", "Jay Chou")
	cuzlog.Warnf("Hello %s", "Jay Chou")
	cuzlog.Warnln("Hello", "Jay Chou", "and", "Miachael Jackson")

	cuzlog.SetFile("error.log")
	cuzlog.Errorf("Hello %s", "Jay Chou")
	//cuzlog.Fatalf("Hello %s", "Jay Chou")
}
