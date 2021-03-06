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
  "homepage_url": "https://github.com/ulumuri/mattermost-plugin-nagios",
  "support_url": "https://github.com/ulumuri/mattermost-plugin-nagios/issues",
  "release_notes_url": "https://github.com/ulumuri/mattermost-plugin-nagios/releases/tag/v1.0.1",
  "icon_path": "assets/orbit-467260.svg",
  "version": "1.0.1",
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
    "header": "Having problems configuring the plugin? [Check the configuration guide](https://github.com/ulumuri/mattermost-plugin-nagios/#configuring-the-plugin).",
    "footer": "To report an issue, make a suggestion or a contribution, or fork your own version of the plugin, [check the repository](https://github.com/ulumuri/mattermost-plugin-nagios).",
    "settings": [
      {
        "key": "NagiosURL",
        "display_name": "Nagios URL",
        "type": "text",
        "help_text": "The URL for your Nagios instance. Must start with http:// or https://.",
        "placeholder": "",
        "default": null
      },
      {
        "key": "Token",
        "display_name": "Token",
        "type": "generated",
        "help_text": "The token for the configuration files watcher.",
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
