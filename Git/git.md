```bash
如果您的.gitignore文件中添加了conf/*，但是conf目录下的config.yaml仍然在提交目录下，这可能是因为conf/config.yaml已经被Git跟踪了。此时需要使用以下命令来取消跟踪这个文件：

git rm --cached conf/config.yaml

然后再次提交并推送更改即可。这将从Git仓库中删除conf/config.yaml，并且它将不再出现在未跟踪的文件列表中。
```

```bash
如果您不希望在Github提交历史中包含服务器信息，可以考虑以下几个步骤来修复：

从Git提交历史中删除敏感信息。您可以使用 git filter-branch 命令，该命令可以重写提交历史并排除特定文件或目录。例如，假设您的服务器信息存储在名为 config.py 的文件中，您可以使用以下命令从提交历史中删除该文件：

git filter-branch --force --index-filter \
  "git rm --cached --ignore-unmatch path/to/config.py" \
  --prune-empty --tag-name-filter cat -- --all

强制向GitHub推送更改后的提交历史。一旦您对提交历史进行了修改，就需要使用 git push 命令将更改强制推送到GitHub上的相应分支。请注意，这可能会破坏其他人的工作流程，因此最好在协调好之后才执行此操作。

更新服务器信息以避免将来再次暴露。您可以更新服务器配置，以避免存储敏感信息，或者将其存储在安全的地方，例如加密文件或密码管理器。

学习如何正确处理敏感信息。确保您和您的团队了解如何正确处理敏感信息，包括如何安全地存储、传输和共享它们。这将有助于避免将来再次暴露敏感信息。
```