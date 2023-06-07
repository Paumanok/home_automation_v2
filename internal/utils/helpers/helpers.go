package helpers

import (
	"strings"
	"path"
	"time"
)

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
// credit to this blog post https://blog.merovius.de/posts/2017-06-18-how-not-to-use-an-http-router/
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p) //normalizes the given path, adds a leading slash
	i := strings.Index(p[1:], "/") + 1 //finds first occurance of "/", increments by 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

type SyncTimer struct {
	TimerInterval int
	TimerVal int
}

func (st *SyncTimer) GetNextDelay() int {
	return st.TimerVal
}

func (st *SyncTimer) Timer() {
	for {
		for st.TimerVal != 0 {
			st.TimerVal--
			time.Sleep(1 * time.Second)
		}
		st.TimerVal = st.TimerInterval
	}

}
