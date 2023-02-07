package publish

import (
	"fmt"
	"testing"
)

func TestSnowflakeCreate(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(s.create())
	}
}
