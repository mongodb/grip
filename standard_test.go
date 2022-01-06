package grip

import (
	"fmt"
	"testing"

	"github.com/mongodb/grip/level"
	"github.com/mongodb/grip/logging"
	"github.com/mongodb/grip/message"
	"github.com/mongodb/grip/send"
	"github.com/stretchr/testify/suite"
)

const testMessage = "hello world"

type (
	basicMethod func(interface{})
	lnMethod    func(...interface{})
	fMethod     func(string, ...interface{})
	whenMethod  func(bool, interface{})
)

type LoggingMethodSuite struct {
	logger        Journaler
	loggingSender *send.InternalSender
	stdSender     *send.InternalSender
	suite.Suite
}

func TestLoggingMethodSuite(t *testing.T) {
	suite.Run(t, new(LoggingMethodSuite))
}

func (s *LoggingMethodSuite) SetupTest() {
	s.logger = logging.NewGrip("test")

	s.stdSender = send.MakeInternalLogger()
	s.NoError(SetSender(s.stdSender))
	s.Exactly(GetSender(), s.stdSender)

	s.loggingSender = send.MakeInternalLogger()
	s.NoError(s.logger.SetSender(s.loggingSender))
	s.Exactly(s.logger.GetSender(), s.loggingSender)
}

func (s *LoggingMethodSuite) TestWhenMethods() {
	cases := map[string][]whenMethod{
		"emergency": {EmergencyWhen, s.logger.EmergencyWhen},
		"alert":     {AlertWhen, s.logger.AlertWhen},
		"critical":  {CriticalWhen, s.logger.CriticalWhen},
		"error":     {ErrorWhen, s.logger.ErrorWhen},
		"warning":   {WarningWhen, s.logger.WarningWhen},
		"notice":    {NoticeWhen, s.logger.NoticeWhen},
		"info":      {InfoWhen, s.logger.InfoWhen},
		"debug":     {DebugWhen, s.logger.DebugWhen},
	}

	for kind, loggers := range cases {
		s.Len(loggers, 2)
		loggers[0](true, testMessage)
		loggers[1](true, testMessage)

		s.True(s.loggingSender.HasMessage())
		s.True(s.stdSender.HasMessage())
		lgrMsg := s.loggingSender.GetMessage()
		s.True(lgrMsg.Logged)
		stdMsg := s.stdSender.GetMessage()
		s.True(stdMsg.Logged)

		s.Equal(lgrMsg.Rendered, stdMsg.Rendered,
			fmt.Sprintf("%s: \n\tlogger: %+v \n\tstandard: %+v", kind, lgrMsg, stdMsg))

		loggers[0](false, testMessage)
		loggers[1](false, testMessage)

		lgrMsg = s.loggingSender.GetMessage()
		s.False(lgrMsg.Logged)
		stdMsg = s.stdSender.GetMessage()
		s.False(stdMsg.Logged)

	}
}

func (s *LoggingMethodSuite) TestBasicMethod() {
	cases := map[string][]basicMethod{
		"emergency": {Emergency, s.logger.Emergency},
		"alert":     {Alert, s.logger.Alert},
		"critical":  {Critical, s.logger.Critical},
		"error":     {Error, s.logger.Error},
		"warning":   {Warning, s.logger.Warning},
		"notice":    {Notice, s.logger.Notice},
		"info":      {Info, s.logger.Info},
		"debug":     {Debug, s.logger.Debug},
	}

	inputs := []interface{}{true, false, []string{"a", "b"}, message.Fields{"a": 1}, 1, "foo"}

	for kind, loggers := range cases {
		s.Len(loggers, 2)
		s.False(s.loggingSender.HasMessage())
		s.False(s.stdSender.HasMessage())

		for _, msg := range inputs {
			loggers[0](msg)
			loggers[1](msg)

			s.True(s.loggingSender.HasMessage())
			s.True(s.stdSender.HasMessage())
			lgrMsg := s.loggingSender.GetMessage()
			stdMsg := s.stdSender.GetMessage()
			s.Equal(lgrMsg.Rendered, stdMsg.Rendered,
				fmt.Sprintf("%s: \n\tlogger: %+v \n\tstandard: %+v", kind, lgrMsg, stdMsg))
		}
	}

}

func (s *LoggingMethodSuite) TestlnMethods() {
	cases := map[string][]lnMethod{
		"emergency": {Emergencyln, s.logger.Emergencyln},
		"alert":     {Alertln, s.logger.Alertln},
		"critical":  {Criticalln, s.logger.Criticalln},
		"error":     {Errorln, s.logger.Errorln},
		"warning":   {Warningln, s.logger.Warningln},
		"notice":    {Noticeln, s.logger.Noticeln},
		"info":      {Infoln, s.logger.Infoln},
		"debug":     {Debugln, s.logger.Debugln},
	}

	for kind, loggers := range cases {
		s.Len(loggers, 2)
		s.False(s.loggingSender.HasMessage())
		s.False(s.stdSender.HasMessage())

		loggers[0](true, testMessage, testMessage)
		loggers[1](true, testMessage, testMessage)

		s.True(s.loggingSender.HasMessage())
		s.True(s.stdSender.HasMessage())
		lgrMsg := s.loggingSender.GetMessage()
		stdMsg := s.stdSender.GetMessage()
		s.Equal(lgrMsg.Rendered, stdMsg.Rendered,
			fmt.Sprintf("%s: \n\tlogger: %+v \n\tstandard: %+v", kind, lgrMsg, stdMsg))
	}
}

func (s *LoggingMethodSuite) TestfMethods() {
	cases := map[string][]fMethod{
		"emergency": {Emergencyf, s.logger.Emergencyf},
		"alert":     {Alertf, s.logger.Alertf},
		"critical":  {Criticalf, s.logger.Criticalf},
		"error":     {Errorf, s.logger.Errorf},
		"warning":   {Warningf, s.logger.Warningf},
		"notice":    {Noticef, s.logger.Noticef},
		"info":      {Infof, s.logger.Infof},
		"debug":     {Debugf, s.logger.Debugf},
	}

	for kind, loggers := range cases {
		s.Len(loggers, 2)
		s.False(s.loggingSender.HasMessage())
		s.False(s.stdSender.HasMessage())

		loggers[0]("%s: %d", testMessage, 3)
		loggers[1]("%s: %d", testMessage, 3)

		s.True(s.loggingSender.HasMessage())
		s.True(s.stdSender.HasMessage())
		lgrMsg := s.loggingSender.GetMessage()
		stdMsg := s.stdSender.GetMessage()
		s.Equal(lgrMsg.Rendered, stdMsg.Rendered,
			fmt.Sprintf("%s: \n\tlogger: %+v \n\tstandard: %+v", kind, lgrMsg, stdMsg))
	}
}

func (s *LoggingMethodSuite) TestProgramaticLevelMethods() {
	type (
		lgwhen   func(bool, level.Priority, interface{})
		lgwhenln func(bool, level.Priority, ...interface{})
		lgwhenf  func(bool, level.Priority, string, ...interface{})
		lg       func(level.Priority, interface{})
		lgln     func(level.Priority, ...interface{})
		lgf      func(level.Priority, string, ...interface{})
	)

	cases := map[string]interface{}{
		"when": []lgwhen{LogWhen, s.logger.LogWhen},
		"lg":   []lg{Log, s.logger.Log},
		"lgln": []lgln{Logln, s.logger.Logln},
		"lgf":  []lgf{Logf, s.logger.Logf},
	}

	const l = level.Emergency

	for kind, loggers := range cases {
		s.Len(loggers, 2)
		s.False(s.loggingSender.HasMessage())
		s.False(s.stdSender.HasMessage())

		switch log := loggers.(type) {
		case []lgwhen:
			log[0](true, l, testMessage)
			log[1](true, l, testMessage)
		case []lgwhenln:
			log[0](true, l, testMessage, "->", testMessage)
			log[1](true, l, testMessage, "->", testMessage)
		case []lgwhenf:
			log[0](true, l, "%T: (%s) %s", log, kind, testMessage)
			log[1](true, l, "%T: (%s) %s", log, kind, testMessage)
		case []lg:
			log[0](l, testMessage)
			log[1](l, testMessage)
		case []lgln:
			log[0](l, testMessage, "->", testMessage)
			log[1](l, testMessage, "->", testMessage)
		case []lgf:
			log[0](l, "%T: (%s) %s", log, kind, testMessage)
			log[1](l, "%T: (%s) %s", log, kind, testMessage)
		default:
			panic("testing error")
		}

		s.True(s.loggingSender.HasMessage())
		s.True(s.stdSender.HasMessage())
		lgrMsg := s.loggingSender.GetMessage()
		stdMsg := s.stdSender.GetMessage()
		s.Equal(lgrMsg.Rendered, stdMsg.Rendered,
			fmt.Sprintf("%s: \n\tlogger: %+v \n\tstandard: %+v", kind, lgrMsg, stdMsg))
	}
}
