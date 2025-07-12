package branch

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var CheckoutLastBranch = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Checkout to the last branch using the checkout last branch functionality",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.
			CreateNCommits(3).
			NewBranch("last-branch").
			EmptyCommit("last commit").
			Checkout("master").
			EmptyCommit("master commit")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Branches().
			Focus().
			Lines(
				Contains("master").IsSelected(),
				Contains("last-branch"),
			)

		// Press the checkout last branch key (should checkout last-branch)
		t.Views().Branches().
			Press(keys.Branches.CheckoutLastBranch).
			Lines(
				Contains("last-branch").IsSelected(),
				Contains("master"),
			)

		// Verify we're on last-branch
		t.Git().CurrentBranchName("last-branch")

		// Press again to go back to master
		t.Views().Branches().
			Press(keys.Branches.CheckoutLastBranch).
			Lines(
				Contains("master").IsSelected(),
				Contains("last-branch"),
			)

		// Verify we're back on master
		t.Git().CurrentBranchName("master")
	},
})
