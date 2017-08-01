# CI Test Report Card 

## Test Summary 
### Passed: 
* [iosxr](#iosxr)

### Failed: 
* [junos](#junos)
* [ios](#ios)
* [nxos](#nxos)

# Results 
## junos
*The following tests have failed:* 
[_See the .log files for details_](./junos.log) 
```
############ FAILURE ############
[junos_netconf : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_netconf/tests/cli/changeport.yaml:49

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.failed == true
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[junos_logging : Configure logging parameters] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_logging/tests/netconf/basic.yaml:196

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'files': 40, u'ssh_keyfile': None, u'src_addr': None, u'name': None, u'facility': None, u'dest': None, u'level': None, u'host': None, u'purge': False, u'aggregate': None, u'state': u'present', u'timeout': 10, u'provider': {u'username': u'ansible', u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'netconf'}, u'active': True, u'rotate_frequency': 20, u'password': None, u'port': None, u'transport': None, u'size': 65536}}
[msg]=> <?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X49/junos" xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="urn:uuid:b149e2e1-3722-40a4-8181-44b877372894">
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
[junos_linkagg : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_linkagg/tests/netconf/basic.yaml:154

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> '<disable/>' not in config.xml
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[junos_lldp : setup - Disable lldp and remove it's configuration] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_lldp/tests/netconf/basic.yaml:4

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'interval': None, u'state': u'absent', u'purge': False, u'host': None, u'disable': True, u'transport': None, u'timeout': 10, u'provider': {u'username': u'ansible', u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'netconf'}, u'active': True, u'password': None, u'hold_multiplier': None, u'port': None, u'transmit_delay': None}}
[msg]=> <?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X49/junos" xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="urn:uuid:f8763550-fdaa-41aa-8fdc-585ac474ecef">
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
[msg]=> <?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X49/junos" xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="urn:uuid:a310aa24-5a2f-41cb-b5fa-5ee10df98880">
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

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> '+   test-1' in result.diff.prepared
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[Has any previous test failed?] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/junos.yaml:179

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

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[ios_facts : test getting all facts] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/ios_facts/tests/cli/all_facts.yaml:5

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[err]=> [Errno 2] No such file or directory
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'authorize': None, u'ssh_keyfile': None, u'auth_pass': None, u'host': None, u'gather_subset': [u'all'], u'timeout': None, u'provider': None, u'password': None, u'port': None}}
[msg]=> unable to connect to socket



############ FAILURE ############
[ios_user : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/ios_user/tests/cli/basic.yaml:12

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == true
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



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
*All tests passed!*

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
[nxos_interface : setup] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_interface/tests/nxapi/set_state_absent.yaml:4

[status]=> -1
[_ansible_parsed]=> True
[_ansible_no_log]=> False
[url]=> http://nxos01:80/ins
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'force': False, u'diff_against': None, u'replace': u'line', u'running_config': None, u'use_ssl': False, u'port': 80, u'transport': u'nxapi', u'before': None, u'parents': None, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'host': u'nxos01', u'timeout': 60, u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}, u'save': False, u'match': u'line', u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'timeout': 10, u'src': None, u'after': None, u'host': u'nxos01', u'ssh_keyfile': None, u'defaults': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'diff_ignore_lines': None, u'url_password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'save_when': u'never', u'backup': False, u'lines': [u'interface Loopback1'], u'intended_config': None, u'url_username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True}}
[msg]=> Connection failure: timed out



############ FAILURE ############
[nxos_user : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_user/tests/cli/basic.yaml:10

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == true
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



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
[date]=> Sun, 06 Aug 2017 00:21:00 GMT
[content-type]=> text/html
[www-authenticate]=> Basic realm="Secure Zone"
[msg]=> HTTP Error 401: Unauthorized



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
[features_currently_supported]=> {u'mvpn': u'disabled', u'pim': u'enabled', u'ldp': u'disabled', u'hsrp_engine': u'enabled', u'itd': u'disabled', u'scheduler': u'disabled', u'fabric_mcast': u'disabled', u'glbp': u'disabled', u'private-vlan': u'disabled', u'eth_port_sec': u'disabled', u'mvrp': u'disabled', u'pong': u'disabled', u'dot1x': u'disabled', u'ptp': u'disabled', u'lldp': u'disabled', u'mpls_te': u'disabled', u'evmed': u'disabled', u'elo': u'disabled', u'imp': u'disabled', u'vmtracker': u'disabled', u'interface-vlan': u'disabled', u'cable-management': u'disabled', u'vtp': u'disabled', u'vrrp': u'disabled', u'ldap': u'disabled', u'privilege': u'disabled', u'dhcp': u'disabled', u'sla_responder': u'disabled', u'fabric-access': u'enabled', u'pim6': u'disabled', u'bulkstat': u'disabled', u'vrrpv3': u'disabled', u'wccp': u'disabled', u'rise': u'disabled', u'l3vpn': u'disabled', u'vni': u'enabled', u'onep': u'disabled', u'hmm': u'disabled', u'nve': u'enabled', u'nxapi': u'enabled', u'rip': u'disabled', u'vpc': u'enabled', u'sftpServer': u'disabled', u'netflow': u'disabled', u'l2vpn': u'disabled', u'msdp': u'disabled', u'otv': u'disabled', u'cts': u'disabled', u'smart-channel': u'disabled', u'ngoam': u'disabled', u'isis': u'disabled', u'evc': u'disabled', u'pbr': u'disabled', u'lisp': u'disabled', u'tacacs': u'disabled', u'evb': u'disabled', u'bfd': u'disabled', u'sla_sender': u'disabled', u'ospfv3': u'disabled', u'bgp': u'disabled', u'tunnel': u'disabled', u'eigrp': u'disabled', u'vnseg_vlan': u'disabled', u'msrp': u'disabled', u'scpServer': u'enabled', u'ospf': u'enabled', u'telnetServer': u'enabled', u'sshServer': u'enabled', u'bfd_app': u'disabled'}
[msg]=> Invalid feature name.



############ FAILURE ############
[Has any previous test failed?] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/nxos.yaml:236

[msg]=> One or more tests failed, check log for details
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



```
