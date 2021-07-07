package api

import "io"

// CloseQuietly 关闭io, 忽略错误
func CloseQuietly(closer io.Closer) {
	_ = closer.Close()
}
