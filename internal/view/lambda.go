package view

import (
	"github.com/derailed/tcell/v2"
	"github.com/one2nc/cloudlens/internal/ui"
)

type Lambda struct {
	ResourceViewer
}

func NewLambda(resource string) ResourceViewer {
	var l Lambda
	l.ResourceViewer = NewBrowser(resource)
	l.AddBindKeysFn(l.bindKeys)
	return &l
}

func (l *Lambda) bindKeys(aa ui.KeyActions) {
	aa.Add(ui.KeyActions{
		ui.KeyShiftN:    ui.NewKeyAction("Sort Function-Name", l.GetTable().SortColCmd("Function-Name", true), true),
		ui.KeyShiftR:    ui.NewKeyAction("Sort Role", l.GetTable().SortColCmd("Role", true), true),
		ui.KeyShiftA:    ui.NewKeyAction("Sort Function-Arn", l.GetTable().SortColCmd("Function-Arn", true), true),
		ui.KeyShiftS:    ui.NewKeyAction("Sort Code-Size", l.GetTable().SortColCmd("Code-Size", true), true),
		ui.KeyShiftM:    ui.NewKeyAction("Sort Last-Modified", l.GetTable().SortColCmd("Last-Modified", true), true),
		tcell.KeyEscape: ui.NewKeyAction("Back", l.App().PrevCmd, false),
	})
}
