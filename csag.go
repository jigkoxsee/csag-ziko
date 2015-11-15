package main

import "fmt"
import "bufio"
import "os"
import "time"
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
	color.Unset()

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	android := []string{
		"                             MM                            MM:                                         ",
		"                              MM                          MM,                                          ",
		"                               +M:                       MM                                            ",
		"                                ~M=   MMMMMMMMMMMMM?,   MM                                             ",
		"                                 :MMMMMMMMMMMMMMMMMMMMMMM                                              ",
		"                               MMMMMMMMMMMMMMMMMMMMMMMMMMMM:                                           ",
		"                             M?MMMMMMMMMMMMMMMMMMMMMMMMMMMMMM+                                         ",
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

	docker := []string{
		"                                        .IMMMMMMMM.                                                 ",
		"                                        .IMMMMMMMM.                                                 ",
		"                                        .IMMMMMMIM.                                                 ",
		"                                        .IMMMMMMIM.                                                 ",
		"                                        .MMIMMMMMM.                                                 ",
		"                                        .IMMMMMMMM.                                                 ",
		"                    ..,..................,.........                                                 ",
		"                    .MMIMMMMI?.MMMMMMMMI.IMIMMMMMI.                           II                    ",
		"                    .MMMMMMMM?.MMMMMMMMM.IMMMMMMMI.                          IMMM~                  ",
		"                    .MMMMMMMM?.MMMMMMMMI.IMMMMMMIM.                         ,MMMMMI                 ",
		"                    .MMMMMMMM?.MMMMMMMMI.IMMMMMMIM.                         IIMMMMMI                ",
		"                    .MMMMMMMM?.MMMMMMMMI.IMMMMMMIM.                         MIMMMIMM+               ",
		"          ...........???????I+.?????????.??I??????...........               MMMMMMMMM              .",
		"          :MMMMMMMI=.MMIMMMMM?.MMMMMMMMI.IMIMMMMMI.?MMMMMMMM.               MMMMMMMMM=..,~===:      ",
		"          :MMMMMMMI~.MMMMMMMM?.MMMMMMMMI.IMMMMMMMM.IMMMMMMMM.               +MMMMMMMMIMMMMMMMMMMI+  ",
		"          :MMMMMMMM~.MMMMMMMM?.MMMMMMMMI.IMMMMMMMM.IMMMMMMMM.                MMMMMMMMMMMMMMMMMMMMM  ",
		"          :MMMMMMMI~.MMMMMMMM?.MMMMMMMMI.IMMMMMMMM.IMMMMMMMM.                :MIMMMMMIMMMMMMMMMMI   ",
		"          :MMMMMIMM~.MMMMMMMI?.IMMMMMMMI.IMMMMMMII.?MMMMMMMM.                 .MMMMMMMMMIMMMMMM:    ",
		"          :MMMMMMMM~.IMMMMMMM?.MMMMMMMMM.IMIMMMMIM.IMMMMMMMM.               :MMMMMMMMMMMMIMMI,      ",
		".......................................................................,=IMIMMMMIMI?III?=,          ",
		".MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMIIMMMMMMMMM.                ",
		".IMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMIMMMMMMI.                 ",
		".MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMI,.                 ",
		".MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM=.                  ",
		".MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMI?.                   ",
		".?MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMI.                    ",
		".:MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMI.                     ",
		"..IIMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMI.                      ",
		" .~MMMMMMMMMMMMMMMMMMMMMMMMMM   MMMIMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMII+.                       ",
		"  .?MMMMMMMMMMMMMMMMMMMMMMMM MI  MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMIMMMI.                         ",
		"   .MMMMMMMMMMMMMMMMMMMMMMMM MM  MMIMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMIMM=.                          ",
		"    .MIMMMMMMMMMMMMMMMMMMMMMMMIMIMIMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMI..                           ",
		"     .MMMMIIMMMMMMMIMMMI~,IIMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMIMM,.                             ",
		"      .IMI                :MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMIM~..                              ",
		"       .~MM=.              :MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM~                                  ",
		"         .IMM~              .MMIMMMMMMMMMMMMMMMMMMMMMMMMIIMMMII                                     ",
		"           .IMM?.  .         .?MMMMMMMMMMMMMMMMMMMMMMMMMMMMM:                                       ",
		"              .+IMM+            ,MMIMMMMMMMMMMMMMMMMMMM?,                                           ",
		"                 .=MMIII+,        .+MMMIMMMMMMMMMMM~                                                ",
		"                      ,=IMMMMMMIMMMIMMMMMMMI                                                        "}
	ubuntu := []string{
		"                                                                ?OOOOO                             ",
		"                                                              OOOOOOOOOO+                          ",
		"                                                             OOOOOOOOOOOOO                         ",
		"                                              ,7$$$$7I      OOOOOOOOOOOOOO=                        ",
		"                                        $$$$$$$$$$$$$$$$$   OOOOOOOOOOOOOOO                        ",
		"                                    7$$$$$$$$$$$$$$$$$$$$   OOOOOOOOOOOOOOO                        ",
		"                                  $$$$$$$$$$$$$$$$$$$$$$   OOOOOOOOOOOOOO                          ",
		"                               I:   $$$$$$$$$$$$$$$$$$$$$7   OOOOOOOOOOOO?                         ",
		"                             7III,   $$$$$$$$$$$$$$$$$$$$$~   7OOOOOOOOO                           ",
		"                           =IIIIII,   $$$$$$$$$$$$$$$$$$$$$7     IOOO                              ",
		"                          IIIIIIIII,   $$$$$$$$$$$$$$$$$$$$$$=           7                         ",
		"                        IIIIIIIIIII,   $$$$$$77:, ::I$$$$$$$$$$$?~ +$$$$$$+                        ",
		"                        IIIIIIIIIIIII    $+               7$$$$$$$$$$$$$$$$$I                      ",
		"                       IIIIIIIIIIIIIII                       $$$$$$$$$$$$$$$$I                     ",
		"                      IIIIIIIIIIIIIIII                         7$$$$$$$$$$$$$$:                    ",
		"                     7IIIIIIIIIIIIII                             7$$$$$$$$$$$$$                    ",
		"                    :IIIIIIIIIIIIII                               $$$$$$$$$$$$$$                   ",
		"                    7IIIIIIIIIIIII                                 $$$$$$$$$$$$$~                  ",
		"                    =IIIIIIIIIIII                                   $$$$$$$$$$$$$                  ",
		"             I$$I     ~IIIIIIIII                                    +$$$$$$$$$$$$                  ",
		"          $7$$$$$$$7    IIIIIIII                                     $$$$$$$$$$$$?                 ",
		"         $$$$$$$$$$$$   ,IIIIII                                      $$$$$$$$$$$$$                 ",
		"        $$$$$$$$$$$$$$   IIIIII                                      ,$$$$$$$$$$$$                 ",
		"      ,$$$$$$$$$$$$$$   IIIIII                                                                     ",
		"       =$$$$$$$$$$$$$$   IIIIII                                                                    ",
		"        $$$$$$$$$$$$$$   7IIIII                                      IOOOOOOOOOOOO                 ",
		"        ?$$$$$$$$$$$$~   IIIIII+                                     OOOOOOOOOOOOO                 ",
		"         ,$$$$$$$$$$    IIIIIIII                                     OOOOOOOOOOOO:                 ",
		"            $$$$$$     IIIIIIIII=                                   OOOOOOOOOOOOO                  ",
		"                     IIIIIIIIIIII                                  ?OOOOOOOOOOOOO                  ",
		"                    IIIIIIIIIIIIII                                ,OOOOOOOOOOOOO                   ",
		"                     IIIIIIIIIIIIII                              IOOOOOOOOOOOOOI                   ",
		"                     :IIIIIIIIIIIIIII                           OOOOOOOOOOOOOOO                    ",
		"                      IIIIIIIIIIIIIIII=                       ZOOOOOOOOOOOOOOO                     ",
		"                       IIIIIIIIIIIIIII?                     OOOOOOOOOOOOOOOOO                      ",
		"                        ?IIIIIIIIIIIII   OO8,           IOOOOOOOOOOOOOOOOOOO                       ",
		"                         ~IIIIIIIIIII   ZOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO                        ",
		"                           7IIIIIIII   OOOOOOOOOOOOOOOOOOOOOOOOI       ,O=                         ",
		"                            ?IIIIII   ZOOOOOOOOOOOOOOOOOOOOOO                                      ",
		"                              IIII   $OOOOOOOOOOOOOOOOOOOOO=    IIIIIII,                           ",
		"                                I   ~OOOOOOOOOOOOOOOOOOOOO=   IIIIIIIIIII                          ",
		"                                   ~OOOOOOOOOOOOOOOOOOOOOO   IIIIIIIIIIIII                         ",
		"                                      8OOOOOOOOOOOOOOOOOO   ?IIIIIIIIIIIIII                        ",
		"                                          ,ZOOOOOOOOOOOOO   IIIIIIIIIIIIIII                        ",
		"                                                            IIIIIIIIIIIIIII                        ",
		"                                                             IIIIIIIIIIIIII                        ",
		"                                                             ,IIIIIIIIIIII                         ",
		"                                                               7IIIIIIII                           "}

	color.Set(color.FgRed, color.Bold)
	defer color.Unset()
	for i := 0; i < 10; i++ {
		fmt.Println()
	}
	for i := 0; i < len(ziko); i++ {
		fmt.Println(ziko[i])
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println()
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println()
	for i := 0; i < len(ce); i++ {
		fmt.Println(ce[i])
		time.Sleep(10 * time.Millisecond)
	}
	for i := 0; i < 10; i++ {
		fmt.Println()
	}
	text, _ = reader.ReadString('\n')
	fmt.Println()

	color.Set(color.FgGreen, color.Bold)
	defer color.Unset()
	for i := 0; i < len(android); i++ {
		fmt.Println(android[i])
		time.Sleep(10 * time.Millisecond)
	}
	text, _ = reader.ReadString('\n')
	fmt.Println()

	color.Set(color.FgCyan, color.Bold)
	defer color.Unset()
	for i := 0; i < len(ubuntu); i++ {
		fmt.Println(ubuntu[i])
		time.Sleep(10 * time.Millisecond)
	}
	text, _ = reader.ReadString('\n')
	fmt.Println()
	color.Unset()
	//	rhel := []string{
	fmt.Println("                                 +MMMMMMMMMMMMMMMMMMMMMMMMM~                                        ")
	fmt.Println("                             $MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMI                                    ")
	fmt.Println("                          MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN,                                ")
	fmt.Println("                       MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMD                              ")
	fmt.Println("                    =MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM,                           ")
	fmt.Println("                  MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNNMMMMMMMMMMMMMMMMMMMM                         ")
	fmt.Print("                NMMMMMMMMMMMMMMD")

	red := color.New(color.FgRed, color.Bold)
	//red.Print("This prints bold cyan %s\n", "too!.")

	red.Print("OOOOOOOO")
	fmt.Print("DNMMMMN8")
	red.Print("OOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMM$                      \n")
	fmt.Print("              OMMMMMMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMN")
	red.Print("OOOO")
	fmt.Print("MMMMMMMMMMMMMMMM~                    \n")
	fmt.Print("             MMMMMMMMMMMMMMMN")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMOOOMMMMMMMMMMMMMMMMM                   \n")
	fmt.Print("           $MMMMMMMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMD")
	red.Print("OOO")
	fmt.Print("MMMMMMMMMMMMMMMM:                 \n")
	fmt.Print("          MMMMMMMMMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMMM                \n")
	fmt.Print("        :MMMMMMMMMMMMMMMMMM")
	red.Print("OO")
	fmt.Print("MMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMM               \n")
	fmt.Print("       ,MMMMMMMMMMMMMMMMMMM")
	red.Print("OO")
	fmt.Print("MMMMMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMMM              \n")
	fmt.Print("       MMMMMMMMMMMMMMMMMMMM")
	red.Print("OOO")
	fmt.Print("MMMMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMMM             \n")
	fmt.Print("      MMMMMMMMMMMMMMMMMMMMD")
	red.Print("OOOOO")
	fmt.Print("MMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMMM            \n")
	fmt.Print("     MMMMMMMMMMMMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMMMM           \n")
	fmt.Print("    MMMMMMMMMMMMMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMMMM          \n")
	fmt.Print("   $MMMMMMMMMMMMMNDDDDDNMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMMMM          \n")
	fmt.Print("  +MMMMMMMMMD")
	red.Print("OOOOOOOOOO")
	fmt.Print("MMMD")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMMMM         \n")
	fmt.Print("  MMMMMMMM")
	red.Print("OOOOOOOOOOOOO")
	fmt.Print("MMMMMMD")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMMMMM        \n")
	fmt.Print(" IMMMMMMM")
	red.Print("OOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMD")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMMMM        \n")
	fmt.Print(" MMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMMMMM       \n")
	fmt.Print(" MMMMMMMMD")
	red.Print("OOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("M")
	red.Print("OOOOO")
	fmt.Print("MMMMMMMMMMMMM       \n")
	fmt.Print("7MMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MD")
	red.Print("OOOOOOO")
	fmt.Print("MMMMMMMMMMM       \n")
	fmt.Print("MMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMD")
	red.Print("OOOOOOOOOOOOOOOOOO")
	fmt.Print("N")
	red.Print("OOOOOOOOOOO")
	fmt.Print("MMMMMMMMMO      \n")
	fmt.Print("MMMMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMM      \n")
	fmt.Print("MMMMMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMM      \n")
	fmt.Print("MMMMMMMMMMMMMMMM8")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("D")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMM=     \n")
	fmt.Print("MMMMMMMMMMMMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMM7     \n")
	fmt.Print("MMMMMMMMMMMMMMMMMMMMD")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMM7     \n")
	fmt.Print("MMMMMMMMMMMMMMMMMMMMMMMN")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMI     \n")
	fmt.Print("MMMMMMMMMMMMMMMMMMMMMMMMMM8")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMM=     \n")
	fmt.Print("MMMMMMMMMMMMMMMMMMMMM:   ?MMMD")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMM      \n")
	fmt.Print("MMMMMMMMMMMMMMMMMMMM       MMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMM      \n")
	fmt.Print("MMMMMMMMMMMMMMMMMMMD        MMMMMMMMD")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMN      \n")
	fmt.Print("7MMMMMMMMMMMMMMMMMMM         ~M,  MMMMMMMM")
	red.Print("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMM       \n")
	fmt.Print(" MMMMMMMMMMMMMMMMMMM$              MMMMMMMMMMMMD")
	red.Print("OOOOOOOOOOOOOOOOOOOOOO")
	fmt.Print("MMMMMMMMMMMMMMMMMMMM       \n")
	fmt.Println(" MMMMMMMMMMMMMDMMMMMMD              NMMMMMMMMMMMMMMMMMMMMMMNNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM        ")
	fmt.Println(" IMMMMMMMMM       MMMMM              ,MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM         ")
	fmt.Println("                   MMMMM8               MMMMMMMMMMOI=DMMMM+,OMMMMM, MMMMMMMMMMMMMMMMMMMMMMM         ")
	fmt.Println("                     IMM?                                          ?MMMMMMMMMMMMMMMMMMMMMM          ")
	fmt.Println("                                                                   MMMMMMMMMMMMMMMMMMMMMM,          ")
	fmt.Println("                                                                  OMMMMMMMMMMMMMMMMMMMMMM           ")
	fmt.Println("                                                 MMM    M         MMMMMMMMMMMMMMMMMMMMMM            ")
	fmt.Println("                                                 MMMMMMMM         MMMMMMMMMMMMMMMMMMMMM,            ")
	fmt.Println("                                                     OO          IMMMMMMMMM7,      :7M              ")
	fmt.Println("                                                                ,MMMMMM                             ")
	fmt.Println("                                                               8MMMMMMM                             ")
	fmt.Println("                                                              NMMMMM,                               ")
	//fmt.Println("                                                             OMMMO                                  ")
	//fmt.Println("                                                            MMM8                                    ")
	//fmt.Println("                                                           MMM?                                     ")
	//fmt.Println("                                                          8~                                        ")
	text, _ = reader.ReadString('\n')
	color.Set(color.FgCyan, color.Bold)
	defer color.Unset()
	for i := 0; i < len(docker); i++ {
		fmt.Println(docker[i])
		time.Sleep(10 * time.Millisecond)
	}
	text, _ = reader.ReadString('\n')
	fmt.Println()
	fmt.Println(text)
}
