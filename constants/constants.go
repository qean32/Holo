package constants

import (
	"os"
	"path/filepath"

	errorlist "main/constants/error-list"
	"main/constants/sgr"
)

const projectName = "holoproject"

var (
	Root = `private`
	Scenario  = `scenario`
)

func InitRoot() error {
	base := os.TempDir()
	if base == "" {
		return errorlist.ErrTempDirEmpty
	}
	Root = filepath.Join(base, projectName)
	if err := os.MkdirAll(Root, 0o755); err != nil {
		return errorlist.ErrInitRoot
	}
	return nil
}

var BinaryPROJECT_INIT = `
%v01101000 01101111 01101100 01101111 01110000 01110010 01101111 01101010 01100101 01100011 01110100
`

var PROJECT_INIT = `
%v __  __    ______    __        ______      ______  ______    ______      __    ______    ______    ______     
%v/\ \_\ \  /\  __ \  /\ \      /\  __ \    /\  == \/\  == \  /\  __ \    /\ \  /\  ___\  /\  ___\  /\__  _\    
%v\ \  __ \ \ \ \/\ \ \ \ \____ \ \ \/\ \   \ \  _-/\ \  __<  \ \ \/\ \  _\_\ \ \ \  __\  \ \ \____ \/_/\ \/    
%v \ \_\ \_\ \ \_____\ \ \_____\ \ \_____\   \ \_\   \ \_\ \_\ \ \_____\/\_____\ \ \_____\ \ \_____\   \ \_\    
%v  \/_/\/_/  \/_____/  \/_____/  \/_____/    \/_/    \/_/ /_/  \/_____/\/_____/  \/_____/  \/_____/    \/_/    
%v
`

// ASCII
// https://translated.turbopages.org/proxy_u/en-ru.ru.6c604ba4-6a11ef9b-b4e5d92b-74722d776562/https/student.cs.uwaterloo.ca/~cs452/terminal.html
// https://www.txtiiart.eu/text-to-ascii-art respect
// https://patorjk.com/software/taag/#p=testall&f=Broadway&t=holoproject+2&x=none&v=4&h=4&w=80&we=false

var HelpMessage = `
%v  __  __    ______    __        ______      __  __    ______    __        ______  
%v /\ \_\ \  /\  __ \  /\ \      /\  __ \    /\ \_\ \  /\  ___\  /\ \      /\  == \ 
%v \ \  __ \ \ \ \/\ \ \ \ \____ \ \ \/\ \   \ \  __ \ \ \  __\  \ \ \____ \ \  _-/ 
%v  \ \_\ \_\ \ \_____\ \ \_____\ \ \_____\   \ \_\ \_\ \ \_____\ \ \_____\ \ \_\   
%v   \/_/\/_/  \/_____/  \/_____/  \/_____/    \/_/\/_/  \/_____/  \/_____/  \/_/   
%v
%v crypto         encryption
%v ecrypto        decryption
%v
%v gmaster        generate master key
%v gkey           generate key
%v
%v test           unit tests
%v
%v
%v menu, clog, log, drop, stop, help,
%v
%v sc             add scenario                          sc keyWord -p:payload
%v list           list scenarios
%v run            run scenario                          run keyWord
%v rsc            remove scenario                       rsc keyWord
%v
`

const CHAR_SELECTED_ITEM = "> "
const CHAR_UN_SELECTED_ITEM = "│  "

const FIELDPREFIX = "> "

var StyleDim = []int{sgr.DIM}

var AcceptableYeaOrNot = []string{"yes", "y", "yea", "да"}
