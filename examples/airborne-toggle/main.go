package main

import (
	"io"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/scax/logitech-simple-airborne-toggle"
)

func init() {
	log.Logger = log.Output(&zerolog.ConsoleWriter{Out: colorable.NewColorableStdout()}).With().Logger()
}

type Config struct {
	InfoUpdate     time.Duration `toml:"UpdateInfo" toml-default:"5s"`
	LogitechUpdate time.Duration `toml:"UpdateLogitech" toml-default:"0.050s"`
	LogFilePath    string        `toml:"Log"`
	logFile        io.WriteCloser
	ProgramExe     string `toml:"Programm"`
}

var config Config

func (cfg *Config) OpenLog() error {
	if cfg.LogFilePath == "" {
		return nil
	}
	file, err := os.Create(cfg.LogFilePath)
	if err != nil {
		return err
	}
	cfg.logFile = file
	return nil
}

func (cfg *Config) CloseLog() error {
	if cfg.logFile == nil {
		return nil
	}
	return cfg.logFile.Close()
}

func createLogger(writer io.Writer) {
	if writer == nil {
		return
	}
	log.Logger = log.Output(zerolog.MultiLevelWriter(&zerolog.ConsoleWriter{Out: colorable.NewColorableStdout()}, &zerolog.ConsoleWriter{Out: writer, NoColor: true})).With().Logger()
}

func initLogitech(exe string) {

}

func main() {
	err := cleanenv.ReadConfig("config.toml", &config)
	if err != nil {
		log.Panic().Err(err).Msg("read config")
	}
	err = config.OpenLog()
	if err != nil {
		log.Panic().Err(err).Interface("config", &config).Msg("open log file")
	}
	defer config.CloseLog()

	createLogger(config.logFile)

	updateTicker := time.NewTicker(50 * time.Millisecond)
	infoTicker := time.NewTicker(config.InfoUpdate)
	stopChan := make(chan bool)

	initLogitech(config.ProgramExe)

loop:
	for {

		select {
		case <-updateTicker.C:
			update(&config)
		case <-infoTicker.C:
			info(&config)
		case stop := <-stopChan:
			if stop {
				break loop
			}
		}

	}

	defer logitech.LogiSteeringShutdown()

}

func Test() {

	bo, err := logitech.LogiUpdate()
	log.Info().Bool("bool", bo).Err(err).Msg("LogiUpdate")

	bo, err = logitech.LogiSteeringInitialize(true)
	log.Info().Bool("bool", bo).Err(err).Msg("LogiSteeringInitialize")

	bo, err = logitech.LogiUpdate()
	log.Info().Bool("bool", bo).Err(err).Msg("LogiUpdate")

	bo, err = logitech.LogiSteeringInitialize(true)
	log.Info().Bool("bool", bo).Err(err).Msg("LogiSteeringInitialize")

	bo, err = logitech.LogiUpdate()
	log.Info().Bool("bool", bo).Err(err).Msg("LogiUpdate")

	dij, err := logitech.LogiGetState(0)
	log.Info().Interface("dij", dij).Int("index", 0).Err(err).Msg("LogiGetState")

	dij, err = logitech.LogiGetState(1)
	log.Info().Interface("dij", dij).Int("index", 1).Err(err).Msg("LogiGetState")

	dij, err = logitech.LogiGetState(2)
	log.Info().Interface("dij", dij).Int("index", 2).Err(err).Msg("LogiGetState")

	dijEng, err := logitech.LogiGetStateENGINES(0)
	log.Info().Interface("dijEng", dijEng).Int("index", 0).Err(err).Msg("LogiGetStateENGINES")

	dijEng, err = logitech.LogiGetStateENGINES(1)
	log.Info().Interface("dijEng", dijEng).Int("index", 1).Err(err).Msg("LogiGetStateENGINES")

	dijEng, err = logitech.LogiGetStateENGINES(2)
	log.Info().Interface("dijEng", dijEng).Int("index", 2).Err(err).Msg("LogiGetStateENGINES")

	name, err := logitech.LogiGetFriendlyProductName(0, 256)
	log.Info().Str("name", name).Int("index", 0).Err(err).Msg("LogiGetStateENGINES")

	logitech.LogiGetFriendlyProductName(1, 256)
	log.Info().Str("name", name).Int("index", 1).Err(err).Msg("LogiGetStateENGINES")

	logitech.LogiGetFriendlyProductName(2, 256)
	log.Info().Str("name", name).Int("index", 2).Err(err).Msg("LogiGetStateENGINES")

}
