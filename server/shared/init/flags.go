package init

import "flag"

type Flags struct {
	Kind    string
	Address string
}

func LoadFlags() *Flags {
	flags := &Flags{}

	flag.StringVar(&flags.Kind, "kind", "", "Kind of micro service")
	flag.StringVar(&flags.Address, "address", "", "Address of application")

	flag.Parse()

	return flags
}
