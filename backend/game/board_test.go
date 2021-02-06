package game

import (
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	type args struct {
		boardname string
	}
	tests := []struct {
		name    string
		args    args
		want    *Board
		wantErr bool
	}{
		{
			name: "Test Normal",
			args: args{
				boardname: "狼王守衛",
			},
			want: &Board{
				Name: "狼王守衛",
				Characters: []*BoardCharacter{
					{
						Name: "狼王",
						Team: 1,
					},
					{
						Name: "狼人",
						Team: 1,
					},
					{
						Name: "狼人",
						Team: 1,
					},
					{
						Name: "狼人",
						Team: 1,
					},
					{
						Name: "預言家",
						Team: 2,
					},
					{
						Name: "女巫",
						Team: 2,
					},
					{
						Name: "獵人",
						Team: 2,
					},
					{
						Name: "守衛",
						Team: 2,
					},
					{
						Name: "平民",
						Team: 3,
					},
					{
						Name: "平民",
						Team: 3,
					},
					{
						Name: "平民",
						Team: 3,
					},
					{
						Name: "平民",
						Team: 3,
					},
				},
				NightFlow: []string{
					"狼人請睜眼。請選擇你們要擊殺的對象，確定是 ＿ 嗎？狼人請閉眼。",
					"預言家請睜眼。請選擇你要查驗的對象。他的身份是 ＿ 。預言家請閉眼。",
					"女巫請睜眼。今晚 ＿ 倒牌，請問你要使用解藥嗎？請問你要使用毒藥嗎？解藥給手勢，毒藥給數字。女巫請閉眼。",
					"獵人請睜眼。你的開槍狀態為 ＿ ？獵人請閉眼。",
					"守衛請睜眼。請選擇你要守護的對象，確定是 ＿ 嗎？守衛請閉眼。",
				},
				HasSheriff: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBoard(tt.args.boardname)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
