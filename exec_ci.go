package main

import (
  "fmt"
  "os/exec"
  "os"
  "sync"
  "strings"
  "io/ioutil"
  "time"
  "flag"
)

type test struct{
  name string
  playbook string
  user string
}

var wg sync.WaitGroup

func exec_net_ci(tst_data test, integration_dir string, inv string, limit string) {
  defer wg.Done()

  if limit != ""{
    if limit != tst_data.name{
      fmt.Println("Skipping: ", tst_data.name)
      return
    }
  }

  fmt.Println("Starting test: ", tst_data.name)

  start := time.Now()
  cmd_str := fmt.Sprintf("ansible-playbook -i %s -vvvv -u %s %s", inv, tst_data.user, tst_data.playbook)

  command := strings.Split(cmd_str, " ")
  cmd := exec.Command(command[0], command[1:]...)
  cmd.Dir = integration_dir
  out, err := cmd.CombinedOutput()

  elapsed := time.Since(start)

  msg := ""

  if err != nil {
    msg = msg + tst_data.name + " failed "
    msg = msg + "after " + elapsed.String()
    msg = msg + "\n    with command:  " + cmd_str
    msg = msg + "\n    executed from: " + integration_dir
  } else {
    msg = msg + tst_data.name + " succeeded "
    msg = msg + "after " + elapsed.String()
  }

  fmt.Println(msg)

  ioutil.WriteFile("results/" + tst_data.name + ".log", out, 0777)
}

func compile_report_card(tests []test){
  fmt.Println("Compiling Report Card")
  report_card := ""
  failed := ""
  passed := ""

  report_card = report_card + "# Results \n"
  for _, tst := range tests{
    dat, _ := ioutil.ReadFile("results/" + tst.name + ".err")
    report_card = report_card + "## " + tst.name + "\n"
    if len(strings.Split(string(dat), "\n")) < 4 {
      passed = passed + "* [" + tst.name + "](#" + tst.name + ")\n"
      report_card = report_card + "*All tests passed!*" + "\n\n"
    } else {
      failed = failed + "* [" + tst.name + "](#" + tst.name + ")\n"
      report_card = report_card + "*The following tests have failed:* \n"
      report_card = report_card + "[_See the .log files for details_](./" + tst.name + ".log) \n"
      report_card = report_card + "```\n"
      report_card = report_card + string(dat)
      report_card = report_card + "```\n"
    }
  }

  title := "# CI Test Report Card \n\n"
  summary := "## Test Summary \n"
  summary = summary + "### Passed: \n" + passed + "\n"
  summary = summary + "### Failed: \n" + failed + "\n"

  ioutil.WriteFile("results/report_card.md", []byte(title + summary + report_card), 0777)
}

func main() {
  // CONFIGURE FILE PATHS
  integration_dir := "/Users/dnewswan/code/concurrent_ci/ansible/test/integration/"
  exec_ci_dir := "/Users/dnewswan/code/concurrent_ci/"
  inv  := "#inventory.with_pass"

  // Configure environment (this shouldn't have to be updated on an individual basis)
  os.Setenv("ANSIBLE_ROLES_PATH", integration_dir + "targets/")
  os.Setenv("ANSIBLE_CALLBACK_PLUGINS", exec_ci_dir + "ansible_plugins")
  os.Setenv("ANSIBLE_CALLBACK_WHITELIST", "test_plug")
  os.Setenv("TASK_ERROR_DIR", exec_ci_dir+"/results/")

  // These should help avoid most 'unable to open shell problems'
  os.Setenv("ANSIBLE_CONNECT_TIMEOUT", "60")
  os.Setenv("ANSIBLE_TIMEOUT", "60")
  os.Setenv("ANSIBLE_HOST_KEY_CHECKING", "false")
  os.Setenv("ANSIBLE_COMMAND_TIMEOUT", "60")

  // os.Setenv("ANSIBLE_LOG_PATH", exec_ci_dir + "log")

  tests := []test{
    // test {
    //   name: "vyos",
    //   playbook: "vyos.yaml",
    //   user: "vyos",
    // },
    // test {
    //   name: "eos",
    //   playbook: "eos.yaml",
    //   user: "admin",
    // },
    test {
      name: "junos",
      playbook: "junos.yaml --diff",
      user: "ansible",
    },
    test {
      name: "ios",
      playbook: "ios.yaml",
      user: "cisco",
    },
    test {
      name: "iosxr",
      playbook: "iosxr.yaml",
      user: "root",
    },
    test {
      name: "nxos",
      playbook: "nxos.yaml",
      user: "admin",
    },
  }

  var limit string;

  flag.StringVar(&limit, "l", "", "limit to single test")
  flag.Parse()

  wg.Add(len(tests))
  for _, tst := range tests{
    go exec_net_ci(tst, integration_dir, inv, limit)
  }
  wg.Wait()

  compile_report_card(tests)
}
