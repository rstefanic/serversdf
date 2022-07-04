package output

import "fmt"

type PipeOutput struct {
	LabelPadding int
}

func NewPipeOutput(longestServerName int) *PipeOutput {
	return &PipeOutput{longestServerName}
}

func (po PipeOutput) WriteInfo(label, status string, index int) {
	po.writeLineWithStatusPrefix(label, status, "INFO")
}

func (po PipeOutput) WriteError(label, status string, index int) {
	po.writeLineWithStatusPrefix(label, status, "FAILURE")
}

func (po PipeOutput) WriteSuccess(label, status string, index int) {
	po.writeLineWithStatusPrefix(label, status, "SUCCESS")
}

func (po PipeOutput) writeLineWithStatusPrefix(label, status, prefix string) {
	fmt.Println(FormatLine(label, prefix+": "+status, po.LabelPadding))
}
