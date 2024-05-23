package dial

import (
	"context"
	"net"
)

var ErrNotPermitted = &net.AddrError{Err: "No permitted address found"}

type IPFilteringWrapper struct {
	delegate *net.Dialer
}

func NewIPFilteringWrapper(delegate *net.Dialer) *IPFilteringWrapper {
	return &IPFilteringWrapper{delegate: delegate}
}

// DialContext 클러스터 내부로의 dial 을 막기 위해 사설 아이피와 기타 대역을 차단합니다.
// 연결이 차단되었을 경우 ErrNotPermitted 를 반환합니다.
func (d *IPFilteringWrapper) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	conn, err := d.delegate.DialContext(ctx, network, addr)
	if err != nil {
		return nil, err
	}

	host, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
	ip := net.ParseIP(host)

	if !isPermittedIP(ip) {
		_ = conn.Close()
		return nil, ErrNotPermitted
	}

	return conn, nil
}

func isPermittedIP(ip net.IP) bool {
	if ip == nil {
		return false
	}
	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false
	}
	return !ip.IsPrivate()
}
