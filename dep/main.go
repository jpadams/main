package main

type Dep struct{}

func (m *Dep) TcpTest() string {
	return "TcpTest"
}
func (m *Dep) UdpTest() string {
	return "UdpTest"
}
func (m *Dep) IpTest() string {
	return "IpTest"
}
func (m *Dep) HttpTest() string {
	return "HttpTest"
}
func (m *Dep) HttpsTest() string {
	return "HttpsTest"
}
func (m *Dep) FooTest() string {
	return "HttpsTest"
}
