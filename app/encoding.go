package app

import (
	"github.com/cosmos/cosmos-sdk/std"
	appparams "github.com/dotneko/checkers/app/params"
	"github.com/ignite-hq/cli/ignite/pkg/cosmoscmd"
)

// MakeTestEncodingConfig creates an EncodingConfig for testing. This function
// should be used only in tests or when creating a new app instance (New*()).
// App user shouldn't create new codecs - use the app.AppCodec instead.
// [DEPRECATED]
func MakeTestEncodingConfig() cosmoscmd.EncodingConfig {
	encodingConfig := appparams.MakeTestEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
