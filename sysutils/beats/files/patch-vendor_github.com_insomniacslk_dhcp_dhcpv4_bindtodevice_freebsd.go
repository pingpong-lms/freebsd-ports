--- /dev/null	2019-03-06 16:46:20.128821000 +0100
+++ vendor/github.com/insomniacslk/dhcp/dhcpv4/bindtodevice_freebsd.go	2019-03-06 16:43:25.453458000 +0100
@@ -0,0 +1,14 @@
+package dhcpv4
+
+import (
+	"net"
+	"syscall"
+)
+
+func BindToInterface(fd int, ifname string) error {
+	iface, err := net.InterfaceByName(ifname)
+	if err != nil {
+		return err
+	}
+	return syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_RECVIF, iface.Index)
+}
