package config

var basicConfig = map[string]string{
	// basic config...
}

var base = mergeConfig(
	basicConfig,
	httpConfig,
)

var baseInterface = mergeConfigInterface(
	httpInterfaceConfig,
)
