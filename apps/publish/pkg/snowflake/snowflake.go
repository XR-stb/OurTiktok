package snowflake

import (
	"strconv"
	"sync"
	"time"
)

const (
	epoch          = int64(1400000000000)              // 设置起始时间
	timestampBits  = uint(41)                          // 时间戳占用位数
	workeridBits   = uint(10)                          // 机器id所占位数
	sequenceBits   = uint(12)                          // 序列所占的位数
	timestampMax   = int64(-1 ^ (-1 << timestampBits)) // 时间戳最大值
	workeridMax    = int64(-1 ^ (-1 << workeridBits))  // 支持的最大机器id数量
	sequenceMask   = int64(-1 ^ (-1 << sequenceBits))  // 支持的最大序列id数量
	workeridShift  = sequenceBits                      // 机器id左移位数
	timestampShift = sequenceBits + workeridBits       // 时间戳左移位数
)

type Snowflake struct {
	sync.Mutex          //锁
	lasttimestamp int64 //时间戳
	workerid      int64 //机器id
	sequence      int64 //序列
}

func (s *Snowflake) New() string {
	s.Lock()
	now := time.Now().UnixNano() / 1000000
	if s.lasttimestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			for now <= s.lasttimestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		s.Unlock()
		return ""
	}
	s.lasttimestamp = now
	s.Unlock()
	shift := int64((t)<<timestampShift | (s.workerid << workeridShift) | (s.sequence))
	res := strconv.FormatInt(shift, 10)
	return res
}
