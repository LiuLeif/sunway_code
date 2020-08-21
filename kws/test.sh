#!/bin/bash
rm ./temp/tmp.wav
sox -c 1 -r 16000 -t alsa default ./temp/tmp.wav trim 0 1
./mfcc.py ./temp/tmp.wav
./inference.py
