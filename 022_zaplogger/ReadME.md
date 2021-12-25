# 1: 说明

- `Zap` 日志包
- `Lumberjack`是一个 Go 包，用于将日志写入滚动文件。

zap 不支持文件归档，如果要支持文件按大小或者时间归档，需要使用 lumberjack，lumberjack 也是 zap 官方推荐的。
