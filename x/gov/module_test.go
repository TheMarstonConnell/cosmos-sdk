package gov_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/simapp"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	db := dbm.NewMemDB()

	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = simapp.DefaultNodeHome
	appOptions[server.FlagInvCheckPeriod] = 5

	app := simapp.NewSimApp(log.NewNopLogger(), db, nil, true, appOptions)

	genesisState := simapp.GenesisStateWithSingleValidator(t, app)

	stateBytes, err := sdk.MarshalGenesisStateToJSON(genesisState, app.AppCodec(), false)
	require.NoError(t, err)

	app.InitChain(
		abcitypes.RequestInitChain{
			AppStateBytes: stateBytes,
			ChainId:       "test-chain-id",
		},
	)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	acc := app.AccountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	require.NotNil(t, acc)
}
