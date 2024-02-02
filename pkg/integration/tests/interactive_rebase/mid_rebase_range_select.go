package interactive_rebase

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var MidRebaseRangeSelect = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Do various things with range selection in the commits view when mid-rebase",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.
			CreateNCommits(10)
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Commits().
			Focus().
			TopLines(
				Contains("commit 10").IsSelected(),
			).
			NavigateToLine(Contains("commit 07")).
			Press(keys.Universal.RangeSelectDown).
			TopLines(
				Contains("commit 10"),
				Contains("commit 09"),
				Contains("commit 08"),
				Contains("commit 07").IsSelected(),
				Contains("commit 06").IsSelected(),
				Contains("commit 05"),
				Contains("commit 04"),
			).
			// Verify we can't perform an edit on multiple commits (it's not supported
			// yet)
			Press(keys.Universal.Edit).
			Tap(func() {
				// This ought to be a toast but I'm too lazy to implement that right now.
				t.ExpectPopup().Alert().
					Title(Equals("Error")).
					Content(Contains("Action does not support range selection, please select a single item")).
					Confirm()
			}).
			NavigateToLine(Contains("commit 05")).
			// Start a rebase
			Press(keys.Universal.Edit).
			TopLines(
				Contains("pick").Contains("commit 10"),
				Contains("pick").Contains("commit 09"),
				Contains("pick").Contains("commit 08"),
				Contains("pick").Contains("commit 07"),
				Contains("pick").Contains("commit 06"),
				Contains("<-- YOU ARE HERE --- commit 05").IsSelected(),
				Contains("commit 04"),
			).
			SelectPreviousItem().
			// perform various actions on a range of commits
			Press(keys.Universal.RangeSelectUp).
			TopLines(
				Contains("pick").Contains("commit 10"),
				Contains("pick").Contains("commit 09"),
				Contains("pick").Contains("commit 08"),
				Contains("pick").Contains("commit 07").IsSelected(),
				Contains("pick").Contains("commit 06").IsSelected(),
				Contains("<-- YOU ARE HERE --- commit 05"),
				Contains("commit 04"),
			).
			Press(keys.Commits.Fixup).
			TopLines(
				Contains("pick").Contains("commit 10"),
				Contains("pick").Contains("commit 09"),
				Contains("pick").Contains("commit 08"),
				Contains("fixup").Contains("commit 07").IsSelected(),
				Contains("fixup").Contains("commit 06").IsSelected(),
				Contains("<-- YOU ARE HERE --- commit 05"),
				Contains("commit 04"),
			).
			Press(keys.Commits.Pick).
			TopLines(
				Contains("pick").Contains("commit 10"),
				Contains("pick").Contains("commit 09"),
				Contains("pick").Contains("commit 08"),
				Contains("pick").Contains("commit 07").IsSelected(),
				Contains("pick").Contains("commit 06").IsSelected(),
				Contains("<-- YOU ARE HERE --- commit 05"),
				Contains("commit 04"),
			).
			Press(keys.Universal.Edit).
			TopLines(
				Contains("pick").Contains("commit 10"),
				Contains("pick").Contains("commit 09"),
				Contains("pick").Contains("commit 08"),
				Contains("edit").Contains("commit 07").IsSelected(),
				Contains("edit").Contains("commit 06").IsSelected(),
				Contains("<-- YOU ARE HERE --- commit 05"),
				Contains("commit 04"),
			).
			Press(keys.Commits.Squash).
			TopLines(
				Contains("pick").Contains("commit 10"),
				Contains("pick").Contains("commit 09"),
				Contains("pick").Contains("commit 08"),
				Contains("squash").Contains("commit 07").IsSelected(),
				Contains("squash").Contains("commit 06").IsSelected(),
				Contains("<-- YOU ARE HERE --- commit 05"),
				Contains("commit 04"),
			).
			Press(keys.Commits.MoveDownCommit).
			TopLines(
				Contains("pick").Contains("commit 10"),
				Contains("pick").Contains("commit 09"),
				Contains("pick").Contains("commit 08"),
				Contains("squash").Contains("commit 07").IsSelected(),
				Contains("squash").Contains("commit 06").IsSelected(),
				Contains("<-- YOU ARE HERE --- commit 05"),
				Contains("commit 04"),
			).
			Tap(func() {
				t.ExpectToast(Contains("Disabled: Cannot move any further"))
			}).
			Press(keys.Commits.MoveUpCommit).
			TopLines(
				Contains("pick").Contains("commit 10"),
				Contains("pick").Contains("commit 09"),
				Contains("squash").Contains("commit 07").IsSelected(),
				Contains("squash").Contains("commit 06").IsSelected(),
				Contains("pick").Contains("commit 08"),
				Contains("<-- YOU ARE HERE --- commit 05"),
				Contains("commit 04"),
			).
			Press(keys.Commits.MoveUpCommit).
			TopLines(
				Contains("pick").Contains("commit 10"),
				Contains("squash").Contains("commit 07").IsSelected(),
				Contains("squash").Contains("commit 06").IsSelected(),
				Contains("pick").Contains("commit 09"),
				Contains("pick").Contains("commit 08"),
				Contains("<-- YOU ARE HERE --- commit 05"),
				Contains("commit 04"),
			).
			Press(keys.Commits.MoveUpCommit).
			TopLines(
				Contains("squash").Contains("commit 07").IsSelected(),
				Contains("squash").Contains("commit 06").IsSelected(),
				Contains("pick").Contains("commit 10"),
				Contains("pick").Contains("commit 09"),
				Contains("pick").Contains("commit 08"),
				Contains("<-- YOU ARE HERE --- commit 05"),
				Contains("commit 04"),
			).
			Press(keys.Commits.MoveUpCommit).
			Tap(func() {
				t.ExpectToast(Contains("Disabled: Cannot move any further"))
			}).
			TopLines(
				Contains("squash").Contains("commit 07").IsSelected(),
				Contains("squash").Contains("commit 06").IsSelected(),
				Contains("pick").Contains("commit 10"),
				Contains("pick").Contains("commit 09"),
				Contains("pick").Contains("commit 08"),
				Contains("<-- YOU ARE HERE --- commit 05"),
				Contains("commit 04"),
			).
			// Verify we can't perform an action on a range that includes both
			// TODO and non-TODO commits
			NavigateToLine(Contains("commit 08")).
			Press(keys.Universal.RangeSelectDown).
			TopLines(
				Contains("squash").Contains("commit 07"),
				Contains("squash").Contains("commit 06"),
				Contains("pick").Contains("commit 10"),
				Contains("pick").Contains("commit 09"),
				Contains("pick").Contains("commit 08").IsSelected(),
				Contains("<-- YOU ARE HERE --- commit 05").IsSelected(),
				Contains("commit 04"),
			).
			Press(keys.Commits.Fixup).
			Tap(func() {
				t.ExpectToast(Contains("Disabled: When rebasing, this action only works on a selection of TODO commits."))
			}).
			TopLines(
				Contains("squash").Contains("commit 07"),
				Contains("squash").Contains("commit 06"),
				Contains("pick").Contains("commit 10"),
				Contains("pick").Contains("commit 09"),
				Contains("pick").Contains("commit 08").IsSelected(),
				Contains("<-- YOU ARE HERE --- commit 05").IsSelected(),
				Contains("commit 04"),
			).
			// continue the rebase
			Tap(func() {
				t.Common().ContinueRebase()
			}).
			TopLines(
				Contains("commit 10"),
				Contains("commit 09"),
				Contains("commit 08"),
				Contains("commit 05"),
				// selected indexes are retained, though we may want to clear it
				// in future (not sure what the best behaviour is right now)
				Contains("commit 04").IsSelected(),
				Contains("commit 03").IsSelected(),
			)
	},
})
