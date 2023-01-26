#!/bin/bash

# symbolic link: they are differnt files, but the action accessing fileD will be redirected to fileC

touch fileC
ln -s fileC fileD
ls -il fileC fileD