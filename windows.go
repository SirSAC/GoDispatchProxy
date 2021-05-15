// Made by SirSAC for Network.
package main
//
import (
	. "context"
	"log"
	. "net"
	"os/exec"
	"runtime"
)
//
func handle_internet(local_connection Conn, target_address string, processor_thread int, pipe_size int, try_count int, keep_alive bool, serial_order bool) {
	if serial_order == true {
		sync_group.Add(processor_thread)
	}
	load_balancer := get_load_balancer(serial_order)
	local_address, _ := ResolveTCPAddr("tcp", load_balancer.interface_address)
	network_dialer.LocalAddr = local_address
	remote_connection, err := network_dialer.DialContext(Background(), "tcp", target_address)
	if err != nil {
		try_again := 0
		for try_again < try_count {
			load_balancer := get_load_balancer(serial_order)
			local_address, _ := ResolveTCPAddr("tcp", load_balancer.interface_address)
			network_dialer.LocalAddr = local_address
			remote_connection, err := network_dialer.DialContext(Background(), "tcp", target_address)
			if err == nil {
				go handle_proxy(local_connection, remote_connection, local_address, target_address, pipe_size, keep_alive)
				if serial_order == true {
					defer sync_group.Wait()
				}
				return
			}
			try_again++
		}
		log.Println(string(COLOR_YELLOW), "[!]", local_address, "<=>", target_address, string(COLOR_RESET))
		local_connection.Write([]byte {5, NETWORK_UNREACHABLE, 0, 1, 0, 0, 0, 0, 0, 0})
		local_connection.Close()
		if serial_order == true {
			defer sync_group.Wait()
			defer sync_group.Done()
		}
		runtime.Goexit()
	}
	go handle_proxy(local_connection, remote_connection, local_address, target_address, pipe_size, keep_alive)
	if serial_order == true {
		defer sync_group.Wait()
	}
}
//
func execute_command(mtu_size string, option_setting bool) {
	exec.Command("cmd.exe", "/c", "wmic.exe process where 'name='DisPro.bin'' call setpriority realtime").Run()
	exec.Command("cmd.exe", "/c", "net.exe stop /y RemoteAccess").Run()
	exec.Command("cmd.exe", "/c", "netsh.exe interface ipv6 set interface interface=1 metric=1 store=active").Run()
	exec.Command("cmd.exe", "/c", "netsh.exe interface ipv6 set subinterface interface=1 mtu=", mtu_size, "store=active").Run()
	exec.Command("cmd.exe", "/c", "netsh.exe interface ipv6 set address interface=1 type=anycast store=active").Run()
	exec.Command("cmd.exe", "/c", "netsh.exe interface ipv6 set address interface=1 address=::1 type=unicast validlifetime=infinite preferredlifetime=infinite store=active").Run()
	exec.Command("cmd.exe", "/c", "netsh.exe interface ipv4 set interface interface=1 metric=1 store=active").Run()
	exec.Command("cmd.exe", "/c", "netsh.exe interface ipv4 set subinterface interface=1 mtu=", mtu_size, "store=active").Run()
	exec.Command("cmd.exe", "/c", "netsh.exe interface ipv4 set address name=1 source=dhcp type=anycast store=active").Run()
	exec.Command("cmd.exe", "/c", "netsh.exe interface ipv4 set address name=1 source=static address=127.0.0.1 mask=255.255.255.255 gwmetric=1 type=unicast store=active").Run()
	if option_setting == true {
		exec.Command("cmd.exe", "/c", "sc.exe config RemoteAccess start=disabled").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe interface ipv6 set dnsservers name=1 source=dhcp register=both validate=yes").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe interface ipv4 set dnsservers name=1 source=dhcp register=both validate=yes").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe interface tcp set global ecncapability=enabled").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe interface tcp set global fastopen=enabled").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe interface tcp set global timestamps=enabled").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe interface portproxy set v4tov6 listenport=8118 connectaddress=::1 connectport=8118 listenaddress=127.0.0.1").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe interface portproxy set v4tov6 listenport=1080 connectaddress=::1 connectport=1080 listenaddress=127.0.0.1").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe lan set autoconfig enabled=yes interface=*").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe lan set allowexplicitcreds allow=yes").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe wlan set randomization enabled=yes interface=*").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe wlan set autoconfig enabled=yes interface=*").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe wlan set hostednetwork mode=allow").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe wlan set blockednetworks display=show").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe wlan set allowexplicitcreds allow=yes").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe winhttp set proxy proxy-server=[::1]:8118").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe http add iplisten ipaddress=::1").Run()
		exec.Command("cmd.exe", "/c", "netsh.exe http add iplisten ipaddress=127.0.0.1").Run()
	}
	exec.Command("cmd.exe", "/c", "ipconfig.exe /registerdns").Run()
	exec.Command("cmd.exe", "/c", "ipconfig.exe /flushdns").Run()
}
