package snowFlake

import (
	"strconv"
	"time"
)
var (
	machineID     int64 // 机器 id 占2位, 十进制范围是 [ 0, 3 ]
	sn            int64 // 序列号占 8 位,十进制范围是 [ 0, 255 ]
	lastTimeStamp string // 上次的时间戳(毫秒级), 1秒=1000毫秒, 1毫秒=1000微秒,1微秒=1000纳秒
)
//获取当前时间，时间格式为：202020215，年月日时 10位
func curTime () string{
	year := strconv.Itoa(time.Now().Year())
	month := strconv.Itoa(int(time.Now().Month()))
	if len(month) == 1 {
		month = "0" + month
	}
	var day string
	var hour string
	if time.Now().Day() < 10 {
		day = "0" + strconv.Itoa(time.Now().Day())
	} else {
		day = strconv.Itoa(time.Now().Day())
	}
	if time.Now().Hour() < 10 {
		hour = "0" + strconv.Itoa(time.Now().Hour())
	} else {
		hour = strconv.Itoa(time.Now().Hour())
	}
	return year + month + day + hour
}
func init() {
	lastTimeStamp = curTime()
}

func SetMachineId(mid int64) {
	// 把机器 id 左移 8 位,让出 8 位空间给序列号使用
	machineID = mid << 8
}

func GetSnowflakeId() int64 {
	//curTimeStamp := time.Now().UnixNano() / 1000000
	curTimeStamp := curTime()
	// 同一毫秒
	if curTimeStamp == lastTimeStamp {
		sn++
		// 序列号占 12 位,十进制范围是 [ 0, 4095 ]
		if sn > 256 {
			time.Sleep(time.Millisecond)
			curTimeStamp = curTime()
			lastTimeStamp = curTimeStamp
			sn = 0
		}

		// 取 64 位的二进制数 0000000000 0000000000 0000000000 000000000 ..... 00000011 1111111111 ( 这里共 10 个 1 )和时间戳进行并操作
		// 并结果( 右数 )第 54 位必然是 0,  低 10 位也就是时间戳的低 10 位
		temp, _ := strconv.ParseInt(curTimeStamp, 10, 64)
		rightBinValue := temp & 0x3FF
		// 机器 id 占用10位空间,序列号占用10位空间,所以左移 10 位; 经过上面的并操作,左移后的第 1 位,必然是 0
		rightBinValue <<= 10
		machineID <<= 8
		id := rightBinValue | machineID | sn
		return id
	}


	if curTimeStamp > lastTimeStamp {
		sn = 0
		lastTimeStamp = curTimeStamp
		// 取 64 位的二进制数 0000000000 0000000000 0000000000 000000000 ..... 00000011 1111111111 ( 这里共 10 个 1 )和时间戳进行并操作
		// 并结果( 右数 )第 54 位必然是 0,  低 10 位也就是时间戳的低 10 位
		temp, _ := strconv.ParseInt(curTimeStamp, 10, 64)
		rightBinValue := temp & 0x3FF
		// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
		rightBinValue <<= 10
		machineID <<= 8
		id := rightBinValue | machineID | sn
		return id
	}


	if curTimeStamp < lastTimeStamp {
		return 0
	}
	return 0
}