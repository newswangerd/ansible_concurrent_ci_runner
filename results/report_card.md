# CI Test Report Card 

## Test Summary 
### Passed: 
* [eos](#eos)

### Failed: 
* [vyos](#vyos)
* [junos](#junos)
* [ios](#ios)
* [iosxr](#iosxr)
* [nxos](#nxos)

# Results 
## vyos
*The following tests have failed:* 
[_See the .log files for details_](./vyos.log) 
```
############ FAILURE ############
vyos_l3_interface : Set IPv6 address/Users/dnewswan/code/ansible/test/integration/targets/vyos_l3_interface/tests/cli/basic.yaml:33

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'name': u'eth1', u'provider': None, u'state': u'present', u'purge': False, u'host': None, u'ipv4': None, u'timeout': None, u'ipv6': u'fd5d:12c9:2201:1::1/64', u'aggregate': None, u'password': None, u'port': None}}
[msg]=> commit failed: commit
[ interfaces ethernet eth1 address fd5d:12c9:2201:1::1/64 ]
RTNETLINK answers: No buffer space available

[[interfaces ethernet eth1]] failed
Commit failed
[edit]
vyos@vyos02# 



############ FAILURE ############
vyos_logging : assert/Users/dnewswan/code/ansible/test/integration/targets/vyos_logging/tests/cli/basic.yaml:11

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == true
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
vyos_interface : assert/Users/dnewswan/code/ansible/test/integration/targets/vyos_interface/tests/cli/basic.yaml:141

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> "set interfaces ethernet eth1 description test-interface-1" in result.commands
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
Has any previous test failed?/Users/dnewswan/code/ansible/test/integration/vyos.yaml:121

[msg]=> One or more tests failed, check log for details
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



```
# Results 
## eos
*All tests passed!*

# Results 
## junos
*The following tests have failed:* 
[_See the .log files for details_](./junos.log) 
```
############ FAILURE ############
[junos_command : test contains operator with xml encoding] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_command/tests/netconf_xml/contains.yaml:4

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'retries': 10, u'commands': [u'show version', u'show interfaces fxp0'], u'ssh_keyfile': None, u'rpcs': None, u'format': u'xml', u'interval': 1, u'wait_for': [u'result[0].rpc-reply.software-information.host-name contains vsrx01', u'result[1].rpc-reply.interface-information.physical-interface.name contains fxp0'], u'display': u'xml', u'host': None, u'match': u'all', u'timeout': 10, u'provider': {u'username': u'ansible', u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'netconf'}, u'password': None, u'port': None, u'transport': None}}
[msg]=> jxmlease is required but does not appear to be installed. It can be installed using `pip install jxmlease`



############ FAILURE ############
[junos_config : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_config/tests/netconf/multiple.yaml:20

[msg]=> The conditional check ''host-name;' not in result.diff.prepared' failed. The error was: error while evaluating conditional ('host-name;' not in result.diff.prepared): 'dict object' has no attribute 'diff'
[failed]=> True



############ FAILURE ############
[junos_netconf : Change port] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_netconf/tests/cli/changeport.yaml:13

[_ansible_parsed]=> False
[_ansible_no_log]=> False
[module_stderr]=> Traceback (most recent call last):
  File "/var/folders/ql/d9_087wj4cd10grvwzxmk4hc0000gn/T/ansible_IAwLym/ansible_module_junos_netconf.py", line 214, in <module>
    main()
  File "/var/folders/ql/d9_087wj4cd10grvwzxmk4hc0000gn/T/ansible_IAwLym/ansible_module_junos_netconf.py", line 202, in main
    commit_configuration(module)
  File "/var/folders/ql/d9_087wj4cd10grvwzxmk4hc0000gn/T/ansible_IAwLym/ansible_modlib.zip/ansible/module_utils/junos.py", line 158, in commit_configuration
  File "/var/folders/ql/d9_087wj4cd10grvwzxmk4hc0000gn/T/ansible_IAwLym/ansible_modlib.zip/ansible/module_utils/netconf.py", line 62, in send_request
  File "src/lxml/lxml.etree.pyx", line 3228, in lxml.etree.fromstring (src/lxml/lxml.etree.c:79594)
  File "src/lxml/parser.pxi", line 1848, in lxml.etree._parseMemoryDocument (src/lxml/lxml.etree.c:119113)
  File "src/lxml/parser.pxi", line 1736, in lxml.etree._parseDoc (src/lxml/lxml.etree.c:117793)
  File "src/lxml/parser.pxi", line 1102, in lxml.etree._BaseParser._parseDoc (src/lxml/lxml.etree.c:112037)
  File "src/lxml/parser.pxi", line 595, in lxml.etree._ParserContext._handleParseResultDoc (src/lxml/lxml.etree.c:105881)
  File "src/lxml/parser.pxi", line 706, in lxml.etree._handleParseResult (src/lxml/lxml.etree.c:107589)
  File "src/lxml/parser.pxi", line 646, in lxml.etree._raiseParseError (src/lxml/lxml.etree.c:106735)
  File "<string>", line 0
lxml.etree.XMLSyntaxError

[changed]=> False
[module_stdout]=> 
[failed]=> True
[rc]=> 0
[msg]=> MODULE FAILURE



############ FAILURE ############
[junos_template : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_template/tests/netconf/backup.yaml:31

[msg]=> The conditional check ''1.1.1.1/32' in result.diff.prepared' failed. The error was: error while evaluating conditional ('1.1.1.1/32' in result.diff.prepared): 'dict object' has no attribute 'diff'
[failed]=> True



############ FAILURE ############
[junos_interface : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_interface/tests/netconf/basic.yaml:145

[msg]=> The conditional check ''[edit interfaces ge-0/0/1]
-   disable;' in result.diff.prepared' failed. The error was: error while evaluating conditional ('[edit interfaces ge-0/0/1]
-   disable;' in result.diff.prepared): 'dict object' has no attribute 'diff'
[failed]=> True



############ FAILURE ############
[junos_logging : Configure logging parameters] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_logging/tests/netconf/basic.yaml:196

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'files': 40, u'ssh_keyfile': None, u'src_addr': None, u'name': None, u'facility': None, u'dest': None, u'level': None, u'host': None, u'purge': False, u'aggregate': None, u'state': u'present', u'timeout': 10, u'provider': {u'username': u'ansible', u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'netconf'}, u'active': True, u'rotate_frequency': 20, u'password': None, u'port': None, u'transport': None, u'size': 65536}}
[msg]=> <?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X49/junos" xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="urn:uuid:e143f076-afd5-4aa4-b0b6-4690a823c7c8">
<load-configuration-results>
<rpc-error>
<error-type>protocol</error-type>
<error-tag>operation-failed</error-tag>
<error-severity>error</error-severity>
<error-message>
could not set size
</error-message>
</rpc-error>
<rpc-error>
<error-type>protocol</error-type>
<error-tag>operation-failed</error-tag>
<error-severity>error</error-severity>
<error-message>
could not set files
</error-message>
</rpc-error>
<rpc-error>
<error-type>protocol</error-type>
<error-tag>operation-failed</error-tag>
<error-severity>error</error-severity>
<error-message>
could not set log-rotate-frequency
</error-message>
</rpc-error>
<ok/>
</load-configuration-results>
</rpc-reply>



############ FAILURE ############
[junos_user : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_user/tests/netconf/basic.yaml:25

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> '<class>read-only</class>' in config.xml
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[junos_static_route : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_static_route/tests/netconf/basic.yaml:70

[msg]=> The conditional check ''inactive: route 1.1.1.0/24' in result.diff.prepared' failed. The error was: error while evaluating conditional ('inactive: route 1.1.1.0/24' in result.diff.prepared): 'dict object' has no attribute 'diff'
[failed]=> True



############ FAILURE ############
[junos_linkagg : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_linkagg/tests/netconf/basic.yaml:135

[msg]=> The conditional check ''+   disable;' in result.diff.prepared' failed. The error was: error while evaluating conditional ('+   disable;' in result.diff.prepared): 'dict object' has no attribute 'diff'
[failed]=> True



############ FAILURE ############
[junos_l3_interface : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_l3_interface/tests/netconf/basic.yaml:27

[msg]=> The conditional check ''+       address 1.1.1.1/32;' in result.diff.prepared' failed. The error was: error while evaluating conditional ('+       address 1.1.1.1/32;' in result.diff.prepared): 'dict object' has no attribute 'diff'
[failed]=> True



############ FAILURE ############
[junos_lldp : setup - Disable lldp and remove it's configuration] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_lldp/tests/netconf/basic.yaml:4

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'interval': None, u'state': u'absent', u'purge': False, u'host': None, u'disable': True, u'transport': None, u'timeout': 10, u'provider': {u'username': u'ansible', u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'netconf'}, u'active': True, u'password': None, u'hold_multiplier': None, u'port': None, u'transmit_delay': None}}
[msg]=> <?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X49/junos" xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="urn:uuid:c585b8e1-b6c3-4067-bc9e-00663da64543">
<load-configuration-results>
<rpc-error>
<error-type>protocol</error-type>
<error-tag>operation-failed</error-tag>
<error-severity>error</error-severity>
<error-message>syntax error</error-message>
<error-info>
<bad-element>lldp</bad-element>
</error-info>
</rpc-error>
<rpc-error>
<error-type>protocol</error-type>
<error-tag>operation-failed</error-tag>
<error-severity>error</error-severity>
<error-message>syntax error</error-message>
<error-info>
<bad-element>lldp</bad-element>
</error-info>
</rpc-error>
<load-error-count>2</load-error-count>
</load-configuration-results>
</rpc-reply>



############ FAILURE ############
[junos_lldp_interface : setup - Remove lldp interface configuration] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_lldp_interface/tests/netconf/basic.yaml:4

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'name': u'ge-0/0/5', u'state': u'absent', u'purge': False, u'active': True, u'host': None, u'disable': True, u'timeout': 10, u'provider': {u'username': u'ansible', u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'netconf'}, u'aggregate': None, u'password': None, u'port': None, u'transport': None}}
[msg]=> <?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X49/junos" xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="urn:uuid:8f09b8ae-1dec-4670-9080-da68eaa28c31">
<load-configuration-results>
<rpc-error>
<error-type>protocol</error-type>
<error-tag>operation-failed</error-tag>
<error-severity>error</error-severity>
<error-message>syntax error</error-message>
<error-info>
<bad-element>lldp</bad-element>
</error-info>
</rpc-error>
<rpc-error>
<error-type>protocol</error-type>
<error-tag>operation-failed</error-tag>
<error-severity>error</error-severity>
<error-message>syntax error</error-message>
<error-info>
<bad-element>lldp</bad-element>
</error-info>
</rpc-error>
<load-error-count>2</load-error-count>
</load-configuration-results>
</rpc-reply>



############ FAILURE ############
[junos_vrf : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_vrf/tests/netconf/basic.yaml:23

[msg]=> The conditional check ''+   test-1' in result.diff.prepared' failed. The error was: error while evaluating conditional ('+   test-1' in result.diff.prepared): 'dict object' has no attribute 'diff'
[failed]=> True



############ FAILURE ############
[Has any previous test failed?] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/junos.yaml:182

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
prepare_ios_tests : Ensure we have loopback 888 for testing/Users/dnewswan/code/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

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
prepare_ios_tests : Ensure we have loopback 888 for testing/Users/dnewswan/code/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

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
prepare_ios_tests : Ensure we have loopback 888 for testing/Users/dnewswan/code/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

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
prepare_ios_tests : Ensure we have loopback 888 for testing/Users/dnewswan/code/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

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
prepare_ios_tests : Ensure we have loopback 888 for testing/Users/dnewswan/code/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

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
prepare_ios_tests : Ensure we have loopback 888 for testing/Users/dnewswan/code/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

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
prepare_ios_tests : Ensure we have loopback 888 for testing/Users/dnewswan/code/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

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
prepare_ios_tests : Ensure we have loopback 888 for testing/Users/dnewswan/code/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

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
Has any previous test failed?/Users/dnewswan/code/ansible/test/integration/ios.yaml:102

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
iosxr_config : assert/Users/dnewswan/code/ansible/test/integration/targets/iosxr_config/tests/cli/toplevel_nonidempotent.yaml:26

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == true
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
iosxr_template : configure device with config/Users/dnewswan/code/ansible/test/integration/targets/iosxr_template/tests/cli/backup.yaml:24

[msg]=> Unexpected failure during module execution.
[failed]=> True
[stdout]=> 



############ FAILURE ############
iosxr_banner : Set login/Users/dnewswan/code/ansible/test/integration/targets/iosxr_banner/tests/cli/basic-login.yaml:8

[_ansible_parsed]=> True
[commands]=> [u'banner login this is my login banner\nthat has a multiline\nstring']
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[rc]=> 1
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'text': u'this is my login banner\nthat has a multiline\nstring\n', u'state': u'present', u'host': None, u'timeout': None, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'host': u'iosxr01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'cli'}, u'password': None, u'banner': u'login', u'port': None}}
[msg]=> banner login this is my login banner
Enter TEXT message. End with the character 't'
that has a multiline
RP/0/0/CPU0:iosxr01(config)#string
                              ^
% Invalid input detected at '^' marker.
RP/0/0/CPU0:iosxr01(config)#



############ FAILURE ############
iosxr_logging : Set up host logging/Users/dnewswan/code/ansible/test/integration/targets/iosxr_logging/tests/cli/basic.yaml:2

[_ansible_parsed]=> False
[_ansible_no_log]=> False
[module_stderr]=> Traceback (most recent call last):
  File "/var/folders/ql/d9_087wj4cd10grvwzxmk4hc0000gn/T/ansible_gGC5A3/ansible_module_iosxr_logging.py", line 358, in <module>
    main()
  File "/var/folders/ql/d9_087wj4cd10grvwzxmk4hc0000gn/T/ansible_gGC5A3/ansible_module_iosxr_logging.py", line 344, in main
    have = map_config_to_obj(module)
  File "/var/folders/ql/d9_087wj4cd10grvwzxmk4hc0000gn/T/ansible_gGC5A3/ansible_module_iosxr_logging.py", line 240, in map_config_to_obj
    if match.group(1) in dest_group:
AttributeError: 'NoneType' object has no attribute 'group'

[changed]=> False
[module_stdout]=> 
[failed]=> True
[rc]=> 0
[msg]=> MODULE FAILURE



############ FAILURE ############
Has any previous test failed?/Users/dnewswan/code/ansible/test/integration/iosxr.yaml:94

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
nxos_command : Run show running-config bgp - should pass/Users/dnewswan/code/ansible/test/integration/targets/nxos_command/tests/nxapi/sanity.yaml:42

[status]=> -1
[_ansible_parsed]=> True
[_ansible_no_log]=> False
[url]=> http://nxos01:80/ins
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'url_password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'commands': [u'sh running-config bgp'], u'ssh_keyfile': None, u'timeout': 10, u'retries': 10, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'interval': 1, u'host': u'nxos01', u'match': u'all', u'url_username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'host': u'nxos01', u'timeout': 60, u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}, u'use_ssl': False, u'wait_for': None, u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}}
[msg]=> Connection failure: timed out



############ FAILURE ############
nxos_config : assert/Users/dnewswan/code/ansible/test/integration/targets/nxos_config/tests/cli/defaults.yaml:37

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == false
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
nxos_mtu : assert/Users/dnewswan/code/ansible/test/integration/targets/nxos_mtu/tests/cli/set_sysmtu.yaml:26

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == false
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
nxos_system : configure name_servers/Users/dnewswan/code/ansible/test/integration/targets/nxos_system/tests/cli/set_name_servers.yaml:13

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
nxos_interface : assert/Users/dnewswan/code/ansible/test/integration/targets/nxos_interface/tests/cli/set_state_absent.yaml:16

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == true
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
nxos_user : Create user/Users/dnewswan/code/ansible/test/integration/targets/nxos_user/tests/cli/basic.yaml:2

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'authorize': True, u'name': u'netend', u'roles': u'network-operator', u'state': u'present', u'provider': None, u'transport': u'cli'}}
[msg]=> Unsupported parameters for (nxos_user) module: authorize. Supported parameters include: aggregate,host,name,password,port,provider,purge,roles,ssh_keyfile,sshkey,state,timeout,transport,update_password,use_ssl,username,validate_certs



############ FAILURE ############
nxos_banner : assert/Users/dnewswan/code/ansible/test/integration/targets/nxos_banner/tests/cli/basic-no-exec.yaml:33

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == false
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
Has any previous test failed?/Users/dnewswan/code/ansible/test/integration/nxos.yaml:156

[msg]=> One or more tests failed, check log for details
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



```
