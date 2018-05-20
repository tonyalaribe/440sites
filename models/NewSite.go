package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/tonyalaribe/440sites/config"
	"github.com/tonyalaribe/440sites/utils"
)

var (
	CommandSiteDirFn = func(c *exec.Cmd) { c.Dir = config.Get().SitesDir }
)

func Command(commandString string, configSetters ...func(*exec.Cmd)) error {
	command := strings.Fields(commandString)
	var cmd *exec.Cmd

	if len(command) == 1 {
		cmd = exec.Command(command[0])
	} else if len(command) > 1 {
		cmd = exec.Command(command[0], command[1:]...)
	} else {
		return errors.New("no command to execute")
	}

	var err error
	for _, v := range configSetters {
		v(cmd)
	}
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return err
	}

	// do something with output
	fmt.Printf("%s\n", stdoutStderr)
	return nil
}

//NewSite runs the hugo new command and generates a new site at given location
func NewSite(shopid string) error {
	commandStr := fmt.Sprintf("hugo new site %s -f json --quiet", shopid)
	err := Command(commandStr, CommandSiteDirFn)
	if err != nil {
		log.Println(err)
	}

	destinationURL := filepath.Join(config.Get().SitesDir, shopid, "themes", "shop440-default")
	err = utils.CopyDir(config.Get().ThemesDir+"/shop440-default", destinationURL)
	if err != nil {
		log.Println(err)
	}

	conf := map[string]interface{}{}
	configTomlFile := filepath.Join(config.Get().SitesDir, shopid, "config.json")
	fileByte, err := ioutil.ReadFile(configTomlFile)
	if err != nil {
		log.Println(err)
	}

	conf["theme"] = "shop440-default"
	conf["baseURL"] = "/"

	fileByte, err = json.MarshalIndent(conf, "", "\t")
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(configTomlFile, fileByte, 0655)
	if err != nil {
		log.Println(err)
	}

	err = Command("hugo", CommandSiteDirFn)
	if err != nil {
		log.Println(err)
	}
	return nil
}
