package filter_by_path

import (
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

func commonSetup(shell *Shell) {
	shell.CreateFileAndAdd("filterFile", "original filterFile content")
	shell.CreateFileAndAdd("otherFile", "original otherFile content")
	shell.Commit("both files")

	shell.UpdateFileAndAdd("otherFile", "new otherFile content")
	shell.Commit("only otherFile")

	shell.UpdateFileAndAdd("filterFile", "new filterFile content")
	shell.Commit("only filterFile")

	shell.EmptyCommit("none of the two")
}

func postFilterTest(t *TestDriver) {
	t.Views().Information().Content(Contains("Filtering by 'filterFile'"))

	t.Views().Commits().
		IsFocused().
		Lines(
			Contains(`only filterFile`).IsSelected(),
			Contains(`both files`),
		).
		SelectNextItem()

	// we only show the filtered file's changes in the main view
	t.Views().Main().
		Content(Contains("filterFile").DoesNotContain("otherFile"))

	t.Views().Commits().
		PressEnter()

	// when you click into the commit itself, you see all files from that commit
	t.Views().CommitFiles().
		IsFocused().
		Lines(
			Equals("▼ /"),
			Contains(`filterFile`),
			Contains(`otherFile`),
		)
}
