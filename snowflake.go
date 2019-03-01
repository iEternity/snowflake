package snowflake

import (
	"errors"
	"sync"
	"time"
)

const (
	TIME_BITS     = 0x1ffffffffff // 41个1，最高位是符号位
	MACHINE_BITS  = 0x3ff         // 10个1
	SEQUENCE_BITS = 0xfff         // 12个1

	// 时间基准为2018.1.1.0.0.0
	EPOCH = 1514736000000

	TIME_SHIFT    = 24
	MACHINE_SHIFT = 12

	MAX_SEQUENCE = 4095
)

var (
	sequence      int64 = 0
	lastTimestamp int64 = 0
	mutex         sync.Mutex
)

func genGBID(machineID int16) (int64, error) {
	if machineID > MACHINE_BITS {
		return 0, errors.New("machineID exceed max value:1023")
	}

	mutex.Lock()
	defer mutex.Unlock()

	millsec := time.Now().UnixNano() / 1000000

	if millsec == lastTimestamp {
		sequence++

		if sequence > MAX_SEQUENCE {
			for millsec <= lastTimestamp {
				millsec = time.Now().UnixNano() / 1000000
			}
			sequence = 0
			lastTimestamp = millsec
		}
	} else {
		sequence = 0
		lastTimestamp = millsec
	}

	millsec -= EPOCH
	millsec &= TIME_BITS
	return (millsec << TIME_SHIFT) | (int64(machineID) << MACHINE_SHIFT) | sequence, nil
}
