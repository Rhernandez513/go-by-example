package main

import "fmt"
import "io/ioutil"
import "os/exec"

// NOTE: This program is intended for UNIX based machines, will not work on Windows
// Let's look at a few examples of spawning processes in Go
func main () {

	// We’ll start with a simple command that takes no arguments
	// Or input and just prints something to stdout.
	// The exec.Command helper creates an object to represent this
	// external process.
	dateCmd := exec.Command("date")

	// .Output() is another helper that handles the common case of 
	// running a command, waiting for it to finish, and collecting 
	// its output. If there were no errors, dateOut will hold bytes 
	// with the date info.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Next we’ll look at a slightly more involved case where we pipe
	// data to the external process on its stdin and collect the results
	// from its stdout.

	// Here we explicitly grab input/output pipes, start the process, 
	// write some input to it, read the resulting output, and finally
	// wait for the process to exit.
	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbyegrep"))	
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Note that when spawning commands we need to provide an explicitly delineated command
	// and argument array, vs. being able to just pass in one command-line string. 
	// If you want to spawn a full command with a string, you can use bash’s -c option:

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

// EOF
