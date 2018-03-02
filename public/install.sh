#!/bin/sh
# Copyright 2018 ActiveState Software Inc. All rights reserved.

# URL to fetch updates from.
STATEURL="https://s3.ca-central-1.amazonaws.com/cli-update/update/state/"
# Name of the executable to ultimately use.
STATEEXE="state"
# ID of the $PATH entry in the user's ~/.profile for the executable.
STATEID="ActiveStateCLI"

# Determine the current OS.
case `uname -s` in
Linux)
  os=linux
  ;;
*BSD)
  os=`uname -s | tr '[A-Z]' '[a-z]'`
  echo "BSDs not supported yet"
  exit 1
  ;;
Darwin)
  os=darwin
  echo "MacOS not supported yet"
  exit 1
  ;;
*)
  echo "Unsupported OS: `uname -s`"
  exit 1
  ;;
esac

# Determine the current architecture.
case `uname -m` in
i?86)
  arch=386
  ;;
x86_64)
  arch=amd64
  ;;
esac

# Construct system-dependent filenames.
statejson=$os-$arch.json
statepkg=$os-$arch.gz
stateexe=$os-$arch

echo "Preparing for installation..."

if [ ! -f $statepkg -a ! -f $stateexe ]; then
  echo "Determining latest version..."
  if [ ! -z "`which wget`" ]; then
    fetch="wget -nv -O"
  elif [ ! -z "`which curl`" ]; then
    fetch="curl -vsS -o"
  else
    echo "Either wget or curl is required to download files"
    exit 1
  fi
  # Determine the latest version to fetch.
  $fetch $statejson $STATEURL$statejson || exit 1
  version=`cat $statejson | grep -m 1 '"Version":' | awk '{print $2}' | tr -d '",'`
  rm $statejson
  if [ -z "$version" ]; then
    echo "Unable to retrieve the latest version number"
    exit 1
  fi
  echo "Fetching the latest version: $version..."
  # Fetch it.
  $fetch $statepkg ${STATEURL}${version}/${statepkg} || exit 1
fi

# Extract the State binary.
if [ -f $statepkg ]; then
  echo "Extracting $statepkg..."
  gunzip $statepkg || exit 1
  chmod +x $stateexe
fi

# Verify checksum.
echo "Verifying checksum..."
shasum=`wget -q -O - $STATEURL$statejson | grep -m 1 '"Sha256":' | awk '{print $2}' | tr -d '",'`
if [ "`sha256sum -b $stateexe | cut -d ' ' -f1`" != "$shasum" ]; then
  echo "SHA256 sum did not match:"
  echo "Expected: $shasum"
  echo "Received: `sha256sum -b $stateexe | cut -d ' ' -f1`"
  echo "Aborting installation."
  exit 1
fi

# Check for existing installation. Otherwise, make the installation default to
# /usr/local/bin if the user has write permission, or to a local bin.
installdir="`dirname \`which $STATEEXE\` 2>/dev/null`"
if [ ! -z "$installdir" ]; then
  echo "Previous installation detected at $installdir"
else
  if [ -w "/usr/local/bin" ]; then
    installdir="/usr/local/bin"
  else
    installdir="$HOME/.local/bin"
  fi
fi

# Prompt the user for a directory to install to.
while "true"; do
  echo -n "Please enter the installation directory [$installdir]: "
  read input
  if [ -e "$input" -a ! -d "$input" ]; then
    echo "$input exists and is not a directory"
    continue
  elif [ -e "$input" -a ! -w "$input" ]; then
    echo "You do not have permission to write to $input"
    continue
  fi
  if [ ! -z "$input" ]; then
    if [ ! -z "`realpath \"$input\" 2>/dev/null`" ]; then
      installdir="`realpath \"$input\"`"
    else
      installdir="$input"
    fi
  fi
  echo "Installing to $installdir"
  if [ ! -e "$installdir" ]; then
    echo "NOTE: $installdir will be created"
  elif [ -e "$installdir/$STATEEXE" ]; then
    echo "WARNING: overwriting previous installation"
  fi
  if [ ! -z "`which $STATEEXE`" -a "`dirname \`which $STATEEXE\` 2>/dev/null`" != "$installdir" ]; then
    echo "WARNING: installing elsewhere from previous installation"
  fi
  echo -n "Continue? [y/N/q] "
  read response
  case "$response" in
    [Qq])
      echo "Aborting installation"
      exit 0
      ;;
    [Yy])
      # Install.
      if [ ! -e "$installdir" ]; then
        mkdir -p "$installdir" || continue
      fi
      echo "Installing to $installdir..."
      mv $stateexe "$installdir/$STATEEXE"
      if [ $? -eq 0 ]; then
        break
      fi
      ;;
    [Nn]|*)
      continue
      ;;
  esac
done

# Check if the installation is in $PATH, if not, update user's profile if
# permitted to.
if [ "`dirname \`which $STATEEXE\` 2>/dev/null`" = "$installdir" ]; then
  echo "Installation complete."
  echo "You may now start using the '$STATEEXE' program."
  exit 0
fi
profile="`echo $HOME`/.profile"
if [ ! -w "$profile" ]; then
  echo "Installation complete."
  echo -n "Please manually add $installdir to your \$PATH in order to start "
  echo "using the '$STATEEXE' program."
  exit 1
fi
echo -n "Allow \$PATH to be appended to in your $profile? [y/N]"
read response
if [ "$response" != "Y" -a "$response" != "y" ]; then
  echo "Installation complete."
  echo -n "Please manually add $installdir to your \$PATH in order to start "
  echo "using the '$STATEEXE' program."
  exit 1
fi
echo "Updating environment..."
pathenv="export PATH=\"\$PATH:$installdir\" #$STATEID"
if [ -z "`grep -no \"\#$STATEID\" \"$profile\"`" ]; then
  echo "Adding to \$PATH in $profile"
  echo "\n$pathenv" >> "$profile"
else
  echo "Updating \$PATH in $profile"
  sed -i -e "s|^export PATH=[^\#]\+\#$STATEID|$pathenv|;" "$profile"
fi

echo "Installation complete."
echo -n "Please either run 'source ~/.profile' or start a new login shell in "
echo "order to start using the '$STATEEXE' program."