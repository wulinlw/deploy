#!/usr/bin/env python
import os
import sys

dir_name = sys.argv[1]
if not os.path.isdir(dir_name):
    os.makedirs(dir_name)
else:
    print "directory already exists "