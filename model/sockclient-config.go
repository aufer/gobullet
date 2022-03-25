package model

import "fmt"

type SockClientConfig struct {
	Prot string
	Host string
	Port string
}

func (scc *SockClientConfig) ConnUrl() string {
	return fmt.Sprintf("%s://%s:%s", scc.Prot, scc.Host, scc.Port)
}

func (scc *SockClientConfig) ConnSrv() string {
	return fmt.Sprintf("%s:%s", scc.Host, scc.Port)
}
