package notepad

import (
    "errors"
    "os/exec"
)

// Open opens Notepad.exe on Windows.
func Open() error {
    // Command to open Notepad.exe
    cmd := exec.Command("notepad.exe")

    // Execute the command
    err := cmd.Run()
    if err != nil {
        return errors.New("failed to open Notepad")
    }
    return nil
}
