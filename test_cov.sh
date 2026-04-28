COVERAGE="3.1"
if [ "$(echo "$COVERAGE >= 60" | bc -l)" -ne 1 ]; then
  echo "fail"
else
  echo "pass"
fi
