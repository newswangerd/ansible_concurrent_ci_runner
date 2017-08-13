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
[junos_linkagg : configure linkagg] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_linkagg/tests/netconf/basic.yaml:15

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[err]=> [Errno 2] No such file or directory
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'description': u'configured by junos_linkagg', u'provider': {u'username': None, u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': None, u'port': None, u'transport': u'netconf'}, u'host': None, u'device_count': 4, u'purge': False, u'aggregate': None, u'state': u'present', u'mode': u'active', u'timeout': None, u'min_links': None, u'active': True, u'password': None, u'members': [u'ge-0/0/6', u'ge-0/0/7'], u'port': None, u'transport': None, u'name': u'ae0'}}
[msg]=> unable to connect to socket



############ FAILURE ############
[junos_lldp : setup - Disable lldp and remove it's configuration] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/junos_lldp/tests/netconf/basic.yaml:4

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'interval': None, u'state': u'absent', u'purge': False, u'host': None, u'transport': None, u'timeout': None, u'provider': {u'username': None, u'ssh_keyfile': None, u'host': u'vsrx01', u'timeout': None, u'password': None, u'port': None, u'transport': u'netconf'}, u'active': True, u'password': None, u'hold_multiplier': None, u'port': None, u'transmit_delay': None}}
[msg]=> <?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X49/junos" xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="urn:uuid:f9729b9b-d132-4b73-832d-d29e8cb15c08">
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
[msg]=> <?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X49/junos" xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="urn:uuid:762f0f96-315d-444e-be8c-906ac4cd97e5">
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



Passed: junos_banner, junos_config, junos_system, junos_l3_interface, junos_vrf, junos_netconf, junos_logging, junos_interface, junos_facts, junos_rpc, junos_static_route, junos_user, junos_command, junos_vlan, 
Failed: junos_linkagg, junos_lldp, junos_lldp_interface, 
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
[iosxr_config : configure top level command] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/iosxr_config/tests/cli/toplevel_nonidempotent.yaml:9

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[err]=> [Errno 2] No such file or directory
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'comment': u'configured by iosxr_config', u'username': None, u'commands': [u'banner motd "hello world"', u'hostname foo'], u'ssh_keyfile': None, u'force': False, u'src': None, u'admin': False, u'backup': False, u'after': None, u'lines': [u'banner motd "hello world"', u'hostname foo'], u'replace': u'line', u'host': None, u'parents': None, u'timeout': None, u'provider': None, u'password': None, u'config': None, u'port': None, u'match': u'strict', u'before': None}}
[msg]=> unable to connect to socket



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



Passed: iosxr_user, iosxr_logging, iosxr_facts, iosxr_system, iosxr_command, iosxr_banner, 
Failed: iosxr_interface, iosxr_config, 
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
[nxos_evpn_vni : Enable nv overlay evpn] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_evpn_vni/tests/nxapi/sanity.yaml:18

[status]=> -1
[_ansible_parsed]=> True
[_ansible_no_log]=> False
[url]=> http://nxos01:80/ins
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'force': False, u'diff_against': None, u'replace': u'line', u'running_config': None, u'use_ssl': False, u'port': 80, u'transport': u'nxapi', u'before': None, u'parents': None, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'host': u'nxos01', u'timeout': 60, u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}, u'save': False, u'match': u'none', u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'timeout': 60, u'src': None, u'after': None, u'host': u'nxos01', u'ssh_keyfile': None, u'defaults': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'diff_ignore_lines': None, u'url_password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'save_when': u'never', u'backup': False, u'lines': [u'nv overlay evpn'], u'intended_config': None, u'url_username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True}}
[msg]=> An unknown error occurred: ''



############ FAILURE ############
[prepare_nxos_tests : Collect interface list] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/prepare_nxos_tests/tasks/main.yml:39

[status]=> -1
[_ansible_parsed]=> True
[_ansible_no_log]=> False
[url]=> http://nxos01:80/ins
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'url_password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'commands': [u'show interface brief | json'], u'ssh_keyfile': None, u'url_username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'retries': 10, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'interval': 1, u'host': u'nxos01', u'transport': u'nxapi', u'timeout': 60, u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'host': u'nxos01', u'timeout': 60, u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}, u'use_ssl': False, u'wait_for': None, u'validate_certs': True, u'port': 80, u'match': u'all'}}
[msg]=> Request failed: <urlopen error [Errno 60] Operation timed out>



############ FAILURE ############
[nxos_feature : assert] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_feature/tests/nxapi/configure.yaml:28

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
[nxos_system : configure domain_list] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_system/tests/cli/set_domain_list.yaml:12

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'ssh_keyfile': None, u'lookup_source': None, u'domain_search': [u'ansible.com', u'redhat.com'], u'hostname': None, u'system_mtu': None, u'domain_name': None, u'domain_lookup': None, u'host': None, u'state': u'present', u'timeout': None, u'provider': None, u'name_servers': None, u'use_ssl': None, u'password': None, u'validate_certs': None, u'port': None, u'transport': u'cli'}}
[msg]=> ip domain-list redhat.com
           ^
% Invalid command at '^' marker.
nxos01# 



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
[nxos_vrf_interface : Remove Idempotence] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_vrf_interface/tests/common/sanity.yaml:47

[status]=> -1
[_ansible_parsed]=> True
[_ansible_no_log]=> False
[url]=> http://nxos01:80/ins
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'url_password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'timeout': 60, u'state': u'absent', u'host': u'nxos01', u'vrf': u'ntc', u'url_username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'host': u'nxos01', u'timeout': 60, u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}, u'interface': u'Ethernet2/1', u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}}
[msg]=> Connection failure: timed out



############ FAILURE ############
[nxos_portchannel : Check Idempotence] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_portchannel/tests/common/sanity.yaml:53

[_ansible_parsed]=> True
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'force': u'true', u'ssh_keyfile': None, u'group': u'99', u'timeout': 60, u'validate_certs': None, u'password': None, u'provider': None, u'host': None, u'include_defaults': u'False', u'state': u'present', u'mode': u'active', u'members': [u'Ethernet2/1', u'Ethernet2/2'], u'min_links': None, u'use_ssl': None, u'save': False, u'config': None, u'port': None, u'transport': u'cli'}}
[msg]=> interface Ethernet2/2
          ^
% Invalid command at '^' marker.
nxos01# 



############ FAILURE ############
[nxos_portchannel : Disable feature LACP] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_portchannel/tests/common/sanity.yaml:87

[_ansible_parsed]=> True
[invalid_feature]=> lacp
[_ansible_no_log]=> False
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'username': None, u'password': None, u'ssh_keyfile': None, u'validate_certs': None, u'feature': u'lacp', u'include_defaults': None, u'state': u'disabled', u'timeout': 60, u'provider': None, u'host': None, u'use_ssl': None, u'save': None, u'config': None, u'port': None, u'transport': u'cli'}}
[features_currently_supported]=> {u'mvpn': u'disabled', u'pim': u'disabled', u'ldp': u'disabled', u'hsrp_engine': u'disabled', u'itd': u'disabled', u'scheduler': u'disabled', u'fabric_mcast': u'disabled', u'glbp': u'disabled', u'private-vlan': u'disabled', u'eth_port_sec': u'disabled', u'mvrp': u'disabled', u'pong': u'disabled', u'dot1x': u'disabled', u'ptp': u'disabled', u'lldp': u'disabled', u'mpls_te': u'disabled', u'evmed': u'disabled', u'elo': u'disabled', u'imp': u'disabled', u'vmtracker': u'disabled', u'interface-vlan': u'disabled', u'cable-management': u'disabled', u'vtp': u'disabled', u'vrrp': u'disabled', u'ldap': u'disabled', u'privilege': u'disabled', u'dhcp': u'disabled', u'sla_responder': u'disabled', u'fabric-access': u'disabled', u'pim6': u'disabled', u'bulkstat': u'disabled', u'vrrpv3': u'disabled', u'wccp': u'disabled', u'rise': u'disabled', u'l3vpn': u'disabled', u'vni': u'disabled', u'onep': u'disabled', u'hmm': u'enabled', u'nve': u'enabled', u'nxapi': u'enabled', u'rip': u'enabled', u'vpc': u'enabled', u'sftpServer': u'disabled', u'netflow': u'disabled', u'l2vpn': u'disabled', u'msdp': u'disabled', u'otv': u'disabled', u'cts': u'disabled', u'smart-channel': u'disabled', u'ngoam': u'disabled', u'isis': u'enabled', u'evc': u'disabled', u'pbr': u'disabled', u'lisp': u'disabled', u'tacacs': u'disabled', u'evb': u'disabled', u'bfd': u'disabled', u'sla_sender': u'disabled', u'ospfv3': u'enabled', u'bgp': u'enabled', u'tunnel': u'disabled', u'eigrp': u'enabled', u'vnseg_vlan': u'disabled', u'msrp': u'disabled', u'scpServer': u'enabled', u'ospf': u'disabled', u'telnetServer': u'enabled', u'sshServer': u'enabled', u'bfd_app': u'disabled'}
[msg]=> Invalid feature name.



############ FAILURE ############
[prepare_nxos_tests : Gather platform info] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/prepare_nxos_tests/tasks/main.yml:58

[msg]=> unable to open shell. Please see: https://docs.ansible.com/ansible/network_debug_troubleshooting.html#unable-to-open-shell
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



############ FAILURE ############
[nxos_switchport : tag vlan Idempotence] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/targets/nxos_switchport/tests/common/sanity.yaml:83

[status]=> -1
[_ansible_parsed]=> True
[_ansible_no_log]=> False
[url]=> http://nxos01:80/ins
[changed]=> False
[failed]=> True
[invocation]=> {u'module_args': {u'native_vlan': u'10', u'access_vlan': None, u'ssh_keyfile': None, u'timeout': 60, u'url_password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'host': u'nxos01', u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'state': u'present', u'trunk_allowed_vlans': None, u'trunk_vlans': u'2-50', u'mode': u'trunk', u'url_username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'provider': {u'username': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'ssh_keyfile': None, u'host': u'nxos01', u'timeout': 60, u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}, u'interface': u'Ethernet2/1', u'use_ssl': False, u'password': u'VALUE_SPECIFIED_IN_NO_LOG_PARAMETER', u'validate_certs': True, u'port': 80, u'transport': u'nxapi'}}
[msg]=> An unknown error occurred: ''



############ FAILURE ############
[Has any previous test failed?] 
/Users/dnewswan/code/concurrent_ci/ansible/test/integration/nxos.yaml:293

[msg]=> One or more tests failed, check log for details
[failed]=> True
[changed]=> False
[_ansible_no_log]=> False



Passed: nxos_acl, nxos_bgp, nxos_acl_interface, nxos_interface, nxos_vrrp, nxos_vtp_version, nxos_vrf, nxos_facts, nxos_bgp_neighbor, nxos_bgp_af, nxos_logging, nxos_vtp_domain, nxos_vtp_password, nxos_vxlan_vtep, nxos_interface_ospf, nxos_user, nxos_rollback, nxos_ospf, nxos_nxapi, 
Failed: nxos_config, nxos_feature, nxos_switchport, nxos_command, nxos_system, nxos_mtu, nxos_evpn_vni, nxos_vrf_interface, nxos_portchannel, nxos_banner, 
```
