package inject

import (
	"context"

	"go.viam.com/utils"

	"go.viam.com/core/board"
	pb "go.viam.com/core/proto/api/v1"
)

// Board is an injected board.
type Board struct {
	board.Board
	MotorByNameFunc            func(name string) (board.Motor, bool)
	ServoByNameFunc            func(name string) (board.Servo, bool)
	SPIByNameFunc              func(name string) (board.SPI, bool)
	AnalogReaderByNameFunc     func(name string) (board.AnalogReader, bool)
	DigitalInterruptByNameFunc func(name string) (board.DigitalInterrupt, bool)
	MotorNamesFunc             func() []string
	ServoNamesFunc             func() []string
	SPINamesFunc               func() []string
	AnalogReaderNamesFunc      func() []string
	DigitalInterruptNamesFunc  func() []string
	CloseFunc                  func() error
	ConfigFunc                 func(ctx context.Context) (board.Config, error)
	StatusFunc                 func(ctx context.Context) (*pb.BoardStatus, error)
	GPIOSetFunc                func(ctx context.Context, pin string, high bool) error
	GPIOGetFunc                func(ctx context.Context, pin string) (bool, error)
	PWMSetFunc                 func(ctx context.Context, pin string, dutyCycle byte) error
	PWMSetFreqFunc             func(ctx context.Context, pin string, freq uint) error
}

// MotorByName calls the injected MotorByName or the real version.
func (b *Board) MotorByName(name string) (board.Motor, bool) {
	if b.MotorByNameFunc == nil {
		return b.Board.MotorByName(name)
	}
	return b.MotorByNameFunc(name)
}

// ServoByName calls the injected ServoByName or the real version.
func (b *Board) ServoByName(name string) (board.Servo, bool) {
	if b.ServoByNameFunc == nil {
		return b.Board.ServoByName(name)
	}
	return b.ServoByNameFunc(name)
}

// SPIByName calls the injected SPIByName or the real version.
func (b *Board) SPIByName(name string) (board.SPI, bool) {
	if b.SPIByNameFunc == nil {
		return b.Board.SPIByName(name)
	}
	return b.SPIByNameFunc(name)
}

// AnalogReaderByName calls the injected AnalogReaderByName or the real version.
func (b *Board) AnalogReaderByName(name string) (board.AnalogReader, bool) {
	if b.AnalogReaderByNameFunc == nil {
		return b.Board.AnalogReaderByName(name)
	}
	return b.AnalogReaderByNameFunc(name)
}

// DigitalInterruptByName calls the injected DigitalInterruptByName or the real version.
func (b *Board) DigitalInterruptByName(name string) (board.DigitalInterrupt, bool) {
	if b.DigitalInterruptByNameFunc == nil {
		return b.Board.DigitalInterruptByName(name)
	}
	return b.DigitalInterruptByNameFunc(name)
}

// MotorNames calls the injected MotorNames or the real version.
func (b *Board) MotorNames() []string {
	if b.MotorNamesFunc == nil {
		return b.Board.MotorNames()
	}
	return b.MotorNamesFunc()
}

// ServoNames calls the injected ServoNames or the real version.
func (b *Board) ServoNames() []string {
	if b.ServoNamesFunc == nil {
		return b.Board.ServoNames()
	}
	return b.ServoNamesFunc()
}

// SPINames calls the injected SPINames or the real version.
func (b *Board) SPINames() []string {
	if b.SPINamesFunc == nil {
		return b.Board.SPINames()
	}
	return b.SPINamesFunc()
}

// AnalogReaderNames calls the injected AnalogReaderNames or the real version.
func (b *Board) AnalogReaderNames() []string {
	if b.AnalogReaderNamesFunc == nil {
		return b.Board.AnalogReaderNames()
	}
	return b.AnalogReaderNamesFunc()
}

// DigitalInterruptNames calls the injected DigitalInterruptNames or the real version.
func (b *Board) DigitalInterruptNames() []string {
	if b.DigitalInterruptNamesFunc == nil {
		return b.Board.DigitalInterruptNames()
	}
	return b.DigitalInterruptNamesFunc()
}

// Close calls the injected Close or the real version.
func (b *Board) Close() error {
	if b.CloseFunc == nil {
		return utils.TryClose(b.Board)
	}
	return b.CloseFunc()
}

// Status calls the injected Status or the real version.
func (b *Board) Status(ctx context.Context) (*pb.BoardStatus, error) {
	if b.StatusFunc == nil {
		return b.Board.Status(ctx)
	}
	return b.StatusFunc(ctx)
}

// GPIOSet calls the injected GPIOSet or the real version.
func (b *Board) GPIOSet(ctx context.Context, pin string, high bool) error {
	if b.GPIOSetFunc == nil {
		return b.Board.GPIOSet(ctx, pin, high)
	}
	return b.GPIOSetFunc(ctx, pin, high)
}

// GPIOGet calls the injected GPIOGet or the real version.
func (b *Board) GPIOGet(ctx context.Context, pin string) (bool, error) {
	if b.GPIOGetFunc == nil {
		return b.Board.GPIOGet(ctx, pin)
	}
	return b.GPIOGetFunc(ctx, pin)
}

// PWMSet calls the injected PWMSet or the real version.
func (b *Board) PWMSet(ctx context.Context, pin string, dutyCycle byte) error {
	if b.PWMSetFunc == nil {
		return b.Board.PWMSet(ctx, pin, dutyCycle)
	}
	return b.PWMSetFunc(ctx, pin, dutyCycle)
}

// PWMSetFreq calls the injected PWMSetFreq or the real version.
func (b *Board) PWMSetFreq(ctx context.Context, pin string, freq uint) error {
	if b.PWMSetFreqFunc == nil {
		return b.Board.PWMSetFreq(ctx, pin, freq)
	}
	return b.PWMSetFreqFunc(ctx, pin, freq)
}
