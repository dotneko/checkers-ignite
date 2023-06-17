package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	checkersapp "github.com/dotneko/checkers/app"
	"github.com/dotneko/checkers/x/checkers/testutil"
	"github.com/dotneko/checkers/x/checkers/types"
	"github.com/stretchr/testify/suite"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
	carol = testutil.Carol
)
const (
	balAlice = 50000000
	balBob   = 20000000
	balCarol = 10000000
)

type IntegrationTestSuite struct {
	suite.Suite

	app         *checkersapp.App
	msgServer   types.MsgServer
	ctx         sdk.Context
	queryClient types.QueryClient
}

var (
	checkersModuleAddress string
)

func TestCheckersKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
