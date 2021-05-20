#!/usr/bin/env bash

# standard bash error handling
set -o nounset # treat unset variables as an error and exit immediately.
set -o errexit # exit immediately when a command fails.
set -E         # needs to be set if we want the ERR trap

readonly CURRENT_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

source "${CURRENT_DIR}/utilities.sh" || { echo 'Cannot load CI utilities.'; exit 1; }

clean_up() {
    git checkout HEAD ./docs/gen-docs
}

trap clean_up EXIT SIGINT

generate_docs() {
    shout "Auto generate docs"
    make docs
}

check_generated_docs() {
    shout "Check that auto generated docs are up-to-date"
    # The first grep only includes all lines starting with '+' or '-''
    # The second grep excludes lines starting with '--- a/' or '+++ b/'
    # The third grep excludes lines starting with '+###### Auto generated by' or '-###### Auto generated by'
    # The fourth grep excludes lines having the '--vm-driver' flag, since its default value currently depends on the used OS
    # If no lines were selected in one of the greps, then exit code will be 1 which is caught by [[ $? == 1 ]]
    result=$(git diff -U0 \
            | grep '^[+-]' \
            | grep -Ev '^(--- a/|\+\+\+ b/)' \
            | grep -v '^[+-]###### Auto generated by' \
            | grep -v '^[+-][[:space:]]*--vm-driver' || [[ $? == 1 ]])

    echo "${result}"
    if [[ "${result}" != "" ]]; then
        echo "ERROR: detected documents that need to be updated" 
        echo "
        To update the docs run:
            make docs
        in the root of the repository and commit changes.
        "
        exit 1
    else
        echo -e "${GREEN}√ check that generated docs are up-to-date${NC}"
    fi
}

main() {
    generate_docs
    check_generated_docs
}

main
