package game

import (
	"reflect"
	"testing"
)

func TestNewCharacter(t *testing.T) {
	type args struct {
		charactername string
	}
	tests := []struct {
		name    string
		args    args
		want    *Character
		wantErr bool
	}{
		{
			name: "Test Normal",
			args: args{
				charactername: "女巫",
			},
			want: &Character{
				Name: "女巫",
				Team: 2,
				Hint: "女巫請睜眼。今晚 ＿ 倒牌，請問你要使用解藥嗎？請問你要使用毒藥嗎？解藥給手勢，毒藥給數字。女巫請閉眼。",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCharacter(tt.args.charactername)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCharacter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
