package cmd

import (
	"context"
	"coscli/util"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tencentyun/cos-go-sdk-v5"
	"os"
)

var mbCmd = &cobra.Command{
	Use:   "mb",
	Short: "Create bucket",
	Long:  `Create bucket

Format:
  ./coscli mb cos://<bucket-name>-<app-id> -r region

Example:
  ./coscli mb cos://examplebucket-1234567890 -r ap-shanghai`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		bucketIDName, cosPath := util.ParsePath(args[0])
		if bucketIDName == "" || cosPath != "" {
			return fmt.Errorf("Invalid arguments! ")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		createBucket(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(mbCmd)

	mbCmd.Flags().StringP("region", "r", "", "Region")

	_ = mbCmd.MarkFlagRequired("region")
}

func createBucket(cmd *cobra.Command, args []string) {
	flagRegion, _ := cmd.Flags().GetString("region")
	bucketIDName, _ := util.ParsePath(args[0])
	c := util.CreateClient(&config, bucketIDName, flagRegion)

	opt := &cos.BucketPutOptions{
		XCosACL:                   "",
		XCosGrantRead:             "",
		XCosGrantWrite:            "",
		XCosGrantFullControl:      "",
		XCosGrantReadACP:          "",
		XCosGrantWriteACP:         "",
		CreateBucketConfiguration: nil,
	}

	_, err := c.Bucket.Put(context.Background(), opt)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Create a new bucket! name: %s region: %s\n", bucketIDName, flagRegion)
}
