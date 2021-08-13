package eval

import (
	"github.com/chen-keinan/go-command-eval/utils"
	"testing"
)

func TestCommandParams(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want map[int][]string
	}{
		{name: "two command and one param", cmd: []string{" aaa", "bb #1"}, want: map[int][]string{1: {"1"}}},
		{name: "two command and 2 params on 2 commands", cmd: []string{" aaa", "bb #1", "cc #2"}, want: map[int][]string{1: {"1"}, 2: {"2"}}},
		{name: "two command and 2 params on one command", cmd: []string{" aaa", "bb #1", "cc #1 #2"}, want: map[int][]string{1: {"1"}, 2: {"1", "2"}}},
		{name: "two command no params", cmd: []string{" aaa", "bb ", "cc"}, want: map[int][]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CommandParams(tt.cmd)
			if len(tt.want) != len(got) {
				t.Errorf("CommandParams() = %v, want %v", got, tt.want)
			}
			for key, value := range tt.want {
				if val, ok := got[key]; ok {
					for k, v := range val {
						if v != value[k] {
							t.Errorf("CommandParams() = %v, want %v", got, tt.want)
						}
					}
				} else {
					{
						t.Errorf("CommandParams() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}

func TestEvalExpression(t *testing.T) {

	tests := []struct {
		name        string
		commandRes  []string
		commResSize int
		testFailure int
		evalExpr    string
		want        int
		wantErr     error
	}{
		{name: "one command res and one param", commandRes: []string{"/etc/hosts"}, commResSize: 1, testFailure: 0, evalExpr: "'$0' == '/etc/hosts'", want: 0, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdEval := cmd{cmdExprBuilder: utils.UpdateCmdExprParam, evalExpr: tt.evalExpr}
			got, err := cmdEval.evalExpression(tt.commandRes, tt.commResSize, make([]string, 0), tt.testFailure)
			if tt.want != got && err != tt.wantErr {
				t.Errorf("evalExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
