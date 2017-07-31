# CI Test Report Card 

## Test Summary 
### Passed: 

### Failed: 
* [junos](#junos)
* [ios](#ios)
* [iosxr](#iosxr)
* [nxos](#nxos)

# Results 
## junos
*The following tests have failed:* 
[_See the .log files for details_](./junos.log) 
```
############ FAILURE ############
[junos_command : test contains operator with xml encoding] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_command/tests/netconf_xml/contains.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_config : setup] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_config/tests/netconf/backup.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_facts : Collect default facts from device] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_facts/tests/netconf/facts.yaml:5

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_netconf : Setup] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_netconf/tests/cli/changeport.yaml:5

[_ansible_parsed]=> False
[_ansible_no_log]=> False
[module_stderr]=> Traceback (most recent call last):
  File "/var/folders/ql/d9_087wj4cd10grvwzxmk4hc0000gn/T/ansible_t2EJiP/ansible_module_junos_netconf.py", line 214, in <module>
    main()
  File "/var/folders/ql/d9_087wj4cd10grvwzxmk4hc0000gn/T/ansible_t2EJiP/ansible_module_junos_netconf.py", line 202, in main
    commit_configuration(module)
  File "/var/folders/ql/d9_087wj4cd10grvwzxmk4hc0000gn/T/ansible_t2EJiP/ansible_modlib.zip/ansible/module_utils/junos.py", line 158, in commit_configuration
  File "/var/folders/ql/d9_087wj4cd10grvwzxmk4hc0000gn/T/ansible_t2EJiP/ansible_modlib.zip/ansible/module_utils/netconf.py", line 44, in send_request
  File "src/lxml/lxml.etree.pyx", line 3228, in lxml.etree.fromstring (src/lxml/lxml.etree.c:79594)
  File "src/lxml/parser.pxi", line 1848, in lxml.etree._parseMemoryDocument (src/lxml/lxml.etree.c:119113)
  File "src/lxml/parser.pxi", line 1736, in lxml.etree._parseDoc (src/lxml/lxml.etree.c:117793)
  File "src/lxml/parser.pxi", line 1102, in lxml.etree._BaseParser._parseDoc (src/lxml/lxml.etree.c:112037)
  File "src/lxml/parser.pxi", line 595, in lxml.etree._ParserContext._handleParseResultDoc (src/lxml/lxml.etree.c:105881)
  File "src/lxml/parser.pxi", line 706, in lxml.etree._handleParseResult (src/lxml/lxml.etree.c:107589)
  File "src/lxml/parser.pxi", line 635, in lxml.etree._raiseParseError (src/lxml/lxml.etree.c:106443)
  File "<string>", line 2
lxml.etree.XMLSyntaxError: Extra content at the end of the document, line 2, column 17

[changed]=> False
[module_stdout]=> 
[failed]=> True
[rc]=> 0
[msg]=> MODULE FAILURE



############ FAILURE ############
[junos_rpc : Execute RPC on device] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_rpc/tests/netconf/rpc.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_vlan : setup - remove vlan] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_vlan/tests/netconf/basic.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_interface : setup - remove interface] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_interface/tests/netconf/basic.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_banner : setup - remove login banner] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_banner/tests/netconf/basic.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_system : setup - remove hostname] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_system/tests/netconf/basic.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_logging : setup - remove file logging] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_logging/tests/netconf/basic.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_user : setup - remove user] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_user/tests/netconf/basic.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_static_route : setup - remove static route] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_static_route/tests/netconf/basic.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_linkagg : setup - remove linkagg] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_linkagg/tests/netconf/basic.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_l3_interface : setup - remove interface address] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_l3_interface/tests/netconf/basic.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_lldp : setup - Disable lldp and remove it's configuration] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_lldp/tests/netconf/basic.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_lldp_interface : setup - Remove lldp interface configuration] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_lldp_interface/tests/netconf/basic.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[junos_vrf : setup - remove vrf] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_vrf/tests/netconf/basic.yaml:4

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[Has any previous test failed?] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/junos.yaml:174

[msg]=> One or more tests failed, check log for details
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



```
# Results 
## ios
*The following tests have failed:* 
[_See the .log files for details_](./ios.log) 
```
############ FAILURE ############
[prepare_ios_tests : Ensure we have loopback 888 for testing] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

[_ansible_parsed]=> True
[stderr_lines]=> [u'show running-config', u'            ^', u"% Invalid input detected at '^' marker.", u'', u'ios01>']
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[stderr]=> show running-config
            ^
% Invalid input detected at '^' marker.

ios01>
[invocation]=> {u'module_args': {u'multiline_delimiter': u'@', u'authorize': None, u'force': False, u'diff_against': None, u'replace': u'line', u'running_config': None, u'save_when': u'never', u'port': None, u'before': None, u'auth_pass': None, u'parents': None, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'authorize': None, u'ssh_keyfile': None, u'auth_pass': None, u'host': u'ios01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'cli'}, u'save': False, u'match': u'line', u'username': None, u'timeout': None, u'after': None, u'host': None, u'password': None, u'diff_ignore_lines': None, u'src': u'interface Loopback888\n description test for ansible\n shutdown\n\n', u'ssh_keyfile': None, u'lines': None, u'intended_config': None, u'defaults': False, u'backup': False}}
[msg]=> unable to retrieve current config



############ FAILURE ############
[prepare_ios_tests : Ensure we have loopback 888 for testing] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

[_ansible_parsed]=> True
[stderr_lines]=> [u'show running-config', u'            ^', u"% Invalid input detected at '^' marker.", u'', u'ios01>']
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[stderr]=> show running-config
            ^
% Invalid input detected at '^' marker.

ios01>
[invocation]=> {u'module_args': {u'multiline_delimiter': u'@', u'authorize': None, u'force': False, u'diff_against': None, u'replace': u'line', u'running_config': None, u'save_when': u'never', u'port': None, u'before': None, u'auth_pass': None, u'parents': None, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'authorize': None, u'ssh_keyfile': None, u'auth_pass': None, u'host': u'ios01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'cli'}, u'save': False, u'match': u'line', u'username': None, u'timeout': None, u'after': None, u'host': None, u'password': None, u'diff_ignore_lines': None, u'src': u'interface Loopback888\n description test for ansible\n shutdown\n\n', u'ssh_keyfile': None, u'lines': None, u'intended_config': None, u'defaults': False, u'backup': False}}
[msg]=> unable to retrieve current config



############ FAILURE ############
[prepare_ios_tests : Ensure we have loopback 888 for testing] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

[_ansible_parsed]=> True
[stderr_lines]=> [u'show running-config', u'            ^', u"% Invalid input detected at '^' marker.", u'', u'ios01>']
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[stderr]=> show running-config
            ^
% Invalid input detected at '^' marker.

ios01>
[invocation]=> {u'module_args': {u'multiline_delimiter': u'@', u'authorize': None, u'force': False, u'diff_against': None, u'replace': u'line', u'running_config': None, u'save_when': u'never', u'port': None, u'before': None, u'auth_pass': None, u'parents': None, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'authorize': None, u'ssh_keyfile': None, u'auth_pass': None, u'host': u'ios01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'cli'}, u'save': False, u'match': u'line', u'username': None, u'timeout': None, u'after': None, u'host': None, u'password': None, u'diff_ignore_lines': None, u'src': u'interface Loopback888\n description test for ansible\n shutdown\n\n', u'ssh_keyfile': None, u'lines': None, u'intended_config': None, u'defaults': False, u'backup': False}}
[msg]=> unable to retrieve current config



############ FAILURE ############
[prepare_ios_tests : Ensure we have loopback 888 for testing] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

[_ansible_parsed]=> True
[stderr_lines]=> [u'show running-config', u'            ^', u"% Invalid input detected at '^' marker.", u'', u'ios01>']
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[stderr]=> show running-config
            ^
% Invalid input detected at '^' marker.

ios01>
[invocation]=> {u'module_args': {u'multiline_delimiter': u'@', u'authorize': None, u'force': False, u'diff_against': None, u'replace': u'line', u'running_config': None, u'save_when': u'never', u'port': None, u'before': None, u'auth_pass': None, u'parents': None, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'authorize': None, u'ssh_keyfile': None, u'auth_pass': None, u'host': u'ios01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'cli'}, u'save': False, u'match': u'line', u'username': None, u'timeout': None, u'after': None, u'host': None, u'password': None, u'diff_ignore_lines': None, u'src': u'interface Loopback888\n description test for ansible\n shutdown\n\n', u'ssh_keyfile': None, u'lines': None, u'intended_config': None, u'defaults': False, u'backup': False}}
[msg]=> unable to retrieve current config



############ FAILURE ############
[prepare_ios_tests : Ensure we have loopback 888 for testing] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

[_ansible_parsed]=> True
[stderr_lines]=> [u'show running-config', u'            ^', u"% Invalid input detected at '^' marker.", u'', u'ios01>']
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[stderr]=> show running-config
            ^
% Invalid input detected at '^' marker.

ios01>
[invocation]=> {u'module_args': {u'multiline_delimiter': u'@', u'authorize': None, u'force': False, u'diff_against': None, u'replace': u'line', u'running_config': None, u'save_when': u'never', u'port': None, u'before': None, u'auth_pass': None, u'parents': None, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'authorize': None, u'ssh_keyfile': None, u'auth_pass': None, u'host': u'ios01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'cli'}, u'save': False, u'match': u'line', u'username': None, u'timeout': None, u'after': None, u'host': None, u'password': None, u'diff_ignore_lines': None, u'src': u'interface Loopback888\n description test for ansible\n shutdown\n\n', u'ssh_keyfile': None, u'lines': None, u'intended_config': None, u'defaults': False, u'backup': False}}
[msg]=> unable to retrieve current config



############ FAILURE ############
[prepare_ios_tests : Ensure we have loopback 888 for testing] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

[_ansible_parsed]=> True
[stderr_lines]=> [u'show running-config', u'            ^', u"% Invalid input detected at '^' marker.", u'', u'ios01>']
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[stderr]=> show running-config
            ^
% Invalid input detected at '^' marker.

ios01>
[invocation]=> {u'module_args': {u'multiline_delimiter': u'@', u'authorize': None, u'force': False, u'diff_against': None, u'replace': u'line', u'running_config': None, u'save_when': u'never', u'port': None, u'before': None, u'auth_pass': None, u'parents': None, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'authorize': None, u'ssh_keyfile': None, u'auth_pass': None, u'host': u'ios01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'cli'}, u'save': False, u'match': u'line', u'username': None, u'timeout': None, u'after': None, u'host': None, u'password': None, u'diff_ignore_lines': None, u'src': u'interface Loopback888\n description test for ansible\n shutdown\n\n', u'ssh_keyfile': None, u'lines': None, u'intended_config': None, u'defaults': False, u'backup': False}}
[msg]=> unable to retrieve current config



############ FAILURE ############
[prepare_ios_tests : Ensure we have loopback 888 for testing] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

[_ansible_parsed]=> True
[stderr_lines]=> [u'show running-config', u'            ^', u"% Invalid input detected at '^' marker.", u'', u'ios01>']
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[stderr]=> show running-config
            ^
% Invalid input detected at '^' marker.

ios01>
[invocation]=> {u'module_args': {u'multiline_delimiter': u'@', u'authorize': None, u'force': False, u'diff_against': None, u'replace': u'line', u'running_config': None, u'save_when': u'never', u'port': None, u'before': None, u'auth_pass': None, u'parents': None, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'authorize': None, u'ssh_keyfile': None, u'auth_pass': None, u'host': u'ios01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'cli'}, u'save': False, u'match': u'line', u'username': None, u'timeout': None, u'after': None, u'host': None, u'password': None, u'diff_ignore_lines': None, u'src': u'interface Loopback888\n description test for ansible\n shutdown\n\n', u'ssh_keyfile': None, u'lines': None, u'intended_config': None, u'defaults': False, u'backup': False}}
[msg]=> unable to retrieve current config



############ FAILURE ############
[Has any previous test failed?] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/ios.yaml:93

[msg]=> One or more tests failed, check log for details
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



```
# Results 
## iosxr
*The following tests have failed:* 
[_See the .log files for details_](./iosxr.log) 
```
############ FAILURE ############
[iosxr_config : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/iosxr_config/tests/cli/toplevel_nonidempotent.yaml:26

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == true
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[Has any previous test failed?] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/iosxr.yaml:85

[msg]=> One or more tests failed, check log for details
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



```
# Results 
## nxos
*The following tests have failed:* 
[_See the .log files for details_](./nxos.log) 
```
############ FAILURE ############
[nxos_command : Run show running-config bgp - should pass] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_command/tests/nxapi/sanity.yaml:42

[status]=> -1
[_ansible_parsed]=> True
[_ansible_no_log]=> False
[url]=> http://nxos01:80/ins
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'url_password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'commands': [u'sh running-config bgp'], u'ssh_keyfile': None, u'timeout': 10, u'retries': 10, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'interval': 1, u'host': u'nxos01', u'match': u'all', u'url_username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'host': u'nxos01', u'timeout': 60, u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}, u'use_ssl': False, u'wait_for': None, u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}}
[msg]=> Connection failure: timed out



############ FAILURE ############
[nxos_config : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_config/tests/cli/defaults.yaml:37

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == false
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[nxos_mtu : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_mtu/tests/cli/set_sysmtu.yaml:26

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == false
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[nxos_system : configure name_servers] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_system/tests/cli/set_name_servers.yaml:13

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'lookup_source': None, u'domain_search': None, u'hostname': None, u'system_mtu': None, u'domain_name': None, u'name_servers': [u'1.1.1.1', u'2.2.2.2', u'3.3.3.3'], u'host': None, u'state': u'present', u'timeout': 10, u'provider': None, u'domain_lookup': None, u'use_ssl': None, u'password': None, u'validate_certs': None, u'port': None, u'transport': u'cli'}}
[msg]=> no ip name-server use-vrf
                                  ^
% Invalid ip address at '^' marker.
nxos01(config)# 



############ FAILURE ############
[nxos_interface : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_interface/tests/cli/set_state_absent.yaml:16

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == true
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[nxos_user : Create user] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_user/tests/cli/basic.yaml:2

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'authorize': True, u'name': u'netend', u'roles': u'network-operator', u'state': u'present', u'provider': None, u'transport': u'cli'}}
[msg]=> Unsupported parameters for (nxos_user) module: authorize. Supported parameters include: aggregate,host,name,password,port,provider,purge,roles,ssh_keyfile,sshkey,state,timeout,transport,update_password,use_ssl,username,validate_certs



############ FAILURE ############
[nxos_banner : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_banner/tests/cli/basic-no-exec.yaml:33

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == false
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[nxos_vtp_password : configure vtp password] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_vtp_password/tests/nxapi/sanity.yaml:11

[body]=> <html>
<head><title>401 Authorization Required</title></head>
<body bgcolor="white">
<center><h1>401 Authorization Required</h1></center>
<hr><center>nginx/1.4.5</center>
</body>
</html>

[status]=> 401
[content-length]=> 194
[_ansible_no_log]=> False
[url]=> http://nxos01:80/ins
[changed]=> False
[server]=> nginx/1.4.5
[failed]=> True
[connection]=> close
[_ansible_parsed]=> True
[invocation]=> {u'module_args': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'url_password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'vtp_password': None, u'ssh_keyfile': None, u'timeout': 10, u'state': u'present', u'host': u'nxos01', u'url_username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'host': u'nxos01', u'timeout': 60, u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}, u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}}
[date]=> Fri, 04 Aug 2017 22:43:23 GMT
[content-type]=> text/html
[www-authenticate]=> Basic realm="Secure Zone"
[msg]=> HTTP Error 401: Unauthorized



############ FAILURE ############
[nxos_vtp_version : disable feature vtp] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_vtp_version/tests/cli/sanity.yaml:17

[_ansible_parsed]=> True
[invalid_feature]=> vtp
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'password': None, u'ssh_keyfile': None, u'validate_certs': None, u'feature': u'vtp', u'include_defaults': None, u'state': u'disabled', u'timeout': 10, u'provider': None, u'host': None, u'use_ssl': None, u'save': None, u'config': None, u'port': None, u'transport': u'cli'}}
[features_currently_supported]=> {}
[msg]=> Invalid feature name.



############ FAILURE ############
[nxos_portchannel : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_portchannel/tests/common/sanity.yaml:57

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == false
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[nxos_portchannel : Disable feature LACP] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_portchannel/tests/common/sanity.yaml:87

[_ansible_parsed]=> True
[invalid_feature]=> lacp
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'password': None, u'ssh_keyfile': None, u'validate_certs': None, u'feature': u'lacp', u'include_defaults': None, u'state': u'disabled', u'timeout': 60, u'provider': None, u'host': None, u'use_ssl': None, u'save': None, u'config': None, u'port': None, u'transport': u'cli'}}
[features_currently_supported]=> {u'mvpn': u'disabled', u'pim': u'enabled', u'ldp': u'disabled', u'hsrp_engine': u'enabled', u'itd': u'disabled', u'scheduler': u'disabled', u'fabric_mcast': u'disabled', u'glbp': u'disabled', u'private-vlan': u'disabled', u'eth_port_sec': u'disabled', u'mvrp': u'disabled', u'pong': u'disabled', u'dot1x': u'disabled', u'ptp': u'disabled', u'lldp': u'enabled', u'mpls_te': u'disabled', u'evmed': u'disabled', u'elo': u'disabled', u'imp': u'disabled', u'vmtracker': u'disabled', u'interface-vlan': u'disabled', u'cable-management': u'disabled', u'vtp': u'enabled', u'vrrp': u'disabled', u'ldap': u'disabled', u'privilege': u'disabled', u'dhcp': u'disabled', u'sla_responder': u'disabled', u'fabric-access': u'disabled', u'pim6': u'disabled', u'bulkstat': u'disabled', u'vrrpv3': u'disabled', u'wccp': u'disabled', u'rise': u'disabled', u'l3vpn': u'disabled', u'vni': u'enabled', u'onep': u'disabled', u'hmm': u'disabled', u'nve': u'enabled', u'nxapi': u'enabled', u'rip': u'disabled', u'vpc': u'enabled', u'sftpServer': u'disabled', u'netflow': u'disabled', u'l2vpn': u'disabled', u'msdp': u'disabled', u'otv': u'disabled', u'cts': u'disabled', u'smart-channel': u'disabled', u'ngoam': u'disabled', u'isis': u'disabled', u'evc': u'disabled', u'pbr': u'disabled', u'lisp': u'disabled', u'tacacs': u'disabled', u'evb': u'disabled', u'bfd': u'disabled', u'sla_sender': u'disabled', u'ospfv3': u'disabled', u'bgp': u'disabled', u'tunnel': u'disabled', u'eigrp': u'disabled', u'vnseg_vlan': u'disabled', u'msrp': u'disabled', u'scpServer': u'enabled', u'ospf': u'enabled', u'telnetServer': u'enabled', u'sshServer': u'enabled', u'bfd_app': u'disabled'}
[msg]=> Invalid feature name.



############ FAILURE ############
[Has any previous test failed?] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/nxos.yaml:227

[msg]=> One or more tests failed, check log for details
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



```
