package eval

import (
	"testing"
)

func TestReverseString1(t *testing.T) {
	res := New()
	tests := []struct {
		name     string
		cmd      []string
		evalExpr string
		want     bool
	}{
		{name: "single command and evalExpr match", cmd: []string{"ls /etc/hosts | awk -F \" \" '{print $1}' |awk 'FNR <= 1'"}, evalExpr: "'$0' == '/etc/hosts'", want: true},
		{name: "two command and evalExpr match", cmd: []string{"ls /etc/hosts | awk -F \" \" '{print $1}' |awk 'FNR <= 1'","ls /etc/group | awk -F \" \" '{print $1}' |awk 'FNR <= 1'"}, evalExpr: "'$0' == '/etc/hosts'; && '$1' == '/etc/group';", want: true},
		{name: "two command and evalExpr do not match", cmd: []string{"ls /etc/hosts | awk -F \" \" '{print $1}' |awk 'FNR <= 1'","ls /etc/group | awk -F \" \" '{print $1}' |awk 'FNR <= 1'"}, evalExpr: "'$0' == '/etc/hosts'; && '$1' == '/etc/group1';", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := res.EvalCommand(tt.cmd, tt.evalExpr); got.Match != tt.want {
				t.Errorf("CvssScoreToSeverity() = %v, want %v", got, tt.want)
			}
		})
	}
}
