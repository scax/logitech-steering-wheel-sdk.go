package logitech

import (
	"unsafe"

	"github.com/rs/zerolog/log"
	"golang.org/x/sys/windows"
)

var (
	logitech = windows.NewLazyDLL("LogitechSteeringWheelEnginesWrapper.dll")
)

var (
	pLogiPlayCarAirborne              = logitech.NewProc("LogiPlayCarAirborne")
	pLogiStopCarAirborne              = logitech.NewProc("LogiStopCarAirborne")
	pLogiUpdate                       = logitech.NewProc("LogiUpdate")
	pLogiGetState                     = logitech.NewProc("LogiGetState")
	pLogiGetStateENGINES              = logitech.NewProc("LogiGetStateENGINES")
	pLogiSteeringInitializeWithWindow = logitech.NewProc("LogiSteeringInitializeWithWindow")
	pLogiSteeringInitialize           = logitech.NewProc("LogiSteeringInitialize")
	pLogiGetFriendlyProductName       = logitech.NewProc("LogiGetFriendlyProductName")
	pLogiSteeringShutdown             = logitech.NewProc("LogiSteeringShutdown")
	pLogiIsPlaying                    = logitech.NewProc("LogiIsPlaying")
)

func LogiIsPlaying(index, forceType int) (bool, error) {

	r0, r1, err := pLogiIsPlaying.Call(uintptr(index), uintptr(forceType))

	log.Trace().Uint64("r0", uint64(r0)).Uint64("r1", uint64(r1)).Err(err).Msg("LogiIsPlaying")

	if err != windows.Errno(0) {
		return false, err
	}

	return r0 != 0, nil

}

func LogiSteeringShutdown() (bool, error) {

	r0, r1, err := pLogiSteeringShutdown.Call()

	log.Trace().Uint64("r0", uint64(r0)).Uint64("r1", uint64(r1)).Err(err).Msg("LogiSteeringShutdown")

	if err != windows.Errno(0) {
		return false, err
	}

	return r0 != 0, nil
}

func LogiGetFriendlyProductName(index int, buffSize int) (string, error) {
	buf := make([]uint16, buffSize)

	r0, r1, err := pLogiGetFriendlyProductName.Call(uintptr(index), uintptr(unsafe.Pointer(&buf[0])), uintptr(buffSize))

	log.Trace().Uint64("r0", uint64(r0)).Uint64("r1", uint64(r1)).Err(err).Int("index", index).Int("buffSize", buffSize).Msg("LogiGetFriendlyProductName")

	if err != windows.Errno(0) {
		return "", err
	}

	return windows.UTF16ToString(buf), nil
}

func LogiPlayCarAirborne(index int) (bool, error) {

	r0, r1, err := pLogiPlayCarAirborne.Call(uintptr(index))

	log.Trace().Uint64("r0", uint64(r0)).Uint64("r1", uint64(r1)).Err(err).Int("index", index).Msg("LogiPlayCarAirborne")

	if err != windows.Errno(0) {
		return false, err
	}

	return r0 != 0, nil

}
func LogiStopCarAirborne(index int) (bool, error) {

	r0, r1, err := pLogiStopCarAirborne.Call(uintptr(index))

	log.Trace().Uint64("r0", uint64(r0)).Uint64("r1", uint64(r1)).Err(err).Int("index", index).Msg("LogiStopCarAirborne")

	if err != windows.Errno(0) {
		return false, err
	}

	return r0 != 0, nil

}

func LogiGetStateENGINES(index int) (*DIJOYSTATE2ENGINES, error) {

	r0, r1, err := pLogiGetStateENGINES.Call(uintptr(index))

	if err != windows.Errno(0) {
		return nil, err
	}

	log.Trace().Uint64("r0", uint64(r0)).Uint64("r1", uint64(r1)).Err(err).Msg("LogiGetState")

	return ((*DIJOYSTATE2ENGINES)(unsafe.Pointer(r0))), nil

}
func LogiGetState(index int) (*DIJOYSTATE2, error) {

	r0, r1, err := pLogiGetState.Call(uintptr(index))

	if err != windows.Errno(0) {
		return nil, err
	}

	log.Trace().Uint64("r0", uint64(r0)).Uint64("r1", uint64(r1)).Err(err).Msg("LogiGetState")

	return ((*DIJOYSTATE2)(unsafe.Pointer(r0))), nil

}
func LogiUpdate() (bool, error) {

	r0, _, err := pLogiUpdate.Call()

	if err != windows.Errno(0) {
		return false, err
	}

	//	log.Trace().Uint64("r0", uint64(r0)).Uint64("r1", uint64(r1)).Err(err).Msg("LogiUpdate")

	return r0 != 0, nil

}

func LogiSteeringInitializeWithWindow(ignoreXInput bool, windowsHwd uintptr) (bool, error) {
	b := 0
	if ignoreXInput {
		b = 1
	}
	r0, _, err := pLogiSteeringInitializeWithWindow.Call(uintptr(b), uintptr(windowsHwd))
	//	log.Trace().Uint64("r0", uint64(r0)).Uint64("r1", uint64(r1)).Err(err).Msg("LogiSteeringInitializeWithWindow")
	if err != windows.Errno(0) {
		return false, err
	}
	return r0 != 0, nil
}
func LogiSteeringInitialize(ignoreXInput bool) (bool, error) {
	b := 0
	if ignoreXInput {
		b = 1
	}
	r0, _, err := pLogiSteeringInitialize.Call(uintptr(b))
	//	log.Trace().Uint64("r0", uint64(r0)).Uint64("r1", uint64(r1)).Err(err).Msg("LogiSteeringInitialize")
	if err != windows.Errno(0) {
		return false, err
	}
	return r0 != 0, nil
}
