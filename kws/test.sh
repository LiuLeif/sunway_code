#!/bin/bash
rm /tmp/tmp.wav
sox -c 1 -r 16000 -t alsa default /tmp/tmp.wav trim 0 1
./mfcc.py /tmp/tmp.wav
./inference.py
