package app

import (
	"io/ioutil"
	"testing"

	"github.com/commander-cli/commander/pkg/runtime"
	"github.com/stretchr/testify/assert"
)

const TestSuiteFile = "/tmp/commander_test.yaml"

func Test_CommanderFile(t *testing.T) {
	tests := []byte(`
tests:
    echo hello:
        exit-code: 0
`)
	err := ioutil.WriteFile(TestSuiteFile, tests, 0755)

	assert.Nil(t, err)

	got := TestCommand(TestSuiteFile, TestCommandContext{})
	assert.Nil(t, got)
}

func Test_FailingSuite(t *testing.T) {
	tests := []byte(`
tests:
    echo hello:
        exit-code: 1
`)
	err := ioutil.WriteFile(TestSuiteFile, tests, 0755)

	assert.Nil(t, err)

	got := TestCommand(TestSuiteFile, TestCommandContext{})
	assert.Equal(t, "Test suite failed, use --verbose for more detailed output", got.Error())

}

func Test_WithTitle(t *testing.T) {
	tests := []byte(`
tests:
    my title:
        command: echo hello
        exit-code: 0
    another:
        command: echo another
        exit-code: 1
`)
	err := ioutil.WriteFile(TestSuiteFile, tests, 0755)

	assert.Nil(t, err)

	context := TestCommandContext{}
	context.Filters = runtime.Filters{"my title"}
	got := TestCommand(TestSuiteFile, context)
	assert.Nil(t, got)
}
