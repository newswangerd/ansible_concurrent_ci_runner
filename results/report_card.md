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

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'retries': 10, u'commands': [u'show version', u'show interfaces lo0'], u'ssh_keyfile': None, u'rpcs': None, u'format': u'xml', u'interval': 1, u'wait_for': [u'result[0].rpc-reply.software-information.host-name contains vsrx01', u'result[1].rpc-reply.interface-information.physical-interface.name contains lo0'], u'display': u'xml', u'host': None, u'match': u'all', u'timeout': None, u'provider': {u'username': None, u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': None, u'port': None, u'transport': u'netconf'}, u'password': None, u'port': None, u'transport': None}}
[msg]=> jxmlease is required but does not appear to be installed. It can be installed using `pip install jxmlease`



############ FAILURE ############
[junos_netconf : Ensure we can communicate over 8080] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_netconf/tests/cli/changeport.yaml:35

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'retries': 10, u'commands': None, u'ssh_keyfile': None, u'rpcs': [u'get-software-information'], u'password': None, u'interval': 1, u'display': None, u'host': None, u'match': u'all', u'timeout': None, u'provider': {u'username': None, u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': None, u'port': 8080, u'transport': u'netconf'}, u'wait_for': None, u'port': 8080, u'transport': None}}
[msg]=> jxmlease is required but does not appear to be installed. It can be installed using `pip install jxmlease`



############ FAILURE ############
[junos_linkagg : configure linkagg] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_linkagg/tests/netconf/basic.yaml:15

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[err]=> [Errno 2] No such file or directory
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'description': u'configured by junos_linkagg', u'min_links': None, u'host': None, u'bundle': u'ae0', u'collection': None, u'purge': None, u'state': u'present', u'disable': False, u'device_count': 4, u'timeout': None, u'provider': {u'username': None, u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': None, u'port': None, u'transport': u'netconf'}, u'active': True, u'mode': u'active', u'password': None, u'members': [u'ge-0/0/6', u'ge-0/0/7'], u'port': None, u'transport': None, u'name': u'ge-0/0/7'}}
[msg]=> unable to connect to socket



############ FAILURE ############
[junos_lldp : setup - Disable lldp and remove it's configuration] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_lldp/tests/netconf/basic.yaml:4

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'interval': None, u'state': u'absent', u'purge': False, u'host': None, u'disable': True, u'transport': None, u'timeout': None, u'provider': {u'username': None, u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': None, u'port': None, u'transport': u'netconf'}, u'active': True, u'password': None, u'hold_multiplier': None, u'port': None, u'transmit_delay': None}}
[msg]=> <?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X49/junos" xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="urn:uuid:5fd8aabd-5858-4403-bf67-e705c564b6a3">
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
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'name': u'ge-0/0/5', u'state': u'absent', u'purge': False, u'active': True, u'host': None, u'disable': True, u'timeout': None, u'provider': {u'username': None, u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': None, u'port': None, u'transport': u'netconf'}, u'aggregate': None, u'password': None, u'port': None, u'transport': None}}
[msg]=> <?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X49/junos" xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="urn:uuid:2c36f326-a9bb-4336-8c6c-499b01f39f71">
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
[Has any previous test failed?] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/junos.yaml:174

[msg]=> One or more tests failed, check log for details
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



Passed: junos_banner, junos_config, junos_system, junos_l3_interface, junos_vrf, junos_logging, junos_interface, junos_facts, junos_rpc, junos_static_route, junos_user, junos_vlan, 
Failed: junos_linkagg, junos_lldp, junos_netconf, junos_lldp_interface, junos_command, 
```
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
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/ios.yaml:101

[msg]=> One or more tests failed, check log for details
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



Passed: ios_interface, ios_command, ios_logging, ios_banner, ios_static_route, ios_system, 
Failed: ios_facts, ios_user, 
```
## iosxr
*The following tests have failed:* 
[_See the .log files for details_](./iosxr.log) 
```
############ FAILURE ############
[iosxr_banner : Set motd] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/iosxr_banner/tests/cli/basic-motd.yaml:8

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[err]=> [Errno 2] No such file or directory
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'text': u'this is my motd banner\nthat has a multiline\nstring\n', u'state': u'present', u'host': None, u'timeout': None, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'host': u'iosxr01', u'timeout': None, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'port': None, u'transport': u'cli'}, u'password': None, u'banner': u'motd', u'port': None}}
[msg]=> unable to connect to socket



############ FAILURE ############
[Has any previous test failed?] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/iosxr.yaml:94

[msg]=> One or more tests failed, check log for details
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



Passed: iosxr_user, iosxr_interface, iosxr_logging, iosxr_facts, iosxr_system, iosxr_config, iosxr_command, 
Failed: iosxr_banner, 
```
## nxos
*The following tests have failed:* 
[_See the .log files for details_](./nxos.log) 
```
############ FAILURE ############
[nxos_command : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_command/tests/nxapi/sanity.yaml:49

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> '65535' in result.stdout[0]
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[nxos_config : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_config/tests/cli/save.yaml:21

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == true
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[nxos_evpn_vni : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_evpn_vni/tests/nxapi/sanity.yaml:54

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == true
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[prepare_nxos_tests : Collect interface list] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/prepare_nxos_tests/tasks/main.yml:21

[status]=> -1
[_ansible_parsed]=> True
[_ansible_no_log]=> False
[url]=> http://nxos01:80/ins
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'url_password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'commands': [u'show interface brief | json'], u'ssh_keyfile': None, u'url_username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'retries': 10, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'interval': 1, u'host': u'nxos01', u'transport': u'nxapi', u'timeout': 60, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'host': u'nxos01', u'timeout': 60, u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}, u'use_ssl': False, u'wait_for': None, u'validate_certs': True, u'port': 80, u'match': u'all'}}
[msg]=> An unknown error occurred: ''



############ FAILURE ############
[nxos_system : configure name_servers] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_system/tests/cli/set_name_servers.yaml:13

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'lookup_source': None, u'domain_search': None, u'hostname': None, u'system_mtu': None, u'domain_name': None, u'name_servers': [u'1.1.1.1', u'2.2.2.2', u'3.3.3.3'], u'host': None, u'state': u'present', u'timeout': None, u'provider': None, u'domain_lookup': None, u'use_ssl': None, u'password': None, u'validate_certs': None, u'port': None, u'transport': u'cli'}}
[msg]=> no ip name-server use-vrf
                                  ^
% Invalid ip address at '^' marker.
nxos01(config)# 



############ FAILURE ############
[nxos_interface : setup] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_interface/tests/nxapi/set_state_present.yaml:4

[_ansible_parsed]=> True
[code]=> 400
[clierror]=> Error: Invalid range: loopback1

[_ansible_no_log]=> False
[url]=> http://nxos01:80/ins
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'force': False, u'diff_against': None, u'replace': u'line', u'running_config': None, u'use_ssl': False, u'port': 80, u'transport': u'nxapi', u'before': None, u'parents': None, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'host': u'nxos01', u'timeout': 60, u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}, u'save': False, u'match': u'line', u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'timeout': 60, u'src': None, u'after': None, u'host': u'nxos01', u'ssh_keyfile': None, u'defaults': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'diff_ignore_lines': None, u'url_password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'save_when': u'never', u'backup': False, u'lines': [u'no interface Loopback1'], u'intended_config': None, u'url_username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True}}
[output]=> {u'msg': u'CLI execution error', u'code': u'400', u'clierror': u'Error: Invalid range: loopback1\n'}
[msg]=> CLI execution error



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
[nxos_acl_interface : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_acl_interface/tests/cli/sanity.yaml:53

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == false
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[nxos_acl_interface : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_acl_interface/tests/nxapi/sanity.yaml:53

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
[invocation]=> {u'module_args': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'url_password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'vtp_password': None, u'ssh_keyfile': None, u'timeout': 60, u'state': u'present', u'host': u'nxos01', u'url_username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'host': u'nxos01', u'timeout': 60, u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}, u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}}
[date]=> Tue, 08 Aug 2017 01:56:15 GMT
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
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/nxos.yaml:254

[msg]=> One or more tests failed, check log for details
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



Passed: nxos_acl, nxos_evpn_global, nxos_feature, nxos_vrrp, nxos_vtp_version, nxos_nxapi, nxos_vrf, nxos_facts, nxos_logging, nxos_vtp_domain, nxos_vxlan_vtep, nxos_vlan, nxos_user, nxos_rollback, nxos_vrf_interface, nxos_switchport, 
Failed: nxos_config, nxos_acl_interface, nxos_interface, nxos_command, nxos_system, nxos_vtp_password, nxos_evpn_vni, nxos_portchannel, nxos_banner, 
```
