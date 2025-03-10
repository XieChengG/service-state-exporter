package collector

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// 服务端口map
var servicePorts = map[string]int{
	"mysql":    13306,
	"redis":    6379,
	"rabbitmq": 15672,
	"gitlab":   8081,
	"harbor":   8084,
	"jenkins":  8085,
	"nacos":    8848,
	"mongodb":  27017,
}

// 服务状态map
var serviceState = map[string]int{}

// 执行netstat命令并过滤出指定端口的行
func runCommandWithFilter(baseCmd string, args []string, filter func(string) bool) ([]string, error) {
	cmd := exec.Command(baseCmd, args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("command failed: %v", err)
	}

	var results []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if filter(line) {
			results = append(results, line)
		}
	}
	return results, nil
}

// 获取服务状态
func getServiceState() (map[string]int, error) {
	for sn, sp := range servicePorts {
		lines, err := runCommandWithFilter("netstat", []string{"-lntup"}, func(line string) bool {
			return strings.Contains(line, ":"+strconv.Itoa(sp))
		})

		if err != nil {
			return nil, err
		}

		if len(lines) > 0 {
			serviceState[sn] = 1
		} else {
			serviceState[sn] = 0
		}
	}
	return serviceState, nil
}
