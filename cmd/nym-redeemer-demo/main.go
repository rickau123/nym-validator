// main.go - redeemer entrypoint.
// Copyright (C) 2019  Nym Authors.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import "github.com/tav/golly/optparse"

func main() {
	var logo = `
  ____                            _        _   _                 
 / ___|___   ___ ___  _ __  _   _| |_     | \ | |_   _ _ __ ___  
| |   / _ \ / __/ _ \| '_ \| | | | __|____|  \| | | | | '_ \ _ \ 
| |___ (_) | (__ (_) | | | | |_| | |______| |\  | |_| | | | | | |
 \____\___/ \___\___/|_| |_|\__,_|\__|    |_| \_|\__, |_| |_| |_|
             (nym-redeemer-demo)                 |___/           
										 
`
	cmds := map[string]func([]string, string){
		"run": cmdRun,
	}
	info := map[string]string{
		"run": "Run a persistent demo nym token redeemer",
	}
	optparse.Commands("nym-redeemer-demo", "0.12.8", cmds, info, logo)
}
