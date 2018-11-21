package init

import "flag"

type Flags struct {
	Kind    string
	EnvFile string
	Address string
}

func LoadFlags() *Flags {
	flags := &Flags{}

	flag.StringVar(&flags.Kind, "kind", "", "Kind of micro service")
	flag.StringVar(&flags.EnvFile, "env", "", "Environment file")
	flag.StringVar(&flags.Address, "address", ":8080", "Web API address")

	flag.Parse()

	return flags
}
