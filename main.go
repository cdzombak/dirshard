package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

var version = "<dev>"

// Environment variables used to set default behavior.
const (
	NEnvVar    = "DIRSHARD_N"
	CIEnvVar   = "DIRSHARD_CI"
	SkipEnvVar = "DIRSHARD_SKIP"
)

// Flag names.
const (
	FlagN    = "n"
	FlagCi   = "ci"
	FlagSkip = "skip"
)

func main() {
	binName := filepath.Base(os.Args[0])

	n := flag.Int(FlagN, 1, fmt.Sprintf("Number of shards to produce. "+
		"Can be set by the %s environment variable; this flag overrides the env var.", NEnvVar))
	ci := flag.Bool(FlagCi, false, fmt.Sprintf("Case-insensitive: letters will be converted to lowercase. "+
		"Can be set by the %s environment variable; this flag overrides the env var.", CIEnvVar))
	skip := flag.Bool(FlagSkip, false, fmt.Sprintf("Skip non-alphanumeric characters entirely, rather than converting them to underscores. "+
		"Can be set by the %s environment variable; this flag overrides the env var.", SkipEnvVar))
	printVersion := flag.Bool("version", false, "Print version and exit.")
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage:\n      %s [OPTIONS] -- some_object_key\n (or) %s [OPTIONS] < cat object_list.txt\n", binName, binName)
		_, _ = fmt.Fprintf(os.Stderr, "\nProduce a path fragment consisting of the first N alphanumeric "+
			"characters of the given object key, separated by a path separator.\n")
		_, _ = fmt.Fprintf(os.Stderr, "No leading or trailing slash is produced.\n")
		_, _ = fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
		_, _ = fmt.Fprintf(os.Stderr, "\nVersion:\n  dirshard %s\n", version)
		_, _ = fmt.Fprintf(os.Stderr, "\nGitHub:\n  https://github.com/cdzombak/dirshard\n")
		_, _ = fmt.Fprintf(os.Stderr, "\nAuthor:\n  Chris Dzombak <https://www.dzombak.com>\n")
	}
	flag.Parse()

	if *printVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if !wasFlagGiven(FlagN) {
		nStrFromEnv := os.Getenv(NEnvVar)
		if nStrFromEnv != "" {
			nFromEnv, err := strconv.Atoi(nStrFromEnv)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dirshard: env var %s='%s' is invalid\n", NEnvVar, nStrFromEnv)
				os.Exit(1)
			}
			n = &nFromEnv
		}
	}

	if !wasFlagGiven(FlagCi) {
		ciStrFromEnv := os.Getenv(CIEnvVar)
		if ciStrFromEnv != "" {
			ciFromEnv, err := strToBool(ciStrFromEnv)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dirshard: env var %s='%s' is invalid\n", CIEnvVar, ciStrFromEnv)
				os.Exit(1)
			}
			ci = &ciFromEnv
		}
	}

	if !wasFlagGiven(FlagSkip) {
		skipStrFromEnv := os.Getenv(SkipEnvVar)
		if skipStrFromEnv != "" {
			skipFromEnv, err := strToBool(skipStrFromEnv)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dirshard: env var %s='%s' is invalid\n", SkipEnvVar, skipStrFromEnv)
				os.Exit(1)
			}
			skip = &skipFromEnv
		}
	}

	if flag.Arg(0) != "" {
		fmt.Print(dirshard(flag.Arg(0), *n, *ci, *skip))
		os.Exit(0)
	}

	didProcessData := false
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(dirshard(scanner.Text(), *n, *ci, *skip))
		didProcessData = true
	}

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "dirshard: error reading stdin: %s\n", err)
		os.Exit(1)
	}
	if !didProcessData {
		_, _ = fmt.Fprintf(os.Stderr, "dirshard: no input provided (usage: %s -- myobjectkey OR %s < filelist)\n", binName, binName)
		os.Exit(1)
	}
}

func dirshard(str string, n int, ci, skip bool) string {
	result := make([]string, n)
	completedIdx := -1
	for _, r := range str {
		if completedIdx == n-1 {
			break
		}
		if !isAlphanumeric(r) {
			if skip {
				continue
			}
			r = '_'
		}
		if ci {
			r = unicode.ToLower(r)
		}
		completedIdx++
		result[completedIdx] = string(r)
	}
	if completedIdx < n-1 {
		for i := completedIdx + 1; i < n; i++ {
			result[i] = "_"
		}
	}
	return strings.Join(result, string(os.PathSeparator))
}

func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func wasFlagGiven(flagName string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == flagName {
			found = true
		}
	})
	return found
}

func strToBool(s string) (bool, error) {
	switch strings.ToLower(s) {
	case "true", "1":
		return true, nil
	case "false", "0":
		return false, nil
	default:
		return false, fmt.Errorf("cannot convert given value '%s' to boolean", s)
	}
}
