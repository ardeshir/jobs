package main

import(
  "fmt"
  "log"
  "os/exec"
)
  
  
func GetCmd() []string {

// Step one confure WinRM for ansible
//     return []string{"powershell", "iwr","https://raw.githubusercontent.com/ansible/ansible/devel/examples/scripts/ConfigureRemotingForAnsible.ps1","-UseBasicParsing","|","iex"}
// Might need this too.     
//   return []string{"powershell", "Set-ExecutionPolicy","-ExecutionPolicy","Unrestricted", "-Scope", "CurrentUser"}
     return []string{"powershell", "get-host" }
}

// https://blogs.technet.microsoft.com/heyscriptingguy/2015/06/11/table-of-basic-powershell-commands/
// https://github.com/ansible/ansible/blob/devel/examples/scripts/ConfigureRemotingForAnsible.ps1

func main() {
     
     cmd := GetCmd()
     out, err := exec.Command(cmd[0], cmd[1:]...).Output()
	 
     // out, err := exec.Command("go version").Output()
	 
	 if err != nil {
	      log.Fatal(err)
     }
	 
	 fmt.Printf("%s\n", out)
}

