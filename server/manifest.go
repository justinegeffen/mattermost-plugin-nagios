// This file is automatically generated. Do not modify it manually.

package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "nagios",
  "name": "Nagios",
  "description": "Nagios plugin for Mattermost",
  "version": "1.0.0",
  "min_server_version": "5.12.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "settings_schema": {
    "header": "To report an issue, make a suggestion or a contribution, or fork your own version of the plugin, [check the repository](https://github.com/ulumuri/mattermost-plugin-nagios).",
    "footer": "",
    "settings": [
      {
        "key": "NagiosURL",
        "display_name": "Nagios URL",
        "type": "text",
        "help_text": "The URL for your Nagios instance. Must start with http:// or https://. For example: https://nagios.example.com.",
        "placeholder": "",
        "default": null
      }
    ]
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}