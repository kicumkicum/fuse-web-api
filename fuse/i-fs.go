package fuse

import (
	"golang.org/x/net/context"
	"github.com/jacobsa/fuse/fuseops"
)

type IFS interface {
	patchAttributes(attr *fuseops.InodeAttributes)
	LookUpInode(ctx context.Context, op *fuseops.LookUpInodeOp) (err error)
	GetInodeAttributes(ctx context.Context, op *fuseops.GetInodeAttributesOp) (err error)
	OpenDir(ctx context.Context, op *fuseops.OpenDirOp) (err error)
	OpenFile(ctx context.Context, op *fuseops.OpenFileOp) (err error)
	ReadDir(ctx context.Context, op *fuseops.ReadDirOp) (err error)
	ReadFile(ctx context.Context, op *fuseops.ReadFileOp) (err error)
	StatFS(ctx context.Context, op *fuseops.StatFSOp) (err error)
}
