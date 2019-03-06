--- vendor/github.com/tsg/gopacket/pcap/pcap.go.orig	2019-03-01 02:07:58 UTC
+++ vendor/github.com/tsg/gopacket/pcap/pcap.go
@@ -170,7 +170,7 @@ type InterfaceAddress struct {
 // BPF is a compiled filter program, useful for offline packet matching.
 type BPF struct {
 	orig string
-	bpf  _Ctype_struct_bpf_program // takes a finalizer, not overriden by outsiders
+	bpf  C.struct_bpf_program // takes a finalizer, not overriden by outsiders
 }
 
 // BlockForever, when passed into OpenLive/SetTimeout, causes it to block forever
@@ -382,7 +382,7 @@ func (p *Handle) Error() error {
 
 // Stats returns statistics on the underlying pcap handle.
 func (p *Handle) Stats() (stat *Stats, err error) {
-	var cstats _Ctype_struct_pcap_stat
+	var cstats C.struct_pcap_stat
 	if -1 == C.pcap_stats(p.cptr, &cstats) {
 		return nil, p.Error()
 	}
@@ -443,7 +443,7 @@ func (p *Handle) SetBPFFilter(expr strin
 		}
 	}
 
-	var bpf _Ctype_struct_bpf_program
+	var bpf C.struct_bpf_program
 	cexpr := C.CString(expr)
 	defer C.free(unsafe.Pointer(cexpr))
 
@@ -486,7 +486,7 @@ func (b *BPF) String() string {
 }
 
 // BPF returns the compiled BPF program.
-func (b *BPF) BPF() _Ctype_struct_bpf_program {
+func (b *BPF) BPF() C.struct_bpf_program {
 	return b.bpf
 }
 
@@ -549,10 +549,10 @@ func FindAllDevs() (ifs []Interface, err
 	return
 }
 
-func findalladdresses(addresses *_Ctype_struct_pcap_addr) (retval []InterfaceAddress) {
+func findalladdresses(addresses *C.struct_pcap_addr) (retval []InterfaceAddress) {
 	// TODO - make it support more than IPv4 and IPv6?
 	retval = make([]InterfaceAddress, 0, 1)
-	for curaddr := addresses; curaddr != nil; curaddr = (*_Ctype_struct_pcap_addr)(curaddr.next) {
+	for curaddr := addresses; curaddr != nil; curaddr = (*C.struct_pcap_addr)(curaddr.next) {
 		var a InterfaceAddress
 		var err error
 		// In case of a tun device on Linux the link layer has no curaddr.addr.
@@ -818,7 +818,7 @@ func (h *Handle) NewDumper(file string) 
 // Writes a packet to the file. The return values of ReadPacketData
 // can be passed to this function as arguments.
 func (d *Dumper) WritePacketData(data []byte, ci gopacket.CaptureInfo) (err error) {
-	var pkthdr _Ctype_struct_pcap_pkthdr
+	var pkthdr C.struct_pcap_pkthdr
 	pkthdr.caplen = C.bpf_u_int32(ci.CaptureLength)
 	pkthdr.len = C.bpf_u_int32(ci.Length)
 
