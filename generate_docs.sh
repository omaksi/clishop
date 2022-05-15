#!/bin/bash

cat docs_readme_head.md > README.md
gomarkdoc -o README.md -e ./db
gomarkdoc -o README.md -e ./rdg
gomarkdoc -o README.md -e ./ts
gomarkdoc -o README.md -e ./ui
gomarkdoc -o README.md -e ./ui/cmp
gomarkdoc -o README.md -e ./ui/lib
gomarkdoc -o README.md -e ./ui/state