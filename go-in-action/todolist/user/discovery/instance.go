package discovery

import (
	"encoding/json"
	"errors"
	"strings"

	"google.golang.org/grpc/resolver"
)

type Server struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Version string `json:"version"`
	Weight  int64  `json:"weight"`
}

func BuildPrefix(server Server) string {
	if server.Version == "" {
		return "/" + server.Name + "/"
	}
	return "/" + server.Name + "/" + server.Version + "/"
}

func BuildRegisterPath(server Server) string {
	return BuildPrefix(server) + server.Address
}

func ParseValue(val []byte) (Server, error) {
	var server Server
	err := json.Unmarshal(val, &server)
	return server, err
}

func SplitPath(path string) (Server, error) {
	var server Server
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return server, errors.New("invalid path")
	}

	server.Address = parts[len(parts)-1]
	return server, nil
}

func Exist(ls []resolver.Address, addr resolver.Address) bool {
	for i := range ls {
		if ls[i].Addr == addr.Addr {
			return true
		}
	}
	return false
}

func Remove(addrs []resolver.Address, addr resolver.Address) ([]resolver.Address, bool) {
	for i := range addrs {
		if addrs[i].Addr == addr.Addr {
			addrs[i] = addrs[len(addrs)-1]
			return addrs[:len(addrs)-1], true
		}
	}
	return nil, false
}

func BuildResolverUrl(app string) string {
	return schema + ":///" + app
}
