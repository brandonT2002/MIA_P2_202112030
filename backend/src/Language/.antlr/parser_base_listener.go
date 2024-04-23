// Code generated from /home/jefferson/Escritorio/MIA/MIA_P1_202112030/src/Language/Parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Parser

import "github.com/antlr4-go/antlr/v4"

// BaseParserListener is a complete listener for a parse tree produced by ParserParser.
type BaseParserListener struct{}

var _ ParserListener = &BaseParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInit is called when production init is entered.
func (s *BaseParserListener) EnterInit(ctx *InitContext) {}

// ExitInit is called when production init is exited.
func (s *BaseParserListener) ExitInit(ctx *InitContext) {}

// EnterCommands is called when production commands is entered.
func (s *BaseParserListener) EnterCommands(ctx *CommandsContext) {}

// ExitCommands is called when production commands is exited.
func (s *BaseParserListener) ExitCommands(ctx *CommandsContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseParserListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseParserListener) ExitCommand(ctx *CommandContext) {}

// EnterExecute is called when production execute is entered.
func (s *BaseParserListener) EnterExecute(ctx *ExecuteContext) {}

// ExitExecute is called when production execute is exited.
func (s *BaseParserListener) ExitExecute(ctx *ExecuteContext) {}

// EnterMkdisk is called when production mkdisk is entered.
func (s *BaseParserListener) EnterMkdisk(ctx *MkdiskContext) {}

// ExitMkdisk is called when production mkdisk is exited.
func (s *BaseParserListener) ExitMkdisk(ctx *MkdiskContext) {}

// EnterMkdiskparams is called when production mkdiskparams is entered.
func (s *BaseParserListener) EnterMkdiskparams(ctx *MkdiskparamsContext) {}

// ExitMkdiskparams is called when production mkdiskparams is exited.
func (s *BaseParserListener) ExitMkdiskparams(ctx *MkdiskparamsContext) {}

// EnterMkdiskparam is called when production mkdiskparam is entered.
func (s *BaseParserListener) EnterMkdiskparam(ctx *MkdiskparamContext) {}

// ExitMkdiskparam is called when production mkdiskparam is exited.
func (s *BaseParserListener) ExitMkdiskparam(ctx *MkdiskparamContext) {}

// EnterRmdisk is called when production rmdisk is entered.
func (s *BaseParserListener) EnterRmdisk(ctx *RmdiskContext) {}

// ExitRmdisk is called when production rmdisk is exited.
func (s *BaseParserListener) ExitRmdisk(ctx *RmdiskContext) {}

// EnterFdisk is called when production fdisk is entered.
func (s *BaseParserListener) EnterFdisk(ctx *FdiskContext) {}

// ExitFdisk is called when production fdisk is exited.
func (s *BaseParserListener) ExitFdisk(ctx *FdiskContext) {}

// EnterFdiskparams is called when production fdiskparams is entered.
func (s *BaseParserListener) EnterFdiskparams(ctx *FdiskparamsContext) {}

// ExitFdiskparams is called when production fdiskparams is exited.
func (s *BaseParserListener) ExitFdiskparams(ctx *FdiskparamsContext) {}

// EnterFdiskparam is called when production fdiskparam is entered.
func (s *BaseParserListener) EnterFdiskparam(ctx *FdiskparamContext) {}

// ExitFdiskparam is called when production fdiskparam is exited.
func (s *BaseParserListener) ExitFdiskparam(ctx *FdiskparamContext) {}

// EnterMount is called when production mount is entered.
func (s *BaseParserListener) EnterMount(ctx *MountContext) {}

// ExitMount is called when production mount is exited.
func (s *BaseParserListener) ExitMount(ctx *MountContext) {}

// EnterMountparams is called when production mountparams is entered.
func (s *BaseParserListener) EnterMountparams(ctx *MountparamsContext) {}

// ExitMountparams is called when production mountparams is exited.
func (s *BaseParserListener) ExitMountparams(ctx *MountparamsContext) {}

// EnterMountparam is called when production mountparam is entered.
func (s *BaseParserListener) EnterMountparam(ctx *MountparamContext) {}

// ExitMountparam is called when production mountparam is exited.
func (s *BaseParserListener) ExitMountparam(ctx *MountparamContext) {}

// EnterUnmount is called when production unmount is entered.
func (s *BaseParserListener) EnterUnmount(ctx *UnmountContext) {}

// ExitUnmount is called when production unmount is exited.
func (s *BaseParserListener) ExitUnmount(ctx *UnmountContext) {}

// EnterMkfs is called when production mkfs is entered.
func (s *BaseParserListener) EnterMkfs(ctx *MkfsContext) {}

// ExitMkfs is called when production mkfs is exited.
func (s *BaseParserListener) ExitMkfs(ctx *MkfsContext) {}

// EnterMkfsparams is called when production mkfsparams is entered.
func (s *BaseParserListener) EnterMkfsparams(ctx *MkfsparamsContext) {}

// ExitMkfsparams is called when production mkfsparams is exited.
func (s *BaseParserListener) ExitMkfsparams(ctx *MkfsparamsContext) {}

// EnterMkfsparam is called when production mkfsparam is entered.
func (s *BaseParserListener) EnterMkfsparam(ctx *MkfsparamContext) {}

// ExitMkfsparam is called when production mkfsparam is exited.
func (s *BaseParserListener) ExitMkfsparam(ctx *MkfsparamContext) {}

// EnterLogin is called when production login is entered.
func (s *BaseParserListener) EnterLogin(ctx *LoginContext) {}

// ExitLogin is called when production login is exited.
func (s *BaseParserListener) ExitLogin(ctx *LoginContext) {}

// EnterLoginparams is called when production loginparams is entered.
func (s *BaseParserListener) EnterLoginparams(ctx *LoginparamsContext) {}

// ExitLoginparams is called when production loginparams is exited.
func (s *BaseParserListener) ExitLoginparams(ctx *LoginparamsContext) {}

// EnterLoginparam is called when production loginparam is entered.
func (s *BaseParserListener) EnterLoginparam(ctx *LoginparamContext) {}

// ExitLoginparam is called when production loginparam is exited.
func (s *BaseParserListener) ExitLoginparam(ctx *LoginparamContext) {}

// EnterLogout is called when production logout is entered.
func (s *BaseParserListener) EnterLogout(ctx *LogoutContext) {}

// ExitLogout is called when production logout is exited.
func (s *BaseParserListener) ExitLogout(ctx *LogoutContext) {}

// EnterPause is called when production pause is entered.
func (s *BaseParserListener) EnterPause(ctx *PauseContext) {}

// ExitPause is called when production pause is exited.
func (s *BaseParserListener) ExitPause(ctx *PauseContext) {}

// EnterMkgrp is called when production mkgrp is entered.
func (s *BaseParserListener) EnterMkgrp(ctx *MkgrpContext) {}

// ExitMkgrp is called when production mkgrp is exited.
func (s *BaseParserListener) ExitMkgrp(ctx *MkgrpContext) {}

// EnterMkusr is called when production mkusr is entered.
func (s *BaseParserListener) EnterMkusr(ctx *MkusrContext) {}

// ExitMkusr is called when production mkusr is exited.
func (s *BaseParserListener) ExitMkusr(ctx *MkusrContext) {}

// EnterMkuserparams is called when production mkuserparams is entered.
func (s *BaseParserListener) EnterMkuserparams(ctx *MkuserparamsContext) {}

// ExitMkuserparams is called when production mkuserparams is exited.
func (s *BaseParserListener) ExitMkuserparams(ctx *MkuserparamsContext) {}

// EnterMkuserparam is called when production mkuserparam is entered.
func (s *BaseParserListener) EnterMkuserparam(ctx *MkuserparamContext) {}

// ExitMkuserparam is called when production mkuserparam is exited.
func (s *BaseParserListener) ExitMkuserparam(ctx *MkuserparamContext) {}

// EnterMkfile is called when production mkfile is entered.
func (s *BaseParserListener) EnterMkfile(ctx *MkfileContext) {}

// ExitMkfile is called when production mkfile is exited.
func (s *BaseParserListener) ExitMkfile(ctx *MkfileContext) {}

// EnterMkfileparams is called when production mkfileparams is entered.
func (s *BaseParserListener) EnterMkfileparams(ctx *MkfileparamsContext) {}

// ExitMkfileparams is called when production mkfileparams is exited.
func (s *BaseParserListener) ExitMkfileparams(ctx *MkfileparamsContext) {}

// EnterMkfileparam is called when production mkfileparam is entered.
func (s *BaseParserListener) EnterMkfileparam(ctx *MkfileparamContext) {}

// ExitMkfileparam is called when production mkfileparam is exited.
func (s *BaseParserListener) ExitMkfileparam(ctx *MkfileparamContext) {}

// EnterMkdir is called when production mkdir is entered.
func (s *BaseParserListener) EnterMkdir(ctx *MkdirContext) {}

// ExitMkdir is called when production mkdir is exited.
func (s *BaseParserListener) ExitMkdir(ctx *MkdirContext) {}

// EnterMkdirparams is called when production mkdirparams is entered.
func (s *BaseParserListener) EnterMkdirparams(ctx *MkdirparamsContext) {}

// ExitMkdirparams is called when production mkdirparams is exited.
func (s *BaseParserListener) ExitMkdirparams(ctx *MkdirparamsContext) {}

// EnterMkdirparam is called when production mkdirparam is entered.
func (s *BaseParserListener) EnterMkdirparam(ctx *MkdirparamContext) {}

// ExitMkdirparam is called when production mkdirparam is exited.
func (s *BaseParserListener) ExitMkdirparam(ctx *MkdirparamContext) {}

// EnterCat is called when production cat is entered.
func (s *BaseParserListener) EnterCat(ctx *CatContext) {}

// ExitCat is called when production cat is exited.
func (s *BaseParserListener) ExitCat(ctx *CatContext) {}

// EnterCatfiles is called when production catfiles is entered.
func (s *BaseParserListener) EnterCatfiles(ctx *CatfilesContext) {}

// ExitCatfiles is called when production catfiles is exited.
func (s *BaseParserListener) ExitCatfiles(ctx *CatfilesContext) {}

// EnterCatfile is called when production catfile is entered.
func (s *BaseParserListener) EnterCatfile(ctx *CatfileContext) {}

// ExitCatfile is called when production catfile is exited.
func (s *BaseParserListener) ExitCatfile(ctx *CatfileContext) {}

// EnterRep is called when production rep is entered.
func (s *BaseParserListener) EnterRep(ctx *RepContext) {}

// ExitRep is called when production rep is exited.
func (s *BaseParserListener) ExitRep(ctx *RepContext) {}

// EnterRepparams is called when production repparams is entered.
func (s *BaseParserListener) EnterRepparams(ctx *RepparamsContext) {}

// ExitRepparams is called when production repparams is exited.
func (s *BaseParserListener) ExitRepparams(ctx *RepparamsContext) {}

// EnterRepparam is called when production repparam is entered.
func (s *BaseParserListener) EnterRepparam(ctx *RepparamContext) {}

// ExitRepparam is called when production repparam is exited.
func (s *BaseParserListener) ExitRepparam(ctx *RepparamContext) {}

// EnterName is called when production name is entered.
func (s *BaseParserListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseParserListener) ExitName(ctx *NameContext) {}

// EnterCommentary is called when production commentary is entered.
func (s *BaseParserListener) EnterCommentary(ctx *CommentaryContext) {}

// ExitCommentary is called when production commentary is exited.
func (s *BaseParserListener) ExitCommentary(ctx *CommentaryContext) {}

// EnterExit is called when production exit is entered.
func (s *BaseParserListener) EnterExit(ctx *ExitContext) {}

// ExitExit is called when production exit is exited.
func (s *BaseParserListener) ExitExit(ctx *ExitContext) {}
