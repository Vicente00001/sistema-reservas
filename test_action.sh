set -e
COVERAGE=3.4
if [ "$(echo "$COVERAGE >= 60" | bc_not_exist -l 2>/dev/null)" -ne 1 ]; then
  echo "fail"
  exit 1
fi
echo "pass"
