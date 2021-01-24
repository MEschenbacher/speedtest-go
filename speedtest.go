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
	serverList  ServerList
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
	user, err := FetchUserInfo()
	if err != nil {
		s.logger.Print(err)
		return
	}
	s.user = *user
	list, err := FetchServerList(user)
	if err != nil {
		s.logger.Print(err)
		return
	}
	s.serverList = list
}

func (s *Speedtest) ShowUser() {
	s.logger.Print(s.user.String())
}
func (s *Speedtest) ShowList() {
	s.logger.Print(s.serverList.String())
}

func (s *Speedtest) ShowResult(serverIds []int) {
	targets, err := s.serverList.FindServer(serverIds)
	if err != nil {
		s.logger.Print(err)
		return
	}
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
