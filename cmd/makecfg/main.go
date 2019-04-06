package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	lib      string
	target   string
	makeEnum bool
	outPath  string
	preset   string
)

var (
	prefix_function = []byte("// FUNCTION_NAME")
	prefix_enum     = []byte("// ENUM_NAME")
)

func init() {
	flag.StringVar(&lib, "lib", "gtk2", "GUI lib, gtk1, gtk2, gtk3, xforms, motif")
	flag.StringVar(&target, "os", "", "Target OS, unix, mac, win")
	flag.BoolVar(&makeEnum, "with-enum", false, "export enum to cfg")
	flag.StringVar(&outPath, "o", "gtk-server.cfg", "output of gtk-server.cfg")
	flag.StringVar(&preset, "preset", "preset.ini", "preset.ini")
}

func main() {
	flag.Parse()
	if len(target) == 0 {
		switch runtime.GOOS {
		case "windows":
			target = "win"
		case "darwin":
			target = "mac"
		default:
			target = "unix"
		}
	}

	srcPathSect := lib
	libDepCmnSect := target + ".common"
	libDepSect := target + "." + lib
	if strings.HasPrefix(lib, "gtk") {
		srcPathSect = "gtk"
	}

	cfg, err := ini.LoadSources(ini.LoadOptions{
		UnparseableSections: []string{libDepSect, libDepCmnSect},
		AllowShadows:        true,
	},
		preset)
	if err != nil {
		log.Fatal(err)
	}

	srcPaths := cfg.Section(srcPathSect).Key("SRC_PATH").ValueWithShadows()
	libDep := cfg.Section(libDepSect).Body()

	srcPaths = append(srcPaths, cfg.Section("common").Key("SRC_PATH").ValueWithShadows()...)
	libDep = libDep + "\n" + cfg.Section(libDepCmnSect).Body()

	w, err := os.OpenFile(outPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	fmt.Fprintf(w, "# makecfg with os=%v, lib=%v, with-enum=%v\n", target, lib, makeEnum)

	_, _ = w.WriteString(libDep)

	fmt.Fprintln(w)

	for _, srcPath := range srcPaths {
		goFiles, err := filepath.Glob(filepath.Join(srcPath, "*.go"))
		if err != nil {
			log.Fatal(err)
		}

		for _, goFile := range goFiles {
			if lib == "gtk1" || lib == "gtk2" {
				if strings.HasSuffix(goFile, ".gtk3.go") {
					continue
				}
			} else if lib == "gtk3" {
				if strings.HasSuffix(goFile, ".gtk1.go") || strings.HasSuffix(goFile, ".gtk2.go") {
					continue
				}
			}

			r, err := os.Open(goFile)
			if err != nil {
				log.Fatal(err)
			}

			s := bufio.NewScanner(r)
			for s.Scan() {
				line := s.Bytes()
				if bytes.HasPrefix(line, prefix_function) || (makeEnum && bytes.HasPrefix(line, prefix_enum)) {
					_, _ = w.Write(line[3:])
					_, _ = w.WriteString("\n")
				}
			}

			r.Close()
		}
	}
}
