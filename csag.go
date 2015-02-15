package main

import "fmt"
import "bufio"
import "os"

//import "time"
import "github.com/fatih/color"

func main() {
	color.Set(color.FgBlue, color.Bold)
	defer color.Unset() // Use it in your function
	fmt.Println("                                                            M")
	fmt.Println("                                                           MMM")
	fmt.Println("                                                           MMMM")
	fmt.Println("                                                        MMMMMMMM  ")
	fmt.Println("                                                      M  MMM MMM MM  ")
	fmt.Println("                                                   MM   MMM   MMM   M   ")
	fmt.Println("                                                  M    MMM     MMM    MM")
	fmt.Println("                                                  M    MMM     MMM     M")
	fmt.Println("                                                  M   MMM       MMM    M")
	fmt.Println("                                                  M  MMM         MMM   M")
	fmt.Println("    MMMMMMMMMMMMMMMMMM     NMMMMMMMMMMMMMMMMM     M  MM           MM   M       MMMMMMMMMMMMMMMMMM")
	fmt.Println("   MMMMMMMMMMMMMMMMMMM    MMMMMMMMMMMMMMMMMM      M MMM           MMM  M     MMMMMMMMMMMMMMMMMMMM")
	fmt.Println(" MMMM                   MMMM                      MMMM             MMM M    MMMM")
	fmt.Println(" MMMM                   MMMM                      MMMM             MMM M    MMMM")
	fmt.Println(" MMMM                   MMMM                      MMM               MMMM    MMMM")
	fmt.Println(" MMMM                   MMMM                      MM                MMMM    MMMM")
	fmt.Println(" MMMM                    MMMMMMMMMMMMMMMMMM      MMM                 MMM    MMMM         MMMMMM")
	fmt.Println(" MMMM                      MMMMMMMMMMMMMMMMMM    MM           MMM    MMM    MMMM         MMMMMMMM")
	fmt.Println(" MMMM                                     MMM   MMM         MMMMMMM   MM    MMMM              MMM")
	fmt.Println(" MMMM                                     MMM   MMM      MMMMM   MMM  MMM   MMMM              MMM")
	fmt.Println(" MMMM                                     MMM  MMMM    MMMM        MMM MM   MMMM              MMM")
	fmt.Println(" MMMM                                     MMM  MM M  MMMM            MMMMM  MMMM              MMM")
	fmt.Println("   MMMMMMMMMMMMMMMMMMM  MMMMMMMMMMMMMMMMMMMM  MMM MMMMM               MMMM   MMMMMMMMMMMMMMMMMMMM")
	fmt.Println("    MMMMMMMMMMMMMMMMMM  MMMMMMMMMMMMMMMMMM    MM MMMM                  MMM     MMMMMMMMMMMMMMM  M")
	fmt.Println("                                             MMMMMM                    M MM")
	fmt.Println("                                             MMMM M                    M  M")
	fmt.Println("                                             MM   M                  M")
	fmt.Println("                                                  M               MM")
	fmt.Println("                                                  M             MM")
	fmt.Println("                                                   MM         M")
	fmt.Println("                                                      MM    M")
	fmt.Println("                                                        MM")
	fmt.Println("")
	fmt.Println("All text will now be bold magenta.")
	color.Unset()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	android := []string{
		"                             MM                            MM:                                         ",
		"                              7M                          MM,                                          ",
		"                               +M:                       MM                                            ",
		"                                ~M=   7MMMMMMMMMMMM?,   MM                                             ",
		"                                 :MMMMMMMMMMMMMMMMMMMMMMM                                              ",
		"                               MMMMMMMMMMMMMMMMMMMMMMMMMMMM:                                           ",
		"                             7?MMMMMMMMMMMMMMMMMMMMMMMMMMMMMM+                                         ",
		"                           MMMMMMMM?MMMMMMMMMMMMMMMMMMMMMMMMMMM+                                       ",
		"                         ~MMMMMM,    MMMMMMMMMMMMMMMM?    MMMMMMM                                      ",
		"                        MMMMMMMM     ?MMMMMMMMMMMMMMM      MMMMMMM:                                    ",
		"                       MMMMMMMMM+    MMMMMMMMMMMMMMMM?    ?MMMMMMMM,                                   ",
		"                      =MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM                                   ",
		"                      MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM                                  ",
		"                     MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM                                  ",
		"                     MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM?                                 ",
		"                     MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM?                                 ",
		"                                                                                                       ",
		"         MMMMMM,     MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM     MMMMMMM                     ",
		"       MMMMMMMMM?    MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM    MMMMMMMMM                    ",
		"       MMMMMMMMMM    MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM   MMMMMMMMMMM                   ",
		"       MMMMMMMMMM    MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM   MMMMMMMMMMM                   ",
		"       MMMMMMMMMM    MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM   MMMMMMMMMMM                   ",
		"       MMMMMMMMMM    MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM   MMMMMMMMMMM                   ",
		"       MMMMMMMMMM    MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM   MMMMMMMMMMM                   ",
		"       MMMMMMMMMM    MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM   MMMMMMMMMMM                   ",
		"       MMMMMMMMMM    MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM   MMMMMMMMMMM                   ",
		"       MMMMMMMMMM    MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM   MMMMMMMMMMM                   ",
		"       MMMMMMMMMM    MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM   MMMMMMMMMMM                   ",
		"       MMMMMMMMMM    MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM   MMMMMMMMMMM                   ",
		"       MMMMMMMMMM    MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM   MMMMMMMMMMM                   "}

	ziko := []string{
		"                       ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄    ▄  ▄▄▄▄▄▄▄▄▄▄▄ ",
		"                      ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌  ▐░▌▐░░░░░░░░░░░▌",
		"                       ▀▀▀▀▀▀▀▀▀█░▌ ▀▀▀▀█░█▀▀▀▀ ▐░▌ ▐░▌ ▐░█▀▀▀▀▀▀▀█░▌",
		"                                ▐░▌     ▐░▌     ▐░▌▐░▌  ▐░▌       ▐░▌",
		"                       ▄▄▄▄▄▄▄▄▄█░▌     ▐░▌     ▐░▌░▌   ▐░▌       ▐░▌",
		"                      ▐░░░░░░░░░░░▌     ▐░▌     ▐░░▌    ▐░▌       ▐░▌",
		"                      ▐░█▀▀▀▀▀▀▀▀▀      ▐░▌     ▐░▌░▌   ▐░▌       ▐░▌",
		"                      ▐░▌               ▐░▌     ▐░▌▐░▌  ▐░▌       ▐░▌",
		"                      ▐░█▄▄▄▄▄▄▄▄▄  ▄▄▄▄█░█▄▄▄▄ ▐░▌ ▐░▌ ▐░█▄▄▄▄▄▄▄█░▌",
		"                      ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌  ▐░▌▐░░░░░░░░░░░▌",
		"                       ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀    ▀  ▀▀▀▀▀▀▀▀▀▀▀ "}
	name := []string{
		"                      ╔═╗┬ ┬┌─┐┌─┐┬┌─┌─┐┬─┐┌┐┌  ╦ ╦┌─┐┬─┐┌─┐┌┬┐┌─┐┌┬┐",
		"                      ╚═╗│ │├─┘├─┤├┴┐│ │├┬┘│││  ║║║├─┤├┬┘│ │ │││ ││││",
		"                      ╚═╝└─┘┴  ┴ ┴┴ ┴└─┘┴└─┘└┘  ╚╩╝┴ ┴┴└─└─┘─┴┘└─┘┴ ┴"}
	ce := []string{
		" __   __         __       ___  ___  __      ___       __          ___  ___  __          __  ",
		"/  ` /  \\  |\\/| |__) |  |  |  |__  |__)    |__  |\\ | / _` | |\\ | |__  |__  |__) | |\\ | / _` ",
		"\\__, \\__/  |  | |    \\__/  |  |___ |  \\    |___ | \\| \\__> | | \\| |___ |___ |  \\ | | \\| \\__> "}

	color.Set(color.FgGreen, color.Bold)
	defer color.Unset()
	for i := 0; i < len(android); i++ {
		fmt.Println(android[i])
	}
	for i := 0; i < len(ziko); i++ {
		fmt.Println(ziko[i])
	}
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}
	for i := 0; i < len(ce); i++ {
		fmt.Println(ce[i])
	}
}
