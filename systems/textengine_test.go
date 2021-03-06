package systems_test

import (
	"fmt"
	"github.com/Lathie/fulcrum/systems"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

//Taken from @benbjohnson
// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

//Test that the NewEngine function works correctly
func TestNewEngine(t *testing.T) {
	engine := systems.NewEngine(true)
	assert(t, engine.Running, "Engine is not Running")
}

//Test that the Update function for textengine works correctly
func TestUpdate(t *testing.T) {
	engine := systems.NewEngine(true)
	boolean := engine.Update()
	assert(t, boolean, "Update cycle failed")
}
