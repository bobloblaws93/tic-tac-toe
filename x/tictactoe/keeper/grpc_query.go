package keeper

import (
	"github.com/assignment/tictactoe/x/tictactoe/types"
)

var _ types.QueryServer = Keeper{}
