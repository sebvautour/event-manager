package client

import (
	"errors"

	"github.com/sebvautour/event-manager/pkg/config"
)

type Script struct {
}

func (c *Client) GetScriptList() (scriptNames []string, err error) {
	return []string{"test_script_1"}, nil
}

func (c *Client) GetScript(name string) (script *config.Script, err error) {
	scripts, err := c.GetScripts()
	if err != nil {
		return nil, err
	}
	for _, script := range scripts {
		if script.Name == name {
			return &script, nil
		}
	}
	return nil, errors.New("Script not found")
}

func (c *Client) GetScripts() (scripts []config.Script, err error) {
	return []config.Script{
		config.Script{
			Config: config.Config{
				Name:       "test_script_1",
				ConfigType: config.ConfigTypeScript},
			Metadata: config.ScriptMetadata{
				Description:   "Example script",
				Type:          "test_script",
				InitialAuthor: "svautour",
			},
		},
	}, nil

}
