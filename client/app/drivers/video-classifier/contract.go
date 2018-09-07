package video_classifier

import "trans/client/app/contracts"

type Classifier interface {
	contracts.EventProducer
	contracts.Runnable
}
