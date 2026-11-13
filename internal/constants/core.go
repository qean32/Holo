package constants

import (
	"os"
	"path/filepath"

	errorlist "github.com/qean32/holo/internal/constants/error-list"
	"github.com/qean32/holo/internal/constants/sgr"
)

const projectName = "holoproject"

var (
	Root   = `private`
	Script = `script`
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
	if err := os.MkdirAll(filepath.Join(Root, Script), 0o755); err != nil {
		return errorlist.ErrInitRoot
	}
	return nil
}

var BinaryPROJECT_INIT = `
%v01100110 01101111 01111000 00100000 01101000 01101001 01100100 01101001 01101110 01100111
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
%v list, clog, log, drop, stop, help,
%v
%v sc             add script                          sc keyWord -p:payload
%v list           list scripts
%v run            run script                          run keyWord
%v rsc            remove script                       rsc keyWord
%v
`

const CHAR_SELECTED_ITEM = "> "
const CHAR_UN_SELECTED_ITEM = "│  "

const FIELDPREFIX = "> "

var StyleDim = []int{sgr.DIM}

var AcceptableYeaOrNot = []string{"yes", "y", "yea", "да"}
