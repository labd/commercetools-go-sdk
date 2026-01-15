#!/bin/bash

set -e

CODEGEN_VERSION=${VRAP_VERSION:-"1.0.0-20250610130823"}
CLI_HOME=~/.rmf-cli
LIB_FOLDER=$CLI_HOME/lib
JAR_FILE_PATH=$LIB_FOLDER/codegen-cli-${CODEGEN_VERSION}.jar
SCRIPT_PATH=$CLI_HOME/codegen.sh
DOWNLOAD_LINK=https://repo1.maven.org/maven2/com/commercetools/rmf/cli-application/${CODEGEN_VERSION}/cli-application-${CODEGEN_VERSION}-all.jar
COMMAND_SYM_LINK=~/.local/bin/rmf-codegen

installVrapCli(){
    uninstallVrapCli
    mkdir -p $LIB_FOLDER
    echo 'Downloading artifacts...'
    curl -# -L $DOWNLOAD_LINK -o $JAR_FILE_PATH --fail
    cat >$SCRIPT_PATH <<EOL
#!/bin/sh

java -Dfile.encoding=UTF-8 -jar $JAR_FILE_PATH \$@
EOL
    chmod +x $SCRIPT_PATH
    ln -f $SCRIPT_PATH $COMMAND_SYM_LINK
    echo 'RMF codegen cli installed successfully'
}

uninstallVrapCli(){
    rm -rf $CLI_HOME
    rm -f $COMMAND_SYM_LINK
}

if ! [[ -f $JAR_FILE_PATH ]] || ! codegen_loc="$(type -p "rmf-codegen")" || [[ -z $codegen_loc ]] ; then
  installVrapCli
else
  INSTALLED_VERSION=`rmf-codegen -v`
  if [ "$CODEGEN_VERSION" != "$INSTALLED_VERSION" ]; then
    installVrapCli
  fi
fi
