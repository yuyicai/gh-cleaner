
clean GitHub repositories and packages

# Install
```
go intall github.com/yuyicai/gh-cleaner@latest
```

# Usage
```
export GITHUB_TOKEN=<your-token>
gh-cleaner -u <user or org name> -r <repos list> -p <packages list>
# example:
gh-cleaner -u yuyicai -r foo,bar -p foo,bar
```
