package dcd

var kernel *DCDKernel = NewDCDKernel()

type DCDKernel struct {
	stat map[string]DataStatistical
}

func NewDCDKernel() *DCDKernel {
	return &DCDKernel{
		stat: make(map[string]DataStatistical),
	}
}
