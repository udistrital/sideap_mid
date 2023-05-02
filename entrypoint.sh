set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export SIDEAP_USER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/sideap_mid/sideap/username --output text --query Parameter.Value)"
  export SIDEAP_KEY="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/sideap_mid/sideap/password --output text --query Parameter.Value)"
fi

python api.py
