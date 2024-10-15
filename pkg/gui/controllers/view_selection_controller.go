package controllers

import (
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/tasks"
	"github.com/samber/lo"
)

type ViewSelectionControllerFactory struct {
	c                    *ControllerCommon
	viewBufferManagerMap *map[string]*tasks.ViewBufferManager
}

func NewViewSelectionControllerFactory(c *ControllerCommon, viewBufferManagerMap *map[string]*tasks.ViewBufferManager) *ViewSelectionControllerFactory {
	return &ViewSelectionControllerFactory{
		c:                    c,
		viewBufferManagerMap: viewBufferManagerMap,
	}
}

func (self *ViewSelectionControllerFactory) Create(context types.Context) types.IController {
	return &ViewSelectionController{
		baseController:       baseController{},
		c:                    self.c,
		context:              context,
		viewBufferManagerMap: self.viewBufferManagerMap,
	}
}

type ViewSelectionController struct {
	baseController
	c *ControllerCommon

	context              types.Context
	viewBufferManagerMap *map[string]*tasks.ViewBufferManager
}

func (self *ViewSelectionController) Context() types.Context {
	return self.context
}

func (self *ViewSelectionController) GetKeybindings(opts types.KeybindingsOpts) []*types.Binding {
	return []*types.Binding{
		{Tag: "navigation", Key: opts.GetKey(opts.Config.Universal.PrevItem), Handler: self.handlePrevLine},
		{Tag: "navigation", Key: opts.GetKey(opts.Config.Universal.PrevItemAlt), Handler: self.handlePrevLine},
		{Tag: "navigation", Key: opts.GetKey(opts.Config.Universal.NextItem), Handler: self.handleNextLine},
		{Tag: "navigation", Key: opts.GetKey(opts.Config.Universal.NextItemAlt), Handler: self.handleNextLine},
		{Tag: "navigation", Key: opts.GetKey(opts.Config.Universal.PrevPage), Handler: self.handlePrevPage, Description: self.c.Tr.PrevPage},
		{Tag: "navigation", Key: opts.GetKey(opts.Config.Universal.NextPage), Handler: self.handleNextPage, Description: self.c.Tr.NextPage},
		{Tag: "navigation", Key: opts.GetKey(opts.Config.Universal.GotoTop), Handler: self.handleGotoTop, Description: self.c.Tr.GotoTop},
		{Tag: "navigation", Key: opts.GetKey(opts.Config.Universal.GotoBottom), Description: self.c.Tr.GotoBottom, Handler: self.handleGotoBottom},
	}
}

func (self *ViewSelectionController) GetMouseKeybindings(opts types.KeybindingsOpts) []*gocui.ViewMouseBinding {
	return []*gocui.ViewMouseBinding{}
}

func (self *ViewSelectionController) handleLineChange(delta int) {
	if delta > 0 {
		if manager, ok := (*self.viewBufferManagerMap)[self.context.GetViewName()]; ok {
			manager.ReadLines(delta)
		}
	}

	v := self.Context().GetView()
	lineIdxBefore := v.CursorY() + v.OriginY()
	lineIdxAfter := lo.Clamp(lineIdxBefore+delta, 0, v.LinesHeight()-1)
	if delta == -1 {
		checkScrollUp(self.Context().GetViewTrait(), self.c.UserConfig(), lineIdxBefore, lineIdxAfter)
	} else if delta == 1 {
		checkScrollDown(self.Context().GetViewTrait(), self.c.UserConfig(), lineIdxBefore, lineIdxAfter)
	}
	v.FocusPoint(0, lineIdxAfter)
}

func (self *ViewSelectionController) handlePrevLine() error {
	self.handleLineChange(-1)
	return nil
}

func (self *ViewSelectionController) handleNextLine() error {
	self.handleLineChange(1)
	return nil
}

func (self *ViewSelectionController) handlePrevPage() error {
	self.handleLineChange(-self.context.GetViewTrait().PageDelta())
	return nil
}

func (self *ViewSelectionController) handleNextPage() error {
	self.handleLineChange(self.context.GetViewTrait().PageDelta())
	return nil
}

func (self *ViewSelectionController) handleGotoTop() error {
	v := self.Context().GetView()
	v.FocusPoint(0, 0)
	return nil
}

func (self *ViewSelectionController) handleGotoBottom() error {
	v := self.Context().GetView()
	v.FocusPoint(0, v.LinesHeight()-1)
	return nil
}
