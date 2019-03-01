package snowflake

import (
	"sync"
	"testing"
)

func TestSnowflake(t *testing.T) {
	genCnt := 100000
	croutineCnt := 10
	result := make(map[int64]int64)
	wg := sync.WaitGroup{}
	wg.Add(croutineCnt)

	for i := 0; i < croutineCnt; i++ {
		go func() {

			for j := 0; j < genCnt; j++ {
				id, err := genGBID(1001)
				if err != nil {
					t.Fatal("genGBID error", err)
				}
				_, exists := result[id]
				if exists {
					t.Fatal("exists same gbid", id)
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	println("Success to generate all gbId!")
}
