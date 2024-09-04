package proxyapi

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/util"
)

type rule hasSubPaths
type node interface{}

type hasSubPaths map[uriElem]node
type uriElem interface {
	uri() string
}

type version string

var defaultVersion = version("v1")

func (v version) uri() string {
	return string(v)
}

type path string

const here = path(".")

func (p path) uri() string {
	return string(p)
}

type hasFunctions []function
type function string

func (f function) uri() string {
	return string(f)
}

type routeConfig struct {
	svc  util.ServiceKey
	rule rule
}

type configBuilder routeConfig

func toService(key util.ServiceKey) configBuilder {
	return configBuilder{
		svc: key,
	}
}
func (b configBuilder) withRule(rule rule) routeConfig {
	return routeConfig{svc: b.svc, rule: rule}
}

type parseState struct {
	current node
	version version
	url     string
}

func parse(state parseState, result router) error {
	switch state.current.(type) {
	case hasFunctions:
		for _, fn := range state.current.(hasFunctions) {
			if result[fn] == nil {
				result[fn] = make(map[version]string)
			}

			if _, exists := result[fn][state.version]; exists {
				return fmt.Errorf("router for function %s with version %s already exists", fn, state.version)
			}
			result[fn][state.version] = join(state.url, fn.uri())
		}
		return nil
	case hasSubPaths:
		subPaths := state.current.(hasSubPaths)

		// iterate subpaths and parse recursively
		for subPath, next := range subPaths {

			// copy and construct subState
			subState := state
			if v, isVersion := subPath.(version); isVersion {
				subState.version = v
			}
			subState.url = join(state.url, subPath.uri())
			subState.current = next

			if err := parse(subState, result); err != nil {
				return err
			}
		}
		return nil
	}
	return errors.New("unknown type")
}

type router map[function]map[version]string

const protocol = "https"

func newRouter(svcMap util.ServiceMap, configs ...routeConfig) router {
	ret := make(router)
	for _, c := range configs {
		if _, exists := svcMap[c.svc]; !exists {
			panic(fmt.Errorf("svc: %s does not exists", c.svc))
		}
		mapRoute(protocol+"://"+svcMap[c.svc], c.rule, ret)
	}
	return ret
}

func mapRoute(host util.Host, target rule, to router) {
	initialState := parseState{current: hasSubPaths(target), version: defaultVersion, url: host.String()}
	if err := parse(initialState, to); err != nil {
		panic(err)
	}
}

func (r router) route(fn function, v version) (string, bool) {
	urls, ok := r[fn]
	if !ok {
		return "", false
	}

	url, ok := urls[v]
	if !ok {
		return "", false
	}
	return url, true
}

func join(a string, b string) string {
	if a == here.uri() {
		return b
	}

	if b == here.uri() {
		return a
	}

	return strings.TrimSuffix(a, "/") + "/" + strings.TrimPrefix(b, "/")
}
