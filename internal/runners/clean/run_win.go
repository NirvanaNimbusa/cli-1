// +build windows

package clean

import (
	"fmt"
	"os/exec"

	"github.com/gobuffalo/packr"

	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/language"
	"github.com/ActiveState/cli/internal/scriptfile"
)

func removeConfig(configPath string) error {
	config.SkipSave(true)
	return runScript("removeConfig", configPath)
}

func removeInstall(installPath string) error {
	return runScript("removeInstall", installPath)
}

func runScript(scriptName, path string) error {
	box := packr.NewBox("../../../assets/scripts/")
	scriptBlock := box.String(fmt.Sprintf("%s.bat", scriptName))
	sf, err := scriptfile.New(language.Batch, scriptName, scriptBlock)
	if err != nil {
		return err
	}

	cmd := exec.Command("cmd.exe", "/C", sf.Filename(), path)
	err = cmd.Start()
	if err != nil {
		return err
	}

	return nil
}
