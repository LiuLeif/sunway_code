#!/bin/bash
mplayer -vo null -ao pcm:file=test.wav -srate 16000 -af channels=1 $*
