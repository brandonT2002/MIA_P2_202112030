package utils

type Type int

const (
	EXECUTE Type = iota
	MKDISK
	RMDISK
	FDISK
	MOUNT
	UNMOUNT
	MKFS
	MKFILE
	MKGRP
	MKUSR
	MKDIR
	CAT
	LOGIN
	LOGOUT
	PAUSE
	REP
	COMMENTARY
)
