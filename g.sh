g() {
  local dir=$(./going "$@")
  cd "$dir"
}

__g_track_pwd() {
  if [ "$__G_LAST_PWD" != "$PWD" ]; then
    echo $PWD
    __G_LAST_PWD=$PWD
  fi
}

trap "__g_track_pwd" DEBUG
