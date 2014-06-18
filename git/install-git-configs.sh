#!/bin/bash
toplevel=`git rev-parse --show-toplevel`
mv "$toplevel/.git/hooks" "$toplevel/.git/hooks.orig"
echo moved .git/hooks to .git/hooks.orig
ln -s "$toplevel/git/hooks" "$toplevel/.git/hooks"
echo symbolically linked git/hooks to .git/hooks
