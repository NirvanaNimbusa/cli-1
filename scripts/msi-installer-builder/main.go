package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/google/uuid"
)

type languagePreset int

// Language presets
const (
	Perl languagePreset = iota
	Python
	Unknown
)

func (lp languagePreset) String() string {
	if lp == Perl {
		return "Perl"
	}
	if lp == Python {
		return "Python"
	}
	return pad("PRESET")
}

// InFile is the template file for the Product.wxs file
var InFile string = "installers/msi-language/Product.p.wxs"

// OutFile is the path to the generated file
var OutFile string = "installers/msi-language/Product.wxs"

type config struct {
	Preset              string
	ID                  string
	ProjectName         string
	Version             string
	ReleaseNotes        string
	Icon                string
	ProjectOwnerAndName string
}

func seededUUID(seed string) string {
	bytes := []byte(seed)
	hash := sha256.New()
	hash.Write(bytes)

	uuid := uuid.NewHash(hash, uuid.UUID{}, bytes, 0)
	return uuid.String()
}

func parsePreset(p string) (languagePreset, error) {
	if strings.ToLower(p) == "perl" {
		return Perl, nil
	}
	if strings.ToLower(p) == "python" {
		return Python, nil
	}
	return Unknown, fmt.Errorf("Invalid language preset: %s", p)
}

func icon(p languagePreset) (string, error) {
	if p == Perl {
		return "assets/perl.ico", nil
	}
	return "", fmt.Errorf("No icon for language Preset %v", p)
}

func normalize(preset languagePreset, c *config) (*config, error) {
	parts := strings.SplitN(c.ProjectOwnerAndName, "/", 2)
	if len(parts) != 2 {
		return c, fmt.Errorf("Second argument must be of type owner/project")
	}

	c.ProjectName = parts[1]
	c.ID = seededUUID(c.ProjectOwnerAndName)

	ic, err := icon(preset)
	if err != nil {
		return c, err
	}
	c.Icon = ic
	return c, nil
}

func pad(s string) string {
	return s + strings.Repeat("-", 246-4-len(s))
}

func baseConfig() *config {
	return &config{
		ID:                  "FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF",
		Version:             "255.255.255",
		Icon:                "./assets/as.ico",
		Preset:              Unknown.String(),
		ProjectOwnerAndName: pad("PROJECT_OWNER_AND_NAME"),
		ReleaseNotes:        pad("RELEASE_NOTES"),
		ProjectName:         pad("PROJECT_NAME"),
	}
}

func parseArgs(args []string) (*config, error) {
	if len(os.Args) == 1 {
		// empty configuration for quick testing
		return normalize(Perl, &config{
			Preset:              Perl.String(),
			ProjectOwnerAndName: "ActiveState/ActivePerl-5.26",
			Version:             "5.26.3000",
			ReleaseNotes:        "http://docs.activestate.com/activeperl/5.26/get/relnotes/",
		})
	}
	if len(os.Args) == 5 {
		preset, err := parsePreset(os.Args[1])
		if err != nil {
			return nil, err
		}
		return normalize(preset, &config{
			Preset:              preset.String(),
			ProjectOwnerAndName: os.Args[2],
			Version:             os.Args[3],
			ReleaseNotes:        os.Args[4],
		})
	}

	if len(os.Args) == 2 && os.Args[1] == "base" {
		return baseConfig(), nil
	}

	return nil, fmt.Errorf("invalid arguments: Expected <preset> <owner/name> <version> <relNotes> | \"base\"")
}

func run(args []string) error {
	c, err := parseArgs(args)
	if err != nil {
		return err
	}

	in, err := ioutil.ReadFile(filepath.FromSlash(InFile))
	if err != nil {
		return err
	}
	tmpl, err := template.New("Product.wxs").Parse(string(in))
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filepath.FromSlash(OutFile), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0533)
	if err != nil {
		return err
	}
	defer f.Close()
	err = tmpl.Execute(f, c)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := run(os.Args)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

}
