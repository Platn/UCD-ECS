#!/bin/bash

# hard link: files point to a same inode

touch fileA
ln fileA fileB
ls -il fileA fileB