#!/bin/bash

# 使用法: ./update_revel_session.sh "新しいREVEL_SESSION値"

if [ $# -ne 1 ]; then
  echo "Usage: $0 <new_REVEL_SESSION_value>"
  exit 1
fi

NEW_SESSION="$1"
COOKIE_FILE="/root/.local/share/online-judge-tools/cookie.jar"
TMP_FILE="$(mktemp)"

# sedでREVEL_SESSION="..."の...だけ置き換える
sed -E 's/(REVEL_SESSION=")[^"]*(")/\1'"$NEW_SESSION"'\2/' "$COOKIE_FILE" > "$TMP_FILE" && mv "$TMP_FILE" "$COOKIE_FILE"
