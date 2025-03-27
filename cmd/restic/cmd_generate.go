package main

import (
	"io"
	"os"
	"time"

	"github.com/restic/restic/api/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/pflag"
)

func newGenerateCommand() *cobra.Command {
	var opts generateOptions

	cmd := &cobra.Command{
		Use:   "generate [flags]",
		Short: "Generate manual pages and auto-completion files (bash, fish, zsh, powershell)",
		Long: `
The "generate" command writes automatically generated files (like the man pages
and the auto-completion files for bash, fish and zsh).

EXIT STATUS
===========

Exit status is 0 if the command was successful.
Exit status is 1 if there was any error.
`,
		DisableAutoGenTag: true,
		RunE: func(_ *cobra.Command, args []string) error {
			return runGenerate(opts, args)
		},
	}
	opts.AddFlags(cmd.Flags())
	return cmd
}

type generateOptions struct {
	ManDir                   string
	BashCompletionFile       string
	FishCompletionFile       string
	ZSHCompletionFile        string
	PowerShellCompletionFile string
}

func (opts *generateOptions) AddFlags(f *pflag.FlagSet) {
	f.StringVar(&opts.ManDir, "man", "", "write man pages to `directory`")
	f.StringVar(&opts.BashCompletionFile, "bash-completion", "", "write bash completion `file` (`-` for stdout)")
	f.StringVar(&opts.FishCompletionFile, "fish-completion", "", "write fish completion `file` (`-` for stdout)")
	f.StringVar(&opts.ZSHCompletionFile, "zsh-completion", "", "write zsh completion `file` (`-` for stdout)")
	f.StringVar(&opts.PowerShellCompletionFile, "powershell-completion", "", "write powershell completion `file` (`-` for stdout)")
}

func writeManpages(root *cobra.Command, dir string) error {
	// use a fixed date for the man pages so that generating them is deterministic
	date, err := time.Parse("Jan 2006", "Jan 2017")
	if err != nil {
		return err
	}

	header := &doc.GenManHeader{
		Title:   "restic backup",
		Section: "1",
		Source:  "generated by `restic generate`",
		Date:    &date,
	}

	Verbosef("writing man pages to directory %v\n", dir)
	return doc.GenManTree(root, header, dir)
}

func writeCompletion(filename string, shell string, generate func(w io.Writer) error) (err error) {
	if stdoutIsTerminal() {
		Verbosef("writing %s completion file to %v\n", shell, filename)
	}
	var outWriter io.Writer
	if filename != "-" {
		var outFile *os.File
		outFile, err = os.Create(filename)
		if err != nil {
			return
		}
		defer func() { err = outFile.Close() }()
		outWriter = outFile
	} else {
		outWriter = globalOptions.stdout
	}

	err = generate(outWriter)
	return
}

func checkStdoutForSingleShell(opts generateOptions) error {
	completionFileOpts := []string{
		opts.BashCompletionFile,
		opts.FishCompletionFile,
		opts.ZSHCompletionFile,
		opts.PowerShellCompletionFile,
	}
	seenIsStdout := false
	for _, completionFileOpt := range completionFileOpts {
		if completionFileOpt == "-" {
			if seenIsStdout {
				return errors.Fatal("the generate command can generate shell completions to stdout for single shell only")
			}
			seenIsStdout = true
		}
	}
	return nil
}

func runGenerate(opts generateOptions, args []string) error {
	if len(args) > 0 {
		return errors.Fatal("the generate command expects no arguments, only options - please see `restic help generate` for usage and flags")
	}

	cmdRoot := newRootCommand()

	if opts.ManDir != "" {
		err := writeManpages(cmdRoot, opts.ManDir)
		if err != nil {
			return err
		}
	}

	err := checkStdoutForSingleShell(opts)
	if err != nil {
		return err
	}

	if opts.BashCompletionFile != "" {
		err := writeCompletion(opts.BashCompletionFile, "bash", cmdRoot.GenBashCompletion)
		if err != nil {
			return err
		}
	}

	if opts.FishCompletionFile != "" {
		err := writeCompletion(opts.FishCompletionFile, "fish", func(w io.Writer) error { return cmdRoot.GenFishCompletion(w, true) })
		if err != nil {
			return err
		}
	}

	if opts.ZSHCompletionFile != "" {
		err := writeCompletion(opts.ZSHCompletionFile, "zsh", cmdRoot.GenZshCompletion)
		if err != nil {
			return err
		}
	}

	if opts.PowerShellCompletionFile != "" {
		err := writeCompletion(opts.PowerShellCompletionFile, "powershell", cmdRoot.GenPowerShellCompletion)
		if err != nil {
			return err
		}
	}

	var empty generateOptions
	if opts == empty {
		return errors.Fatal("nothing to do, please specify at least one output file/dir")
	}

	return nil
}
