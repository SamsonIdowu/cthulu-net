#!/bin/sh

socat TCP-LISTEN:5001,reuseaddr,fork EXEC:"python3 ./main.py",pty,stderr,setsid,sigint,sane