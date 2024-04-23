// Code generated from /home/jefferson/Escritorio/MIA/MIA_P1_202112030/src/Language/Parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Parser

import "github.com/antlr4-go/antlr/v4"

// ParserListener is a complete listener for a parse tree produced by ParserParser.
type ParserListener interface {
	antlr.ParseTreeListener

	// EnterInit is called when entering the init production.
	EnterInit(c *InitContext)

	// EnterCommands is called when entering the commands production.
	EnterCommands(c *CommandsContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterExecute is called when entering the execute production.
	EnterExecute(c *ExecuteContext)

	// EnterMkdisk is called when entering the mkdisk production.
	EnterMkdisk(c *MkdiskContext)

	// EnterMkdiskparams is called when entering the mkdiskparams production.
	EnterMkdiskparams(c *MkdiskparamsContext)

	// EnterMkdiskparam is called when entering the mkdiskparam production.
	EnterMkdiskparam(c *MkdiskparamContext)

	// EnterRmdisk is called when entering the rmdisk production.
	EnterRmdisk(c *RmdiskContext)

	// EnterFdisk is called when entering the fdisk production.
	EnterFdisk(c *FdiskContext)

	// EnterFdiskparams is called when entering the fdiskparams production.
	EnterFdiskparams(c *FdiskparamsContext)

	// EnterFdiskparam is called when entering the fdiskparam production.
	EnterFdiskparam(c *FdiskparamContext)

	// EnterMount is called when entering the mount production.
	EnterMount(c *MountContext)

	// EnterMountparams is called when entering the mountparams production.
	EnterMountparams(c *MountparamsContext)

	// EnterMountparam is called when entering the mountparam production.
	EnterMountparam(c *MountparamContext)

	// EnterUnmount is called when entering the unmount production.
	EnterUnmount(c *UnmountContext)

	// EnterMkfs is called when entering the mkfs production.
	EnterMkfs(c *MkfsContext)

	// EnterMkfsparams is called when entering the mkfsparams production.
	EnterMkfsparams(c *MkfsparamsContext)

	// EnterMkfsparam is called when entering the mkfsparam production.
	EnterMkfsparam(c *MkfsparamContext)

	// EnterLogin is called when entering the login production.
	EnterLogin(c *LoginContext)

	// EnterLoginparams is called when entering the loginparams production.
	EnterLoginparams(c *LoginparamsContext)

	// EnterLoginparam is called when entering the loginparam production.
	EnterLoginparam(c *LoginparamContext)

	// EnterLogout is called when entering the logout production.
	EnterLogout(c *LogoutContext)

	// EnterPause is called when entering the pause production.
	EnterPause(c *PauseContext)

	// EnterMkgrp is called when entering the mkgrp production.
	EnterMkgrp(c *MkgrpContext)

	// EnterMkusr is called when entering the mkusr production.
	EnterMkusr(c *MkusrContext)

	// EnterMkuserparams is called when entering the mkuserparams production.
	EnterMkuserparams(c *MkuserparamsContext)

	// EnterMkuserparam is called when entering the mkuserparam production.
	EnterMkuserparam(c *MkuserparamContext)

	// EnterMkfile is called when entering the mkfile production.
	EnterMkfile(c *MkfileContext)

	// EnterMkfileparams is called when entering the mkfileparams production.
	EnterMkfileparams(c *MkfileparamsContext)

	// EnterMkfileparam is called when entering the mkfileparam production.
	EnterMkfileparam(c *MkfileparamContext)

	// EnterMkdir is called when entering the mkdir production.
	EnterMkdir(c *MkdirContext)

	// EnterMkdirparams is called when entering the mkdirparams production.
	EnterMkdirparams(c *MkdirparamsContext)

	// EnterMkdirparam is called when entering the mkdirparam production.
	EnterMkdirparam(c *MkdirparamContext)

	// EnterCat is called when entering the cat production.
	EnterCat(c *CatContext)

	// EnterCatfiles is called when entering the catfiles production.
	EnterCatfiles(c *CatfilesContext)

	// EnterCatfile is called when entering the catfile production.
	EnterCatfile(c *CatfileContext)

	// EnterRep is called when entering the rep production.
	EnterRep(c *RepContext)

	// EnterRepparams is called when entering the repparams production.
	EnterRepparams(c *RepparamsContext)

	// EnterRepparam is called when entering the repparam production.
	EnterRepparam(c *RepparamContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterCommentary is called when entering the commentary production.
	EnterCommentary(c *CommentaryContext)

	// EnterExit is called when entering the exit production.
	EnterExit(c *ExitContext)

	// ExitInit is called when exiting the init production.
	ExitInit(c *InitContext)

	// ExitCommands is called when exiting the commands production.
	ExitCommands(c *CommandsContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitExecute is called when exiting the execute production.
	ExitExecute(c *ExecuteContext)

	// ExitMkdisk is called when exiting the mkdisk production.
	ExitMkdisk(c *MkdiskContext)

	// ExitMkdiskparams is called when exiting the mkdiskparams production.
	ExitMkdiskparams(c *MkdiskparamsContext)

	// ExitMkdiskparam is called when exiting the mkdiskparam production.
	ExitMkdiskparam(c *MkdiskparamContext)

	// ExitRmdisk is called when exiting the rmdisk production.
	ExitRmdisk(c *RmdiskContext)

	// ExitFdisk is called when exiting the fdisk production.
	ExitFdisk(c *FdiskContext)

	// ExitFdiskparams is called when exiting the fdiskparams production.
	ExitFdiskparams(c *FdiskparamsContext)

	// ExitFdiskparam is called when exiting the fdiskparam production.
	ExitFdiskparam(c *FdiskparamContext)

	// ExitMount is called when exiting the mount production.
	ExitMount(c *MountContext)

	// ExitMountparams is called when exiting the mountparams production.
	ExitMountparams(c *MountparamsContext)

	// ExitMountparam is called when exiting the mountparam production.
	ExitMountparam(c *MountparamContext)

	// ExitUnmount is called when exiting the unmount production.
	ExitUnmount(c *UnmountContext)

	// ExitMkfs is called when exiting the mkfs production.
	ExitMkfs(c *MkfsContext)

	// ExitMkfsparams is called when exiting the mkfsparams production.
	ExitMkfsparams(c *MkfsparamsContext)

	// ExitMkfsparam is called when exiting the mkfsparam production.
	ExitMkfsparam(c *MkfsparamContext)

	// ExitLogin is called when exiting the login production.
	ExitLogin(c *LoginContext)

	// ExitLoginparams is called when exiting the loginparams production.
	ExitLoginparams(c *LoginparamsContext)

	// ExitLoginparam is called when exiting the loginparam production.
	ExitLoginparam(c *LoginparamContext)

	// ExitLogout is called when exiting the logout production.
	ExitLogout(c *LogoutContext)

	// ExitPause is called when exiting the pause production.
	ExitPause(c *PauseContext)

	// ExitMkgrp is called when exiting the mkgrp production.
	ExitMkgrp(c *MkgrpContext)

	// ExitMkusr is called when exiting the mkusr production.
	ExitMkusr(c *MkusrContext)

	// ExitMkuserparams is called when exiting the mkuserparams production.
	ExitMkuserparams(c *MkuserparamsContext)

	// ExitMkuserparam is called when exiting the mkuserparam production.
	ExitMkuserparam(c *MkuserparamContext)

	// ExitMkfile is called when exiting the mkfile production.
	ExitMkfile(c *MkfileContext)

	// ExitMkfileparams is called when exiting the mkfileparams production.
	ExitMkfileparams(c *MkfileparamsContext)

	// ExitMkfileparam is called when exiting the mkfileparam production.
	ExitMkfileparam(c *MkfileparamContext)

	// ExitMkdir is called when exiting the mkdir production.
	ExitMkdir(c *MkdirContext)

	// ExitMkdirparams is called when exiting the mkdirparams production.
	ExitMkdirparams(c *MkdirparamsContext)

	// ExitMkdirparam is called when exiting the mkdirparam production.
	ExitMkdirparam(c *MkdirparamContext)

	// ExitCat is called when exiting the cat production.
	ExitCat(c *CatContext)

	// ExitCatfiles is called when exiting the catfiles production.
	ExitCatfiles(c *CatfilesContext)

	// ExitCatfile is called when exiting the catfile production.
	ExitCatfile(c *CatfileContext)

	// ExitRep is called when exiting the rep production.
	ExitRep(c *RepContext)

	// ExitRepparams is called when exiting the repparams production.
	ExitRepparams(c *RepparamsContext)

	// ExitRepparam is called when exiting the repparam production.
	ExitRepparam(c *RepparamContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitCommentary is called when exiting the commentary production.
	ExitCommentary(c *CommentaryContext)

	// ExitExit is called when exiting the exit production.
	ExitExit(c *ExitContext)
}
