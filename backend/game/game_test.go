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
					Username: "test_member",
				},
			},
			want: &Game{
				Day: 0,
				Member: map[string]*Member{
					"test_member": {
						Username: "test_member",
					},
				},
				Host: &Member{
					Username: "test_member",
				},
				Player: []*Player{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
				End:    false,
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
			assert.Equal(t, tt.want.End, got.End)
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
					Username: "user",
				},
			},
			expectMember: map[string]*Member{
				"admin": {
					Username: "admin",
				},
				"user": {
					Username: "user",
				},
			},
		},
		{
			name: "Test Exist",
			args: args{
				member: &Member{
					Username: "admin",
				},
			},
			expectMember: map[string]*Member{
				"admin": {
					Username: "admin",
				},
			},
			wantErr: fmt.Errorf("Player Already In Room"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := NewGame("狼王守衛", &Member{
				Username: "admin",
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
					Username: "user",
				},
			},
			expectMember: map[string]*Member{
				"admin": {
					Username: "admin",
				},
			},
		},
		{
			name: "Test Normal With Seat",
			args: args{
				member: &Member{
					Username: "user",
				},
			},
			expectMember: map[string]*Member{
				"admin": {
					Username: "admin",
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
					Username: "non-user",
				},
			},
			expectMember: map[string]*Member{
				"admin": {
					Username: "admin",
				},
				"user": {
					Username: "user",
				},
			},
			wantErr: fmt.Errorf("Player Not In Room"),
		},
		{
			name: "Test Host Exit",
			args: args{
				member: &Member{
					Username: "admin",
				},
			},
			expectMember: map[string]*Member{
				"admin": {
					Username: "admin",
				},
				"user": {
					Username: "user",
				},
			},
			wantErr: fmt.Errorf("Host Cannot Exit Room"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := NewGame("狼王守衛", &Member{
				Username: "admin",
			})
			assert.Nil(t, err)
			err = game.JoinRoom(&Member{
				Username: "user",
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
					Username: "user",
				},
			},
			expectHost: &Member{
				Username: "user",
			},
		},
		{
			name: "Test Not Exist",
			args: args{
				member: &Member{
					Username: "non-user",
				},
			},
			expectHost: &Member{
				Username: "admin",
			},
			wantErr: fmt.Errorf("Player Not In Room"),
		},
		{
			name: "Test Change Host With Player",
			args: args{
				member: &Member{
					Username: "player",
				},
			},
			expectHost: &Member{
				Username: "admin",
			},
			wantErr: fmt.Errorf("Cannot Set Player As Host"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := NewGame("狼王守衛", &Member{
				Username: "admin",
			})
			assert.Nil(t, err)
			err = game.JoinRoom(&Member{
				Username: "user",
			})
			assert.Nil(t, err)
			err = game.JoinRoom(&Member{
				Username: "player",
			})
			assert.Nil(t, err)
			err = game.SetSeat(&Member{
				Username: "player",
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
					Username: "user",
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
					Username: "user",
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
					Username: "user",
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
					Username: "non-user",
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
					Username: "admin",
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
					Username: "user",
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
					Username: "user",
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
				Username: "admin",
			})
			assert.Nil(t, err)
			if tt.originPlayer != nil {
				game.Player = tt.originPlayer
			}
			err = game.JoinRoom(&Member{
				Username: "user",
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
				Username: "admin",
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

func TestGame_IsHost(t *testing.T) {
	type args struct {
		member *Member
	}
	tests := []struct {
		name      string
		args      args
		expectRes bool
	}{
		{
			name: "Test Normal",
			args: args{
				member: &Member{
					Username: "admin",
				},
			},
			expectRes: true,
		},
		{
			name: "Test Not Host",
			args: args{
				member: &Member{
					Username: "player",
				},
			},
			expectRes: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := NewGame("狼王守衛", &Member{
				Username: "admin",
			})
			assert.Nil(t, err)

			res := game.IsHost(tt.args.member)
			assert.Equal(t, tt.expectRes, res)
		})
	}
}

func TestGame_GetSeat(t *testing.T) {
	type args struct {
		member *Member
	}
	tests := []struct {
		name         string
		args         args
		originPlayer []*Player
		expectRes    int
		wantErr      error
	}{
		{
			name: "Test Normal 1",
			args: args{
				member: &Member{
					Username: "admin",
				},
			},
			originPlayer: []*Player{{
				Name: "admin",
			}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
			expectRes: 0,
		},
		{
			name: "Test Normal 2",
			args: args{
				member: &Member{
					Username: "admin",
				},
			},
			originPlayer: []*Player{nil, nil, nil, nil, nil, {
				Name: "admin",
			}, nil, nil, nil, nil, nil, nil},
			expectRes: 5,
		},
		{
			name: "Test Not InSeat",
			args: args{
				member: &Member{
					Username: "admin",
				},
			},
			expectRes: -1,
			wantErr:   fmt.Errorf("Not In Seat"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := NewGame("狼王守衛", &Member{
				Username: "admin",
			})
			assert.Nil(t, err)
			if tt.originPlayer != nil {
				game.Player = tt.originPlayer
			}

			res, err := game.GetSeat(tt.args.member)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.expectRes, res)
		})
	}
}

func TestGame_FillTestUser(t *testing.T) {
	game, err := NewGame("狼王守衛", &Member{
		Username: "admin",
	})
	assert.Nil(t, err)
	err = game.JoinRoom(&Member{
		Username: "user",
	})
	assert.Nil(t, err)
	err = game.SetSeat(&Member{
		Username: "user",
	}, 0)
	assert.Nil(t, err)
	game.FillTestUser()
	assert.Equal(t, &Player{Name: "user"}, game.Player[0])
	for _, player := range game.Player {
		assert.NotNil(t, player)
	}

}

func TestGame_Kill(t *testing.T) {
	game, err := NewGame("狼王守衛", &Member{
		Username: "admin",
	})
	assert.Nil(t, err)
	game.FillTestUser()
	err = game.Start()
	assert.Nil(t, err)

	assert.True(t, game.Player[0].Alive)
	game.Kill(0)
	assert.False(t, game.Player[0].Alive)
}
