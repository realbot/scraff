package scraff

import (
	"github.com/golang/glog"
)

const processorVersion = "1.0.0"

type AdProcessor struct {
	Extractors []AdExtractor
	Store      AdStore
}

func (te AdProcessor) Run() int {
	glog.Infof("ad processor version %s\n", processorVersion)

	return ExitCodeOK
}
