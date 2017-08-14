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
[junos_config : setup] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_config/tests/netconf/backup.yaml:4

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[err]=> [Errno 2] No such file or directory
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'comment': u'configured by junos_config', u'username': None, u'ssh_keyfile': None, u'rollback': None, u'timeout': None, u'confirm': 0, u'lines': [u'set system host-name vsrx01', u'delete interfaces lo0'], u'update': u'merge', u'replace': None, u'confirm_commit': False, u'zeroize': False, u'src_format': None, u'provider': {u'username': None, u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': None, u'port': None, u'transport': u'netconf'}, u'host': None, u'src': None, u'password': None, u'backup': False, u'port': None, u'transport': None}}
[msg]=> unable to connect to socket



############ FAILURE ############
[junos_linkagg : configure linkagg] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_linkagg/tests/netconf/basic.yaml:15

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[err]=> [Errno 2] No such file or directory
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'name': u'ae0', u'provider': {u'username': None, u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': None, u'port': None, u'transport': u'netconf'}, u'state': u'present', u'device_count': 4, u'members': [u'ge-0/0/6', u'ge-0/0/7'], u'active': True, u'host': None, u'mode': u'active', u'timeout': None, u'min_links': None, u'aggregate': None, u'password': None, u'port': None, u'transport': None, u'description': u'configured by junos_linkagg'}}
[msg]=> unable to connect to socket



############ FAILURE ############
[junos_lldp : setup - Disable lldp and remove it's configuration] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_lldp/tests/netconf/basic.yaml:4

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'interval': None, u'state': u'absent', u'purge': False, u'host': None, u'transport': None, u'timeout': None, u'provider': {u'username': None, u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': None, u'port': None, u'transport': u'netconf'}, u'active': True, u'password': None, u'hold_multiplier': None, u'port': None, u'transmit_delay': None}}
[msg]=> <?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X49/junos" xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="urn:uuid:9ac9dcd3-6e44-4702-bada-7ce8656ebb25">
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
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'name': u'ge-0/0/5', u'state': u'absent', u'purge': False, u'active': True, u'host': None, u'timeout': None, u'provider': {u'username': None, u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': None, u'port': None, u'transport': u'netconf'}, u'aggregate': None, u'password': None, u'port': None, u'transport': None}}
[msg]=> <?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X49/junos" xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="urn:uuid:fd77a7c8-68c7-4670-be50-a50ac9e9e671">
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



Passed: junos_banner, junos_system, junos_l3_interface, junos_vrf, junos_netconf, junos_logging, junos_interface, junos_facts, junos_rpc, junos_static_route, junos_user, junos_command, junos_vlan, 
Failed: junos_linkagg, junos_lldp, junos_config, junos_lldp_interface, 
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
[prepare_ios_tests : Ensure we have loopback 888 for testing] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/prepare_ios_tests/tasks/main.yml:3

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



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



Passed: ios_interface, ios_banner, ios_command, ios_logging, ios_static_route, ios_facts, 
Failed: ios_user, 
```
## iosxr
*The following tests have failed:* 
[_See the .log files for details_](./iosxr.log) 
```
############ FAILURE ############
[iosxr_interface : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/iosxr_interface/tests/cli/basic.yaml:233

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == true
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[Has any previous test failed?] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/iosxr.yaml:94

[msg]=> One or more tests failed, check log for details
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



Passed: iosxr_user, iosxr_logging, iosxr_facts, iosxr_system, iosxr_config, iosxr_command, iosxr_banner, 
Failed: iosxr_interface, 
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
[nxos_evpn_global : Enable nv overlay evpn] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_evpn_global/tests/cli/sanity.yaml:26

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'password': None, u'ssh_keyfile': None, u'validate_certs': None, u'include_defaults': None, u'host': None, u'timeout': None, u'provider': None, u'nv_overlay_evpn': True, u'use_ssl': None, u'save': None, u'config': None, u'port': None, u'transport': u'cli'}}
[msg]=> timeout trying to send command: show running-config



############ FAILURE ############
[nxos_evpn_global : enable nxapi] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_evpn_global/tasks/nxapi.yaml:11

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'force': False, u'diff_against': None, u'replace': u'line', u'running_config': None, u'save_when': u'never', u'use_ssl': None, u'port': None, u'transport': u'cli', u'before': None, u'parents': None, u'provider': None, u'save': False, u'match': u'line', u'username': None, u'timeout': None, u'after': None, u'host': None, u'password': None, u'diff_ignore_lines': None, u'src': None, u'ssh_keyfile': None, u'backup': False, u'lines': [u'feature nxapi', u'nxapi http port 80'], u'intended_config': None, u'defaults': False, u'validate_certs': None}}
[msg]=> timeout trying to send command: show running-config all



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
[nxos_banner : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_banner/tests/cli/basic-no-exec.yaml:33

[_ansible_no_log]=> False
[changed]=> False
[assertion]=> result.changed == false
[_ansible_verbose_always]=> True
[failed]=> True
[evaluated_to]=> False



############ FAILURE ############
[nxos_vtp_version : enable feature vtp] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_vtp_version/tests/nxapi/sanity.yaml:5

[status]=> -1
[_ansible_parsed]=> True
[_ansible_no_log]=> False
[url]=> http://nxos01:80/ins
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'url_password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'url_username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'feature': u'vtp', u'include_defaults': None, u'state': u'enabled', u'timeout': 60, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'host': u'nxos01', u'timeout': 60, u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}, u'host': u'nxos01', u'use_ssl': False, u'save': None, u'config': None, u'port': 80, u'transport': u'nxapi'}}
[msg]=> Connection failure: timed out



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
[features_currently_supported]=> {u'mvpn': u'disabled', u'pim': u'disabled', u'ldp': u'disabled', u'hsrp_engine': u'disabled', u'itd': u'disabled', u'scheduler': u'disabled', u'fabric_mcast': u'disabled', u'glbp': u'disabled', u'private-vlan': u'disabled', u'eth_port_sec': u'disabled', u'mvrp': u'disabled', u'pong': u'disabled', u'dot1x': u'disabled', u'ptp': u'disabled', u'lldp': u'disabled', u'mpls_te': u'disabled', u'evmed': u'disabled', u'elo': u'disabled', u'imp': u'disabled', u'vmtracker': u'disabled', u'interface-vlan': u'disabled', u'cable-management': u'disabled', u'vtp': u'disabled', u'vrrp': u'disabled', u'ldap': u'disabled', u'privilege': u'disabled', u'dhcp': u'disabled', u'sla_responder': u'disabled', u'fabric-access': u'disabled', u'pim6': u'disabled', u'bulkstat': u'disabled', u'vrrpv3': u'disabled', u'wccp': u'disabled', u'rise': u'disabled', u'l3vpn': u'disabled', u'vni': u'disabled', u'onep': u'disabled', u'hmm': u'enabled', u'nve': u'enabled', u'nxapi': u'enabled', u'rip': u'enabled', u'vpc': u'enabled', u'sftpServer': u'disabled', u'netflow': u'disabled', u'l2vpn': u'disabled', u'msdp': u'disabled', u'otv': u'disabled', u'cts': u'disabled', u'smart-channel': u'disabled', u'ngoam': u'disabled', u'isis': u'enabled', u'evc': u'disabled', u'pbr': u'disabled', u'lisp': u'disabled', u'tacacs': u'disabled', u'evb': u'disabled', u'bfd': u'disabled', u'sla_sender': u'disabled', u'ospfv3': u'enabled', u'bgp': u'disabled', u'tunnel': u'disabled', u'eigrp': u'enabled', u'vnseg_vlan': u'disabled', u'msrp': u'disabled', u'scpServer': u'enabled', u'ospf': u'enabled', u'telnetServer': u'enabled', u'sshServer': u'enabled', u'bfd_app': u'disabled'}
[msg]=> Invalid feature name.



############ FAILURE ############
[Has any previous test failed?] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/nxos.yaml:293

[msg]=> One or more tests failed, check log for details
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



Passed: nxos_acl, nxos_bgp, nxos_feature, nxos_acl_interface, nxos_interface, nxos_vrrp, nxos_nxapi, nxos_vrf, nxos_facts, nxos_bgp_neighbor, nxos_system, nxos_bgp_af, nxos_logging, nxos_vtp_domain, nxos_vtp_password, nxos_vxlan_vtep, nxos_interface_ospf, nxos_vlan, nxos_user, nxos_evpn_vni, nxos_rollback, nxos_vrf_interface, nxos_ospf, nxos_switchport, 
Failed: nxos_config, nxos_evpn_global, nxos_vtp_version, nxos_command, nxos_mtu, nxos_portchannel, nxos_banner, 
```
