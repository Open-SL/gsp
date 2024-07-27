package cli

import (
	"fmt"
	"io"
	"os/exec"
)

func RunCommand(cmd *exec.Cmd) error {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}

	// Stream stdout and stderr in separate goroutines
	go streamOutput(stdout, "stdout")
	go streamOutput(stderr, "stderr")

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to wait for command: %w", err)
	}

	return nil
}

func streamOutput(reader io.Reader, prefix string) {
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			fmt.Printf("[%s] %s", prefix, buf[:n])
		}
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error reading %s: %v\n", prefix, err)
			}
			break
		}
	}
}
