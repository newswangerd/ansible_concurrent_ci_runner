package main

import (
	"fmt"
  "os/exec"
  "os"
  "sync"
  "strings"
  "io/ioutil"
  "time"
)

type test struct{
  name string
  playbook string
  user string
}

var wg sync.WaitGroup

func exec_net_ci(tst_data test, integration_dir string, inv string) {
  defer wg.Done()

  fmt.Println("Starting test for: " + tst_data.name)

  start := time.Now()
  playbook :=integration_dir + tst_data.playbook
  cmd_str := fmt.Sprintf("ansible-playbook -i %s -vvvv -u %s %s", inv, tst_data.user, playbook)

  command := strings.Split(cmd_str, " ")
  cmd := exec.Command(command[0], command[1:]...)
	out, err := cmd.CombinedOutput()

  elapsed := time.Since(start)

  msg := ""

  if err != nil {
    msg = msg + tst_data.name + " failed "
    msg = msg + "after " + elapsed.String()
    msg = msg + "\n    with command: " + cmd_str
  } else {
    msg = msg + tst_data.name + " succeeded "
    msg = msg + "after " + elapsed.String()
  }



  fmt.Println(msg)

  ioutil.WriteFile("results/" + tst_data.name + ".log", out, 0777)
}

func main() {
	// CONFIG FIXME with more dynamic paths
  integration_dir := "/Users/dnewswan/code/ansible/test/integration/"
  inv  := integration_dir + "inventory.network"
	os.Setenv("ANSIBLE_ROLES_PATH", integration_dir + "targets/")
  os.Setenv("ANSIBLE_CALLBACK_PLUGINS", "/Users/dnewswan/code/concurrent_ci/ansible_plugins")
  os.Setenv("ANSIBLE_CALLBACK_WHITELIST", "test_plug")
  os.Setenv("TASK_ERROR_DIR", "/Users/dnewswan/code/concurrent_ci/results/")

  tests := []test{
    test {
      name: "vyos",
      playbook: "vyos.yaml",
      user: "vyos",
    },
    test {
      name: "eos",
      playbook: "eos.yaml",
      user: "admin",
    },
    test {
      name: "junos",
      playbook: "junos.yaml",
      user: "root",
    },
    test {
      name: "ios",
      playbook: "ios.yaml",
      user: "cisco",
    },
    test {
      name: "iosxr",
      playbook: "ios.yaml",
      user: "root",
    },
    test {
      name: "nxos",
      playbook: "nxos.yaml",
      user: "admin",
    },
  }


  commands := [][]string{}
  commands = append(commands, )

  wg.Add(len(tests))
  for _, tst := range tests{
    go exec_net_ci(tst, integration_dir, inv)
  }
  wg.Wait()
}
