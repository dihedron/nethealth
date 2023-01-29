package base

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"

	"github.com/dihedron/go-log-facade/logging"
	"github.com/dihedron/go-log-facade/logging/console"
	"github.com/dihedron/go-log-facade/logging/file"
	"github.com/dihedron/go-log-facade/logging/uber"
)

type Command struct {
	// LogLevel sets the debugging level of the application.
	LogLevel string `short:"D" long:"log-level" description:"The debug level of the application." optional:"yes" choice:"off" choice:"trace" choice:"debug" choice:"info" choice:"warn" choice:"error" default:"warn" env:"GRAPHSTACK_LOG_LEVEL"`
	// LogStream is the type of logger to use.
	LogStream string `short:"L" long:"log-stream" description:"The logger to use." optional:"yes" choice:"zap" choice:"stdout" choice:"stderr" choice:"file" choice:"log" choice:"none" default:"stderr" env:"GRAPHSTACK_LOG_STREAM"`
	// CPUProfile sets the (optional) path of the file for CPU profiling info.
	CPUProfile string `short:"C" long:"cpu-profile" description:"The (optional) path where the CPU profiler will store its data." optional:"yes"`
	// MemProfile sets the (optional) path of the file for memory profiling info.
	MemProfile string `short:"M" long:"mem-profile" description:"The (optional) path where the memory profiler will store its data." optional:"yes"`
	// AutomationFriendly enables automation-friendly JSON output.
	AutomationFriendly bool `short:"A" long:"automation-friendly" description:"Whether to output in automation friendly JSON format." optional:"yes"`
}

// InitLogger initialises the logger.
func (cmd *Command) InitLogger(global bool) logging.Logger {
	switch cmd.LogLevel {
	case "trace":
		logging.SetGlobalLevel(logging.LevelTrace)
	case "debug":
		logging.SetGlobalLevel(logging.LevelDebug)
	case "info":
		logging.SetGlobalLevel(logging.LevelInfo)
	case "warn":
		logging.SetGlobalLevel(logging.LevelWarn)
	case "error":
		logging.SetGlobalLevel(logging.LevelError)
	case "off":
		logging.SetGlobalLevel(logging.LevelOff)
	}
	var logger logging.Logger = &logging.NoOpLogger{}
	switch cmd.LogStream {
	case "none":
		logger = &logging.NoOpLogger{}
	case "stdout":
		logger = console.NewLogger(console.StdOut)
	case "stderr":
		logger = console.NewLogger(console.StdErr)
	case "zap":
		logger, _ = uber.NewLogger()
	case "file":
		exe, _ := os.Executable()
		log := fmt.Sprintf("%s-%d.log", strings.Replace(exe, ".exe", "", -1), os.Getpid())
		logger = file.NewLogger(log)
	}
	if global {
		logging.SetLogger(logger)
	}
	return logger
}

func (cmd *Command) ProfileCPU() *Closer {
	logger := logging.GetLogger()
	var f *os.File
	if cmd.CPUProfile != "" {
		var err error
		f, err = os.Create(cmd.CPUProfile)
		if err != nil {
			logger.Errorf("could not create CPU profile at %s: %v", cmd.CPUProfile, err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			logger.Errorf("could not start CPU profiler: %v", err)
		}
	}
	return &Closer{
		file: f,
	}
}

func (cmd *Command) ProfileMemory() {
	logger := logging.GetLogger()
	if cmd.MemProfile != "" {
		f, err := os.Create(cmd.MemProfile)
		if err != nil {
			logger.Errorf("could not create memory profile at %s: %v", cmd.MemProfile, err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			logger.Errorf("could not write memory profile: %v", err)
		}
	}
}
