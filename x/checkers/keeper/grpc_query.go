package keeper

import (
	"github.com/dotneko/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
