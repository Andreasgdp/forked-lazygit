package i18n

func polishTranslationSet() TranslationSet {
	return TranslationSet{
		NotEnoughSpace:                      "Za mało miejsca do wyświetlenia paneli",
		DiffTitle:                           "Różnice",
		LogTitle:                            "Log",
		FilesTitle:                          "Pliki",
		BranchesTitle:                       "Gałęzie",
		CommitsTitle:                        "Commity",
		StashTitle:                          "Schowek",
		UnstagedChanges:                     `Unstaged Changes`,
		StagedChanges:                       `Staged Changes`,
		MergingMainTitle:                    "Resolve merge conflicts",
		CommitMessage:                       "Wiadomość commita",
		CredentialsUsername:                 "Username",
		CredentialsPassword:                 "Password",
		CredentialsPassphrase:               "Passphrase",
		PassUnameWrong:                      "Password, passphrase and/or username wrong",
		CommitChanges:                       "commituj zmiany",
		AmendLastCommit:                     "zmień ostatnie zatwierdzenie",
		SureToAmend:                         "Czy na pewno chcesz zmienić ostatnie zatwierdzenie? Możesz zmienić komunikat zatwierdzenia z panelu zatwierdzeń.",
		NoCommitToAmend:                     "Nie ma zobowiązania do zmiany.",
		CommitChangesWithEditor:             "commituj zmiany używając edytora z gita",
		StatusTitle:                         "Status",
		GlobalTitle:                         "Globalne",
		LcNavigate:                          "nawiguj",
		LcMenu:                              "menu",
		LcExecute:                           "wykonaj",
		LcOpen:                              "otwórz",
		LcIgnore:                            "ignoruj",
		LcDelete:                            "usuń",
		LcToggleStaged:                      "przełącz zatwierdzenie",
		LcToggleStagedAll:                   "przełącz wszystkie zatwierdzenia",
		LcRefresh:                           "odśwież",
		LcEdit:                              "edytuj",
		LcScroll:                            "przewiń",
		LcAbortMerge:                        "o scalaniu",
		LcResolveMergeConflicts:             "rozwiąż konflikty scalania",
		LcCheckout:                          "przełącz",
		NoChangedFiles:                      "Brak zmienionych plików",
		FileHasNoUnstagedChanges:            "Plik nie zawiera żadnych nieopublikowanych zmian do dodania",
		CannotGitAdd:                        "Nie można git add --patch nieśledzonych plików",
		NoFilesDisplay:                      "Brak pliku do wyświetlenia",
		PullWait:                            "Wciąganie zmian...",
		PushWait:                            "Wypychanie zmian...",
		FetchWait:                           "Fetching...",
		FileNoMergeCons:                     "Ten plik nie powoduje konfliktów scalania",
		SureTo:                              "Jesteś pewny, że chcesz {{.deleteVerb}} {{.fileName}} (stracisz swoje wprowadzone zmiany)?",
		AlreadyCheckedOutBranch:             "Już przęłączono na tą gałąź",
		SureForceCheckout:                   "Jesteś pewny, że chcesz wymusić przełączenie? Stracisz wszystkie lokalne zmiany",
		ForceCheckoutBranch:                 "Wymuś przełączenie gałęzi",
		BranchName:                          "Nazwa gałęzi",
		NewBranchNameBranchOff:              "Nazwa nowej gałęzi (gałąź na bazie {{.branchName}})",
		CantDeleteCheckOutBranch:            "Nie możesz usunąć obecnej przełączonej gałęzi!",
		DeleteBranch:                        "Usuń gałąź",
		DeleteBranchMessage:                 "Jesteś pewien, że chcesz usunąć gałąź {{.selectedBranchName}} ?",
		ForceDeleteBranchMessage:            "Na pewno wymusić usunięcie gałęzi {{.selectedBranchName}}?",
		LcRebaseBranch:                      "rebase branch",
		CantRebaseOntoSelf:                  "You cannot rebase a branch onto itself",
		CantMergeBranchIntoItself:           "Nie możesz scalić gałęzi do samej siebie",
		LcForceCheckout:                     "wymuś przełączenie",
		LcMerge:                             "scal",
		LcCheckoutByName:                    "przełącz używając nazwy",
		LcNewBranch:                         "nowa gałąź",
		LcDeleteBranch:                      "usuń gałąź",
		LcForceDeleteBranch:                 "usuń gałąź (wymuś)",
		NoBranchesThisRepo:                  "Brak gałęzi dla tego repozytorium",
		NoTrackingThisBranch:                "Brak śledzenia dla tej gałęzi",
		CommitMessageConfirm:                "{{.keyBindClose}}: zamknij, {{.keyBindNewLine}}: new line, {{.keyBindConfirm}}: potwierdź",
		CommitWithoutMessageErr:             "Nie możesz commitować bez podania wiadomości",
		CloseConfirm:                        "{{.keyBindClose}}: zamknij, {{.keyBindConfirm}}: potwierdź",
		LcClose:                             "zamknij",
		SureResetThisCommit:                 "Jesteś pewny, że chcesz zresetować ten commit?",
		ResetToCommit:                       "Zresetuj, aby commitować",
		LcSquashDown:                        "ściśnij w dół",
		LcRename:                            "przemianuj",
		LcResetToThisCommit:                 "zresetuj do tego commita",
		LcFixupCommit:                       "napraw commit",
		NoCommitsThisBranch:                 "Brak commitów dla tej gałęzi",
		OnlySquashTopmostCommit:             "Można tylko ścisnąć najwyższy commit",
		YouNoCommitsToSquash:                "Nie masz commitów do ściśnięcia",
		CantFixupWhileUnstagedChanges:       "Nie można wykonać naprawy, kiedy istnieją niezatwierdzone zmiany",
		Fixup:                               "Napraw",
		SureFixupThisCommit:                 "Jesteś pewny, ze chcesz naprawić ten commit? Commit poniżej zostanie ściśnięty w górę wraz z tym",
		OnlyRenameTopCommit:                 "Można przmianować tylko najwyższy commit",
		LcRenameCommit:                      "przemianuj commit",
		LcRenameCommitEditor:                "przemianuj commit w edytorze",
		PotentialErrInGetselectedCommit:     "potencjalny błąd w getSelected Commit (niedopasowane ui i stan)",
		Error:                               "Błąd",
		RunningSubprocess:                   "uruchomiony podproces",
		LcSelectHunk:                        "wybierz kawałek",
		LcNavigateConflicts:                 "nawiguj konflikty",
		LcPickHunk:                          "wybierz kawałek",
		LcPickBothHunks:                     "wybierz oba kawałki",
		LcUndo:                              "cofnij",
		LcPop:                               "wyciągnij",
		LcDrop:                              "porzuć",
		LcApply:                             "zastosuj",
		NoStashEntries:                      "Brak pozycji w schowku",
		StashDrop:                           "Porzuć schowek",
		SureDropStashEntry:                  "Jesteś pewny, że chcesz porzucić tę pozycję w schowku?",
		NoStashTo:                           "Brak schowka dla {{.method}}",
		NoTrackedStagedFilesStash:           "Nie masz śledzonych/zatwierdzonych plików do przechowania",
		StashChanges:                        "Przechowaj zmiany",
		IssntListOfViews:                    "{{.name}} nie jest na liście widoków",
		LcNewFocusedViewIs:                  "nowy skupiony widok to {{.newFocusedView}}",
		MergeAborted:                        "Scalanie anulowane",
		OpenConfig:                          "otwórz plik konfiguracyjny",
		EditConfig:                          "edytuj plik konfiguracyjny",
		ForcePush:                           "Wymuś wypchnięcie",
		ForcePushPrompt:                     "Twoja gałąź rozeszła się z gałęzią zdalną. Wciśnij 'esc' aby anulować lub 'enter' aby wymusić wypchnięcie.",
		ForcePushDisabled:                   "Your branch has diverged from the remote branch and you've disabled force pushing",
		UpdatesRejectedAndForcePushDisabled: "Updates were rejected and you have disabled force pushing",
		LcCheckForUpdate:                    "sprawdź aktualizacje",
		CheckingForUpdates:                  "Sprawdzanie aktualizacji...",
		OnLatestVersionErr:                  "Już posiadasz najnowszą wersję",
		MajorVersionErr:                     "Nowa wersja ({{.newVersion}}) posiada niekompatybilne zmiany w porównaniu do obecnej wersji ({{.currentVersion}})",
		CouldNotFindBinaryErr:               "Nie można znaleźć pliku binarnego w {{.url}}",
		AnonymousReportingTitle:             "Help make lazygit better",
		AnonymousReportingPrompt:            "Włączyć anonimowe raportowanie błędów w celu pomocy w usprawnianiu lazygita (enter/esc)?",
		LcEditFile:                          `edytuj plik`,
		LcOpenFile:                          `otwórz plik`,
		LcIgnoreFile:                        `dodaj do .gitignore`,
		LcRefreshFiles:                      `odśwież pliki`,
		LcMergeIntoCurrentBranch:            `scal do obecnej gałęzi`,
		ConfirmQuit:                         `Na pewno chcesz wyjść z programu?`,
		LcAllBranchesLogGraph:               `pokazywać wszystkie logi branżowe`,
		UnsupportedGitService:               `Nieobsługiwana usługa git`,
		LcCreateOrShowPullRequest:           `utwórz żądanie wyciągnięcia`,
		LcCopyPullRequestURL:                `skopiuj adres URL żądania ściągnięcia do schowka`,
		NoBranchOnRemote:                    `Ta gałąź nie istnieje na zdalnym. Najpierw musisz go odepchnąć na odległość.`,
		LcFetch:                             `fetch`,
		NoAutomaticGitFetchTitle:            `No automatic git fetch`,
		NoAutomaticGitFetchBody:             `Lazygit can't use "git fetch" in a private repo use f in the branches panel to run "git fetch" manually`,
		FileEnter:                           `zatwierdź pojedyncze linie`,
		FileStagingRequirements:             `Można tylko zatwierdzić pojedyncze linie dla śledzonych plików z niezatwierdzonymi zmianami`,
		StagingTitle:                        `Zatwierdzanie`,
		ReturnToFilesPanel:                  `wróć do panelu plików`,
		CantFindHunks:                       `Nie można znaleźć żadnych kawałków w tej łatce`,
		CantFindHunk:                        `Nie można znaleźć kawałka`,
		RebasingTitle:                       "Rebasing",
		MergingTitle:                        "Merging",
		ConfirmRebase:                       "Are you sure you want to rebase {{.checkedOutBranch}} onto {{.selectedBranch}}?",
		ConfirmMerge:                        "Are you sure you want to merge {{.selectedBranch}} into {{.checkedOutBranch}}?",
		FwdNoUpstream:                       "Cannot fast-forward a branch with no upstream",
		FwdCommitsToPush:                    "Cannot fast-forward a branch with commits to push",
		ErrorOccurred:                       "An error occurred! Please create an issue at",
		MainTitle:                           "Main",
		NormalTitle:                         "Normal",
		LcSoftReset:                         "soft reset",
		SureSquashThisCommit:                "Are you sure you want to squash this commit into the commit below?",
		Squash:                              "Squash",
		LcPickCommit:                        "pick commit (when mid-rebase)",
		LcRevertCommit:                      "revert commit",
		LcDeleteCommit:                      "delete commit",
		LcMoveDownCommit:                    "move commit down one",
		LcMoveUpCommit:                      "move commit up one",
		LcEditCommit:                        "edit commit",
		LcAmendToCommit:                     "amend commit with staged changes",
		FoundConflicts:                      "Conflicts! To abort press 'esc', otherwise press 'enter'",
		FoundConflictsTitle:                 "Auto-merge failed",
		Undo:                                "undo",
		PickHunk:                            "pick hunk",
		PickBothHunks:                       "pick both hunks",
		ViewMergeRebaseOptions:              "view merge/rebase options",
		NotMergingOrRebasing:                "You are currently neither rebasing nor merging",
		RecentRepos:                         "recent repositories",
		MergeOptionsTitle:                   "Merge Options",
		RebaseOptionsTitle:                  "Rebase Options",
		ConflictsResolved:                   "all merge conflicts resolved. Continue?",
		NoRoom:                              "Not enough room",
		YouAreHere:                          "YOU ARE HERE",
		LcRewordNotSupported:                "rewording commits while interactively rebasing is not currently supported",
		LcCherryPickCopy:                    "copy commit (cherry-pick)",
		LcCherryPickCopyRange:               "copy commit range (cherry-pick)",
		LcPasteCommits:                      "paste commits (cherry-pick)",
		SureCherryPick:                      "Are you sure you want to cherry-pick the copied commits onto this branch?",
		CherryPick:                          "Cherry-Pick",
		CannotRebaseOntoFirstCommit:         "You cannot interactive rebase onto the first commit",
		CannotSquashOntoSecondCommit:        "You cannot squash/fixup onto the second commit",
		Donate:                              "Donate",
		PrevLine:                            "select previous line",
		NextLine:                            "select next line",
		PrevHunk:                            "select previous hunk",
		NextHunk:                            "select next hunk",
		PrevConflict:                        "select previous conflict",
		NextConflict:                        "select next conflict",
		SelectTop:                           "select top hunk",
		SelectBottom:                        "select bottom hunk",
		ScrollDown:                          "scroll down",
		ScrollUp:                            "scroll up",
		AmendCommitTitle:                    "Amend Commit",
		AmendCommitPrompt:                   "Are you sure you want to amend this commit with your staged files?",
		DeleteCommitTitle:                   "Delete Commit",
		DeleteCommitPrompt:                  "Are you sure you want to delete this commit?",
		SquashingStatus:                     "squashing",
		FixingStatus:                        "fixing up",
		DeletingStatus:                      "deleting",
		MovingStatus:                        "moving",
		RebasingStatus:                      "rebasing",
		AmendingStatus:                      "amending",
		CherryPickingStatus:                 "cherry-picking",
		CommitFiles:                         "Commit files",
		LcViewCommitFiles:                   "view commit's files",
		CommitFilesTitle:                    "Commit files",
		LcGoBack:                            "go back",
		NoCommiteFiles:                      "No files for this commit",
		LcCheckoutCommitFile:                "checkout file",
		LcDiscardOldFileChange:              "discard this commit's changes to this file",
		DiscardFileChangesTitle:             "Discard file changes",
		DiscardFileChangesPrompt:            "Are you sure you want to discard this commit's changes to this file? If this file was created in this commit, it will be deleted",
		DisabledForGPG:                      "Feature not available for users using GPG",
		CreateRepo:                          "Not in a git repository. Create a new git repository? (y/n): ",
		AutoStashTitle:                      "Autostash?",
		AutoStashPrompt:                     "You must stash and pop your changes to bring them across. Do this automatically? (enter/esc)",
		StashPrefix:                         "Auto-stashing changes for ",
		LcViewDiscardOptions:                "view 'discard changes' options",
		LcCancel:                            "cancel",
		LcDiscardAllChanges:                 "discard all changes",
		LcDiscardUnstagedChanges:            "discard unstaged changes",
		LcDiscardAllChangesToAllFiles:       "nuke working tree",
		LcDiscardAnyUnstagedChanges:         "discard unstaged changes",
		LcDiscardUntrackedFiles:             "discard untracked files",
		LcHardReset:                         "hard reset",
		LcViewResetOptions:                  `view reset options`,
		LcCreateFixupCommit:                 `create fixup commit for this commit`,
		LcSquashAboveCommits:                `squash all 'fixup!' commits above selected commits (autosquash)`,
		SquashAboveCommits:                  `Squash all 'fixup!' commits above selected commits (autosquash)`,
		SureSquashAboveCommits:              `Are you sure you want to squash all fixup! commits above {{.commit}}?`,
		CreateFixupCommit:                   `Create fixup commit`,
		SureCreateFixupCommit:               `Are you sure you want to create a fixup! commit for commit {{.commit}}?`,
		LcExecuteCustomCommand:              "execute custom command",
		CustomCommand:                       "Custom Command:",
		LcCommitChangesWithoutHook:          "commit changes without pre-commit hook",
		SkipHookPrefixNotConfigured:         "You have not configured a commit message prefix for skipping hooks. Set `git.skipHookPrefix = 'WIP'` in your config",
		LcResetTo:                           `reset to`,
		PressEnterToReturn:                  "Press enter to return to lazygit",
		LcViewStashOptions:                  "view stash options",
		LcStashAllChanges:                   "przechowaj pliki",
		LcStashStagedChanges:                "stash staged changes",
		LcStashOptions:                      "Stash options",
		NotARepository:                      "Error: must be run inside a git repository",
		LcJump:                              "jump to panel",
		ExitLineByLineMode:                  `exit line-by-line mode`,
		EnterUpstream:                       `Enter upstream as '<remote> <branchname>'`,
		ReturnToRemotesList:                 `return to remotes list`,
		IgnoreTracked:                       "Ignore tracked file",
		IgnoreTrackedPrompt:                 "Are you sure you want to ignore a tracked file?",
		LcCommitPrefixPatternError:          "Error in commitPrefix pattern",
		NoFilesStagedTitle:                  "No files staged",
		NoFilesStagedPrompt:                 "You have not staged any files. Commit all files?",
		BranchNotFoundTitle:                 "Branch not found",
		BranchNotFoundPrompt:                "Branch not found. Create a new branch named",
		PullRequestURLCopiedToClipboard:     "URL żądania ściągnięcia skopiowany do schowka",
		CommitMessageCopiedToClipboard:      "Commit message skopiowany do schowka",
		LcCopiedToClipboard:                 "skopiowany do schowka",
		CreateOrOpenPullRequestOptions:      "Utwórz opcje żądania ściągnięcia",
		LcCreateOrOpenPullRequestOptions:    "utwórz opcje żądania ściągnięcia",
	}
}
