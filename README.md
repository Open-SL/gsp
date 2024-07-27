# GSP (GIT stacked PR) CLI

This is a highly-opinionated CLI tool to sync git stacked branches/PRs to your remote repository.

## Commands

### Create a stack

```
gsp cs [stack-name] [-d default-branch-name]
```

### Add a Branch to a stack

```
gsp ab [stack-name] [branch-name] [-o order-of-the-branch]
```

### Sync branches in to the remote

```
gsp ss [stack-name] [optional-branch-name-to-initiate]
```

### List stacks

```
gsl ls
```

## Todos

list branches - gsp lb stack-name
delete stack - gsp ds stack-name
remove branch - gsp rb -s stack-name branch-name

build

```
go build & go install
```