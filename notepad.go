package notepad

import "os/exec"
// tag
// Open opens Notepad.exe on Windows.
func Open() {
    // Command to open Notepad.exe
    exec.Command("notepad.exe").Run()
}
