package main

import "fmt"

// Logger 接口定义了所有记录器必须实现的方法
type Logger interface {
	Log(message string)
}

// ConsoleLogger 是 Logger 接口的一个实现，表示控制台记录器
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("Console logger: " + message)
}

// FileLogger 是 Logger 接口的一个实现，表示文件记录器
type FileLogger struct{}

func (f *FileLogger) Log(message string) {
	fmt.Println("File logger: " + message)
	// 这里应该包含写入文件的逻辑
}

// LoggerCreator 定义了工厂方法的接口
type LoggerCreator interface {
	CreateLogger() Logger
}

// ConsoleLoggerCreator 是 LoggerCreator 的实现，用于创建 ConsoleLogger
type ConsoleLoggerCreator struct{}

func (c *ConsoleLoggerCreator) CreateLogger() Logger {
	return &ConsoleLogger{}
}

// FileLoggerCreator 是 LoggerCreator 的实现，用于创建 FileLogger
type FileLoggerCreator struct{}

func (f *FileLoggerCreator) CreateLogger() Logger {
	return &FileLogger{}
}

func main() {
	var loggerCreator LoggerCreator

	// 选择使用控制台记录器
	loggerCreator = &ConsoleLoggerCreator{}
	logger := loggerCreator.CreateLogger()
	logger.Log("This is a message.")

	// 选择使用文件记录器
	loggerCreator = &FileLoggerCreator{}
	logger = loggerCreator.CreateLogger()
	logger.Log("This is another message.")
}
