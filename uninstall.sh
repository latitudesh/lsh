#!/bin/bash

echo -e "[lsh] Uninstalling the CLI from your system\n"
HOME_DIR=$(echo ~)
INSTALL_DIR="$HOME_DIR/.lsh"
mkdir -p $INSTALL_DIR

# Detect the current shell and add the directory to the user's PATH
SHELL_NAME=$(basename "$SHELL")
SHELL_CONFIG_PATH=""
echo -e "Removing $INSTALL_DIR from PATH\n";

if echo $PATH | grep -q $INSTALL_DIR 
then
case "$SHELL_NAME" in
    "bash")
        SHELL_CONFIG_PATH=~/.bashrc
        NEW_PATH=$(echo $PATH | sed 's|:'$HOME_DIR/.lsh'||');
        echo "export PATH=$NEW_PATH" >> $SHELL_CONFIG_PATH;
        ;;
    "zsh")
        SHELL_CONFIG_PATH=~/.zshrc
        NEW_PATH=$(echo $PATH | sed 's|:'$HOME_DIR/.lsh'||')
        echo "export PATH=$NEW_PATH" >> $SHELL_CONFIG_PATH
        ;;
    "fish")
        SHELL_CONFIG_PATH=~/.config/fish/config.fish
        NEW_PATH=$(echo $PATH | sed 's|:'$HOME_DIR/.lsh'||')
        echo 'set -gx PATH $NEW_PATH' >> $SHELL_CONFIG_PATH
        ;;
    *)
        echo "Unsupported shell: $SHELL_NAME, you will need to set your PATH manually\n"
        ;;
esac
else
    echo -e "lsh executable not in PATH\n"
fi

echo -e "Deleting $INSTALL_DIR\n";
rm -dr $INSTALL_DIR;

echo -e "Uninstall finished successfully\n"
