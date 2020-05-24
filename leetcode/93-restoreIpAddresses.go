package leetcode

import "strconv"

// "25525511135" ==> ["255.255.11.135", "255.255.111.35"]
func restoreIpAddresses(s string) (result []string) {
	if len(s) < 4 {
		return
	}

	backtrackingRestoreIpAddresses(&result, s, "", 0, 0)

	return
}

func backtrackingRestoreIpAddresses(result *[]string, ip, restored string, idx, count int) {
	if count > 4 {
		return
	}
	if count == 4 && idx == len(ip) {
		*result = append(*result, restored)
		return
	}

	for i := 1; i < 4; i++ {
		if idx+i > len(ip) {
			break
		}

		cur := ip[idx : idx+i]
		if len(cur) > 1 && cur[0] == '0' {
			continue
		}

		v, _ := strconv.Atoi(cur)
		if i == 3 && v >= 256 {
			continue
		}

		part := restored + "." + cur
		if count == 0 {
			part = cur
		}

		backtrackingRestoreIpAddresses(result, ip, part, idx+i, count+1)
	}
}
