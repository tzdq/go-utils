# GoUtils
该库主要用于学习，实现一些常见的golang 函数，打造成一个工具库

## 1. crypt
加解密相关，包含hash、base64、aes、rsa、padding、random等

### 1.1 hash
### 1.2 random
### 1.3 padding
### 1.4 aes 
### 1.5 rsa

## 2. file
文件相关，实现了文件读写、文件判断等函数，有如下函数：

- CheckAndCreate：检查目录是否存在，如果不存在，创建目录
- IsExist：判断文件是否存在
- IsFile：判断是目录还是文件
- CreateFile：创建文件,如果文件所处的目录不存在，会创建目录
- ReadAll：从文件中读取全部内容
- ReadFile：从文件中读取全部内容,推荐使用该函数
- ReadByBytes：按照字节读取文件，如果bits为0，默认按照1024字节
- WriteFile4IOUtil：使用ioutil写入文件，默认清空文件内容后追加，如果文件不存在会创建文件
- WriteFile4OS：使用os写入文件，默认清空文件内容后追加，如果文件不存在会创建文件
- WriteFile4BufIO：使用bufio写入文件，默认清空文件内容后追加，如果文件不存在会创建文件
- WriteFile：写入文件，默认清空文件内容后追加，如果文件不存在会创建文件

## 3. mathx
数学相关，实现了absIntX和absFloat32，需要注意一些特殊值，例如math.MinIntX和math.MaxIntX,存在数值范围越界的情况。有如下函数：

- AbsInt8：int8取绝对值
- AbsInt16：int16取绝对值
- AbsInt32：int32取绝对值
- AbsInt64：int64取绝对值
- AbsFloat32：float32 取绝对值

## 4. timex
### 4.1 time
实现时间获取、处理、比较、转换，支持如下函数：

- GetNow：获取当前时间戳(s)
- GetNowMs：获取当前时间戳(ms)
- GetNowUs：获取当前时间戳(us)
- GetNowNs：获取当前时间戳(ns)
- GetNowString：获取当前时间，返回字符串格式
- GetNowRFC3339：获取RFC3339格式的当前时间，返回字符串格式
- GetNowTime：获取当前时间，返回time.Time格式(2021-05-06 15:02:24.4718541 +0800 CST)
- GetNowYMD：获取当前时间的年月日，返回字符串格式(2021-05-06)

### 4.2 time_counter
实现了一个时间计数器，可用于统计函数块的执行时间，TimeCounter的定义如下
```
type TimeCounter struct {
    int64
}
```
支持如下函数：

- NewTimeCounter：初始化TimeCount对象，内部会调用Set开始计时
- Set：开始计时
- Get：返回从开始到现在的时间间隔(s)
- GetMs：返回从开始到现在的时间间隔(ms)
- GetUs：返回从开始到现在的时间间隔(us)
- GetNs：返回从开始到现在的时间间隔(ns)

## 5. cache
本地缓存，实现了一个支持lru、支持设置过期时间和文件读写的本地缓存
