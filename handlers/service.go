package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/deployithq/deployit/errors"
	"github.com/fatih/color"
	"gopkg.in/urfave/cli.v2"
	"net/http"
	"os"
	"strconv"
)

func ServiceStart(c *cli.Context) error {

	env := NewEnv()

	color.Cyan("Starting %s ...", os.Args[1])

	res, err := http.Post(fmt.Sprintf("%s/service/%s/start", env.HostUrl, os.Args[1]), "application/json", new(bytes.Buffer))
	if err != nil {
		env.Log.Error(err)
		return err
	}

	if res.StatusCode != 200 {
		err = errors.ParseError(res)
		env.Log.Error(err)
		return err
	}

	color.Cyan("Finished!")

	return nil
}

func ServiceStop(c *cli.Context) error {
	env := NewEnv()

	color.Cyan("Stopping %s ...", os.Args[1])

	res, err := http.Post(fmt.Sprintf("%s/service/%s/stop", env.HostUrl, os.Args[1]), "application/json", new(bytes.Buffer))
	if err != nil {
		env.Log.Error(err)
		return err
	}

	if res.StatusCode != 200 {
		err = errors.ParseError(res)
		env.Log.Error(err)
		return err
	}

	color.Cyan("Finished!")

	return nil
}

func ServiceRestart(c *cli.Context) error {

	env := NewEnv()

	color.Cyan("Restarting %s ...", os.Args[1])

	res, err := http.Post(fmt.Sprintf("%s/service/%s/restart", env.HostUrl, os.Args[1]), "application/json", new(bytes.Buffer))
	if err != nil {
		env.Log.Error(err)
		return err
	}

	if res.StatusCode != 200 {
		err = errors.ParseError(res)
		env.Log.Error(err)
		return err
	}

	color.Cyan("Finished!")

	return nil
}

func ServiceDeploy(c *cli.Context) error {

	// TODO Adapt for other services

	env := NewEnv()

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/service/%s", env.HostUrl, os.Args[1]), new(bytes.Buffer))
	if err != nil {
		env.Log.Error(err)
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		env.Log.Error(err)
		return err
	}

	if res.StatusCode != 200 {
		err = errors.ParseError(res)
		env.Log.Error(err)
		return err
	}

	response := struct {
		Port     int64  `json:"port"`
		Password string `json:"password"`
	}{}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		env.Log.Error(err)
		return err
	}

	color.Cyan("Your redis adress: %s:%s", env.HostUrl, strconv.FormatInt(response.Port, 10))
	color.Cyan("Your redis password: %s", response.Password)

	return nil
}

func ServiceRemove(c *cli.Context) error {

	env := NewEnv()

	color.Cyan("Removing %s ...", os.Args[1])

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/service/%s", env.HostUrl, os.Args[1]), new(bytes.Buffer))
	if err != nil {
		env.Log.Error(err)
		return err
	}

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		env.Log.Error(err)
		return err
	}

	if res.StatusCode != 200 {
		err = errors.ParseError(res)
		env.Log.Error(err)
		return err
	}

	color.Cyan("Finished!")

	return nil
}