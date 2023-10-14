package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"

	"github.com/adshao/go-binance/v2"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

const outputDir = "out"

func main() {
	logger, _ := zap.NewProduction()
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Panicf("failed to sync logs: %+v", err)
		}
	}()

	apiKey, ok := os.LookupEnv("BINANCE_API_KEY")
	if !ok {
		logger.Panic("BINANCE_API_KEY environment variable is required")
	}

	secretKey, ok := os.LookupEnv("BINANCE_SECRET_KEY")
	if !ok {
		logger.Panic("BINANCE_SECRET_KEY environment variable is required")
	}

	fileName, ok := os.LookupEnv("FILE_NAME")
	if !ok {
		fileName = "historical_data.csv"
	}

	startTimeRaw, ok := os.LookupEnv("START_TIME")
	if !ok {
		startTimeRaw = time.Now().Add((24 * time.Hour) * 30).Format(time.RFC3339)
	}

	startTime, err := time.Parse(time.RFC3339, startTimeRaw)
	if err != nil {
		logger.Panic("failed to parse START_TIME", zap.Error(err))
	}

	client := binance.NewClient(apiKey, secretKey)

	ctx := context.Background()

	if err := client.NewPingService().Do(ctx); err != nil {
		logger.Panic("failed to ping binance API", zap.Error(err))
	}

	// ensure out directory is created
	_, err = os.Stat(outputDir)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			logger.Panic("failed to check for the out directory", zap.Error(err))
		}

		if err = os.Mkdir(outputDir, os.ModePerm); err != nil {
			logger.Panic("failed to create out dir", zap.Error(err))
		}
	}

	filePath := fmt.Sprintf("%s/%s", outputDir, fileName)

	// ensure file is created
	var file *os.File
	_, err = os.Stat(filePath)
	if err != nil {
		fmt.Printf("error type: %T\n", err)
		if !errors.Is(err, fs.ErrNotExist) {
			logger.Panic("failed to check for file", zap.String("file", filePath), zap.Error(err))
		}

		if file, err = os.Create(filePath); err != nil {
			logger.Panic("failed to create file", zap.Error(err))
		}
	}

	if file == nil {
		file, err = os.OpenFile(filePath, os.O_RDWR, os.ModeType)
		if err != nil {
			logger.Panic("failed to open file", zap.String("file", fileName), zap.Error(err))
		}
	}

	logger.Info("starting...")

	// cleans file before writes
	if err = file.Truncate(0); err != nil {
		logger.Panic("failed to clean file", zap.Error(err))
	}

	// sets the next write to line 0 column 0 (start of file)
	if _, err = file.Seek(0, 0); err != nil {
		logger.Panic("failed to seek for the start of the file", zap.Error(err))
	}

	// writes CSV header
	if _, err = file.WriteString("Date,Close,Open,High,Low,Volume\n"); err != nil {
		logger.Panic("failed to write headers", zap.Error(err))
	}

	isProcessing := true
	endTime := time.Now()

	for isProcessing {
		logger.Info("processing...", zap.String("from", endTime.Format(time.RFC3339)), zap.String("to", endTime.Add(-1000*time.Hour).Format(time.RFC3339)))
		klines := client.NewKlinesService()
		klines.Symbol("BTCUSDT")
		klines.Interval("1h")
		klines.EndTime(endTime.UnixMilli())
		klines.Limit(1000)
		response, err := klines.Do(ctx)
		if err != nil {
			logger.Panic("failed to get klines", zap.Error(err))
		}

		for index := range response {
			current := response[(len(response)-1)-index]
			if _, err = file.WriteString(fmt.Sprintf("%d,%s,%s,%s,%s,%s\n", current.CloseTime, current.Close, current.Open, current.High, current.Low, current.Volume)); err != nil {
				logger.Panic("failed to write line", zap.Error(err))
			}
		}

		endTime = time.UnixMilli(response[0].CloseTime)
		if response[len(response)-1].CloseTime < startTime.UnixMilli() {
			isProcessing = false
		}
	}

	logger.Info("file saved", zap.String("path", filePath))
}
