package judge

import (
)

func Problem1(ans string) bool {
	modelAns := `root:/bin/bash
daemon:/usr/sbin/nologin
bin:/usr/sbin/nologin
sys:/usr/sbin/nologin
sync:/bin/sync
games:/usr/sbin/nologin
man:/usr/sbin/nologin
lp:/usr/sbin/nologin
mail:/usr/sbin/nologin
news:/usr/sbin/nologin
uucp:/usr/sbin/nologin
proxy:/usr/sbin/nologin
www-data:/usr/sbin/nologin
backup:/usr/sbin/nologin
list:/usr/sbin/nologin
irc:/usr/sbin/nologin
gnats:/usr/sbin/nologin
nobody:/usr/sbin/nologin
_apt:/usr/sbin/nologin
nginx:/bin/false`
	if ans == modelAns {
		return true
	}
	return false
}
