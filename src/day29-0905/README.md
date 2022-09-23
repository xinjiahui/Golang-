## 需求分析
- 支持往不同的地方输出日志
- 日志分级别
    - Debug
    - Info
    - Trace
    - Warning
    - Error
    - Fatal
- 日志要支持开关
- 完整的日志记录要包含有时间，行号，文件名，日志级别，日志信息
- 日志文件要切割
    - 按文件大小切割
    - 按日期切割
        - 1.记录上一次切割的小时数
        - 2.在写日志之前检查一下当前时间的小时数和之前保存的是否一致。
    