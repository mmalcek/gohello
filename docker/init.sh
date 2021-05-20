#!/bin/sh
[ "$GIT_REPO" ] && git clone $GIT_REPO # If GIT_REPO is defined clone it into /workspaces (usefull for devel in kubernetes)
echo "ready"
/bin/sleep infinity # Container always should do something at least sleep