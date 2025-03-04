package view

import (
	"github.com/derailed/tcell/v2"
	"github.com/one2nc/cloudlens/internal/ui"
)

type iamUserPloicy struct {
	ResourceViewer
}

func NewIamUserPloicy(resource string) ResourceViewer {
	var up iamUserPloicy
	up.ResourceViewer = NewBrowser(resource)
	up.AddBindKeysFn(up.bindKeys)
	return &up
}

func (up *iamUserPloicy) bindKeys(aa ui.KeyActions) {
	aa.Add(ui.KeyActions{
		tcell.KeyEscape: ui.NewKeyAction("Back", up.App().PrevCmd, true),
		ui.KeyShiftA:    ui.NewKeyAction("Policy-ARN", up.GetTable().SortColCmd("Policy-ARN", true), true),
		ui.KeyShiftN:    ui.NewKeyAction("Policy-Name", up.GetTable().SortColCmd("Policy-Name", true), true),
	})
}
