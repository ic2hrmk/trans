package haar

import "trans/client/app/contracts"

type Classifier interface {
	contracts.EventProducer
	contracts.Runnable
}
