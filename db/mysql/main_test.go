package dbutils

import (
	"testing"
)

func Test_QueryOne(t *testing.T) {
	// NewConn(dsn)
	QueryOne()

}
func Test_QueryMore(t *testing.T) {
	// NewConn(dsn)

	QueryMore(0)
}
func Test_Insert(t *testing.T) {
	// NewConn(dsn)
	dbExec()
}
