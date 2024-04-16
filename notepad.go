package notepad

import "os/exec"

// Open opens Notepad.exe on Windows.
func Open() {
    // Command to open Notepad.exe
    exec.Command("notepad.exe").Run()
}
