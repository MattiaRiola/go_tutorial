package freemail

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var providerTesters = [...]func(t *testing.T){
	testGmail,
	testOutlook,
	testYahoo,
	testStark,
}

const (
	// set this to true to see what happens for all the tests
	FAILING_TESTS = false
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func TestKnownFreemail(t *testing.T) {
	var knownProviders = providerTesters[:3]
	for _, f := range knownProviders {
		t.Log(GetFunctionName(f))
		t.Run(GetFunctionName(f), f)
	}

}

func TestIsFreemail(t *testing.T) {
	if FAILING_TESTS {
		t.Log("Testing some providers")
		for _, f := range providerTesters {
			t.Run(GetFunctionName(f), f)
		}
		t.Log("all providers tested")
		t.Run("TestFailure", testFailure)
		t.Run("TestOk", testOk)
	}
}

func TestOkAndFail(t *testing.T) {
	if FAILING_TESTS {
		t.Run("TestFailure", testFailure)
		t.Run("TestOk", testOk)
	}
}

func testGmail(t *testing.T) {
	gmail := "user@gmail.com"
	if !IsFreemail(gmail) {
		t.Fail()
		t.Log("gmail isn't recognized as freeMail")
	}
}
func testOutlook(t *testing.T) {
	outlook := "user@outlook.com"
	if !IsFreemail(outlook) {
		t.Fail()
		t.Log("outlook isn't recognized as freeMail")
	}
}
func testYahoo(t *testing.T) {
	yahoo := "user@yahoo.com"
	if !IsFreemail(yahoo) {
		t.Fail()
		t.Log("yahoo isn't recognized as freeMail")
	}
}
func testStark(t *testing.T) {
	stark := "user@stark.com"
	if !IsFreemail(stark) {
		t.Fail()
		t.Log("stark isn't recognized as freeMail")
	}
}

func testFailure(t *testing.T) {
	t.Log("Fail wont stop the execution of the test")
	t.Fail()
	//To stop the execution use FailNow or Fatal:
	// t.Log("Log from TestLogFailure")
	// t.FailNow()
	//They are equivalent to:
	t.Fatal("TestLogFailure: Fatal now the execution of the tests will end")
	t.Log("This log shouldn't be displayed")
}
func testOk(t *testing.T) {
	myTrue := alwaysTrue()

	assert.True(t, myTrue)

}

func alwaysTrue() bool {
	return true
}
