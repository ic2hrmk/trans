package reporter

import (
	"trans/client/app/contracts"
)

type Reporter interface {
	contracts.EventListener
	contracts.Runnable
}
