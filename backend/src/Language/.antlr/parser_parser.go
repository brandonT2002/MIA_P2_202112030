// Code generated from /home/jefferson/Escritorio/MIA/MIA_P1_202112030/src/Language/Parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Parser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import (
	commands "mia/Classes/Commands"
	interfaces "mia/Classes/Interfaces"

	"os"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ParserParser struct {
	*antlr.BaseParser
}

var ParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func parserParserInit() {
	staticData := &ParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'mkdisk'", "'rmdisk'", "'fdisk'", "'mount'", "'unmount'", "'mkfs'",
		"'login'", "'logout'", "'mkgrp'", "'rmgrp'", "'mkusr'", "'rmusr'", "'mkfile'",
		"'cat'", "'remove'", "'edit'", "'rename'", "'mkdir'", "'copy'", "'move'",
		"'find'", "'chown'", "'chgrp'", "'chmod'", "'pause'", "'recovery'",
		"'loss'", "'execute'", "'rep'", "'exit'", "'mbr'", "'disk'", "'inode'",
		"'journaling'", "'block'", "'bm_inode'", "'bm_block'", "'tree'", "'sb'",
		"'file'", "'ls'", "'full'", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"'='", "", "", "'\\n'",
	}
	staticData.SymbolicNames = []string{
		"", "RW_mkdisk", "RW_rmdisk", "RW_fdisk", "RW_mount", "RW_unmount",
		"RW_mkfs", "RW_login", "RW_logout", "RW_mkgrp", "RW_rmgrp", "RW_mkusr",
		"RW_rmusr", "RW_mkfile", "RW_cat", "RW_remove", "RW_edit", "RW_rename",
		"RW_mkdir", "RW_copy", "RW_move", "RW_find", "RW_chown", "RW_chgrp",
		"RW_chmod", "RW_pause", "RW_recovery", "RW_loss", "RW_execute", "RW_rep",
		"RW_exit", "RW_mbr", "RW_disk", "RW_inode", "RW_journaling", "RW_block",
		"RW_bm_inode", "RW_bm_block", "RW_tree", "RW_sb", "RW_file", "RW_ls",
		"RW_full", "RW_size", "RW_fit", "RW_unit", "RW_driveletter", "RW_name",
		"RW_type", "RW_delete", "RW_add", "RW_id", "RW_fs", "RW_user", "RW_pass",
		"RW_grp", "RW_path", "RW_r", "RW_cont", "RW_fileN", "RW_destino", "RW_ugo",
		"RW_ruta", "TK_fit", "TK_unit", "TK_type", "TK_fs", "TK_number", "TK_id",
		"TK_path", "TK_equ", "TK_sym", "COMMENTARY", "NEWLINE", "UNUSED_",
	}
	staticData.RuleNames = []string{
		"init", "commands", "command", "execute", "mkdisk", "mkdiskparams",
		"mkdiskparam", "rmdisk", "fdisk", "fdiskparams", "fdiskparam", "mount",
		"mountparams", "mountparam", "unmount", "mkfs", "mkfsparams", "mkfsparam",
		"login", "loginparams", "loginparam", "logout", "pause", "mkgrp", "mkusr",
		"mkuserparams", "mkuserparam", "mkfile", "mkfileparams", "mkfileparam",
		"mkdir", "mkdirparams", "mkdirparam", "cat", "catfiles", "catfile",
		"rep", "repparams", "repparam", "name", "commentary", "exit",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 74, 594, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 91, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 101, 8, 1, 10, 1, 12, 1, 104, 9, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 3, 2, 158, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 167, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 175, 8,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 185, 8, 5, 10,
		5, 12, 5, 188, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 3, 6, 202, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 3, 7, 211, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 219,
		8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 229, 8, 9,
		10, 9, 12, 9, 232, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 266, 8, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 3, 11, 274, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 5, 12, 284, 8, 12, 10, 12, 12, 12, 287, 9, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 297, 8, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 306, 8, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 314, 8, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 324, 8, 16, 10, 16, 12,
		16, 327, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 341, 8, 17, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 3, 18, 349, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 5, 19, 359, 8, 19, 10, 19, 12, 19, 362, 9, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 380, 8, 20, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		3, 23, 395, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 403,
		8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 413,
		8, 25, 10, 25, 12, 25, 416, 9, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 430, 8, 26, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 438, 8, 27, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 448, 8, 28, 10, 28, 12, 28,
		451, 9, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 467, 8, 29, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 475, 8, 30, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 485, 8, 31, 10, 31, 12, 31, 488,
		9, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 496, 8, 32, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 504, 8, 33, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 514, 8, 34, 10, 34, 12,
		34, 517, 9, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 3, 36, 530, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 5, 37, 540, 8, 37, 10, 37, 12, 37, 543, 9, 37,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 562, 8, 38, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		3, 39, 586, 8, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 0,
		11, 2, 10, 18, 24, 32, 38, 50, 56, 62, 68, 74, 42, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 0,
		1, 2, 0, 64, 65, 68, 68, 628, 0, 90, 1, 0, 0, 0, 2, 92, 1, 0, 0, 0, 4,
		157, 1, 0, 0, 0, 6, 166, 1, 0, 0, 0, 8, 174, 1, 0, 0, 0, 10, 176, 1, 0,
		0, 0, 12, 201, 1, 0, 0, 0, 14, 210, 1, 0, 0, 0, 16, 218, 1, 0, 0, 0, 18,
		220, 1, 0, 0, 0, 20, 265, 1, 0, 0, 0, 22, 273, 1, 0, 0, 0, 24, 275, 1,
		0, 0, 0, 26, 296, 1, 0, 0, 0, 28, 305, 1, 0, 0, 0, 30, 313, 1, 0, 0, 0,
		32, 315, 1, 0, 0, 0, 34, 340, 1, 0, 0, 0, 36, 348, 1, 0, 0, 0, 38, 350,
		1, 0, 0, 0, 40, 379, 1, 0, 0, 0, 42, 381, 1, 0, 0, 0, 44, 384, 1, 0, 0,
		0, 46, 394, 1, 0, 0, 0, 48, 402, 1, 0, 0, 0, 50, 404, 1, 0, 0, 0, 52, 429,
		1, 0, 0, 0, 54, 437, 1, 0, 0, 0, 56, 439, 1, 0, 0, 0, 58, 466, 1, 0, 0,
		0, 60, 474, 1, 0, 0, 0, 62, 476, 1, 0, 0, 0, 64, 495, 1, 0, 0, 0, 66, 503,
		1, 0, 0, 0, 68, 505, 1, 0, 0, 0, 70, 518, 1, 0, 0, 0, 72, 529, 1, 0, 0,
		0, 74, 531, 1, 0, 0, 0, 76, 561, 1, 0, 0, 0, 78, 585, 1, 0, 0, 0, 80, 587,
		1, 0, 0, 0, 82, 590, 1, 0, 0, 0, 84, 85, 3, 2, 1, 0, 85, 86, 5, 0, 0, 1,
		86, 87, 6, 0, -1, 0, 87, 91, 1, 0, 0, 0, 88, 89, 5, 0, 0, 1, 89, 91, 6,
		0, -1, 0, 90, 84, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 1, 1, 0, 0, 0, 92,
		93, 6, 1, -1, 0, 93, 94, 3, 4, 2, 0, 94, 95, 6, 1, -1, 0, 95, 102, 1, 0,
		0, 0, 96, 97, 10, 2, 0, 0, 97, 98, 3, 4, 2, 0, 98, 99, 6, 1, -1, 0, 99,
		101, 1, 0, 0, 0, 100, 96, 1, 0, 0, 0, 101, 104, 1, 0, 0, 0, 102, 100, 1,
		0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 3, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0,
		105, 106, 3, 6, 3, 0, 106, 107, 6, 2, -1, 0, 107, 158, 1, 0, 0, 0, 108,
		109, 3, 8, 4, 0, 109, 110, 6, 2, -1, 0, 110, 158, 1, 0, 0, 0, 111, 112,
		3, 14, 7, 0, 112, 113, 6, 2, -1, 0, 113, 158, 1, 0, 0, 0, 114, 115, 3,
		16, 8, 0, 115, 116, 6, 2, -1, 0, 116, 158, 1, 0, 0, 0, 117, 118, 3, 22,
		11, 0, 118, 119, 6, 2, -1, 0, 119, 158, 1, 0, 0, 0, 120, 121, 3, 28, 14,
		0, 121, 122, 6, 2, -1, 0, 122, 158, 1, 0, 0, 0, 123, 124, 3, 30, 15, 0,
		124, 125, 6, 2, -1, 0, 125, 158, 1, 0, 0, 0, 126, 127, 3, 36, 18, 0, 127,
		128, 6, 2, -1, 0, 128, 158, 1, 0, 0, 0, 129, 130, 3, 42, 21, 0, 130, 131,
		6, 2, -1, 0, 131, 158, 1, 0, 0, 0, 132, 133, 3, 46, 23, 0, 133, 134, 6,
		2, -1, 0, 134, 158, 1, 0, 0, 0, 135, 136, 3, 48, 24, 0, 136, 137, 6, 2,
		-1, 0, 137, 158, 1, 0, 0, 0, 138, 139, 3, 54, 27, 0, 139, 140, 6, 2, -1,
		0, 140, 158, 1, 0, 0, 0, 141, 142, 3, 60, 30, 0, 142, 143, 6, 2, -1, 0,
		143, 158, 1, 0, 0, 0, 144, 145, 3, 66, 33, 0, 145, 146, 6, 2, -1, 0, 146,
		158, 1, 0, 0, 0, 147, 148, 3, 44, 22, 0, 148, 149, 6, 2, -1, 0, 149, 158,
		1, 0, 0, 0, 150, 151, 3, 72, 36, 0, 151, 152, 6, 2, -1, 0, 152, 158, 1,
		0, 0, 0, 153, 154, 3, 80, 40, 0, 154, 155, 6, 2, -1, 0, 155, 158, 1, 0,
		0, 0, 156, 158, 3, 82, 41, 0, 157, 105, 1, 0, 0, 0, 157, 108, 1, 0, 0,
		0, 157, 111, 1, 0, 0, 0, 157, 114, 1, 0, 0, 0, 157, 117, 1, 0, 0, 0, 157,
		120, 1, 0, 0, 0, 157, 123, 1, 0, 0, 0, 157, 126, 1, 0, 0, 0, 157, 129,
		1, 0, 0, 0, 157, 132, 1, 0, 0, 0, 157, 135, 1, 0, 0, 0, 157, 138, 1, 0,
		0, 0, 157, 141, 1, 0, 0, 0, 157, 144, 1, 0, 0, 0, 157, 147, 1, 0, 0, 0,
		157, 150, 1, 0, 0, 0, 157, 153, 1, 0, 0, 0, 157, 156, 1, 0, 0, 0, 158,
		5, 1, 0, 0, 0, 159, 160, 5, 28, 0, 0, 160, 161, 5, 56, 0, 0, 161, 162,
		5, 70, 0, 0, 162, 163, 5, 69, 0, 0, 163, 167, 6, 3, -1, 0, 164, 165, 5,
		28, 0, 0, 165, 167, 6, 3, -1, 0, 166, 159, 1, 0, 0, 0, 166, 164, 1, 0,
		0, 0, 167, 7, 1, 0, 0, 0, 168, 169, 5, 1, 0, 0, 169, 170, 3, 10, 5, 0,
		170, 171, 6, 4, -1, 0, 171, 175, 1, 0, 0, 0, 172, 173, 5, 1, 0, 0, 173,
		175, 6, 4, -1, 0, 174, 168, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 175, 9, 1,
		0, 0, 0, 176, 177, 6, 5, -1, 0, 177, 178, 3, 12, 6, 0, 178, 179, 6, 5,
		-1, 0, 179, 186, 1, 0, 0, 0, 180, 181, 10, 2, 0, 0, 181, 182, 3, 12, 6,
		0, 182, 183, 6, 5, -1, 0, 183, 185, 1, 0, 0, 0, 184, 180, 1, 0, 0, 0, 185,
		188, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 11, 1,
		0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 190, 5, 43, 0, 0, 190, 191, 5, 70,
		0, 0, 191, 192, 5, 67, 0, 0, 192, 202, 6, 6, -1, 0, 193, 194, 5, 44, 0,
		0, 194, 195, 5, 70, 0, 0, 195, 196, 5, 63, 0, 0, 196, 202, 6, 6, -1, 0,
		197, 198, 5, 45, 0, 0, 198, 199, 5, 70, 0, 0, 199, 200, 5, 64, 0, 0, 200,
		202, 6, 6, -1, 0, 201, 189, 1, 0, 0, 0, 201, 193, 1, 0, 0, 0, 201, 197,
		1, 0, 0, 0, 202, 13, 1, 0, 0, 0, 203, 204, 5, 2, 0, 0, 204, 205, 5, 46,
		0, 0, 205, 206, 5, 70, 0, 0, 206, 207, 7, 0, 0, 0, 207, 211, 6, 7, -1,
		0, 208, 209, 5, 2, 0, 0, 209, 211, 6, 7, -1, 0, 210, 203, 1, 0, 0, 0, 210,
		208, 1, 0, 0, 0, 211, 15, 1, 0, 0, 0, 212, 213, 5, 3, 0, 0, 213, 214, 3,
		18, 9, 0, 214, 215, 6, 8, -1, 0, 215, 219, 1, 0, 0, 0, 216, 217, 5, 3,
		0, 0, 217, 219, 6, 8, -1, 0, 218, 212, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0,
		219, 17, 1, 0, 0, 0, 220, 221, 6, 9, -1, 0, 221, 222, 3, 20, 10, 0, 222,
		223, 6, 9, -1, 0, 223, 230, 1, 0, 0, 0, 224, 225, 10, 2, 0, 0, 225, 226,
		3, 20, 10, 0, 226, 227, 6, 9, -1, 0, 227, 229, 1, 0, 0, 0, 228, 224, 1,
		0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0,
		0, 231, 19, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 233, 234, 5, 43, 0, 0, 234,
		235, 5, 70, 0, 0, 235, 236, 5, 67, 0, 0, 236, 266, 6, 10, -1, 0, 237, 238,
		5, 46, 0, 0, 238, 239, 5, 70, 0, 0, 239, 240, 7, 0, 0, 0, 240, 266, 6,
		10, -1, 0, 241, 242, 5, 47, 0, 0, 242, 243, 5, 70, 0, 0, 243, 244, 5, 68,
		0, 0, 244, 266, 6, 10, -1, 0, 245, 246, 5, 45, 0, 0, 246, 247, 5, 70, 0,
		0, 247, 248, 5, 64, 0, 0, 248, 266, 6, 10, -1, 0, 249, 250, 5, 48, 0, 0,
		250, 251, 5, 70, 0, 0, 251, 252, 5, 65, 0, 0, 252, 266, 6, 10, -1, 0, 253,
		254, 5, 44, 0, 0, 254, 255, 5, 70, 0, 0, 255, 256, 5, 63, 0, 0, 256, 266,
		6, 10, -1, 0, 257, 258, 5, 49, 0, 0, 258, 259, 5, 70, 0, 0, 259, 260, 5,
		42, 0, 0, 260, 266, 6, 10, -1, 0, 261, 262, 5, 50, 0, 0, 262, 263, 5, 70,
		0, 0, 263, 264, 5, 67, 0, 0, 264, 266, 6, 10, -1, 0, 265, 233, 1, 0, 0,
		0, 265, 237, 1, 0, 0, 0, 265, 241, 1, 0, 0, 0, 265, 245, 1, 0, 0, 0, 265,
		249, 1, 0, 0, 0, 265, 253, 1, 0, 0, 0, 265, 257, 1, 0, 0, 0, 265, 261,
		1, 0, 0, 0, 266, 21, 1, 0, 0, 0, 267, 268, 5, 4, 0, 0, 268, 269, 3, 24,
		12, 0, 269, 270, 6, 11, -1, 0, 270, 274, 1, 0, 0, 0, 271, 272, 5, 4, 0,
		0, 272, 274, 6, 11, -1, 0, 273, 267, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0,
		274, 23, 1, 0, 0, 0, 275, 276, 6, 12, -1, 0, 276, 277, 3, 26, 13, 0, 277,
		278, 6, 12, -1, 0, 278, 285, 1, 0, 0, 0, 279, 280, 10, 2, 0, 0, 280, 281,
		3, 26, 13, 0, 281, 282, 6, 12, -1, 0, 282, 284, 1, 0, 0, 0, 283, 279, 1,
		0, 0, 0, 284, 287, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0,
		0, 286, 25, 1, 0, 0, 0, 287, 285, 1, 0, 0, 0, 288, 289, 5, 46, 0, 0, 289,
		290, 5, 70, 0, 0, 290, 291, 7, 0, 0, 0, 291, 297, 6, 13, -1, 0, 292, 293,
		5, 47, 0, 0, 293, 294, 5, 70, 0, 0, 294, 295, 5, 68, 0, 0, 295, 297, 6,
		13, -1, 0, 296, 288, 1, 0, 0, 0, 296, 292, 1, 0, 0, 0, 297, 27, 1, 0, 0,
		0, 298, 299, 5, 5, 0, 0, 299, 300, 5, 51, 0, 0, 300, 301, 5, 70, 0, 0,
		301, 302, 5, 68, 0, 0, 302, 306, 6, 14, -1, 0, 303, 304, 5, 5, 0, 0, 304,
		306, 6, 14, -1, 0, 305, 298, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 306, 29,
		1, 0, 0, 0, 307, 308, 5, 6, 0, 0, 308, 309, 3, 32, 16, 0, 309, 310, 6,
		15, -1, 0, 310, 314, 1, 0, 0, 0, 311, 312, 5, 6, 0, 0, 312, 314, 6, 15,
		-1, 0, 313, 307, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 314, 31, 1, 0, 0, 0,
		315, 316, 6, 16, -1, 0, 316, 317, 3, 34, 17, 0, 317, 318, 6, 16, -1, 0,
		318, 325, 1, 0, 0, 0, 319, 320, 10, 2, 0, 0, 320, 321, 3, 34, 17, 0, 321,
		322, 6, 16, -1, 0, 322, 324, 1, 0, 0, 0, 323, 319, 1, 0, 0, 0, 324, 327,
		1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 33, 1, 0,
		0, 0, 327, 325, 1, 0, 0, 0, 328, 329, 5, 51, 0, 0, 329, 330, 5, 70, 0,
		0, 330, 331, 5, 68, 0, 0, 331, 341, 6, 17, -1, 0, 332, 333, 5, 48, 0, 0,
		333, 334, 5, 70, 0, 0, 334, 335, 5, 42, 0, 0, 335, 341, 6, 17, -1, 0, 336,
		337, 5, 52, 0, 0, 337, 338, 5, 70, 0, 0, 338, 339, 5, 66, 0, 0, 339, 341,
		6, 17, -1, 0, 340, 328, 1, 0, 0, 0, 340, 332, 1, 0, 0, 0, 340, 336, 1,
		0, 0, 0, 341, 35, 1, 0, 0, 0, 342, 343, 5, 7, 0, 0, 343, 344, 3, 38, 19,
		0, 344, 345, 6, 18, -1, 0, 345, 349, 1, 0, 0, 0, 346, 347, 5, 7, 0, 0,
		347, 349, 6, 18, -1, 0, 348, 342, 1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 349,
		37, 1, 0, 0, 0, 350, 351, 6, 19, -1, 0, 351, 352, 3, 40, 20, 0, 352, 353,
		6, 19, -1, 0, 353, 360, 1, 0, 0, 0, 354, 355, 10, 2, 0, 0, 355, 356, 3,
		40, 20, 0, 356, 357, 6, 19, -1, 0, 357, 359, 1, 0, 0, 0, 358, 354, 1, 0,
		0, 0, 359, 362, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0,
		361, 39, 1, 0, 0, 0, 362, 360, 1, 0, 0, 0, 363, 364, 5, 53, 0, 0, 364,
		365, 5, 70, 0, 0, 365, 366, 5, 68, 0, 0, 366, 380, 6, 20, -1, 0, 367, 368,
		5, 54, 0, 0, 368, 369, 5, 70, 0, 0, 369, 370, 5, 68, 0, 0, 370, 380, 6,
		20, -1, 0, 371, 372, 5, 54, 0, 0, 372, 373, 5, 70, 0, 0, 373, 374, 5, 67,
		0, 0, 374, 380, 6, 20, -1, 0, 375, 376, 5, 51, 0, 0, 376, 377, 5, 70, 0,
		0, 377, 378, 5, 68, 0, 0, 378, 380, 6, 20, -1, 0, 379, 363, 1, 0, 0, 0,
		379, 367, 1, 0, 0, 0, 379, 371, 1, 0, 0, 0, 379, 375, 1, 0, 0, 0, 380,
		41, 1, 0, 0, 0, 381, 382, 5, 8, 0, 0, 382, 383, 6, 21, -1, 0, 383, 43,
		1, 0, 0, 0, 384, 385, 5, 25, 0, 0, 385, 386, 6, 22, -1, 0, 386, 45, 1,
		0, 0, 0, 387, 388, 5, 9, 0, 0, 388, 389, 5, 47, 0, 0, 389, 390, 5, 70,
		0, 0, 390, 391, 5, 68, 0, 0, 391, 395, 6, 23, -1, 0, 392, 393, 5, 9, 0,
		0, 393, 395, 6, 23, -1, 0, 394, 387, 1, 0, 0, 0, 394, 392, 1, 0, 0, 0,
		395, 47, 1, 0, 0, 0, 396, 397, 5, 11, 0, 0, 397, 398, 3, 50, 25, 0, 398,
		399, 6, 24, -1, 0, 399, 403, 1, 0, 0, 0, 400, 401, 5, 11, 0, 0, 401, 403,
		6, 24, -1, 0, 402, 396, 1, 0, 0, 0, 402, 400, 1, 0, 0, 0, 403, 49, 1, 0,
		0, 0, 404, 405, 6, 25, -1, 0, 405, 406, 3, 52, 26, 0, 406, 407, 6, 25,
		-1, 0, 407, 414, 1, 0, 0, 0, 408, 409, 10, 2, 0, 0, 409, 410, 3, 52, 26,
		0, 410, 411, 6, 25, -1, 0, 411, 413, 1, 0, 0, 0, 412, 408, 1, 0, 0, 0,
		413, 416, 1, 0, 0, 0, 414, 412, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415,
		51, 1, 0, 0, 0, 416, 414, 1, 0, 0, 0, 417, 418, 5, 53, 0, 0, 418, 419,
		5, 70, 0, 0, 419, 420, 5, 68, 0, 0, 420, 430, 6, 26, -1, 0, 421, 422, 5,
		54, 0, 0, 422, 423, 5, 70, 0, 0, 423, 424, 5, 68, 0, 0, 424, 430, 6, 26,
		-1, 0, 425, 426, 5, 55, 0, 0, 426, 427, 5, 70, 0, 0, 427, 428, 5, 68, 0,
		0, 428, 430, 6, 26, -1, 0, 429, 417, 1, 0, 0, 0, 429, 421, 1, 0, 0, 0,
		429, 425, 1, 0, 0, 0, 430, 53, 1, 0, 0, 0, 431, 432, 5, 13, 0, 0, 432,
		433, 3, 56, 28, 0, 433, 434, 6, 27, -1, 0, 434, 438, 1, 0, 0, 0, 435, 436,
		5, 13, 0, 0, 436, 438, 6, 27, -1, 0, 437, 431, 1, 0, 0, 0, 437, 435, 1,
		0, 0, 0, 438, 55, 1, 0, 0, 0, 439, 440, 6, 28, -1, 0, 440, 441, 3, 58,
		29, 0, 441, 442, 6, 28, -1, 0, 442, 449, 1, 0, 0, 0, 443, 444, 10, 2, 0,
		0, 444, 445, 3, 58, 29, 0, 445, 446, 6, 28, -1, 0, 446, 448, 1, 0, 0, 0,
		447, 443, 1, 0, 0, 0, 448, 451, 1, 0, 0, 0, 449, 447, 1, 0, 0, 0, 449,
		450, 1, 0, 0, 0, 450, 57, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 452, 453, 5,
		56, 0, 0, 453, 454, 5, 70, 0, 0, 454, 455, 5, 69, 0, 0, 455, 467, 6, 29,
		-1, 0, 456, 457, 5, 43, 0, 0, 457, 458, 5, 70, 0, 0, 458, 459, 5, 67, 0,
		0, 459, 467, 6, 29, -1, 0, 460, 461, 5, 58, 0, 0, 461, 462, 5, 70, 0, 0,
		462, 463, 5, 69, 0, 0, 463, 467, 6, 29, -1, 0, 464, 465, 5, 57, 0, 0, 465,
		467, 6, 29, -1, 0, 466, 452, 1, 0, 0, 0, 466, 456, 1, 0, 0, 0, 466, 460,
		1, 0, 0, 0, 466, 464, 1, 0, 0, 0, 467, 59, 1, 0, 0, 0, 468, 469, 5, 18,
		0, 0, 469, 470, 3, 62, 31, 0, 470, 471, 6, 30, -1, 0, 471, 475, 1, 0, 0,
		0, 472, 473, 5, 18, 0, 0, 473, 475, 6, 30, -1, 0, 474, 468, 1, 0, 0, 0,
		474, 472, 1, 0, 0, 0, 475, 61, 1, 0, 0, 0, 476, 477, 6, 31, -1, 0, 477,
		478, 3, 64, 32, 0, 478, 479, 6, 31, -1, 0, 479, 486, 1, 0, 0, 0, 480, 481,
		10, 2, 0, 0, 481, 482, 3, 64, 32, 0, 482, 483, 6, 31, -1, 0, 483, 485,
		1, 0, 0, 0, 484, 480, 1, 0, 0, 0, 485, 488, 1, 0, 0, 0, 486, 484, 1, 0,
		0, 0, 486, 487, 1, 0, 0, 0, 487, 63, 1, 0, 0, 0, 488, 486, 1, 0, 0, 0,
		489, 490, 5, 56, 0, 0, 490, 491, 5, 70, 0, 0, 491, 492, 5, 69, 0, 0, 492,
		496, 6, 32, -1, 0, 493, 494, 5, 57, 0, 0, 494, 496, 6, 32, -1, 0, 495,
		489, 1, 0, 0, 0, 495, 493, 1, 0, 0, 0, 496, 65, 1, 0, 0, 0, 497, 498, 5,
		14, 0, 0, 498, 499, 3, 68, 34, 0, 499, 500, 6, 33, -1, 0, 500, 504, 1,
		0, 0, 0, 501, 502, 5, 14, 0, 0, 502, 504, 6, 33, -1, 0, 503, 497, 1, 0,
		0, 0, 503, 501, 1, 0, 0, 0, 504, 67, 1, 0, 0, 0, 505, 506, 6, 34, -1, 0,
		506, 507, 3, 70, 35, 0, 507, 508, 6, 34, -1, 0, 508, 515, 1, 0, 0, 0, 509,
		510, 10, 2, 0, 0, 510, 511, 3, 70, 35, 0, 511, 512, 6, 34, -1, 0, 512,
		514, 1, 0, 0, 0, 513, 509, 1, 0, 0, 0, 514, 517, 1, 0, 0, 0, 515, 513,
		1, 0, 0, 0, 515, 516, 1, 0, 0, 0, 516, 69, 1, 0, 0, 0, 517, 515, 1, 0,
		0, 0, 518, 519, 5, 59, 0, 0, 519, 520, 5, 70, 0, 0, 520, 521, 5, 69, 0,
		0, 521, 522, 6, 35, -1, 0, 522, 71, 1, 0, 0, 0, 523, 524, 5, 29, 0, 0,
		524, 525, 3, 74, 37, 0, 525, 526, 6, 36, -1, 0, 526, 530, 1, 0, 0, 0, 527,
		528, 5, 29, 0, 0, 528, 530, 6, 36, -1, 0, 529, 523, 1, 0, 0, 0, 529, 527,
		1, 0, 0, 0, 530, 73, 1, 0, 0, 0, 531, 532, 6, 37, -1, 0, 532, 533, 3, 76,
		38, 0, 533, 534, 6, 37, -1, 0, 534, 541, 1, 0, 0, 0, 535, 536, 10, 2, 0,
		0, 536, 537, 3, 76, 38, 0, 537, 538, 6, 37, -1, 0, 538, 540, 1, 0, 0, 0,
		539, 535, 1, 0, 0, 0, 540, 543, 1, 0, 0, 0, 541, 539, 1, 0, 0, 0, 541,
		542, 1, 0, 0, 0, 542, 75, 1, 0, 0, 0, 543, 541, 1, 0, 0, 0, 544, 545, 5,
		47, 0, 0, 545, 546, 5, 70, 0, 0, 546, 547, 3, 78, 39, 0, 547, 548, 6, 38,
		-1, 0, 548, 562, 1, 0, 0, 0, 549, 550, 5, 56, 0, 0, 550, 551, 5, 70, 0,
		0, 551, 552, 5, 69, 0, 0, 552, 562, 6, 38, -1, 0, 553, 554, 5, 62, 0, 0,
		554, 555, 5, 70, 0, 0, 555, 556, 5, 69, 0, 0, 556, 562, 6, 38, -1, 0, 557,
		558, 5, 51, 0, 0, 558, 559, 5, 70, 0, 0, 559, 560, 5, 68, 0, 0, 560, 562,
		6, 38, -1, 0, 561, 544, 1, 0, 0, 0, 561, 549, 1, 0, 0, 0, 561, 553, 1,
		0, 0, 0, 561, 557, 1, 0, 0, 0, 562, 77, 1, 0, 0, 0, 563, 564, 5, 31, 0,
		0, 564, 586, 6, 39, -1, 0, 565, 566, 5, 32, 0, 0, 566, 586, 6, 39, -1,
		0, 567, 568, 5, 33, 0, 0, 568, 586, 6, 39, -1, 0, 569, 570, 5, 34, 0, 0,
		570, 586, 6, 39, -1, 0, 571, 572, 5, 35, 0, 0, 572, 586, 6, 39, -1, 0,
		573, 574, 5, 36, 0, 0, 574, 586, 6, 39, -1, 0, 575, 576, 5, 37, 0, 0, 576,
		586, 6, 39, -1, 0, 577, 578, 5, 38, 0, 0, 578, 586, 6, 39, -1, 0, 579,
		580, 5, 39, 0, 0, 580, 586, 6, 39, -1, 0, 581, 582, 5, 40, 0, 0, 582, 586,
		6, 39, -1, 0, 583, 584, 5, 41, 0, 0, 584, 586, 6, 39, -1, 0, 585, 563,
		1, 0, 0, 0, 585, 565, 1, 0, 0, 0, 585, 567, 1, 0, 0, 0, 585, 569, 1, 0,
		0, 0, 585, 571, 1, 0, 0, 0, 585, 573, 1, 0, 0, 0, 585, 575, 1, 0, 0, 0,
		585, 577, 1, 0, 0, 0, 585, 579, 1, 0, 0, 0, 585, 581, 1, 0, 0, 0, 585,
		583, 1, 0, 0, 0, 586, 79, 1, 0, 0, 0, 587, 588, 5, 72, 0, 0, 588, 589,
		6, 40, -1, 0, 589, 81, 1, 0, 0, 0, 590, 591, 5, 30, 0, 0, 591, 592, 6,
		41, -1, 0, 592, 83, 1, 0, 0, 0, 37, 90, 102, 157, 166, 174, 186, 201, 210,
		218, 230, 265, 273, 285, 296, 305, 313, 325, 340, 348, 360, 379, 394, 402,
		414, 429, 437, 449, 466, 474, 486, 495, 503, 515, 529, 541, 561, 585,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ParserParserInit initializes any static state used to implement ParserParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewParserParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ParserParserInit() {
	staticData := &ParserParserStaticData
	staticData.once.Do(parserParserInit)
}

// NewParserParser produces a new parser instance for the optional input antlr.TokenStream.
func NewParserParser(input antlr.TokenStream) *ParserParser {
	ParserParserInit()
	this := new(ParserParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Parser.g4"

	return this
}

// ParserParser tokens.
const (
	ParserParserEOF            = antlr.TokenEOF
	ParserParserRW_mkdisk      = 1
	ParserParserRW_rmdisk      = 2
	ParserParserRW_fdisk       = 3
	ParserParserRW_mount       = 4
	ParserParserRW_unmount     = 5
	ParserParserRW_mkfs        = 6
	ParserParserRW_login       = 7
	ParserParserRW_logout      = 8
	ParserParserRW_mkgrp       = 9
	ParserParserRW_rmgrp       = 10
	ParserParserRW_mkusr       = 11
	ParserParserRW_rmusr       = 12
	ParserParserRW_mkfile      = 13
	ParserParserRW_cat         = 14
	ParserParserRW_remove      = 15
	ParserParserRW_edit        = 16
	ParserParserRW_rename      = 17
	ParserParserRW_mkdir       = 18
	ParserParserRW_copy        = 19
	ParserParserRW_move        = 20
	ParserParserRW_find        = 21
	ParserParserRW_chown       = 22
	ParserParserRW_chgrp       = 23
	ParserParserRW_chmod       = 24
	ParserParserRW_pause       = 25
	ParserParserRW_recovery    = 26
	ParserParserRW_loss        = 27
	ParserParserRW_execute     = 28
	ParserParserRW_rep         = 29
	ParserParserRW_exit        = 30
	ParserParserRW_mbr         = 31
	ParserParserRW_disk        = 32
	ParserParserRW_inode       = 33
	ParserParserRW_journaling  = 34
	ParserParserRW_block       = 35
	ParserParserRW_bm_inode    = 36
	ParserParserRW_bm_block    = 37
	ParserParserRW_tree        = 38
	ParserParserRW_sb          = 39
	ParserParserRW_file        = 40
	ParserParserRW_ls          = 41
	ParserParserRW_full        = 42
	ParserParserRW_size        = 43
	ParserParserRW_fit         = 44
	ParserParserRW_unit        = 45
	ParserParserRW_driveletter = 46
	ParserParserRW_name        = 47
	ParserParserRW_type        = 48
	ParserParserRW_delete      = 49
	ParserParserRW_add         = 50
	ParserParserRW_id          = 51
	ParserParserRW_fs          = 52
	ParserParserRW_user        = 53
	ParserParserRW_pass        = 54
	ParserParserRW_grp         = 55
	ParserParserRW_path        = 56
	ParserParserRW_r           = 57
	ParserParserRW_cont        = 58
	ParserParserRW_fileN       = 59
	ParserParserRW_destino     = 60
	ParserParserRW_ugo         = 61
	ParserParserRW_ruta        = 62
	ParserParserTK_fit         = 63
	ParserParserTK_unit        = 64
	ParserParserTK_type        = 65
	ParserParserTK_fs          = 66
	ParserParserTK_number      = 67
	ParserParserTK_id          = 68
	ParserParserTK_path        = 69
	ParserParserTK_equ         = 70
	ParserParserTK_sym         = 71
	ParserParserCOMMENTARY     = 72
	ParserParserNEWLINE        = 73
	ParserParserUNUSED_        = 74
)

// ParserParser rules.
const (
	ParserParserRULE_init         = 0
	ParserParserRULE_commands     = 1
	ParserParserRULE_command      = 2
	ParserParserRULE_execute      = 3
	ParserParserRULE_mkdisk       = 4
	ParserParserRULE_mkdiskparams = 5
	ParserParserRULE_mkdiskparam  = 6
	ParserParserRULE_rmdisk       = 7
	ParserParserRULE_fdisk        = 8
	ParserParserRULE_fdiskparams  = 9
	ParserParserRULE_fdiskparam   = 10
	ParserParserRULE_mount        = 11
	ParserParserRULE_mountparams  = 12
	ParserParserRULE_mountparam   = 13
	ParserParserRULE_unmount      = 14
	ParserParserRULE_mkfs         = 15
	ParserParserRULE_mkfsparams   = 16
	ParserParserRULE_mkfsparam    = 17
	ParserParserRULE_login        = 18
	ParserParserRULE_loginparams  = 19
	ParserParserRULE_loginparam   = 20
	ParserParserRULE_logout       = 21
	ParserParserRULE_pause        = 22
	ParserParserRULE_mkgrp        = 23
	ParserParserRULE_mkusr        = 24
	ParserParserRULE_mkuserparams = 25
	ParserParserRULE_mkuserparam  = 26
	ParserParserRULE_mkfile       = 27
	ParserParserRULE_mkfileparams = 28
	ParserParserRULE_mkfileparam  = 29
	ParserParserRULE_mkdir        = 30
	ParserParserRULE_mkdirparams  = 31
	ParserParserRULE_mkdirparam   = 32
	ParserParserRULE_cat          = 33
	ParserParserRULE_catfiles     = 34
	ParserParserRULE_catfile      = 35
	ParserParserRULE_rep          = 36
	ParserParserRULE_repparams    = 37
	ParserParserRULE_repparam     = 38
	ParserParserRULE_name         = 39
	ParserParserRULE_commentary   = 40
	ParserParserRULE_exit         = 41
)

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC returns the c rule contexts.
	GetC() ICommandsContext

	// SetC sets the c rule contexts.
	SetC(ICommandsContext)

	// GetResult returns the result attribute.
	GetResult() []interfaces.Command

	// SetResult sets the result attribute.
	SetResult([]interfaces.Command)

	// Getter signatures
	EOF() antlr.TerminalNode
	Commands() ICommandsContext

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []interfaces.Command
	c      ICommandsContext
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_init
	return p
}

func InitEmptyInitContext(p *InitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_init
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) GetC() ICommandsContext { return s.c }

func (s *InitContext) SetC(v ICommandsContext) { s.c = v }

func (s *InitContext) GetResult() []interfaces.Command { return s.result }

func (s *InitContext) SetResult(v []interfaces.Command) { s.result = v }

func (s *InitContext) EOF() antlr.TerminalNode {
	return s.GetToken(ParserParserEOF, 0)
}

func (s *InitContext) Commands() ICommandsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandsContext)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Init() (localctx IInitContext) {
	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParserParserRULE_init)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_mkdisk, ParserParserRW_rmdisk, ParserParserRW_fdisk, ParserParserRW_mount, ParserParserRW_unmount, ParserParserRW_mkfs, ParserParserRW_login, ParserParserRW_logout, ParserParserRW_mkgrp, ParserParserRW_mkusr, ParserParserRW_mkfile, ParserParserRW_cat, ParserParserRW_mkdir, ParserParserRW_pause, ParserParserRW_execute, ParserParserRW_rep, ParserParserRW_exit, ParserParserCOMMENTARY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)

			var _x = p.commands(0)

			localctx.(*InitContext).c = _x
		}
		{
			p.SetState(85)
			p.Match(ParserParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*InitContext).result = localctx.(*InitContext).GetC().GetResult()

	case ParserParserEOF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)
			p.Match(ParserParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*InitContext).result = []interfaces.Command{}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandsContext is an interface to support dynamic dispatch.
type ICommandsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() ICommandsContext

	// GetC returns the c rule contexts.
	GetC() ICommandContext

	// SetL sets the l rule contexts.
	SetL(ICommandsContext)

	// SetC sets the c rule contexts.
	SetC(ICommandContext)

	// GetResult returns the result attribute.
	GetResult() []interfaces.Command

	// SetResult sets the result attribute.
	SetResult([]interfaces.Command)

	// Getter signatures
	Command() ICommandContext
	Commands() ICommandsContext

	// IsCommandsContext differentiates from other interfaces.
	IsCommandsContext()
}

type CommandsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []interfaces.Command
	l      ICommandsContext
	c      ICommandContext
}

func NewEmptyCommandsContext() *CommandsContext {
	var p = new(CommandsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_commands
	return p
}

func InitEmptyCommandsContext(p *CommandsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_commands
}

func (*CommandsContext) IsCommandsContext() {}

func NewCommandsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandsContext {
	var p = new(CommandsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_commands

	return p
}

func (s *CommandsContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandsContext) GetL() ICommandsContext { return s.l }

func (s *CommandsContext) GetC() ICommandContext { return s.c }

func (s *CommandsContext) SetL(v ICommandsContext) { s.l = v }

func (s *CommandsContext) SetC(v ICommandContext) { s.c = v }

func (s *CommandsContext) GetResult() []interfaces.Command { return s.result }

func (s *CommandsContext) SetResult(v []interfaces.Command) { s.result = v }

func (s *CommandsContext) Command() ICommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *CommandsContext) Commands() ICommandsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandsContext)
}

func (s *CommandsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Commands() (localctx ICommandsContext) {
	return p.commands(0)
}

func (p *ParserParser) commands(_p int) (localctx ICommandsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewCommandsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICommandsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ParserParserRULE_commands, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)

		var _x = p.Command()

		localctx.(*CommandsContext).c = _x
	}
	localctx.(*CommandsContext).result = []interfaces.Command{localctx.(*CommandsContext).GetC().GetResult()}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCommandsContext(p, _parentctx, _parentState)
			localctx.(*CommandsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_commands)
			p.SetState(96)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(97)

				var _x = p.Command()

				localctx.(*CommandsContext).c = _x
			}
			localctx.(*CommandsContext).SetResult(localctx.(*CommandsContext).GetL().GetResult())
			localctx.(*CommandsContext).result = append(localctx.(*CommandsContext).result, localctx.(*CommandsContext).GetC().GetResult())

		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC1 returns the c1 rule contexts.
	GetC1() IExecuteContext

	// GetC2 returns the c2 rule contexts.
	GetC2() IMkdiskContext

	// GetC3 returns the c3 rule contexts.
	GetC3() IRmdiskContext

	// GetC4 returns the c4 rule contexts.
	GetC4() IFdiskContext

	// GetC5 returns the c5 rule contexts.
	GetC5() IMountContext

	// GetC6 returns the c6 rule contexts.
	GetC6() IUnmountContext

	// GetC7 returns the c7 rule contexts.
	GetC7() IMkfsContext

	// GetC8 returns the c8 rule contexts.
	GetC8() ILoginContext

	// GetC9 returns the c9 rule contexts.
	GetC9() ILogoutContext

	// GetC10 returns the c10 rule contexts.
	GetC10() IMkgrpContext

	// GetC11 returns the c11 rule contexts.
	GetC11() IMkusrContext

	// GetC12 returns the c12 rule contexts.
	GetC12() IMkfileContext

	// GetC13 returns the c13 rule contexts.
	GetC13() IMkdirContext

	// GetC18 returns the c18 rule contexts.
	GetC18() ICatContext

	// GetC19 returns the c19 rule contexts.
	GetC19() IPauseContext

	// GetC20 returns the c20 rule contexts.
	GetC20() IRepContext

	// GetC21 returns the c21 rule contexts.
	GetC21() ICommentaryContext

	// SetC1 sets the c1 rule contexts.
	SetC1(IExecuteContext)

	// SetC2 sets the c2 rule contexts.
	SetC2(IMkdiskContext)

	// SetC3 sets the c3 rule contexts.
	SetC3(IRmdiskContext)

	// SetC4 sets the c4 rule contexts.
	SetC4(IFdiskContext)

	// SetC5 sets the c5 rule contexts.
	SetC5(IMountContext)

	// SetC6 sets the c6 rule contexts.
	SetC6(IUnmountContext)

	// SetC7 sets the c7 rule contexts.
	SetC7(IMkfsContext)

	// SetC8 sets the c8 rule contexts.
	SetC8(ILoginContext)

	// SetC9 sets the c9 rule contexts.
	SetC9(ILogoutContext)

	// SetC10 sets the c10 rule contexts.
	SetC10(IMkgrpContext)

	// SetC11 sets the c11 rule contexts.
	SetC11(IMkusrContext)

	// SetC12 sets the c12 rule contexts.
	SetC12(IMkfileContext)

	// SetC13 sets the c13 rule contexts.
	SetC13(IMkdirContext)

	// SetC18 sets the c18 rule contexts.
	SetC18(ICatContext)

	// SetC19 sets the c19 rule contexts.
	SetC19(IPauseContext)

	// SetC20 sets the c20 rule contexts.
	SetC20(IRepContext)

	// SetC21 sets the c21 rule contexts.
	SetC21(ICommentaryContext)

	// GetResult returns the result attribute.
	GetResult() interfaces.Command

	// SetResult sets the result attribute.
	SetResult(interfaces.Command)

	// Getter signatures
	Execute() IExecuteContext
	Mkdisk() IMkdiskContext
	Rmdisk() IRmdiskContext
	Fdisk() IFdiskContext
	Mount() IMountContext
	Unmount() IUnmountContext
	Mkfs() IMkfsContext
	Login() ILoginContext
	Logout() ILogoutContext
	Mkgrp() IMkgrpContext
	Mkusr() IMkusrContext
	Mkfile() IMkfileContext
	Mkdir() IMkdirContext
	Cat() ICatContext
	Pause() IPauseContext
	Rep() IRepContext
	Commentary() ICommentaryContext
	Exit() IExitContext

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result interfaces.Command
	c1     IExecuteContext
	c2     IMkdiskContext
	c3     IRmdiskContext
	c4     IFdiskContext
	c5     IMountContext
	c6     IUnmountContext
	c7     IMkfsContext
	c8     ILoginContext
	c9     ILogoutContext
	c10    IMkgrpContext
	c11    IMkusrContext
	c12    IMkfileContext
	c13    IMkdirContext
	c18    ICatContext
	c19    IPauseContext
	c20    IRepContext
	c21    ICommentaryContext
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) GetC1() IExecuteContext { return s.c1 }

func (s *CommandContext) GetC2() IMkdiskContext { return s.c2 }

func (s *CommandContext) GetC3() IRmdiskContext { return s.c3 }

func (s *CommandContext) GetC4() IFdiskContext { return s.c4 }

func (s *CommandContext) GetC5() IMountContext { return s.c5 }

func (s *CommandContext) GetC6() IUnmountContext { return s.c6 }

func (s *CommandContext) GetC7() IMkfsContext { return s.c7 }

func (s *CommandContext) GetC8() ILoginContext { return s.c8 }

func (s *CommandContext) GetC9() ILogoutContext { return s.c9 }

func (s *CommandContext) GetC10() IMkgrpContext { return s.c10 }

func (s *CommandContext) GetC11() IMkusrContext { return s.c11 }

func (s *CommandContext) GetC12() IMkfileContext { return s.c12 }

func (s *CommandContext) GetC13() IMkdirContext { return s.c13 }

func (s *CommandContext) GetC18() ICatContext { return s.c18 }

func (s *CommandContext) GetC19() IPauseContext { return s.c19 }

func (s *CommandContext) GetC20() IRepContext { return s.c20 }

func (s *CommandContext) GetC21() ICommentaryContext { return s.c21 }

func (s *CommandContext) SetC1(v IExecuteContext) { s.c1 = v }

func (s *CommandContext) SetC2(v IMkdiskContext) { s.c2 = v }

func (s *CommandContext) SetC3(v IRmdiskContext) { s.c3 = v }

func (s *CommandContext) SetC4(v IFdiskContext) { s.c4 = v }

func (s *CommandContext) SetC5(v IMountContext) { s.c5 = v }

func (s *CommandContext) SetC6(v IUnmountContext) { s.c6 = v }

func (s *CommandContext) SetC7(v IMkfsContext) { s.c7 = v }

func (s *CommandContext) SetC8(v ILoginContext) { s.c8 = v }

func (s *CommandContext) SetC9(v ILogoutContext) { s.c9 = v }

func (s *CommandContext) SetC10(v IMkgrpContext) { s.c10 = v }

func (s *CommandContext) SetC11(v IMkusrContext) { s.c11 = v }

func (s *CommandContext) SetC12(v IMkfileContext) { s.c12 = v }

func (s *CommandContext) SetC13(v IMkdirContext) { s.c13 = v }

func (s *CommandContext) SetC18(v ICatContext) { s.c18 = v }

func (s *CommandContext) SetC19(v IPauseContext) { s.c19 = v }

func (s *CommandContext) SetC20(v IRepContext) { s.c20 = v }

func (s *CommandContext) SetC21(v ICommentaryContext) { s.c21 = v }

func (s *CommandContext) GetResult() interfaces.Command { return s.result }

func (s *CommandContext) SetResult(v interfaces.Command) { s.result = v }

func (s *CommandContext) Execute() IExecuteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExecuteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExecuteContext)
}

func (s *CommandContext) Mkdisk() IMkdiskContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskContext)
}

func (s *CommandContext) Rmdisk() IRmdiskContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmdiskContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmdiskContext)
}

func (s *CommandContext) Fdisk() IFdiskContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskContext)
}

func (s *CommandContext) Mount() IMountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountContext)
}

func (s *CommandContext) Unmount() IUnmountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnmountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnmountContext)
}

func (s *CommandContext) Mkfs() IMkfsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfsContext)
}

func (s *CommandContext) Login() ILoginContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoginContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoginContext)
}

func (s *CommandContext) Logout() ILogoutContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogoutContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogoutContext)
}

func (s *CommandContext) Mkgrp() IMkgrpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkgrpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkgrpContext)
}

func (s *CommandContext) Mkusr() IMkusrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkusrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkusrContext)
}

func (s *CommandContext) Mkfile() IMkfileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfileContext)
}

func (s *CommandContext) Mkdir() IMkdirContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdirContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdirContext)
}

func (s *CommandContext) Cat() ICatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatContext)
}

func (s *CommandContext) Pause() IPauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPauseContext)
}

func (s *CommandContext) Rep() IRepContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepContext)
}

func (s *CommandContext) Commentary() ICommentaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentaryContext)
}

func (s *CommandContext) Exit() IExitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExitContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParserParserRULE_command)
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_execute:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)

			var _x = p.Execute()

			localctx.(*CommandContext).c1 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC1().GetResult()

	case ParserParserRW_mkdisk:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)

			var _x = p.Mkdisk()

			localctx.(*CommandContext).c2 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC2().GetResult()

	case ParserParserRW_rmdisk:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)

			var _x = p.Rmdisk()

			localctx.(*CommandContext).c3 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC3().GetResult()

	case ParserParserRW_fdisk:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(114)

			var _x = p.Fdisk()

			localctx.(*CommandContext).c4 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC4().GetResult()

	case ParserParserRW_mount:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(117)

			var _x = p.Mount()

			localctx.(*CommandContext).c5 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC5().GetResult()

	case ParserParserRW_unmount:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(120)

			var _x = p.Unmount()

			localctx.(*CommandContext).c6 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC6().GetResult()

	case ParserParserRW_mkfs:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(123)

			var _x = p.Mkfs()

			localctx.(*CommandContext).c7 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC7().GetResult()

	case ParserParserRW_login:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(126)

			var _x = p.Login()

			localctx.(*CommandContext).c8 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC8().GetResult()

	case ParserParserRW_logout:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(129)

			var _x = p.Logout()

			localctx.(*CommandContext).c9 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC9().GetResult()

	case ParserParserRW_mkgrp:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(132)

			var _x = p.Mkgrp()

			localctx.(*CommandContext).c10 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC10().GetResult()

	case ParserParserRW_mkusr:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(135)

			var _x = p.Mkusr()

			localctx.(*CommandContext).c11 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC11().GetResult()

	case ParserParserRW_mkfile:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(138)

			var _x = p.Mkfile()

			localctx.(*CommandContext).c12 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC12().GetResult()

	case ParserParserRW_mkdir:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(141)

			var _x = p.Mkdir()

			localctx.(*CommandContext).c13 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC13().GetResult()

	case ParserParserRW_cat:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(144)

			var _x = p.Cat()

			localctx.(*CommandContext).c18 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC18().GetResult()

	case ParserParserRW_pause:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(147)

			var _x = p.Pause()

			localctx.(*CommandContext).c19 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC19().GetResult()

	case ParserParserRW_rep:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(150)

			var _x = p.Rep()

			localctx.(*CommandContext).c20 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC20().GetResult()

	case ParserParserCOMMENTARY:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(153)

			var _x = p.Commentary()

			localctx.(*CommandContext).c21 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC21().GetResult()

	case ParserParserRW_exit:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(156)
			p.Exit()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExecuteContext is an interface to support dynamic dispatch.
type IExecuteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE returns the e token.
	GetE() antlr.Token

	// GetP returns the p token.
	GetP() antlr.Token

	// SetE sets the e token.
	SetE(antlr.Token)

	// SetP sets the p token.
	SetP(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Execute

	// SetResult sets the result attribute.
	SetResult(*commands.Execute)

	// Getter signatures
	RW_path() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	RW_execute() antlr.TerminalNode
	TK_path() antlr.TerminalNode

	// IsExecuteContext differentiates from other interfaces.
	IsExecuteContext()
}

type ExecuteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Execute
	e      antlr.Token
	p      antlr.Token
}

func NewEmptyExecuteContext() *ExecuteContext {
	var p = new(ExecuteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_execute
	return p
}

func InitEmptyExecuteContext(p *ExecuteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_execute
}

func (*ExecuteContext) IsExecuteContext() {}

func NewExecuteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecuteContext {
	var p = new(ExecuteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_execute

	return p
}

func (s *ExecuteContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecuteContext) GetE() antlr.Token { return s.e }

func (s *ExecuteContext) GetP() antlr.Token { return s.p }

func (s *ExecuteContext) SetE(v antlr.Token) { s.e = v }

func (s *ExecuteContext) SetP(v antlr.Token) { s.p = v }

func (s *ExecuteContext) GetResult() *commands.Execute { return s.result }

func (s *ExecuteContext) SetResult(v *commands.Execute) { s.result = v }

func (s *ExecuteContext) RW_path() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_path, 0)
}

func (s *ExecuteContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *ExecuteContext) RW_execute() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_execute, 0)
}

func (s *ExecuteContext) TK_path() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_path, 0)
}

func (s *ExecuteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecuteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Execute() (localctx IExecuteContext) {
	localctx = NewExecuteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ParserParserRULE_execute)
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)

			var _m = p.Match(ParserParserRW_execute)

			localctx.(*ExecuteContext).e = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(ParserParserRW_path)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*ExecuteContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExecuteContext).result = commands.NewExecute((func() int {
			if localctx.(*ExecuteContext).GetE() == nil {
				return 0
			} else {
				return localctx.(*ExecuteContext).GetE().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExecuteContext).GetE() == nil {
				return 0
			} else {
				return localctx.(*ExecuteContext).GetE().GetColumn()
			}
		}()), map[string]string{"path": (func() string {
			if localctx.(*ExecuteContext).GetP() == nil {
				return ""
			} else {
				return localctx.(*ExecuteContext).GetP().GetText()
			}
		}())})

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)

			var _m = p.Match(ParserParserRW_execute)

			localctx.(*ExecuteContext).e = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExecuteContext).result = commands.NewExecute((func() int {
			if localctx.(*ExecuteContext).GetE() == nil {
				return 0
			} else {
				return localctx.(*ExecuteContext).GetE().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExecuteContext).GetE() == nil {
				return 0
			} else {
				return localctx.(*ExecuteContext).GetE().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkdiskContext is an interface to support dynamic dispatch.
type IMkdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m token.
	GetM() antlr.Token

	// SetM sets the m token.
	SetM(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IMkdiskparamsContext

	// SetP sets the p rule contexts.
	SetP(IMkdiskparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Mkdisk

	// SetResult sets the result attribute.
	SetResult(*commands.Mkdisk)

	// Getter signatures
	RW_mkdisk() antlr.TerminalNode
	Mkdiskparams() IMkdiskparamsContext

	// IsMkdiskContext differentiates from other interfaces.
	IsMkdiskContext()
}

type MkdiskContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mkdisk
	m      antlr.Token
	p      IMkdiskparamsContext
}

func NewEmptyMkdiskContext() *MkdiskContext {
	var p = new(MkdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdisk
	return p
}

func InitEmptyMkdiskContext(p *MkdiskContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdisk
}

func (*MkdiskContext) IsMkdiskContext() {}

func NewMkdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdiskContext {
	var p = new(MkdiskContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkdisk

	return p
}

func (s *MkdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdiskContext) GetM() antlr.Token { return s.m }

func (s *MkdiskContext) SetM(v antlr.Token) { s.m = v }

func (s *MkdiskContext) GetP() IMkdiskparamsContext { return s.p }

func (s *MkdiskContext) SetP(v IMkdiskparamsContext) { s.p = v }

func (s *MkdiskContext) GetResult() *commands.Mkdisk { return s.result }

func (s *MkdiskContext) SetResult(v *commands.Mkdisk) { s.result = v }

func (s *MkdiskContext) RW_mkdisk() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mkdisk, 0)
}

func (s *MkdiskContext) Mkdiskparams() IMkdiskparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskparamsContext)
}

func (s *MkdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkdisk() (localctx IMkdiskContext) {
	localctx = NewMkdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ParserParserRULE_mkdisk)
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)

			var _m = p.Match(ParserParserRW_mkdisk)

			localctx.(*MkdiskContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)

			var _x = p.mkdiskparams(0)

			localctx.(*MkdiskContext).p = _x
		}
		localctx.(*MkdiskContext).result = commands.NewMkdisk((func() int {
			if localctx.(*MkdiskContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdiskContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkdiskContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdiskContext).GetM().GetColumn()
			}
		}()), localctx.(*MkdiskContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)

			var _m = p.Match(ParserParserRW_mkdisk)

			localctx.(*MkdiskContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdiskContext).result = commands.NewMkdisk((func() int {
			if localctx.(*MkdiskContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdiskContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkdiskContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdiskContext).GetM().GetColumn()
			}
		}()), map[string]string{"fit": "FF", "unit": "M"})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkdiskparamsContext is an interface to support dynamic dispatch.
type IMkdiskparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IMkdiskparamsContext

	// GetP returns the p rule contexts.
	GetP() IMkdiskparamContext

	// SetL sets the l rule contexts.
	SetL(IMkdiskparamsContext)

	// SetP sets the p rule contexts.
	SetP(IMkdiskparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Mkdiskparam() IMkdiskparamContext
	Mkdiskparams() IMkdiskparamsContext

	// IsMkdiskparamsContext differentiates from other interfaces.
	IsMkdiskparamsContext()
}

type MkdiskparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IMkdiskparamsContext
	p      IMkdiskparamContext
}

func NewEmptyMkdiskparamsContext() *MkdiskparamsContext {
	var p = new(MkdiskparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdiskparams
	return p
}

func InitEmptyMkdiskparamsContext(p *MkdiskparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdiskparams
}

func (*MkdiskparamsContext) IsMkdiskparamsContext() {}

func NewMkdiskparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdiskparamsContext {
	var p = new(MkdiskparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkdiskparams

	return p
}

func (s *MkdiskparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdiskparamsContext) GetL() IMkdiskparamsContext { return s.l }

func (s *MkdiskparamsContext) GetP() IMkdiskparamContext { return s.p }

func (s *MkdiskparamsContext) SetL(v IMkdiskparamsContext) { s.l = v }

func (s *MkdiskparamsContext) SetP(v IMkdiskparamContext) { s.p = v }

func (s *MkdiskparamsContext) GetResult() map[string]string { return s.result }

func (s *MkdiskparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *MkdiskparamsContext) Mkdiskparam() IMkdiskparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskparamContext)
}

func (s *MkdiskparamsContext) Mkdiskparams() IMkdiskparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskparamsContext)
}

func (s *MkdiskparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdiskparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkdiskparams() (localctx IMkdiskparamsContext) {
	return p.mkdiskparams(0)
}

func (p *ParserParser) mkdiskparams(_p int) (localctx IMkdiskparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMkdiskparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMkdiskparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, ParserParserRULE_mkdiskparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)

		var _x = p.Mkdiskparam()

		localctx.(*MkdiskparamsContext).p = _x
	}
	localctx.(*MkdiskparamsContext).result = map[string]string{"fit": "FF", "unit": "M", localctx.(*MkdiskparamsContext).GetP().GetResult()[0]: localctx.(*MkdiskparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMkdiskparamsContext(p, _parentctx, _parentState)
			localctx.(*MkdiskparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_mkdiskparams)
			p.SetState(180)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(181)

				var _x = p.Mkdiskparam()

				localctx.(*MkdiskparamsContext).p = _x
			}
			localctx.(*MkdiskparamsContext).SetResult(localctx.(*MkdiskparamsContext).GetL().GetResult())
			localctx.(*MkdiskparamsContext).result[localctx.(*MkdiskparamsContext).GetP().GetResult()[0]] = localctx.(*MkdiskparamsContext).GetP().GetResult()[1]

		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkdiskparamContext is an interface to support dynamic dispatch.
type IMkdiskparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// GetV3 returns the v3 token.
	GetV3() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// SetV3 sets the v3 token.
	SetV3(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_size() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_number() antlr.TerminalNode
	RW_fit() antlr.TerminalNode
	TK_fit() antlr.TerminalNode
	RW_unit() antlr.TerminalNode
	TK_unit() antlr.TerminalNode

	// IsMkdiskparamContext differentiates from other interfaces.
	IsMkdiskparamContext()
}

type MkdiskparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
	v3     antlr.Token
}

func NewEmptyMkdiskparamContext() *MkdiskparamContext {
	var p = new(MkdiskparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdiskparam
	return p
}

func InitEmptyMkdiskparamContext(p *MkdiskparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdiskparam
}

func (*MkdiskparamContext) IsMkdiskparamContext() {}

func NewMkdiskparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdiskparamContext {
	var p = new(MkdiskparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkdiskparam

	return p
}

func (s *MkdiskparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdiskparamContext) GetV1() antlr.Token { return s.v1 }

func (s *MkdiskparamContext) GetV2() antlr.Token { return s.v2 }

func (s *MkdiskparamContext) GetV3() antlr.Token { return s.v3 }

func (s *MkdiskparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *MkdiskparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *MkdiskparamContext) SetV3(v antlr.Token) { s.v3 = v }

func (s *MkdiskparamContext) GetResult() []string { return s.result }

func (s *MkdiskparamContext) SetResult(v []string) { s.result = v }

func (s *MkdiskparamContext) RW_size() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_size, 0)
}

func (s *MkdiskparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MkdiskparamContext) TK_number() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_number, 0)
}

func (s *MkdiskparamContext) RW_fit() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_fit, 0)
}

func (s *MkdiskparamContext) TK_fit() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_fit, 0)
}

func (s *MkdiskparamContext) RW_unit() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_unit, 0)
}

func (s *MkdiskparamContext) TK_unit() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_unit, 0)
}

func (s *MkdiskparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdiskparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkdiskparam() (localctx IMkdiskparamContext) {
	localctx = NewMkdiskparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ParserParserRULE_mkdiskparam)
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_size:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Match(ParserParserRW_size)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*MkdiskparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdiskparamContext).result = []string{"size", (func() string {
			if localctx.(*MkdiskparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*MkdiskparamContext).GetV1().GetText()
			}
		}())}

	case ParserParserRW_fit:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.Match(ParserParserRW_fit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)

			var _m = p.Match(ParserParserTK_fit)

			localctx.(*MkdiskparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdiskparamContext).result = []string{"fit", (func() string {
			if localctx.(*MkdiskparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*MkdiskparamContext).GetV2().GetText()
			}
		}())}

	case ParserParserRW_unit:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(197)
			p.Match(ParserParserRW_unit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)

			var _m = p.Match(ParserParserTK_unit)

			localctx.(*MkdiskparamContext).v3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdiskparamContext).result = []string{"unit", (func() string {
			if localctx.(*MkdiskparamContext).GetV3() == nil {
				return ""
			} else {
				return localctx.(*MkdiskparamContext).GetV3().GetText()
			}
		}())}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRmdiskContext is an interface to support dynamic dispatch.
type IRmdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetR returns the r token.
	GetR() antlr.Token

	// GetP returns the p token.
	GetP() antlr.Token

	// SetR sets the r token.
	SetR(antlr.Token)

	// SetP sets the p token.
	SetP(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Rmdisk

	// SetResult sets the result attribute.
	SetResult(*commands.Rmdisk)

	// Getter signatures
	RW_driveletter() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	RW_rmdisk() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	TK_unit() antlr.TerminalNode
	TK_type() antlr.TerminalNode

	// IsRmdiskContext differentiates from other interfaces.
	IsRmdiskContext()
}

type RmdiskContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Rmdisk
	r      antlr.Token
	p      antlr.Token
}

func NewEmptyRmdiskContext() *RmdiskContext {
	var p = new(RmdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rmdisk
	return p
}

func InitEmptyRmdiskContext(p *RmdiskContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rmdisk
}

func (*RmdiskContext) IsRmdiskContext() {}

func NewRmdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RmdiskContext {
	var p = new(RmdiskContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_rmdisk

	return p
}

func (s *RmdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *RmdiskContext) GetR() antlr.Token { return s.r }

func (s *RmdiskContext) GetP() antlr.Token { return s.p }

func (s *RmdiskContext) SetR(v antlr.Token) { s.r = v }

func (s *RmdiskContext) SetP(v antlr.Token) { s.p = v }

func (s *RmdiskContext) GetResult() *commands.Rmdisk { return s.result }

func (s *RmdiskContext) SetResult(v *commands.Rmdisk) { s.result = v }

func (s *RmdiskContext) RW_driveletter() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_driveletter, 0)
}

func (s *RmdiskContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *RmdiskContext) RW_rmdisk() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_rmdisk, 0)
}

func (s *RmdiskContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *RmdiskContext) TK_unit() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_unit, 0)
}

func (s *RmdiskContext) TK_type() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_type, 0)
}

func (s *RmdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RmdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Rmdisk() (localctx IRmdiskContext) {
	localctx = NewRmdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ParserParserRULE_rmdisk)
	var _la int

	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)

			var _m = p.Match(ParserParserRW_rmdisk)

			localctx.(*RmdiskContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Match(ParserParserRW_driveletter)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*RmdiskContext).p = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&19) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*RmdiskContext).p = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		localctx.(*RmdiskContext).result = commands.NewRmdisk((func() int {
			if localctx.(*RmdiskContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RmdiskContext).GetR().GetLine()
			}
		}()), (func() int {
			if localctx.(*RmdiskContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RmdiskContext).GetR().GetColumn()
			}
		}()), map[string]string{"driveletter": (func() string {
			if localctx.(*RmdiskContext).GetP() == nil {
				return ""
			} else {
				return localctx.(*RmdiskContext).GetP().GetText()
			}
		}())})

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)

			var _m = p.Match(ParserParserRW_rmdisk)

			localctx.(*RmdiskContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RmdiskContext).result = commands.NewRmdisk((func() int {
			if localctx.(*RmdiskContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RmdiskContext).GetR().GetLine()
			}
		}()), (func() int {
			if localctx.(*RmdiskContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RmdiskContext).GetR().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFdiskContext is an interface to support dynamic dispatch.
type IFdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetF returns the f token.
	GetF() antlr.Token

	// SetF sets the f token.
	SetF(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IFdiskparamsContext

	// SetP sets the p rule contexts.
	SetP(IFdiskparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Fdisk

	// SetResult sets the result attribute.
	SetResult(*commands.Fdisk)

	// Getter signatures
	RW_fdisk() antlr.TerminalNode
	Fdiskparams() IFdiskparamsContext

	// IsFdiskContext differentiates from other interfaces.
	IsFdiskContext()
}

type FdiskContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Fdisk
	f      antlr.Token
	p      IFdiskparamsContext
}

func NewEmptyFdiskContext() *FdiskContext {
	var p = new(FdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdisk
	return p
}

func InitEmptyFdiskContext(p *FdiskContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdisk
}

func (*FdiskContext) IsFdiskContext() {}

func NewFdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FdiskContext {
	var p = new(FdiskContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_fdisk

	return p
}

func (s *FdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *FdiskContext) GetF() antlr.Token { return s.f }

func (s *FdiskContext) SetF(v antlr.Token) { s.f = v }

func (s *FdiskContext) GetP() IFdiskparamsContext { return s.p }

func (s *FdiskContext) SetP(v IFdiskparamsContext) { s.p = v }

func (s *FdiskContext) GetResult() *commands.Fdisk { return s.result }

func (s *FdiskContext) SetResult(v *commands.Fdisk) { s.result = v }

func (s *FdiskContext) RW_fdisk() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_fdisk, 0)
}

func (s *FdiskContext) Fdiskparams() IFdiskparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskparamsContext)
}

func (s *FdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Fdisk() (localctx IFdiskContext) {
	localctx = NewFdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ParserParserRULE_fdisk)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)

			var _m = p.Match(ParserParserRW_fdisk)

			localctx.(*FdiskContext).f = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)

			var _x = p.fdiskparams(0)

			localctx.(*FdiskContext).p = _x
		}
		localctx.(*FdiskContext).result = commands.NewFdisk((func() int {
			if localctx.(*FdiskContext).GetF() == nil {
				return 0
			} else {
				return localctx.(*FdiskContext).GetF().GetLine()
			}
		}()), (func() int {
			if localctx.(*FdiskContext).GetF() == nil {
				return 0
			} else {
				return localctx.(*FdiskContext).GetF().GetColumn()
			}
		}()), localctx.(*FdiskContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(216)

			var _m = p.Match(ParserParserRW_fdisk)

			localctx.(*FdiskContext).f = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskContext).result = commands.NewFdisk((func() int {
			if localctx.(*FdiskContext).GetF() == nil {
				return 0
			} else {
				return localctx.(*FdiskContext).GetF().GetLine()
			}
		}()), (func() int {
			if localctx.(*FdiskContext).GetF() == nil {
				return 0
			} else {
				return localctx.(*FdiskContext).GetF().GetColumn()
			}
		}()), map[string]string{"unit": "K", "fit": "WF", "type": "P"})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFdiskparamsContext is an interface to support dynamic dispatch.
type IFdiskparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IFdiskparamsContext

	// GetP returns the p rule contexts.
	GetP() IFdiskparamContext

	// SetL sets the l rule contexts.
	SetL(IFdiskparamsContext)

	// SetP sets the p rule contexts.
	SetP(IFdiskparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Fdiskparam() IFdiskparamContext
	Fdiskparams() IFdiskparamsContext

	// IsFdiskparamsContext differentiates from other interfaces.
	IsFdiskparamsContext()
}

type FdiskparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IFdiskparamsContext
	p      IFdiskparamContext
}

func NewEmptyFdiskparamsContext() *FdiskparamsContext {
	var p = new(FdiskparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdiskparams
	return p
}

func InitEmptyFdiskparamsContext(p *FdiskparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdiskparams
}

func (*FdiskparamsContext) IsFdiskparamsContext() {}

func NewFdiskparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FdiskparamsContext {
	var p = new(FdiskparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_fdiskparams

	return p
}

func (s *FdiskparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *FdiskparamsContext) GetL() IFdiskparamsContext { return s.l }

func (s *FdiskparamsContext) GetP() IFdiskparamContext { return s.p }

func (s *FdiskparamsContext) SetL(v IFdiskparamsContext) { s.l = v }

func (s *FdiskparamsContext) SetP(v IFdiskparamContext) { s.p = v }

func (s *FdiskparamsContext) GetResult() map[string]string { return s.result }

func (s *FdiskparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *FdiskparamsContext) Fdiskparam() IFdiskparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskparamContext)
}

func (s *FdiskparamsContext) Fdiskparams() IFdiskparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskparamsContext)
}

func (s *FdiskparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FdiskparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Fdiskparams() (localctx IFdiskparamsContext) {
	return p.fdiskparams(0)
}

func (p *ParserParser) fdiskparams(_p int) (localctx IFdiskparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewFdiskparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFdiskparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, ParserParserRULE_fdiskparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)

		var _x = p.Fdiskparam()

		localctx.(*FdiskparamsContext).p = _x
	}
	localctx.(*FdiskparamsContext).result = map[string]string{"unit": "K", "fit": "WF", "type": "P", localctx.(*FdiskparamsContext).GetP().GetResult()[0]: localctx.(*FdiskparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFdiskparamsContext(p, _parentctx, _parentState)
			localctx.(*FdiskparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_fdiskparams)
			p.SetState(224)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(225)

				var _x = p.Fdiskparam()

				localctx.(*FdiskparamsContext).p = _x
			}
			localctx.(*FdiskparamsContext).SetResult(localctx.(*FdiskparamsContext).GetL().GetResult())
			localctx.(*FdiskparamsContext).result[localctx.(*FdiskparamsContext).GetP().GetResult()[0]] = localctx.(*FdiskparamsContext).GetP().GetResult()[1]

		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFdiskparamContext is an interface to support dynamic dispatch.
type IFdiskparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// GetV3 returns the v3 token.
	GetV3() antlr.Token

	// GetV4 returns the v4 token.
	GetV4() antlr.Token

	// GetV5 returns the v5 token.
	GetV5() antlr.Token

	// GetV6 returns the v6 token.
	GetV6() antlr.Token

	// GetV7 returns the v7 token.
	GetV7() antlr.Token

	// GetV8 returns the v8 token.
	GetV8() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// SetV3 sets the v3 token.
	SetV3(antlr.Token)

	// SetV4 sets the v4 token.
	SetV4(antlr.Token)

	// SetV5 sets the v5 token.
	SetV5(antlr.Token)

	// SetV6 sets the v6 token.
	SetV6(antlr.Token)

	// SetV7 sets the v7 token.
	SetV7(antlr.Token)

	// SetV8 sets the v8 token.
	SetV8(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_size() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_number() antlr.TerminalNode
	RW_driveletter() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	TK_unit() antlr.TerminalNode
	TK_type() antlr.TerminalNode
	RW_name() antlr.TerminalNode
	RW_unit() antlr.TerminalNode
	RW_type() antlr.TerminalNode
	RW_fit() antlr.TerminalNode
	TK_fit() antlr.TerminalNode
	RW_delete() antlr.TerminalNode
	RW_full() antlr.TerminalNode
	RW_add() antlr.TerminalNode

	// IsFdiskparamContext differentiates from other interfaces.
	IsFdiskparamContext()
}

type FdiskparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
	v3     antlr.Token
	v4     antlr.Token
	v5     antlr.Token
	v6     antlr.Token
	v7     antlr.Token
	v8     antlr.Token
}

func NewEmptyFdiskparamContext() *FdiskparamContext {
	var p = new(FdiskparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdiskparam
	return p
}

func InitEmptyFdiskparamContext(p *FdiskparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdiskparam
}

func (*FdiskparamContext) IsFdiskparamContext() {}

func NewFdiskparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FdiskparamContext {
	var p = new(FdiskparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_fdiskparam

	return p
}

func (s *FdiskparamContext) GetParser() antlr.Parser { return s.parser }

func (s *FdiskparamContext) GetV1() antlr.Token { return s.v1 }

func (s *FdiskparamContext) GetV2() antlr.Token { return s.v2 }

func (s *FdiskparamContext) GetV3() antlr.Token { return s.v3 }

func (s *FdiskparamContext) GetV4() antlr.Token { return s.v4 }

func (s *FdiskparamContext) GetV5() antlr.Token { return s.v5 }

func (s *FdiskparamContext) GetV6() antlr.Token { return s.v6 }

func (s *FdiskparamContext) GetV7() antlr.Token { return s.v7 }

func (s *FdiskparamContext) GetV8() antlr.Token { return s.v8 }

func (s *FdiskparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *FdiskparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *FdiskparamContext) SetV3(v antlr.Token) { s.v3 = v }

func (s *FdiskparamContext) SetV4(v antlr.Token) { s.v4 = v }

func (s *FdiskparamContext) SetV5(v antlr.Token) { s.v5 = v }

func (s *FdiskparamContext) SetV6(v antlr.Token) { s.v6 = v }

func (s *FdiskparamContext) SetV7(v antlr.Token) { s.v7 = v }

func (s *FdiskparamContext) SetV8(v antlr.Token) { s.v8 = v }

func (s *FdiskparamContext) GetResult() []string { return s.result }

func (s *FdiskparamContext) SetResult(v []string) { s.result = v }

func (s *FdiskparamContext) RW_size() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_size, 0)
}

func (s *FdiskparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *FdiskparamContext) TK_number() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_number, 0)
}

func (s *FdiskparamContext) RW_driveletter() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_driveletter, 0)
}

func (s *FdiskparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *FdiskparamContext) TK_unit() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_unit, 0)
}

func (s *FdiskparamContext) TK_type() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_type, 0)
}

func (s *FdiskparamContext) RW_name() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_name, 0)
}

func (s *FdiskparamContext) RW_unit() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_unit, 0)
}

func (s *FdiskparamContext) RW_type() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_type, 0)
}

func (s *FdiskparamContext) RW_fit() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_fit, 0)
}

func (s *FdiskparamContext) TK_fit() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_fit, 0)
}

func (s *FdiskparamContext) RW_delete() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_delete, 0)
}

func (s *FdiskparamContext) RW_full() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_full, 0)
}

func (s *FdiskparamContext) RW_add() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_add, 0)
}

func (s *FdiskparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FdiskparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Fdiskparam() (localctx IFdiskparamContext) {
	localctx = NewFdiskparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ParserParserRULE_fdiskparam)
	var _la int

	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_size:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(233)
			p.Match(ParserParserRW_size)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*FdiskparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"size", (func() string {
			if localctx.(*FdiskparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV1().GetText()
			}
		}())}

	case ParserParserRW_driveletter:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(237)
			p.Match(ParserParserRW_driveletter)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*FdiskparamContext).v2 = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&19) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*FdiskparamContext).v2 = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		localctx.(*FdiskparamContext).result = []string{"driveletter", (func() string {
			if localctx.(*FdiskparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV2().GetText()
			}
		}())}

	case ParserParserRW_name:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(241)
			p.Match(ParserParserRW_name)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*FdiskparamContext).v3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"name", (func() string {
			if localctx.(*FdiskparamContext).GetV3() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV3().GetText()
			}
		}())}

	case ParserParserRW_unit:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(245)
			p.Match(ParserParserRW_unit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)

			var _m = p.Match(ParserParserTK_unit)

			localctx.(*FdiskparamContext).v4 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"unit", (func() string {
			if localctx.(*FdiskparamContext).GetV4() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV4().GetText()
			}
		}())}

	case ParserParserRW_type:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(249)
			p.Match(ParserParserRW_type)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)

			var _m = p.Match(ParserParserTK_type)

			localctx.(*FdiskparamContext).v5 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"type", (func() string {
			if localctx.(*FdiskparamContext).GetV5() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV5().GetText()
			}
		}())}

	case ParserParserRW_fit:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(253)
			p.Match(ParserParserRW_fit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)

			var _m = p.Match(ParserParserTK_fit)

			localctx.(*FdiskparamContext).v6 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"fit", (func() string {
			if localctx.(*FdiskparamContext).GetV6() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV6().GetText()
			}
		}())}

	case ParserParserRW_delete:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(257)
			p.Match(ParserParserRW_delete)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)

			var _m = p.Match(ParserParserRW_full)

			localctx.(*FdiskparamContext).v7 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"delete", (func() string {
			if localctx.(*FdiskparamContext).GetV7() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV7().GetText()
			}
		}())}

	case ParserParserRW_add:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(261)
			p.Match(ParserParserRW_add)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*FdiskparamContext).v8 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"add", (func() string {
			if localctx.(*FdiskparamContext).GetV8() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV8().GetText()
			}
		}())}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMountContext is an interface to support dynamic dispatch.
type IMountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m token.
	GetM() antlr.Token

	// SetM sets the m token.
	SetM(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IMountparamsContext

	// SetP sets the p rule contexts.
	SetP(IMountparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Mount

	// SetResult sets the result attribute.
	SetResult(*commands.Mount)

	// Getter signatures
	RW_mount() antlr.TerminalNode
	Mountparams() IMountparamsContext

	// IsMountContext differentiates from other interfaces.
	IsMountContext()
}

type MountContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mount
	m      antlr.Token
	p      IMountparamsContext
}

func NewEmptyMountContext() *MountContext {
	var p = new(MountContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mount
	return p
}

func InitEmptyMountContext(p *MountContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mount
}

func (*MountContext) IsMountContext() {}

func NewMountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MountContext {
	var p = new(MountContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mount

	return p
}

func (s *MountContext) GetParser() antlr.Parser { return s.parser }

func (s *MountContext) GetM() antlr.Token { return s.m }

func (s *MountContext) SetM(v antlr.Token) { s.m = v }

func (s *MountContext) GetP() IMountparamsContext { return s.p }

func (s *MountContext) SetP(v IMountparamsContext) { s.p = v }

func (s *MountContext) GetResult() *commands.Mount { return s.result }

func (s *MountContext) SetResult(v *commands.Mount) { s.result = v }

func (s *MountContext) RW_mount() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mount, 0)
}

func (s *MountContext) Mountparams() IMountparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountparamsContext)
}

func (s *MountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mount() (localctx IMountContext) {
	localctx = NewMountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ParserParserRULE_mount)
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(267)

			var _m = p.Match(ParserParserRW_mount)

			localctx.(*MountContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)

			var _x = p.mountparams(0)

			localctx.(*MountContext).p = _x
		}
		localctx.(*MountContext).result = commands.NewMount((func() int {
			if localctx.(*MountContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MountContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MountContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MountContext).GetM().GetColumn()
			}
		}()), localctx.(*MountContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(271)

			var _m = p.Match(ParserParserRW_mount)

			localctx.(*MountContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MountContext).result = commands.NewMount((func() int {
			if localctx.(*MountContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MountContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MountContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MountContext).GetM().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMountparamsContext is an interface to support dynamic dispatch.
type IMountparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IMountparamsContext

	// GetP returns the p rule contexts.
	GetP() IMountparamContext

	// SetL sets the l rule contexts.
	SetL(IMountparamsContext)

	// SetP sets the p rule contexts.
	SetP(IMountparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Mountparam() IMountparamContext
	Mountparams() IMountparamsContext

	// IsMountparamsContext differentiates from other interfaces.
	IsMountparamsContext()
}

type MountparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IMountparamsContext
	p      IMountparamContext
}

func NewEmptyMountparamsContext() *MountparamsContext {
	var p = new(MountparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mountparams
	return p
}

func InitEmptyMountparamsContext(p *MountparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mountparams
}

func (*MountparamsContext) IsMountparamsContext() {}

func NewMountparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MountparamsContext {
	var p = new(MountparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mountparams

	return p
}

func (s *MountparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *MountparamsContext) GetL() IMountparamsContext { return s.l }

func (s *MountparamsContext) GetP() IMountparamContext { return s.p }

func (s *MountparamsContext) SetL(v IMountparamsContext) { s.l = v }

func (s *MountparamsContext) SetP(v IMountparamContext) { s.p = v }

func (s *MountparamsContext) GetResult() map[string]string { return s.result }

func (s *MountparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *MountparamsContext) Mountparam() IMountparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountparamContext)
}

func (s *MountparamsContext) Mountparams() IMountparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountparamsContext)
}

func (s *MountparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mountparams() (localctx IMountparamsContext) {
	return p.mountparams(0)
}

func (p *ParserParser) mountparams(_p int) (localctx IMountparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMountparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMountparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, ParserParserRULE_mountparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)

		var _x = p.Mountparam()

		localctx.(*MountparamsContext).p = _x
	}
	localctx.(*MountparamsContext).result = map[string]string{localctx.(*MountparamsContext).GetP().GetResult()[0]: localctx.(*MountparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMountparamsContext(p, _parentctx, _parentState)
			localctx.(*MountparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_mountparams)
			p.SetState(279)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(280)

				var _x = p.Mountparam()

				localctx.(*MountparamsContext).p = _x
			}
			localctx.(*MountparamsContext).SetResult(localctx.(*MountparamsContext).GetL().GetResult())
			localctx.(*MountparamsContext).result[localctx.(*MountparamsContext).GetP().GetResult()[0]] = localctx.(*MountparamsContext).GetP().GetResult()[1]

		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMountparamContext is an interface to support dynamic dispatch.
type IMountparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_driveletter() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	TK_unit() antlr.TerminalNode
	TK_type() antlr.TerminalNode
	RW_name() antlr.TerminalNode

	// IsMountparamContext differentiates from other interfaces.
	IsMountparamContext()
}

type MountparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
}

func NewEmptyMountparamContext() *MountparamContext {
	var p = new(MountparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mountparam
	return p
}

func InitEmptyMountparamContext(p *MountparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mountparam
}

func (*MountparamContext) IsMountparamContext() {}

func NewMountparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MountparamContext {
	var p = new(MountparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mountparam

	return p
}

func (s *MountparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MountparamContext) GetV1() antlr.Token { return s.v1 }

func (s *MountparamContext) GetV2() antlr.Token { return s.v2 }

func (s *MountparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *MountparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *MountparamContext) GetResult() []string { return s.result }

func (s *MountparamContext) SetResult(v []string) { s.result = v }

func (s *MountparamContext) RW_driveletter() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_driveletter, 0)
}

func (s *MountparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MountparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *MountparamContext) TK_unit() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_unit, 0)
}

func (s *MountparamContext) TK_type() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_type, 0)
}

func (s *MountparamContext) RW_name() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_name, 0)
}

func (s *MountparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mountparam() (localctx IMountparamContext) {
	localctx = NewMountparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ParserParserRULE_mountparam)
	var _la int

	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_driveletter:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(288)
			p.Match(ParserParserRW_driveletter)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MountparamContext).v1 = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&19) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*MountparamContext).v1 = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		localctx.(*MountparamContext).result = []string{"driveletter", (func() string {
			if localctx.(*MountparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*MountparamContext).GetV1().GetText()
			}
		}())}

	case ParserParserRW_name:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)
			p.Match(ParserParserRW_name)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MountparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MountparamContext).result = []string{"name", (func() string {
			if localctx.(*MountparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*MountparamContext).GetV2().GetText()
			}
		}())}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnmountContext is an interface to support dynamic dispatch.
type IUnmountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetU returns the u token.
	GetU() antlr.Token

	// GetP returns the p token.
	GetP() antlr.Token

	// SetU sets the u token.
	SetU(antlr.Token)

	// SetP sets the p token.
	SetP(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Unmount

	// SetResult sets the result attribute.
	SetResult(*commands.Unmount)

	// Getter signatures
	RW_id() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	RW_unmount() antlr.TerminalNode
	TK_id() antlr.TerminalNode

	// IsUnmountContext differentiates from other interfaces.
	IsUnmountContext()
}

type UnmountContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Unmount
	u      antlr.Token
	p      antlr.Token
}

func NewEmptyUnmountContext() *UnmountContext {
	var p = new(UnmountContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_unmount
	return p
}

func InitEmptyUnmountContext(p *UnmountContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_unmount
}

func (*UnmountContext) IsUnmountContext() {}

func NewUnmountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnmountContext {
	var p = new(UnmountContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_unmount

	return p
}

func (s *UnmountContext) GetParser() antlr.Parser { return s.parser }

func (s *UnmountContext) GetU() antlr.Token { return s.u }

func (s *UnmountContext) GetP() antlr.Token { return s.p }

func (s *UnmountContext) SetU(v antlr.Token) { s.u = v }

func (s *UnmountContext) SetP(v antlr.Token) { s.p = v }

func (s *UnmountContext) GetResult() *commands.Unmount { return s.result }

func (s *UnmountContext) SetResult(v *commands.Unmount) { s.result = v }

func (s *UnmountContext) RW_id() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_id, 0)
}

func (s *UnmountContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *UnmountContext) RW_unmount() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_unmount, 0)
}

func (s *UnmountContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *UnmountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnmountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Unmount() (localctx IUnmountContext) {
	localctx = NewUnmountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ParserParserRULE_unmount)
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(298)

			var _m = p.Match(ParserParserRW_unmount)

			localctx.(*UnmountContext).u = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Match(ParserParserRW_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*UnmountContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*UnmountContext).result = commands.NewUnmount((func() int {
			if localctx.(*UnmountContext).GetU() == nil {
				return 0
			} else {
				return localctx.(*UnmountContext).GetU().GetLine()
			}
		}()), (func() int {
			if localctx.(*UnmountContext).GetU() == nil {
				return 0
			} else {
				return localctx.(*UnmountContext).GetU().GetColumn()
			}
		}()), map[string]string{"id": (func() string {
			if localctx.(*UnmountContext).GetP() == nil {
				return ""
			} else {
				return localctx.(*UnmountContext).GetP().GetText()
			}
		}())})

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(303)

			var _m = p.Match(ParserParserRW_unmount)

			localctx.(*UnmountContext).u = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*UnmountContext).result = commands.NewUnmount((func() int {
			if localctx.(*UnmountContext).GetU() == nil {
				return 0
			} else {
				return localctx.(*UnmountContext).GetU().GetLine()
			}
		}()), (func() int {
			if localctx.(*UnmountContext).GetU() == nil {
				return 0
			} else {
				return localctx.(*UnmountContext).GetU().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkfsContext is an interface to support dynamic dispatch.
type IMkfsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m token.
	GetM() antlr.Token

	// SetM sets the m token.
	SetM(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IMkfsparamsContext

	// SetP sets the p rule contexts.
	SetP(IMkfsparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Mkfs

	// SetResult sets the result attribute.
	SetResult(*commands.Mkfs)

	// Getter signatures
	RW_mkfs() antlr.TerminalNode
	Mkfsparams() IMkfsparamsContext

	// IsMkfsContext differentiates from other interfaces.
	IsMkfsContext()
}

type MkfsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mkfs
	m      antlr.Token
	p      IMkfsparamsContext
}

func NewEmptyMkfsContext() *MkfsContext {
	var p = new(MkfsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfs
	return p
}

func InitEmptyMkfsContext(p *MkfsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfs
}

func (*MkfsContext) IsMkfsContext() {}

func NewMkfsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfsContext {
	var p = new(MkfsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkfs

	return p
}

func (s *MkfsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfsContext) GetM() antlr.Token { return s.m }

func (s *MkfsContext) SetM(v antlr.Token) { s.m = v }

func (s *MkfsContext) GetP() IMkfsparamsContext { return s.p }

func (s *MkfsContext) SetP(v IMkfsparamsContext) { s.p = v }

func (s *MkfsContext) GetResult() *commands.Mkfs { return s.result }

func (s *MkfsContext) SetResult(v *commands.Mkfs) { s.result = v }

func (s *MkfsContext) RW_mkfs() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mkfs, 0)
}

func (s *MkfsContext) Mkfsparams() IMkfsparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfsparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfsparamsContext)
}

func (s *MkfsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkfs() (localctx IMkfsContext) {
	localctx = NewMkfsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ParserParserRULE_mkfs)
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(307)

			var _m = p.Match(ParserParserRW_mkfs)

			localctx.(*MkfsContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)

			var _x = p.mkfsparams(0)

			localctx.(*MkfsContext).p = _x
		}
		localctx.(*MkfsContext).result = commands.NewMkfs((func() int {
			if localctx.(*MkfsContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfsContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkfsContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfsContext).GetM().GetColumn()
			}
		}()), localctx.(*MkfsContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(311)

			var _m = p.Match(ParserParserRW_mkfs)

			localctx.(*MkfsContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfsContext).result = commands.NewMkfs((func() int {
			if localctx.(*MkfsContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfsContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkfsContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfsContext).GetM().GetColumn()
			}
		}()), map[string]string{"fs": "2fs"})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkfsparamsContext is an interface to support dynamic dispatch.
type IMkfsparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IMkfsparamsContext

	// GetP returns the p rule contexts.
	GetP() IMkfsparamContext

	// SetL sets the l rule contexts.
	SetL(IMkfsparamsContext)

	// SetP sets the p rule contexts.
	SetP(IMkfsparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Mkfsparam() IMkfsparamContext
	Mkfsparams() IMkfsparamsContext

	// IsMkfsparamsContext differentiates from other interfaces.
	IsMkfsparamsContext()
}

type MkfsparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IMkfsparamsContext
	p      IMkfsparamContext
}

func NewEmptyMkfsparamsContext() *MkfsparamsContext {
	var p = new(MkfsparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfsparams
	return p
}

func InitEmptyMkfsparamsContext(p *MkfsparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfsparams
}

func (*MkfsparamsContext) IsMkfsparamsContext() {}

func NewMkfsparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfsparamsContext {
	var p = new(MkfsparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkfsparams

	return p
}

func (s *MkfsparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfsparamsContext) GetL() IMkfsparamsContext { return s.l }

func (s *MkfsparamsContext) GetP() IMkfsparamContext { return s.p }

func (s *MkfsparamsContext) SetL(v IMkfsparamsContext) { s.l = v }

func (s *MkfsparamsContext) SetP(v IMkfsparamContext) { s.p = v }

func (s *MkfsparamsContext) GetResult() map[string]string { return s.result }

func (s *MkfsparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *MkfsparamsContext) Mkfsparam() IMkfsparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfsparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfsparamContext)
}

func (s *MkfsparamsContext) Mkfsparams() IMkfsparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfsparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfsparamsContext)
}

func (s *MkfsparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfsparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkfsparams() (localctx IMkfsparamsContext) {
	return p.mkfsparams(0)
}

func (p *ParserParser) mkfsparams(_p int) (localctx IMkfsparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMkfsparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMkfsparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, ParserParserRULE_mkfsparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)

		var _x = p.Mkfsparam()

		localctx.(*MkfsparamsContext).p = _x
	}
	localctx.(*MkfsparamsContext).result = map[string]string{"fs": "2fs", localctx.(*MkfsparamsContext).GetP().GetResult()[0]: localctx.(*MkfsparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMkfsparamsContext(p, _parentctx, _parentState)
			localctx.(*MkfsparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_mkfsparams)
			p.SetState(319)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(320)

				var _x = p.Mkfsparam()

				localctx.(*MkfsparamsContext).p = _x
			}
			localctx.(*MkfsparamsContext).SetResult(localctx.(*MkfsparamsContext).GetL().GetResult())
			localctx.(*MkfsparamsContext).result[localctx.(*MkfsparamsContext).GetP().GetResult()[0]] = localctx.(*MkfsparamsContext).GetP().GetResult()[1]

		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkfsparamContext is an interface to support dynamic dispatch.
type IMkfsparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// GetV3 returns the v3 token.
	GetV3() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// SetV3 sets the v3 token.
	SetV3(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_id() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	RW_type() antlr.TerminalNode
	RW_full() antlr.TerminalNode
	RW_fs() antlr.TerminalNode
	TK_fs() antlr.TerminalNode

	// IsMkfsparamContext differentiates from other interfaces.
	IsMkfsparamContext()
}

type MkfsparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
	v3     antlr.Token
}

func NewEmptyMkfsparamContext() *MkfsparamContext {
	var p = new(MkfsparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfsparam
	return p
}

func InitEmptyMkfsparamContext(p *MkfsparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfsparam
}

func (*MkfsparamContext) IsMkfsparamContext() {}

func NewMkfsparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfsparamContext {
	var p = new(MkfsparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkfsparam

	return p
}

func (s *MkfsparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfsparamContext) GetV1() antlr.Token { return s.v1 }

func (s *MkfsparamContext) GetV2() antlr.Token { return s.v2 }

func (s *MkfsparamContext) GetV3() antlr.Token { return s.v3 }

func (s *MkfsparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *MkfsparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *MkfsparamContext) SetV3(v antlr.Token) { s.v3 = v }

func (s *MkfsparamContext) GetResult() []string { return s.result }

func (s *MkfsparamContext) SetResult(v []string) { s.result = v }

func (s *MkfsparamContext) RW_id() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_id, 0)
}

func (s *MkfsparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MkfsparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *MkfsparamContext) RW_type() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_type, 0)
}

func (s *MkfsparamContext) RW_full() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_full, 0)
}

func (s *MkfsparamContext) RW_fs() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_fs, 0)
}

func (s *MkfsparamContext) TK_fs() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_fs, 0)
}

func (s *MkfsparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfsparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkfsparam() (localctx IMkfsparamContext) {
	localctx = NewMkfsparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ParserParserRULE_mkfsparam)
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_id:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(328)
			p.Match(ParserParserRW_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MkfsparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfsparamContext).result = []string{"id", (func() string {
			if localctx.(*MkfsparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*MkfsparamContext).GetV1().GetText()
			}
		}())}

	case ParserParserRW_type:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(332)
			p.Match(ParserParserRW_type)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)

			var _m = p.Match(ParserParserRW_full)

			localctx.(*MkfsparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfsparamContext).result = []string{"type", (func() string {
			if localctx.(*MkfsparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*MkfsparamContext).GetV2().GetText()
			}
		}())}

	case ParserParserRW_fs:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(336)
			p.Match(ParserParserRW_fs)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)

			var _m = p.Match(ParserParserTK_fs)

			localctx.(*MkfsparamContext).v3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfsparamContext).result = []string{"fs", (func() string {
			if localctx.(*MkfsparamContext).GetV3() == nil {
				return ""
			} else {
				return localctx.(*MkfsparamContext).GetV3().GetText()
			}
		}())}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoginContext is an interface to support dynamic dispatch.
type ILoginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l token.
	GetL() antlr.Token

	// SetL sets the l token.
	SetL(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() ILoginparamsContext

	// SetP sets the p rule contexts.
	SetP(ILoginparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Login

	// SetResult sets the result attribute.
	SetResult(*commands.Login)

	// Getter signatures
	RW_login() antlr.TerminalNode
	Loginparams() ILoginparamsContext

	// IsLoginContext differentiates from other interfaces.
	IsLoginContext()
}

type LoginContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Login
	l      antlr.Token
	p      ILoginparamsContext
}

func NewEmptyLoginContext() *LoginContext {
	var p = new(LoginContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_login
	return p
}

func InitEmptyLoginContext(p *LoginContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_login
}

func (*LoginContext) IsLoginContext() {}

func NewLoginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoginContext {
	var p = new(LoginContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_login

	return p
}

func (s *LoginContext) GetParser() antlr.Parser { return s.parser }

func (s *LoginContext) GetL() antlr.Token { return s.l }

func (s *LoginContext) SetL(v antlr.Token) { s.l = v }

func (s *LoginContext) GetP() ILoginparamsContext { return s.p }

func (s *LoginContext) SetP(v ILoginparamsContext) { s.p = v }

func (s *LoginContext) GetResult() *commands.Login { return s.result }

func (s *LoginContext) SetResult(v *commands.Login) { s.result = v }

func (s *LoginContext) RW_login() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_login, 0)
}

func (s *LoginContext) Loginparams() ILoginparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoginparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoginparamsContext)
}

func (s *LoginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Login() (localctx ILoginContext) {
	localctx = NewLoginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ParserParserRULE_login)
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(342)

			var _m = p.Match(ParserParserRW_login)

			localctx.(*LoginContext).l = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)

			var _x = p.loginparams(0)

			localctx.(*LoginContext).p = _x
		}
		localctx.(*LoginContext).result = commands.NewLogin((func() int {
			if localctx.(*LoginContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*LoginContext).GetL().GetLine()
			}
		}()), (func() int {
			if localctx.(*LoginContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*LoginContext).GetL().GetColumn()
			}
		}()), localctx.(*LoginContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)

			var _m = p.Match(ParserParserRW_login)

			localctx.(*LoginContext).l = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginContext).result = commands.NewLogin((func() int {
			if localctx.(*LoginContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*LoginContext).GetL().GetLine()
			}
		}()), (func() int {
			if localctx.(*LoginContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*LoginContext).GetL().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoginparamsContext is an interface to support dynamic dispatch.
type ILoginparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() ILoginparamsContext

	// GetP returns the p rule contexts.
	GetP() ILoginparamContext

	// SetL sets the l rule contexts.
	SetL(ILoginparamsContext)

	// SetP sets the p rule contexts.
	SetP(ILoginparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Loginparam() ILoginparamContext
	Loginparams() ILoginparamsContext

	// IsLoginparamsContext differentiates from other interfaces.
	IsLoginparamsContext()
}

type LoginparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      ILoginparamsContext
	p      ILoginparamContext
}

func NewEmptyLoginparamsContext() *LoginparamsContext {
	var p = new(LoginparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loginparams
	return p
}

func InitEmptyLoginparamsContext(p *LoginparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loginparams
}

func (*LoginparamsContext) IsLoginparamsContext() {}

func NewLoginparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoginparamsContext {
	var p = new(LoginparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_loginparams

	return p
}

func (s *LoginparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *LoginparamsContext) GetL() ILoginparamsContext { return s.l }

func (s *LoginparamsContext) GetP() ILoginparamContext { return s.p }

func (s *LoginparamsContext) SetL(v ILoginparamsContext) { s.l = v }

func (s *LoginparamsContext) SetP(v ILoginparamContext) { s.p = v }

func (s *LoginparamsContext) GetResult() map[string]string { return s.result }

func (s *LoginparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *LoginparamsContext) Loginparam() ILoginparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoginparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoginparamContext)
}

func (s *LoginparamsContext) Loginparams() ILoginparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoginparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoginparamsContext)
}

func (s *LoginparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoginparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Loginparams() (localctx ILoginparamsContext) {
	return p.loginparams(0)
}

func (p *ParserParser) loginparams(_p int) (localctx ILoginparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewLoginparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILoginparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, ParserParserRULE_loginparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)

		var _x = p.Loginparam()

		localctx.(*LoginparamsContext).p = _x
	}
	localctx.(*LoginparamsContext).result = map[string]string{localctx.(*LoginparamsContext).GetP().GetResult()[0]: localctx.(*LoginparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLoginparamsContext(p, _parentctx, _parentState)
			localctx.(*LoginparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_loginparams)
			p.SetState(354)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(355)

				var _x = p.Loginparam()

				localctx.(*LoginparamsContext).p = _x
			}
			localctx.(*LoginparamsContext).SetResult(localctx.(*LoginparamsContext).GetL().GetResult())
			localctx.(*LoginparamsContext).result[localctx.(*LoginparamsContext).GetP().GetResult()[0]] = localctx.(*LoginparamsContext).GetP().GetResult()[1]

		}
		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoginparamContext is an interface to support dynamic dispatch.
type ILoginparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// GetV3 returns the v3 token.
	GetV3() antlr.Token

	// GetV4 returns the v4 token.
	GetV4() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// SetV3 sets the v3 token.
	SetV3(antlr.Token)

	// SetV4 sets the v4 token.
	SetV4(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_user() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	RW_pass() antlr.TerminalNode
	TK_number() antlr.TerminalNode
	RW_id() antlr.TerminalNode

	// IsLoginparamContext differentiates from other interfaces.
	IsLoginparamContext()
}

type LoginparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
	v3     antlr.Token
	v4     antlr.Token
}

func NewEmptyLoginparamContext() *LoginparamContext {
	var p = new(LoginparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loginparam
	return p
}

func InitEmptyLoginparamContext(p *LoginparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loginparam
}

func (*LoginparamContext) IsLoginparamContext() {}

func NewLoginparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoginparamContext {
	var p = new(LoginparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_loginparam

	return p
}

func (s *LoginparamContext) GetParser() antlr.Parser { return s.parser }

func (s *LoginparamContext) GetV1() antlr.Token { return s.v1 }

func (s *LoginparamContext) GetV2() antlr.Token { return s.v2 }

func (s *LoginparamContext) GetV3() antlr.Token { return s.v3 }

func (s *LoginparamContext) GetV4() antlr.Token { return s.v4 }

func (s *LoginparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *LoginparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *LoginparamContext) SetV3(v antlr.Token) { s.v3 = v }

func (s *LoginparamContext) SetV4(v antlr.Token) { s.v4 = v }

func (s *LoginparamContext) GetResult() []string { return s.result }

func (s *LoginparamContext) SetResult(v []string) { s.result = v }

func (s *LoginparamContext) RW_user() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_user, 0)
}

func (s *LoginparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *LoginparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *LoginparamContext) RW_pass() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_pass, 0)
}

func (s *LoginparamContext) TK_number() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_number, 0)
}

func (s *LoginparamContext) RW_id() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_id, 0)
}

func (s *LoginparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoginparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Loginparam() (localctx ILoginparamContext) {
	localctx = NewLoginparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ParserParserRULE_loginparam)
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(363)
			p.Match(ParserParserRW_user)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*LoginparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginparamContext).result = []string{"user", (func() string {
			if localctx.(*LoginparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetV1().GetText()
			}
		}())}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(367)
			p.Match(ParserParserRW_pass)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*LoginparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginparamContext).result = []string{"pass", (func() string {
			if localctx.(*LoginparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetV2().GetText()
			}
		}())}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(371)
			p.Match(ParserParserRW_pass)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*LoginparamContext).v3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginparamContext).result = []string{"pass", (func() string {
			if localctx.(*LoginparamContext).GetV3() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetV3().GetText()
			}
		}())}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(375)
			p.Match(ParserParserRW_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*LoginparamContext).v4 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginparamContext).result = []string{"id", (func() string {
			if localctx.(*LoginparamContext).GetV4() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetV4().GetText()
			}
		}())}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogoutContext is an interface to support dynamic dispatch.
type ILogoutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP returns the p token.
	GetP() antlr.Token

	// SetP sets the p token.
	SetP(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Logout

	// SetResult sets the result attribute.
	SetResult(*commands.Logout)

	// Getter signatures
	RW_logout() antlr.TerminalNode

	// IsLogoutContext differentiates from other interfaces.
	IsLogoutContext()
}

type LogoutContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Logout
	p      antlr.Token
}

func NewEmptyLogoutContext() *LogoutContext {
	var p = new(LogoutContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_logout
	return p
}

func InitEmptyLogoutContext(p *LogoutContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_logout
}

func (*LogoutContext) IsLogoutContext() {}

func NewLogoutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogoutContext {
	var p = new(LogoutContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_logout

	return p
}

func (s *LogoutContext) GetParser() antlr.Parser { return s.parser }

func (s *LogoutContext) GetP() antlr.Token { return s.p }

func (s *LogoutContext) SetP(v antlr.Token) { s.p = v }

func (s *LogoutContext) GetResult() *commands.Logout { return s.result }

func (s *LogoutContext) SetResult(v *commands.Logout) { s.result = v }

func (s *LogoutContext) RW_logout() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_logout, 0)
}

func (s *LogoutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogoutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Logout() (localctx ILogoutContext) {
	localctx = NewLogoutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ParserParserRULE_logout)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)

		var _m = p.Match(ParserParserRW_logout)

		localctx.(*LogoutContext).p = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*LogoutContext).result = commands.NewLogout((func() int {
		if localctx.(*LogoutContext).GetP() == nil {
			return 0
		} else {
			return localctx.(*LogoutContext).GetP().GetLine()
		}
	}()), (func() int {
		if localctx.(*LogoutContext).GetP() == nil {
			return 0
		} else {
			return localctx.(*LogoutContext).GetP().GetColumn()
		}
	}()))

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPauseContext is an interface to support dynamic dispatch.
type IPauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP returns the p token.
	GetP() antlr.Token

	// SetP sets the p token.
	SetP(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Pause

	// SetResult sets the result attribute.
	SetResult(*commands.Pause)

	// Getter signatures
	RW_pause() antlr.TerminalNode

	// IsPauseContext differentiates from other interfaces.
	IsPauseContext()
}

type PauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Pause
	p      antlr.Token
}

func NewEmptyPauseContext() *PauseContext {
	var p = new(PauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_pause
	return p
}

func InitEmptyPauseContext(p *PauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_pause
}

func (*PauseContext) IsPauseContext() {}

func NewPauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PauseContext {
	var p = new(PauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_pause

	return p
}

func (s *PauseContext) GetParser() antlr.Parser { return s.parser }

func (s *PauseContext) GetP() antlr.Token { return s.p }

func (s *PauseContext) SetP(v antlr.Token) { s.p = v }

func (s *PauseContext) GetResult() *commands.Pause { return s.result }

func (s *PauseContext) SetResult(v *commands.Pause) { s.result = v }

func (s *PauseContext) RW_pause() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_pause, 0)
}

func (s *PauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Pause() (localctx IPauseContext) {
	localctx = NewPauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ParserParserRULE_pause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(384)

		var _m = p.Match(ParserParserRW_pause)

		localctx.(*PauseContext).p = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*PauseContext).result = commands.NewPause((func() int {
		if localctx.(*PauseContext).GetP() == nil {
			return 0
		} else {
			return localctx.(*PauseContext).GetP().GetLine()
		}
	}()), (func() int {
		if localctx.(*PauseContext).GetP() == nil {
			return 0
		} else {
			return localctx.(*PauseContext).GetP().GetColumn()
		}
	}()))

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkgrpContext is an interface to support dynamic dispatch.
type IMkgrpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m token.
	GetM() antlr.Token

	// GetV returns the v token.
	GetV() antlr.Token

	// SetM sets the m token.
	SetM(antlr.Token)

	// SetV sets the v token.
	SetV(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Mkgrp

	// SetResult sets the result attribute.
	SetResult(*commands.Mkgrp)

	// Getter signatures
	RW_name() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	RW_mkgrp() antlr.TerminalNode
	TK_id() antlr.TerminalNode

	// IsMkgrpContext differentiates from other interfaces.
	IsMkgrpContext()
}

type MkgrpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mkgrp
	m      antlr.Token
	v      antlr.Token
}

func NewEmptyMkgrpContext() *MkgrpContext {
	var p = new(MkgrpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkgrp
	return p
}

func InitEmptyMkgrpContext(p *MkgrpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkgrp
}

func (*MkgrpContext) IsMkgrpContext() {}

func NewMkgrpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkgrpContext {
	var p = new(MkgrpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkgrp

	return p
}

func (s *MkgrpContext) GetParser() antlr.Parser { return s.parser }

func (s *MkgrpContext) GetM() antlr.Token { return s.m }

func (s *MkgrpContext) GetV() antlr.Token { return s.v }

func (s *MkgrpContext) SetM(v antlr.Token) { s.m = v }

func (s *MkgrpContext) SetV(v antlr.Token) { s.v = v }

func (s *MkgrpContext) GetResult() *commands.Mkgrp { return s.result }

func (s *MkgrpContext) SetResult(v *commands.Mkgrp) { s.result = v }

func (s *MkgrpContext) RW_name() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_name, 0)
}

func (s *MkgrpContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MkgrpContext) RW_mkgrp() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mkgrp, 0)
}

func (s *MkgrpContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *MkgrpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkgrpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkgrp() (localctx IMkgrpContext) {
	localctx = NewMkgrpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ParserParserRULE_mkgrp)
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(387)

			var _m = p.Match(ParserParserRW_mkgrp)

			localctx.(*MkgrpContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			p.Match(ParserParserRW_name)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(389)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MkgrpContext).v = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkgrpContext).result = commands.NewMkgrp((func() int {
			if localctx.(*MkgrpContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkgrpContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkgrpContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkgrpContext).GetM().GetColumn()
			}
		}()), map[string]string{"name": (func() string {
			if localctx.(*MkgrpContext).GetV() == nil {
				return ""
			} else {
				return localctx.(*MkgrpContext).GetV().GetText()
			}
		}())})

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(392)

			var _m = p.Match(ParserParserRW_mkgrp)

			localctx.(*MkgrpContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkgrpContext).result = commands.NewMkgrp((func() int {
			if localctx.(*MkgrpContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkgrpContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkgrpContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkgrpContext).GetM().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkusrContext is an interface to support dynamic dispatch.
type IMkusrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m token.
	GetM() antlr.Token

	// SetM sets the m token.
	SetM(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IMkuserparamsContext

	// SetP sets the p rule contexts.
	SetP(IMkuserparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Mkusr

	// SetResult sets the result attribute.
	SetResult(*commands.Mkusr)

	// Getter signatures
	RW_mkusr() antlr.TerminalNode
	Mkuserparams() IMkuserparamsContext

	// IsMkusrContext differentiates from other interfaces.
	IsMkusrContext()
}

type MkusrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mkusr
	m      antlr.Token
	p      IMkuserparamsContext
}

func NewEmptyMkusrContext() *MkusrContext {
	var p = new(MkusrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkusr
	return p
}

func InitEmptyMkusrContext(p *MkusrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkusr
}

func (*MkusrContext) IsMkusrContext() {}

func NewMkusrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkusrContext {
	var p = new(MkusrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkusr

	return p
}

func (s *MkusrContext) GetParser() antlr.Parser { return s.parser }

func (s *MkusrContext) GetM() antlr.Token { return s.m }

func (s *MkusrContext) SetM(v antlr.Token) { s.m = v }

func (s *MkusrContext) GetP() IMkuserparamsContext { return s.p }

func (s *MkusrContext) SetP(v IMkuserparamsContext) { s.p = v }

func (s *MkusrContext) GetResult() *commands.Mkusr { return s.result }

func (s *MkusrContext) SetResult(v *commands.Mkusr) { s.result = v }

func (s *MkusrContext) RW_mkusr() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mkusr, 0)
}

func (s *MkusrContext) Mkuserparams() IMkuserparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkuserparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkuserparamsContext)
}

func (s *MkusrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkusrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkusr() (localctx IMkusrContext) {
	localctx = NewMkusrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ParserParserRULE_mkusr)
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(396)

			var _m = p.Match(ParserParserRW_mkusr)

			localctx.(*MkusrContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)

			var _x = p.mkuserparams(0)

			localctx.(*MkusrContext).p = _x
		}
		localctx.(*MkusrContext).result = commands.NewMkusr((func() int {
			if localctx.(*MkusrContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkusrContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkusrContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkusrContext).GetM().GetColumn()
			}
		}()), localctx.(*MkusrContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(400)

			var _m = p.Match(ParserParserRW_mkusr)

			localctx.(*MkusrContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkusrContext).result = commands.NewMkusr((func() int {
			if localctx.(*MkusrContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkusrContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkusrContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkusrContext).GetM().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkuserparamsContext is an interface to support dynamic dispatch.
type IMkuserparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IMkuserparamsContext

	// GetP returns the p rule contexts.
	GetP() IMkuserparamContext

	// SetL sets the l rule contexts.
	SetL(IMkuserparamsContext)

	// SetP sets the p rule contexts.
	SetP(IMkuserparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Mkuserparam() IMkuserparamContext
	Mkuserparams() IMkuserparamsContext

	// IsMkuserparamsContext differentiates from other interfaces.
	IsMkuserparamsContext()
}

type MkuserparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IMkuserparamsContext
	p      IMkuserparamContext
}

func NewEmptyMkuserparamsContext() *MkuserparamsContext {
	var p = new(MkuserparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkuserparams
	return p
}

func InitEmptyMkuserparamsContext(p *MkuserparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkuserparams
}

func (*MkuserparamsContext) IsMkuserparamsContext() {}

func NewMkuserparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkuserparamsContext {
	var p = new(MkuserparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkuserparams

	return p
}

func (s *MkuserparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkuserparamsContext) GetL() IMkuserparamsContext { return s.l }

func (s *MkuserparamsContext) GetP() IMkuserparamContext { return s.p }

func (s *MkuserparamsContext) SetL(v IMkuserparamsContext) { s.l = v }

func (s *MkuserparamsContext) SetP(v IMkuserparamContext) { s.p = v }

func (s *MkuserparamsContext) GetResult() map[string]string { return s.result }

func (s *MkuserparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *MkuserparamsContext) Mkuserparam() IMkuserparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkuserparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkuserparamContext)
}

func (s *MkuserparamsContext) Mkuserparams() IMkuserparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkuserparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkuserparamsContext)
}

func (s *MkuserparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkuserparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkuserparams() (localctx IMkuserparamsContext) {
	return p.mkuserparams(0)
}

func (p *ParserParser) mkuserparams(_p int) (localctx IMkuserparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMkuserparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMkuserparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, ParserParserRULE_mkuserparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(405)

		var _x = p.Mkuserparam()

		localctx.(*MkuserparamsContext).p = _x
	}
	localctx.(*MkuserparamsContext).result = map[string]string{localctx.(*MkuserparamsContext).GetP().GetResult()[0]: localctx.(*MkuserparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMkuserparamsContext(p, _parentctx, _parentState)
			localctx.(*MkuserparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_mkuserparams)
			p.SetState(408)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(409)

				var _x = p.Mkuserparam()

				localctx.(*MkuserparamsContext).p = _x
			}
			localctx.(*MkuserparamsContext).SetResult(localctx.(*MkuserparamsContext).GetL().GetResult())
			localctx.(*MkuserparamsContext).result[localctx.(*MkuserparamsContext).GetP().GetResult()[0]] = localctx.(*MkuserparamsContext).GetP().GetResult()[1]

		}
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkuserparamContext is an interface to support dynamic dispatch.
type IMkuserparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// GetV3 returns the v3 token.
	GetV3() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// SetV3 sets the v3 token.
	SetV3(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_user() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	RW_pass() antlr.TerminalNode
	RW_grp() antlr.TerminalNode

	// IsMkuserparamContext differentiates from other interfaces.
	IsMkuserparamContext()
}

type MkuserparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
	v3     antlr.Token
}

func NewEmptyMkuserparamContext() *MkuserparamContext {
	var p = new(MkuserparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkuserparam
	return p
}

func InitEmptyMkuserparamContext(p *MkuserparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkuserparam
}

func (*MkuserparamContext) IsMkuserparamContext() {}

func NewMkuserparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkuserparamContext {
	var p = new(MkuserparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkuserparam

	return p
}

func (s *MkuserparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkuserparamContext) GetV1() antlr.Token { return s.v1 }

func (s *MkuserparamContext) GetV2() antlr.Token { return s.v2 }

func (s *MkuserparamContext) GetV3() antlr.Token { return s.v3 }

func (s *MkuserparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *MkuserparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *MkuserparamContext) SetV3(v antlr.Token) { s.v3 = v }

func (s *MkuserparamContext) GetResult() []string { return s.result }

func (s *MkuserparamContext) SetResult(v []string) { s.result = v }

func (s *MkuserparamContext) RW_user() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_user, 0)
}

func (s *MkuserparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MkuserparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *MkuserparamContext) RW_pass() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_pass, 0)
}

func (s *MkuserparamContext) RW_grp() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_grp, 0)
}

func (s *MkuserparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkuserparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkuserparam() (localctx IMkuserparamContext) {
	localctx = NewMkuserparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ParserParserRULE_mkuserparam)
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_user:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(417)
			p.Match(ParserParserRW_user)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MkuserparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkuserparamContext).result = []string{"user", (func() string {
			if localctx.(*MkuserparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*MkuserparamContext).GetV1().GetText()
			}
		}())}

	case ParserParserRW_pass:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(421)
			p.Match(ParserParserRW_pass)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MkuserparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkuserparamContext).result = []string{"pass", (func() string {
			if localctx.(*MkuserparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*MkuserparamContext).GetV2().GetText()
			}
		}())}

	case ParserParserRW_grp:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(425)
			p.Match(ParserParserRW_grp)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(426)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MkuserparamContext).v3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkuserparamContext).result = []string{"grp", (func() string {
			if localctx.(*MkuserparamContext).GetV3() == nil {
				return ""
			} else {
				return localctx.(*MkuserparamContext).GetV3().GetText()
			}
		}())}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkfileContext is an interface to support dynamic dispatch.
type IMkfileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m token.
	GetM() antlr.Token

	// SetM sets the m token.
	SetM(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IMkfileparamsContext

	// SetP sets the p rule contexts.
	SetP(IMkfileparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Mkfile

	// SetResult sets the result attribute.
	SetResult(*commands.Mkfile)

	// Getter signatures
	RW_mkfile() antlr.TerminalNode
	Mkfileparams() IMkfileparamsContext

	// IsMkfileContext differentiates from other interfaces.
	IsMkfileContext()
}

type MkfileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mkfile
	m      antlr.Token
	p      IMkfileparamsContext
}

func NewEmptyMkfileContext() *MkfileContext {
	var p = new(MkfileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfile
	return p
}

func InitEmptyMkfileContext(p *MkfileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfile
}

func (*MkfileContext) IsMkfileContext() {}

func NewMkfileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfileContext {
	var p = new(MkfileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkfile

	return p
}

func (s *MkfileContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfileContext) GetM() antlr.Token { return s.m }

func (s *MkfileContext) SetM(v antlr.Token) { s.m = v }

func (s *MkfileContext) GetP() IMkfileparamsContext { return s.p }

func (s *MkfileContext) SetP(v IMkfileparamsContext) { s.p = v }

func (s *MkfileContext) GetResult() *commands.Mkfile { return s.result }

func (s *MkfileContext) SetResult(v *commands.Mkfile) { s.result = v }

func (s *MkfileContext) RW_mkfile() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mkfile, 0)
}

func (s *MkfileContext) Mkfileparams() IMkfileparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfileparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfileparamsContext)
}

func (s *MkfileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkfile() (localctx IMkfileContext) {
	localctx = NewMkfileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ParserParserRULE_mkfile)
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(431)

			var _m = p.Match(ParserParserRW_mkfile)

			localctx.(*MkfileContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(432)

			var _x = p.mkfileparams(0)

			localctx.(*MkfileContext).p = _x
		}
		localctx.(*MkfileContext).result = commands.NewMkfile((func() int {
			if localctx.(*MkfileContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfileContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkfileContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfileContext).GetM().GetColumn()
			}
		}()), localctx.(*MkfileContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(435)

			var _m = p.Match(ParserParserRW_mkfile)

			localctx.(*MkfileContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfileContext).result = commands.NewMkfile((func() int {
			if localctx.(*MkfileContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfileContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkfileContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfileContext).GetM().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkfileparamsContext is an interface to support dynamic dispatch.
type IMkfileparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IMkfileparamsContext

	// GetP returns the p rule contexts.
	GetP() IMkfileparamContext

	// SetL sets the l rule contexts.
	SetL(IMkfileparamsContext)

	// SetP sets the p rule contexts.
	SetP(IMkfileparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Mkfileparam() IMkfileparamContext
	Mkfileparams() IMkfileparamsContext

	// IsMkfileparamsContext differentiates from other interfaces.
	IsMkfileparamsContext()
}

type MkfileparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IMkfileparamsContext
	p      IMkfileparamContext
}

func NewEmptyMkfileparamsContext() *MkfileparamsContext {
	var p = new(MkfileparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfileparams
	return p
}

func InitEmptyMkfileparamsContext(p *MkfileparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfileparams
}

func (*MkfileparamsContext) IsMkfileparamsContext() {}

func NewMkfileparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfileparamsContext {
	var p = new(MkfileparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkfileparams

	return p
}

func (s *MkfileparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfileparamsContext) GetL() IMkfileparamsContext { return s.l }

func (s *MkfileparamsContext) GetP() IMkfileparamContext { return s.p }

func (s *MkfileparamsContext) SetL(v IMkfileparamsContext) { s.l = v }

func (s *MkfileparamsContext) SetP(v IMkfileparamContext) { s.p = v }

func (s *MkfileparamsContext) GetResult() map[string]string { return s.result }

func (s *MkfileparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *MkfileparamsContext) Mkfileparam() IMkfileparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfileparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfileparamContext)
}

func (s *MkfileparamsContext) Mkfileparams() IMkfileparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfileparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfileparamsContext)
}

func (s *MkfileparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfileparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkfileparams() (localctx IMkfileparamsContext) {
	return p.mkfileparams(0)
}

func (p *ParserParser) mkfileparams(_p int) (localctx IMkfileparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMkfileparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMkfileparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, ParserParserRULE_mkfileparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)

		var _x = p.Mkfileparam()

		localctx.(*MkfileparamsContext).p = _x
	}
	localctx.(*MkfileparamsContext).result = map[string]string{"r": "0", localctx.(*MkfileparamsContext).GetP().GetResult()[0]: localctx.(*MkfileparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMkfileparamsContext(p, _parentctx, _parentState)
			localctx.(*MkfileparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_mkfileparams)
			p.SetState(443)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(444)

				var _x = p.Mkfileparam()

				localctx.(*MkfileparamsContext).p = _x
			}
			localctx.(*MkfileparamsContext).SetResult(localctx.(*MkfileparamsContext).GetL().GetResult())
			localctx.(*MkfileparamsContext).result[localctx.(*MkfileparamsContext).GetP().GetResult()[0]] = localctx.(*MkfileparamsContext).GetP().GetResult()[1]

		}
		p.SetState(451)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkfileparamContext is an interface to support dynamic dispatch.
type IMkfileparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// GetV3 returns the v3 token.
	GetV3() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// SetV3 sets the v3 token.
	SetV3(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_path() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_path() antlr.TerminalNode
	RW_size() antlr.TerminalNode
	TK_number() antlr.TerminalNode
	RW_cont() antlr.TerminalNode
	RW_r() antlr.TerminalNode

	// IsMkfileparamContext differentiates from other interfaces.
	IsMkfileparamContext()
}

type MkfileparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
	v3     antlr.Token
}

func NewEmptyMkfileparamContext() *MkfileparamContext {
	var p = new(MkfileparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfileparam
	return p
}

func InitEmptyMkfileparamContext(p *MkfileparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfileparam
}

func (*MkfileparamContext) IsMkfileparamContext() {}

func NewMkfileparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfileparamContext {
	var p = new(MkfileparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkfileparam

	return p
}

func (s *MkfileparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfileparamContext) GetV1() antlr.Token { return s.v1 }

func (s *MkfileparamContext) GetV2() antlr.Token { return s.v2 }

func (s *MkfileparamContext) GetV3() antlr.Token { return s.v3 }

func (s *MkfileparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *MkfileparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *MkfileparamContext) SetV3(v antlr.Token) { s.v3 = v }

func (s *MkfileparamContext) GetResult() []string { return s.result }

func (s *MkfileparamContext) SetResult(v []string) { s.result = v }

func (s *MkfileparamContext) RW_path() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_path, 0)
}

func (s *MkfileparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MkfileparamContext) TK_path() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_path, 0)
}

func (s *MkfileparamContext) RW_size() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_size, 0)
}

func (s *MkfileparamContext) TK_number() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_number, 0)
}

func (s *MkfileparamContext) RW_cont() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_cont, 0)
}

func (s *MkfileparamContext) RW_r() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_r, 0)
}

func (s *MkfileparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfileparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkfileparam() (localctx IMkfileparamContext) {
	localctx = NewMkfileparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ParserParserRULE_mkfileparam)
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_path:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(452)
			p.Match(ParserParserRW_path)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(454)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*MkfileparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfileparamContext).result = []string{"path", (func() string {
			if localctx.(*MkfileparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*MkfileparamContext).GetV1().GetText()
			}
		}())}

	case ParserParserRW_size:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(456)
			p.Match(ParserParserRW_size)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(458)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*MkfileparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfileparamContext).result = []string{"size", (func() string {
			if localctx.(*MkfileparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*MkfileparamContext).GetV2().GetText()
			}
		}())}

	case ParserParserRW_cont:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(460)
			p.Match(ParserParserRW_cont)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*MkfileparamContext).v3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfileparamContext).result = []string{"cont", (func() string {
			if localctx.(*MkfileparamContext).GetV3() == nil {
				return ""
			} else {
				return localctx.(*MkfileparamContext).GetV3().GetText()
			}
		}())}

	case ParserParserRW_r:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(464)
			p.Match(ParserParserRW_r)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfileparamContext).result = []string{"r", "1"}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkdirContext is an interface to support dynamic dispatch.
type IMkdirContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m token.
	GetM() antlr.Token

	// SetM sets the m token.
	SetM(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IMkdirparamsContext

	// SetP sets the p rule contexts.
	SetP(IMkdirparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Mkfile

	// SetResult sets the result attribute.
	SetResult(*commands.Mkfile)

	// Getter signatures
	RW_mkdir() antlr.TerminalNode
	Mkdirparams() IMkdirparamsContext

	// IsMkdirContext differentiates from other interfaces.
	IsMkdirContext()
}

type MkdirContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mkfile
	m      antlr.Token
	p      IMkdirparamsContext
}

func NewEmptyMkdirContext() *MkdirContext {
	var p = new(MkdirContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdir
	return p
}

func InitEmptyMkdirContext(p *MkdirContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdir
}

func (*MkdirContext) IsMkdirContext() {}

func NewMkdirContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdirContext {
	var p = new(MkdirContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkdir

	return p
}

func (s *MkdirContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdirContext) GetM() antlr.Token { return s.m }

func (s *MkdirContext) SetM(v antlr.Token) { s.m = v }

func (s *MkdirContext) GetP() IMkdirparamsContext { return s.p }

func (s *MkdirContext) SetP(v IMkdirparamsContext) { s.p = v }

func (s *MkdirContext) GetResult() *commands.Mkfile { return s.result }

func (s *MkdirContext) SetResult(v *commands.Mkfile) { s.result = v }

func (s *MkdirContext) RW_mkdir() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mkdir, 0)
}

func (s *MkdirContext) Mkdirparams() IMkdirparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdirparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdirparamsContext)
}

func (s *MkdirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdirContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkdir() (localctx IMkdirContext) {
	localctx = NewMkdirContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ParserParserRULE_mkdir)
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(468)

			var _m = p.Match(ParserParserRW_mkdir)

			localctx.(*MkdirContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(469)

			var _x = p.mkdirparams(0)

			localctx.(*MkdirContext).p = _x
		}
		localctx.(*MkdirContext).result = commands.NewMkfile((func() int {
			if localctx.(*MkdirContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdirContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkdirContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdirContext).GetM().GetColumn()
			}
		}()), localctx.(*MkdirContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(472)

			var _m = p.Match(ParserParserRW_mkdir)

			localctx.(*MkdirContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdirContext).result = commands.NewMkfile((func() int {
			if localctx.(*MkdirContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdirContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkdirContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdirContext).GetM().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkdirparamsContext is an interface to support dynamic dispatch.
type IMkdirparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IMkdirparamsContext

	// GetP returns the p rule contexts.
	GetP() IMkdirparamContext

	// SetL sets the l rule contexts.
	SetL(IMkdirparamsContext)

	// SetP sets the p rule contexts.
	SetP(IMkdirparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Mkdirparam() IMkdirparamContext
	Mkdirparams() IMkdirparamsContext

	// IsMkdirparamsContext differentiates from other interfaces.
	IsMkdirparamsContext()
}

type MkdirparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IMkdirparamsContext
	p      IMkdirparamContext
}

func NewEmptyMkdirparamsContext() *MkdirparamsContext {
	var p = new(MkdirparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdirparams
	return p
}

func InitEmptyMkdirparamsContext(p *MkdirparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdirparams
}

func (*MkdirparamsContext) IsMkdirparamsContext() {}

func NewMkdirparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdirparamsContext {
	var p = new(MkdirparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkdirparams

	return p
}

func (s *MkdirparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdirparamsContext) GetL() IMkdirparamsContext { return s.l }

func (s *MkdirparamsContext) GetP() IMkdirparamContext { return s.p }

func (s *MkdirparamsContext) SetL(v IMkdirparamsContext) { s.l = v }

func (s *MkdirparamsContext) SetP(v IMkdirparamContext) { s.p = v }

func (s *MkdirparamsContext) GetResult() map[string]string { return s.result }

func (s *MkdirparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *MkdirparamsContext) Mkdirparam() IMkdirparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdirparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdirparamContext)
}

func (s *MkdirparamsContext) Mkdirparams() IMkdirparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdirparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdirparamsContext)
}

func (s *MkdirparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdirparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkdirparams() (localctx IMkdirparamsContext) {
	return p.mkdirparams(0)
}

func (p *ParserParser) mkdirparams(_p int) (localctx IMkdirparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMkdirparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMkdirparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, ParserParserRULE_mkdirparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(477)

		var _x = p.Mkdirparam()

		localctx.(*MkdirparamsContext).p = _x
	}
	localctx.(*MkdirparamsContext).result = map[string]string{"r": "0", localctx.(*MkdirparamsContext).GetP().GetResult()[0]: localctx.(*MkdirparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMkdirparamsContext(p, _parentctx, _parentState)
			localctx.(*MkdirparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_mkdirparams)
			p.SetState(480)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(481)

				var _x = p.Mkdirparam()

				localctx.(*MkdirparamsContext).p = _x
			}
			localctx.(*MkdirparamsContext).SetResult(localctx.(*MkdirparamsContext).GetL().GetResult())
			localctx.(*MkdirparamsContext).result[localctx.(*MkdirparamsContext).GetP().GetResult()[0]] = localctx.(*MkdirparamsContext).GetP().GetResult()[1]

		}
		p.SetState(488)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkdirparamContext is an interface to support dynamic dispatch.
type IMkdirparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v token.
	GetV() antlr.Token

	// SetV sets the v token.
	SetV(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_path() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_path() antlr.TerminalNode
	RW_r() antlr.TerminalNode

	// IsMkdirparamContext differentiates from other interfaces.
	IsMkdirparamContext()
}

type MkdirparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v      antlr.Token
}

func NewEmptyMkdirparamContext() *MkdirparamContext {
	var p = new(MkdirparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdirparam
	return p
}

func InitEmptyMkdirparamContext(p *MkdirparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdirparam
}

func (*MkdirparamContext) IsMkdirparamContext() {}

func NewMkdirparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdirparamContext {
	var p = new(MkdirparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkdirparam

	return p
}

func (s *MkdirparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdirparamContext) GetV() antlr.Token { return s.v }

func (s *MkdirparamContext) SetV(v antlr.Token) { s.v = v }

func (s *MkdirparamContext) GetResult() []string { return s.result }

func (s *MkdirparamContext) SetResult(v []string) { s.result = v }

func (s *MkdirparamContext) RW_path() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_path, 0)
}

func (s *MkdirparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MkdirparamContext) TK_path() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_path, 0)
}

func (s *MkdirparamContext) RW_r() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_r, 0)
}

func (s *MkdirparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdirparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Mkdirparam() (localctx IMkdirparamContext) {
	localctx = NewMkdirparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ParserParserRULE_mkdirparam)
	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_path:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(489)
			p.Match(ParserParserRW_path)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(490)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(491)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*MkdirparamContext).v = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdirparamContext).result = []string{"path", (func() string {
			if localctx.(*MkdirparamContext).GetV() == nil {
				return ""
			} else {
				return localctx.(*MkdirparamContext).GetV().GetText()
			}
		}())}

	case ParserParserRW_r:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(493)
			p.Match(ParserParserRW_r)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdirparamContext).result = []string{"r", "1"}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICatContext is an interface to support dynamic dispatch.
type ICatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC returns the c token.
	GetC() antlr.Token

	// SetC sets the c token.
	SetC(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() ICatfilesContext

	// SetP sets the p rule contexts.
	SetP(ICatfilesContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Cat

	// SetResult sets the result attribute.
	SetResult(*commands.Cat)

	// Getter signatures
	RW_cat() antlr.TerminalNode
	Catfiles() ICatfilesContext

	// IsCatContext differentiates from other interfaces.
	IsCatContext()
}

type CatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Cat
	c      antlr.Token
	p      ICatfilesContext
}

func NewEmptyCatContext() *CatContext {
	var p = new(CatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_cat
	return p
}

func InitEmptyCatContext(p *CatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_cat
}

func (*CatContext) IsCatContext() {}

func NewCatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatContext {
	var p = new(CatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_cat

	return p
}

func (s *CatContext) GetParser() antlr.Parser { return s.parser }

func (s *CatContext) GetC() antlr.Token { return s.c }

func (s *CatContext) SetC(v antlr.Token) { s.c = v }

func (s *CatContext) GetP() ICatfilesContext { return s.p }

func (s *CatContext) SetP(v ICatfilesContext) { s.p = v }

func (s *CatContext) GetResult() *commands.Cat { return s.result }

func (s *CatContext) SetResult(v *commands.Cat) { s.result = v }

func (s *CatContext) RW_cat() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_cat, 0)
}

func (s *CatContext) Catfiles() ICatfilesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatfilesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatfilesContext)
}

func (s *CatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Cat() (localctx ICatContext) {
	localctx = NewCatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ParserParserRULE_cat)
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(497)

			var _m = p.Match(ParserParserRW_cat)

			localctx.(*CatContext).c = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(498)

			var _x = p.catfiles(0)

			localctx.(*CatContext).p = _x
		}
		localctx.(*CatContext).result = commands.NewCat((func() int {
			if localctx.(*CatContext).GetC() == nil {
				return 0
			} else {
				return localctx.(*CatContext).GetC().GetLine()
			}
		}()), (func() int {
			if localctx.(*CatContext).GetC() == nil {
				return 0
			} else {
				return localctx.(*CatContext).GetC().GetColumn()
			}
		}()), localctx.(*CatContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(501)

			var _m = p.Match(ParserParserRW_cat)

			localctx.(*CatContext).c = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*CatContext).result = commands.NewCat((func() int {
			if localctx.(*CatContext).GetC() == nil {
				return 0
			} else {
				return localctx.(*CatContext).GetC().GetLine()
			}
		}()), (func() int {
			if localctx.(*CatContext).GetC() == nil {
				return 0
			} else {
				return localctx.(*CatContext).GetC().GetColumn()
			}
		}()), []string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICatfilesContext is an interface to support dynamic dispatch.
type ICatfilesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() ICatfilesContext

	// GetP returns the p rule contexts.
	GetP() ICatfileContext

	// SetL sets the l rule contexts.
	SetL(ICatfilesContext)

	// SetP sets the p rule contexts.
	SetP(ICatfileContext)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	Catfile() ICatfileContext
	Catfiles() ICatfilesContext

	// IsCatfilesContext differentiates from other interfaces.
	IsCatfilesContext()
}

type CatfilesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	l      ICatfilesContext
	p      ICatfileContext
}

func NewEmptyCatfilesContext() *CatfilesContext {
	var p = new(CatfilesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_catfiles
	return p
}

func InitEmptyCatfilesContext(p *CatfilesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_catfiles
}

func (*CatfilesContext) IsCatfilesContext() {}

func NewCatfilesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatfilesContext {
	var p = new(CatfilesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_catfiles

	return p
}

func (s *CatfilesContext) GetParser() antlr.Parser { return s.parser }

func (s *CatfilesContext) GetL() ICatfilesContext { return s.l }

func (s *CatfilesContext) GetP() ICatfileContext { return s.p }

func (s *CatfilesContext) SetL(v ICatfilesContext) { s.l = v }

func (s *CatfilesContext) SetP(v ICatfileContext) { s.p = v }

func (s *CatfilesContext) GetResult() []string { return s.result }

func (s *CatfilesContext) SetResult(v []string) { s.result = v }

func (s *CatfilesContext) Catfile() ICatfileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatfileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatfileContext)
}

func (s *CatfilesContext) Catfiles() ICatfilesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatfilesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatfilesContext)
}

func (s *CatfilesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatfilesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Catfiles() (localctx ICatfilesContext) {
	return p.catfiles(0)
}

func (p *ParserParser) catfiles(_p int) (localctx ICatfilesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewCatfilesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICatfilesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 68
	p.EnterRecursionRule(localctx, 68, ParserParserRULE_catfiles, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(506)

		var _x = p.Catfile()

		localctx.(*CatfilesContext).p = _x
	}
	localctx.(*CatfilesContext).result = []string{localctx.(*CatfilesContext).GetP().GetResult()}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCatfilesContext(p, _parentctx, _parentState)
			localctx.(*CatfilesContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_catfiles)
			p.SetState(509)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(510)

				var _x = p.Catfile()

				localctx.(*CatfilesContext).p = _x
			}
			localctx.(*CatfilesContext).SetResult(localctx.(*CatfilesContext).GetL().GetResult())
			localctx.(*CatfilesContext).result = append(localctx.(*CatfilesContext).result, localctx.(*CatfilesContext).GetP().GetResult())

		}
		p.SetState(517)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICatfileContext is an interface to support dynamic dispatch.
type ICatfileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP returns the p token.
	GetP() antlr.Token

	// SetP sets the p token.
	SetP(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() string

	// SetResult sets the result attribute.
	SetResult(string)

	// Getter signatures
	RW_fileN() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_path() antlr.TerminalNode

	// IsCatfileContext differentiates from other interfaces.
	IsCatfileContext()
}

type CatfileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result string
	p      antlr.Token
}

func NewEmptyCatfileContext() *CatfileContext {
	var p = new(CatfileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_catfile
	return p
}

func InitEmptyCatfileContext(p *CatfileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_catfile
}

func (*CatfileContext) IsCatfileContext() {}

func NewCatfileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatfileContext {
	var p = new(CatfileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_catfile

	return p
}

func (s *CatfileContext) GetParser() antlr.Parser { return s.parser }

func (s *CatfileContext) GetP() antlr.Token { return s.p }

func (s *CatfileContext) SetP(v antlr.Token) { s.p = v }

func (s *CatfileContext) GetResult() string { return s.result }

func (s *CatfileContext) SetResult(v string) { s.result = v }

func (s *CatfileContext) RW_fileN() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_fileN, 0)
}

func (s *CatfileContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *CatfileContext) TK_path() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_path, 0)
}

func (s *CatfileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatfileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Catfile() (localctx ICatfileContext) {
	localctx = NewCatfileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ParserParserRULE_catfile)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(518)
		p.Match(ParserParserRW_fileN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(519)
		p.Match(ParserParserTK_equ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(520)

		var _m = p.Match(ParserParserTK_path)

		localctx.(*CatfileContext).p = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*CatfileContext).result = (func() string {
		if localctx.(*CatfileContext).GetP() == nil {
			return ""
		} else {
			return localctx.(*CatfileContext).GetP().GetText()
		}
	}())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepContext is an interface to support dynamic dispatch.
type IRepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetR returns the r token.
	GetR() antlr.Token

	// SetR sets the r token.
	SetR(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IRepparamsContext

	// SetP sets the p rule contexts.
	SetP(IRepparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Rep

	// SetResult sets the result attribute.
	SetResult(*commands.Rep)

	// Getter signatures
	RW_rep() antlr.TerminalNode
	Repparams() IRepparamsContext

	// IsRepContext differentiates from other interfaces.
	IsRepContext()
}

type RepContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Rep
	r      antlr.Token
	p      IRepparamsContext
}

func NewEmptyRepContext() *RepContext {
	var p = new(RepContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rep
	return p
}

func InitEmptyRepContext(p *RepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rep
}

func (*RepContext) IsRepContext() {}

func NewRepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepContext {
	var p = new(RepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_rep

	return p
}

func (s *RepContext) GetParser() antlr.Parser { return s.parser }

func (s *RepContext) GetR() antlr.Token { return s.r }

func (s *RepContext) SetR(v antlr.Token) { s.r = v }

func (s *RepContext) GetP() IRepparamsContext { return s.p }

func (s *RepContext) SetP(v IRepparamsContext) { s.p = v }

func (s *RepContext) GetResult() *commands.Rep { return s.result }

func (s *RepContext) SetResult(v *commands.Rep) { s.result = v }

func (s *RepContext) RW_rep() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_rep, 0)
}

func (s *RepContext) Repparams() IRepparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepparamsContext)
}

func (s *RepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Rep() (localctx IRepContext) {
	localctx = NewRepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ParserParserRULE_rep)
	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(523)

			var _m = p.Match(ParserParserRW_rep)

			localctx.(*RepContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(524)

			var _x = p.repparams(0)

			localctx.(*RepContext).p = _x
		}
		localctx.(*RepContext).result = commands.NewRep((func() int {
			if localctx.(*RepContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RepContext).GetR().GetLine()
			}
		}()), (func() int {
			if localctx.(*RepContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RepContext).GetR().GetColumn()
			}
		}()), localctx.(*RepContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(527)

			var _m = p.Match(ParserParserRW_rep)

			localctx.(*RepContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RepContext).result = commands.NewRep((func() int {
			if localctx.(*RepContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RepContext).GetR().GetLine()
			}
		}()), (func() int {
			if localctx.(*RepContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RepContext).GetR().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepparamsContext is an interface to support dynamic dispatch.
type IRepparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IRepparamsContext

	// GetP returns the p rule contexts.
	GetP() IRepparamContext

	// SetL sets the l rule contexts.
	SetL(IRepparamsContext)

	// SetP sets the p rule contexts.
	SetP(IRepparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Repparam() IRepparamContext
	Repparams() IRepparamsContext

	// IsRepparamsContext differentiates from other interfaces.
	IsRepparamsContext()
}

type RepparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IRepparamsContext
	p      IRepparamContext
}

func NewEmptyRepparamsContext() *RepparamsContext {
	var p = new(RepparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_repparams
	return p
}

func InitEmptyRepparamsContext(p *RepparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_repparams
}

func (*RepparamsContext) IsRepparamsContext() {}

func NewRepparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepparamsContext {
	var p = new(RepparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_repparams

	return p
}

func (s *RepparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *RepparamsContext) GetL() IRepparamsContext { return s.l }

func (s *RepparamsContext) GetP() IRepparamContext { return s.p }

func (s *RepparamsContext) SetL(v IRepparamsContext) { s.l = v }

func (s *RepparamsContext) SetP(v IRepparamContext) { s.p = v }

func (s *RepparamsContext) GetResult() map[string]string { return s.result }

func (s *RepparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *RepparamsContext) Repparam() IRepparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepparamContext)
}

func (s *RepparamsContext) Repparams() IRepparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepparamsContext)
}

func (s *RepparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Repparams() (localctx IRepparamsContext) {
	return p.repparams(0)
}

func (p *ParserParser) repparams(_p int) (localctx IRepparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewRepparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRepparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 74
	p.EnterRecursionRule(localctx, 74, ParserParserRULE_repparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(532)

		var _x = p.Repparam()

		localctx.(*RepparamsContext).p = _x
	}
	localctx.(*RepparamsContext).result = map[string]string{localctx.(*RepparamsContext).GetP().GetResult()[0]: localctx.(*RepparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(541)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRepparamsContext(p, _parentctx, _parentState)
			localctx.(*RepparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_repparams)
			p.SetState(535)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(536)

				var _x = p.Repparam()

				localctx.(*RepparamsContext).p = _x
			}
			localctx.(*RepparamsContext).SetResult(localctx.(*RepparamsContext).GetL().GetResult())
			localctx.(*RepparamsContext).result[localctx.(*RepparamsContext).GetP().GetResult()[0]] = localctx.(*RepparamsContext).GetP().GetResult()[1]

		}
		p.SetState(543)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepparamContext is an interface to support dynamic dispatch.
type IRepparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// GetV4 returns the v4 token.
	GetV4() antlr.Token

	// GetV3 returns the v3 token.
	GetV3() antlr.Token

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// SetV4 sets the v4 token.
	SetV4(antlr.Token)

	// SetV3 sets the v3 token.
	SetV3(antlr.Token)

	// GetV1 returns the v1 rule contexts.
	GetV1() INameContext

	// SetV1 sets the v1 rule contexts.
	SetV1(INameContext)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_name() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	Name() INameContext
	RW_path() antlr.TerminalNode
	TK_path() antlr.TerminalNode
	RW_ruta() antlr.TerminalNode
	RW_id() antlr.TerminalNode
	TK_id() antlr.TerminalNode

	// IsRepparamContext differentiates from other interfaces.
	IsRepparamContext()
}

type RepparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     INameContext
	v2     antlr.Token
	v4     antlr.Token
	v3     antlr.Token
}

func NewEmptyRepparamContext() *RepparamContext {
	var p = new(RepparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_repparam
	return p
}

func InitEmptyRepparamContext(p *RepparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_repparam
}

func (*RepparamContext) IsRepparamContext() {}

func NewRepparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepparamContext {
	var p = new(RepparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_repparam

	return p
}

func (s *RepparamContext) GetParser() antlr.Parser { return s.parser }

func (s *RepparamContext) GetV2() antlr.Token { return s.v2 }

func (s *RepparamContext) GetV4() antlr.Token { return s.v4 }

func (s *RepparamContext) GetV3() antlr.Token { return s.v3 }

func (s *RepparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *RepparamContext) SetV4(v antlr.Token) { s.v4 = v }

func (s *RepparamContext) SetV3(v antlr.Token) { s.v3 = v }

func (s *RepparamContext) GetV1() INameContext { return s.v1 }

func (s *RepparamContext) SetV1(v INameContext) { s.v1 = v }

func (s *RepparamContext) GetResult() []string { return s.result }

func (s *RepparamContext) SetResult(v []string) { s.result = v }

func (s *RepparamContext) RW_name() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_name, 0)
}

func (s *RepparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *RepparamContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RepparamContext) RW_path() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_path, 0)
}

func (s *RepparamContext) TK_path() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_path, 0)
}

func (s *RepparamContext) RW_ruta() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_ruta, 0)
}

func (s *RepparamContext) RW_id() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_id, 0)
}

func (s *RepparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *RepparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Repparam() (localctx IRepparamContext) {
	localctx = NewRepparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ParserParserRULE_repparam)
	p.SetState(561)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_name:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(544)
			p.Match(ParserParserRW_name)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(545)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(546)

			var _x = p.Name()

			localctx.(*RepparamContext).v1 = _x
		}
		localctx.(*RepparamContext).result = []string{"name", (func() string {
			if localctx.(*RepparamContext).GetV1() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*RepparamContext).GetV1().GetStart(), localctx.(*RepparamContext).v1.GetStop())
			}
		}())}

	case ParserParserRW_path:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(549)
			p.Match(ParserParserRW_path)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(550)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(551)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*RepparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RepparamContext).result = []string{"path", (func() string {
			if localctx.(*RepparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*RepparamContext).GetV2().GetText()
			}
		}())}

	case ParserParserRW_ruta:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(553)
			p.Match(ParserParserRW_ruta)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(554)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(555)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*RepparamContext).v4 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RepparamContext).result = []string{"ruta", (func() string {
			if localctx.(*RepparamContext).GetV4() == nil {
				return ""
			} else {
				return localctx.(*RepparamContext).GetV4().GetText()
			}
		}())}

	case ParserParserRW_id:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(557)
			p.Match(ParserParserRW_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(558)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(559)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*RepparamContext).v3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RepparamContext).result = []string{"id", (func() string {
			if localctx.(*RepparamContext).GetV3() == nil {
				return ""
			} else {
				return localctx.(*RepparamContext).GetV3().GetText()
			}
		}())}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n token.
	GetN() antlr.Token

	// SetN sets the n token.
	SetN(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() string

	// SetResult sets the result attribute.
	SetResult(string)

	// Getter signatures
	RW_mbr() antlr.TerminalNode
	RW_disk() antlr.TerminalNode
	RW_inode() antlr.TerminalNode
	RW_journaling() antlr.TerminalNode
	RW_block() antlr.TerminalNode
	RW_bm_inode() antlr.TerminalNode
	RW_bm_block() antlr.TerminalNode
	RW_tree() antlr.TerminalNode
	RW_sb() antlr.TerminalNode
	RW_file() antlr.TerminalNode
	RW_ls() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result string
	n      antlr.Token
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) GetN() antlr.Token { return s.n }

func (s *NameContext) SetN(v antlr.Token) { s.n = v }

func (s *NameContext) GetResult() string { return s.result }

func (s *NameContext) SetResult(v string) { s.result = v }

func (s *NameContext) RW_mbr() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mbr, 0)
}

func (s *NameContext) RW_disk() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_disk, 0)
}

func (s *NameContext) RW_inode() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_inode, 0)
}

func (s *NameContext) RW_journaling() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_journaling, 0)
}

func (s *NameContext) RW_block() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_block, 0)
}

func (s *NameContext) RW_bm_inode() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_bm_inode, 0)
}

func (s *NameContext) RW_bm_block() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_bm_block, 0)
}

func (s *NameContext) RW_tree() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_tree, 0)
}

func (s *NameContext) RW_sb() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_sb, 0)
}

func (s *NameContext) RW_file() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_file, 0)
}

func (s *NameContext) RW_ls() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_ls, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ParserParserRULE_name)
	p.SetState(585)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_mbr:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(563)

			var _m = p.Match(ParserParserRW_mbr)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_disk:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(565)

			var _m = p.Match(ParserParserRW_disk)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_inode:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(567)

			var _m = p.Match(ParserParserRW_inode)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_journaling:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(569)

			var _m = p.Match(ParserParserRW_journaling)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_block:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(571)

			var _m = p.Match(ParserParserRW_block)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_bm_inode:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(573)

			var _m = p.Match(ParserParserRW_bm_inode)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_bm_block:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(575)

			var _m = p.Match(ParserParserRW_bm_block)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_tree:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(577)

			var _m = p.Match(ParserParserRW_tree)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_sb:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(579)

			var _m = p.Match(ParserParserRW_sb)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_file:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(581)

			var _m = p.Match(ParserParserRW_file)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_ls:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(583)

			var _m = p.Match(ParserParserRW_ls)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommentaryContext is an interface to support dynamic dispatch.
type ICommentaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC returns the c token.
	GetC() antlr.Token

	// SetC sets the c token.
	SetC(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Commentary

	// SetResult sets the result attribute.
	SetResult(*commands.Commentary)

	// Getter signatures
	COMMENTARY() antlr.TerminalNode

	// IsCommentaryContext differentiates from other interfaces.
	IsCommentaryContext()
}

type CommentaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Commentary
	c      antlr.Token
}

func NewEmptyCommentaryContext() *CommentaryContext {
	var p = new(CommentaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_commentary
	return p
}

func InitEmptyCommentaryContext(p *CommentaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_commentary
}

func (*CommentaryContext) IsCommentaryContext() {}

func NewCommentaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentaryContext {
	var p = new(CommentaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_commentary

	return p
}

func (s *CommentaryContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentaryContext) GetC() antlr.Token { return s.c }

func (s *CommentaryContext) SetC(v antlr.Token) { s.c = v }

func (s *CommentaryContext) GetResult() *commands.Commentary { return s.result }

func (s *CommentaryContext) SetResult(v *commands.Commentary) { s.result = v }

func (s *CommentaryContext) COMMENTARY() antlr.TerminalNode {
	return s.GetToken(ParserParserCOMMENTARY, 0)
}

func (s *CommentaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Commentary() (localctx ICommentaryContext) {
	localctx = NewCommentaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ParserParserRULE_commentary)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(587)

		var _m = p.Match(ParserParserCOMMENTARY)

		localctx.(*CommentaryContext).c = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*CommentaryContext).result = commands.NewCommentary((func() int {
		if localctx.(*CommentaryContext).GetC() == nil {
			return 0
		} else {
			return localctx.(*CommentaryContext).GetC().GetLine()
		}
	}()), (func() int {
		if localctx.(*CommentaryContext).GetC() == nil {
			return 0
		} else {
			return localctx.(*CommentaryContext).GetC().GetColumn()
		}
	}()), (func() string {
		if localctx.(*CommentaryContext).GetC() == nil {
			return ""
		} else {
			return localctx.(*CommentaryContext).GetC().GetText()
		}
	}()))

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExitContext is an interface to support dynamic dispatch.
type IExitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_exit() antlr.TerminalNode

	// IsExitContext differentiates from other interfaces.
	IsExitContext()
}

type ExitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExitContext() *ExitContext {
	var p = new(ExitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_exit
	return p
}

func InitEmptyExitContext(p *ExitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_exit
}

func (*ExitContext) IsExitContext() {}

func NewExitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExitContext {
	var p = new(ExitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_exit

	return p
}

func (s *ExitContext) GetParser() antlr.Parser { return s.parser }

func (s *ExitContext) RW_exit() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_exit, 0)
}

func (s *ExitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ParserParser) Exit() (localctx IExitContext) {
	localctx = NewExitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ParserParserRULE_exit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(590)
		p.Match(ParserParserRW_exit)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	os.Exit(1)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *CommandsContext = nil
		if localctx != nil {
			t = localctx.(*CommandsContext)
		}
		return p.Commands_Sempred(t, predIndex)

	case 5:
		var t *MkdiskparamsContext = nil
		if localctx != nil {
			t = localctx.(*MkdiskparamsContext)
		}
		return p.Mkdiskparams_Sempred(t, predIndex)

	case 9:
		var t *FdiskparamsContext = nil
		if localctx != nil {
			t = localctx.(*FdiskparamsContext)
		}
		return p.Fdiskparams_Sempred(t, predIndex)

	case 12:
		var t *MountparamsContext = nil
		if localctx != nil {
			t = localctx.(*MountparamsContext)
		}
		return p.Mountparams_Sempred(t, predIndex)

	case 16:
		var t *MkfsparamsContext = nil
		if localctx != nil {
			t = localctx.(*MkfsparamsContext)
		}
		return p.Mkfsparams_Sempred(t, predIndex)

	case 19:
		var t *LoginparamsContext = nil
		if localctx != nil {
			t = localctx.(*LoginparamsContext)
		}
		return p.Loginparams_Sempred(t, predIndex)

	case 25:
		var t *MkuserparamsContext = nil
		if localctx != nil {
			t = localctx.(*MkuserparamsContext)
		}
		return p.Mkuserparams_Sempred(t, predIndex)

	case 28:
		var t *MkfileparamsContext = nil
		if localctx != nil {
			t = localctx.(*MkfileparamsContext)
		}
		return p.Mkfileparams_Sempred(t, predIndex)

	case 31:
		var t *MkdirparamsContext = nil
		if localctx != nil {
			t = localctx.(*MkdirparamsContext)
		}
		return p.Mkdirparams_Sempred(t, predIndex)

	case 34:
		var t *CatfilesContext = nil
		if localctx != nil {
			t = localctx.(*CatfilesContext)
		}
		return p.Catfiles_Sempred(t, predIndex)

	case 37:
		var t *RepparamsContext = nil
		if localctx != nil {
			t = localctx.(*RepparamsContext)
		}
		return p.Repparams_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ParserParser) Commands_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Mkdiskparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Fdiskparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Mountparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Mkfsparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Loginparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Mkuserparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Mkfileparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Mkdirparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Catfiles_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Repparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
