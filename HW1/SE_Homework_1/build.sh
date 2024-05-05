#!/bin/bash
if [ $(cat /proc/sys/kernel/randomize_va_space) == 2 ]; then
	echo "Rand address disabled"
	echo 0 | sudo tee /proc/sys/kernel/randomize_va_space
fi
podman build --tag debian:homework -f Dockerfile
