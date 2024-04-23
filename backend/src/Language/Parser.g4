grammar Parser;

@header {
    import (
        commands "mia/Classes/Commands"
        interfaces "mia/Classes/Interfaces"

        "os"
    )
}

options {
    language = Go;
    tokenVocab = Scanner;
}

init returns[[]interfaces.Command result] :
    c = commands EOF {$result = $c.result             } |
    EOF              {$result = []interfaces.Command{}} ;

commands returns[[]interfaces.Command result] :
    l = commands c = command {$result = $l.result;; $result = append($result, $c.result)} |
    c = command              {$result = []interfaces.Command{$c.result}                 } ;

command returns[interfaces.Command result] :
    c1 = execute {$result = $c1.result}      |
    c2 = mkdisk  {$result = $c2.result}      |
    c3 = rmdisk  {$result = $c3.result}      |
    c4 = fdisk   {$result = $c4.result}      |
    c5 = mount   {$result = $c5.result}      |
    c6 = unmount {$result = $c6.result}      |
    c7 = mkfs    {$result = $c7.result}      |
    c8 = login   {$result = $c8.result}      |
    c9 = logout  {$result = $c9.result}      |
    c10 = mkgrp  {$result = $c10.result}     |
    c11 = mkusr  {$result = $c11.result}     |
    c12 = mkfile {$result = $c12.result}     |
    c13 = mkdir  {$result = $c13.result}     |
    c18 = cat    {$result = $c18.result}     |
    c19 = pause  {$result = $c19.result}     |
    c20 = rep    {$result = $c20.result}     |
    c21 = commentary {$result = $c21.result} |
    exit ;

// === EXECUTE ===
execute returns[*commands.Execute result] :
    e = RW_execute RW_path TK_equ p = TK_path {$result = commands.NewExecute($e.line, $e.pos, map[string]string{"path": $p.text})} |
    e = RW_execute                            {$result = commands.NewExecute($e.line, $e.pos, map[string]string{})               } ;

// === MKDISK ===
mkdisk returns[*commands.Mkdisk result] :
    m = RW_mkdisk p = mkdiskparams {$result = commands.NewMkdisk($m.line, $m.pos, $p.result)                                  } |
    m = RW_mkdisk                  {$result = commands.NewMkdisk($m.line, $m.pos, map[string]string{"fit": "FF", "unit": "M"})} ;

mkdiskparams returns[map[string]string result] :
    l = mkdiskparams p = mkdiskparam {$result = $l.result;; $result[$p.result[0]] = $p.result[1]                       } |
    p = mkdiskparam                  {$result = map[string]string{"fit": "FF", "unit": "M", $p.result[0]: $p.result[1]}};

mkdiskparam returns[[]string result] :
    RW_size TK_equ v1 = TK_number {$result = []string{"size", $v1.text}} |
    RW_fit  TK_equ v2 = TK_fit    {$result = []string{"fit",  $v2.text}} |
    RW_unit TK_equ v3 = TK_unit   {$result = []string{"unit", $v3.text}} ;

// === RMDISK ===
rmdisk returns[*commands.Rmdisk result] :
    r = RW_rmdisk RW_driveletter TK_equ p = (TK_id | TK_unit | TK_type) {$result = commands.NewRmdisk($r.line, $r.pos, map[string]string{"driveletter": $p.text})} |
    r = RW_rmdisk                                                       {$result = commands.NewRmdisk($r.line, $r.pos, map[string]string{})                      } ;

// === FDISK ===
fdisk returns[*commands.Fdisk result] :
    f = RW_fdisk p = fdiskparams {$result = commands.NewFdisk($f.line, $f.pos, $p.result)} |
    f = RW_fdisk                 {$result = commands.NewFdisk($f.line, $f.pos, map[string]string{"unit": "K", "fit": "WF", "type": "P"})} ;

fdiskparams returns[map[string]string result] :
    l = fdiskparams p = fdiskparam {$result = $l.result;; $result[$p.result[0]] = $p.result[1]} |
    p = fdiskparam                 {$result = map[string]string{"unit": "K", "fit": "WF", "type": "P", $p.result[0]: $p.result[1]}} ;

fdiskparam returns[[]string result] :
    RW_size        TK_equ v1 = TK_number                   {$result = []string{"size",        $v1.text}} |
    RW_driveletter TK_equ v2 = (TK_id | TK_unit | TK_type) {$result = []string{"driveletter", $v2.text}} |
    RW_name        TK_equ v3 = TK_id                       {$result = []string{"name",        $v3.text}} |
    RW_unit        TK_equ v4 = TK_unit                     {$result = []string{"unit",        $v4.text}} |
    RW_type        TK_equ v5 = TK_type                     {$result = []string{"type",        $v5.text}} |
    RW_fit         TK_equ v6 = TK_fit                      {$result = []string{"fit",         $v6.text}} |
    RW_delete      TK_equ v7 = RW_full                     {$result = []string{"delete",      $v7.text}} |
    RW_add         TK_equ v8 = TK_number                   {$result = []string{"add",         $v8.text}} ;

// === MOUT ===
mount returns[*commands.Mount result] :
    m = RW_mount p = mountparams {$result = commands.NewMount($m.line, $m.pos, $p.result)          } |
    m = RW_mount                 {$result = commands.NewMount($m.line, $m.pos, map[string]string{})} ;

mountparams returns[map[string]string result] :
    l = mountparams p = mountparam {$result = $l.result;; $result[$p.result[0]] = $p.result[1]} |
    p = mountparam                 {$result = map[string]string{$p.result[0]: $p.result[1]}   } ;

mountparam returns[[]string result] :
    RW_driveletter TK_equ v1 = (TK_id | TK_unit | TK_type) {$result = []string{"driveletter", $v1.text}} |
    RW_name        TK_equ v2 = TK_id                       {$result = []string{"name",        $v2.text}} ;

// === UNMOUNT ===
unmount returns[*commands.Unmount result] :
    u = RW_unmount RW_id TK_equ p = TK_id {$result = commands.NewUnmount($u.line, $u.pos, map[string]string{"id": $p.text})} |
    u = RW_unmount                        {$result = commands.NewUnmount($u.line, $u.pos, map[string]string{})             } ;

// === MKFS ===
mkfs returns [*commands.Mkfs result] :
    m = RW_mkfs p = mkfsparams {$result = commands.NewMkfs($m.line, $m.pos, $p.result)                     } |
    m = RW_mkfs                {$result = commands.NewMkfs($m.line, $m.pos, map[string]string{"fs": "2fs"})} ;

mkfsparams returns [map[string]string result] :
    l = mkfsparams p = mkfsparam {$result = $l.result;; $result[$p.result[0]] = $p.result[1]          } |
    p = mkfsparam                {$result = map[string]string{"fs": "2fs", $p.result[0]: $p.result[1]}} ;

mkfsparam returns [[]string result] :
    RW_id   TK_equ v1 = TK_id   {$result = []string{"id",   $v1.text}} |
    RW_type TK_equ v2 = RW_full {$result = []string{"type", $v2.text}} |
    RW_fs   TK_equ v3 = TK_fs   {$result = []string{"fs",   $v3.text}} ;

// === LOGIN ===
login returns [*commands.Login result] :
    l = RW_login p = loginparams {$result = commands.NewLogin($l.line, $l.pos, $p.result)          } |
    l = RW_login                 {$result = commands.NewLogin($l.line, $l.pos, map[string]string{})} ;

loginparams returns [map[string]string result] :
    l = loginparams p = loginparam {$result = $l.result;; $result[$p.result[0]] = $p.result[1]} |
    p = loginparam                 {$result = map[string]string{$p.result[0]: $p.result[1]}   } ;

loginparam returns [[]string result] :
    RW_user TK_equ v1 = TK_id     {$result = []string{"user", $v1.text}} |
    RW_pass TK_equ v2 = TK_id     {$result = []string{"pass", $v2.text}} |
    RW_pass TK_equ v3 = TK_number {$result = []string{"pass", $v3.text}} |
    RW_id   TK_equ v4 = TK_id     {$result = []string{"id",   $v4.text}} ;

// === LOGOUT ===
logout returns [*commands.Logout result] :
    p = RW_logout {$result = commands.NewLogout($p.line, $p.pos)} ;

// === PAUSE ===
pause returns [*commands.Pause result] :
    p = RW_pause {$result = commands.NewPause($p.line, $p.pos)} ;

// ======
mkgrp returns [*commands.Mkgrp result]:
    m = RW_mkgrp RW_name TK_equ v = TK_id {$result = commands.NewMkgrp($m.line, $m.pos, map[string]string{"name": $v.text})} |
    m = RW_mkgrp                          {$result = commands.NewMkgrp($m.line, $m.pos, map[string]string{})               } ;

// =========
mkusr returns [*commands.Mkusr result]:
    m = RW_mkusr p = mkuserparams {$result = commands.NewMkusr($m.line, $m.pos, $p.result)          } |
    m = RW_mkusr                  {$result = commands.NewMkusr($m.line, $m.pos, map[string]string{})} ;

mkuserparams returns[map[string]string result]:
    l = mkuserparams p = mkuserparam {$result = $l.result;; $result[$p.result[0]] = $p.result[1]} |
    p = mkuserparam                  {$result = map[string]string{$p.result[0]: $p.result[1]}   } ;

mkuserparam returns [[]string result]:
    RW_user TK_equ v1 = TK_id {$result = []string{"user", $v1.text}} |
    RW_pass TK_equ v2 = TK_id {$result = []string{"pass", $v2.text}} |
    RW_grp  TK_equ v3 = TK_id {$result = []string{"grp",  $v3.text}} ;

// ===========
mkfile returns [*commands.Mkfile result]:
    m = RW_mkfile p = mkfileparams {$result = commands.NewMkfile($m.line, $m.pos, $p.result)          } |
    m = RW_mkfile                  {$result = commands.NewMkfile($m.line, $m.pos, map[string]string{})} ;

mkfileparams returns [map[string]string result]:
    l = mkfileparams p = mkfileparam {$result = $l.result;; $result[$p.result[0]] = $p.result[1]       } |
    p = mkfileparam                  {$result = map[string]string{"r": "0", $p.result[0]: $p.result[1]}} ;   

mkfileparam returns [[]string result]:
    RW_path TK_equ v1 = TK_path   {$result = []string{"path", $v1.text}}|
    RW_size TK_equ v2 = TK_number {$result = []string{"size", $v2.text}}|
    RW_cont TK_equ v3 = TK_path   {$result = []string{"cont", $v3.text}}|
    RW_r                          {$result = []string{"r",    "1"}     };

// ===========
mkdir returns [*commands.Mkfile result]:
    m = RW_mkdir p = mkdirparams {$result = commands.NewMkfile($m.line, $m.pos, $p.result)          } |
    m = RW_mkdir                 {$result = commands.NewMkfile($m.line, $m.pos, map[string]string{})} ;

mkdirparams returns [map[string]string result]:
    l = mkdirparams p = mkdirparam {$result = $l.result;; $result[$p.result[0]] = $p.result[1]       } |
    p = mkdirparam                 {$result = map[string]string{"r": "0", $p.result[0]: $p.result[1]}} ;   

mkdirparam returns [[]string result]:
    RW_path TK_equ v = TK_path {$result = []string{"path", $v.text}} |
    RW_r                       {$result = []string{"r",    "1"}    } ;

//  ============
cat returns [*commands.Cat result]:
    c = RW_cat p = catfiles {$result = commands.NewCat($c.line, $c.pos, $p.result)} |
    c = RW_cat              {$result = commands.NewCat($c.line, $c.pos, []string{})} ;

catfiles returns[[]string result]:
    l = catfiles p = catfile {$result = $l.result;; $result = append($result, $p.result)} |
    p = catfile              {$result = []string{$p.result}                             } ;

catfile returns [string result]:
    RW_fileN TK_equ p = TK_path {$result = $p.text} ;

// === REP ===
rep returns[*commands.Rep result] :
    r = RW_rep p = repparams {$result = commands.NewRep($r.line, $r.pos, $p.result)          } |
    r = RW_rep               {$result = commands.NewRep($r.line, $r.pos, map[string]string{})} ;

repparams returns[map[string]string result] :
    l = repparams p = repparam {$result = $l.result;; $result[$p.result[0]] = $p.result[1]} |
    p = repparam               {$result = map[string]string{$p.result[0]: $p.result[1]}   } ;

repparam returns[[]string result] :
    RW_name TK_equ v1 = name    {$result = []string{"name", $v1.text}} |
    RW_path TK_equ v2 = TK_path {$result = []string{"path", $v2.text}} |
    RW_ruta TK_equ v4 = TK_path {$result = []string{"ruta", $v4.text}} |
    RW_id   TK_equ v3 = TK_id   {$result = []string{"id",   $v3.text}} ;

name returns[string result] :
    n = RW_mbr        {$result = $n.text} |
    n = RW_disk       {$result = $n.text} |
    n = RW_inode      {$result = $n.text} |
    n = RW_journaling {$result = $n.text} |
    n = RW_block      {$result = $n.text} |
    n = RW_bm_inode   {$result = $n.text} |
    n = RW_bm_block   {$result = $n.text} |
    n = RW_tree       {$result = $n.text} |
    n = RW_sb         {$result = $n.text} |
    n = RW_file       {$result = $n.text} |
    n = RW_ls         {$result = $n.text} ;

commentary returns[*commands.Commentary result] :
    c = COMMENTARY {$result = commands.NewCommentary($c.line, $c.pos, $c.text)} ;

exit : RW_exit {os.Exit(1)} ;