package ignitehello_test

import (
	"testing"

	keepertest "github.com/albacg5/ignite-hello/testutil/keeper"
	"github.com/albacg5/ignite-hello/testutil/nullify"
	"github.com/albacg5/ignite-hello/x/ignitehello"
	"github.com/albacg5/ignite-hello/x/ignitehello/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IgnitehelloKeeper(t)
	ignitehello.InitGenesis(ctx, *k, genesisState)
	got := ignitehello.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
