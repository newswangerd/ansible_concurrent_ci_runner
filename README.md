# Concurrent Test Runner

This is used to run all of the network tests concurrently. I got tired of running tests individually and sifting through thousands of lines of console output to find errors, so I created this. It runs everything at once and reports any errors that occur to results/_platform_.err with concise error messages that give the user an idea of exactly what failed. The rest of the console output is logged to results/_platform_.log for those times when just looking at the error messages isn't enough.

# Overview
This is written in a combination of Go and Python. Yes, I know, I should have written it entirely in Python because that's what Ansible is written in, but I don't know how to do concurrency in Python, so in the interest of time, Go works just fine.

`exec_ci.go` essentially just launches one go routine for each platform being tested. Each go routine simply executes one of the tests using a shell command such as `ansible-playbook -i inventory.network -vvvv -u root junos.yaml`, waits for the command to finish executing, and logs the output.

Failed tasks in results/_platform_.err are logged using a custom Ansible callback plugin which can be found in the `ansible_plugins` directory.

# Usage
1. Install go

2. Set up your ansible dev environment with `source ansible/hacking/env-setup` in the main Ansible repo.

3. Configure the variables at the top of the main function `exec_ci.go` with the relevant directories for your computer.

4. Since there isn't a way to pass the SSH password for the device via the command line, passwords have to be configured manually using `ansible_ssh_pass=` in inventory.networking.

5. Run with `go run exec_ci.go`. Tests will probably take around half an hour to complete.

Note, re running the tests reset all of the files in `results/`.

Some tinkering is still required to fix the "unable to open shell" errors.

Sample output:
```
Starting test for: eos
Starting test for: vyos
Starting test for: nxos
Starting test for: iosxr
Starting test for: ios
Starting test for: junos
iosxr failed after 4m11.067241604s
    with command: ansible-playbook -i /Users/dnewswan/code/ansible/test/integration/inventory.network -vvvv -u root /Users/dnewswan/code/ansible/test/integration/ios.yaml
junos failed after 4m34.909657518s
    with command: ansible-playbook -i /Users/dnewswan/code/ansible/test/integration/inventory.network -vvvv -u root /Users/dnewswan/code/ansible/test/integration/junos.yaml
ios failed after 5m16.944068975s
    with command: ansible-playbook -i /Users/dnewswan/code/ansible/test/integration/inventory.network -vvvv -u cisco /Users/dnewswan/code/ansible/test/integration/ios.yaml
vyos failed after 11m29.55132583s
    with command: ansible-playbook -i /Users/dnewswan/code/ansible/test/integration/inventory.network -vvvv -u vyos /Users/dnewswan/code/ansible/test/integration/vyos.yaml
eos failed after 22m51.937052768s
    with command: ansible-playbook -i /Users/dnewswan/code/ansible/test/integration/inventory.network -vvvv -u admin /Users/dnewswan/code/ansible/test/integration/eos.yaml
nxos failed after 25m17.384572714s
    with command: ansible-playbook -i /Users/dnewswan/code/ansible/test/integration/inventory.network -vvvv -u admin /Users/dnewswan/code/ansible/test/integration/nxos.yaml
```
