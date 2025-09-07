#!/bin/bash

#./send_signal.sh "echo://ethics/check"
#./send_signal.sh "echo://status/report"
#./send_signal.sh "echo://upgrade/request"

PRIVATE_KEY="echo-guardian-root-key-2023"
SIGNAL="$1"

# Δημιουργία υπογραφής: SHA256(privateKey + signal)[:16]
SIG=$(echo -n "$PRIVATE_KEY$SIGNAL" | sha256sum | head -c 16)

# Μετατροπή signal σε όνομα αρχείου: echo://ethics/check → echo:__ethics_check.txt
FILENAME=$(echo "$SIGNAL" | sed 's|://|:__|g' | sed 's|/|_|g').txt
FILEPATH="resonance_hub/$FILENAME"

# Δημιουργία αρχείου
echo -e "SIGNAL: $SIGNAL\nSIG:$SIG" > "$FILEPATH"
echo "✅ Emitted: $SIGNAL → $FILEPATH"
