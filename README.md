# GoUtils
该库主要用于学习，实现一些常见的golang 函数，打造成一个工具库

## 1. crypt
加解密相关，包含hash、base64、aes、rsa、padding、random等

### 1.1 hash
hash函数，封装了go提供的所有hash函数和增加了一些函数，有如下函数

- Md5File：对文件进行md5操作，理论上也可以采用shaX等函数来处理
- HashBytes：使用指定的hash函数对传入的数据进行hash，返回原始的[]byte，支持的hash函数包括如下：
  ```
  const (
        HtMD5 HashType = iota + 1
        HtSha1
        HtSha224
        HtSha256
        HtSha384
        HtSha512
        HtFnv32
        HtFnvA32
        HtFnv64
        HtFnvA64
        HtFnv128
        HtFnvA128
        HtTime33
        HtAdler32
        HtCrc32
        HtCrc64ISO
        HtCrc64ECMA
  )
  ```
- HmacBytes：使用指定的hmacXXX函数对传入的数据进行hash，返回原始的[]byte
- ToHexString：[]byte转换string
- PBKDF2：PBKDF2哈希算法
- Time33：Time33哈希算法
- HashUInt32：使用指定的hash函数对传入的数据进行hash，返回uint32，目前hash算法函数只支持HtFnv32, HtFnvA32, HtAdler32, HtCrc32, HtTime33。
- HashUInt64：使用指定的hash函数对传入的数据进行hash，返回uint64，目前hash算法函数只支持HtFnv32, HtFnvA32, HtAdler32, HtCrc32, HtTime33,HtFnv64,HtFnvA64,HtCrc64ISO,HtCrc64ECMA
- JumpConsistentHash：jump consistent hash算法，返回uint32

### 1.2 random
一些随机函数，包含整型、整型范围、[]byte和字符串类型，有如下函数：

- RandInt：返回一个非负的随机整数，范围为[0, ∞)
- RandIntN：返回一个非负的随机整数，范围为[0, n-1]
- RandIntRange：返回一个非负的随机整数，范围为[min, max-1]
- RandBytes：返回一个指定长度的随机字节切片，每个字节的取值范围为[0x00,0xff]
- RandString：返回指定长度的随机字符串，字符串内容由大小写字母、数字和特殊字符组成

### 1.3 padding
实现了加解密算法中常见的填充算法和去填充算法，有如下函数：

- X923Padding：ANSI x9.23填充算法
- X923UnPadding：ANSI x9.23去填充算法
- ISO10126Padding：ISO 10126填充算法
- ISO10126UnPadding：ISO 10126去填充算法
- ISO7816Padding：ISO 7816-4填充函数
- ISO7816UnPadding：ISO 7816-4去填充函数
- PKCS7Padding：PKCS7填充算法，理论上使用PKCS5的也可以直接使用这个函数，推荐使用该算法
- PKCS7UnPadding：PKCS7去填充算法
- ZeroPadding：Zero填充算法，不推荐此填充算法，如果原始数据有末尾0x00，会导致在去填充的时候出问题
- ZeroUnPadding：Zero填充算法

### 1.4 aes 
实现了aes加解密算法的5种模式，理论上对于des算法也是可以适用的，支持的模式包括CBC、ECB、CFB、CTR、OFB。有如下函数：

- AESEncrypt：选择特定模式进行aes加密
- AESDecrypt：选择特定模式进行aes解密

### 1.5 rsa
实现了rsa加解密算法，包含公私钥生成、加解密、签名等操作，同时公私钥支持参数传入和文件读取2种方式，有如下函数：

- RSAGenKeyToFile：生成公钥和私钥，并保存到指定的文件中
- RSAGenKey：生成公钥和私钥
- RSAEncrypt：使用RSA公钥加密
- RSADecrypt：使用RSA私钥解密
- RSAEncryptFromFile：使用RSA公钥加密，公钥从文件读取
- RSADecryptFromFile：使用RSA私钥解密，私钥从文件读取
- RSASign：使用RSA私钥对信息签名  
- RSAVerySign：使用RSA公钥对信息进行签名验证
- RSASignFromFile：使用RSA私钥对信息签名，私钥从文件读取
- RSAVerySignFromFile：使用RSA公钥对信息进行签名验证，公钥从文件读取

### 1.6 base64
封装了下go src提供的base64算法，有如下函数：

- Base64Encode：base64编码
- Base64Decode：base64解码
- Base64UrlEncode：url-safe base64编码
- Base64UrlDecode：url-safe base64解码

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
- GetNowString：获取当前时间，返回字符串格式(2021-06-24 20:28:00)
- GetNowRFC3339：获取RFC3339格式的当前时间，返回字符串格式(2021-06-24T20:28:21+08:00)
- GetNowTime：获取当前时间，返回time.Time格式(2021-05-06 15:02:24.4718541 +0800 CST)
- GetNowYMD：获取当前时间的年月日，返回字符串格式(2021-05-06)
- GetNowHMS：获取当前时间的时分秒，返回字符串格式(20:28:40)
- IsLeapYear：判断是否是闰年
- CompareTime：比较2个时间戳是否在一定范围，目前只支持判断是否在同一年、同一个月、同一周、同一天、同一小时、同一分钟,判断时compareRange的取值如下：
  ```
  IsOneYear = iota
  IsOneMonth
  IsOneWeek
  IsOneDay
  IsOneHour
  IsOneMinute
  ```
- CalcIntervalDays：计算两个时间戳之间间隔天数，向下取整。begin和end表示时间戳，单位秒
- CalcIntervalHours：计算两个时间戳之间的间隔小时数，向下取整。begin和end表示时间戳，单位秒
- DayBeginTime：获取指定时间当天开始时间，例如2021-06-24的开始时间为2021-06-24 00:00:00 +0800 CST
- DayEndTime：获取指定时间当天结束时间，例如2021-06-24的开始时间为2021-06-24 23:59:59.999999999 +0800 CST
- DaySecs: 判断指定时间是当天的第几秒,注意，当天开始时间为当天的第0秒，当天结束时间为当天的第86399秒
- Before：当前时间减去n秒的时间，返回字符串格式
- After：当前时间加上n秒后的时间，返回字符串格式
- GetWeekDate：获取当周某天,返回t时刻日期，返回格式为time.Time
- GetWeekDateStr：获取当周某天的年月日字符串格式， 如20180402
- ToTime：2006-01-02 15:04:05字符串格式转换Time格式
- ToUnix：2006-01-02 15:04:05字符串格式转换时间戳
- ToRFC3339：2006-01-02 15:04:05字符串格式转换RFC3339格式
- RFC3339ToUnix：2006-01-02T15:04:05Z07:00字符串格式转换时间戳
- ToString：格式化time.Time为字符串,支持的格式化选项如下：
  ```
  MM - month - 01
  M - month - 1, single bit
  DD - day - 02
  D - day 2
  YYYY - year - 2006
  YY - year - 06
  HH - 24 hours - 03
  H - 24 hours - 3
  hh - 12 hours - 03
  h - 12 hours - 3
  mm - minute - 04
  m - minute - 4
  ss - second - 05
  s - second = 5
  ```

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

## 5. compress
实现了gzip、zlib、frate三种压缩、解压缩相关的函数

- GzipCompress：gzip压缩
- GzipDecompress：gzip解压缩 
- ZlibCompress：zlib压缩
- ZlibDecompress：zlib解压缩
- FlateCompress：flate压缩
- FlateDecompress：flate解压缩

## 6. convert
类型转换相关的函数

- ToString：interface{}转换成字符串
- ToBool：interface{}转换成bool
- ToFloat32：interface{}转换成float32,对于string、[]byte、float64等类型的数据而言，如果转换的数据超出math.MaxFloat32,会报错
- ToFloat64：interface{}转换成float64，对于string、[]byte等类型的数据而言，如果转换的数据超出math.MaxFloat64,会报错；对于float32数据而言，使用float64可能存在精度丢失，目前使用strconv.ParseFloat(fmt.Sprintf("%f", f), 64)，暂时还没有发现精度丢失的情况
- ToInt**：interface{}转换成int64、int32、int16、int8、int等类型，如果转换的数据超出math.MaxInt**，转换的结果有问题，不会报错；
- ToUint**：interface{}转换成uint64、uint32、uint16、uint8、uint等类型，对于负数，会报错；对于int**、uint**、float32、float64等类型的数据而言，如果转换的数据超出math.MaxUint**，转换的结果有问题，不会报错；而string、[]byte等类型的数据会报错