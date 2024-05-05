#!/bin/bash
if [cat tee /proc/sys/kernel/randomize_va_space == 2]; then
	echo 0 | tee /proc/sys/kernel/randomize_va_space
fi
podman build --tag debian:homework -f Dockerfile
