package os_example

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(srcFileName, destFileName string) error {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		return err
	}

	dstFile, err := os.Open(destFileName)
	if err != nil {
		return err
	}

	n, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	fmt.Printf("Copied %d bytes from %s to %s", n, srcFileName, destFileName)

	return nil
}
