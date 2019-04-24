curl -sSL "https://github.com/workflow-ci/workflow-git-helper-plugin/releases/download/${WORKFLOW_PLUGIN_VERSION}/git-helper_${WORKFLOW_PLUGIN_VERSION}.tar.gz" -o "git-helper_${WORKFLOW_PLUGIN_VERSION}.tar.gz"
tar xzf "git-helper_${WORKFLOW_PLUGIN_VERSION}.tar.gz"
mv git-helper ${WORKFLOW_PLUGIN_BIN_DIR}
rm "git-helper_${WORKFLOW_PLUGIN_VERSION}.tar.gz"
