package main

import (
	"log"

	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/cmd"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const serverName = "government_exam"

func main() {
	rootCmd := &cobra.Command{
		Use:   serverName,
		Short: serverName,
		Long:  "GovernmentExam project server",
	}
	rootCmd.AddCommand(cmd.NewCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("cmd execute error", zap.Error(err))
	}
}
