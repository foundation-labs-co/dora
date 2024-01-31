package config

import (
	"embed"
)

// explorer config
//
//go:embed default.config.yml
var DefaultConfigYml string

// chain presets
//
//go:embed preset-mainnet.chain.yml
var MainnetPresetYml string

//go:embed preset-minimal.chain.yml
var MinimalPresetYml string

//go:embed preset-gnosis.chain.yml
var GnosisPresetYml string

// chain configs
//
//go:embed mainnet.chain.yml
var MainnetChainYml string

//go:embed prater.chain.yml
var PraterChainYml string

//go:embed sepolia.chain.yml
var SepoliaChainYml string

//go:embed holesky.chain.yml
var HoleskyChainYml string

//go:embed gnosis.chain.yml
var GnosisChainYml string

//go:embed devnet.chain.yml
var DevnetChainYml string

// validator names
//
//go:embed *.names.yml
var ValidatorNamesYml embed.FS
