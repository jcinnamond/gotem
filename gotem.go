package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type partials map[string]string

func main() {
	include_path := flag.String("I", "", "search `directory` for includes")

	flag.Parse()

	var p partials
	var err error
	if *include_path == "" {
		p = make(partials)
	} else {
		p, err = loadPartials(*include_path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error load partials: %s", err)
			os.Exit(1)
		}
	}

	var in io.ReadCloser
	var out io.WriteCloser

	switch flag.NArg() {
	case 0:
		in = os.Stdin
		out = os.Stdout
	case 1:
		in = openIn(flag.Arg(0))
		defer in.Close()
		out = os.Stdout
	case 2:
		in = openIn(flag.Arg(0))
		defer in.Close()
		out = openOut(flag.Arg(1))
		defer out.Close()
	default:
		fmt.Fprintln(os.Stderr, "Too many arguments.")
		fmt.Fprintf(os.Stderr, "Usage: %s [input [output]] [arguments]\n\n", path.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(3)
	}

	compileTemplate(in, out, p)
}

// Tries to open an input file or stdin, panics on error.
func openIn(path string) io.ReadCloser {
	if path == "-" {
		return os.Stdin
	}

	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open input file %s: %s\n", path, err)
		os.Exit(1)
	}
	return f
}

// Tries to open an output file or stdout, panics on error.
func openOut(path string) io.WriteCloser {
	if path == "-" {
		return os.Stdout
	}

	f, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open output file %s: %s\n", path, err)
		os.Exit(1)
	}
	return f
}

// Return a list name/content partials from a given directory
func loadPartials(path string) (partials, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	p := make(partials)
	for _, f := range files {
		name := f.Name()
		if filepath.Ext(f.Name()) == ".template" {
			name = strings.TrimSuffix(f.Name(), ".template")
			fullpath := strings.Join([]string{path, f.Name()}, "/")
			content, err := ioutil.ReadFile(fullpath)
			if err != nil {
				return nil, err
			}
			p[name] = string(content)
		}
	}
	return p, nil
}

// Loads and compiles a single template
func compileTemplate(in io.Reader, out io.Writer, partials partials) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	// Writing this as a function literal creates a closure around partials
	include := func(args ...interface{}) template.HTML {
		if len(args) < 1 {
			return ""
		}
		name, ok := args[0].(string)
		if !ok {
			return ""
		}
		return template.HTML(partials[name])
	}

	t := template.New("main")
	t = t.Funcs(template.FuncMap{"include": include})
	t = template.Must(t.Parse(string(src)))
	return t.Execute(out, "")
}
