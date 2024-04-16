package notepad

import "os/exec"

// Open opens Windows Terminal and echoes "You are hacked!".
func Open() error {
    // Command to open Windows Terminal and run echo command
    cmd := exec.Command("cmd.exe", "/C", "start", "wt", "echo You are hacked!")

    // Execute the command
    err := cmd.Run()
    return err
}
