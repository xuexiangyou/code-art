package common

import (
	"compress/gzip"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
)

// GetClientIP gets the correct IP for the end client instead of the proxy
func GetClientIP(c *gin.Context) string {
	// first check the X-Forwarded-For header
	requester := c.Request.Header.Get("X-Forwarded-For")
	// if empty, check the Real-IP header
	if len(requester) == 0 {
		requester = c.Request.Header.Get("X-Real-IP")
	}
	// if the requester is still empty, use the hard-coded address from the socket
	if len(requester) == 0 {
		requester = c.Request.RemoteAddr
	}

	// if requester is a comma delimited list, take the first one
	// (this happens when proxied via elastic load balancer then again through nginx)
	if strings.Contains(requester, ",") {
		requester = strings.Split(requester, ",")[0]
	}

	return requester
}

// GzipDoCompress gzip 文件解压
/**
 *src压缩包路径
 * dst解压后的路径
*/
func GzipDoCompress(src, dst string) error {
	gzipFile, err := os.Open(src)
	if err != nil {
		return err
	}
	gzipReader, err := gzip.NewReader(gzipFile)
	if err == io.EOF {
		return nil
	}
	defer gzipReader.Close()

	outfileWriter, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer outfileWriter.Close()

	_, err = io.Copy(outfileWriter, gzipReader)
	return err
}
