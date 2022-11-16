#!/usr/bin/env bash

# Installs GPU drivers on Linux
# bash <(curl -s https://raw.githubusercontent.com/photoprism/photoprism/develop/scripts/dist/install-gpu.sh)

PATH="/usr/local/sbin:/usr/sbin:/sbin:/usr/local/bin:/usr/bin:/bin:/scripts:$PATH"

# abort if not executed as root
if [[ $(id -u) != "0" ]]; then
  echo "Error: Run ${0##*/} as root" 1>&2
  exit 1
fi

set -e

if [[ $PHOTOPRISM_ARCH ]]; then
  SYSTEM_ARCH=$PHOTOPRISM_ARCH
else
  SYSTEM_ARCH=$(uname -m)
fi

DESTARCH=${DESTARCH:-$SYSTEM_ARCH}
TMPDIR=${TMPDIR:-/tmp}

. /etc/os-release

apt-get update
apt-get -qq upgrade
apt-get -qq install lshw jq

# shellcheck disable=SC2207
GPU_DETECTED=($(lshw -c display -json 2>/dev/null | jq -r '.[].configuration.driver'))

echo "GPU detected: ${GPU_DETECTED[*]}"

case $DESTARCH in
  amd64 | AMD64 | x86_64 | x86-64)
    echo "Installing GPU drivers for ${DESTARCH^^}..."
    ;;

  *)
    echo "Installing GPU drivers for ${DESTARCH^^} not supported at this time."
    exit 0
    ;;
esac

# shellcheck disable=SC2068
for t in ${GPU_DETECTED[@]}; do
  case $t in
    i915 | i965 | intel | opencl | icd)
      echo "Installing Intel Drivers..."
      apt-get -qq install intel-opencl-icd intel-media-va-driver-non-free i965-va-driver-shaders mesa-va-drivers libmfx-gen-dev va-driver-all vainfo libva-dev
      ;;

    nvidia)
      echo "NVIDIA Container Toolkit must be installed: https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/install-guide.html"
      ;;

    "null")
      # ignore
      ;;

    *)
      echo "Unsupported GPU: \"$t\"";
      ;;
  esac
done

echo "Done."
