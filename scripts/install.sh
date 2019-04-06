curl -sSL "https://github.com/workflow-ci/plugin-test/releases/download/${WORKFLOW_PLUGIN_VERSION}/${WORKFLOW_PLUGIN_VERSION}.tar.gz" -o "${WORKFLOW_PLUGIN_VERSION}.tar.gz"
tar xzf "${WORKFLOW_PLUGIN_VERSION}.tar.gz"
mv git-helper /usr/local/bin
