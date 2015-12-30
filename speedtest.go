package speedtest

import (
	"log"
	"os"
)

func exists(pathtofile string) bool {
	_, err := os.Stat(pathtofile)
	return err == nil
}

type Speedtest struct {
	logger      *log.Logger
	logToFile   bool
	logfilepath string
	user        User
	serverList  List
}

func (s *Speedtest) ToggleLogDest() {
	s.logToFile = !s.logToFile
	if s.logToFile {
		var (
			f   *os.File
			err error
		)
		if exists(s.logfilepath) {
			f, err = os.OpenFile(s.logfilepath, os.O_RDWR, 0777)
		} else {
			f, err = os.Create(s.logfilepath)
		}
		if err != nil {
			panic(err)
		}
		s.logger.SetOutput(f)
	} else {
		s.logger.SetOutput(os.Stdout)
	}
}

func (s *Speedtest) SetLogFilePath(newFilePath string) {
	s.logfilepath = newFilePath
}

func (s *Speedtest) FetchServers() {
	s.user = FetchUserInfo()
	s.serverList = FetchServerList(s.user)
}

func (s *Speedtest) ShowUser() {
	s.user.Show()
}
func (s *Speedtest) ShowList() {
	s.serverList.Show()
}

func (s *Speedtest) ShowResult(serverIds []int) {
	targets := s.serverList.FindServer(serverIds)
	targets.StartTest(s.logger)
	targets.ShowResult(s.logger)
}

func New() *Speedtest {
	return &Speedtest{
		logger:      log.New(os.Stdout, "", log.Ltime),
		logToFile:   false,
		logfilepath: "./logger.log",
	}
}
