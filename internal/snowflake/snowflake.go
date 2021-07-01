package snowflake

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/sony/sonyflake"
)

var Generator *sonyflake.Sonyflake

func init() {
	// 从命令行中解析出机器ID
	if len(os.Args) < 3 {
		log.Fatal("请指定机器ID和服务端口")
	}
	id, err := strconv.ParseUint(os.Args[1], 10, 16)
	if err != nil {
		log.Fatal("机器ID格式不正确")
	}

	st := sonyflake.Settings{
		StartTime: time.Date(2020, 11, 01, 0, 0, 0, 0, time.Local),
		MachineID: func() (uint16, error) {
			return uint16(id), nil
		},
		CheckMachineID: nil,
	}

	Generator = sonyflake.NewSonyflake(st)
}
