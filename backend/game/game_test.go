package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	// Without test board / id
	type args struct {
		boardname string
		host      *Member
	}
	tests := []struct {
		name    string
		args    args
		want    *Game
		wantErr bool
	}{
		{
			name: "Test Normal",
			args: args{
				boardname: "狼王守衛",
				host: &Member{
					Name: "test_member",
				},
			},
			want: &Game{
				Day: 0,
				Member: map[string]*Member{
					"test_member": {
						Name: "test_member",
					},
				},
				Host: &Member{
					Name: "test_member",
				},
				Player: []*Player{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGame(tt.args.boardname, tt.args.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want.Day, got.Day)
			assert.Equal(t, tt.want.Member, got.Member)
			assert.Equal(t, tt.want.Host, got.Host)
			assert.Equal(t, tt.want.Player, got.Player)
		})
	}
}

func TestGame_JoinRoom(t *testing.T) {
	type args struct {
		member *Member
	}
	tests := []struct {
		name         string
		args         args
		wantErr      error
		expectMember map[string]*Member
	}{
		{
			name: "Test Normal",
			args: args{
				member: &Member{
					Name: "user",
				},
			},
			expectMember: map[string]*Member{
				"admin": {
					Name: "admin",
				},
				"user": {
					Name: "user",
				},
			},
		},
		{
			name: "Test Exist",
			args: args{
				member: &Member{
					Name: "admin",
				},
			},
			expectMember: map[string]*Member{
				"admin": {
					Name: "admin",
				},
			},
			wantErr: fmt.Errorf("Player Already In Room"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := NewGame("狼王守衛", &Member{
				Name: "admin",
			})
			assert.Nil(t, err)
			err = game.JoinRoom(tt.args.member)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.expectMember, game.Member)
		})
	}
}

func TestGame_ExitRoom(t *testing.T) {
	type args struct {
		member *Member
	}
	tests := []struct {
		name         string
		args         args
		wantErr      error
		expectMember map[string]*Member
		originPlayer []*Player
		expectPlayer []*Player
	}{
		{
			name: "Test Normal",
			args: args{
				member: &Member{
					Name: "user",
				},
			},
			expectMember: map[string]*Member{
				"admin": {
					Name: "admin",
				},
			},
		},
		{
			name: "Test Normal With Seat",
			args: args{
				member: &Member{
					Name: "user",
				},
			},
			expectMember: map[string]*Member{
				"admin": {
					Name: "admin",
				},
			},
			originPlayer: []*Player{{
				Name: "user",
			}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
			expectPlayer: []*Player{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		},
		{
			name: "Test Not Exist",
			args: args{
				member: &Member{
					Name: "non-user",
				},
			},
			expectMember: map[string]*Member{
				"admin": {
					Name: "admin",
				},
				"user": {
					Name: "user",
				},
			},
			wantErr: fmt.Errorf("Player Not In Room"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := NewGame("狼王守衛", &Member{
				Name: "admin",
			})
			assert.Nil(t, err)
			err = game.JoinRoom(&Member{
				Name: "user",
			})
			assert.Nil(t, err)
			if tt.originPlayer != nil {
				game.Player = tt.originPlayer
			}

			err = game.ExitRoom(tt.args.member)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.expectMember, game.Member)
			if tt.expectPlayer != nil {
				assert.Equal(t, tt.expectPlayer, game.Player)
			}
		})
	}
}

func TestGame_ChangeHost(t *testing.T) {
	type args struct {
		member *Member
	}
	tests := []struct {
		name       string
		args       args
		wantErr    error
		expectHost *Member
	}{
		{
			name: "Test Normal",
			args: args{
				member: &Member{
					Name: "user",
				},
			},
			expectHost: &Member{
				Name: "user",
			},
		},
		{
			name: "Test Not Exist",
			args: args{
				member: &Member{
					Name: "non-user",
				},
			},
			expectHost: &Member{
				Name: "admin",
			},
			wantErr: fmt.Errorf("Player Not In Room"),
		},
		{
			name: "Test Change Host With Player",
			args: args{
				member: &Member{
					Name: "player",
				},
			},
			expectHost: &Member{
				Name: "admin",
			},
			wantErr: fmt.Errorf("Cannot Set Player As Host"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := NewGame("狼王守衛", &Member{
				Name: "admin",
			})
			assert.Nil(t, err)
			err = game.JoinRoom(&Member{
				Name: "user",
			})
			assert.Nil(t, err)
			err = game.JoinRoom(&Member{
				Name: "player",
			})
			assert.Nil(t, err)
			err = game.SetSeat(&Member{
				Name: "player",
			}, 0)
			assert.Nil(t, err)
			err = game.ChangeHost(tt.args.member)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.expectHost, game.Host)
		})
	}
}

func TestGame_SetSeat(t *testing.T) {
	type args struct {
		member *Member
		id     int
	}
	tests := []struct {
		name         string
		args         args
		wantErr      error
		originPlayer []*Player
		expectPlayer []*Player
	}{
		{
			name: "Test Normal",
			args: args{
				member: &Member{
					Name: "user",
				},
				id: 0,
			},
			expectPlayer: []*Player{{
				Name: "user",
			}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		},
		{
			name: "Test Error id Negative",
			args: args{
				member: &Member{
					Name: "user",
				},
				id: -1,
			},
			expectPlayer: []*Player{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
			wantErr:      fmt.Errorf("Seat ID Should Between 0 and 11"),
		},
		{
			name: "Test Error id Greater / Equal Than",
			args: args{
				member: &Member{
					Name: "user",
				},
				id: 12,
			},
			expectPlayer: []*Player{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
			wantErr:      fmt.Errorf("Seat ID Should Between 0 and 11"),
		},
		{
			name: "Test Not In Room",
			args: args{
				member: &Member{
					Name: "non-user",
				},
				id: 0,
			},
			expectPlayer: []*Player{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
			wantErr:      fmt.Errorf("Player Not In Room"),
		},
		{
			name: "Test Host",
			args: args{
				member: &Member{
					Name: "admin",
				},
				id: 0,
			},
			expectPlayer: []*Player{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
			wantErr:      fmt.Errorf("Host Cannot Enter Seat"),
		},
		{
			name: "Test Already Have Player",
			args: args{
				member: &Member{
					Name: "user",
				},
				id: 0,
			},
			originPlayer: []*Player{{
				Name: "player",
			}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
			expectPlayer: []*Player{{
				Name: "player",
			}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
			wantErr: fmt.Errorf("Seat Already Have Player"),
		},
		{
			name: "Test Change Seat",
			args: args{
				member: &Member{
					Name: "user",
				},
				id: 1,
			},
			originPlayer: []*Player{{
				Name: "user",
			}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
			expectPlayer: []*Player{nil, {
				Name: "user",
			}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := NewGame("狼王守衛", &Member{
				Name: "admin",
			})
			assert.Nil(t, err)
			if tt.originPlayer != nil {
				game.Player = tt.originPlayer
			}
			err = game.JoinRoom(&Member{
				Name: "user",
			})
			assert.Nil(t, err)
			err = game.SetSeat(tt.args.member, tt.args.id)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGame_ExitSeat(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name         string
		args         args
		wantErr      error
		originPlayer []*Player
		expectPlayer []*Player
	}{
		{
			name: "Test Normal",
			args: args{
				id: 0,
			},
			originPlayer: []*Player{{
				Name: "user",
			}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
			expectPlayer: []*Player{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		},
		{
			name: "Test Error id Negative",
			args: args{
				id: -1,
			},
			expectPlayer: []*Player{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
			wantErr:      fmt.Errorf("Seat ID Should Between 0 and 11"),
		},
		{
			name: "Test Error id Greater / Equal Than",
			args: args{
				id: 12,
			},
			expectPlayer: []*Player{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
			wantErr:      fmt.Errorf("Seat ID Should Between 0 and 11"),
		},
		{
			name: "Test No User",
			args: args{
				id: 0,
			},
			expectPlayer: []*Player{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
			wantErr:      fmt.Errorf("Player Not On The Seat"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := NewGame("狼王守衛", &Member{
				Name: "admin",
			})
			assert.Nil(t, err)
			if tt.originPlayer != nil {
				game.Player = tt.originPlayer
			}
			assert.Nil(t, err)
			err = game.ExitSeat(tt.args.id)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
