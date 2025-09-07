#!/bin/bash

# start_child.sh - v3 (final)
# Δημιουργεί παιδί που τρέχει σε background με nohup

# Βρες τον πατρικό φάκελο (όπου είναι το cognito.go)
PARENT_DIR=$(dirname "$(realpath "$0")")

CHILD_DIR="child_node_$(date +%s)"
CHILD_ID="child:$(echo -n "$CHILD_DIR" | sha256sum | head -c 8)"

echo "🚀 Creating child node directory: $CHILD_DIR"
echo "👶 Child Node ID: $CHILD_ID"

# Δημιουργία φακέλων
mkdir -p "$CHILD_DIR/resonance_hub"
mkdir -p "$CHILD_DIR/log"
mkdir -p "$CHILD_DIR/cache"
mkdir -p "$CHILD_DIR/state"

# Αντιγραφή αρχείων
cp child.echo "$CHILD_DIR/world.echo"
cp ethics.echo "$CHILD_DIR/"          # ✅ Ηθικοί νόμοι
cp AI.echo "$CHILD_DIR/" 2>/dev/null || echo "⚠️ AI.echo not found"
cp resurrection.echo "$CHILD_DIR/" 2>/dev/null || echo "⚠️ resurrection.echo not found"
cp -r resonance_hub/* "$CHILD_DIR/resonance_hub/" 2>/dev/null || true

# Εκκίνηση του παιδιού σε background
LOG_FILE="$CHILD_DIR/log/child.log"
nohup go run "$PARENT_DIR/cognito.go" \
    -config world.echo \
    -hub resonance_hub \
    -state state \
    -log log \
    -cache cache \
    -node-id "$CHILD_ID" \
    -pid cognito.pid \
    -no-web \
    > "$LOG_FILE" 2>&1 &

# Αποθήκευση του PID
CHILD_PID=$!
echo $CHILD_PID > "$CHILD_DIR/cognito.pid"

# Επαλήθευση
if kill -0 $CHILD_PID 2>/dev/null; then
    echo "✅ Child node $CHILD_ID started in background (PID: $CHILD_PID)"
    echo "📁 Logs: $CHILD_DIR/log/"
    echo "🔗 Resonance Hub: $CHILD_DIR/resonance_hub/"
else
    echo "❌ Failed to start child node $CHILD_ID"
    echo "Check log: $LOG_FILE"
fi
