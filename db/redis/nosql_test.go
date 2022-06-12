package redis

import (
	"context"
	"testing"
)

func TestNosql(t *testing.T) {
	M.redisDb.Set(context.Background(), "nosql_test", "hhhhh", 0)
}
