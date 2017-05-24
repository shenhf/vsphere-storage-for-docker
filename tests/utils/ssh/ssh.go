// Copyright 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This util exposes util to invoke remove commands using ssh

package ssh

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/vmware/docker-volume-vsphere/tests/utils/misc"
)

// sshIdentity an array variable to prepare ssh input parameter to pass identify value
var sshIdentity = []string{strings.Split(os.Getenv("SSH_KEY_OPT"), " ")[0], strings.Split(os.Getenv("SSH_KEY_OPT"), " ")[1], "-q", "-kTax", "-o StrictHostKeyChecking=no"}

// InvokeCommand - can be consumed by test directly to invoke
// any command on the remote host.
// remoteHostIP: remote machine address to execute on the machine
// cmd: A command string to be executed on the remote host as per
// remoteHostIP value
func InvokeCommand(ip, cmd string) ([]byte, error) {
	return exec.Command("/usr/bin/ssh", append(sshIdentity, "root@"+ip, cmd)...).CombinedOutput()
}

// InvokeCommandLocally - can be consumed by test directly to invoke
// any command on the remote host.
// remoteHostIP: remote machine address to execute on the machine
// cmd: A command string to be executed on the remote host as per
// remoteHostIP value
func InvokeCommandLocally(cmd string) string {
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to invoke command [%s]: %v", cmd, err)
	}
	return strings.TrimSpace(misc.FormatOutput(out[:]))
}