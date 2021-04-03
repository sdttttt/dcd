package huck

var kernel *HuckKernel = NewDCDKernel()

type HuckKernel struct {
	stat map[string]DataStatistical
}

func NewDCDKernel() *HuckKernel {
	return &HuckKernel{
		stat: make(map[string]DataStatistical),
	}
}
